package dictionary

func (d *Dictionary[T, K]) Clone() *Dictionary[T, K] {
	res := New(map[T]K{})
	for key, value := range d.data {
		res.Add(key, value)
	}
	return res
}
