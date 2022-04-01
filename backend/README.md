#To get started with the frontend

###please do the tasks accordingly

#### ** please install docker and postgres to get started then follow the points below

1 - Clone this repository then navigate backend folder from your terminal

2 - Type `make postgres` to create container from postgres image  

3 - Type `make createdb` to create database name `products`  

4 - Type `make migrateup` to make migration on postgres database which creates your database  

5 - After all the steps above please run `go run main.go` 

6 - Once it is run then the API is ready on http://localhost:8080/

7 - Finally open http://localhost:3000/ on your browser or use postman to access the data

enjoy...
