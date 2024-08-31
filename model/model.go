package model

type Company struct {
	Address      string
	LogoLocation string
	Name         string
}

type Ticket struct {
	ID                 int
	ShowName           string
	ShowTime           string
	Language           string
	ShowVenue          string
	SeatNumber         string
	TicketCost         float64
	Screen             string
	TicketCount        int
	ShowPosterLocation string
}
