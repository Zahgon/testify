package assert

import (
	"reflect"
	"time"

	"github.com/stretchr/testify/internal/spew"
)

//go:generate sh -c "cd ../_codegen && go build && cd - && ../_codegen/_codegen -output-package=assert -template=assertion_format.go.tmpl"

type TestingT interface {
	Errorf(format string, args ...interface{})
}

type ComparisonAssertionFunc = func(TestingT, interface{}, interface{}, ...interface{}) bool

type ValueAssertionFunc = func(TestingT, interface{}, ...interface{}) bool

type BoolAssertionFunc = func(TestingT, bool, ...interface{}) bool

type ErrorAssertionFunc = func(TestingT, error, ...interface{}) bool

type PanicAssertionFunc = func(t TestingT, f PanicTestFunc, msgAndArgs ...interface{}) bool

type Comparison func() (success bool)

func ObjectsAreEqual(expected, actual interface{}) bool { _ = "STUB: not implemented"; return false }

func copyExportedFields(expected interface{}) interface{} { _ = "STUB: not implemented"; return nil }

func ObjectsExportedFieldsAreEqual(expected, actual interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func ObjectsAreEqualValues(expected, actual interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func isNumericType(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

func CallerInfo() []string { _ = "STUB: not implemented"; return nil }

func isTest(name, prefix string) bool { _ = "STUB: not implemented"; return false }

func messageFromMsgAndArgs(msgAndArgs ...interface{}) string { _ = "STUB: not implemented"; return "" }

func indentMessageLines(message string, longestLabelLen int) string {
	_ = "STUB: not implemented"
	return ""
}

type failNower interface {
	FailNow()
}

func FailNow(t TestingT, failureMessage string, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Fail(t TestingT, failureMessage string, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

type labeledContent struct {
	label   string
	content string
}

func labeledOutput(content ...labeledContent) string { _ = "STUB: not implemented"; return "" }

func Implements(t TestingT, interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NotImplements(t TestingT, interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func isType(expectedType, object interface{}) bool { _ = "STUB: not implemented"; return false }

func IsType(t TestingT, expectedType, object interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func IsNotType(t TestingT, theType, object interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Equal(t TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func validateEqualArgs(expected, actual interface{}) error { _ = "STUB: not implemented"; return nil }

func Same(t TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NotSame(t TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func samePointers(first, second interface{}) (same bool, ok bool) {
	_ = "STUB: not implemented"
	return false, false
}

func formatUnequalValues(expected, actual interface{}) (e string, a string) {
	_ = "STUB: not implemented"
	return "", ""
}

func truncatingFormat(format string, data interface{}) string { _ = "STUB: not implemented"; return "" }

func EqualValues(t TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func EqualExportedValues(t TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Exactly(t TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NotNil(t TestingT, object interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func isNil(object interface{}) bool { _ = "STUB: not implemented"; return false }

func Nil(t TestingT, object interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func isEmpty(object interface{}) bool { _ = "STUB: not implemented"; return false }

func isEmptyValue(objValue reflect.Value) bool { _ = "STUB: not implemented"; return false }

func Empty(t TestingT, object interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NotEmpty(t TestingT, object interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func getLen(x interface{}) (length int, ok bool) { _ = "STUB: not implemented"; return 0, false }

func Len(t TestingT, object interface{}, length int, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func True(t TestingT, value bool, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func False(t TestingT, value bool, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NotEqual(t TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NotEqualValues(t TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func containsElement(list interface{}, element interface{}) (ok, found bool) {
	_ = "STUB: not implemented"
	return false, false
}

func Contains(t TestingT, s, contains interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NotContains(t TestingT, s, contains interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Subset(t TestingT, list, subset interface{}, msgAndArgs ...interface{}) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

func NotSubset(t TestingT, list, subset interface{}, msgAndArgs ...interface{}) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

func ElementsMatch(t TestingT, listA, listB interface{}, msgAndArgs ...interface{}) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

func isList(t TestingT, list interface{}, msgAndArgs ...interface{}) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

func diffLists(listA, listB interface{}) (extraA, extraB []interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}

func formatListDiff(listA, listB interface{}, extraA, extraB []interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

func NotElementsMatch(t TestingT, listA, listB interface{}, msgAndArgs ...interface{}) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

func Condition(t TestingT, comp Comparison, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

type PanicTestFunc func()

func didPanic(f PanicTestFunc) (didPanic bool, message interface{}, stack string) {
	_ = "STUB: not implemented"
	return false, nil, ""
}

func Panics(t TestingT, f PanicTestFunc, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func PanicsWithValue(t TestingT, expected interface{}, f PanicTestFunc, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func PanicsWithError(t TestingT, errString string, f PanicTestFunc, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NotPanics(t TestingT, f PanicTestFunc, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func WithinDuration(t TestingT, expected, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func WithinRange(t TestingT, actual, start, end time.Time, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func toFloat(x interface{}) (float64, bool) { _ = "STUB: not implemented"; return 0, false }

func InDelta(t TestingT, expected, actual interface{}, delta float64, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func InDeltaSlice(t TestingT, expected, actual interface{}, delta float64, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func InDeltaMapValues(t TestingT, expected, actual interface{}, delta float64, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func calcRelativeError(expected, actual interface{}) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func InEpsilon(t TestingT, expected, actual interface{}, epsilon float64, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func InEpsilonSlice(t TestingT, expected, actual interface{}, epsilon float64, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NoError(t TestingT, err error, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Error(t TestingT, err error, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func EqualError(t TestingT, theError error, errString string, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func ErrorContains(t TestingT, theError error, contains string, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func matchRegexp(rx interface{}, str interface{}) bool { _ = "STUB: not implemented"; return false }

func Regexp(t TestingT, rx interface{}, str interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NotRegexp(t TestingT, rx interface{}, str interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Zero(t TestingT, i interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NotZero(t TestingT, i interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func FileExists(t TestingT, path string, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NoFileExists(t TestingT, path string, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func DirExists(t TestingT, path string, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NoDirExists(t TestingT, path string, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func JSONEq(t TestingT, expected string, actual string, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func YAMLEq(t TestingT, expected string, actual string, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func typeAndKind(v interface{}) (reflect.Type, reflect.Kind) {
	_ = "STUB: not implemented"
	return *new(reflect.Type), *new(reflect.Kind)
}

func diff(expected interface{}, actual interface{}) string { _ = "STUB: not implemented"; return "" }

func isFunction(arg interface{}) bool { _ = "STUB: not implemented"; return false }

var spewConfig = spew.ConfigState{
	Indent:                  " ",
	DisablePointerAddresses: true,
	DisableCapacities:       true,
	SortKeys:                true,
	DisableMethods:          true,
	MaxDepth:                10,
}

var spewConfigStringerEnabled = spew.ConfigState{
	Indent:                  " ",
	DisablePointerAddresses: true,
	DisableCapacities:       true,
	SortKeys:                true,
	MaxDepth:                10,
}

type tHelper = interface {
	Helper()
}

func Eventually(t TestingT, condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

type CollectT struct {
	errors []error
}

func (CollectT) Helper() { _ = "STUB: not implemented"; return }

func (c *CollectT) Errorf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (c *CollectT) FailNow() { _ = "STUB: not implemented"; return }

func (*CollectT) Reset() { _ = "STUB: not implemented"; return }

func (*CollectT) Copy(TestingT) { _ = "STUB: not implemented"; return }

func (c *CollectT) fail() { _ = "STUB: not implemented"; return }

func (c *CollectT) failed() bool { _ = "STUB: not implemented"; return false }

func EventuallyWithT(t TestingT, condition func(collect *CollectT), waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Never(t TestingT, condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func ErrorIs(t TestingT, err, target error, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NotErrorIs(t TestingT, err, target error, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func ErrorAs(t TestingT, err error, target interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func NotErrorAs(t TestingT, err error, target interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func unwrapAll(err error) (errs []error) { _ = "STUB: not implemented"; return nil }

func buildErrorChainString(err error, withType bool) string { _ = "STUB: not implemented"; return "" }
