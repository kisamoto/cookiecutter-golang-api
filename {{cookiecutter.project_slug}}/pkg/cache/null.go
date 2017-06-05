package cache

/*
 * NewNull returns a Cache implementation that does nothing.
 * It serves as a way to pass in a cache implementation rather
 * than using `nil` and can be easily substituted in the future
 * with a storage backed Cache implementation.
 */
func NewNull() *nullCache {
	return &nullCache{}
}

type nullCache struct {
}

// Get returns `nil` and not found
func (n *nullCache) Get(Cacheable) (cacheable Cacheable, found bool) {
	return nil, false
}

// Put returns nil
func (n *nullCache) Put(Marshaler) error {
	return nil
}

// Delete returns nil
func (n *nullCache) Delete(Marshaler) error {
	return nil
}

// GetRaw returns an empty byte array and not found
func (n *nullCache) GetRaw(key []byte) (value []byte, found bool) {
	return []byte{}, false
}

// PutRaw returns nil
func (n *nullCache) PutRaw(key, value []byte) error {
	return nil
}

// DeleteRaw returns nil
func (n *nullCache) DeleteRaw(key []byte) error {
	return nil
}
