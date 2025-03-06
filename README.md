# Tennis Club Website

A modern web application for managing a local tennis club. Built with **Go (Golang)** for the backend, **PostgreSQL** as the database, and **React.js** for the frontend.

## Table of Contents

1. [Features](#features)
2. [Tech Stack](#tech-stack)
3. [Project Structure](#project-structure)
4. [Setup Instructions](#setup-instructions)
   - [Backend](#backend-setup)
   - [Frontend](#frontend-setup)
5. [Running the Application](#running-the-application)
6. [API Documentation](#api-documentation)
7. [Deployment](#deployment)
8. [Contributing](#contributing)
9. [License](#license)

---

## Features

- **User Authentication**: Secure login and registration for players, coaches, and admins.
- **Court Booking System**: Reserve tennis courts with a user-friendly calendar interface.
- **Admin Dashboard**: Manage users, bookings, and announcements.
- **Responsive Design**: Optimized for mobile, tablet, and desktop devices.
- **Announcements**: Display news and updates for club members.

---

## Tech Stack

- **Backend**:
  - Language: Go (Golang)
  - Framework: Gin or Fiber
  - Database: PostgreSQL
  - Authentication: JWT (JSON Web Tokens)
- **Frontend**:
  - Framework: React.js
  - State Management: React Context or Redux
  - Styling: TailwindCSS or Bootstrap
  - Routing: React Router
- **Other Tools**:
  - API Testing: Postman or Insomnia
  - CI/CD: GitHub Actions
  - Deployment: Docker, AWS, Heroku, or Vercel

---

## Project Structure

### Backend

```
tennis-club-backend/
├── cmd/                 # Application entry point
├── internal/            # Core application logic
│   ├── config/          # Configuration setup
│   ├── handlers/        # HTTP route handlers
│   ├── models/          # Database models
│   ├── repositories/    # Database interaction logic
│   ├── services/        # Business logic
│   └── middleware/      # Middleware for auth, logging, etc.
├── migrations/          # Database migration scripts
└── pkg/                 # Shared utilities and helpers
```

### Frontend

```
tennis-club-frontend/
├── public/              # Static assets
├── src/                 # Source code
│   ├── components/      # Reusable UI components
│   ├── pages/           # Page-level components
│   ├── context/         # Global state management
│   ├── hooks/           # Custom React hooks
│   ├── services/        # API service layer
│   └── utils/           # Utility functions
└── package.json         # Node.js dependencies
```

---

## Setup Instructions

### Backend Setup

1. **Install Dependencies**:
   ```bash
   go mod download
   ```

2. **Set Up PostgreSQL**:
   - Install PostgreSQL and create a database named `tennis_club`.
   - Update the `.env` file with your database credentials:
     ```
     DB_HOST=localhost
     DB_PORT=5432
     DB_USER=your_username
     DB_PASSWORD=your_password
     DB_NAME=tennis_club
     ```

3. **Run Migrations**:
   Use `golang-migrate` to apply database migrations:
   ```bash
   migrate -path ./migrations -database "postgres://username:password@localhost:5432/tennis_club?sslmode=disable" up
   ```

4. **Start the Server**:
   ```bash
   go run cmd/main.go
   ```

### Frontend Setup

1. **Install Dependencies**:
   ```bash
   npm install
   ```

2. **Set Environment Variables**:
   Create a `.env` file in the `tennis-club-frontend` directory:
   ```
   REACT_APP_API_URL=http://localhost:8080
   ```

3. **Start the Development Server**:
   ```bash
   npm start
   ```

---

## Running the Application

1. Start the backend server:
   ```bash
   go run cmd/main.go
   ```

2. Start the frontend development server:
   ```bash
   npm start
   ```

3. Access the application at `http://localhost:3000`.

---

## API Documentation

The backend exposes RESTful APIs for various functionalities:

- **Authentication**:
  - `POST /api/auth/register`: Register a new user.
  - `POST /api/auth/login`: Log in and receive a JWT token.
- **Courts**:
  - `GET /api/courts`: Fetch all available courts.
  - `POST /api/bookings`: Create a new booking.
- **Admin**:
  - `GET /api/admin/users`: Fetch all users (admin-only).
  - `DELETE /api/admin/bookings/:id`: Delete a booking (admin-only).

For detailed API documentation, use tools like Swagger or Postman.

---

## Deployment

1. **Backend**:
   - Build the Go binary:
     ```bash
     go build -o tennis-club-backend cmd/main.go
     ```
   - Deploy using Docker or directly on a server.

2. **Frontend**:
   - Build the React app:
     ```bash
     npm run build
     ```
   - Deploy the `build` folder to a static hosting service like Netlify, Vercel, or AWS S3.

3. **Database**:
   - Host PostgreSQL on a managed service (e.g., AWS RDS) or self-host it.

4. **CI/CD**:
   - Set up a pipeline using GitHub Actions or GitLab CI for automated testing and deployment.

---

## Contributing

We welcome contributions! To contribute:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature`).
3. Commit your changes (`git commit -m "Add your feature"`).
4. Push to the branch (`git push origin feature/your-feature`).
5. Open a pull request.

Please ensure your code adheres to the project's coding standards and includes appropriate tests.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---
