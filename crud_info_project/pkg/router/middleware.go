package router

import (
	"net/http"
	"strings"

	"info_project_service/pkg/server"
)

// Router CORS Configuration Struct
type routerCORSConfig struct {
	Origins string
	Methods string
	Headers string
}

// Router CORS Configuration Variable
var routerCORSCfg routerCORSConfig

// RouterCORS Function
func routerCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add Header for CORS
		w.Header().Set("Access-Control-Allow-Origin", routerCORSCfg.Origins)
		w.Header().Set("Access-Control-Allow-Methods", routerCORSCfg.Methods)
		w.Header().Set("Access-Control-Allow-Headers", routerCORSCfg.Headers)
		next.ServeHTTP(w, r)
	})
}

// RouterRealIP Function
func routerRealIP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		XForwardedFor := http.CanonicalHeaderKey("X-Forwarded-For")
		XRealIP := http.CanonicalHeaderKey("X-Real-IP")
		// Get Real IP from Cannoical Header
		if XForwardedFor != "" {
			dataIndex := strings.Index(XForwardedFor, ", ")
			if dataIndex == -1 {
				dataIndex = len(XForwardedFor)
			}
			r.RemoteAddr = XForwardedFor[:dataIndex]
		} else if XRealIP != "" {
			r.RemoteAddr = XRealIP
		}
		next.ServeHTTP(w, r)
	})
}

// RouterEntitySize Function
func routerEntitySize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate Entity Size
		r.Body = http.MaxBytesReader(w, r.Body, server.Config.GetInt64("SERVER_UPLOAD_LIMIT"))
		next.ServeHTTP(w, r)
	})
}
