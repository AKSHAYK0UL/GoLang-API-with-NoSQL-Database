GoLang API with JSON Database

Description

This project is a GoLang API that serves as a simple backend system for managing courses. It utilizes a JSON file as its database backend, providing endpoints for performing CRUD (Create, Read, Update, Delete) operations on course data.

The API is built using the Go programming language and the Gorilla Mux router for handling HTTP requests. It follows RESTful principles, with each endpoint corresponding to a specific action on the course data.

Features:

Create: Add new courses to the database.

Read: Retrieve information about all courses or a specific course by its ID.

Update: Modify existing course data.

Delete: Remove courses from the database, either individually or all at once.

The API supports JSON format for data exchange, making it easy to integrate with frontend or other backend systems. It also includes functionality to keep the server alive by periodically sending requests to a specified URL.

Endpoints:

GET /: Home page.

GET /get: Retrieve all courses.

POST /post: Add a new course.

GET /search/{id}: Search for a course by its ID.

PUT /update/{id}: Update an existing course.

DELETE /del/{id}: Delete a course by its ID.

DELETE /del: Delete all courses from the database.

Database:

The API uses a JSON file (courses.json) as its database. Each course is stored as a JSON object in the file, allowing for easy data management and persistence.
