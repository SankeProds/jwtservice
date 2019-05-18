# JWTservice
A JWT service for a secret app I'm creating. There is no business logic here so I thought "why not share?"

Users will be able to register and log in with this service. The user of the other services of the app will require a JWT provided by this service `/login` endpoint.

this services state is not final. I'm now focused on other services for the app but this repo still has more versions/features in its future. 

# Prerequisites:

- A running Postgres DB.
- Set up your database
- Run `db/mig_*.sql` before using the service.

# Build the service:

```
go get
go build
```

# Running it:

The following script runs the service (read the contents to see configuration options):
```
./run.sh
```

The following script does the same thing but runs a `go build` before 
```
./buildandrun.sh
```

# two endpoints 

## [POST] /register

registers a user to use the other services. 

- id: a user id  
- data: a json with users data  
- authData: the data to use to verify the user identity  
  - AuthMethod is the only mandatory field of authdata. Only supported method right now is "password"

```
application/json
body:
{
    "id": "A_USER_ID", 
    "data": {
        "age": 33,
        "And": "Any other data you waned to store"
    },
    "authData": {
        "AuthMethod": "password",
	"Password": "1"
    }
}
```

## [POST] /login

logs in a user that is going to use the other services

- id: a user id
- loginData: the data to log the user in. This is compared/validated with the data at registration. fields depend on authMethod set at `/register`

```
application/json
body:
{
    "id": "A_USER_ID",
    "loginData": {
        "password":"1"
    }
}
```
returns
```
A VERIFIABLE JSON WEB TOKEN
```

# ROAD MAP

- DOCKER IMAGE
- JENKINS FILE
- [GET] /user endpoint
- tests
- probably more
