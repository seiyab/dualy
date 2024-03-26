package dualy_test

import (
	"image"
	"testing"

	"github.com/seiyab/dualy"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

func TestDualY(t *testing.T) {
	d := dualy.New()
	l1, err := plotter.NewLine(funcXYer(func(x float64) float64 { return x }))
	if err != nil {
		t.Fatal(err)
	}
	d.LeftPlot.Add(l1)

	l2, err := plotter.NewLine(funcXYer(func(x float64) float64 { return x * x * 1.5 }))
	if err != nil {
		t.Fatal(err)
	}
	d.RightPlot.Add(l2)

	dpi := 96
	img := image.NewRGBA(image.Rect(0, 0, 4*dpi, 3*dpi))
	c := vgimg.NewWith(vgimg.UseImage(img))
	d.Draw(draw.New(c))

	matchSnapshot(t, c.Image(), "snapshots/dualy.png")
}

func funcXYer(fn func(x float64) float64) plotter.XYer {
	var xys plotter.XYs
	for i := float64(0); i <= 1; i += 0.01 {
		xys = append(xys, struct{ X, Y float64 }{X: i, Y: fn(i)})
	}
	return xys
}
