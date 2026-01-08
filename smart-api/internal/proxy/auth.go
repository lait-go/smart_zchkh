package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ProxyAuthService() http.Handler {
	targer, err := url.Parse("http://localhost:8082")
	if err != nil {
		log.Fatal("invalid auth_service URL: %v", err)
	}
	return httputil.NewSingleHostReverseProxy(targer)
}