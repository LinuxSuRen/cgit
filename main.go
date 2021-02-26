package main

import (
	"context"
	"fmt"
	"github.com/linuxsuren/cgit/cmd"
	"github.com/linuxsuren/cgit/pkg"
	ext "github.com/linuxsuren/cobra-extension"
	extver "github.com/linuxsuren/cobra-extension/version"
	aliasCmd "github.com/linuxsuren/go-cli-alias/pkg/cmd"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
)

const (
	TargetCLI = "git"
	AliasCLI  = "cgit"
)

func main() {
	command := &cobra.Command{
		Use: AliasCLI,
		RunE: func(command *cobra.Command, args []string) (err error) {
			fmt.Println(args, "sdfs")
			preHook(args)

			command.Println(args)

			var gitBinary string
			if gitBinary, err = exec.LookPath(TargetCLI); err == nil {
				err = pkg.ExecCommandInDir(gitBinary, "", args...)
			}
			return
		},
	}

	command.AddCommand(extver.NewVersionCmd("linuxsuren", AliasCLI, AliasCLI, nil))

	aliasCmd.AddAliasCmd(command, getAliasList())

	command.AddCommand(ext.NewCompletionCmd(command),
		cmd.NewMirrorCmd(context.TODO()))

	aliasCmd.Execute(command, TargetCLI, getAliasList(), preHook)
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
	if !strings.HasPrefix(address, "http") && !strings.HasPrefix(address, "git@") {
		args[1] = fmt.Sprintf("https://github.com.cnpmjs.org/%s", address)
	}
}

func useMirror(args []string) {
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
