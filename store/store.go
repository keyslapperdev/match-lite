package store

type memKV map[string]string

type Storer interface {
	Write(key, data string)
	Read(key string) string
}

func NewMemoryStore() *memKV {
	return new(memKV)
}

func (mk memKV) Write(key, data string) {
	mk[key] = data
}

func (mk memKV) Read(key string) string {
	return mk[key]
}
