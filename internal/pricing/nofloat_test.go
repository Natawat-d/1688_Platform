package pricing

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestNoFloats keeps the money path integer-only. A float that creeps into a
// price is not a rounding curiosity, it is a cent that goes missing on one
// order in ten thousand and cannot be reproduced afterwards, so this fails the
// build rather than a review.
func TestNoFloats(t *testing.T) {
	entries, err := os.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}
	fset := token.NewFileSet()
	parsed := 0
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") {
			continue
		}
		f, err := parser.ParseFile(fset, filepath.Join(".", name), nil, 0)
		if err != nil {
			t.Fatalf("parse %s: %v", name, err)
		}
		parsed++
		ast.Inspect(f, func(n ast.Node) bool {
			if id, ok := n.(*ast.Ident); ok && (id.Name == "float64" || id.Name == "float32") {
				t.Errorf("%s: %s is banned in package pricing", fset.Position(id.Pos()), id.Name)
			}
			return true
		})
	}
	if parsed == 0 {
		t.Fatal("no non-test source files found to check")
	}
}
