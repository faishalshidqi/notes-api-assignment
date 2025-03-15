# Installation Guide
> [!CAUTION]
> You must have `go` installed. Use this command to ensure you have it installed.
```shell
go version
```
The terminal should print the go version installed on your machine.
```console
go version go1.24.1 windows/amd64
```

1. Clone this repository
```git clone https://github.com/faishalshidqi/notes-api-assignment```
2. Add a .env file
3. Add this environment variables to .env
   - SERVER_ADDRESS (should be filled with an address the server will be running from. e.g.: localhost:9000 or 0.0.0.0:5000)
   - CONTEXT_TIMEOUT (should be filled with an integer, context timeout is in seconds. ideally within 3-5 seconds.)
   - Access_TOKEN_KEY (should be filled with a hex for access token key)
   - ACCESS_TOKEN_AGE (should be filled with an integer, access token age is in hours)
   - DBUSER (should be filled with the database username)
   - DBPASSWORD (should be filled with the database password)
   - DBHOST (should be filled with the database hostname or IP address)
   - DBPORT (should be filled with the database port)
   - DBDatabase (should be filled with the database name)
4. This project implements database migration. Please install any database migration tool for golang
One of the tools you can use is [goose](https://github.com/pressly/goose). Install goose with this command ```go install github.com/pressly/goose/v3/cmd/goose@latest```
5. If your mysql instance authenticate using url + username + password. Migrate the database up with
```shell
goose mysql "<mysql username>:<mysql password>@tcp(<mysql ip>:<mysql port>)/<mysql name>" up
```
6. run this command to build and run notes api
```shell
go run main.go
```

