package emv

import (
	"fmt"
	"strconv"
	"strings"
)

//Parse -
func (t *Tags) Parse(payload string) {
	offset := 0
	for {
		tagFound, tagExists := hasTagOnPayload(payload, offset)

		if tagExists {
			offset += len(tagFound.name) + 2 + tagFound.size*2
			fmt.Printf("\n\tTag Name: %s", tagFound.name)
			fmt.Printf("\n\tTag Value: %s", tagFound.value)
			fmt.Printf("\n\tTag Size: %d", tagFound.size)

			fmt.Printf("\n\n\tOffset: %d", offset)

			t.items[tagFound.name] = *tagFound
		} else {
			return
		}
	}
}

func hasTagOnPayload(payload string, offset int) (*tag, bool) {

	for tagName := range tagMap {
		payloadPosition := payload[offset:]

		if strings.HasPrefix(payloadPosition, tagName) {
			tagSize, _ := strconv.ParseInt(payloadPosition[len(tagName):len(tagName)+2], 16, 64)
			//tagHexSize, _ := strconv.Atoi(payloadPosition[len(tagName) : len(tagName)+2])
			//tagSize := str
			tagValue := payloadPosition[len(tagName)+2 : len(tagName)+2+int(tagSize)*2]
			tagFound := &tag{
				name:  tagName,
				value: tagValue,
				size:  int(tagSize),
			}

			return tagFound, true
		}
	}
	return &tag{}, false
}
