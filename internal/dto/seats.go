package dto

type SeatAvailabilityItem struct {
	ID         int    `json:"id"`
	Row        string `json:"row"`
	SeatNumber string `json:"seat_number"`
	SeatType   string `json:"seat_type"`
	Status     string `json:"status"` // available, booked
}

type SeatAvailabilityResponse struct {
	CinemaID       int                    `json:"cinema_id"`
	Date           string                 `json:"date"`
	Time           string                 `json:"time"`
	TotalSeats     int                    `json:"total_seats"`
	AvailableSeats int                    `json:"available_seats"`
	BookedSeats    int                    `json:"booked_seats"`
	Seats          []SeatAvailabilityItem `json:"seats"`
}
