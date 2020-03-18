package testing

import (
	"testing"

	"data.structs/trees"
)

func TestInitializeTree(t *testing.T) {
	r := trees.InitializeTree()
	if r.Root == nil {
		t.Error("Fail")
	}
}
