package auth

import (
	"encoding/json"
	"explore_address/pkg/router"
	"explore_address/pkg/server"
	"io"
	"net/http"
)

var (
	_googleReCaptchaSecret = server.Config.GetString("GOOGLE_RECAPTCHA_SECRET")
)

const (
	_googleReCaptchaVerifyAPI = "https://www.google.com/recaptcha/api/siteverify?"
)

func ReCaptcha(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reCaptchaResponse := r.Header.Get("ReCaptchaResponse")

		request := _googleReCaptchaVerifyAPI + "secret=" + _googleReCaptchaSecret + "&response=" + reCaptchaResponse + "&remoteip=" + r.RemoteAddr
		response, err := http.Get(request)
		if err != nil {
			router.ResponseUnauthorized(w, "google re-captcha is invalid")
			return
		}
		if response.StatusCode != http.StatusOK {
			router.ResponseUnauthorized(w, "google re-captcha is invalid")
			return
		}
		body, err := io.ReadAll(response.Body)
		if err != nil {
			router.ResponseUnauthorized(w, "google re-captcha is invalid")
			return
		}

		data := map[string]any{}
		err = json.Unmarshal(body, &data)
		if err != nil {
			router.ResponseUnauthorized(w, "google re-captcha is invalid")
			return
		}

		if !data["success"].(bool) {
			router.ResponseUnauthorized(w, "google re-captcha is invalid")
			return
		}

		next.ServeHTTP(w, r)
	})
}
