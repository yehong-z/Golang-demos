package Immutable

type ImmutableMap interface {
	Put(k string, v any) ImmutableMap
	Get(k string) (v any, b bool)
}

type PersistentMap struct {
	data map[string]interface{}
}

func NewPersistentMap() *PersistentMap {
	return &PersistentMap{
		data: make(map[string]interface{}),
	}
}

func (m *PersistentMap) Put(key string, value any) ImmutableMap {
	newData := make(map[string]interface{})
	for k, v := range m.data {
		newData[k] = v
	}
	newData[key] = value
	return &PersistentMap{data: newData}
}

func (m *PersistentMap) Get(key string) (any, bool) {
	value, ok := m.data[key]
	return value, ok
}
