package opt

type Option[T any] interface {
	apply(*T)
}

type OptionFunc[T any] func(*T)

func (f OptionFunc[T]) apply(o *T) {
	f(o)
}

func Each[T any](options *T, opts ...Option[T]) {
	for _, opt := range opts {
		opt.apply(options)
	}
}
