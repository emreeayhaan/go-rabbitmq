1) docker run -d --hostname my-rabbit --name rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management

2) http://localhost:15672/#/ 

3) cd producer
go run producer.go

4) cd consumer
go run consumer.go