package lifthus

import "net/http"

// CookieMaker takes name and value and generates default lifthus auth cookie.
func AuthCookieMaker(name string, value string) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		Domain:   ".lifthus.com",
		HttpOnly: true,
		Secure:   CookieSecure,
		SameSite: CookieSameSite,
	}
}
