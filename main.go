package main

import (
	"fluctuations/graphics"
	"fmt"
	m "math"
	c "math/cmplx"
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

func calcPhiComplex(a, o, t float64) complex128 {
	ac := complex(a, 0)
	oc := complex(o, 0)
	tc := complex(t, 0)

	sc := c.Sqrt(oc*oc - ac*ac)
	res := 2 * c.Atan((ac-sc*c.Tan((-1)*tc*sc)/2)/oc)
	return res
}

func printD(omega, a, step float64) {
	diffs := make([]float64, 0, int(m.Round(2*m.Pi/step)))
	phis := make([]float64, 0, int(m.Round(2*m.Pi/step)))

	for phi := 0.0; phi < 2*m.Pi; phi += step {
		diffs = append(diffs, calcDiff(a, omega, phi))
		phis = append(phis, phi)
	}

	graphics.PlotGraph(phis, diffs, fmt.Sprintf("Фазовый портрет для A = %.2f, Omega = %.2f", a, omega))

}

func printC(omega, a, step, tmin, tmax float64) {
	phis := make([]float64, 0, int(m.Round((tmax-tmin)/step)))
	ts := make([]float64, 0, int(m.Round((tmax-tmin)/step)))

	for t := tmin; t < tmax; t += step {
		phis = append(phis, calcPhi(a, omega, t))
		ts = append(ts, t)
	}

	//graphics.PlotGraph(ts, phis, fmt.Sprintf("График решений для A = %.1f, Omega = %.1f", a, omega))
	graphics.PlotGraph(ts, phis, fmt.Sprintf("График решений"))
}

func complexPrintC(omega, a, step, tmin, tmax float64) {
	reals := make([]float64, 0, int(m.Round((tmax-tmin)/step)))
	imags := make([]float64, 0, int(m.Round((tmax-tmin)/step)))
	ts := make([]float64, 0, int(m.Round((tmax-tmin)/step)))

	for t := tmin; t < tmax; t += step {
		phi := calcPhiComplex(a, omega, t)
		reals = append(reals, real(phi))
		imags = append(imags, imag(phi))
		ts = append(ts, t)
	}

	//graphics.PlotGraph(ts, phis, fmt.Sprintf("График решений для A = %.1f, Omega = %.1f", a, omega))
	graphics.PlotGraph(ts, reals, fmt.Sprintf("График решений комплексный для A = %.2f, Omega = %.2f", a, omega))
}

func main() {
	A := 5.0
	Omega := 1.0
	printD(Omega, A, 0.1)
	//printC(Omega, A, 0.02, -5, 5)
	complexPrintC(Omega, A, 0.02, -5, 5)
}
