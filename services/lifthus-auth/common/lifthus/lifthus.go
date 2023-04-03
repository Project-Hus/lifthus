package lifthus

import (
	"lifthus-auth/ent"
	"log"
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
	ok1, ok2 := false, false
	GoogleClientID, ok1 = os.LookupEnv("GOOGLE_CLIENT_ID")
	HusSecretKey, ok2 = os.LookupEnv("HUS_SECRET_KEY")
	if !ok1 || !ok2 {
		log.Fatalf("GOOGLE_CLIENT_ID or HUS_SECRET_KEY is not set")
	}
	if goenv == "production" {
		Host = "lifthus.com"
		URL = "https://lifthus.com"
		Origins = []string{"https://cloudhus.com", "https://lifthus.com", "https://surfhus.com",
			"https://www.cloudhus.com", "https://www.lifthus.com", "https://www.surfhus.com"}
		CookieDomain = ".lifthus.com"
		AuthURL = "https://auth.lifthus.com"
		ApiURL = "https://api.lifthus.com"
	} else { // development or native
		Host = "localhost:9091"
		URL = "http://localhost:9091"
		Origins = []string{"http://localhost:3000", "http://localhost:9090", "http://localhost:9091"}
		CookieDomain = ""
		AuthURL = "http://localhost:9091"
		ApiURL = "http://localhost:9091"
	}
}
