package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"udemy.com/aml/dcm"
	"udemy.com/aml/euler"
	"udemy.com/aml/quaternion"
	"udemy.com/aml/vector"
)

func main() {
	attitude_xyz := euler.New(
		dcm.DegreesToRadians(0.0),
		dcm.DegreesToRadians(0.0),
		dcm.DegreesToRadians(0.0),
	)
	omega_body, _ := vector.New(
		dcm.DegreesToRadians(-1.),
		dcm.DegreesToRadians(15.),
		dcm.DegreesToRadians(-2.),
	)

	dt := 0.01
	timeValues := []float64{}
	phiValues := []float64{}
	thetaValues := []float64{}
	siValues := []float64{}

	q, _ := quaternion.Angles2Quat(attitude_xyz)

	for t := 0.; t < 20.+dt; t += dt {
		q_dot, _ := quaternion.KinematicRates_BodyRates(q, omega_body)
		q, _ = quaternion.Integrate(q, q_dot, dt)
		q.Normalize()
		attitude_new, _ := q.ToAngles("XYZ")
		timeValues = append(timeValues, t)
		phiValues = append(phiValues, dcm.RadiansToDegrees(attitude_new.Phi))
		thetaValues = append(thetaValues, dcm.RadiansToDegrees(attitude_new.Theta))
		siValues = append(siValues, dcm.RadiansToDegrees(attitude_new.Si))
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
