package gpsapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type GPSData struct {
	Latitude  float64 `json: "latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude"`
}
type Server struct {
	*mux.Router
}

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}
	s.routes()
	return s
}
func (s *Server) getGpsData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data GPSData
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Println("Data is received\nLatitude:", data.Latitude, "\nLongitude:", data.Longitude,"\nAltitude:",data.Altitude)
	}
}
func (s *Server) routes() {
	s.HandleFunc("/getGPS", s.getGpsData()).Methods("POST")
}
