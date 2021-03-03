#!/bin/bash

token=$(curl --location --request POST 'http://localhost:8765/api/login' \
--header 'Content-Type: application/json' \
--data '{
    "username": "superuser",
    "password": "NexVisor"
}' 2> /dev/null | jq -r '.token')

echo $token