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

package main

import (
	"log"
	"github.com/SmartEnergyPlatform/marketplace-rating/lib"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found", err)
	}
	err = lib.InitEventHandling()
	if err != nil {
		log.Fatal("unable to start amqp", err)
	}
	lib.DbInit()
	defer lib.DbClose()
	m := lib.NewMigration(lib.GetDB())
	m.Migrate()
	lib.CreateServer()

}
