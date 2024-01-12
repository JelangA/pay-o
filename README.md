Description Team:
##Techbyte Innovation
Krisna Purnama as Hacker
Jelang Anugrah Raharjo as Hacker
Hanif Ahmad Rizqullah as Hustler
Fanza Atsila Fizarli as Hipster
 
 Pay-O API Documentation

This documentation provides information about the Pay-O API, which handles user registration, login, and validation using JWT tokens.

## Table of Contents
- [Signup](#signup)
- [Login](#login)
- [Validate](#validate)

## Signup
Registers a new user with the system.

### Endpoint
`POST /signup`

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
}`

### Response Body
`json
{ "token": "your_generated_token" }`

###Failure (HTTP 400 Bad Request):
`json{
  "error" : "message",
}`

## Sigin
Registers a new user with the system.

### Endpoint
POST `/login`

### Request Body
- `Email` (string, required): User's email address.
- `Password` (string, required): User's password.

#### Example Request
`json
{
  "Email": "john@example.com",
  "Password": "password123",
}`
