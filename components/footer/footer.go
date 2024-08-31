package footer

import (
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func GetPageFooter() core.Row {
	return row.New(2).Add(
		col.New(12).Add(
			text.New("Powered by BookmyShow", props.Text{
				Style: fontstyle.Italic,
				Size:  8,
				Align: align.Center,
				Color: &props.Color{Red: 255, Green: 120, Blue: 218},
			}),
		),
	)
}
