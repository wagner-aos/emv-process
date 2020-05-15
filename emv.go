package emv

import "fmt"

//EMV -
type EMV interface {
	GetEMV() Tags
	GetTag(name string) (string, string, int, error)
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

//GetTag - it returs tag name, value, size and error
func (t *Tags) GetTag(name string) (string, string, int, error) {
	tag := t.items[name]
	if tag.name == "" {
		err := fmt.Errorf("[emv] error - tag: %s not found", name)
		return "", "", 0, err
	}
	return tag.name, tag.value, tag.size, nil
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
