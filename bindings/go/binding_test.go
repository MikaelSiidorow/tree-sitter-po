package tree_sitter_gettext_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_gettext "github.com/mikaelsiidorow/tree-sitter-gettext/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_gettext.Language())
	if language == nil {
		t.Errorf("Error loading PO grammar")
	}
}
