package types

// Language the type of the Language
type Language string

func (r Language) String() string {
	return string(r)
}

// FixedString the type of some FixedString
type FixedString string

func (r FixedString) String() string {
	return string(r)
}
