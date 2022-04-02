package util

func MapKeys[K comparable, V any](d map[K]V) []K {
	var res []K
	for k := range d {
		res = append(res, k)
	}
	return res
}

func MapValues[K comparable, V any](d map[K]V) []V {
	var res []V
	for _, v := range d {
		res = append(res, v)
	}
	return res
}

func MapToList[K comparable, V any](d map[K]V) []Tuple2[K, V] {
	var res []Tuple2[K, V]
	for k, v := range d {
		res = append(res, NewTuple2(k, v))
	}
	return res
}

func MapAdd[K comparable, V any](d map[K]V, k K, v V) {
	d[k] = v
}

func MapDel[K comparable, V any](d map[K]V, k K) {
	if MapHas(d, k) {
		delete(d, k)
	}
}

func MapHas[K comparable, V any](d map[K]V, k K) bool {
	_, ok := d[k]
	return ok
}

func MapGet[K comparable, V any](d map[K]V, k K) V {
	v, ok := d[k]
	if !ok {
		panic("Dictionary not found key")
	}
	return v
}

func MapFind[K comparable, V any](d map[K]V, k K) Optional[V] {
	v, ok := d[k]
	if ok {
		return Some(v)
	} else {
		return None[V]()
	}
}
