package tree_sitter_po_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_po "github.com/mikaelsiidorow/tree-sitter-po/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_po.Language())
	if language == nil {
		t.Errorf("Error loading PO grammar")
	}
}
