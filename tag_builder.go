package emv

//Data -
type Data interface {
	GetName() string
	GetValue() string
	GetSize() int
}

//Tag -
type tag struct {
	name        string
	description string
	value       string
	source      string
	format      string
	size        int
	minSize     int
	maxSize     int
}

//TagBuilder -
type TagBuilder interface {
	SetTag(string, string) TagBuilder
	BuildTag() Data
}

// New -
func New() TagBuilder {
	return &tag{}
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

//GetName - return name data from tag object
func (t *tag) GetName() string {
	return t.name
}

//GetValue - return value data from tag object
func (t *tag) GetValue() string {
	return t.value
}

//GetSize - return size of value data from tag object
func (t *tag) GetSize() int {
	return len(t.value)
}
