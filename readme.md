# Event Booking API

This project is an Event Booking API built with Go. It provides functionality for CRUD operations on events, user authentication, and user registration for events. The project uses Gin for the HTTP web framework, JWT for JSON Web Tokens, and SQLite for the database.
Will add soon a frontend app in NextJS to combine it with this api. 🤞

## Features

- **CRUD Operations for Events**: Create, Read, Update, and Delete events.
- **User Authentication**: Secure user authentication using JWT.
- **User Registration for Events**: Users can register for events.

## Technologies Used

- **Go**: Programming language.
- **Gin**: HTTP web framework.
- **JWT**: JSON Web Tokens for authentication.
- **SQLite**: Lightweight database for storage.
- **Bcrypt**: Password hashing.

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/Badara-Senpai/event-booking-api.git
    cd event-booking-api
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Run the application:
    ```bash
    go run main.go
    ```

## API Endpoints

### Authentication

- **Register**: `POST /signin`
    - Request Body:
      ```json
      {
          "full_name": "ABC XYZ",
          "email": "example@example.com",
          "password": "*********"
      }
      ```
    - Response: 201 Created

- **Login**: `POST /signup`
    - Request Body:
      ```json
      {
          "email": "example@example.com",
          "password": "*********"
      }
      ```
    - Response: 200 OK (returns JWT token)

### Events

- **Create Event**: `POST /events`
    - Must be authenticated
    - Request Body:
      ```json
      {
          "name": "Event Name",
          "date": "2024-07-21",
          "location": "Event Location"
      }
      ```
    - Response: 201 Created

- **Get All Events**: `GET /events`
    - Response: 200 OK (returns list of events)

- **Get Event by ID**: `GET /events/:id`
    - Must be authenticated
    - Response: 200 OK (returns event details)

- **Update Event**: `PUT /events/:id`
    - Must be authenticated
    - Must Own the Event
    - Request Body:
      ```json
      {
          "name": "Updated Event Name",
          "date": "2024-08-01",
          "location": "Updated Location"
      }
      ```
    - Response: 200 OK

- **Delete Event**: `DELETE /events/:id`
    - Must be authenticated
    - Must Own the Event
    - Response: 200 OK

### User Registration for Events

- **Register for Event**: `POST /events/:id/register`
    - Must be authenticated
    - Response: 200 OK

- **Cancel Registration for Event**: `DELETE /events/:id/register`
    - Must be authenticated
    - Response: 200 OK

## Contact

If you have any questions or suggestions, feel free to reach out to me at [badara.samb@outlook.com].
