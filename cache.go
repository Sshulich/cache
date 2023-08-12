package cache

type InMemoryCache interface {
	Set()
	Get() interface{}
	Delete()
}

type Cache struct {
	mapping map[string]interface{}
}

func New() *Cache {
	return &Cache{make(map[string]interface{})}
}

func (c *Cache) Set(k string, v interface{}) {
	c.mapping[k] = v
}

func (c *Cache) Get(k string) interface{} {
	return c.mapping[k]
}

func (c *Cache) Delete(k string) {
	delete(c.mapping, k)
}
