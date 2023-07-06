package ffcore

import (
	"context"
	"log"
	"os"
	"time"

	"samfeatureflag/tracer"

	ffclient "github.com/thomaspoignant/go-feature-flag"
	"github.com/thomaspoignant/go-feature-flag/exporter/logsexporter"
	"github.com/thomaspoignant/go-feature-flag/ffuser"
	"github.com/thomaspoignant/go-feature-flag/retriever/fileretriever"
	"go.uber.org/zap"
)

func CreateFeatureFlagClient(configPath string) {
	ffclient.Init(ffclient.Config{
		Logger:          log.New(os.Stdout, "", 0),
		Context:         context.Background(),
		PollingInterval: 3 * time.Second,
		Retriever: &fileretriever.Retriever{
			Path: configPath,
		},
		DataExporter: ffclient.DataExporter{
			FlushInterval:    10,
			MaxEventInMemory: 2,
			Exporter: &logsexporter.Exporter{
				LogFormat: "[{{ .FormattedDate}}] user=\"{{ .UserKey}}\", flag=\"{{ .Key}}\", value=\"{{ .Value}}\", variation=\"{{ .Variation}}\"",
			},
		},
	})
	defer ffclient.Close()
}

func EvaluateSimpleFlag(flagKey, aUser string) error {

	user := ffuser.NewUser(aUser)
	hasFlag, _ := ffclient.BoolVariation(flagKey, user, false)
	//ffuser.NewAnonymousUser()

	tracer.Logger.Info("result", zap.Bool("value", hasFlag))
	if hasFlag {

	}
	return nil

}

func EvaluateABtestingFlag(flagKey string) error {

	user1 := ffuser.NewUserBuilder("123").Build()
	user2 := ffuser.NewUserBuilder("456").Build()
	var1, _ := ffclient.StringVariation(flagKey, user1, "error")
	var2, _ := ffclient.StringVariation(flagKey, user2, "error")

	tracer.Logger.Info("result", zap.String("user var1", var1), zap.String("user 2 var", var2))

	return nil

}
