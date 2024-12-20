package tree_sitter_mxml_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-xml"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_xml.Language())
	if language == nil {
		t.Errorf("Error loading Xml grammar")
	}
}
