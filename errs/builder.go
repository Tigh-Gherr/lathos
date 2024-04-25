package errs

import "github.com/theflyingcodr/lathos"

type Builder struct {
	code   string
	detail string
	fn     func(code, detail string) error
}

func Build() *Builder {
	return &Builder{}
}

func (b *Builder) NotFound() *Builder {
	b.fn = wrap(NewErrNotFound)
	return b
}

func (b *Builder) NotDuplicate() *Builder {
	b.fn = wrap(NewErrDuplicate)
	return b
}

func (b *Builder) NotNotAuthenticated() *Builder {
	b.fn = wrap(NewErrNotAuthenticated)
	return b
}

func (b *Builder) NotNotAuthorised() *Builder {
	b.fn = wrap(NewErrNotAuthorised)
	return b
}

func (b *Builder) NotAvailable() *Builder {
	b.fn = wrap(NewErrNotAvailable)
	return b
}

func (b *Builder) Unprocessable() *Builder {
	b.fn = wrap(NewErrUnprocessable)
	return b
}

func (b *Builder) TooManyRequests() *Builder {
	b.fn = wrap(NewErrTooManyRequests)
	return b
}

func (b *Builder) Conflict() *Builder {
	b.fn = wrap(NewErrConflict)
	return b
}

func wrap[T lathos.ClientError](fn func(string, string) T) func(string, string) error {
	return func(code, detail string) error {
		return fn(code, detail)
	}
}
