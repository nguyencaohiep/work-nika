package address

import (
	"explore_address/pkg/router"
	"explore_address/pkg/server"
	"explore_address/service/address/controller"
	"net"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

var ExploreAddressCrawlerSubRouter = chi.NewRouter()

func init() {

	ExploreAddressCrawlerSubRouter.Group(func(r chi.Router) {
		ExploreAddressCrawlerSubRouter.With(CheckLegitIPCrawler).Get("/address/check-exist", controller.CheckExistAddress)
	})
}

func CheckLegitIPCrawler(next http.Handler) http.Handler { // for transfer, holder, fund, venture
	// Return Next HTTP Handler Function, If Authorization is Valid
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		addr := getRealAddr(r)

		// log.Println(log.LogLevelInfo, "CheckLegitIPAdminServer", addr)

		legitIP := server.Config.GetString("LEGIT_IP_SERVER_CRAWLER")
		// log.Println(log.LogLevelInfo, addr, "")
		// the actual vaildation - replace with whatever you want
		if addr != legitIP {
			router.ResponseUnauthorized(w, "")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func getRealAddr(r *http.Request) string {

	remoteIP := ""
	// the default is the originating ip. but we try to find better options because this is almost
	// never the right IP
	if parts := strings.Split(r.RemoteAddr, ":"); len(parts) == 2 {
		remoteIP = parts[0]
	}
	// If we have a forwarded-for header, take the address from there
	if xff := strings.Trim(r.Header.Get("X-Forwarded-For"), ","); len(xff) > 0 {
		addrs := strings.Split(xff, ",")
		lastFwd := addrs[len(addrs)-1]
		if ip := net.ParseIP(lastFwd); ip != nil {
			remoteIP = ip.String()
		}
		// parse X-Real-Ip header
	} else if xri := r.Header.Get("X-Real-Ip"); len(xri) > 0 {
		if ip := net.ParseIP(xri); ip != nil {
			remoteIP = ip.String()
		}
	}

	return remoteIP

}
