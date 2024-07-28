package main

import (
	"log"
	"net/http"
	"regexp"

	"github.com/skip2/go-qrcode"
)

func isValidUPI(upiID string) bool {
	// Basic UPI ID format check: username@bankname
	// Username: alphanumeric characters, can include '.', '_', '-'
	// Bankname: alphanumeric characters
	var validUPI = regexp.MustCompile(`^[a-zA-Z0-9.\-_]+@[a-zA-Z0-9]+$`)
	return validUPI.MatchString(upiID)
}

func generateQR(w http.ResponseWriter, r *http.Request) {
	upiID := r.URL.Query().Get("upi_id")
	if upiID == "" {
		http.Error(w, "UPI ID is required", http.StatusBadRequest)
		return
	}

	if !isValidUPI(upiID) {
		http.Error(w, "Invalid UPI ID format", http.StatusBadRequest)
		return
	}

	// Generate UPI URI
	upiURI := "upi://pay?pa=" + upiID + "&pn=Payee"

	// Generate QR Code
	png, err := qrcode.Encode(upiURI, qrcode.Medium, 256)
	if err != nil {
		http.Error(w, "Failed to generate QR code", http.StatusInternalServerError)
		return
	}

	// Return the QR code as a PNG image
	w.Header().Set("Content-Type", "image/png")
	w.Write(png)
}

func main() {
	// Serve static files from the current directory
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	http.HandleFunc("/generate_qr", generateQR)

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
