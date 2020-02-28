//
//  Practicing Memcached
//
//  Copyright © 2020. All rights reserved.
//

package main

import (
	ap "github.com/moemoe89/practicing-memcached-golang/api"
	conf "github.com/moemoe89/practicing-memcached-golang/config"
	"github.com/moemoe89/practicing-memcached-golang/routers"

	"fmt"
	"net/http"
)

func main() {
	client, err := conf.InitMemcached()
	if err != nil {
		panic(err)
	}

	log := conf.InitLog()

	repo := ap.NewMemcachedRepository(client)
	svc := ap.NewService(log, repo)

	app := routers.GetRouter(log, svc)
	err = http.ListenAndServe(":" + conf.Configuration.Port, app)
	if err != nil {
		panic(fmt.Sprintf("Can't start the app: %s", err.Error()))
	}
}
