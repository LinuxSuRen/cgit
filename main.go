package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main()  {
	cmd := &cobra.Command{
		Use: "cgit",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			env := os.Environ()

			preferGitHub(args)
			useMirror(args)

			var gitBinary string
			if gitBinary, err = exec.LookPath("git"); err == nil {
				syscall.Exec(gitBinary, append([]string{"git"}, args...), env)
			}
			return
		},
	}

	cmd.Execute()
}

func preferGitHub(args []string) {
	if len(args) <= 1 || args[0] != "clone" {
		return
	}

	address := args[1]
	if !strings.HasPrefix(address, "http") {
		args[1] = fmt.Sprintf("https://github.com.cnpmjs.org/%s", address)
	}
}

func useMirror(args []string) {
	for i, arg := range args {
		if strings.Contains(arg, "github.com") && !strings.Contains(arg, "github.com.cnpmjs.org") {
			args[i] = strings.ReplaceAll(arg, "github.com", "github.com.cnpmjs.org")
			break
		}
	}
}

