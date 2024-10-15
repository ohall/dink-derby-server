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

Refer to the `api-spec.yml` file for detailed API documentation. Key endpoints include:

- `GET /auth/google`: Initiates Google SSO login
- `GET /auth/google/callback`: Handles Google OAuth2 callback
- `GET /users/{userId}`: Retrieves user profile
- `POST /derbies`: Creates a new derby
- `POST /derbies/{derbyId}/join`: Joins a derby
- `GET /derbies/{derbyId}`: Retrieves derby details
- `POST /derbies/{derbyId}/results`: Posts derby results
- `GET /derbies/{derbyId}/leaderboard`: Retrieves derby leaderboard
- `GET /derbies/search`: Searches derbies by location
- `POST /notifications`: Sends notifications to users

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request for any changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

