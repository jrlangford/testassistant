package toolbox

import (
	"testing"
	"math"
)

func failOnFloatMismatch(expectedFloat, receivedFloat float64, t *testing.T) {
	epsilon := 1e-9
	if math.Abs(expectedFloat-receivedFloat) >= epsilon {
		t.Errorf("\nExpected float: %.20f\nReceived float: %.20f\nEpsilon:%g\n", expectedFloat, receivedFloat, epsilon)
	}
}
