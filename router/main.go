package main

import (
	// license "licensing/license/licenseServer"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)
func main() {
	// // Router = routes.NewRouterServer()
	// http.ListenAndServe(":8080",nil)
	routes := map[string]string{
		"/getGPS":"http://localhost:8000",
		// "/prod-eval/activatekey":"http://localhost:8000",
		// "/prod-eval/createuser":"http://localhost:8000",
		// "/prod-eval/getallcustomer":"http://localhost:8000",
		// "/prod-eval/getlicensedata":"http://localhost:8000",
		// "/prod-eval/updatelicensedata":"http://localhost:8000",
		// "/prod-eval/deactivatelicense":"http://localhost:8000",
	}
	for prefix, target := range routes{
		url, err := url.Parse(target)
		if err != nil{
			log.Fatal(err.Error())
		} 
		handler := httputil.NewSingleHostReverseProxy(url)
		fmt.Println(prefix)
		http.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
			handler.ServeHTTP(w, r)
		})
	}
	// http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {w.Write([]byte("Got request"))})
	http.ListenAndServe(":9000",nil)
}
// func Callproxy(l *license.Server1){
// 	router.ProxyServer(l)
// }