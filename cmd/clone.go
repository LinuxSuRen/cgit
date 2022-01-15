package cmd

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/linuxsuren/cgit/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os/exec"
	"path"
	"strings"
)

type cloneOption struct {
	ws bool
}

// NewCloneCommand returns the clone command
func NewCloneCommand() (cmd *cobra.Command) {
	opt := &cloneOption{}

	cmd = &cobra.Command{
		Use:   "clone",
		Short: "A smart way to clone repositories from GitHub",
		RunE:  opt.runE,
	}

	flags := cmd.Flags()
	flags.BoolVarP(&opt.ws, "ws", "", false, "Clone the code into ~/ws/github/org/repo if it is true")
	return
}

func (o *cloneOption) runE(_ *cobra.Command, args []string) (err error) {
	output := func(arg string) string {
		if orgAndRepo := strings.Split(arg, "/"); len(orgAndRepo) == 2 {
			return path.Join(viper.GetString("ws"), arg)
		}
		return ""
	}
	if !o.ws {
		output = nil
	}
	args = pkg.ParseShortCode(args, output)

	var targetDir string
	gitAddress := args[0]
	if len(args) >= 2 {
		targetDir = args[1]
	}

	var gitBinary string
	if gitBinary, err = exec.LookPath("git"); err == nil {
		gitArgs := []string{"clone"}
		gitArgs = append(gitArgs, args...)
		pkg.UseMirror(gitArgs)
		if err = pkg.ExecCommandInDir(gitBinary, "", gitArgs...); err == nil {
			err = pkg.ExecCommandInDir(gitBinary, targetDir, "remote", "set-url", "origin", gitAddress)
		}
	}
	if err != nil {
		return
	}

	var ghBinary string
	if ghBinary, err = exec.LookPath("gh"); err == nil {
		prompt := &survey.Confirm{
			Message: "do you want to fork it?",
		}
		var ok bool
		if err = survey.AskOne(prompt, &ok); err == nil && ok {
			err = pkg.ExecCommandInDir(ghBinary, targetDir, "repo", "fork", "--remote")
		}
	}
	return
}
