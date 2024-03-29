build:
	docker build -t emaily:0.0.13 .

compile:
	GOOS=linux GOARCH=amd64 go build cmd/email_sender/main.go
	mv main bin/emaily

run:
	docker run --rm --publish 8081:8081 --name emaily emaily:0.0.13
