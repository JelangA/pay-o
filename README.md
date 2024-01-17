Description Team:

`Techbyte Innovation`

Krisna Purnama as Hacker <br>
Jelang Anugrah Raharjo as Hacker <br>
Hanif Ahmad Rizqullah as Hustler <br>
Fanza Atsila Fizarli as Hipster <br>
 
 Pay-O API Documentation

This documentation provides information about the Pay-O API, which handles user registration, login, and validation using JWT tokens.

instalation 

clone `https://github.com/JelangA/pay-o` <br>
set `.env.example to .env` <br>
go to folder `cd pay-o` <br>
run command `compiledaemon --command="./pay-o"` <br>



## Table of Contents
- [Signup](#signup)
- [Login](#sigin)
- [GetDataUser](#tes)
- [Validate](#validate)

## Signup
Registers a new user with the system.

### Endpoint
`POST` `/signup`

### Request Body
- `Name` (string, required): User's name.
- `Email` (string, required): User's email address.
- `Password` (string, required): User's password.
- `Phone` (int, required): User's phone number.

#### Example Request
```json
{
  "Name": "John Doe",
  "Email": "john@example.com",
  "Password": "password123",
  "Phone": 123456789
}
```
### Response Body
```json
{ "token": "your_generated_token" }
```
###Failure (HTTP 401 Bad Request):
```json
{
  "error" : "message",
}
```

## Sigin
Registers a new user with the system.

### Endpoint
`POST` `/login`

### Request Body
- `Email` (string, required): User's email address.
- `Password` (string, required): User's password.

### Example Request
```json
{
  "Email": "john@example.com",
  "Password": "password123",
}
```
### Response Body
```json
{ "token": "your_generated_token" }
```
###Failure (HTTP 401 Bad Request):
```json
{
  "error" : "message",
}
```

## Tes
Get data user login

### Endpoint
`GET` `/tes`

### Request Header
```json
{
   "Authorization": "your_token"
}
```

### Response
```json
{
   "Data": "Data"
}
```

###Failure (HTTP 401 Bad Request):
```json
{
  "error" : "message",
}
```

###Forbiden (HTTP 403 Forbidden content)
redirect to forbidden content page 



