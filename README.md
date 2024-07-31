# AfyaChain

AfyaChain is an electronic medical record and healthcare management information system designed to securely handle patient records using blockchain technology for data integrity and Go for the backend. The application manages patient data, medical records, and user access while ensuring data security and scalability.

## Table of Contents

- [Features](#features)
- [Technologies](#technologies)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Database Schema](#database-schema)
- [Folder Structure](#folder-structure)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Patient Management**: Create, read, update, and delete patient records.
- **Medical Records**: Track and manage medical records associated with patients.
- **User Access Logs**: Log user actions for auditing and security purposes.
- **Blockchain Integration**: Utilize blockchain technology to ensure patient data's integrity and immutability.
- **Responsive UI**: Serve web pages with dynamic content through HTML templates.
- **API Endpoints**: RESTful endpoints for interaction with the application.

## Technologies

- **Backend**: Go
- **Database**: MySQL
- **Blockchain**: Custom blockchain implementation using Go
- **Frontend**: HTML, CSS, JavaScript
- **Containerization**: Docker
- **ORM**: GORM (for Go)
- **Migrations**: Sequelize (for database schema management)

## Installation

### Prerequisites

- Go (1.18 or later)
- MySQL
- Docker (optional, for containerization)

### Setting Up Locally

1. **Clone the Repository**

   ```bash
   git clone https://github.com/Athooh/AfyaChain.git
   cd AfyaChain
   ```

2. **Install Go Dependencies**

   ```bash
   go mod AfyaChain
   ```

3. **Set Up the Database**

   - Update `database.go` with your database credentials.
   - Run the SQL initialization script to set up the database schema:

     ```bash
     mysql -u [username] -p [database_name] < init.sql
     ```

4. **Build the Application**

   ```bash
   go build -o afyachain main.go
   ```

5. **Run the Application**

   ```bash
   ./afyachain
   ```

## Usage

- **Home Page**: Access the home page at `http://localhost:8081/`
- **Patient Management**:
  - **Create Patient**: Submit patient information via the `/create-patient` endpoint.
  - **Get Patient**: Retrieve patient details via the `/patient?id=[id]` endpoint.
  - **Get All Patients**: List all patients via the `/patients` endpoint.
- **Medical Records**:
  - **Create Record**: Create a new medical record via the `/create-record` endpoint.
  - **Get Record**: Retrieve medical records via the `/record?id=[id]` endpoint.

## API Endpoints

### Patient Management

- **POST /create-patient**
  - **Request Body**: `firstname`, `lastname`, `date` (YYYY-MM-DD), `phone`, `email`, `address`, `gender`
  - **Response**: `Patient Created`

- **GET /patient**
  - **Query Parameters**: `id` (int)
  - **Response**: Patient details in JSON format

- **GET /patients**
  - **Response**: List of all patients in JSON format

### Medical Records

- **POST /create-record**
  - **Request Body**: `patientID`, `recordDate` (YYYY-MM-DD), `condition`, `treatment`, `notes`
  - **Response**: `Medical Record Created`

- **GET /record**
  - **Query Parameters**: `id` (int)
  - **Response**: Medical record details in JSON format

## Database Schema

### Patients Table

- **ID**: Integer, Primary Key, Auto Increment
- **FirstName**: Varchar(100)
- **LastName**: Varchar(100)
- **DOB**: Date
- **Gender**: Varchar(10)
- **Email**: Varchar(100), Unique
- **Phone**: Varchar(15)
- **Address**: Varchar(255)
- **CreatedAt**: Timestamp
- **UpdatedAt**: Timestamp

### MedicalRecords Table

- **ID**: Integer, Primary Key, Auto Increment
- **PatientID**: Integer, Foreign Key
- **RecordDate**: Date
- **Condition**: Text
- **Treatment**: Text
- **Notes**: Text
- **CreatedAt**: Timestamp
- **UpdatedAt**: Timestamp

## Folder Structure

- **`backend`**: Contains the core backend logic, including blockchain and database operations.
- **`database`**: Manages database connections and CRUD operations.
- **`migrations`**: Holds migration scripts for setting up the database schema.
- **`models`**: Defines data models for patients and medical records.
- **`static`**: Stores static files like CSS, JS, and images.
- **`templates`**: HTML templates for rendering web pages.
- **`Dockerfile`**: Configuration file for building the Docker image.
- **`main.go`**: Entry point for the Go application.
- **`Makefile`**: Defines build and management commands.
- **`README.md`**: Documentation for the project.

## Contributing

Contributions are welcome! Please follow these steps to contribute:

1. Fork the repository.
2. Create a feature branch (`git checkout -b feature/YourFeature`).
3. Commit your changes (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature/YourFeature`).
5. Open a pull request.

##Authors

1. [Seth Athooh](https://github.com/Athooh)
2. [Bravian Nyatoro](https://github.com/bravian1)
3. [Raymond Caleb](https://github.com/raymond9734)
4. [Vincent Odhiambo](https://github.com/Vincent-Omondi)
5. [Stella Oiro](https://github.com/Stella-Achar-Oiro)

## License

The project is licensed under the MIT License. Please look at the [LICENSE](LICENSE) file for details.

