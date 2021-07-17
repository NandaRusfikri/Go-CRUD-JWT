
# API Spec Go-CRUD-JWT 


## Create User

Request :
- Method : POST
- Endpoint : `localhost:3000/user`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```yaml
form-data:
  name:nanda
  phone:12345678
  email:saya@nanda.dev
  username:nanda
  password:nanda

```

Response :

```json 
{
    "data": {
        "id": 3,
        "name": "nanda",
        "phone": "12345678",
        "email": "saya@nanda.dev",
        "username": "nanda",
        "password": "nanda",
        "CreatedAt": "2021-07-17 20:52:28.205326"
    }
}
```

## Login

Request :
- Method : POST
- Endpoint : `localhost:3000/login`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```yaml
form-data:
  username:nanda
  password:nanda

```

Response :

```json 
{
    "message": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.9xaLm7gtDvsvJ4wgmsAvZoOZfAgN7HUAc2htDY9hT2g",
    "data": [
        {
            "id": 3,
            "name": "kucing",
            "phone": "12345678",
            "email": "kucing@kucing.com",
            "username": "kucing",
            "gender": false
        }
    ]
}
```


## Get List User

Request :
- Method : GET
- Endpoint : `localhost:3000/users`
- Header :
    - Authorization: token
    - Accept: application/json

- Params :

```yaml
Query Params:
  limit: 2
  offset: 2
  name: "nanda"
  
```

Response :

```json 
{
    "count": 1,
    "data": [
        {
            "id": 1,
            "name": "nandarusfikri",
            "phone": "123456",
            "email": "nanda@nanda.com",
            "username": "bebek",
            "gender": false
        }
    ]
}
```



## Delete User

Request :
- Method : DELETE
- Endpoint : `localhost:3000/user?id=1`
- Header :
    - Authorization: token

Response :

```json 
{
    "data": {
        "id": 2,
        "DeletedAt": "2021-07-17 20:52:28"
    }
}
```


## Update User

Request :
- Method : PUT
- Endpoint : `localhost:3000/user`
- Header :
    - Authorization: token

- Body :

```yaml
form-data:
  name:nanda
  phone:12345678
  email:saya@nanda.dev
  username:nanda
  password:nanda

```

Response :

```json 
{
    "data": {
        "id": 3,
        "name": "nanda",
        "phone": "12345678",
        "email": "saya@nanda.dev",
        "username": "nanda",
        "password": "nanda",
        "CreatedAt": "2021-07-17 20:52:28.205326"
    }
}
```

## Setup Database
create database nandarusfikri 