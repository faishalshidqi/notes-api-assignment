# Installation Guide
> [!CAUTION]
> You must have `go`, `docker`, and `npm` installed. Use this command to ensure you have it installed.
```shell
go version;docker -v;npm -v
```
The terminal should print the version of go and docker installed on your machine.
```console
go version go1.24.1 windows/amd64
Docker version 28.0.1, build 068a01e
10.7.0
```

1. Clone this repository
```git clone https://github.com/faishalshidqi/notes-api-assignment```
2. Add a .env file
3. Add this environment variables to .env
   - SERVER_ADDRESS (should be filled with a port number in ":<port number>" format. e.g. :5000 if you want to run it on port 5000)
   - CONTEXT_TIMEOUT (should be filled with an integer, context timeout is in seconds. ideally within 3-5 seconds.)
   - ACCESS_TOKEN_KEY (should be filled with a hex for access token key)
   - ACCESS_TOKEN_AGE (should be filled with an integer, access token age is in hours)
   - MYSQL_ROOT_USER (should be filled with the database username)
   - MYSQL_ROOT_PASSWORD (should be filled with the database password of your choosing)
   - DBHOST (should be filled with the database hostname or IP address. Fill with host.docker.internal if you run with docker)
   - DBPORT (should be filled with the database port. Fill with 3307 if you run with docker)
   - MYSQL_DATABASE (should be filled with the database name)
4. run this command to build and run notes api
```shell
docker compose up -d
```
5. This project implements database migration. Please install any database migration tool for golang
One of the tools you can use is [goose](https://github.com/pressly/goose). Install goose with this command ```go install github.com/pressly/goose/v3/cmd/goose@latest```
6. Navigate to the infrastructures/sql/schema folder.
7. If your mysql instance authenticates using url + username + password. Migrate up with
```shell
goose mysql "<mysql username>:<mysql password>@tcp(<mysql ip>:<mysql port>)/<mysql name>" up
```
For example, if you run this with docker compose. Use this command,
```shell
goose mysql "root:<the MYSQL_ROOT_PASSWORD value>@tcp(localhost:3307)/notes_api" up
```
8. Import [Notes API Test.postman_collection.json](Notes%20API%20Test.postman_collection.json) and [Notes API Test.postman_environment.json](Notes%20API%20Test.postman_environment.json) to Postman so you can test the API. And don't forget to switch the environment on before sending requests.
> [!IMPORTANT]
> Swagger UI is only visible at http://localhost:5000/docs/index.html if the API isn't run with Docker. To see swagger ui, just `go run main.go`. Don't forget to adjust the env vars.

> [!IMPORTANT]
> swagger.json documentation is in the docs folder.

> [!IMPORTANT]
> frontend can be visited at http://localhost:5000
