package tlv

import "fmt"

// TLV -
type TLV interface {
	GetTLV() Tags
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

//GetTLV -
func (t *Tags) GetTLV() Tags {
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

func formatting(t tag) string {
	return fmt.Sprintf("*** TAG %s NOT FOUND ***", t.GetName())
}

// GetTags return list of tags emv
func (t *Tags) GetTags() map[string]tag {
	return t.items
}
