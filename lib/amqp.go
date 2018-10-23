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
	"log"
	"encoding/json"
	"github.com/SmartEnergyPlatform/amqp-wrapper-lib"
)

var AmqpConn *amqp_wrapper_lib.Connection

var PROCESS_RATING_TOPIC = "processrating"

func InitEventHandling() (err error) {
	amqpurl := GetEnv("AMQP_URL", "")
	AmqpConn, err = amqp_wrapper_lib.Init(amqpurl, []string{PROCESS_RATING_TOPIC}, 10)
	return
}

func sendEvent(topic string, event interface{}) error {
	payload, err := json.Marshal(event)
	if err != nil {
		log.Println("ERROR: event marshaling:", err)
		return err
	}
	log.Println("DEBUG: send amqp event: ", topic, string(payload))
	return AmqpConn.Publish(topic, payload)
}

