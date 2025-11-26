package protoc

type Options map[string][]string

func (o Options) Set(name, value string) error {
	o[name] = append(o[name], value)
	return nil
}

func (o Options) Get(name string) string {
	if val := o[name]; len(val) > 0 {
		return val[0]
	}
	return ""
}

func (o Options) Exist(name string) bool {
	_, ok := o[name]
	return ok
}

func (o Options) Gets(name string) []string {
	return o[name]
}

func NewOptions() Options { return Options{} }
