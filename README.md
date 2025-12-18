
# Backend Task

this project contains an api service built using go fiber, postgresql and sqlc.
this serives manages user data and suports CRUD operations along with paginated support for listing users.

# Technologies Used
- Go Fiber
- PostgreSQL
- SQLC
- go-playground/validator 

# Project Structure
```
/cmd/server/main.go
/config/
/db/migrations/
/db/query/
/db/sqlc/<generated>
/internal/
├── handler/
├── repository/
├── service/
├── routes/
├── middleware/
├── models/
└── logger/
```

# Setup Instructions
1. Clone the repository
2. Set up PostgreSQL database
3. Configure environment variables in a .env file with database connection details
4. generate sqlc code using `sqlc generate`
5. Run the application using `go run cmd/server/main.go`

# API Endpoints
- POST /users - Create a new user
- GET /users/:id - Get user by ID
- PUT /users/:id - Update user by ID
- DELETE /users/:id - Delete user by ID
- GET /users - List users with pagination support

# Pagination
The List Users endpoint supports pagination through query parameters:
- page: The page number (default: 1)
- limit: The number of users per page (default: 10) 

The response includes the current page, limit, total number of users, and the list of users for the requested page.
# Example Request
GET /users?page=2&limit=5
```
Response:
{
  "page": 1,
  "limit": 10,
  "total": 1,
  "data": [
    {
      "id": 1,
      "name": "John Doe",
      "dob": "1990-05-15",
      "age": 33
    }
  ]
}
```



