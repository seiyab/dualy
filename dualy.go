package dualy

import (
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

type DualY struct {
	LeftPlot  SidePlot
	RightPlot SidePlot
	X         plot.Axis
	LeftY     plot.Axis
	RightY    plot.Axis
}

func New() *DualY {
	p := plot.New()
	return &DualY{
		LeftPlot:  SidePlot{},
		RightPlot: SidePlot{},
		X:         p.X,
		LeftY:     p.Y,
		RightY:    p.Y,
	}
}

func (d *DualY) Draw(c draw.Canvas) {
	c.SetColor(color.White)
	c.Fill(c.Rectangle.Path())

	size := c.Size()
	width := size.X
	height := size.Y
	dataC := draw.Canvas{
		Canvas: c.Canvas,
		Rectangle: vg.Rectangle{
			Min: vg.Point{
				X: c.Min.X + width*0.1,
				Y: c.Min.Y + height*0.1,
			},
			Max: vg.Point{
				X: c.Max.X - width*0.1,
				Y: c.Max.Y - height*0.1,
			},
		},
	}

	for _, pr := range d.LeftPlot.plotters {
		dr, ok := pr.(plot.DataRanger)
		if !ok {
			continue
		}
		xmin, xmax, ymin, ymax := dr.DataRange()
		d.X.Min = math.Min(d.X.Min, xmin)
		d.X.Max = math.Max(d.X.Max, xmax)
		d.LeftY.Min = math.Min(d.LeftY.Min, ymin)
		d.LeftY.Max = math.Max(d.LeftY.Max, ymax)
	}
	for _, pr := range d.RightPlot.plotters {
		dr, ok := pr.(plot.DataRanger)
		if !ok {
			continue
		}
		xmin, xmax, ymin, ymax := dr.DataRange()
		d.X.Min = math.Min(d.X.Min, xmin)
		d.X.Max = math.Max(d.X.Max, xmax)
		d.RightY.Min = math.Min(d.RightY.Min, ymin)
		d.RightY.Max = math.Max(d.RightY.Max, ymax)
	}
	for _, pr := range d.LeftPlot.plotters {
		p := plot.New()
		p.X = d.X
		p.Y = d.LeftY
		pr.Plot(dataC, p)
	}
	for _, pr := range d.RightPlot.plotters {
		p := plot.New()
		p.X = d.X
		p.Y = d.RightY
		pr.Plot(dataC, p)
	}
}
