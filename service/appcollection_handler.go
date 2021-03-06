/*
Copyright 2016 Pivotal

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package service

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pivotalservices/app-usage-nozzle/usageevents"
	"github.com/unrolled/render"
)

func appCollectionHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", req.Header.Get("Origin"))
		w.Header().Add("Access-Control-Allow-Methods", "GET")
		formatter.JSON(w, http.StatusOK, usageevents.AppStats)
	}
}

func singleAppHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", req.Header.Get("Origin"))
		w.Header().Add("Access-Control-Allow-Methods", "GET")
		vars := mux.Vars(req)
		app := vars["app"]
		org := vars["org"]
		space := vars["space"]
		key := usageevents.GetMapKeyFromAppData(org, space, app)

		fmt.Printf("Retrieving app at key : %s\n", key)
		stat, exists := usageevents.AppStats[key]
		if exists {
			formatter.JSON(w, http.StatusOK, usageevents.CalculateDetailedStat(stat))
		} else {
			formatter.JSON(w, http.StatusNotFound, "No such app")
		}
	}
}
