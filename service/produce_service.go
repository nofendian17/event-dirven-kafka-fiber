package service

import "kafka-fibre/model"

type ProduceService interface {
	Create(request model.EventRequest) (err error)
}