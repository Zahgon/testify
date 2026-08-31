package assert

import (
	http "net/http"
	url "net/url"
	time "time"
)

func Conditionf(t TestingT, comp Comparison, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Containsf(t TestingT, s interface{}, contains interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func DirExistsf(t TestingT, path string, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func ElementsMatchf(t TestingT, listA interface{}, listB interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Emptyf(t TestingT, object interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Equalf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func EqualErrorf(t TestingT, theError error, errString string, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func EqualExportedValuesf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func EqualValuesf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Errorf(t TestingT, err error, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func ErrorAsf(t TestingT, err error, target interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func ErrorContainsf(t TestingT, theError error, contains string, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func ErrorIsf(t TestingT, err error, target error, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Eventuallyf(t TestingT, condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func EventuallyWithTf(t TestingT, condition func(collect *CollectT), waitFor time.Duration, tick time.Duration, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Exactlyf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Failf(t TestingT, failureMessage string, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func FailNowf(t TestingT, failureMessage string, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Falsef(t TestingT, value bool, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func FileExistsf(t TestingT, path string, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Greaterf(t TestingT, e1 interface{}, e2 interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func GreaterOrEqualf(t TestingT, e1 interface{}, e2 interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func HTTPBodyContainsf(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func HTTPBodyNotContainsf(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func HTTPErrorf(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func HTTPRedirectf(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func HTTPStatusCodef(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func HTTPSuccessf(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Implementsf(t TestingT, interfaceObject interface{}, object interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func InDeltaf(t TestingT, expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func InDeltaMapValuesf(t TestingT, expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func InDeltaSlicef(t TestingT, expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func InEpsilonf(t TestingT, expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func InEpsilonSlicef(t TestingT, expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func IsDecreasingf(t TestingT, object interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func IsIncreasingf(t TestingT, object interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func IsNonDecreasingf(t TestingT, object interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func IsNonIncreasingf(t TestingT, object interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func IsNotTypef(t TestingT, theType interface{}, object interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func IsTypef(t TestingT, expectedType interface{}, object interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func JSONEqf(t TestingT, expected string, actual string, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Lenf(t TestingT, object interface{}, length int, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Lessf(t TestingT, e1 interface{}, e2 interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func LessOrEqualf(t TestingT, e1 interface{}, e2 interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Negativef(t TestingT, e interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Neverf(t TestingT, condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Nilf(t TestingT, object interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NoDirExistsf(t TestingT, path string, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NoErrorf(t TestingT, err error, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NoFileExistsf(t TestingT, path string, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NotContainsf(t TestingT, s interface{}, contains interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NotElementsMatchf(t TestingT, listA interface{}, listB interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NotEmptyf(t TestingT, object interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NotEqualf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NotEqualValuesf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NotErrorAsf(t TestingT, err error, target interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NotErrorIsf(t TestingT, err error, target error, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NotImplementsf(t TestingT, interfaceObject interface{}, object interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NotNilf(t TestingT, object interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NotPanicsf(t TestingT, f PanicTestFunc, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NotRegexpf(t TestingT, rx interface{}, str interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NotSamef(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NotSubsetf(t TestingT, list interface{}, subset interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NotZerof(t TestingT, i interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Panicsf(t TestingT, f PanicTestFunc, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func PanicsWithErrorf(t TestingT, errString string, f PanicTestFunc, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func PanicsWithValuef(t TestingT, expected interface{}, f PanicTestFunc, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Positivef(t TestingT, e interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Regexpf(t TestingT, rx interface{}, str interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Samef(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Subsetf(t TestingT, list interface{}, subset interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Truef(t TestingT, value bool, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func WithinDurationf(t TestingT, expected time.Time, actual time.Time, delta time.Duration, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func WithinRangef(t TestingT, actual time.Time, start time.Time, end time.Time, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func YAMLEqf(t TestingT, expected string, actual string, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Zerof(t TestingT, i interface{}, msg string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}
