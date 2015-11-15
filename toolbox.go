package toolbox

import (
	"math"
	"testing"
)

func FailOnFloatMismatch(expectedFloat, receivedFloat float64, t *testing.T) {
	epsilon := 1e-9
	if math.Abs(expectedFloat-receivedFloat) >= epsilon {
		t.Errorf("\nExpected float: %.20f\nReceived float: %.20f\nEpsilon:%g\n", expectedFloat, receivedFloat, epsilon)
	}
}

func FailOnFloatToleranceExceedance(expectedFloat, tolerancePercentage, receivedFloat float64, t *testing.T) {
	error := 1 - receivedFloat/expectedFloat
	errorPercentage := 100 * error
	if errorPercentage > tolerancePercentage {
		t.Errorf("\nExpected float: %.20f\nReceived float: %.20f\nTolerance:%f%%\nError:%f%% (%f), \n", expectedFloat, receivedFloat, tolerancePercentage, errorPercentage, error)
	}
}
