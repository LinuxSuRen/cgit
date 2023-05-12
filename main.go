package main

import (
	"context"
	"fmt"
	"github.com/linuxsuren/cgit/cmd"
	"github.com/linuxsuren/cgit/pkg"
	ext "github.com/linuxsuren/cobra-extension/pkg"
	extver "github.com/linuxsuren/cobra-extension/version"
	aliasCmd "github.com/linuxsuren/go-cli-alias/pkg/cmd"
	"github.com/linuxsuren/http-downloader/pkg/installer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
)

const (
	// TargetCLI is the target command for alias
	TargetCLI = "git"
	// AliasCLI is the alias command
	AliasCLI = "cgit"
)

func main() {
	command := &cobra.Command{
		Use: AliasCLI,
		RunE: func(command *cobra.Command, args []string) (err error) {
			preHook(args)

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
		cmd.NewMirrorCmd(ctx),
		cmd.NewCloneCommand())

	// do the dep checking
	if err := installDepTools(); err != nil {
		panic(err)
	}

	aliasCmd.ExecuteContextV2(command, context.TODO(), TargetCLI, getAliasList(), preHook)
}

func installDepTools() (err error){
	is := &installer.Installer{
		Provider: "github",
		OS:       runtime.GOOS,
		Arch:     runtime.GOARCH,
		Fetch:    true,
	}
	dep := map[string]string{
		"git":        "git",
	}

	err = is.CheckDepAndInstall(dep)
	return
}

func preHook(args []string) []string {
	args = parseShortCode(args)
	useMirror(args)
	return args
}

func parseShortCode(args []string) []string {
	if len(args) <= 1 || args[0] != "clone" {
		return args
	}

	address := args[1]
	if strings.HasPrefix(address, "gitee.com") {
		args[1] = fmt.Sprintf("https://%s", address)

		if len(args) == 2 {
			args = append(args, path.Join(strings.ReplaceAll(viper.GetString("ws"), "github", "gitee"),
				strings.ReplaceAll(address, "gitee.com", "")))
		}
	} else if strings.HasPrefix(address, "https://gitee.com/") {
		if len(args) == 2 {
			args = append(args, path.Join(strings.ReplaceAll(viper.GetString("ws"), "github", "gitee"),
				strings.ReplaceAll(address, "https://gitee.com/", "")))
		}

	} else if !strings.HasPrefix(address, "http") && !strings.HasPrefix(address, "git@") {
		args[1] = fmt.Sprintf("https://ghproxy.com/%s", address)

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
		if strings.Contains(arg, "github.com") && !strings.Contains(arg, "ghproxy.com") {
			args[i] = strings.ReplaceAll(arg, "github.com", "ghproxy.com")
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
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(err)
		}
	}
	loadDefaults()
}

func loadDefaults() {
	viper.SetDefault("ws", os.ExpandEnv("$HOME/ws/github/"))
}
