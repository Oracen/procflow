package stringhandle

import (
	"testing"

	"github.com/Oracen/procflow/core/constants"
	"github.com/stretchr/testify/assert"
)

func TestConvertToGraphPackage(t *testing.T) {

	t.Run(
		"test node name conversion functions",
		func(t *testing.T) {
			parent := "parent"
			child := "child"
			got := PackNames(parent, child)
			want := len(parent) + len(child) + len(constants.StandardDelimiter)
			assert.Len(t, got, want)

			got = UnpackNames(parent, got)
			assert.Equal(t, child, got)

		},
	)

}
