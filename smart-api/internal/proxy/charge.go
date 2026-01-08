package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ChargeServerProxy() http.Handler {
	targer, err := url.Parse("http://localhost:7070")
	if err != nil {
		log.Fatal("invalid charge_service URL: %v", err)
	}
	return httputil.NewSingleHostReverseProxy(targer)
}