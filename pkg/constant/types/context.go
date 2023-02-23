package types

type ContextKey string

func (r ContextKey) Key() string {
	return string(r)
}
