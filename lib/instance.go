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
	"github.com/SmartEnergyPlatform/marketplace-rating/lib/model"

	"math"

	"github.com/satori/go.uuid"
	"github.com/jinzhu/gorm"
	"errors"
)

func CreateInstance(req model.RatingRequest, userId string, throwEvent bool) (err error) {
	id := uuid.NewV4()
	instance := model.Instance{
		ID:         id,
		UserId:     userId,
		ObjectId:   req.ObjectId,
		ObjectType: req.ObjectType,
		Stars:      req.Stars,
	}
	tx := DB.Begin()
	tx.Where(model.Instance{UserId: userId, ObjectId: req.ObjectId, ObjectType: req.ObjectType}).Assign(model.Instance{Stars: req.Stars}).FirstOrCreate(&instance)
	if tx.Error != nil {
		return tx.Error
	}
	if throwEvent {
		if instance.ObjectType == Process {
			event, err := getRating(tx, instance)
			if err == nil {
				err = sendEvent(PROCESS_RATING_TOPIC, event)
			}
		}
	}
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return tx.Error
}

func getRating(tx *gorm.DB, instance model.Instance) (event model.RatingEvent, err error) {
	temp, err := calculateRatingFor(tx, instance.ObjectType, instance.ObjectId)
	if err != nil {
		return event, err
	}
	event = model.RatingEvent{
		Command: "PUT",
		Id:instance.ObjectId,
		Rating:temp.Rating,

	}
	return
}

func getUserAndProcessRating(objectType string, user string, id string) (userAndProcessRating model.UserAndObjectRating) {
	var userRating model.ResultUserRating

	DB.Table("instances").Select("stars").Where("object_type = ? AND user_id = ? AND object_id = ?", objectType, user, id).Scan(&userRating)
	userAndProcessRating.StarsUser = userRating.Stars

	var sqlResult model.Results
	var stars, count int = 0, 0
	var starRating [5]model.StarRating

	for i := 5; i > 0; i-- {
		index := 5 - i
		starRating[index].Star = i
		starRating[index].Rating = 0
		starRating[index].StarRatio = 0
	}

	DB.Table("instances").Select("object_id, stars, count(stars) as count ").Group("object_id, object_type, stars").Having("object_type = ? and object_id = ?", objectType, id).Scan(&sqlResult)

	if len(sqlResult) != 0 {
		for _, sqlRow := range sqlResult {
			stars += sqlRow.Stars * sqlRow.Count
			count += sqlRow.Count
			starRating[5-sqlRow.Stars].Rating = sqlRow.Count
		}

		var max int
		for i := 0; i < len(starRating); i++ {
			if max < starRating[i].Rating {
				max = starRating[i].Rating
			}
		}

		for i := 0; i < len(starRating); i++ {
			starRating[i].StarRatio = starRating[i].Rating * 100 / max
		}

		userAndProcessRating.StarRatings = starRating
		userAndProcessRating.StarsObject = math.Round(float64(stars)/float64(count)*10) / 10
		userAndProcessRating.Rating = count
	}

	return

}

func calculateRating(objectType string) (ergebnis model.ObjectRatings) {
	return calculateRatingDb(DB, objectType)
}

func calculateRatingDb(db *gorm.DB, objectType string) (ergebnis model.ObjectRatings) {
	var results model.Results
	var item model.ObjectRating
	db.Raw("select object_id, sum(stars) as stars, count(stars) as count from instances group by object_id, object_type having object_type = ?;", objectType).Scan(&results)
	if len(results) != 0 {
		for _, sqlRow := range results {
			item.Stars = math.Round(float64(sqlRow.Stars)/float64(sqlRow.Count)*10) / 10
			item.Rating = sqlRow.Count
			item.ObjectId = sqlRow.ObjectId
			ergebnis = append(ergebnis, item)
		}
	}
	return
}

func calculateRatingFor(db *gorm.DB, objectType string, objectId string)(result model.ObjectRating, err error){
	var results model.Results
	db.Raw("select sum(stars) as stars, count(stars) as count from instances where object_type = ? and object_id = ?;", objectType, objectId).Scan(&results)
	if len(results) != 0 {
		for _, sqlRow := range results {
			result.Stars = math.Round(float64(sqlRow.Stars)/float64(sqlRow.Count)*10) / 10
			result.Rating = sqlRow.Count
			return
		}
	}
	return result, errors.New("no rating found for "+objectType+" "+ objectId)
}