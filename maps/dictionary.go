package maps

type Dictionary map[string]string

func (d Dictionary) Search(key string) string {
	return d[key]
}
