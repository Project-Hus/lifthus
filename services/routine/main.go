package main

import (
	"lifthus-auth/common/lifthus"
	"net/http"
	"time"

	"log"
	"os"
	"routine/common/db"
	"routine/ent"

	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	envbyjson "github.com/lifthus/envbyjson/go"

	lmw "lifthus-auth/common/middleware"
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

	// connect to lifthus_routine_db
	dbClient, err := db.ConnectToLifthusRoutine()
	if err != nil {
		log.Fatalf("[F]connecting db failed:%v", err)
	}
	defer dbClient.Close()

	// create new http.Client from routineApi
	routineHttpClient := &http.Client{
		Timeout: time.Second * 5,
	}

	// create echo web server instance an set CORS headers
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: lifthus.Origins,
		AllowHeaders: []string{
			echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAuthorization, echo.HeaderAccessControlAllowOrigin,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderAccessControlAllowMethods,
			echo.HeaderXRequestedWith,
		},
		ExposeHeaders: []string{
			echo.HeaderAuthorization,
		},
		AllowCredentials: true,
		AllowMethods: []string{
			http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions, http.MethodPatch,
		},
	}))

	// set uid to context if the user is signed
	e.Pre(lmw.UidSetter(dbClient))
}
