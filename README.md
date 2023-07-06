# CRM API

This is a simple CRM (Customer Relationship Management) API written in Go. It provides basic CRUD (Create, Read, Update, Delete) operations for managing leads.

## Prerequisites

- Go programming language installed
- MySQL database server running
- Fiber and GORM libraries

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/crm-api.git
   ```

2. Change to the project directory:

   ```bash
   cd crm-api
   ```

3. Install the required dependencies:

   ```bash
   go get -u github.com/gofiber/fiber
   go get -u gorm.io/gorm
   go get -u gorm.io/driver/mysql
   ```

4. Configure the database connection in the `main.go` file:

   ```go
   dsn := "your-username:your-password@tcp(127.0.0.1:3306)/crm?charset=utf8mb4&parseTime=True&loc=Local"
   ```

   Replace `your-username` and `your-password` with your MySQL credentials.

5. Run the application:

   ```bash
   go run main.go
   ```

6. The API server should now be running on `http://localhost:3000`.

## API Endpoints

### Get all leads

```http
GET /api/v1/lead
```

Retrieves a list of all leads.

### Get a lead by ID

```http
GET /api/v1/lead/:id
```

Retrieves a specific lead by its ID.

### Create a new lead

```http
POST /api/v1/lead/:id
```

Creates a new lead. The lead data should be provided in the request body.

### Delete a lead

```http
DELETE /api/v1/lead/:id
```

Deletes a lead with the specified ID.

## Database Migration

The application uses the GORM library to manage the database schema. When the application starts, it automatically migrates the `Lead` model to the database. If the table doesn't exist, it will be created.

## Contributing

Contributions are welcome! If you find any issues or want to add new features, please submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
