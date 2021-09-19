package service

import (
	"kafka-fibre/entitiy"
	"kafka-fibre/model"
	"kafka-fibre/repository"
)

func NewProduceService(produceRepository *repository.ProduceRepository) ProduceService {
	return &produceServiceImpl{
		ProduceRepository: *produceRepository,
	}
}

type produceServiceImpl struct {
	ProduceRepository repository.ProduceRepository
}

func (service *produceServiceImpl) Create(request model.EventRequest) (err error) {
	event := &entitiy.Event{
		Key:   request.Key,
		Value: request.Value,
	}
	err = service.ProduceRepository.Produce(event)
	if err != nil {
		return err
	}
	return nil
}
