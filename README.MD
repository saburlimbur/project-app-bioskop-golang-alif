# Bootcamp Mini Project - Bioskop Cinema Booking API Documentation

API untuk sistem pemesanan tiket bioskop yang dibangun dengan Go, Chi Router, dan PostgreSQL.

## Base URL

```
http://localhost:8080/api/v1
```

## Authentication

API ini menggunakan authentication berbasis token. Setelah login, token akan diberikan dan harus disertakan dalam header untuk endpoint yang memerlukan autentikasi.

```
Authorization: Bearer <your_token>
```

## Endpoints

### Health Check

#### GET /

Mengecek status API.

**Response:**

```json
{
  "status": true,
  "message": "OK",
  "data": {
    "message": "Bioskop Cinema Booking API Running"
  }
}
```

---

### User Management

#### POST /api/v1/users/register

Mendaftarkan user baru (Register).

**Request Body:**

```json
{
  "username": "something",
  "email": "something@mail.com",
  "password": "something",
  "full_name": "something",
  "phone_number": "08913121312"
}
```

**Response:**

```json
{
  "status": true,
  "message": "Register successfully, Please verify your email",
  "data": {
    "id": 1,
    "username": "something",
    "email": "something@mail.com",
    "password": "something",
    "full_name": "something",
    "phone_number": "08913121312",
    "email_verified": false,
    "created_at": "2026-01-16T15:45:26.428015Z"
  }
}
```

#### POST /api/v1/users/login

Login user.

**Request Body:**

```json
{
  "email": "something@mail.com",
  "password": "something"
}
```

**Response:**

```json
{
  "status": true,
  "message": "Login successfully",
  "data": {
    "user": {
      "id": 1,
      "username": "something",
      "email": "something@mail.com",
      "password": "something",
      "full_name": "something",
      "phone_number": "08913121312",
      "email_verified": true,
      "created_at": "2026-01-16T15:45:26.428015Z"
    },
    "token": "some13thing12token",
    "expires_at": "2026-01-17T16:01:13.2499842+07:00"
  }
}
```

#### POST /api/v1/users/logout

Logout user.

**Headers:**

```
Authorization: Bearer <token>
```

**Response:**

```json
{
  "status": true,
  "message": "Logout successfully"
}
```

#### POST /api/v1/users/verify-email

Verifikasi email user.

**Request Body:**

```json
{
  "email": "something@mail.com",
  "otp": "001312"
}
```

**Response:**

```json
{
  "status": true,
  "message": "Email verified successfully",
  "data": {
    "id": 1,
    "username": "something",
    "email": "something@mail.com",
    "password": "something",
    "full_name": "something",
    "phone_number": "08913121312",
    "email_verified": true,
    "created_at": "2026-01-16T15:45:26.428015Z"
  }
}
```

---

### Cinema

#### GET /api/v1/cinemas/?page=1

Mendapatkan daftar cinema.

**Headers:**

```
Authorization: Bearer <token>
```

**Query Parameters:**

- `page` (optional): Nomor halaman (default: 1)
- `limit` (optional): Jumlah data per halaman (default: 10)
- `search` (optional): Kata kunci pencarian

**Response:**

```json
{
  "status": true,
  "message": "success get data",
  "data": [
    {
      "id": 1,
      "name": "Cineplex Galaxy",
      "address": "Jl. Sudirman No. 10",
      "city": "Jakarta",
      "phone_number": "021-5551234",
      "total_seats": 150,
      "created_at": "2026-01-12T16:17:29.931439Z",
      "updated_at": "2026-01-12T16:17:29.931439Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 1,
    "total_rows": 6,
    "total_pages": 2
  }
}
```

#### GET /api/v1/cinemas/{id}

Mendapatkan detail cinema berdasarkan ID.

**Headers:**

```
Authorization: Bearer <token>
```

**Response:**

```json
{
  "status": true,
  "message": "success",
  "data": {
    "id": 1,
    "name": "Mall Taman Anggrek XXI",
    "address": "Jl. Letjen S. Parman No. 21",
    "city": "Jakarta",
    "phone_number": "021-5559012",
    "total_seats": 200,
    "created_at": "2026-01-12T16:17:29.931439Z"
  }
}
```

#### GET /api/v1/cinemas/{id}/seats

Mendapatkan ketersediaan kursi pada cinema tertentu.

### Contoh Endpoint

#### GET /cinemas/1/seats?date=2026-01-20&time=19:00

**Headers:**

```
Authorization: Bearer <token>
```

**Query Parameters:**

- `showtime_id` (required): ID dari jadwal tayang

**Response:**

```json
{
  "status": true,
  "message": "success get seat availability",
  "data": {
    "cinema_id": 1,
    "date": "2026-01-20",
    "time": "19:00",
    "total_seats": 3,
    "available_seats": 2,
    "booked_seats": 0,
    "seats": [
      {
        "id": 1,
        "row": "A",
        "seat_number": "A1",
        "seat_type": "regular",
        "status": "available"
      },
      {
        "id": 2,
        "row": "A",
        "seat_number": "A2",
        "seat_type": "regular",
        "status": "available"
      }
    ]
  }
}
```

---

### Booking

#### POST /api/v1/booking

Membuat booking baru.

**Headers:**

```
Authorization: Bearer <token>
```

**Request Body:**

```json
{
  "showtime_id": 1,
  "seat_id": 3
}
```

**Response:**

```json
{
  "status": true,
  "message": "Booking created successfully",
  "data": {
    "id": 1,
    "user_id": 1,
    "showtime_id": 1,
    "seat_id": 3,
    "booking_code": "BK-363330",
    "booking_status": "pending",
    "total_price": 50000,
    "expired_at": "2026-01-16T16:30:22.024545Z",
    "created_at": "2026-01-16T16:15:22.023738Z",
    "updated_at": "2026-01-16T16:15:22.023738Z"
  }
}
```

#### GET /api/v1/booking/{id}

Mendapatkan detail booking berdasarkan ID.

**Headers:**

```
Authorization: Bearer <token>
```

**Response:**

```json
{
  "status": true,
  "message": "successfully get booking detail",
  "data": {
    "id": 4,
    "user_id": 8,
    "showtime_id": 1,
    "seat_id": 3,
    "booking_code": "BK-363330",
    "booking_status": "pending",
    "total_price": 50000,
    "expired_at": "2026-01-16T16:30:22.024545Z",
    "created_at": "2026-01-16T16:15:22.023738Z"
  }
}
```

<!-- #### GET /api/v1/user/bookings-lists

Mendapatkan history booking user.

**Headers:**

```
Authorization: Bearer <token>
```

**Query Parameters:**

- `page` (optional): Nomor halaman (default: 1)
- `limit` (optional): Jumlah data per halaman (default: 10)
- `status` (optional): Filter berdasarkan status (pending, confirmed, cancelled)

**Response:**

```json
{
  "status": "success",
  "message": "Booking history retrieved successfully",
  "data": {
    "bookings": [
      {
        "booking_id": "uuid",
        "movie_title": "Movie Title",
        "cinema_name": "Cinema XXI",
        "showtime": "2025-01-17T14:00:00Z",
        "seats": ["A1", "A2"],
        "total_price": 100000,
        "status": "confirmed",
        "created_at": "2025-01-16T10:00:00Z"
      }
    ],
    "pagination": {
      "page": 1,
      "limit": 10,
      "total": 25
    }
  }
}
``` -->

#### GET /api/v1/user/bookings

Mendapatkan history booking user.

**Headers:**

```
Authorization: Bearer <token>
```

**Response:**

```json
{
  "status": true,
  "message": "successfully get booking history",
  "data": [
    {
      "id": 4,
      "booking_code": "BK-363330",
      "booking_status": "pending",
      "total_price": 50000,
      "expired_at": "2026-01-16T16:30:22.024545Z",
      "created_at": "2026-01-16T16:15:22.023738Z",
      "movie_title": "Avengers: Endgame",
      "cinema_name": "Cineplex Galaxy",
      "show_date": "2026-01-20",
      "show_time": "19:00",
      "seat_number": "A3",
      "seat_row": "A",
      "seat_type": "regular",
      "payment_status": "pending",
      "payment_amount": 0
    }
  ]
}
```

---

### Payment

#### GET /api/v1/payment-methods

Mendapatkan daftar metode pembayaran yang tersedia.

**Headers:**

```
Authorization: Bearer <token>
```

**Response:**

```json
{
  "status": true,
  "message": "successfully get payment methods",
  "data": [
    {
      "id": 3,
      "name": "GoPay",
      "code": "GOPAY"
    },
    {
      "id": 2,
      "name": "OVO",
      "code": "OVO"
    },
    {
      "id": 1,
      "name": "Transfer Bank",
      "code": "BANK_TRANSFER"
    }
  ]
}
```

#### POST /api/v1/payment

Membuat pembayaran untuk booking.

**Headers:**

```
Authorization: Bearer <token>
```

**Request Body:**

```json
{
  "booking_id": 4,
  "payment_method_id": 1
}
```

**Response:**

```json
{
  "status": true,
  "message": "Payment created succesfully",
  "data": {
    "id": 8,
    "booking_id": 4,
    "amount": 50000,
    "payment_status": "success",
    "payment_date": "2026-01-16T16:22:49.686959Z",
    "transaction_id": "TRX-20260116162249-735135",
    "created_at": "2026-01-16T16:22:49.680189Z",
    "updated_at": "0001-01-01T00:00:00Z",
    "payment_method": {
      "id": 1,
      "name": "Transfer Bank",
      "code": "BANK_TRANSFER",
      "description": ""
    }
  }
}
```

---

## Error Responses

API akan mengembalikan error dalam format berikut:

```json
{
  "status": "error",
  "message": "Error message here",
  "errors": {
    "field_name": ["validation error message"]
  }
}
```

### HTTP Status Codes

- `200 OK`: Request berhasil
- `201 Created`: Resource berhasil dibuat
- `400 Bad Request`: Request tidak valid
- `401 Unauthorized`: Authentication diperlukan atau gagal
- `403 Forbidden`: Akses ditolak
- `404 Not Found`: Resource tidak ditemukan
- `422 Unprocessable Entity`: Validation error
- `500 Internal Server Error`: Server error

---

## Notes

1. Semua endpoint kecuali health check, register, login, verify-email, dan logout memerlukan authentication token
2. Token diberikan setelah login berhasil dan harus disertakan di header `Authorization: Bearer <token>`
3. Booking memiliki waktu kadaluarsa (biasanya 15 menit), pastikan untuk melakukan pembayaran sebelum booking expire
4. Semua timestamp menggunakan format ISO 8601 (UTC)
5. Harga dalam satuan Rupiah (IDR)
