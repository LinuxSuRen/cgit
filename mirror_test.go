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

	// no need to insert mirror again
	abc = []string{"clone", "https://github.com.cnpmjs.org/cli/cli"}
	useMirror(abc)
	assert.Equal(t, abc[1], "https://github.com.cnpmjs.org/cli/cli")
}

func TestPreferGitHub(t *testing.T) {
	args := []string{"clone", "a/b"}
	preferGitHub(args)
	assert.Equal(t, args[1], "https://github.com.cnpmjs.org/a/b")
}
