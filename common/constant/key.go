package constant

type ContKey string

func (r ContKey) Key() string {
	return string(r)
}
