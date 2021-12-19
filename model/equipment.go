package model

type Item struct {
	Name   string
	Fields []*Field
}

func (i *Item) HasField(f *Field) bool {
	for _, field := range i.Fields {
		if field.Name == f.Name {
			return true
		}
	}
	return false
}
