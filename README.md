# Movie CRUD API

A simple RESTful API built with Go that provides CRUD (Create, Read, Update, Delete) operations for managing a movie collection.

## 🚀 Features

- **GET** `/movies` - Retrieve all movies
- **GET** `/movies/{id}` - Retrieve a specific movie by ID
- **POST** `/movies` - Create a new movie
- **PUT** `/movies/{id}` - Update an existing movie
- **DELETE** `/movies/{id}` - Delete a movie by ID

## 🛠️ Tech Stack

- **Language**: Go (1.24.4)
- **Router**: Gorilla Mux
- **Data Format**: JSON
- **Storage**: In-memory (slice)

## 📋 Prerequisites

- Go 1.24.4 or higher
- Git (for cloning the repository)

## 🔧 Installation

1. **Clone the repository**
   ```bash
   git clone <your-repo-url>
   cd CRUD-API
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Run the application**
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`

## 📡 API Endpoints

### Get All Movies
```http
GET /movies
```

**Response:**
```json
[
  {
    "id": "1",
    "isbn": "4311",
    "title": "F1",
    "director": {
      "firstname": "John",
      "lastname": "Cage"
    }
  }
]
```

### Get Movie by ID
```http
GET /movies/{id}
```

**Response:**
```json
{
  "id": "1",
  "isbn": "4311",
  "title": "F1",
  "director": {
    "firstname": "John",
    "lastname": "Cage"
  }
}
```

### Create New Movie
```http
POST /movies
```

**Request Body:**
```json
{
  "isbn": "1234",
  "title": "New Movie",
  "director": {
    "firstname": "Director",
    "lastname": "Name"
  }
}
```

**Response:**
```json
{
  "id": "generated-id",
  "isbn": "1234",
  "title": "New Movie",
  "director": {
    "firstname": "Director",
    "lastname": "Name"
  }
}
```

### Update Movie
```http
PUT /movies/{id}
```

**Request Body:**
```json
{
  "isbn": "5678",
  "title": "Updated Movie",
  "director": {
    "firstname": "Updated",
    "lastname": "Director"
  }
}
```

### Delete Movie
```http
DELETE /movies/{id}
```

**Response:** Returns the updated list of movies after deletion.

## 🧪 Testing the API

You can test the API using various tools:

### Using cURL

**Get all movies:**
```bash
curl -X GET http://localhost:8080/movies
```

**Get a specific movie:**
```bash
curl -X GET http://localhost:8080/movies/1
```

**Create a new movie:**
```bash
curl -X POST http://localhost:8080/movies \
  -H "Content-Type: application/json" \
  -d '{
    "isbn": "9999",
    "title": "Test Movie",
    "director": {
      "firstname": "Test",
      "lastname": "Director"
    }
  }'
```

**Update a movie:**
```bash
curl -X PUT http://localhost:8080/movies/1 \
  -H "Content-Type: application/json" \
  -d '{
    "isbn": "1111",
    "title": "Updated Movie",
    "director": {
      "firstname": "Updated",
      "lastname": "Director"
    }
  }'
```

**Delete a movie:**
```bash
curl -X DELETE http://localhost:8080/movies/1
```

### Using Postman

1. Import the endpoints into Postman
2. Set the base URL to `http://localhost:8080`
3. Test each endpoint with appropriate request bodies

## 📁 Project Structure

```
CRUD-API/
├── main.go        # Main application file with all handlers
├── go.mod         # Go module dependencies
├── go.sum         # Dependency checksums
└── README.md      # Project documentation
```

## 🏗️ Data Models

### Movie
```go
type Movie struct {
    ID       string    `json:"id"`
    Isbn     string    `json:"isbn"`
    Title    string    `json:"title"`
    Director *Director `json:"director"`
}
```

### Director
```go
type Director struct {
    FirstName string `json:"firstname"`
    LastName  string `json:"lastname"`
}
```

## 📚 Sample Data

The application comes pre-loaded with sample movies:

1. **F1** - Directed by John Cage
2. **Mad Max** - Directed by Peter Kevin
3. **End Game** - Directed by Kevin Fiege
4. **Thor Ragnarok** - Directed by Alex Hales
5. **The Revenant** - Directed by Alejandro González
6. **Interstellar** - Directed by Christopher Nolan

## 🔮 Future Improvements

- [ ] Add database integration (PostgreSQL, MongoDB)
- [ ] Implement proper error handling and validation
- [ ] Add authentication and authorization
- [ ] Implement pagination for large datasets
- [ ] Add unit and integration tests
- [ ] Add logging middleware
- [ ] Implement proper ID generation (UUID)
- [ ] Add API documentation with Swagger
- [ ] Add rate limiting
- [ ] Implement proper HTTP status codes

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📞 Contact

Your Name - your.email@example.com

Project Link: [https://github.com/yourusername/CRUD-API](https://github.com/yourusername/CRUD-API)

---

⭐ If you found this project helpful, please give it a star!