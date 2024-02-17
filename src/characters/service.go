package characters

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	CharacterService struct {
		repository Repository
		logger     *log.Logger
	}

	Service interface {
		Create(data Character) (*Character, error)
		GetAll() ([]Character, error)
		Get(id string) (*Character, error)
		Delete(id string) error
		Update(id string, data *Character) (int, error)
	}
)

func NewCharacterService(clientDB *mongo.Client, logger *log.Logger) *CharacterService {
	repository := NewCharacterRepository(clientDB, logger)
	return &CharacterService{
		repository: repository,
		logger:     logger,
	}
}

func (service CharacterService) Create(character Character) (*Character, error) {
	service.logger.Println("[!] Getting into Create User Method on service layer")

	result, err := service.repository.Create(&character)
	if err != nil {
		service.logger.Println("[x] Something went wrong with the respository response")
		return nil, err
	}
	return result, nil
}

func (service CharacterService) GetAll() ([]Character, error) {
	service.logger.Println("[!] Getting into Get All Characters Method on service layer")

	result, err := service.repository.GetAll()
	if err != nil {
		service.logger.Println("[x] Something went wrong with the respository response")
		return nil, err
	}
	return result, nil
}

func (service CharacterService) Get(id string) (*Character, error) {
	service.logger.Println("[!] Getting into Get Character Method on service layer")

	result, err := service.repository.Get(id)
	if err != nil {
		service.logger.Println("[x] Something went wrong with the respository response")
		return nil, err
	}
	return result, nil
}

func (service CharacterService) Delete(id string) error {
	service.logger.Println("[!] Getting into Delete Character Method on service layer")

	err := service.repository.Delete(id)
	if err != nil {
		service.logger.Println("[x] Something went wrong with the respository response")
		return err
	}
	return nil
}

func (service CharacterService) Update(id string, character *Character) (int, error) {
	service.logger.Println("[!] Getting into Update Character Method on service layer")

	result, err := service.repository.Update(id, character)
	if err != nil {
		service.logger.Println("[x] Something went wrong with the respository response")
		return 0, err
	}
	return result, nil
}
