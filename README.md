# GOLANG ECHO FRAMEWORK REST API

## UI
- https://github.com/kapit4n/workout-plus

## Start it
- go run server.go

## Resources
- GET ```/```
- GET ```/categories```
- GET ```/intervals```
- GET ```/users```
- GET ```/users/:id```
- POST ```/users```
- PUT ```/users```
- DELETE ```/users```

## tools
curl --header "Content-Type: application/json" --request POST --data '{"fullname":"Luis Arce","email":"code1@gmail.com", "phoneNumber": "1231312", "picture": "https://i2.wp.com/wipy.tv/wp-content/uploads/2020/09/nuevos-personajes-de-avatar-2.jpg", "username": "code1", "password": "password123"}' http://localhost:1323/users
