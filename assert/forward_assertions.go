package assert

type Assertions struct {
	t TestingT
}

func New(t TestingT) *Assertions { _ = "STUB: not implemented"; return nil }

//go:generate sh -c "cd ../_codegen && go build && cd - && ../_codegen/_codegen -output-package=assert -template=assertion_forward.go.tmpl -include-format-funcs"
