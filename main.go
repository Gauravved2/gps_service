package main

import (
	"gpgps/gpsapi"
	"net/http"
)

func main() {
	srv := gpsapi.NewServer()
	if err:=http.ListenAndServe(":8000", srv);err!=nil{
		panic(err.Error())
	}
}
