# BingeBuster-MovieService
This service is part of the application BingeBuster, it's responsible for handling all the user functionality in the app. Like keeping track of watch history, liked movies or personal details. The service is written in the programming language Go and will be provided with its own database.


[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=Portfolio-Advanced-software_BingeBuster-MovieService&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=Portfolio-Advanced-software_BingeBuster-MovieService)


## Release V1.0.0
This release contains the basic code to send and receive messages through a RabbitMQ server. 


## Notes

To create proto files use ```protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/movie.proto```