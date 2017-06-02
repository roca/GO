package pack

import "testing"

func TestIntegrateConst(t *testing.T) {
	pi := PolyIntegrator{}

	result := pi.Integrate(0, 10, 3)

	if result != 30 {
		t.Error("Failed to integrate constant polynomial")
	}
}
