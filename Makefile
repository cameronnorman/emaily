build:
	GOOS=linux GOARCH=amd64 go build cmd/email_sender/main.go
	mv main bin/emaily

	docker build -t emaily:0.0.7 .

run:
	docker run --rm --publish 8081:8081 --detach --name emaily emaily:0.0.7
