# Gunakan image Go yang sudah diatur dengan baik sebagai builder
FROM golang:1.21 AS builder

# Set variabel lingkungan
ENV GOPATH /go

# Direktori kerja dalam kontainer
WORKDIR /app

ENV BUILD_ENV=development

# Salin kode aplikasi ke dalam kontainer
COPY . .

RUN go mod download

# pindah ke cmd/server untuk build app
WORKDIR /app/cmd/server

# Kompilasi aplikasi
RUN CGO_ENABLED=auto GOOS=linux go build -a -o main .

# Buat gambar Docker akhir
FROM alpine:latest

# Install dependensi yang diperlukan
RUN apk --no-cache add ca-certificates

# Salin biner yang telah dikompilasi ke gambar akhir
COPY --from=builder /app/cmd/server/main .

# Tetapkan direktori kerja
# WORKDIR /app

# Port yang akan digunakan oleh aplikasi (sesuaikan sesuai kebutuhan Anda)
# EXPOSE 8080

# Perintah yang akan dijalankan saat kontainer dimulai
CMD ["./main"]
