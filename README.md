# Finchat payments API
## Simple application to demonstrate Golang API development with Stripe integration skills
Containts endpoints to create new customers, create payments and retreive all payments by customer.
This application is based on Fiber library.

### How to run
- copy .env.dist to .env and set stripe secret key and other parameters.
- go run finchat.go
OR
- go build && ./finchat

### Design Decisions
Error handling:
Uses similar http codes as stripe in response.
By default returning same stripe error object from handler results in HTTP 500 status.
Therefore, the handlers return new error by copying stripe status code with error message.
This way invalid key or bad request result in HTTP status 4xx rather than 500.

Uses middlware/recover to handle any panic calls as they are made.
- Time taken to develop: 5:40 hours.
