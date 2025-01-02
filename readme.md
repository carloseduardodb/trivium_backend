# Crypto Finance

A Go-based cryptocurrency finance tracking and analysis application.

## Description

This project is a robust cryptocurrency finance application built with Go, designed to track and analyze cryptocurrency market data, manage portfolios, and provide insights for crypto investments.

## Features

- Real-time cryptocurrency price tracking
- Portfolio management
- Market analysis tools
- Transaction history
- Price alerts
- Performance analytics

## Prerequisites

- Go 1.21 or higher
- Git
- Internet connection for API calls

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/crypto-finance.git
cd crypto-finance
```

2. Install dependencies:
```bash
go mod tidy
```

3. Build the application:
```bash
go build
```

## Build and Run with Docker

To build and run the application using Docker, follow the steps below:

### Prerequisites

- Make sure you have Docker installed on your machine.

### Build the Image

Run the following command in the root of the project:

```bash
docker build -t crypto-finance .
```

### Run the Container

After building the image, you can run the container with the following command:

```bash
docker run -p 8080:8080 -p 3000:3000 crypto-finance
```

This will expose the backend on port 8080 and the frontend on port 3000.

## Usage

Run the application:
```bash
./crypto-finance
```

## Configuration

Create a `.env` file in the root directory with your configuration:

```env
API_KEY=your_api_key
PORT=8080
```

## Contributing

1. The project is under construction.

## License

This project is licensed under the [Creative Commons Non-Commercial License (CC BY-NC)](LICENSE).
