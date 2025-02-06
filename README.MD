# go-microservice
Dummy microservice on Golang

## Starting server with Docker
```
docker build --no-cache -t go-microservice .
docker run -p 8095:8095 go-microservice
```

## Trying out API endpoints

Request to get current time in Halifax, Nova Scotia
```
GET http://localhost:8095/api/time-in-halifax
```
Response:
```
{"time_in_halifax":"2025-02-06 16:13:46 AST"}
```


Request to post some content 
```
POST http://localhost:8095/api/content -H "Content-Type: application/json" -d '{"text": "Hello, world!"}'
```

Successful response:
```
{"message":"Content processed successfully","status":"success"}
```

Failure:
```
{"message":"Something went wrong","status":"failure"}
```
