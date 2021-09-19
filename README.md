# Simple implementation kafka & go

## system req
- apache kafka
- go 1.6 newer

## installation

- golang brew

```
$ brew install go
```

- kafka brew
```
$ brew cask install java
$ brew install kafka
$ brew services start kafka
$ brew services start zookeeper
```
- kafka configuration
```
$ vim /usr/local/etc/kafka/server.properties 
```
- add **auto.create.topics.enable=true**
- restart kafka (``` $ brew services restart kafka```)

- create new topic with 3 partition (assumption has 3 app consumer)
```
$ kafka-topics --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 3 --topic test-topic
```

- run application
```
$ cd project-dir
$ go run main.go serve
```

- test rest api
```
$ curl --location --request POST 'http://127.0.0.1:3000/api/event' \
--header 'Content-Type: application/json' \
--data-raw '{
    "key" :"events.kafka.create",
    "value" : "this is value"
}'
```

- consume event and run 3 consumer
```
$ cd project dir
$ go run main.go consume
```
