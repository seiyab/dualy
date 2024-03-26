package dualy

import "gonum.org/v1/plot"

type SidePlot struct {
	plotters []plot.Plotter
}

func (s *SidePlot) Add(p plot.Plotter) {
	s.plotters = append(s.plotters, p)
}
