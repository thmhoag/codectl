package update

import (
	"errors"
	"fmt"
	"github.com/blang/semver"
	"github.com/manifoldco/promptui"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"time"
)

type updateOpts struct {
	auth string
}

func NewUpdateCmd(ctx Ctx) *cobra.Command {
	opts := &updateOpts{}

	version := ctx.CurrentVersion()
	log := ctx.Log()

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update the {{.RepoName}} binary",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			v, err := semver.ParseTolerant(version.Semver)
			if err != nil {
				cmd.PrintErrln("Unable to get current version")
				log.WithError(err).Fatalln("Unable to get current version")
			}

			updater, err := selfupdate.NewUpdater(selfupdate.Config{
				APIToken: opts.auth,
			})

			if err != nil {
				log.WithError(err).Fatalln("Unable to create updater")
			}

			githubRepo := "{{.OrgName}}/{{.RepoName}}"
			log.WithField("repo", githubRepo).Debugln("detecting version from repo")
			latest, found, err := updater.DetectLatest("{{.OrgName}}/{{.RepoName}}")
			if err != nil {
				cmd.PrintErr("Unable to detect latest version")
				log.WithError(err).Fatal("error detecting version")
			}

			log.Debugln("version found: ", found)
			if !found || latest.Version.LTE(v) {
				cmd.Println("Current version is the latest")
				return
			}

			label := fmt.Sprintf("Do you want to update to %s? (y/n)", latest.Version)
			prompt := promptui.Prompt{
				Label:     label,
				Default:   "n",
				AllowEdit: false,
				Validate:  validateInput,
			}

			input, err := prompt.Run()
			if err != nil {
				cmd.PrintErrln("Error confirming to update")
				log.WithError(err).Fatal("error getting user input to confirm update")
			}

			input = strings.ToLower(input)
			if input == "n" {
				log.Debugln("user elected not to update")
				return
			}

			exe, err := os.Executable()
			if err != nil {
				cmd.PrintErrln("Error updating, unable to detect executable")
				log.WithError(err).Fatalln("unable to detect executable")
			}

			cmd.Print("Updating...")
			result := make(chan error)
			defer close(result)

			go func() {
				if err := updater.UpdateTo(latest, exe); err != nil {
					result <- err
				}

				result <- nil
			}()

			finished := false
			for !finished {
				select {
				case err := <- result:
					finished = true
					cmd.Println() // to end the "loading" line

					if err != nil {
						cmd.PrintErrln("Error occurred updating binary")
						log.WithError(err).Fatalln("error updating binary")
					}
				case <- time.After(time.Second/2):
					cmd.Print(".")
				}
			}

			cmd.Println("Successfully updated to version ", latest.Version)
		},
	}

	cmd.Flags().StringVar(&opts.auth, "auth", "", "Personal access token for Github API auth")

	return cmd
}

func validateInput(input string) error {
	i := strings.ToLower(input)
	if i != "y" && i != "n" {
		return errors.New("answer must be yes (y) or no (n)")
	}

	return nil
}
