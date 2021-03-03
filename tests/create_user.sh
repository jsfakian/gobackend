#!/bin/bash

curl --location --request POST 'http://localhost:8765/api/users' \
--header 'Content-Type: application/json' \
--data '{
    "username": "admin",
    "password": "12345678",
    "email": "foo@foo.gr",
    "role": "admin",
    "company": "sfakasgr.corp",
    "phone": "123456789",
    "birthday": "10/12/2003"
}'
