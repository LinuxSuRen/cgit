package main

import (
	"context"
	"fmt"
	"github.com/linuxsuren/cgit/cmd"
	"github.com/linuxsuren/cgit/pkg"
	ext "github.com/linuxsuren/cobra-extension/pkg"
	extver "github.com/linuxsuren/cobra-extension/version"
	aliasCmd "github.com/linuxsuren/go-cli-alias/pkg/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"path"
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

	ctx := context.TODO()
	command.AddCommand(ext.NewCompletionCmd(command),
		cmd.NewMirrorCmd(ctx))

	aliasCmd.ExecuteContextV2(command, context.TODO(), TargetCLI, getAliasList(), preHook)
}

func preHook(args []string) []string {
	args = preferGitHub(args)
	useMirror(args)
	return args
}

func preferGitHub(args []string) []string {
	if len(args) <= 1 || args[0] != "clone" {
		return args
	}

	address := args[1]
	if !strings.HasPrefix(address, "http") && !strings.HasPrefix(address, "git@") {
		args[1] = fmt.Sprintf("https://github.com.cnpmjs.org/%s", address)

		if len(args) == 2 {
			args = append(args, path.Join(viper.GetString("ws"), address))
		}
	}
	return args
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

func init() {
	viper.SetConfigName("cgit")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			return
		} else {
			panic(err)
		}
	}
	loadDefaults()
	return
}

func loadDefaults() {
	viper.SetDefault("ws", os.ExpandEnv("$HOME/ws/github/"))
}
