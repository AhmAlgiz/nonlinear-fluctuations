package main

import (
	"fluctuations/graphics"
	"fmt"
	m "math"
)

func calcPhi(a, o, t float64) float64 {
	s := m.Sqrt(o*o - a*a)
	return 2 * m.Atan((a-s*m.Tan((-1)*t*s)/2)/o)
}

func calcDiff(a, o, phi float64) float64 {
	return o - a*m.Sin(phi)
}

func calcDiff0(a, o float64) float64 {
	return m.Asin(o / a)
}

func printD(omega, a, step float64) {
	diffs := make([]float64, 0, int(m.Round(2*m.Pi/step)))
	phis := make([]float64, 0, int(m.Round(2*m.Pi/step)))

	for phi := 0.0; phi < 2*m.Pi; phi += step {
		diffs = append(diffs, calcDiff(a, omega, phi))
		phis = append(phis, phi)
	}

	graphics.PlotGraph(phis, diffs, fmt.Sprintf("Фазовый портрет для A = %.1f, Omega = %.1f", a, omega))

}

func printC(omega, a, step, tmin, tmax float64) {
	phis := make([]float64, 0, int(m.Round((tmax-tmin)/step)))
	ts := make([]float64, 0, int(m.Round((tmax-tmin)/step)))

	for t := tmin; t < tmax; t += step {
		phis = append(phis, calcPhi(a, omega, t))
		ts = append(ts, t)
	}

	graphics.PlotGraph(ts, phis, fmt.Sprintf("График решений для A = %.1f, Omega = %.1f", a, omega))

}

func main() {
	A := 10.0
	Omega := 1.0
	//printD(Omega, A, 0.1)
	printC(Omega, A, 0.1, -5, 5)
}
