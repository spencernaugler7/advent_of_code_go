package envpattern

import "sync"

type Registrar struct {
	values map[string]any
	sync.Mutex
}

func (r *Registrar) Register(name string, value any) error {
	if r == nil {
		return nil
	}

	r.Lock()
	defer r.Unlock()

	r.values[name] = value

	return nil
}

func (r *Registrar) Lookup(k string) any {
	if r == nil {
		return nil
	}

	r.Lock()
	defer r.Unlock()
	return r.values[k]
}
