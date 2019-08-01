package completion

import "github.com/spf13/cobra"

func NewCompletionCmd(ctx interface{}) *cobra.Command {

	completionCmd := &cobra.Command{
		Use:   "completion",
		Short: "generates bash completion scripts",
		Long: `Outputs kind shell completion for the given shell (bash or zsh)
This depends on the bash-completion binary.  Example installation instructions:
# for bash users
        $ source <({{.RepoName}} completion bash)

# for zsh users
        % {{.RepoName}} completion zsh > /usr/local/share/zsh/site-functions/_{{.RepoName}}
        % autoload -U compinit && compinit

Additionally, you may want to output the completion to a file and source in your .bashrc

# ~/.bashrc or ~/.profile
. <({{.RepoName}} completion {bash/zsh})

Note for zsh users: [1] zsh completions are only supported in versions of zsh >= 5.2
`,
	}

	bashCmd := &cobra.Command{
		Use: "bash",
		Short: "Output shell completions for bash",
		Run: func(cmd *cobra.Command, args []string) {
			completionCmd.Parent().GenBashCompletion(cmd.OutOrStdout())
		},
	}

	zshCmd := &cobra.Command{
		Use: "zsh",
		Short: "Output shell completions for zsh",
		Run: func(cmd *cobra.Command, args []string) {
			completionCmd.Parent().GenZshCompletion(cmd.OutOrStdout())
		},
	}

	completionCmd.AddCommand(bashCmd)
	completionCmd.AddCommand(zshCmd)

	return completionCmd
}