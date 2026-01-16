-- Table Users (Customer)
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    full_name VARCHAR(100) NOT NULL,
    phone_number VARCHAR(20),
    is_verified BOOLEAN DEFAULT FALSE
    email_verified_at TIMESTAMP;
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table Sessions (untuk menyimpan token login)
CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token VARCHAR(255) UNIQUE NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    ip_address VARCHAR(45),
    device_info TEXT,
    revoked_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table Cinemas (Bioskop)
CREATE TABLE cinemas (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    address TEXT NOT NULL,
    city VARCHAR(50) NOT NULL,
    phone_number VARCHAR(20),
    total_seats INTEGER NOT NULL DEFAULT 50,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table Movies (Film yang diputar)
CREATE TABLE movies (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    genre VARCHAR(100),
    duration INTEGER NOT NULL, -- dalam menit
    rating VARCHAR(10),
    description TEXT,
    thumbnail TEXT NOT NULL,
    poster TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table Showtimes (Jadwal tayang film di bioskop)
CREATE TABLE showtimes (
    id SERIAL PRIMARY KEY,
    cinema_id INTEGER NOT NULL REFERENCES cinemas(id) ON DELETE CASCADE,
    movie_id INTEGER NOT NULL REFERENCES movies(id) ON DELETE CASCADE,
    show_date DATE NOT NULL,
    show_time TIME NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(cinema_id, movie_id, show_date, show_time)
);

-- Table Seats (Kursi di setiap bioskop)
CREATE TABLE seats (
    id SERIAL PRIMARY KEY,
    cinema_id INTEGER NOT NULL REFERENCES cinemas(id) ON DELETE CASCADE,
    seat_number VARCHAR(10) NOT NULL,
    row_number VARCHAR(5) NOT NULL,
    seat_type VARCHAR(20) DEFAULT 'regular', -- regular, vip, premium
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(cinema_id, seat_number)
);

-- Table Payment Methods (Metode Pembayaran)
CREATE TABLE payment_methods (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    code VARCHAR(20) UNIQUE NOT NULL,
    description TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table Bookings (Pemesanan Tiket)
CREATE TABLE bookings (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    showtime_id INTEGER NOT NULL REFERENCES showtimes(id) ON DELETE CASCADE,
    seat_id INTEGER NOT NULL REFERENCES seats(id) ON DELETE CASCADE,
    booking_code VARCHAR(20) UNIQUE NOT NULL,
    booking_status VARCHAR(20) DEFAULT 'pending', -- 'confirmed', 'cancelled',
    total_price DECIMAL(10, 2) NOT NULL,
    expired_at TIMESTAMP NOT NULL, -- kadaluarsa booking
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(showtime_id, seat_id) -- Satu kursi tidak bisa dipesan 2x untuk showtime yang sama
);

-- Table Payments (Pembayaran)
CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    booking_id INTEGER NOT NULL REFERENCES bookings(id) ON DELETE CASCADE,
    payment_method_id INTEGER NOT NULL REFERENCES payment_methods(id),
    amount DECIMAL(10, 2) NOT NULL,
    payment_status VARCHAR(20) DEFAULT 'pending', -- success, failed
    payment_date TIMESTAMP,
    transaction_id VARCHAR(100) UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE otp_verifications (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    otp_code VARCHAR(6) NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    is_used BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_otp_user 
ON otp_verifications(user_id);


CREATE INDEX idx_sessions_token 
ON sessions(token);

CREATE INDEX idx_sessions_user_id 
ON sessions(user_id);

CREATE INDEX idx_seats_cinema 
ON seats(cinema_id);

-- ==

CREATE INDEX idx_showtimes_cinema_date 
ON showtimes (cinema_id, show_date);

CREATE INDEX idx_bookings_user 
ON bookings (user_id);

CREATE INDEX idx_bookings_showtime 
ON bookings (showtime_id);

CREATE INDEX idx_payments_booking 
ON payments (booking_id);

-- ===