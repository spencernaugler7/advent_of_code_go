package envpattern

import "sync"

type Registrar struct {
	values map[string]interface{}
	sync.Mutex
}

func (r *Registrar) Register(name string, value interface{}) error {
	if r == nil {
		return nil
	}

	r.Lock()
	defer r.Unlock()

	r.values[name] = value

	return nil
}

func (r *Registrar) Lookup(name string) (interface{}, error) {
	if r == nil {
		return nil, 
	}

	r.Lock()
	value := r.values[name]
	r.Unlock()
	return value, nil
}