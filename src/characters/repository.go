package characters

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	CharacterRepository struct {
		ClientDB *mongo.Client
		Logger   *log.Logger
	}

	Repository interface {
		Create(data *Character) (*Character, error)
		GetAll() ([]Character, error)
		Get(id string) (*Character, error)
		Delete(id string) error
		Update(id string, data *Character) (int, error)
	}
)

func NewCharacterRepository(client *mongo.Client, logger *log.Logger) *CharacterRepository {
	return &CharacterRepository{
		ClientDB: client,
		Logger:   logger,
	}
}

func (repository *CharacterRepository) GetAll() ([]Character, error) {
	repository.Logger.Println("[!] Getting all characters")
	collection := repository.ClientDB.Database("kamehamehapi").Collection("characters")
	var characters []Character
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	cursor.All(context.TODO(), &characters)

	return characters, nil
}

func (repository *CharacterRepository) Get(id string) (*Character, error) {
	repository.Logger.Println("[!] Getting one character")
	filter := bson.D{{Key: "_id", Value: id}}
	var character Character
	collection := repository.ClientDB.Database("kamehamehapi").Collection("characters")
	err := collection.FindOne(context.TODO(), filter).Decode(&character)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
	}

	return &character, nil
}

func (repository *CharacterRepository) Create(data *Character) (*Character, error) {
	repository.Logger.Println("[!] Creating a character")
	collection := repository.ClientDB.Database("kamehamehapi").Collection("characters")
	result, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		return nil, err
	}

	data.ID = result.InsertedID.(primitive.ObjectID)

	return data, nil
}

func (repository *CharacterRepository) Update(id string, data *Character) (int, error) {
	repository.Logger.Println("[!] Updating a character")
	filter := bson.D{{Key: "_id", Value: id}}
	collection := repository.ClientDB.Database("kamehamehapi").Collection("characters")
	updated, err := collection.ReplaceOne(context.TODO(), filter, data)
	if err != nil {
		return 0, err
	}

	return int(updated.ModifiedCount), nil
}

func (repository *CharacterRepository) Delete(id string) error {
	repository.Logger.Println("[!] Deleting a character")
	filter := bson.D{{Key: "_id", Value: id}}
	collection := repository.ClientDB.Database("kamehamehapi").Collection("characters")
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}
