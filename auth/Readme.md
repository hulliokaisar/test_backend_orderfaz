# Simple Authentication App

This is a simple authentication application built with Go Programming Language.

## Description

The application provides the following functionalities:

- User registration with unique username and MSISDN (phone number)
- Password encryption and storage in the database
- User login with MSISDN and password
- Generation of JWT token upon successful login

## Requirements

- Go Programming Language
- MySQL Database

## Installation

1. Clone the repository:
git clone https://github.com/hulliokaisar/authentication-app.git

2. Set up the MySQL database with the following configurations:
   - Host: localhost
   - Port: 3306
   - Username: root
   - Password: [your-password]
   - Database Name: orderfaz1

3. Build the Go application and create a Docker image:

4. Run the Docker container:
docker run -p 8080:8080 -d authentication-app


5. Access the application at http://localhost:8080 in your web browser.

## Usage

- Registration: Fill in the registration form with your name, username, MSISDN, and password. Click "Register" to create a new user.
- Login: Fill in the login form with your MSISDN and password. Click "Login" to authenticate and obtain a JWT token.

## License

This project is licensed under the [MIT License](LICENSE).
