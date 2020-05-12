package emv

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"unicode/utf8"
)

// const ...
const (
	TagApplicationDefinitionFileName = "4F"
	TagApplicationLabel              = "50"
	TagTrack2EquivalentData          = "57"
	TagApplicationPAN                = "5A"
	TagCardholderName                = "5F20"
	TagLanguagePreference            = "5F2D"
	TagIssuerURL                     = "5F50"
	TagApplicationVersionNumber      = "9F08"
	TagIssuerApplicationData         = "9F10"
	TagTokenRequestorID              = "9F19"
	TagPaymentAccountReference       = "9F24"
	TagLast4DigitsOfPAN              = "9F25"
	TagApplicationCryptogram         = "9F26"
	TagApplicationTransactionCounter = "9F36"
	TagUnpredictableNumber           = "9F37"
)

//EMV -
type EMV interface {
	GetEMV() tags
	RemoveTag(name string)
	//GeneratePayload() string
}

//Builder -
type Builder interface {
	AddTag(name, value string) Builder
	Build() EMV
}

// tags -
type tags struct {
	items map[string]tag
}

// NewEMV -
func NewEMV() Builder {
	return &tags{
		items: map[string]tag{},
	}
}

//AddTag -
func (t *tags) AddTag(name, value string) Builder {
	builder := New()
	data := builder.
		SetTag(name, value).
		BuildTag()

	t.items[data.GetName()] = tag{
		value: data.GetValue(),
	}
	return t
}

//AddTag -
func (t *tags) RemoveTag(name string) {
	delete(t.items, name)
}

//GetEMV
func (t *tags) GetEMV() tags {
	return *t
}

//Build -
func (t *tags) Build() EMV {
	return &tags{
		items: t.items,
	}
}

//TAG

//Data -
type Data interface {
	GetName() string
	GetValue() string
}

//Tag -
type tag struct {
	name        string
	description string
	value       string
	source      string
	format      string
	minSize     int
	maxSize     int
}

// New -
func New() TagBuilder {
	return &tag{}
}

//SetValue -
func (t *tag) AddTag() {

}

//GetName - return name data from tag object
func (t *tag) GetName() string {
	return t.name
}

//GetValue - return value data from tag object
func (t *tag) GetValue() string {
	return t.value
}

// // BERTLV ...
// type BERTLV struct {
// 	DataApplicationDefinitionFileName Tag // "4F"
// 	DataApplicationLabel              Tag // "50"
// 	DataTrack2EquivalentData          Tag // "57"
// 	DataApplicationPAN                Tag // "5A"
// 	DataCardholderName                Tag // "5F20"
// 	DataLanguagePreference            Tag // "5F2D"
// 	DataIssuerURL                     Tag // "5F50"
// 	DataApplicationVersionNumber      Tag // "9F08"
// 	DataIssuerApplicationData         Tag // "9F10"
// 	DataTokenRequestorID              Tag // "9F19"
// 	DataPaymentAccountReference       Tag // "9F24"
// 	DataLast4DigitsOfPAN              Tag // "9F25"
// 	DataApplicationCryptogram         Tag // "9F26"
// 	DataApplicationTransactionCounter Tag // "9F36"
// 	DataUnpredictableNumber           Tag // "9F37"
// }

// GeneratePayload ...
//func (e *TAGS) GeneratePayload() (string, error) {
//s := ""

// if len(c.CommonDataTemplates) > 0 {
// 	for _, t := range c.CommonDataTemplates {
// 		template := formattingTemplate(t.BERTLV)
// 		s += template
// 	}
// }

// decoded, err := hex.DecodeString(s)
// if err != nil {
// 	return "", err
// }
// s = base64.StdEncoding.EncodeToString([]byte(string(decoded)))
//return s, nil
//return string(decoded), nil
//}

//Format TLV
func format(id, value string) string {
	length := utf8.RuneCountInString(value) / 2
	lengthStr := strconv.Itoa(length)
	lengthStr = "00" + fmt.Sprintf("%X", length)
	return id + lengthStr[len(lengthStr)-2:] + value
}

func toHex(s string) string {
	src := []byte(s)
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return string(dst)
}

// func formattingTemplate(t BERTLV) string {
// 	template := ""
// 	if t.DataApplicationDefinitionFileName.Name != "" {
// 		template += format(TagApplicationDefinitionFileName, t.DataApplicationDefinitionFileName)
// 	}
// 	if t.DataApplicationLabel != "" {
// 		template += format(TagApplicationLabel, toHex(t.DataApplicationLabel))
// 	}
// 	if t.DataTrack2EquivalentData != "" {
// 		template += format(TagTrack2EquivalentData, t.DataTrack2EquivalentData)
// 	}
// 	if t.DataApplicationPAN != "" {
// 		template += format(TagApplicationPAN, t.DataApplicationPAN)
// 	}
// 	if t.DataCardholderName != "" {
// 		template += format(TagCardholderName, toHex(t.DataCardholderName))
// 	}
// 	if t.DataLanguagePreference != "" {
// 		template += format(TagLanguagePreference, toHex(t.DataLanguagePreference))
// 	}
// 	if t.DataIssuerURL != "" {
// 		template += format(TagIssuerURL, toHex(t.DataIssuerURL))
// 	}
// 	if t.DataApplicationVersionNumber != "" {
// 		template += format(TagApplicationVersionNumber, t.DataApplicationVersionNumber)
// 	}
// 	if t.DataIssuerApplicationData != "" {
// 		template += format(TagIssuerApplicationData, t.DataIssuerApplicationData)
// 	}
// 	if t.DataTokenRequestorID != "" {
// 		template += format(TagTokenRequestorID, t.DataTokenRequestorID)
// 	}
// 	if t.DataPaymentAccountReference != "" {
// 		template += format(TagPaymentAccountReference, t.DataPaymentAccountReference)
// 	}
// 	if t.DataLast4DigitsOfPAN != "" {
// 		template += format(TagLast4DigitsOfPAN, t.DataLast4DigitsOfPAN)
// 	}
// 	if t.DataApplicationCryptogram != "" {
// 		template += format(TagApplicationCryptogram, t.DataApplicationCryptogram)
// 	}
// 	if t.DataApplicationTransactionCounter != "" {
// 		template += format(TagApplicationTransactionCounter, t.DataApplicationTransactionCounter)
// 	}
// 	if t.DataUnpredictableNumber != "" {
// 		template += format(TagUnpredictableNumber, t.DataUnpredictableNumber)
// 	}
// 	return template
// }
