package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"lifthus-auth/ent"

	"lifthus-auth/common/lifthus"
	"lifthus-auth/db"

	"lifthus-auth/api/auth"

	_ "lifthus-auth/docs"

	echoSwagger "github.com/swaggo/echo-swagger"

	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
)

var echoLambda *echoadapter.EchoLambdaV2
var dbClient *ent.Client

// @title Lifthus user server
// @version 0.0.0
// @description This is Project-Hus's subservice Lifthus's user management server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url lifthus531@gmail.com
// @contact.email lifthus531@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host lifthus.com
// @BasePath /
func main() {
	// GOENV
	// production : production for aws lambda
	// development : sam local environment
	// native : native go environment
	goenv, ok := os.LookupEnv("GOENV")
	if !ok {
		log.Fatalf("GOENV is not set")
	}

	// in production environment, env vars comes from parameter store.
	// in development environment, env vars comes from env.json.
	// in native Go environment, load env vars from .env
	if goenv == "native" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("loading .env file failed : %s", err)
		}
	}

	// connecting to lifthus_user_db with ent
	dbClient, err := db.ConnectToLifthusAuth()
	if err != nil {
		log.Fatalf("[F]connecting db failed:%v", err)
	}
	if goenv == "native" {
		defer dbClient.Close()
	}

	// initialize Lifthus common variables
	lifthus.InitLifthusVars(goenv, dbClient)

	// create new http.Client for authApi
	authHttpClient := &http.Client{
		Timeout: time.Second * 5,
	}
	authApiControllerParams := auth.AuthApiControllerParams{
		DbClient:   dbClient,
		HttpClient: authHttpClient,
	}

	// create echo web server instance and set CORS headers
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// If your Backend is deployed in AWS and using API Gateway to call through,
		// then all these headers need to be applied in API Gateway level also.
		AllowOrigins: lifthus.Origins,

		// to allow all headers
		AllowHeaders: []string{
			echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization,
			echo.HeaderAccessControlAllowOrigin, echo.HeaderAccessControlAllowHeaders, echo.HeaderAccessControlAllowMethods,
			echo.HeaderXRequestedWith,
		},
		AllowCredentials: true,
		AllowMethods: []string{
			http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions, http.MethodPatch,
		},
	}))
	e = auth.NewAuthApiController(e, authApiControllerParams)

	e.GET("/auth/openapi/*", echoSwagger.WrapHandler)

	if goenv == "native" {
		e.Logger.Fatal(e.Start(":9091"))
	} else {
		echoLambda = echoadapter.NewV2(e)
		lambda.Start(Handler)
	}

	// Run the server
	e.Logger.Fatal(e.Start(":9091"))
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	resp, err := echoLambda.ProxyWithContext(ctx, req)
	hst := req.Headers["Host"]
	pth := req.RequestContext.HTTP.Path

	org := req.Headers["Origin"]
	fmt.Println("RESPONSE==========", hst, pth, org)
	fmt.Println(fmt.Sprintf("%+v", resp))
	fmt.Println("err:", err)
	return resp, err
}
