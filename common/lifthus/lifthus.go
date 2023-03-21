package lifthus

import (
	"fmt"
	"lifthus-auth/ent"
	"os"
)

var GoogleClientID = ""
var HusSecretKey = ""

var Host = ""
var URL = ""
var Origins = []string{}
var CookieDomain = ""
var AuthURL = ""
var ApiURL = ""

var LifthusURL = "http://localhost:3000"

func InitLifthusVars(goenv string, _ *ent.Client) {
	//common
	GoogleClientID = os.Getenv("GOOGLE_CLIENT_ID")
	HusSecretKey = os.Getenv("HUS_SECRET_KEY")
	if goenv == "production" {
		Host = "lifthus.com"
		URL = "https://lifthus.com"
		Origins = []string{"https://cloudhus.com", "https://lifthus.com", "https://surfhus.com"}
		CookieDomain = ".lifthus.com"
		AuthURL = "https://auth.lifthus.com"
		ApiURL = "https://api.lifthus.com"
	} else { // development
		Host = "localhost:9091"
		URL = "http://localhost:9091"
		Origins = []string{"https://localhost:3000", "http://localhost:9090", "http://localhost:9091"}
		CookieDomain = ""
		AuthURL = "http://localhost:9091"
		ApiURL = "http://localhost:9091"
	}
	fmt.Println("initialized Lifthus vars for" + goenv)
}
