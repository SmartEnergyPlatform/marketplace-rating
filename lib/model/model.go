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

package model

import "github.com/satori/go.uuid"

type ObjectRating struct {
	ObjectId string  `json:"objectId,omitempty"`
	Stars    float64 `json:"stars,omitempty"`
	Rating   int     `json:"rating,omitempty"`
}
type ObjectRatings []ObjectRating

type Instance struct {
	ID         uuid.UUID `gorm:"type:char(36);column:id"`
	UserId     string    `gorm:"primary_key;type:varchar(255)"`
	ObjectId   string    `gorm:"primary_key;type:varchar(255)"`
	ObjectType string    `gorm:"primary_key;type:varchar(10)"`
	Stars      int       `gorm:"type:decimal(1,0)"`
}

type Instances []Instance

type RatingRequest struct {
	ObjectId   string `json:"objectId,omitempty"`
	ObjectType string `json:"objectType,omitempty"`
	Stars      int    `json:"stars,omitempty"`
}

type Result struct {
	ObjectId string
	Stars    int
	Count    int
}

type Results []Result

type Response struct {
	Message string `json:"message,omitempty"`
}

type ResultUserRating struct {
	Stars int
}

type UserAndObjectRating struct {
	StarsUser   int           `json:"starsUser"`
	ObjectId    string        `json:"objectId,omitempty"`
	StarsObject float64       `json:"starsObject,omitempty"`
	Rating      int           `json:"rating,omitempty"`
	StarRatings [5]StarRating `json:"starRatings,omitempty"`
}

type StarRating struct {
	Star      int `json:"star,omitempty"`
	Rating    int `json:"rating"`
	StarRatio int `json:"starRatio"`
}

type RatingEvent struct {
	Command string `json:"command"`
	Id		string	`json:"id"`
	Stars   float64	`json:"stars"`
	Rating	int		`json:"rating"`
}