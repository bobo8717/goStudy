package server

type Options struct {
	Codecs map[string]string
	Name   string
	Tag    string
}

type Option func(*Options)

// Server name
func Name(n string) Option {
	return func(o *Options) {
		o.Name = n
	}
}
func Tag(n string) Option {
	return func(o *Options) {
		o.Tag = n
	}
}

func newOptions(opt ...Option) Options {
	opts := Options{
		Codecs: make(map[string]string),
	}

	for _, o := range opt {
		o(&opts)
	}

	if opts.Tag == "" {
		opts.Tag = "default-Tag"
	}
	return opts
}
