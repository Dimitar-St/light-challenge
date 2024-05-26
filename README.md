# light-challenge
light-challenge

#### Start database container using docker:

`docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=12345 postgres`

#### Start the app server:

Go to the back-end/bin folder via terminal and start one of the binaries

### Start the client:

Go to front-end folder and type `npm start`, everything should start successfully

#### If it does not work with the client you can try with this example curl command:

`curl --location 'localhost:8080/invoice' \
--header 'Content-Type: application/json' \
--data '{
  "amount": 10002,
  "manager_approval": false,
  "department": "marketing"
}
'`