package lifthus

import (
	"lifthus-auth/ent"
	"log"
	"net/http"
	"os"
	"time"
)

var GoogleClientID = ""
var HusSecretKey = ""
var HusSecretKeyBytes []byte

var Host = ""
var URL = ""
var Origins = []string{}
var CookieDomain = ""
var AuthURL = ""
var ApiURL = ""

var Http *http.Client

var LifthusURL = "http://localhost:3000"

var CookieSecure = true
var CookieSameSite = http.SameSiteLaxMode

var LifthusServiceName = "lifthus"

func GetLstExp() int64 {
	return time.Now().Add(time.Minute * 10).Unix() // 10 min basically
}

func InitLifthusVars(husenv string, _ *ent.Client) {
	Http = &http.Client{
		Timeout: time.Second * 5,
	}

	ok1, ok2 := false, false
	GoogleClientID, ok1 = os.LookupEnv("GOOGLE_CLIENT_ID")
	HusSecretKey, ok2 = os.LookupEnv("HUS_SECRET_KEY")
	HusSecretKeyBytes = []byte(HusSecretKey)
	if !ok1 || !ok2 {
		log.Fatalf("GOOGLE_CLIENT_ID or HUS_SECRET_KEY is not set")
	}
	switch husenv {
	case "production":
		Host = "lifthus.com"
		URL = "https://lifthus.com"
		Origins = []string{"http://localhost:3000", "https://cloudhus.com", "https://lifthus.com", "https://surfhus.com",
			"https://www.cloudhus.com", "https://www.lifthus.com", "https://www.surfhus.com"}
		CookieDomain = ".lifthus.com"
		AuthURL = "https://auth.lifthus.com"
		ApiURL = "https://api.lifthus.com"
	case "development":
		Host = "localhost:9100"
		URL = "http://localhost:9100"
		Origins = []string{"http://localhost:3000", "http://localhost:9000", "http://localhost:9200"}
		CookieDomain = ""
		AuthURL = "http://localhost:9091"
		ApiURL = "http://localhost:9091"
		CookieSecure = false
	case "native":
		Host = "localhost:9100"
		URL = "http://localhost:9101"
		Origins = []string{
			"http://localhost:3000",
			"http://localhost:9001",
			"http://localhost:9002",
			"http://localhost:9101",
			"http://localhost:9102",
		}
		CookieDomain = ""
		AuthURL = "http://localhost:9091"
		ApiURL = "http://localhost:9091"
		CookieSecure = false
	default:
		log.Fatal("HUS_ENV must be set(production|development|native)")
	}
	// for development server
	_, ok := os.LookupEnv("DEV_LIFTHUS_ACM_ARN")
	if !ok {
		LifthusServiceName = "lifthus-dev"
	}
}
