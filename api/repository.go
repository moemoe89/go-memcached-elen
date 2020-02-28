//
//  Practicing Memcached
//
//  Copyright © 2020. All rights reserved.
//

package api

import (
	"github.com/bradfitz/gomemcache/memcache"
)

// Repository represent the repositories
type Repository interface {
	Set(key, val string) error
	Get(key string) (string, error)
	GetMulti(keys []string) (map[string]*memcache.Item, error)
	Delete(key string) error
}

type memcachedRepository struct {
	Client *memcache.Client
}

// NewMemcachedRepository will create an object that represent the Repository interface
func NewMemcachedRepository(Client *memcache.Client) Repository {
	return &memcachedRepository{Client}
}

func (m *memcachedRepository) Set(key, val string) (error) {
	err := m.Client.Set(&memcache.Item{Key: key, Value: []byte(val)})
	return err
}

func (m *memcachedRepository) Get(key string) (string, error) {
	val, err := m.Client.Get(key)
	if err != nil {
		return "", err
	}
	return string(val.Value), nil
}

func (m *memcachedRepository) GetMulti(keys []string) (map[string]*memcache.Item, error) {
	it, err := m.Client.GetMulti(keys)
	if err != nil {
		return nil, err
	}
	return it, nil
}

func (m *memcachedRepository) Delete(key string) error {
	err := m.Client.Delete(key)
	return err
}
