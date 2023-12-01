package entity

type Category struct {
	ID   uint
	Name string
}

func (c *Category) IsEmpty() bool {
	return c.Name == ""
}
