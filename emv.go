package emv

//EMV -
type EMV interface {
	GetEMV() Tags
	RemoveTag(name string)
	ToTLV() string
	Parse(payload string)
	//GeneratePayload() string
}

// Tags -
type Tags struct {
	items map[string]tag
}

//GetEMV -
func (t *Tags) GetEMV() Tags {
	return *t
}

//RemoveTag -
func (t *Tags) RemoveTag(name string) {
	delete(t.items, name)
}

//ToTLV -
func (t *Tags) ToTLV() string {
	s := ""
	if len(t.items) > 0 {
		for _, t := range t.items {
			s += formatting(t)
		}
	}
	return s
}
