package main

import (
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

func useMirror(args []string) {
	for i, arg := range args {
		if strings.Contains(arg, "github.com") {
			args[i] = strings.ReplaceAll(arg, "github.com", "github.com.cnpmjs.org")
			break
		}
	}
}

