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
	"testing"
	"log"
	"github.com/joho/godotenv"
	"github.com/SmartEnergyPlatform/marketplace-rating/lib/model"
)

func checkErr(t *testing.T, err error){
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
}

func TestCalculateRatingFor(t *testing.T){
	err := godotenv.Load("./../.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	DbInit()
	defer DbClose()
	m := NewMigration(GetDB())
	m.Migrate()

	err = CreateInstance(model.RatingRequest{
		ObjectId:"a",
		ObjectType:Process,
		Stars:3,
	}, "test", false)
	checkErr(t, err)

	err = CreateInstance(model.RatingRequest{
		ObjectId:"a",
		ObjectType:Process,
		Stars:4,
	}, "test", false)
	checkErr(t, err)

	err = CreateInstance(model.RatingRequest{
		ObjectId:"a",
		ObjectType:Process,
		Stars:2,
	}, "asd", false)
	checkErr(t, err)

	result, err := calculateRatingFor(DB, Process, "a")
	checkErr(t, err)
	if !(result.Stars == 3.0 && result.Rating == 2) {
		t.Fatal(result)
	}

}
