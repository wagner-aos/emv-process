package emv

import (
	"strconv"
	"strings"

	"github.com/kataras/golog"
)

//Parse -
func (t *Tags) Parse(payload string) {
	tags := map[string]tag{}
	offset := 0
	tagFound, tagExists := hasTagOnPayloadStart(payload, offset)

	if tagExists {
		//tag := tagMap[tagName]
		offset = len(tagFound.name) + tagFound.size
		golog.Infof("%s", tagFound.name)
		golog.Infof("%d", offset)
	}

	t.items = tags

}

func hasTagOnPayloadStart(payload string, offset int) (*tag, bool) {
	tagFound := &tag{}
	tagSize := 0
	for tagName := range tagMap {
		if strings.HasPrefix(payload[offset:], tagName) {
			tagSize, _ = strconv.Atoi(payload[len(tagName) : len(tagName)+2])
			tagValue := payload[len(tagName)+tagSize : len(tagName)+tagSize+tagSize]
			tagFound = &tag{
				name:  tagName,
				value: tagValue,
				size:  tagSize,
			}

			golog.Debugf("Tag: %s Value: %s- size: %d not found on tags available", tagFound.GetName(), tagFound.GetValue(), tagFound.GetSize())
			return tagFound, true
		}
	}
	golog.Error("Tag not found on tags available")
	return tagFound, false
}

// func extractTag(payload, stag string) *tag {

// }
