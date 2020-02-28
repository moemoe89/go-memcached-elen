//
//  Practicing Memcached
//
//  Copyright © 2020. All rights reserved.
//

package routers

import (
	ap "github.com/moemoe89/practicing-memcached-golang/api"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// GetRouter will create a variable that represent the gin.Engine
func GetRouter(log *logrus.Entry, svc ap.Service) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", ap.Ping).Methods("GET")
	r.HandleFunc("/ping", ap.Ping).Methods("GET")

	ctrl := ap.NewCtrl(log, svc)

	r.HandleFunc("/set", ctrl.Set).Methods("POST")
	r.HandleFunc("/get/{key}", ctrl.Get).Methods("GET")
	r.HandleFunc("/get/multi/{keys}", ctrl.GetMulti).Methods("GET")
	r.HandleFunc("/delete/{key}", ctrl.Delete).Methods("DELETE")

	return r
}
