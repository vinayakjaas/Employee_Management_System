# Employee Management System

This is a comprehensive Employee Management System built with Go, Fiber, MongoDB, Kubernetes, and Helm.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
  - [Installation](#installation)
  - [Configuration](#configuration)
  - [Usage](#usage)
- [Deployment with Kubernetes and Helm](#deployment-with-kubernetes-and-helm)
- [Contributing](#contributing)
- [License](#license)

## Introduction

The Employee Management System is a web-based application designed to manage employee information efficiently. It provides functionalities to add, view, update, and delete employee records, making HR management tasks streamlined and organized.

## Features

- Add new employees with details such as name, salary, and age.
- View a list of all employees with their respective details.
- Update employee information.
- Delete employee records.
- Secure authentication and authorization.

## Prerequisites

Before you begin, ensure you have the following installed:

- Go
- MongoDB
- Kubernetes (for deployment)
- Helm (for Helm deployment)
- Docker (optional, for containerization)

## Getting Started

### Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/yourusername/employee-management-system.git

2. Install dependencies:

    ```bash
   cd employee-management-system
   go mod download

## Configuration
1.MongoDB Setup:
 Ensure MongoDB is running on your local machine or a remote server.
 Update the MongoDB connection string in main.go to point to your MongoDB instance.

2.Environment Variables:
 Create a .env file in the project root and add any necessary environment variables.

## Usage
1.Run the application:
  ```bash
  go run main.go 
  ```

2.Access the application in your web browser at http://localhost:3000.

## Deployment with Kubernetes and Helm
To deploy the Employee Management System to a Kubernetes cluster using Helm:

1.Ensure you have a Kubernetes cluster configured and Helm installed.

2.Modify the Helm chart values in the charts/employee-management directory to suit your environment.

3.Install the Helm chart:
 ```bash
   helm install my-go-app my-go-app/
```
4.Access the application using the provided service URL.

## Screenshots

![Screenshot 1](src/Screenshot%20from%202024-03-26%2022-47-45.png)

![Screenshot 1](src/Screenshot%20from%202024-03-26%2022-49-46.png)

![Screenshot 2](src/Screenshot%20from%202024-03-26%2022-51-21.png)


## Contributing
Contributions are welcome! Feel free to open issues or pull requests to improve the project.



