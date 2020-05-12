package emv

//TagBuilder -
type TagBuilder interface {
	SetTag(string, string) TagBuilder
	BuildTag() Data
}

//SetName -
func (t *tag) SetTag(name, value string) TagBuilder {
	t.name = name
	t.value = value
	return t
}

//Build -
func (t *tag) BuildTag() Data {
	return &tag{
		name:  t.name,
		value: t.value,
	}
}
