package config

// Constants for oauth.
const (
	HostURL    = "http://localhost:3000"
	APIRoot    = "/api"
	APIVersion = "v1"
	APIURL     = APIRoot + "/" + APIVersion
	OauthURL                  = HostURL + APIURL + "/oauth"
	OauthGoogleClientID       = ""
	OauthGoogleClientSecret   = ""
	OauthGoogleRedirectURL    = OauthURL + "/google/redirect"
	OauthGithubClientID       = ""
	OauthGithubClientSecret   = ""
	OauthGithubRedirectURL    = OauthURL + "/github/redirect"
	OauthYahooClientID        = ""
	OauthYahooClientSecret    = ""
	OauthYahooRedirectURL     = OauthURL + "/yahoo/redirect"
	OauthFacebookClientID     = ""
	OauthFacebookClientSecret = ""
	OauthFacebookRedirectURL  = OauthURL + "/facebook/redirect"
	OauthTwitterClientID      = ""
	OauthTwitterClientSecret  = ""
	OauthTwitterRedirectURL   = OauthURL + "/twitter/redirect"
	OauthLinkedinClientID     = ""
	OauthLinkedinClientSecret = ""
	OauthLinkedinRedirectURL  = OauthURL + "/linkedin/redirect"
	OauthKakaoClientID        = ""
	OauthKakaoClientSecret    = ""
	OauthKakaoRedirectURL     = OauthURL + "/kakao/redirect"
	OauthNaverClientID        = ""
	OauthNaverClientSecret    = ""
	OauthNaverRedirectURL     = OauthURL + "/naver/redirect"
)
