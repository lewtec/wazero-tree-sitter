package main

import (
	"archive/tar"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// tree-sitter core lives in the workspaced github source cache:
//
//	~/.cache/workspaced/sources/github/sha256("v4:repo:tree-sitter/tree-sitter@HEAD")
var treeSitterSource = workspacedGithubSource{
	Repo:        "tree-sitter/tree-sitter",
	Version:     "HEAD",
	Marker:      "lib/src/lib.c",
	EnvOverride: "TREE_SITTER_PATH",
}

type workspacedGithubSource struct {
	Repo        string
	Version     string
	Marker      string
	EnvOverride string
}

func (s workspacedGithubSource) version() string {
	v := strings.TrimSpace(s.Version)
	if v == "" {
		return "HEAD"
	}
	return v
}

func (s workspacedGithubSource) repo() string {
	return strings.Trim(strings.TrimSpace(s.Repo), "/")
}

func (s workspacedGithubSource) cacheKey() string {
	return "v4:repo:" + s.repo() + "@" + s.version()
}

func (s workspacedGithubSource) CachePath() (string, error) {
	sum := sha256.Sum256([]byte(s.cacheKey()))
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".cache", "workspaced", "sources", "github", hex.EncodeToString(sum[:])), nil
}

func resolveTreeSitterPath() (string, error) {
	return treeSitterSource.Resolve()
}

func (s workspacedGithubSource) Resolve() (string, error) {
	if s.EnvOverride != "" {
		if p := strings.TrimSpace(os.Getenv(s.EnvOverride)); p != "" {
			if err := s.checkReady(p); err != nil {
				return "", err
			}
			return p, nil
		}
	}
	return s.ensure()
}

func (s workspacedGithubSource) checkReady(root string) error {
	mark := filepath.Join(root, s.Marker)
	if _, err := os.Stat(mark); err != nil {
		return fmt.Errorf("%s root %s: %w (need %s)", s.repo(), root, err, s.Marker)
	}
	return nil
}

func (s workspacedGithubSource) ensure() (string, error) {
	cache, err := s.CachePath()
	if err != nil {
		return "", err
	}
	if err := s.checkReady(cache); err == nil {
		return cache, nil
	}
	digest, err := lockDigestForGithubSource(s.repo())
	if err != nil {
		return "", fmt.Errorf("%s cache miss at %s: %w\nrun: mise run grammars:lock", s.repo(), cache, err)
	}
	url := fmt.Sprintf("https://codeload.github.com/%s/tar.gz/%s", s.repo(), digest)
	if err := fetchGithubTarball(url, cache); err != nil {
		return "", err
	}
	if err := s.checkReady(cache); err != nil {
		return "", err
	}
	return cache, nil
}

func lockDigestForGithubSource(repo string) (string, error) {
	lockPath, err := findUp("workspaced.lock.json")
	if err != nil {
		return "", err
	}
	b, err := os.ReadFile(lockPath)
	if err != nil {
		return "", err
	}
	var lock struct {
		Dependencies []struct {
			Kind          string `json:"kind"`
			Ref           string `json:"ref"`
			DepName       string `json:"depName"`
			CurrentDigest string `json:"currentDigest"`
		} `json:"dependencies"`
	}
	if err := json.Unmarshal(b, &lock); err != nil {
		return "", err
	}
	want := "github:" + repo
	for _, d := range lock.Dependencies {
		if d.Kind != "source" {
			continue
		}
		if d.Ref == want || d.DepName == repo {
			if dig := strings.TrimSpace(d.CurrentDigest); dig != "" {
				return dig, nil
			}
		}
	}
	return "", fmt.Errorf("no locked source %s in %s", want, lockPath)
}

func fetchGithubTarball(url, dest string) error {
	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return err
	}
	tmp := dest + ".tmp"
	_ = os.RemoveAll(tmp)
	if err := os.MkdirAll(tmp, 0o755); err != nil {
		return err
	}
	defer os.RemoveAll(tmp)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("GET %s: %s", url, resp.Status)
	}
	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		return err
	}
	defer gz.Close()
	tr := tar.NewReader(gz)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		name := hdr.Name
		if i := strings.IndexByte(name, '/'); i >= 0 {
			name = name[i+1:]
		} else {
			continue
		}
		if name == "" {
			continue
		}
		target := filepath.Join(tmp, name)
		switch hdr.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, 0o755); err != nil {
				return err
			}
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
				return err
			}
			f, err := os.OpenFile(target, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
			if err != nil {
				return err
			}
			if _, err := io.Copy(f, tr); err != nil {
				f.Close()
				return err
			}
			if err := f.Close(); err != nil {
				return err
			}
		}
	}
	_ = os.RemoveAll(dest)
	return os.Rename(tmp, dest)
}
