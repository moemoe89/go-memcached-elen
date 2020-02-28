//
//  Practicing Memcached
//
//  Copyright © 2020. All rights reserved.
//

package config

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

// InitMemcached will create a variable that represent the memcache.Client
func InitMemcached() (*memcache.Client, error) {
	client := memcache.New(Configuration.Memcached.URL)
	err := client.Ping()
	if err != nil {
		return nil, fmt.Errorf("Failed to ping connection to memcached: %s", err.Error())
	}

	return client, nil
}
