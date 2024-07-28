# UPI QR Code Generator

A simple, privacy-first web application to generate UPI QR codes without storing any user information. The application is built using Go for the backend and plain HTML, CSS, and JavaScript for the frontend.

## Features

- Generate UPI QR codes by entering a UPI ID.
- Immediate client-side validation for UPI ID format.
- Dark/Light theme toggle with modern UI design.
- Stateless server: does not store user information or generated QR codes.

## Prerequisites

- [Go](https://golang.org/dl/) (version 1.22.5 or later)
- [Git](https://git-scm.com/downloads)

## Installation

1. **Clone the repository:**

   ```sh
   git clone https://github.com/yourusername/go-upi-qr.git
   cd go-upi-qr
   ```

2. **Initialize Go modules:**

    ```sh
    go mod tidy
    ```

3. **Usage:**

    3.1. Run the server:

    ```sh
    go run main.go
    ```

    3.2. Open your browser and navigate to:

    http://localhost:8080/index.html
    
    3.3. Enter your UPI ID and click "Generate QR Code" to see the QR code.

    3.4. Toggle between Dark and Light themes using the sun/moon icon.
