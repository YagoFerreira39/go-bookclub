package book_repository

// import (
// 	"log"

// 	"github.com/YagoFerreira39/go-bookclub/src/domain/models"
// 	"github.com/YagoFerreira39/go-bookclub/src/externals/infrastructure/mongodb"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// type BookRepository struct {
// 	MongoDBInfrastructure *mongodb.MongoDBInfrastructure
// }

// func (repository *BookRepository) CreateBook(model *models.BookModel) (*models.BookModel, error) {
// 	client := repository.MongoDBInfrastructure.GetMongoDBClient()
// 	ctx := repository.MongoDBInfrastructure.GetContext()
// 	collection := client.Database("BookClub").Collection("Books")

// 	// Convert the model to a BSON document
// 	documentToInsert, err := bson.Marshal(model)
// 	if err != nil {
// 		log.Printf("Error marshaling BookModel: %v", err)
// 		return model, err
// 	}

// 	result, err := collection.InsertOne(ctx, documentToInsert)
// 	if err != nil {
// 		log.Printf("Unable to insert document: %v.", err)
// 		return model, err
// 	}

// 	// Convert the InsertedID to primitive.ObjectID
// 	insertedID, ok := result.InsertedID.(primitive.ObjectID)
// 	if !ok {
// 		log.Println("Error converting InsertedID to primitive.ObjectID")
// 		return model, err
// 	}

// 	model.ID = &insertedID

// 	return model, nil
// }

// func (repository *BookRepository) GetBookById(bookId string) (*models.BookModel, error) {
// 	client := repository.MongoDBInfrastructure.GetMongoDBClient()
// 	ctx := repository.MongoDBInfrastructure.GetContext()
// 	collection := client.Database("BookClub").Collection("Books")

// 	var modelResult models.BookModel
// 	docId, _ := primitive.ObjectIDFromHex(bookId)
// 	filter := bson.M{"_id": docId}

// 	error := collection.FindOne(ctx, filter).Decode(&modelResult)
// 	if error != nil {
// 		log.Println("Fail to retrieve book from database.")
// 	}

// 	return &modelResult, error
// }
