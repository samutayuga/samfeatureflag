package cmd

import (
	"log"
	"samfeatureflag/ffcore"
	"samfeatureflag/tracer"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	flagConfigPath string
	flagKey        string
	user           string
	masterCmd      = &cobra.Command{
		Use: "mastercommand",
		Args: func(cmd *cobra.Command, args []string) error {
			if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
				return err
			}
			tracer.Logger.Info("all are good")
			return nil
		},

		Short: "a parent command that holds the rest of the command",
		Long: `Cobra is a CLI library for Go that empowers applications.
		This application is a tool to experiment with the feature flag evaluation.`,
		TraverseChildren: true,
		// Run: func(cmd *cobra.Command, args []string) {
		// },
	}
)

var SimpleBoolFlagCmd = &cobra.Command{
	Use:   "simple",
	Short: "boolean feature flag",
	Long:  `To launch boolean feature flag`,
	Args: func(cmd *cobra.Command, args []string) error {

		tracer.Logger.Info("all are good", zap.Any("args", args), zap.String("flag", flagConfigPath))

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ffcore.EvaluateSimpleFlag(flagKey, user)
		return nil
	},
}

var ABTestingCmd = &cobra.Command{
	Use:   "abtesting",
	Short: "AB Testing flag",
	Long: `Perform the evaluation on the AB testing flag
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ffcore.EvaluateABtestingFlag(flagKey)
		return nil
	},
}

func init() {
	//zap.NewDevelopment()
	cobra.OnInitialize(func() {
		log.Printf("Initialize")
		ffcore.CreateFeatureFlagClient("config/flags-config.yaml")

	})
	masterCmd.AddCommand(SimpleBoolFlagCmd)
	masterCmd.AddCommand(ABTestingCmd)
	//masterCmd.AddCommand(helloCmd)
	//SimpleBoolFlagCmd.Flags().StringVarP(&flagConfigPath, "flagsPath", "c", "config/flags-config.yaml", "--flagsPath config/flags-config.yaml")
	SimpleBoolFlagCmd.Flags().StringVarP(&flagKey, "flagKey", "k", "test-flag", "--flagKey test-flag")
	SimpleBoolFlagCmd.Flags().StringVarP(&user, "user", "u", "user123", "--user user123")
	SimpleBoolFlagCmd.MarkFlagRequired("flagKey")
	ABTestingCmd.Flags().StringVarP(&flagKey, "flagKey", "k", "experimentation-flag", "--flagKey experimentation-flag")
	ABTestingCmd.MarkFlagRequired("flagKey")
}

func Execute() error {
	return masterCmd.Execute()
}
