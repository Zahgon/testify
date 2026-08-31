package require

import (
	http "net/http"
	url "net/url"
	time "time"

	assert "github.com/stretchr/testify/assert"
)

func Condition(t TestingT, comp assert.Comparison, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Conditionf(t TestingT, comp assert.Comparison, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Contains(t TestingT, s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Containsf(t TestingT, s interface{}, contains interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func DirExists(t TestingT, path string, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func DirExistsf(t TestingT, path string, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func ElementsMatch(t TestingT, listA interface{}, listB interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func ElementsMatchf(t TestingT, listA interface{}, listB interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Empty(t TestingT, object interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Emptyf(t TestingT, object interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Equal(t TestingT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func EqualError(t TestingT, theError error, errString string, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func EqualErrorf(t TestingT, theError error, errString string, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func EqualExportedValues(t TestingT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func EqualExportedValuesf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func EqualValues(t TestingT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func EqualValuesf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Equalf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Error(t TestingT, err error, msgAndArgs ...interface{}) { _ = "STUB: not implemented"; return }

func ErrorAs(t TestingT, err error, target interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func ErrorAsf(t TestingT, err error, target interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func ErrorContains(t TestingT, theError error, contains string, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func ErrorContainsf(t TestingT, theError error, contains string, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func ErrorIs(t TestingT, err error, target error, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func ErrorIsf(t TestingT, err error, target error, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Errorf(t TestingT, err error, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Eventually(t TestingT, condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func EventuallyWithT(t TestingT, condition func(collect *assert.CollectT), waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func EventuallyWithTf(t TestingT, condition func(collect *assert.CollectT), waitFor time.Duration, tick time.Duration, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Eventuallyf(t TestingT, condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Exactly(t TestingT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Exactlyf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Fail(t TestingT, failureMessage string, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func FailNow(t TestingT, failureMessage string, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func FailNowf(t TestingT, failureMessage string, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Failf(t TestingT, failureMessage string, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func False(t TestingT, value bool, msgAndArgs ...interface{}) { _ = "STUB: not implemented"; return }

func Falsef(t TestingT, value bool, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func FileExists(t TestingT, path string, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func FileExistsf(t TestingT, path string, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Greater(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func GreaterOrEqual(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func GreaterOrEqualf(t TestingT, e1 interface{}, e2 interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Greaterf(t TestingT, e1 interface{}, e2 interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func HTTPBodyContains(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func HTTPBodyContainsf(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func HTTPBodyNotContains(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func HTTPBodyNotContainsf(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func HTTPError(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func HTTPErrorf(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func HTTPRedirect(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func HTTPRedirectf(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func HTTPStatusCode(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func HTTPStatusCodef(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func HTTPSuccess(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func HTTPSuccessf(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Implements(t TestingT, interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Implementsf(t TestingT, interfaceObject interface{}, object interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func InDelta(t TestingT, expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func InDeltaMapValues(t TestingT, expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func InDeltaMapValuesf(t TestingT, expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func InDeltaSlice(t TestingT, expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func InDeltaSlicef(t TestingT, expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func InDeltaf(t TestingT, expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func InEpsilon(t TestingT, expected interface{}, actual interface{}, epsilon float64, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func InEpsilonSlice(t TestingT, expected interface{}, actual interface{}, epsilon float64, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func InEpsilonSlicef(t TestingT, expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func InEpsilonf(t TestingT, expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func IsDecreasing(t TestingT, object interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func IsDecreasingf(t TestingT, object interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func IsIncreasing(t TestingT, object interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func IsIncreasingf(t TestingT, object interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func IsNonDecreasing(t TestingT, object interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func IsNonDecreasingf(t TestingT, object interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func IsNonIncreasing(t TestingT, object interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func IsNonIncreasingf(t TestingT, object interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func IsNotType(t TestingT, theType interface{}, object interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func IsNotTypef(t TestingT, theType interface{}, object interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func IsType(t TestingT, expectedType interface{}, object interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func IsTypef(t TestingT, expectedType interface{}, object interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func JSONEq(t TestingT, expected string, actual string, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func JSONEqf(t TestingT, expected string, actual string, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Len(t TestingT, object interface{}, length int, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Lenf(t TestingT, object interface{}, length int, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Less(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func LessOrEqual(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func LessOrEqualf(t TestingT, e1 interface{}, e2 interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Lessf(t TestingT, e1 interface{}, e2 interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Negative(t TestingT, e interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Negativef(t TestingT, e interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Never(t TestingT, condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Neverf(t TestingT, condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Nil(t TestingT, object interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Nilf(t TestingT, object interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NoDirExists(t TestingT, path string, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NoDirExistsf(t TestingT, path string, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NoError(t TestingT, err error, msgAndArgs ...interface{}) { _ = "STUB: not implemented"; return }

func NoErrorf(t TestingT, err error, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NoFileExists(t TestingT, path string, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NoFileExistsf(t TestingT, path string, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NotContains(t TestingT, s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NotContainsf(t TestingT, s interface{}, contains interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NotElementsMatch(t TestingT, listA interface{}, listB interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NotElementsMatchf(t TestingT, listA interface{}, listB interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NotEmpty(t TestingT, object interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NotEmptyf(t TestingT, object interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NotEqual(t TestingT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NotEqualValues(t TestingT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NotEqualValuesf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NotEqualf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NotErrorAs(t TestingT, err error, target interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NotErrorAsf(t TestingT, err error, target interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NotErrorIs(t TestingT, err error, target error, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NotErrorIsf(t TestingT, err error, target error, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NotImplements(t TestingT, interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NotImplementsf(t TestingT, interfaceObject interface{}, object interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NotNil(t TestingT, object interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NotNilf(t TestingT, object interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NotPanics(t TestingT, f assert.PanicTestFunc, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NotPanicsf(t TestingT, f assert.PanicTestFunc, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NotRegexp(t TestingT, rx interface{}, str interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NotRegexpf(t TestingT, rx interface{}, str interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NotSame(t TestingT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NotSamef(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NotSubset(t TestingT, list interface{}, subset interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NotSubsetf(t TestingT, list interface{}, subset interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NotZero(t TestingT, i interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func NotZerof(t TestingT, i interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Panics(t TestingT, f assert.PanicTestFunc, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func PanicsWithError(t TestingT, errString string, f assert.PanicTestFunc, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func PanicsWithErrorf(t TestingT, errString string, f assert.PanicTestFunc, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func PanicsWithValue(t TestingT, expected interface{}, f assert.PanicTestFunc, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func PanicsWithValuef(t TestingT, expected interface{}, f assert.PanicTestFunc, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Panicsf(t TestingT, f assert.PanicTestFunc, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Positive(t TestingT, e interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Positivef(t TestingT, e interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Regexp(t TestingT, rx interface{}, str interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Regexpf(t TestingT, rx interface{}, str interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Same(t TestingT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Samef(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Subset(t TestingT, list interface{}, subset interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Subsetf(t TestingT, list interface{}, subset interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func True(t TestingT, value bool, msgAndArgs ...interface{}) { _ = "STUB: not implemented"; return }

func Truef(t TestingT, value bool, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func WithinDuration(t TestingT, expected time.Time, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func WithinDurationf(t TestingT, expected time.Time, actual time.Time, delta time.Duration, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func WithinRange(t TestingT, actual time.Time, start time.Time, end time.Time, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func WithinRangef(t TestingT, actual time.Time, start time.Time, end time.Time, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func YAMLEq(t TestingT, expected string, actual string, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func YAMLEqf(t TestingT, expected string, actual string, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Zero(t TestingT, i interface{}, msgAndArgs ...interface{}) { _ = "STUB: not implemented"; return }

func Zerof(t TestingT, i interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}
