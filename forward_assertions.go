package assert

// Assertions provides assertion methods around the
// TestingT interface.
type Assertions struct {
	t         TestingT
	failRerun int
}

// New makes a new Assertions object for the specified TestingT.
func New(t TestingT, options ...CliOption) *Assertions {
	as := &Assertions{
		t:         t,
		failRerun: 1,
	}
	for _, opt := range options {
		opt(as)
	}
	return as
}

type CliOption func(r *Assertions)

func WithFailRerun(times int) CliOption {
	return func(r *Assertions) {
		r.failRerun = times
	}
}

//go:generate sh -c "cd ./.github/_codegen && go build && cd - && ./.github/_codegen/_codegen -output-package=assert -template=assertion_forward.go.tmpl -include-format-funcs && gofumpt -l -w ."
