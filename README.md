# Fishing Derby App API

This is the API server for the Fishing Derby App, which allows users to manage fishing derbies, including user accounts, derby creation, joining derbies, posting results, and viewing leaderboards. The API uses Google OAuth2 for authentication.

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)
- [License](#license)

## Features

- User authentication via Google SSO
- Create and manage fishing derbies
- Join derbies and post results
- View derby leaderboards
- Search derbies by location
- Send notifications to users

## Requirements

- Go 1.18 or later
- Access to a Google Cloud project for OAuth2 setup
- [OpenAPI Generator](https://openapi-generator.tech/) for generating server code

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/fishing-derby-app.git
   cd fishing-derby-app
   ```

2. **Install dependencies:**

   Ensure you have Go modules enabled and run:

   ```bash
   go mod tidy
   ```

3. **Set up Google OAuth2:**

   - Create a Google Cloud project and set up OAuth2 credentials.
   - Update your environment variables with the client ID and secret.

4. **Generate server code (if needed):**

   If you need to regenerate server code from the OpenAPI spec:

   ```bash
   openapi-generator generate -i api-spec.yml -g go-server -o ./generated-server
   ```

## Usage

1. **Run the server:**

   ```bash
   go run main.go
   ```

2. **Access the API:**

   The server will be running at `http://localhost:8080`. You can use tools like Postman to interact with the API.

## API Endpoints

### Users
- **GET** `/api/users`: Get all users
- **POST** `/api/users`: Create a new user
- **GET** `/api/users/{id}`: Get a user by ID
- **PUT** `/api/users/{id}`: Update a user
- **DELETE** `/api/users/{id}`: Delete a user
- **GET** `/api/users/{id}/derbies`: Get all derbies a user has participated in
- **GET** `/api/users/{id}/catches`: Get all catches by a user

### Derbies
- **GET** `/api/derbies`: Get all derbies
- **POST** `/api/derbies`: Create a new derby
- **GET** `/api/derbies/{id}`: Get a derby by ID
- **PUT** `/api/derbies/{id}`: Update a derby
- **DELETE** `/api/derbies/{id}`: Delete a derby
- **GET** `/api/derbies/{id}/catches`: Get all catches in a derby
- **GET** `/api/derbies/{id}/participants`: Get all participants in a derby

### Catches
- **GET** `/api/catches`: Get all catches
- **POST** `/api/catches`: Record a new catch
- **GET** `/api/catches/{id}`: Get a catch by ID
- **PUT** `/api/catches/{id}`: Update a catch
- **DELETE** `/api/catches/{id}`: Delete a catch
- **GET** `/api/catches/by-user/{userId}`: Get all catches by a user
- **GET** `/api/catches/by-derby/{derbyId}`: Get all catches in a derby
- **GET** `/api/catches/by-location/{locationId}`: Get all catches at a specific location

### Locations
- **GET** `/api/locations`: Get all locations
- **POST** `/api/locations`: Create a new location
- **GET** `/api/locations/{id}`: Get a location by ID
- **PUT** `/api/locations/{id}`: Update a location
- **DELETE** `/api/locations/{id}`: Delete a location

## Example Requests

1. **Get All Users**
    ```bash
    curl -X GET http://localhost:3000/api/users
    ```

2. **Create a New User**
    ```bash
    curl -X POST http://localhost:3000/api/users \
    -H "Content-Type: application/json" \
    -d '{"name": "John Doe", "email": "john@example.com"}'
    ```

3. **Get All Catches in a Derby**
    ```bash
    curl -X GET http://localhost:3000/api/derbies/{derbyId}/catches
    ```


## Entities
[Diagram](https://app.excalidraw.com/s/9DFvsCgFtWB/3ZlpiEi6uj1)
![image](https://github.com/user-attachments/assets/88a969b0-6a4d-47d2-8399-33966d37b0e8)


## Contributing

Contributions are welcome! Please fork the repository and submit a pull request for any changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

