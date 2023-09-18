# stlReader-DevStorm - Beta

## Overview

This is a sample README for an application that uses Docker Compose to run the "challenge-client" (React.js) and "challenge-server" (Golang API). This document provides an overview of the project, setup instructions, and basic usage information.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Docker: [Install Docker](https://docs.docker.com/get-docker/)
- Docker Compose: [Install Docker Compose](https://docs.docker.com/compose/install/)

## Getting Started

To get the project up and running, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/BrenoSantanaBruno/stlReader-DevStorm
Navigate to the project directory:

```bash
Copy code
cd stlReader-DevStorm
Build and start the Docker containers:
```

```bash
docker-compose up --build
```
This command will build and start the "challenge-client" and "challenge-server" containers defined in the docker-compose.yml file.

Usage
Once the Docker containers are running, you can access the application as follows:

React Client (Frontend): http://localhost:3000
Golang API (Backend): http://localhost:8080
Feel free to access the client and make API requests as needed.

Configuration
You can customize the configuration of both the React client and Golang server in their respective project directories. Refer to the project-specific documentation for more information on configuring each component.

Troubleshooting
If you encounter issues or have questions, refer to the project's documentation or check the GitHub repository for any known issues or updates.

Contributing
Contributions are welcome! If you'd like to contribute to the project, please follow the guidelines in the project's CONTRIBUTING.md file.