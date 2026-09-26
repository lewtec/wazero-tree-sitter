// Host ABI for wazero. Every export takes and returns wasm32 integers.
// TSNode is passed as a pointer into wasm memory (sizeof reported by wts_node_sizeof).
#include <stddef.h>
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

#include <tree_sitter/api.h>

extern const TSLanguage *wts_language(void);

#define EXPORT __attribute__((used, visibility("default")))

EXPORT uint32_t wts_abi_version(void) { return 1; }

EXPORT uint32_t wts_node_sizeof(void) { return (uint32_t)sizeof(TSNode); }

static TSNode *node_at(uint32_t p) { return (TSNode *)(uintptr_t)p; }

EXPORT uint32_t wts_parser_new(void) {
  return (uint32_t)(uintptr_t)ts_parser_new();
}

EXPORT void wts_parser_delete(uint32_t parser) {
  ts_parser_delete((TSParser *)(uintptr_t)parser);
}

EXPORT uint32_t wts_parser_set_language(uint32_t parser) {
  const TSLanguage *lang = wts_language();
  if (lang == NULL) {
    return 0;
  }
  return ts_parser_set_language((TSParser *)(uintptr_t)parser, lang) ? 1 : 0;
}

EXPORT uint32_t wts_parser_parse(uint32_t parser, uint32_t src, uint32_t len) {
  return (uint32_t)(uintptr_t)ts_parser_parse_string(
      (TSParser *)(uintptr_t)parser, NULL, (const char *)(uintptr_t)src, len);
}

EXPORT void wts_tree_delete(uint32_t tree) {
  if (tree != 0) {
    ts_tree_delete((TSTree *)(uintptr_t)tree);
  }
}

EXPORT void wts_tree_root(uint32_t tree, uint32_t out_node) {
  *node_at(out_node) = ts_tree_root_node((const TSTree *)(uintptr_t)tree);
}

EXPORT uint32_t wts_node_is_null(uint32_t node) {
  return ts_node_is_null(*node_at(node)) ? 1 : 0;
}

EXPORT uint32_t wts_node_type(uint32_t node) {
  return (uint32_t)(uintptr_t)ts_node_type(*node_at(node));
}

EXPORT uint32_t wts_node_child_count(uint32_t node) {
  return ts_node_child_count(*node_at(node));
}

EXPORT void wts_node_child(uint32_t node, uint32_t index, uint32_t out_node) {
  *node_at(out_node) = ts_node_child(*node_at(node), index);
}

EXPORT uint32_t wts_node_field_name_for_child(uint32_t node, uint32_t index) {
  return (uint32_t)(uintptr_t)ts_node_field_name_for_child(*node_at(node), index);
}

EXPORT uint32_t wts_node_named_child_count(uint32_t node) {
  return ts_node_named_child_count(*node_at(node));
}

EXPORT void wts_node_named_child(uint32_t node, uint32_t index, uint32_t out_node) {
  *node_at(out_node) = ts_node_named_child(*node_at(node), index);
}

EXPORT uint32_t wts_node_start_byte(uint32_t node) {
  return ts_node_start_byte(*node_at(node));
}

EXPORT uint32_t wts_node_end_byte(uint32_t node) {
  return ts_node_end_byte(*node_at(node));
}

EXPORT uint32_t wts_node_is_named(uint32_t node) {
  return ts_node_is_named(*node_at(node)) ? 1 : 0;
}

EXPORT uint32_t wts_node_is_extra(uint32_t node) {
  return ts_node_is_extra(*node_at(node)) ? 1 : 0;
}

EXPORT uint32_t wts_node_is_error(uint32_t node) {
  return ts_node_is_error(*node_at(node)) ? 1 : 0;
}

EXPORT uint32_t wts_node_has_error(uint32_t node) {
  return ts_node_has_error(*node_at(node)) ? 1 : 0;
}

EXPORT uint32_t wts_node_has_changes(uint32_t node) {
  return ts_node_has_changes(*node_at(node)) ? 1 : 0;
}
