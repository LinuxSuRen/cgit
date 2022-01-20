package pkg

import "strings"

// UseMirror modify the args with a GitHub mirror
func UseMirror(args []string) {
	// only for git clone
	if len(args) < 2 || args[0] != "clone" {
		return
	}
	for i, arg := range args {
		if strings.Contains(arg, "github.com") && !strings.Contains(arg, "github.com.cnpmjs.org") {
			args[i] = strings.ReplaceAll(arg, "github.com", "github.com.cnpmjs.org")
			break
		}
	}
}
