package header

import (
	"github.com/MYK12397/go-pdf/model"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func GetPageHeader(c model.Company) core.Row {
	return row.New(16).Add(
		image.NewFromFileCol(4, c.LogoLocation, props.Rect{
			Center:  true,
			Percent: 100,
		}),
		col.New(2),
		col.New(6).Add(
			text.New(c.Name, props.Text{
				Style: fontstyle.Bold,
				Size:  10,
			}),
			text.New(c.Address, props.Text{
				Top:  6,
				Size: 10,
			}),
		),
	)
}
