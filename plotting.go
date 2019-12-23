package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
)

var (
	p, _ = plot.New()
	p5, _ = plot.New()
)

func plotting(xPoints []float64, yPoints []float64, k int, name string, filename string,
	r uint8, g uint8, b uint8) {
	p.Title.Text = "functions"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"



	l, err := plotter.NewLine(funcPoints(xPoints, yPoints, k))
	if err != nil {
		panic(err)
	}
	l.LineStyle.Color = color.RGBA{r, g, b, 255}

	p.Add(l)
	p.Legend.Add(name, l)
	// Save the plot to a PNG file.
	if err := p.Save(10*vg.Inch, 8*vg.Inch, filename); err != nil {
		panic(err)
	}
}

func plotting5(xPoints []float64, yPoints []float64, k int, name string, filename string,
	r uint8, g uint8, b uint8) {
	p5.Title.Text = "functions"
	p5.X.Label.Text = "X"
	p5.Y.Label.Text = "Y"



	l, err := plotter.NewLine(funcPoints(xPoints, yPoints, k))
	if err != nil {
		panic(err)
	}
	l.LineStyle.Color = color.RGBA{r, g, b, 255}

	p5.Add(l)
	p5.Legend.Add(name, l)
	// Save the plot to a PNG file.
	if err := p5.Save(10*vg.Inch, 8*vg.Inch, filename); err != nil {
		panic(err)
	}
}

func funcPoints(xPoints []float64, yPoints []float64, k int) plotter.XYs {
	pts := make(plotter.XYs, k)
	for i := range pts {
		pts[i].X = xPoints[i]
		pts[i].Y = yPoints[i]
	}
	return pts
}
