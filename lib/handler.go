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
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func CreateServer() {
	fmt.Println("Start Server")
	router := mux.NewRouter()
	e := NewEndpoint()
	router.HandleFunc("/process", e.getAllProcesses).Methods("GET")
	router.HandleFunc("/process/{id}", e.getProcess).Methods("GET")
	router.HandleFunc("/data", e.getAllData).Methods("GET")
	router.HandleFunc("/rating", e.storeRating).Methods("PUT")
	c := cors.New(
		cors.Options{
			AllowedHeaders: []string{"Content-Type", "Authorization"},
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		})
	handler := c.Handler(router)
	logger := NewLogger(handler, "CALL")
	log.Fatal(http.ListenAndServe(":8080", logger))

}
