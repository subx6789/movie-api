# Movie API

A simple RESTful API for managing movies built using Go (Golang) with features like getting movie details, creating, updating, and deleting movies. It uses the Gorilla Mux package for routing and UUID for unique ID generation.

## Features

- Get all movies
- Get a specific movie by ID
- Add a new movie
- Update an existing movie by ID
- Delete a movie by ID
- Error handling for invalid inputs and non-existent resources

## Tech Stack

- **Go (Golang)**: Backend programming language.
- **Gorilla Mux**: HTTP request router and dispatcher for matching incoming requests to their handlers.
- **UUID**: For generating unique IDs.
- **godotenv**: For loading environment variables from a `.env` file.

## Prerequisites

Make sure you have the following installed:

- Go (v1.16+)
- [Git](https://git-scm.com/)

You can download Go from [here](https://go.dev/dl/).

## Project Setup

### 1. Clone the Repository

```bash
 git clone https://github.com/subx6789/movie-api.git
 cd movie-api
```

### 2. Install Dependencies

Run the following command to install the required Go packages:

```bash
 go get -u github.com/gorilla/mux
 go get -u github.com/google/uuid
 go get -u github.com/joho/godotenv
```

### 3. Set up the Environment

Create a .env file in the root of the project and specify the port for the API server:

```bash
 PORT=8080
```

If you don't specify a port, the server will default to 8080.

### 4. Run the Application

To start the server, run:

```bash
 go run main.go
```

You should see:

```bash
 Server starting on port 8080
```

The server will be available at http://localhost:8080

### 5. Endpoints

- Get All Movies

GET `/movies`: Returns a list of all movies in JSON format.

- Get Movie by ID

GET `/movie/{id}`: Returns the movie with the specified ID.

- Add a New Movie

POST `/movies`: Request Body:

```json
{
  "isbn": "123456",
  "title": "Movie Title",
  "overview": "Movie overview description",
  "director": {
    "firstName": "Director First Name",
    "lastName": "Director Last Name"
  }
}
```

- Update a Movie by ID

PUT `/movie/{id}`: You can partially update any of the movie fields by specifying them in the request body.

Example Request Body:

```json
{
  "title": "Updated Movie Title",
  "director": {
    "firstName": "Updated First Name"
  }
}
```

- Delete a Movie by ID

DELETE `/movie/{id}`: Deletes the movie with the specified ID.

### 6. Error Handling

- **400 Bad Request:** Returned when invalid input is provided in the request body (e.g., missing required fields).
- **404 Not Found:** Returned when the movie with the specified ID is not found.
- **500 Internal Server Error:** Returned for any server-side issues, such as JSON encoding errors.

## Contributing

If you'd like to contribute, please open an issue or submit a pull request.

- Fork the repository
- Create your feature branch (`git checkout -b feature/my-feature`)
- Commit your changes (`git commit -m 'Add some feature'`)
- Push to the branch (`git push origin feature/my-feature`)
- Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

## Project deployed with Render

- **Link:** [Movie-API](https://movie-api-0ckg.onrender.com)
