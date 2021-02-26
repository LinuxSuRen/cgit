package cmd

import (
	"context"
	"fmt"
	"github.com/linuxsuren/cgit/pkg"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
)

type mirrorOption struct {
	enable bool
	remote string
}

// NewMirrorCmd returns the mirror command
func NewMirrorCmd(ctx context.Context) (cmd *cobra.Command) {
	opt := mirrorOption{}

	cmd = &cobra.Command{
		Use:   "mirror",
		Short: "Toggle the git mirror",
		RunE:  opt.runE,
	}

	flags := cmd.Flags()
	flags.BoolVarP(&opt.enable, "enable", "e", true, "Enable/disable the git mirror")
	flags.StringVarP(&opt.remote, "remote", "r", "origin", "The remote of git repository")
	return
}

func (o *mirrorOption) runE(cmd *cobra.Command, args []string) (err error) {
	var remoteURLStr string
	var remotePushURLStr string
	if remoteURLStr, err = pkg.ExecCommandWithOutput("git", "", "remote", "get-url", o.remote); err != nil {
		return
	}

	if remotePushURLStr, err = pkg.ExecCommandWithOutput("git", "", "remote", "get-url", "--push", o.remote); err != nil {
		return
	}
	// remove the .git tail
	remoteURLStr = strings.TrimSuffix(remoteURLStr, ".git")

	gitProtocol := strings.HasPrefix(remoteURLStr, "git@")
	if gitProtocol {
		remoteURLStr = strings.ReplaceAll(remoteURLStr, "git@github.com:", "https://github.com/")
	}

	var remoteURL *url.URL
	if remoteURL, err = url.Parse(remoteURLStr); err != nil {
		cmd.Println("error with parse URL", remoteURLStr)
		return
	}

	if o.enable {
		remoteURL.Host = "github.com.cnpmjs.org"
	} else {
		remoteURL.Host = "github.com"
	}

	targetRemoteURLStr := remoteURL.String()
	if strings.HasPrefix(remotePushURLStr, "git@") && !o.enable {
		// the mirror git server does not support git protocol
		targetRemoteURLStr = strings.ReplaceAll(targetRemoteURLStr, fmt.Sprintf("https://%s/", remoteURL.Host),
			fmt.Sprintf("git@%s:", remoteURL.Host))
		targetRemoteURLStr += ".git"
	}

	if err = pkg.ExecCommandInDir("git", "", "remote", "set-url", o.remote, targetRemoteURLStr); err != nil {
		return
	}
	if err = pkg.ExecCommandInDir("git", "", "remote", "set-url", "--push", o.remote, remotePushURLStr); err != nil {
		return
	}
	return
}
