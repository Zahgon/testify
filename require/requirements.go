package require

type TestingT interface {
	Errorf(format string, args ...interface{})
	FailNow()
}

type tHelper = interface {
	Helper()
}

type ComparisonAssertionFunc = func(TestingT, interface{}, interface{}, ...interface{})

type ValueAssertionFunc = func(TestingT, interface{}, ...interface{})

type BoolAssertionFunc = func(TestingT, bool, ...interface{})

type ErrorAssertionFunc = func(TestingT, error, ...interface{})

//go:generate sh -c "cd ../_codegen && go build && cd - && ../_codegen/_codegen -output-package=require -template=require.go.tmpl -include-format-funcs"
