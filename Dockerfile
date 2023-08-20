# Gunakan image Golang versi 1.19 sebagai base image
FROM golang:1.19-alpine AS builder

# Set working directory di dalam container
WORKDIR /app

# Menyalin file kode Go dan file go.mod/go.sum ke dalam container
COPY . .

# Build aplikasi
RUN go build -o main .

# Jalankan aplikasi
CMD ["./main"]
