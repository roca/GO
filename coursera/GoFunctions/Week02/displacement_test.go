package main

import "testing"

func TestGenDisplaceFn(t *testing.T) {
	fn := GenDisplaceFn(10, 2, 1)
	displacement := fn(3)

	expected_dispalcement := 52.0

	if displacement != expected_dispalcement {
		t.Errorf("%v != %v", displacement, expected_dispalcement)
	}
}

func TestConvertStringToFloats(t *testing.T) {
	floats := ConvertStringToFloats("0 1 2 3 4 ")
	if len(floats) != 5 {
		t.Errorf("floats is oncrrect length %d insted of 5", len(floats))
	}

	for i, v := range floats {
		if v != float64(i) {
			t.Errorf("%f != %f\n", v, float64(i))
		}
	}

}
