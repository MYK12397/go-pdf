package body

import (
	"fmt"
	"strconv"

	"github.com/MYK12397/go-pdf/model"
	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func GetShowDetails(t model.Ticket) []core.Row {
	rows := []core.Row{
		row.New(30).Add(
			image.NewFromFileCol(4, t.ShowPosterLocation, props.Rect{
				Center:  false,
				Percent: 100,
			}),
			col.New(8).Add(
				text.New(t.ShowName, props.Text{
					Style: fontstyle.Bold,
					Size:  10,
				}),
				text.New(t.Language, props.Text{
					Top:   6,
					Style: fontstyle.Normal,
					Size:  8,
					Color: &props.Color{Red: 95, Green: 95, Blue: 95},
				}),
				text.New(t.ShowTime, props.Text{
					Top:   12,
					Style: fontstyle.Bold,
					Size:  10,
				}),
				text.New(t.ShowVenue, props.Text{
					Top:   18,
					Style: fontstyle.Normal,
					Size:  8,
					Color: &props.Color{Red: 95, Green: 95, Blue: 95},
				}),
			),
		),
		row.New(6),
		row.New(1).Add(
			line.NewCol(12, props.Line{
				Thickness:   0.2,
				Color:       &props.Color{Red: 200, Green: 200, Blue: 200},
				SizePercent: 100,
				Style:       linestyle.Dashed,
			}),
		),
		row.New(3),
		row.New(16).Add(
			col.New(2).Add(
				text.New(strconv.Itoa(t.TicketCount), props.Text{
					Style: fontstyle.Bold,
					Size:  24,
					Align: align.Center,
				}),
				text.New("Tickets", props.Text{
					Top:   12,
					Style: fontstyle.Normal,
					Size:  8,
					Color: &props.Color{Red: 95, Green: 95, Blue: 95},
					Align: align.Center,
				}),
			),
			col.New(2),
			col.New(8).Add(
				text.New(t.Screen, props.Text{
					Size:  8,
					Color: &props.Color{Red: 95, Green: 95, Blue: 95},
				}),
				text.New(t.SeatNumber, props.Text{
					Top:   6,
					Style: fontstyle.Bold,
					Size:  14,
				}),
			),
		),
		row.New(3),
		row.New(1).Add(
			line.NewCol(12, props.Line{
				Thickness:   0.2,
				Color:       &props.Color{Red: 200, Green: 200, Blue: 200},
				SizePercent: 100,
				Style:       linestyle.Dashed,
			}),
		),
		row.New(6),
		row.New(20).Add(
			code.NewQrCol(12,
				fmt.Sprintf("%v\n%v\n%v\n%v", t.ID, t.ShowName, t.ShowTime, t.ShowVenue),
				props.Rect{
					Center:  true,
					Percent: 100,
				},
			),
		),
		row.New(10).Add(
			col.New(12).Add(text.New(fmt.Sprintf("Booking ID: %v", t.ID), props.Text{
				Style: fontstyle.Normal,
				Size:  8,
				Align: align.Center,
				Top:   2,
			})),
		),
		row.New(1).Add(
			line.NewCol(12, props.Line{
				Thickness:   0.2,
				Color:       &props.Color{Red: 200, Green: 200, Blue: 200},
				SizePercent: 100,
				Style:       linestyle.Solid,
			}),
		),
		row.New(3),
		row.New(10).Add(
			code.NewBarCol(12, strconv.Itoa(t.ID),
				props.Barcode{
					Center:  true,
					Percent: 100,
				},
			),
		),
	}

	return rows
}
