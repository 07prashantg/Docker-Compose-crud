Overview

This project implements a simple CRUD (Create, Read, Update, Delete) application using Golang, SQL, Docker and Docker-Compose. It provides an interface to create, retrieve, update, and delete records in a database. The application is containerized using Docker and launched using Docker Compose.The API can be tested using Postman.



Tech-Stack Used

    - Golang
    - SQL
    - Docker
    - Docker Compose
    - Docker Hub
    - PostMan
    

Features

This project implements the following features:

- Create a new record in the database
- Read a record from the database
- Update a record in the database
- Delete a record from the database



Testing the API with Postman

To test the API, you can use Postman. The endpoints for the API are:
- "/" for home page
- GET /users to create a new record
- POST /detailById to retrieve a record using Id
- POST /insert to insert a new record
- POST /update to update a record
- POST /deleteById to delete a record by Id

You can use Postman to send HTTP requests to these endpoints and verify that the API is working correctly.


Conclusion

This project demonstrates how to create a simple CRUD application using Golang and SQL, containerize it using Docker, and launch it using Docker Compose. You can use this as a starting point for building more complex applications.
