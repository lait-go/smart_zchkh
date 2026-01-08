package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func TaskServiceProxy() http.Handler {
	targer, err := url.Parse("http://localhost:8081")
	if err != nil {
		log.Fatal("invalid task_service URL: %v", err)
	}
	return httputil.NewSingleHostReverseProxy(targer)
}