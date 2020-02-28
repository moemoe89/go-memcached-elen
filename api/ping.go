//
//  Practicing Memcached
//
//  Copyright © 2020. All rights reserved.
//

package api

import (
	"encoding/json"
	"net/http"
	"time"
)

var starTime = time.Now()

// Ping will handle the ping endpoint
func Ping(w http.ResponseWriter, r *http.Request) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	startTime := starTime.In(loc)

	res := map[string]string{
		"start_time": startTime.Format("[02 January 2006] 15:04:05 MST"),
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
