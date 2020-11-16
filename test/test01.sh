#!/usr/bin/env bash

curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"data":[
    "https://reqres.in/api/products/3",
    "http://date.jsontest.com/",
    "http://date.jsontest.com/",
    "http://date.jsontest.com/",
    "http://date.jsontest.com/",
    "http://date.jsontest.com/",
    "http://date.jsontest.com/",
    "http://date.jsontest.com/",
    "http://date.jsontest.com/",
    "http://date.jsontest.com/",
    "http://date.jsontest.com/",
    "http://date.jsontest.com/",
    "http://date.jsontest.com/",
    "http://date.jsontest.com/",
    "http://date.jsontest.com/",
    "http://date.jsontest.com/",
    "http://date.jsontest.com/"
    ]}' \
  http://localhost:8080/