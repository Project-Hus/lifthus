package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	lmw "lifthus-auth/common/middleware"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	envbyjson "github.com/lifthus/envbyjson/go"

	"lifthus-auth/ent"

	"lifthus-auth/common/lifthus"
	"lifthus-auth/db"

	"lifthus-auth/api/auth"
	"lifthus-auth/api/user"

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
// @host auth.lifthus.com
// @BasePath /auth
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

	// initialize Lifthus common variables
	lifthus.InitLifthusVars(husenv, dbClient)

	// connecting to lifthus_user_db with ent
	dbClient, err := db.ConnectToLifthusAuth()
	if err != nil {
		log.Fatalf("[F]connecting db failed:%v", err)
	}
	// main's defer in lambda environment is actually not executed.
	defer dbClient.Close()

	// create new http.Client for authApi
	authHttpClient := &http.Client{
		Timeout: time.Second * 5,
	}
	authApiControllerParams := auth.AuthApiControllerParams{
		DbClient:   dbClient,
		HttpClient: authHttpClient,
	}
	userApiControllerParams := user.UserApiControllerParams{
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

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// get request ip address
			rip := c.RealIP()
			org := c.Request().Header.Get("Origin")
			fmt.Println("REQUEST from Origin", org, rip)
			fmt.Println(c.Request())
			return next(c)
		}
	})

	e = auth.NewAuthApiController(e, authApiControllerParams)
	e = user.NewUserApiController(e, userApiControllerParams)

	e.GET("/auth/openapi/*", echoSwagger.WrapHandler)

	// if the environment is native, run the echo server.
	if husenv == "native" {
		e.Logger.Fatal(e.Start(":" + os.Getenv("AUTH_PORT")))
	} else {
		// if it's in lambda environment, run lambda.Start.
		echoLambda = echoadapter.NewV2(e)
		lambda.Start(Handler)
	}
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	resp, err := echoLambda.ProxyWithContext(ctx, req)
	rip := req.RequestContext.HTTP.SourceIP
	org := req.Headers["Origin"]
	fmt.Println("RESPONSE to Origin", org, rip)
	fmt.Println(resp)
	fmt.Println("err:", err)
	return resp, err
}
