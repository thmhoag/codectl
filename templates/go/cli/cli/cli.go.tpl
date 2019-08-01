package cli

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/{{.OrgName}}/{{.RepoName}}/cli/completion"
	"github.com/{{.OrgName}}/{{.RepoName}}/cli/config"
	"github.com/{{.OrgName}}/{{.RepoName}}/cli/root"
	"github.com/{{.OrgName}}/{{.RepoName}}/cli/update"
	"github.com/{{.OrgName}}/{{.RepoName}}/cli/version"
	"os"
	"strings"
)

const name = "{{.RepoName}}"

func Execute() error {
	logger := logrus.New()
	cfg := *newConfig(name)
	repoCacheDir := fmt.Sprintf("%s/repository", getConfigDirPath(name))

	ctx := &globalCtx{
		log: *logrus.NewEntry(logger),
		config: cfg,
		workingDir: getWorkingDir(),
		version: version.Properties{
			Semver: Semver,
			Commit: Commit,
			Built:  Built,
		},
		repman: repomanager.NewManager(&repomanager.ManagerOpts{
			CacheDir: repoCacheDir,
			PropName: "repositories",
			Config: &cfg,
		}),
		generator: generator.LoadFromPath(repoCacheDir),
	}

	rootCmd := root.NewRootCmd(ctx)
	rootCmd.Version = ctx.CurrentVersion().Semver

	rootCmd.AddCommand(completion.NewCompletionCmd(ctx))
	rootCmd.AddCommand(version.NewVersionCmd(ctx))
	rootCmd.AddCommand(config.NewConfigCmd(ctx))
	rootCmd.AddCommand(update.NewUpdateCmd(ctx))

	return rootCmd.Execute()
}

func getConfigDirPath(appName string) string {
	home := getHomeDir()
	dirPath := fmt.Sprintf("%s/.%s", home, appName)

	return dirPath
}

func newConfig(appName string) *viper.Viper {
	configPath := getConfigDirPath(appName)
	configName := "config.yaml"
	createConfigIfNotExists(configPath, configName)

	v := viper.New()
	v.AddConfigPath(configPath)
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.SetEnvPrefix(strings.ToUpper(appName))
	v.AutomaticEnv()
	v.ReadInConfig()

	return v
}

func getHomeDir() string {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	return home
}

func getWorkingDir() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return dir
}

func createConfigIfNotExists(folderPath string, configFileName string) {
	fullPath := fmt.Sprintf("%s/%s", folderPath, configFileName)
	if !fileExists(fullPath) {
		os.MkdirAll(folderPath, os.ModePerm)
		os.OpenFile(fullPath, os.O_RDONLY|os.O_CREATE, 0666)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}