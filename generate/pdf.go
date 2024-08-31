package generate

import (
	"log"

	"github.com/MYK12397/go-pdf/components/body"
	"github.com/MYK12397/go-pdf/components/footer"
	"github.com/MYK12397/go-pdf/components/header"
	"github.com/MYK12397/go-pdf/model"
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func GetMaroto(c model.Company, t model.Ticket) core.Maroto {
	cfg := config.NewBuilder().WithDimensions(120, 200).Build()

	mrt := maroto.New(cfg)

	err := mrt.RegisterHeader(header.GetPageHeader(c))
	if err != nil {
		log.Println("error registering header")
	}

	mrt.AddRow(6)
	mrt.AddRow(4, line.NewCol(12, props.Line{
		Thickness:   0.2,
		Color:       &props.Color{Red: 200, Green: 200, Blue: 200},
		SizePercent: 100,
	}))

	mrt.AddRow(6)
	mrt.AddRows(body.GetShowDetails(t)...)

	mrt.AddRow(8)
	err = mrt.RegisterFooter(footer.GetPageFooter())
	if err != nil {
		log.Println("error registering footer")
	}

	return mrt
}
