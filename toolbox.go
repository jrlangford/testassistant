//Package toolbox provides functions that simplify the test writing process and provide a standardized way to report failure.
package toolbox

import (
	"fmt"
	"math"
	"regexp"
	"testing"
)

//Debug enables output printing when a test has not failed if it is set to 'true'
var Debug = false

//FailOnFloatMismatch compares two floating point numbers and reports failure if thier difference exceeds 1e-9.
func FailOnFloatMismatch(expectedFloat, receivedFloat float64, t *testing.T) {
	epsilon := 1e-9
	outputString := fmt.Sprintf("\nExpected float: %.20f\nReceived float: %.20f\nEpsilon:%g\n", expectedFloat, receivedFloat, epsilon)
	if math.Abs(expectedFloat-receivedFloat) >= epsilon || math.IsNaN(receivedFloat) {
		t.Error(outputString)
	} else if Debug {
		fmt.Print(outputString)
	}
}

//FailOnFloatToleranceExceedance compares two floating point numbers and reports failure if the error percentage exceeds tolerancePercentage.
func FailOnFloatToleranceExceedance(expectedFloat, tolerancePercentage, receivedFloat float64, t *testing.T) {
	err := 1 - receivedFloat/expectedFloat
	errorPercentage := 100 * err
	outputString := fmt.Sprintf("\nExpected float: %.20f\nReceived float: %.20f\nTolerance:%g%%\nError:%g%% (%f), \n", expectedFloat, receivedFloat, tolerancePercentage, errorPercentage, err)
	if math.Abs(errorPercentage) > tolerancePercentage || math.IsNaN(receivedFloat) {
		t.Error(outputString)
	} else if Debug {
		fmt.Print(outputString)
	}
}

//FailOnIntMismatch compares two ints and reports failure if they are different.
func FailOnIntMismatch(expectedInt, receivedInt int, t *testing.T) {
	outputString := fmt.Sprintf("\nExpected int: %d\nReceived int: %d\n\n", expectedInt, receivedInt)
	if receivedInt != expectedInt {
		t.Error(outputString)
	} else if Debug {
		fmt.Print(outputString)
	}
}

//FailOnStringMismatch compares two strings and reports failure if they are not equal.
func FailOnStringMismatch(expectedString, receivedString string, t *testing.T) {
	if receivedString != expectedString {
		outputString := fmt.Sprintf(
			"String Mismatch\n--------------------------------------------------\n"+
				"Received String:\n%s\n--------------------------------------------------\n"+
				"Expected String:\n%s\n--------------------------------------------------\n",
			receivedString,
			expectedString,
		)
		t.Error(outputString)
	}
}

//FailOnRegexMismatch compares a string with a regex pattern and reports failure if they do not match.
func FailOnRegexMismatch(expectedPattern, receivedString string, t *testing.T) {
	matched, err := regexp.MatchString(expectedPattern, receivedString)
	if err != nil {
		t.Error(err.Error())
	}
	if !matched {
		outputString := fmt.Sprintf(
			"Regex Mismatch\n--------------------------------------------------\n"+
				"Received String:\n%s\n--------------------------------------------------\n"+
				"Expected Pattern:\n%s\n--------------------------------------------------\n",
			receivedString,
			expectedPattern,
		)
		t.Error(outputString)
	}
}
