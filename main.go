package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MYK12397/go-pdf/generate"
	"github.com/MYK12397/go-pdf/model"
)

func main() {
	c := model.Company{
		Name:         "BookmyShow",
		Address:      "Lake City Mall, Noida",
		LogoLocation: "./bookmyshow-logo.png",
	}

	t := model.Ticket{
		ID:                 1,
		ShowName:           "DeadPool 3 (U/A)",
		ShowTime:           "Sat, 07 Sept | 7:00 PM",
		Language:           "English",
		ShowVenue:          "Inox: Lake City Mall",
		SeatNumber:         "Platinum - A1, A2",
		TicketCost:         620.00,
		Screen:             "Screen: RECLINER Q5",
		TicketCount:        2,
		ShowPosterLocation: "./download.jpeg",
	}

	m := generate.GetMaroto(c, t)

	doc, err := m.Generate()
	if err != nil {
		log.Println("error while generating PDF")
	}

	filename := fmt.Sprintf("ticket-%d.pdf", t.ID)

	if _, err := os.Stat("temp"); os.IsNotExist(err) {
		err := os.Mkdir("temp", 0755)
		if err != nil {
			log.Println("error creating directory: ", err)
		}
	}

	err = doc.Save("temp/" + filename)
	if err != nil {
		log.Println("unable to save file: ", err)
	}
}
