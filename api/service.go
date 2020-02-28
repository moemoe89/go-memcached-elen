//
//  Practicing Memcached
//
//  Copyright © 2020. All rights reserved.
//

package api

import (
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/bradfitz/gomemcache/memcache"
)

// Service represent the services
type Service interface {
	Set(key, value string) (int, error)
	Get(key string) (string, int, error)
	GetMulti(keys []string) (map[string]string, int, error)
	Delete(key string) (int, error)
}

type implService struct {
	log        *logrus.Entry
	repository Repository
}

// NewService will create an object that represent the Service interface
func NewService(log *logrus.Entry, r Repository) Service {
	return &implService{log: log, repository: r}
}

func (s *implService) Set(key, value string) (int, error) {
	err := s.repository.Set(key, value)
	if err != nil {
		s.log.Errorf("can't set data: %s", err.Error())
		return http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	return 0, nil
}

func (s *implService) Get(key string) (string, int, error) {
	val, err := s.repository.Get(key)
	if  err == memcache.ErrCacheMiss {
		return val, http.StatusNotFound, errors.New("Key not found")
	}

	if err != nil {
		s.log.Errorf("can't get data: %s", err.Error())
		return val, http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	return val, 0, nil
}

func (s *implService) GetMulti(keys []string) (map[string]string, int, error) {
	datas, err := s.repository.GetMulti(keys)
	if err != nil {
		s.log.Errorf("can't get multi data: %s", err.Error())
		return nil, http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	out := map[string]string{}
	for k, v := range datas {
		out[k] = string(v.Value)
	}

	return out, 0, nil
}

func (s *implService) Delete(key string) (int, error) {
	err := s.repository.Delete(key)
	if err != nil {
		s.log.Errorf("can't delete data: %s", err.Error())
		return http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	return 0, nil
}