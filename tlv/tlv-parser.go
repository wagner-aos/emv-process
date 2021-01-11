package tlv

import (
	"fmt"
	"strings"
)

//Parse -
func (t *Tags) Parse(payload string) {
	offset := 0

	for _, key := range tagMap.Keys() {
		payloadPosition := payload[offset:]
		value, _ := tagMap.Get(key)
		tg := value.(tag)
		if strings.HasPrefix(payloadPosition, tg.name) {

			tagSize := tg.size
			tagValue := payloadPosition[len(tg.name) + tg.def : len(tg.name) + tg.def + tagSize]
			fmt.Print(tagValue)
			tagFound := &tag{
				name:  tg.name,
				value: tagValue,
				size:  tagSize,
			}

			t.items[tagFound.name] = *tagFound
			offset += len(tg.name) + tg.def + tagSize
		}
	}
}
