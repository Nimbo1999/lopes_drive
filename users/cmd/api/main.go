package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/nimbo1999/lopes_drive/users/internal/app"
	"github.com/nimbo1999/lopes_drive/users/internal/config"
)

func main() {
	// Setup the environment variables with viper
	config.ReadEnvs()
	env := config.CreateEnvironment()

	// AWS Session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config:            *aws.NewConfig().WithRegion("us-east-1").WithCredentials(credentials.NewEnvCredentials()),
	}))

	// Create DynamoDB client
	dynamo := dynamodb.New(sess)

	params := &app.AppParams{
		Db:  dynamo,
		Env: env,
	}

	usersApp := app.NewApp(params)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", env.Port), usersApp.Handler); err != nil {
		log.Fatalln(err)
	}
}
