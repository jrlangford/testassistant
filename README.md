# toolbox
--
    import "github.com/jrlangford/toolbox"

Package toolbox provides functions that simplify the test writing process and
provide a standardized way to report failure.

## Usage

```go
var Debug = false
```
Debug enables output printing when a test has not failed if it is set to 'true'

#### func  FailOnFloatMismatch

```go
func FailOnFloatMismatch(expectedFloat, receivedFloat float64, t *testing.T)
```
FailOnFloatMismatch compares two floating point numbers and reports failure if
thier difference exceeds 1e-9.

#### func  FailOnFloatToleranceExceedance

```go
func FailOnFloatToleranceExceedance(expectedFloat, tolerancePercentage, receivedFloat float64, t *testing.T)
```
FailOnFloatToleranceExceedance compares two floating point numbers and reports
failure if the error percentage exceeds tolerancePercentage.

#### func  FailOnIntMismatch

```go
func FailOnIntMismatch(expectedInt, receivedInt int, t *testing.T)
```
FailOnIntMismatch compares two ints and reports failure if they are different.

#### func  FailOnRegexMismatch

```go
func FailOnRegexMismatch(expectedPattern, receivedString string, t *testing.T)
```
FailOnRegexMismatch compares a string with a regex pattern and reports failure
if they do not match.

#### func  FailOnStringMismatch

```go
func FailOnStringMismatch(expectedString, receivedString string, t *testing.T)
```
FailOnStringMismatch compares two strings and reports failure if they are not
equal.
