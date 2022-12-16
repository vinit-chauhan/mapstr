package mapstr

import (
	"strings"
)

type M map[string]interface{}

/*
 *
 * Get (key) : (value interface{}, err error)
 * Update (key, value) : (ok bool, err error)
 * Delete (key) : (err error)
 * HasKey (key) : (exist bool, err error)
 * Validate (func) : (valid bool, err error)
 *
 */

func New() M {
	return make(M)
}

func (m M) Set(d interface{}) bool {
	if msi, ok := d.(M); ok {
		m.Put(msi)
		return true
	}
	return false
}

func (m M) Get(key string) (value interface{}, present bool, err error) {
	value, present, err = m.find(key, m)
	return value, present, err
}

func (m M) Put(d M) {
	for k, v := range d {
		m[k] = v
	}
}

func (m M) Update(key string, value interface{}) (ok bool, err error) {
	return ok, nil
}

func (m M) Delete(key string) (err error) {
	return nil
}

func (m M) HasKey(key string) (exist bool, err error) {
	return exist, nil
}

func (m M) Validate(f func()) (valid bool, err error) {
	return valid, nil
}

func (m M) find(key string, data M) (value interface{}, present bool, err error) {
	for {
		if v, exists := data[key]; exists {
			return v, true, nil
		}

		tmpID := strings.IndexRune(key, '.')
		if tmpID < 0 {
			return data, false, nil
		}

		currKey := key[:tmpID]

		v, exist := data[currKey]
		if !exist {
			return nil, false, ErrKeyNotFound
		}

		ok, d := CastToMapstr(v)
		if !ok {
			return nil, false, ErrUnableToCastToMap
		}

		key = key[tmpID+1:]
		data = d
	}
}

func CastToMapstr(data interface{}) (ok bool, dataMap M) {
	if msi, ok := data.(map[string]interface{}); ok {
		return true, msi
	}
	return false, nil
}
