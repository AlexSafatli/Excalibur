package model

type Skill struct {
	Name   string
	Fields []*Field
}

func (s *Skill) HasField(f *Field) bool {
	for _, field := range s.Fields {
		if field.Name == f.Name {
			return true
		}
	}
	return false
}
