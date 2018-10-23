/*
 *    Copyright 2018 InfAI (CC SES)
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package lib

import (
	"encoding/json"
	"fmt"
	"github.com/SmartEnergyPlatform/marketplace-rating/lib/model"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

const Process = "process"
const Data = "data"

type Endpoint struct {
}

func NewEndpoint() *Endpoint {
	return &Endpoint{}
}

func (e *Endpoint) getAllProcesses(w http.ResponseWriter, r *http.Request) {
	ratings := calculateRating(Process)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(ratings)
}

func (e *Endpoint) getProcess(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	userId := getUserId(r)
	userRating := getUserAndProcessRating(Process, userId, id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(userRating)
}

func (e *Endpoint) getAllData(w http.ResponseWriter, r *http.Request) {
	ratings := calculateRating(Data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(ratings)
}

func (e *Endpoint) storeRating(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var ratingReq model.RatingRequest
	err := decoder.Decode(&ratingReq)
	if err != nil {
		fmt.Println(err)
	}
	userId := getUserId(r)
	defer r.Body.Close()
	CreateInstance(ratingReq, userId, true)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(model.Response{"OK"})

}

func getUserId(r *http.Request) (userId string) {
	userId = r.Header.Get("X-UserId")
	if userId == "" {
		userId = "testUser"
	}
	strings.Replace(userId, "\"", "", -1)
	return
}
