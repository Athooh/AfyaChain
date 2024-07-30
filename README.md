# AfyaChain Blockchain-Based Healthcare System

## Table of Contents

1. [Project Overview](#project-overview)
2. [Features](#features)
3. [Architecture](#architecture)
4. [Setup and Installation](#setup-and-installation)
    - [Prerequisites](#prerequisites)
    - [Backend Setup (Go)](#backend-setup-go)
    - [Smart Contracts Setup (Solidity)](#smart-contracts-setup-solidity)
5. [Usage](#usage)
6. [Directory Structure](#directory-structure)
7. [Contributing](#contributing)
8. [License](#license)

---

## Project Overview

AfyaChain's Blockchain-Based Healthcare System is designed to enhance the management and exchange of healthcare data using blockchain technology. This system aims to provide secure, immutable, and decentralized storage and access to patient health records, integrating various healthcare providers, patients, and insurance companies into a unified network.

## Features

- **Electronic Health Records Management**: Securely create, edit, and store patient health records.
- **Personal Health Records Access**: Patients can access their health data anytime.
- **Medical Claims Management**: Streamlined sharing of health data for insurance claims.
- **Clinical Research and Trials**: Access to anonymized patient data for research.
- **Automated Doctor Referrals Processing**: Efficient handling of doctor referrals.
- **User Authorization and Data Access Verification**: Multi-factor authentication and private ID-based access.
- **Data Security and HIPAA Compliance**: Ensuring data encryption and secure access.

## Architecture

The system is composed of the following key components:
- **Blockchain Network**: A distributed ledger storing all patient health records.
- **Smart Contracts**: Written in Solidity to handle record creation, validation, and access authorization.
- **Backend**: Built with Go, providing APIs for frontend applications.
- **Frontend**: Interfaces for patients and healthcare providers to interact with the system.

## Setup and Installation

### Prerequisites

- Go 1.16 or higher
- Node.js and npm
- Solidity Compiler (solc)

### Backend Setup (Go)

1. **Clone the Repository**:
    ```bash
    git clone https://github.com/yourusername/afya-chain.git
    cd afya-chain
    ```

2. **Install Dependencies**:
    Ensure you have Go installed. Navigate to the `backend` directory and run:
    ```bash
    cd backend
    go mod tidy
    ```

3. **Environment Variables**:
    Create a `.env` file in the `backend` directory with the following content:
    ```env
    PORT=8080
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=yourusername
    DB_PASSWORD=yourpassword
    DB_NAME=yourdbname
    ```

4. **Run the Backend**:
    ```bash
    go run main.go
    ```

### Smart Contracts Setup (Solidity)

1. **Install solc**:
    Ensure you have Node.js installed, then install the Solidity compiler globally:
    ```bash
    npm install -g solc
    ```

2. **Write and Compile Smart Contracts**:
    Create a Solidity file named `HelloWorld.sol`:
    ```solidity
    // SPDX-License-Identifier: MIT
    pragma solidity ^0.8.0;

    contract HelloWorld {
        string public greet = "Hello, World!";
    }
    ```

    Compile the contract:
    ```bash
    solcjs --bin --abi HelloWorld.sol
    ```

3. **Generate Go Bindings for Smart Contracts**:
    Install the `abigen` tool from go-ethereum:
    ```bash
    go get -u github.com/ethereum/go-ethereum/cmd/abigen
    ```

    Generate the Go bindings:
    ```bash
    abigen --sol HelloWorld.sol --pkg main --out HelloWorld.go
    ```

## Usage

1. **Start the Backend**:
    Ensure the backend is running:
    ```bash
    go run main.go
    ```

2. **Interact with the Smart Contract**:
    Use the Go bindings to deploy and interact with the smart contract. Example code can be found in the `examples` directory.

3. **Access the Frontend**:
    Navigate to the frontend directory, install dependencies, and start the frontend application.

## Directory Structure

```
afya-chain/
├── backend/
│   ├── main.go
│   ├── handlers/
│   ├── models/
│   ├── routes/
│   └── .env
├── contracts/
│   ├── HelloWorld.sol
│   └── HelloWorld.go
├── frontend/
│   ├── public/
│   ├── src/
│   └── package.json
└── README.md
```

## Contributing

We welcome contributions! Please read our [contributing guidelines](CONTRIBUTING.md) for details on the code of conduct and the process for submitting pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

********************************

Here is the link to chatgpt response:
https://chatgpt.com/share/25532be1-0310-4042-a301-9087970b77b9


Backend (Go)
cmd/server/
main.go: The entry point of the application, initializing the server, setting up routes, and starting the server.
internal/
api/:

handlers/: Contains handlers for various endpoints.
health_records.go: Handles requests related to health records creation, viewing, and updating.
patient_access.go: Manages patient data access requests.
claims_management.go: Manages medical claims processing.
research_trials.go: Handles requests related to clinical research and trials.
router.go: Defines the routes for the API endpoints and associates them with the appropriate handlers.
config/:

config.go: Configuration settings for the application, such as database credentials, blockchain network settings, etc.
database/:

database.go: Initializes the database connection and provides functions for interacting with the database.
middleware/:

auth.go: Middleware for handling authentication, ensuring that requests are properly authenticated.
logger.go: Middleware for logging HTTP requests and responses.
models/:

health_record.go: Defines the data model for health records.
patient.go: Defines the data model for patient information.
claim.go: Defines the data model for medical claims.
services/:

health_records_service.go: Business logic for managing health records.
patient_service.go: Business logic for managing patient data.
claims_service.go: Business logic for managing medical claims.
utils/:

encryption.go: Utility functions for data encryption.
hash.go: Utility functions for generating and verifying hash values.
pkg/
blockchain/:

client.go: Initializes the blockchain client and connects to the blockchain network.
contracts.go: Contains the Go bindings for the Solidity smart contracts.
transactions.go: Functions for creating and submitting transactions to the blockchain.
auth/:

jwt.go: Functions for generating and verifying JSON Web Tokens (JWT) for authentication.
go.mod and go.sum
go.mod: The Go module file, listing the module's dependencies.
go.sum: Checksums for verifying the integrity of the module's dependencies.
Contracts (Solidity)
HealthRecords.sol: Solidity smart contract for managing health records on the blockchain.
ClaimsManagement.sol: Solidity smart contract for managing medical claims on the blockchain.
Migrations.sol: Solidity contract for handling contract migrations.
Migrations
1_initial_migration.js: JavaScript file for deploying the initial set of contracts to the blockchain.
Frontend
doctor-portal/
public/: Contains static files such as HTML, CSS, and images for the doctor portal.
src/: Contains the source code for the doctor portal.
components/: React components for various parts of the doctor portal.
Login.js: Component for the login page.
Dashboard.js: Component for the main dashboard.
HealthRecord.js: Component for viewing and managing health records.
Claims.js: Component for managing medical claims.
Research.js: Component for accessing clinical research data.
App.js: Main application component for the doctor portal.
index.js: Entry point for the React application.
package.json: Lists dependencies and scripts for building and running the doctor portal.
webpack.config.js: Configuration file for Webpack, specifying how to bundle the frontend assets.
patient-portal/
public/: Contains static files such as HTML, CSS, and images for the patient portal.
src/: Contains the source code for the patient portal.
components/: React components for various parts of the patient portal.
Login.js: Component for the login page.
Dashboard.js: Component for the main dashboard.
HealthData.js: Component for viewing personal health data.
Insurance.js: Component for managing insurance information.
Appointments.js: Component for managing appointments.
App.js: Main application component for the patient portal.
index.js: Entry point for the React application.
package.json: Lists dependencies and scripts for building and running the patient portal.
webpack.config.js: Configuration file for Webpack, specifying how to bundle the frontend assets.
Build
contracts/: Contains the compiled JSON files for the Solidity smart contracts.
HealthRecords.json: Compiled output for the HealthRecords contract.
ClaimsManagement.json: Compiled output for the ClaimsManagement contract.
Migrations.json: Compiled output for the Migrations contract.
Scripts
deploy.js: Script for deploying the smart contracts to the blockchain.
interact.js: Script for interacting with the deployed smart contracts.
