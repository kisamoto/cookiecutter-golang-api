package cache

// NewRedis returns a Redis powered implementation of Cache
func NewRedis() *Redis {
	return &Redis{}
}

// Redis implements the Cache interface using the Redis k/v
// store. It uses thread safe connection pooling.
type Redis struct {
}

func (r *Redis) Get(cacheable Cacheable) (Cacheable, bool) {
	key, err := cacheable.MarshalKey()
	if err != nil {
		return cacheable, false
	}
	value, found := r.GetRaw(key)
	if !found {
		LogMiss(key)
		return cacheable, false
	}
	LogHit(key)
	err = cacheable.UnmarshelValue(value)
	if err != nil {
		return cacheable, false
	}
	return cacheable, true
}

func (r *Redis) Put(marshaler Marshaler) error {
	key, err := marshaler.MarshalKey()
	if err != nil {
		return err
	}
	value, err := marshaler.MarshalValue()
	if err != nil {
		return err
	}
	return r.PutRaw(key, value)
}

func (r *Redis) Delete(marshaler Marshaler) error {
	key, err := marshaler.MarshalKey()
	if err != nil {
		return err
	}
	return r.DeleteRaw(key)
}

func (r *Redis) GetRaw(key []byte) (value []byte, found bool) {
	return []byte{}, true
}

func (r *Redis) PutRaw(key, value []byte) error {
	return nil
}

func (r *Redis) DeleteRaw(key []byte) error {
	return nil
}
