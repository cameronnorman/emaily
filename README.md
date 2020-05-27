# Emaily

## Getting started

### Using binary
`./bin/emaily`

### Using Docker image
`make build && make run`

### Using a different base docker image
```
FROM ubuntu:latest as dev
WORKDIR /usr/src/app
COPY bin/emaily /usr/local/bin/emaily
EXPOSE 8081
CMD /usr/local/bin/emaily
```

## Request example
To send a email make a HTTP request like below

`POST http://localhost:8081`

```
{
	"details": {
		"from": "your-email-address",
		"to": "email-address-you-sending-to",
		"subject": "email-subject",
		"body": "email-body-content" //html supported
	},
	"config": {
		"server": "your-smtp-server",
		"port": "server-port",
		"username": "your-username",
		"password": "your-password"
	}
}
```
