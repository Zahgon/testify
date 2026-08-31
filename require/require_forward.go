package require

import (
	http "net/http"
	url "net/url"
	time "time"

	assert "github.com/stretchr/testify/assert"
)

func (a *Assertions) Condition(comp assert.Comparison, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Conditionf(comp assert.Comparison, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Contains(s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Containsf(s interface{}, contains interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) DirExists(path string, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) DirExistsf(path string, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) ElementsMatch(listA interface{}, listB interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) ElementsMatchf(listA interface{}, listB interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Empty(object interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Emptyf(object interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Equal(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) EqualError(theError error, errString string, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) EqualErrorf(theError error, errString string, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) EqualExportedValues(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) EqualExportedValuesf(expected interface{}, actual interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) EqualValues(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) EqualValuesf(expected interface{}, actual interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Equalf(expected interface{}, actual interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Error(err error, msgAndArgs ...interface{}) { _ = "STUB: not implemented"; return }

func (a *Assertions) ErrorAs(err error, target interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) ErrorAsf(err error, target interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) ErrorContains(theError error, contains string, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) ErrorContainsf(theError error, contains string, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) ErrorIs(err error, target error, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) ErrorIsf(err error, target error, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Errorf(err error, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Eventually(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) EventuallyWithT(condition func(collect *assert.CollectT), waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) EventuallyWithTf(condition func(collect *assert.CollectT), waitFor time.Duration, tick time.Duration, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Eventuallyf(condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Exactly(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Exactlyf(expected interface{}, actual interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Fail(failureMessage string, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) FailNow(failureMessage string, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) FailNowf(failureMessage string, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Failf(failureMessage string, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) False(value bool, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Falsef(value bool, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) FileExists(path string, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) FileExistsf(path string, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Greater(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) GreaterOrEqual(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) GreaterOrEqualf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Greaterf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) HTTPBodyContains(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) HTTPBodyContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) HTTPBodyNotContains(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) HTTPBodyNotContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) HTTPError(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) HTTPErrorf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) HTTPRedirect(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) HTTPRedirectf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) HTTPStatusCode(handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) HTTPStatusCodef(handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) HTTPSuccess(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) HTTPSuccessf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Implements(interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Implementsf(interfaceObject interface{}, object interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) InDelta(expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) InDeltaMapValues(expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) InDeltaMapValuesf(expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) InDeltaSlice(expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) InDeltaSlicef(expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) InDeltaf(expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) InEpsilon(expected interface{}, actual interface{}, epsilon float64, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) InEpsilonSlice(expected interface{}, actual interface{}, epsilon float64, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) InEpsilonSlicef(expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) InEpsilonf(expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) IsDecreasing(object interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) IsDecreasingf(object interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) IsIncreasing(object interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) IsIncreasingf(object interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) IsNonDecreasing(object interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) IsNonDecreasingf(object interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) IsNonIncreasing(object interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) IsNonIncreasingf(object interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) IsNotType(theType interface{}, object interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) IsNotTypef(theType interface{}, object interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) IsType(expectedType interface{}, object interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) IsTypef(expectedType interface{}, object interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) JSONEq(expected string, actual string, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) JSONEqf(expected string, actual string, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Len(object interface{}, length int, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Lenf(object interface{}, length int, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Less(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) LessOrEqual(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) LessOrEqualf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Lessf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Negative(e interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Negativef(e interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Never(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Neverf(condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Nil(object interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Nilf(object interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NoDirExists(path string, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NoDirExistsf(path string, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NoError(err error, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NoErrorf(err error, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NoFileExists(path string, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NoFileExistsf(path string, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NotContains(s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NotContainsf(s interface{}, contains interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NotElementsMatch(listA interface{}, listB interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NotElementsMatchf(listA interface{}, listB interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NotEmpty(object interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NotEmptyf(object interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NotEqual(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NotEqualValues(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NotEqualValuesf(expected interface{}, actual interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NotEqualf(expected interface{}, actual interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NotErrorAs(err error, target interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NotErrorAsf(err error, target interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NotErrorIs(err error, target error, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NotErrorIsf(err error, target error, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NotImplements(interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NotImplementsf(interfaceObject interface{}, object interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NotNil(object interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NotNilf(object interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NotPanics(f assert.PanicTestFunc, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NotPanicsf(f assert.PanicTestFunc, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NotRegexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NotRegexpf(rx interface{}, str interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NotSame(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NotSamef(expected interface{}, actual interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NotSubset(list interface{}, subset interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NotSubsetf(list interface{}, subset interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NotZero(i interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) NotZerof(i interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Panics(f assert.PanicTestFunc, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) PanicsWithError(errString string, f assert.PanicTestFunc, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) PanicsWithErrorf(errString string, f assert.PanicTestFunc, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) PanicsWithValue(expected interface{}, f assert.PanicTestFunc, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) PanicsWithValuef(expected interface{}, f assert.PanicTestFunc, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Panicsf(f assert.PanicTestFunc, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Positive(e interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Positivef(e interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Regexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Regexpf(rx interface{}, str interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Same(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Samef(expected interface{}, actual interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Subset(list interface{}, subset interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Subsetf(list interface{}, subset interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) True(value bool, msgAndArgs ...interface{}) { _ = "STUB: not implemented"; return }

func (a *Assertions) Truef(value bool, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) WithinDuration(expected time.Time, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) WithinDurationf(expected time.Time, actual time.Time, delta time.Duration, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) WithinRange(actual time.Time, start time.Time, end time.Time, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) WithinRangef(actual time.Time, start time.Time, end time.Time, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) YAMLEq(expected string, actual string, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) YAMLEqf(expected string, actual string, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Zero(i interface{}, msgAndArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a *Assertions) Zerof(i interface{}, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}
