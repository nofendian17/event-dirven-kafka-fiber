package repository

import "kafka-fibre/entitiy"

type ProduceRepository interface {
	Produce(event *entitiy.Event) error
}