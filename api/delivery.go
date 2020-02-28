//
//  Practicing Memcached
//
//  Copyright © 2020. All rights reserved.
//

package api

import (
	"github.com/moemoe89/practicing-memcached-golang/api/api_struct/form"
	"github.com/moemoe89/practicing-memcached-golang/api/api_struct/model"
	cons "github.com/moemoe89/practicing-memcached-golang/constant"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// ctrl struct represent the delivery for controller
type ctrl struct {
	log *logrus.Entry
	svc Service
}

// NewCtrl will create an object that represent the ctrl struct
func NewCtrl(log *logrus.Entry, svc Service) *ctrl {
	return &ctrl{log, svc}
}

func (c *ctrl) Set(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req form.SetForm
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.NewGenericResponse(http.StatusInternalServerError, cons.ERR, []string{err.Error()}, nil))
		return
	}

	err = json.Unmarshal(reqBody, &req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.NewGenericResponse(http.StatusInternalServerError, cons.ERR, []string{err.Error()}, nil))
		return
	}

	errs := req.Validate()
	if len(errs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.NewGenericResponse(http.StatusBadRequest, cons.ERR, []string{err.Error()}, nil))
		return
	}

	status, err := c.svc.Set(req.Key, req.Value)
	if err != nil {
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(model.NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(model.NewGenericResponse(http.StatusCreated, cons.ERR, []string{"Data has been saved"}, nil))
}

func (c *ctrl) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	key := mux.Vars(r)["key"]

	value, status, err := c.svc.Get(key)
	if err != nil {
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(model.NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.NewGenericResponse(http.StatusOK, cons.ERR, []string{"Data has been retrieved"}, value))
}

func (c *ctrl) GetMulti(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	keysParam := mux.Vars(r)["keys"]
	keys := strings.Split(keysParam, ",")

	values, status, err := c.svc.GetMulti(keys)
	if err != nil {
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(model.NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.NewGenericResponse(http.StatusOK, cons.ERR, []string{"Data has been retrieved"}, values))
}

func (c *ctrl) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	key := mux.Vars(r)["key"]

	_, status, err := c.svc.Get(key)
	if err != nil {
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(model.NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}

	status, err = c.svc.Delete(key)
	if err != nil {
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(model.NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.NewGenericResponse(http.StatusOK, cons.ERR, []string{"Data has been deleted"}, nil))
}