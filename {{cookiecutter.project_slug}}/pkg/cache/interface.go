package cache

// A Cache Gets, Puts and Deletes []byte values and is used to wrap
// slow database queries.
type Cache interface {
	Get(Cacheable) (cacheable Cacheable, found bool)
	Put(Marshaler) error
	Delete(Marshaler) error
	GetRaw(key []byte) (value []byte, found bool)
	PutRaw(key, value []byte) error
	DeleteRaw(key []byte) error
}

// Marshaler implementations are used in the cache to turn
// contents into []byte
type Marshaler interface {
	MarshalKey() ([]byte, error)
	MarshalValue() ([]byte, error)
}

// Unmarsheler implementations assemble themselves from the
// key and value []byte provided by the cache
type Unmarsheler interface {
	UnmarshelKey([]byte) error
	UnmarshelValue([]byte) error
}

// A Cacheable implementation is both Marshalable and Unmarshalable
type Cacheable interface {
	Marshaler
	Unmarsheler
}
