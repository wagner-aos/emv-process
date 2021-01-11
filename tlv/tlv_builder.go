package tlv

//Builder -
type Builder interface {
	AddTag(name, value string) Builder
	Build() TLV
}

// NewTLV -
func NewTLV(field int) Builder {
	loadTags(field)
	return &Tags{
		items: map[string]tag{},
	}
}

//Build -
func (t *Tags) Build() TLV {
	return &Tags{
		items: t.items,
	}
}

//AddTag -
func (t *Tags) AddTag(name, value string) Builder {
	builder := New()
	data := builder.
		SetTag(name, value).
		BuildTag()

	t.items[data.GetName()] = tag{
		name:  data.GetName(),
		value: data.GetValue(),
		size:  len(data.GetValue()),
	}
	return t
}
