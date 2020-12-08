package main

import (
	"fmt"
	ext "github.com/linuxsuren/cobra-extension"
	extver "github.com/linuxsuren/cobra-extension/version"
	aliasCmd "github.com/linuxsuren/go-cli-alias/pkg/cmd"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

const (
	TargetCLI = "git"
	AliasCLI  = "cgit"
)

func main() {
	cmd := &cobra.Command{
		Use: AliasCLI,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			env := os.Environ()

			preHook(args)

			var gitBinary string
			if gitBinary, err = exec.LookPath(TargetCLI); err == nil {
				syscall.Exec(gitBinary, append([]string{TargetCLI}, args...), env)
			}
			return
		},
	}

	cmd.AddCommand(extver.NewVersionCmd("linuxsuren", AliasCLI, AliasCLI, nil))

	aliasCmd.AddAliasCmd(cmd, getAliasList())

	cmd.AddCommand(ext.NewCompletionCmd(cmd))

	aliasCmd.Execute(cmd, TargetCLI, getAliasList(), preHook)
}

func preHook(args []string) {
	preferGitHub(args)
	useMirror(args)
}

func preferGitHub(args []string) {
	if len(args) <= 1 || args[0] != "clone" {
		return
	}

	address := args[1]
	if !strings.HasPrefix(address, "http") || !strings.HasPrefix(address, "git@") {
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
