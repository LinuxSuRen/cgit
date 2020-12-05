package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestMirror(t *testing.T) {
	abc := []string{"a", "b", "github.com"}
	useMirror(abc)
 	assert.Equal(t, abc[2], "github.com.cnpmjs.org")

	abc = []string{"clone", "https://github.com/cli/cli"}
	useMirror(abc)
	assert.Equal(t, abc[1], "https://github.com.cnpmjs.org/cli/cli")
}
