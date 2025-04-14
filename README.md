
# File CRUD API

This is a **File CRUD** project developed using **Golang**. The API allows performing **create**, **read**, and **delete** operations on files via **HTTP**. The application uses **Docker** to simplify the development environment and ensure the **PostgreSQL** database is correctly configured.

## Features

The API provides the following routes:

- **GET** `/files/` - Lists all stored files.
- **GET** `/files/:id` - Retrieves a specific file by ID.
- **DELETE** `/files/:id` - Deletes a specific file by ID.
- **POST** `/files/` - Uploads a new file.

## Technologies Used

- **Golang** - The main programming language.
- **Docker** - For container creation and orchestration.
- **PostgreSQL** - The database used to store file information.

## Prerequisites

Before running the application, make sure you have the following installed:

- **Docker**: [Installation Link](https://www.docker.com/get-started)
- **Docker Compose**: [Installation Link](https://docs.docker.com/compose/install/)

## How to Run

### 1. Clone the repository:

```bash
git clone https://github.com/LeonardoAValerio/api-files.git
cd api-files
```

### 2. Build and start the containers with Docker Compose:

```bash
docker-compose up --build
```

This command will:
- Build the Go application image.
- Start the PostgreSQL database container.
- Start the Go application, which will be available at `http://localhost:8080`.

### 3. Testing the Routes

After running the command above, you can test the routes in your browser or using tools like **Postman** or **cURL**:

- **GET** `/files/` - Lists all stored files.
- **GET** `/files/:id` - Retrieves a specific file by ID.
- **DELETE** `/files/:id` - Deletes a specific file by ID.
- **POST** `/files/` - Uploads a new file.

## Project Structure

The structure of your project is as follows:

```
api-files/
│
├── docker-compose.yml        # Configuration to orchestrate containers
├── Dockerfile                # File to build the application image
├── go.mod                    # Go dependencies
├── go.sum                    # Go dependency checksums
├── main.go                   # Main application file
├── models/                   # Directory for data models (e.g., File)
├── controllers/              # Directory for route handlers (controllers)
├── database/                 # Directory for database creation and configurations
├── repository/               # Directory for repository code for each data model
└── README.md                 # This file
```

## Contributing

Feel free to open **issues** or **pull requests** if you want to suggest improvements or fix bugs.