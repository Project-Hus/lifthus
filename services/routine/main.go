package main

import (
	"context"
	"lifthus-auth/common/lifthus"
	"net/http"

	"log"
	"os"
	actCommand "routine/internal/app/command/act"
	"routine/internal/ent"
	"routine/pkg/db"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	envbyjson "github.com/lifthus/envbyjson/go"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "routine/docs"

	_ "github.com/go-sql-driver/mysql"

	rmw "routine/pkg/middleware"
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
// @host localhost:9100
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

	// initialize lifthus common variables
	lifthus.InitLifthusVars(husenv, nil)

	// connect to lifthus_routine_db
	dbClient, err := db.ConnectToLifthusRoutine()
	if err != nil {
		log.Fatalf("[F]connecting db failed:%v", err)
	}
	defer dbClient.Close()

	// // create new http.Client from routineApi
	// routineHttpClient := &http.Client{
	// 	Timeout: time.Second * 5,
	// }

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
	e.Pre(rmw.UidSetter())

	// REQUEST LOGGER
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// get request ip address
			rip := c.RealIP()
			org := c.Request().Header.Get("Origin")
			log.Println("REQUEST from Origin", org, rip)
			log.Println(c.Request())
			return next(c)
		}
	})

	// programApiControllerParams := program.ProgramApiControllerParams{
	// 	DbClient:   dbClient,
	// 	HttpClient: routineHttpClient,
	// }

	// recApiControllerParams := rec.RecApiControllerParams(programApiControllerParams)

	// e = program.NewProgramApiController(e, programApiControllerParams)
	// e = rec.NewRecApiController(e, recApiControllerParams)

	e = actCommand.SetActCommandControllerTo(e)

	// swagger
	e.GET("/routine/openapi/*", echoSwagger.WrapHandler)

	// if the environment is native, run the echo server.
	if husenv == "native" {
		e.Logger.Fatal(e.Start(":" + os.Getenv("ROUTINE_PORT")))
	} else {
		// if it's lambda environment, run lambda.Start
		echoLambda = echoadapter.NewV2(e)
		lambda.Start(Handler)
	}
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	resp, err := echoLambda.ProxyWithContext(ctx, req)

	// RESPONSE LOGGER
	rip := req.RequestContext.HTTP.SourceIP
	org := req.Headers["Origin"]
	log.Println("RESPONSE to Origin", org, rip)
	log.Println(resp)
	log.Println("err:", err)

	return resp, err
}
