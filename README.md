# JWTservice
a JWT service

# Prerequisites:

A running postgres DB. Run `db/mig_*.sql` before using the service.

# Build the service:

go get
go build

# Running it:

The following script runs the service (read the contents to see configuration options):
```
./run.sh
```

The followin script does the same thing but runs a `go build` before 
```
./buildandrun.sh
```

# two endpoints 

## [POST] /register

registers an user to use the other services. 

- id: a user id  
- data: a json with users data  
- authData: the data to use to verify the user identity  
  - AuthMethod is the only mandatory field of authdata. Only supported method rigth now is "password"

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

logs in a use that is going to use the other services

- id: a user id
- loginData: the data to log the user in. This is compared/validated with the data at registration. fields depends on authMethod set at `/register`

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
<JWT>
```
