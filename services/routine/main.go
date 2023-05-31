package main

import (
	"lifthus-auth/common/lifthus"

	"log"
	"os"
	"routine/ent"

	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	envbyjson "github.com/lifthus/envbyjson/go"
)

var echoLambda *echoadapter.EchoLambdaV2
var dbClient *ent.Client

// @title Lifthus routine server
// @version 0.0.0
// @description This is Project-Hus's subservice Lifthus's routine management server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url lifthus531@gmail.com
// @contact.email lifthus531@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host auth.lifthus.com
// @BasePath /routine
func main() {
	// HUS_ENV
	// production : production for aws lambda
	// development : sam local environment
	// native : native go environment
	husenv, heok := os.LookupEnv("HUS_ENV")
	if !heok {
		log.Fatal("environment variable HUS_ENV must be set(production|development|native)")
	}
	// if husenv is native, load env.json with envbyjson
	if husenv == "native" {
		envbyjson.LoadProp("../../env.json", "Parameters")
	}

	// initialize lIfthus common variables
	lifthus.InitLifthusVars(husenv, nil)
}
