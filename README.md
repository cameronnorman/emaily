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
### Send normal email with specifying the body
To send a email make a HTTP request like below

`POST http://localhost:8081/send`

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
### Send a template and specify the body
To send a email with a template make a HTTP request like below

`POST http://localhost:8081/send_with_template`

```
{
	"details": {
		"from": "your-email-address",
		"to": "email-address-you-sending-to",
		"subject": "email-subject",
		"body": "" // body is overridden with template 
	},
	"config": {
		"server": "your-smtp-server",
		"port": "server-port",
		"username": "your-username",
		"password": "your-password"
	}
	"template_name": "invoice",
	"data": {
		"logo": "your-cool-logo",
		"invoice_number": "invoice-number",
		"company_name": "company-name",
		"contact_email": "contact-email",
		"contact_number": "contact-number",
		"invoice_items": [
			{
				"name":"cool hat",
				"quantity":1,
				"price":"10",
				"total":"10"
			}
		],
		"total": "10",
		"payment_details": {
			"bank": "bank-name",
			"account_number": "account-number",
			"branch_code": "branch-code",
			"reference": "reference",
			"message": "Cool message to customer"
		}
	}
}
```