package Middleware

import (
	"fmt"
	"net/http"
	"server/Cache"
	"server/Logging"
	"server/Response"
)

func NoAuthLogging(next http.Handler) http.Handler {
	/*
		This middleware will be used to log the incoming requests
		It will take input the footprint of the Handle function which will handle the request
		After logging the request will be redirected to the appropriate Handle function
	*/
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// r.RemoteAddr to get the ip address of the device, the request came from
		// r.RequestURI to get the api end point which the request has hit
		Logging.REQ.Println(r.RemoteAddr, " ", r.Proto, " ", r.Method, " ", r.RequestURI)
		// Redirecting the request to the appropriate Handle function
		next.ServeHTTP(w, r)

	})
}

func AuthLogging(next http.Handler) http.Handler {
	/*
		This middleware will be used to log the incoming requests
		It will take input the footprint of the Handle function which will handle the request
		After logging the request will be redirected to the appropriate Handle function
	*/
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// r.RemoteAddr to get the ip address of the device, the request came from
		// r.RequestURI to get the api end point which the request has hit
		Logging.REQ.Println(r.RemoteAddr, " ", r.Proto, " ", r.Method, " ", r.RequestURI)
		// Redirecting the request to the appropriate Handle function

		if Cache.VerifyAccessToken(r.Header.Get("uid"), r.Header.Get("access_token")) {
			next.ServeHTTP(w, r)
		} else {
			fmt.Println("response is 102")
			Response.Unauthorized(w, r, "102")
		}

	})
}
