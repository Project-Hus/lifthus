package main

import (
	"context"
	"log"
	"net/http"
	"os"

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

var echoLambda *echoadapter.EchoLambda
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
	// set .env
	err := godotenv.Load() // now you can use os.Getenv("VAR_NAME")
	if err != nil {
		log.Fatalf("[F]loading .env file failed:%v", err)
	}

	// connecting to lifthus_user_db with ent
	client, err := db.ConnectToLifthusAuth()
	if err != nil {
		log.Fatalf("[F]connecting db failed:%v", err)
	}
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("[F]creating schema resources failed : %v", err)
	}

	// initialize Lifthus common variables
	lifthus.InitLifthusVars(os.Getenv("GOENV"), client)

	// subdomains
	// hosts := map[string]*Host{}

	//  Create echo web server instance and set CORS headers
	e := auth.NewAuthApiController(client) //echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{

		// If your Backend is deployed in AWS and using API Gateway to call through,
		// then all these headers need to be applied in API Gateway level also.
		AllowOrigins: lifthus.Origins,
		AllowHeaders: []string{
			echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization,
		},
		AllowCredentials: true,
		AllowMethods: []string{
			http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions, http.MethodPatch,
		},
	}))

	// // authApi, which controls auth all over the services
	// userApi := auth.NewAuthApiController(client)
	// hosts["localhost:9091"] = &Host{Echo: userApi} // gonna use auth.cloudhus.com later

	// // get requset and process by its subdomain
	// e.Any("/*", func(c echo.Context) (err error) {
	// 	fmt.Println(c.Request().Host, c.Request().URL.Path)
	// 	req, res := c.Request(), c.Response()
	// 	host, ok := hosts[req.Host] // if the host is not registered, it will be nil.
	// 	if !ok {
	// 		return c.NoContent(http.StatusNotFound)
	// 	} else {
	// 		host.Echo.ServeHTTP(res, req)
	// 	}
	// 	return err
	// })

	// provide api docs with swagger 2.0
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Run the server
	e.Logger.Fatal(e.Start(":9091"))
}

type Host struct {
	Echo *echo.Echo
}
