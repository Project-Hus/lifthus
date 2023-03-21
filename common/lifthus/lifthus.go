package lifthus

import (
	"lifthus-auth/ent"
	"os"
)

var GoogleClientID = "199526293983-r0b7tpmbpcc8nb786v261e451i2vihu3.apps.googleusercontent.com"

var Host = ""
var URL = ""
var Origins = []string{}
var AuthCookieDomain = ""
var AuthURL = ""
var ApiURL = ""

var LifthusURL = "http://localhost:3000"

func InitHusVars(_ *ent.Client) {
	goenv := os.Getenv("GOENV")
	//common
	GoogleClientID = os.Getenv("GOOGLE_CLIENT_ID")
	if goenv == "production" {
		Host = "lifthus.com"
		URL = "https://lifthus.com"
		Origins = []string{"https://cloudhus.com", "https://lifthus.com", "https://surfhus.com"}
		AuthCookieDomain = "auth.lifthus.com"
		AuthURL = "https://auth.lifthus.com"
		ApiURL = "https://api.lifthus.com"
	} else { // development
		Host = "localhost:9091"
		URL = "http://localhost:9091"
		Origins = []string{"https://localhost:3000", "http://localhost:9090", "http://localhost:9091"}
		AuthCookieDomain = ""
		AuthURL = "http://localhost:9091"
		ApiURL = "http://localhost:9091"
	}
}
