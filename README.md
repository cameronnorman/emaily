# Emaily

## Getting started

### Using binary
`./bin/emaily`

### Using Docker image



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
