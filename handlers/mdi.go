package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/TS22082/ts_cli_tool/utils"
	"github.com/charmbracelet/huh"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
)

func MDIHandler(cmd *cobra.Command, args []string) {
	var prompt string
	var dbName string

	dbClient := utils.MongoConnect()
	defer dbClient.Disconnect(context.Background())

	dbs, dbListErr := dbClient.ListDatabases(context.Background(), bson.D{})

	if dbListErr != nil {
		fmt.Println("❌ Error:", dbListErr)
		return
	}

	dbNames := make([]huh.Option[string], len(dbs.Databases))

	for i, db := range dbs.Databases {
		dbNames[i] = huh.NewOption(db.Name, db.Name)
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("Select a database").Options(dbNames...).Value(&dbName),
		),
		huh.NewGroup(
			huh.NewInput().Title("Add your prompt here").Prompt("👾").Value(&prompt),
		),
	)

	formErr := form.Run()

	if formErr != nil {
		fmt.Println("❌ Error:", formErr)
		return
	}

	apiKey := os.Getenv("OPENAI_API_KEY")

	if apiKey == "" {
		fmt.Println("❌ Error: OPENAI_API_KEY is not set, run with --help for more information")
		return
	}

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(`You are a helpful assistant that must return only valid JSON describing a MongoDB operation, e.g. { "operation": "find", "collection": "users", "filter": {...} }`),
			openai.SystemMessage(`The operation should either find, or count.`),
			openai.UserMessage(prompt),
		},
		Model: openai.ChatModelGPT4o,
	})

	if err != nil {
		panic(err.Error())
	}

	content := chatCompletion.Choices[0].Message.Content
	cleanContent := utils.CleanAIResponse(content)
	var queryData utils.MongoQuery

	queryErr := json.Unmarshal([]byte(cleanContent), &queryData)

	if queryErr != nil {
		fmt.Println("❌ Error:", queryErr)
		return
	}

	queryData.Database = dbName
	dbExecuteErr := utils.ExecuteMongoQuery(dbClient, queryData)

	if dbExecuteErr != nil {
		fmt.Println("❌ Error:", dbExecuteErr)
		return
	}
}
