/**
 * @file PO (gettext) grammar for tree-sitter
 * @author Mikael Siidorow <mikael@siidorow.com>
 * @license MIT
 */

/// <reference types="tree-sitter-cli/dsl" />
// @ts-check

module.exports = grammar({
  name: "po",

  rules: {
    // TODO: add the actual grammar rules
    source_file: $ => "hello"
  }
});
