package emv

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"unicode/utf8"
)

//Format TLV - return TLV of a tag
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

func formatting(t tag) string {

	if t.name == TagApplicationDefinitionFileName {
		return format(TagApplicationDefinitionFileName, t.value)
	}
	if t.name == TagApplicationLabel {
		return format(TagApplicationLabel, t.value)
	}
	if t.name == TagTrack2EquivalentData {
		return format(TagTrack2EquivalentData, t.value)
	}
	if t.name == TagApplicationPAN {
		return format(TagApplicationPAN, t.value)
	}
	if t.name == TagApplicationInterchangeProfile {
		return format(TagApplicationInterchangeProfile, t.value)
	}
	if t.name == TagDedicatedFileName {
		return format(TagDedicatedFileName, t.value)
	}
	if t.name == TagTerminalVerificationResults {
		return format(TagTerminalVerificationResults, t.value)
	}
	if t.name == TagTransactionDate {
		return format(TagTransactionDate, t.value)
	}
	if t.name == TagTransactionStatusInformation {
		return format(TagTransactionStatusInformation, t.value)
	}
	if t.name == TagTransactionType {
		return format(TagTransactionType, t.value)
	}
	if t.name == TagCardholderName {
		return format(TagCardholderName, t.value)
	}
	if t.name == TagApplicationExpirationDate {
		return format(TagApplicationExpirationDate, t.value)
	}
	if t.name == TagIssuerCountryCode {
		return format(TagIssuerCountryCode, t.value)
	}
	if t.name == TagTransactionCurrencyCode {
		return format(TagTransactionCurrencyCode, t.value)
	}
	if t.name == TagLanguagePreference {
		return format(TagLanguagePreference, t.value)
	}
	if t.name == TagPANSequenceNumber {
		return format(TagPANSequenceNumber, t.value)
	}
	if t.name == TagIssuerURL {
		return format(TagIssuerURL, t.value)
	}
	if t.name == TagAuthorisedAmount {
		return format(TagAuthorisedAmount, t.value)
	}
	if t.name == TagAmountOther {
		return format(TagAmountOther, t.value)
	}
	if t.name == TagApplicationVersionNumber {
		return format(TagApplicationVersionNumber, t.value)
	}
	if t.name == TagIssuerApplicationData {
		return format(TagIssuerApplicationData, t.value)
	}
	if t.name == TagTokenRequestorID {
		return format(TagTokenRequestorID, t.value)
	}
	if t.name == TagTerminalCountryCode {
		return format(TagTerminalCountryCode, t.value)
	}
	if t.name == TagPaymentAccountReference {
		return format(TagPaymentAccountReference, t.value)
	}
	if t.name == TagLast4DigitsOfPAN {
		return format(TagLast4DigitsOfPAN, t.value)
	}
	if t.name == TagApplicationCryptogram {
		return format(TagApplicationCryptogram, t.value)
	}
	if t.name == TagCryptogramInformationData {
		return format(TagCryptogramInformationData, t.value)
	}
	if t.name == TagTerminalCapabilities {
		return format(TagTerminalCapabilities, t.value)
	}
	if t.name == TagCardholderVerificationMethod {
		return format(TagCardholderVerificationMethod, t.value)
	}
	if t.name == TagTerminalType {
		return format(TagTerminalType, t.value)
	}
	if t.name == TagApplicationTransactionCounter {
		return format(TagApplicationTransactionCounter, t.value)
	}
	if t.name == TagUnpredictableNumber {
		return format(TagUnpredictableNumber, t.value)
	}

	return string("tag not found")
}
