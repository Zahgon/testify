package mock

import (
	"reflect"
	"regexp"
	"runtime"
	"sync"
	"time"

	"github.com/stretchr/objx"

	"github.com/stretchr/testify/internal/spew"
)

var gccgoRE = regexp.MustCompile(`\.pN\d+_`)

type TestingT interface {
	Logf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	FailNow()
}

type Call struct {
	Parent *Mock

	Method string

	Arguments Arguments

	ReturnArguments Arguments

	callerInfo []string

	Repeatability int

	totalCalls int

	optional bool

	WaitFor <-chan time.Time

	waitTime time.Duration

	RunFn func(Arguments)

	PanicMsg *string

	requires []*Call
}

func newCall(parent *Mock, methodName string, callerInfo []string, methodArguments Arguments, returnArguments Arguments) *Call {
	_ = "STUB: not implemented"
	return nil
}

func (c *Call) lock() { _ = "STUB: not implemented"; return }

func (c *Call) unlock() { _ = "STUB: not implemented"; return }

func (c *Call) Return(returnArguments ...interface{}) *Call { _ = "STUB: not implemented"; return nil }

func (c *Call) Panic(msg string) *Call { _ = "STUB: not implemented"; return nil }

func (c *Call) Once() *Call { _ = "STUB: not implemented"; return nil }

func (c *Call) Twice() *Call { _ = "STUB: not implemented"; return nil }

func (c *Call) Times(i int) *Call { _ = "STUB: not implemented"; return nil }

func (c *Call) WaitUntil(w <-chan time.Time) *Call { _ = "STUB: not implemented"; return nil }

func (c *Call) After(d time.Duration) *Call { _ = "STUB: not implemented"; return nil }

func (c *Call) Run(fn func(args Arguments)) *Call { _ = "STUB: not implemented"; return nil }

func (c *Call) Maybe() *Call { _ = "STUB: not implemented"; return nil }

//go:noinline
func (c *Call) On(methodName string, arguments ...interface{}) *Call {
	_ = "STUB: not implemented"
	return nil
}

func (c *Call) Unset() *Call { _ = "STUB: not implemented"; return nil }

func (c *Call) NotBefore(calls ...*Call) *Call { _ = "STUB: not implemented"; return nil }

func InOrder(calls ...*Call) { _ = "STUB: not implemented"; return }

type Mock struct {
	ExpectedCalls []*Call

	Calls []Call

	test TestingT

	testData objx.Map

	mutex sync.Mutex
}

func (m *Mock) String() string { _ = "STUB: not implemented"; return "" }

func (m *Mock) TestData() objx.Map { _ = "STUB: not implemented"; return *new(objx.Map) }

func (m *Mock) Test(t TestingT) { _ = "STUB: not implemented"; return }

func (m *Mock) fail(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (m *Mock) On(methodName string, arguments ...interface{}) *Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *Mock) findExpectedCall(method string, arguments ...interface{}) (int, *Call) {
	_ = "STUB: not implemented"
	return 0, nil
}

type matchCandidate struct {
	call      *Call
	mismatch  string
	diffCount int
}

func (c matchCandidate) isBetterMatchThan(other matchCandidate) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *Mock) findClosestCall(method string, arguments ...interface{}) (*Call, string) {
	_ = "STUB: not implemented"
	return nil, ""
}

func callString(method string, arguments Arguments, includeArgumentValues bool) string {
	_ = "STUB: not implemented"
	return ""
}

func (m *Mock) Called(arguments ...interface{}) Arguments {
	_ = "STUB: not implemented"
	return *new(Arguments)
}

func (m *Mock) MethodCalled(methodName string, arguments ...interface{}) Arguments {
	_ = "STUB: not implemented"
	return *new(Arguments)
}

type assertExpectationiser interface {
	AssertExpectations(TestingT) bool
}

func AssertExpectationsForObjects(t TestingT, testObjects ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *Mock) AssertExpectations(t TestingT) bool { _ = "STUB: not implemented"; return false }

func (m *Mock) checkExpectation(call *Call) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func (m *Mock) AssertNumberOfCalls(t TestingT, methodName string, expectedCalls int) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *Mock) AssertCalled(t TestingT, methodName string, arguments ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *Mock) AssertNotCalled(t TestingT, methodName string, arguments ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *Mock) IsMethodCallable(t TestingT, methodName string, arguments ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func isArgsEqual(expected Arguments, args []interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *Mock) methodWasCalled(methodName string, expected []interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *Mock) expectedCalls() []*Call { _ = "STUB: not implemented"; return nil }

func (m *Mock) calls() []Call { _ = "STUB: not implemented"; return nil }

type Arguments []interface{}

const (
	Anything = "mock.Anything"
)

type AnythingOfTypeArgument = anythingOfTypeArgument

type anythingOfTypeArgument string

func AnythingOfType(t string) AnythingOfTypeArgument {
	_ = "STUB: not implemented"
	return *new(AnythingOfTypeArgument)
}

type IsTypeArgument struct {
	t reflect.Type
}

func IsType(t interface{}) *IsTypeArgument { _ = "STUB: not implemented"; return nil }

type FunctionalOptionsArgument struct {
	values []interface{}
}

func (f *FunctionalOptionsArgument) String() string { _ = "STUB: not implemented"; return "" }

func FunctionalOptions(values ...interface{}) *FunctionalOptionsArgument {
	_ = "STUB: not implemented"
	return nil
}

type argumentMatcher struct {
	fn reflect.Value
}

func (f argumentMatcher) Matches(argument interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func (f argumentMatcher) String() string { _ = "STUB: not implemented"; return "" }

func MatchedBy(fn interface{}) argumentMatcher {
	_ = "STUB: not implemented"
	return *new(argumentMatcher)
}

func (args Arguments) Get(index int) interface{} { _ = "STUB: not implemented"; return nil }

func (args Arguments) Is(objects ...interface{}) bool { _ = "STUB: not implemented"; return false }

type missingArgument struct{}

func (missingArgument) String() string { _ = "STUB: not implemented"; return "" }

var missing missingArgument

func (args Arguments) Diff(objects []interface{}) (string, int) {
	_ = "STUB: not implemented"
	return "", 0
}

func (args Arguments) Assert(t TestingT, objects ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func (args Arguments) String(indexOrNil ...int) string { _ = "STUB: not implemented"; return "" }

func (args Arguments) Int(index int) int { _ = "STUB: not implemented"; return 0 }

func (args Arguments) Error(index int) error { _ = "STUB: not implemented"; return nil }

func (args Arguments) Bool(index int) bool { _ = "STUB: not implemented"; return false }

func safeTypeName(t reflect.Type) string { _ = "STUB: not implemented"; return "" }

func typeAndKind(v interface{}) (reflect.Type, reflect.Kind) {
	_ = "STUB: not implemented"
	return *new(reflect.Type), *new(reflect.Kind)
}

func diffArguments(expected Arguments, actual Arguments) string {
	_ = "STUB: not implemented"
	return ""
}

func diff(expected interface{}, actual interface{}) string { _ = "STUB: not implemented"; return "" }

var spewConfig = spew.ConfigState{
	Indent:                  " ",
	DisablePointerAddresses: true,
	DisableCapacities:       true,
	SortKeys:                true,
}

type tHelper interface {
	Helper()
}

func assertOpts(expected, actual interface{}) (expectedFmt, actualFmt string) {
	_ = "STUB: not implemented"
	return "", ""
}

func runtimeFunc(opt interface{}) *runtime.Func { _ = "STUB: not implemented"; return nil }

func funcName(f *runtime.Func) string { _ = "STUB: not implemented"; return "" }

func isFuncSame(f1, f2 *runtime.Func) bool { _ = "STUB: not implemented"; return false }
