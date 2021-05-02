package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"udemy.com/aml/dcm"
	"udemy.com/aml/vector"
)

func main() {
	attitude, _ := vector.New(
		dcm.DegreesToRadians(0.0), // phi(Role Rate)
		dcm.DegreesToRadians(0.0), // theta(Pitch Rate)
		dcm.DegreesToRadians(0.0), // si(Yaw Rate)
	)
	omegaBody, _ := vector.New( // Angular Rates
		dcm.DegreesToRadians(-1.0), // phiDot
		dcm.DegreesToRadians(15.0), // thetaDot
		dcm.DegreesToRadians(-2.0), // siDot
	)

	dt := 0.01
	timeValues := []float64{}
	phiValues := []float64{}
	thetaValues := []float64{}
	siValues := []float64{}
	for t := 0.; t < 20.+dt; t += dt {
		attitudeDot, _ := dcm.XYZEulerAngleRates(attitude.X, attitude.Y, attitude.Z, omegaBody)
		attitude, _ = dcm.EulerIntergration(attitude, attitudeDot, dt)

		timeValues = append(timeValues, t)
		phiValues = append(phiValues, dcm.RadiansToDegrees(attitude.X))
		thetaValues = append(thetaValues, dcm.RadiansToDegrees(attitude.Y))
		siValues = append(siValues, dcm.RadiansToDegrees(attitude.Z))
	}

	p := plot.New()

	p.Title.Text = "Euler Angels over Time"
	p.X.Label.Text = "Time"
	p.Y.Label.Text = "Angle"

	err := plotutil.AddLines(p,
		"Roll", AnglesOverTime(phiValues, timeValues),
		"Pitch", AnglesOverTime(thetaValues, timeValues),
		"Yaw", AnglesOverTime(siValues, timeValues),
	)
	if err != nil {
		panic(err)
	}
	p.Legend.Left = true
	p.Legend.Top = true

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}
func AnglesOverTime(eulerAngles, timeValues []float64) plotter.XYs {
	pts := make(plotter.XYs, len(eulerAngles))
	for i := range pts {
		pts[i].X = timeValues[i]
		pts[i].Y = eulerAngles[i]
	}
	return pts
}
