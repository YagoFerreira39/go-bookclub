package extensions

import (
	"github.com/YagoFerreira39/go-bookclub/src/domain/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookExtension struct{}

func (bookExtension *BookExtension) FromDatabaseResultToModel(result map[string]interface{}) models.BookModel {
	model := bookExtension.createBookModel(result)

	return model
}

func (bookExtension *BookExtension) FromDatabaseResultToModelList(resultList []map[string]interface{}) []models.BookModel {
	var models []models.BookModel
	for _, result := range resultList {
		model := bookExtension.createBookModel(result)

		models = append(models, model)
	}

	return models
}

func (bookExtension *BookExtension) createBookModel(result map[string]interface{}) models.BookModel {
	id, _ := primitive.ObjectIDFromHex(result["_id"].(string))
	model := models.BookModel{
		ID:        &id,
		Name:      result["name"].(string),
		Author:    result["author"].(string),
		ISBN:      result["isbn"].(string),
		Published: result["published"].(float64),
	}

	return model
}
