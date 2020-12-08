package main

import (
	"context"
	"fmt"
	extver "github.com/linuxsuren/cobra-extension/version"
	alias "github.com/linuxsuren/go-cli-alias/pkg"
	aliasCmd "github.com/linuxsuren/go-cli-alias/pkg/cmd"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main() {
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

	cmd.AddCommand(extver.NewVersionCmd("linuxsuren", "cgit", "cgit", nil))

	var ctx context.Context
	if defMgr, err := alias.GetDefaultAliasMgrWithNameAndInitialData(cmd.Name(), []alias.Alias{
		{Name: "cm", Command: "checkout master"},
	}); err == nil {
		ctx = context.WithValue(context.Background(), alias.AliasKey, defMgr)

		cmd.AddCommand(aliasCmd.NewRootCommand(ctx))
	} else {
		cmd.Println(fmt.Errorf("cannot get default alias manager, error: %v", err))
	}

	cmd.SilenceErrors = true
	a := cmd.Execute()
	if a != nil && strings.Contains(a.Error(), "unknown command") {
		args := os.Args[1:]
		var ctx context.Context
		var defMgr *alias.DefaultAliasManager
		var err error
		if defMgr, err = alias.GetDefaultAliasMgrWithNameAndInitialData(cmd.Name(), []alias.Alias{
			{Name: "cm", Command: "checkout master"},
		}); err == nil {
			ctx = context.WithValue(context.Background(), alias.AliasKey, defMgr)
			var gitBinary string
			var targetCmd []string
			env := os.Environ()

			if gitBinary, err = exec.LookPath("git"); err != nil {
				panic("cannot find git")
			}

			if ok, redirect := aliasCmd.RedirectToAlias(ctx, args); ok {
				args = redirect
			}

			preferGitHub(args)
			useMirror(args)

			targetCmd = append([]string{"git"}, args...)
			_ = syscall.Exec(gitBinary, targetCmd, env) // ignore the errors due to we've no power to deal with it
		} else {
			err = fmt.Errorf("cannot get default alias manager, error: %v", err)
		}
	}
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
