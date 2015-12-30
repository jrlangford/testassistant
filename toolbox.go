package toolbox

import (
	"math"
	"testing"
	"fmt"
)

var Debug = false

func FailOnFloatMismatch(expectedFloat, receivedFloat float64, t *testing.T) {
	epsilon := 1e-9
	outputString := fmt.Sprintf("\nExpected float: %.20f\nReceived float: %.20f\nEpsilon:%g\n", expectedFloat, receivedFloat, epsilon)
	if math.Abs(expectedFloat-receivedFloat) >= epsilon || math.IsNaN(receivedFloat) {
		t.Error(outputString)
	} else if Debug {
		fmt.Print(outputString)
	}
}

func FailOnFloatToleranceExceedance(expectedFloat, tolerancePercentage, receivedFloat float64, t *testing.T) {
	error := 1 - receivedFloat/expectedFloat
	errorPercentage := 100 * error
	outputString := fmt.Sprintf("\nExpected float: %.20f\nReceived float: %.20f\nTolerance:%g%%\nError:%g%% (%f), \n", expectedFloat, receivedFloat, tolerancePercentage, errorPercentage, error)
	if math.Abs(errorPercentage) > tolerancePercentage || math.IsNaN(receivedFloat) {
		t.Error(outputString)
	} else if Debug {
		fmt.Print(outputString)
	}
}

func FailOnIntMismatch(expectedInt, receivedInt int, t *testing.T) {
	outputString := fmt.Sprintf("\nExpected int: %d\nReceived int: %d\n\n", expectedInt, receivedInt)
	if receivedInt != expectedInt {
		t.Error(outputString)
	} else if Debug {
		fmt.Print(outputString)
	}
}
