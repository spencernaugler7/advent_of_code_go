package envpattern

type Env struct {
	Logger
	*Registrar
}

func Environment(logger Logger, registrar *Registrar) *Env {
	if logger == nil {
		logger = NullLogger{}
	}

	return &Env{Logger: logger, Registrar: registrar}
}

func Null() *Env {
	return Environment(nil, nil)
}
