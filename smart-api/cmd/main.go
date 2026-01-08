package main

import (
	"net/http"
	"smart-api/api"
	"smart-api/internal/middleware"
	"smart-api/internal/proxy"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/charges", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			api.ChargesHandler(w, r)
		} else if r.Method == http.MethodPost {
			api.CreateChargeHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/charges/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			api.ChargeHandlerByID(w, r)
		}
		if r.Method == http.MethodPut {
			api.ChargesHandlerPut(w, r)
		}
	})
	mux.HandleFunc("/api/accounts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			api.AccountsHadlerGet(w, r)
		} else if r.Method == http.MethodPost {
			api.AccountsHadlerPost(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/api/accounts/", api.AccountsHandlerByUserID)

	mux.Handle("/api/login", proxy.ProxyAuthService())
	mux.Handle("/api/register", proxy.ProxyAuthService())

	mux.Handle("/api/v1/tasks", proxy.TaskServiceProxy())
	mux.Handle("/api/v1/tasks/", proxy.TaskServiceProxy())

	mux.Handle("/api/charge", proxy.ChargeServerProxy())	
	mux.Handle("/api/charge/", proxy.ChargeServerProxy())	

	handler := middleware.CorsMiddleware(mux)
	http.ListenAndServe(":8080", handler)
}
