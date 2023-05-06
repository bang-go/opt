package opt

type Option[T comparable] interface {
	apply(*T)
}

type OptionFunc[T comparable] func(*T)

func (f OptionFunc[T]) apply(o *T) {
	f(o)
}

func Each[T comparable](options *T, opts ...Option[T]) {
	for _, opt := range opts {
		opt.apply(options)
	}
}
