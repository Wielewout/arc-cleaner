package cmd

import (
	"context"
	"log/slog"
	"os"

	"github.com/actions/actions-runner-controller/apis/actions.github.com/v1alpha1"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/wielewout/arc-cleaner/internal/kubernetes"
	"github.com/wielewout/arc-cleaner/internal/logging"
)

var (
	version string
	commit  string
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "arc-cleaner",
	Short: "A cleaner for GitHub ARC",
	Long: `ARC cleaner is an application to clean up resources from the GitHub
Actions Runner Controller (ARC).`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		logger := slog.Default()
		logger.Info("started arc-cleaner", "version", version, "commit", commit)

		k8sClient, err := kubernetes.NewClient()
		if err != nil {
			logger.Error("failed to create kubernetes client", "error", err.Error())
		}

		namespace := viper.GetString("namespace")

		nsLogger := logger.With("namespace", namespace)
		nsCtx := logging.WithContext(ctx, nsLogger)

		ephemeralRunnerList := getEphemeralRunnerList(nsCtx, k8sClient, namespace)
		for _, ephemeralRunner := range ephemeralRunnerList.Items {
			erLogger := nsLogger.With("name", ephemeralRunner.Name)
			erLogger.Debug("ephemeral runner")
		}
	},
}

func getEphemeralRunnerList(ctx context.Context, k8sClient *kubernetes.Client, namespace string) *v1alpha1.EphemeralRunnerList {
	logger := logging.FromContext(ctx)

	ephemeralRunnerList := new(v1alpha1.EphemeralRunnerList)
	err := k8sClient.List(
		ctx,
		ephemeralRunnerList,
		client.InNamespace(namespace),
	)

	if err != nil {
		logger.Error("failed to list ephemeral runners", "error", err.Error())
		return ephemeralRunnerList
	}

	logger.Debug("listed ephemeral runners", "length", len(ephemeralRunnerList.Items))
	return ephemeralRunnerList
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.arc-cleaner.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".arc-cleaner")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		level := viper.GetString("log.level")
		logging.SetLevel(level)

		slog.Debug("using config file", "path", viper.ConfigFileUsed())
	}
}
