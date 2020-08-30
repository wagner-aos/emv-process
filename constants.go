package emv

import (
	"fmt"
)

// const emv tag names
const (
	TagApplicationDefinitionFileName    = "4F"
	TagApplicationLabel                 = "50"
	TagTrack2EquivalentData             = "57"
	TagApplicationPAN                   = "5A"
	TagIssuerScriptData1                = "71"
	TagIssuerScriptData2                = "72"
	TagApplicationInterchangeProfile    = "82"
	TagDedicatedFileName                = "84"
	TagIssuerAuthenticationData         = "91"
	TagTerminalVerificationResults      = "95"
	TagTransactionDate                  = "9A"
	TagTransactionStatusInformation     = "9B"
	TagTransactionType                  = "9C"
	TagCardholderName                   = "5F20"
	TagApplicationExpirationDate        = "5F24"
	TagIssuerCountryCode                = "5F28"
	TagTransactionCurrencyCode          = "5F2A"
	TagLanguagePreference               = "5F2D"
	TagPANSequenceNumber                = "5F34"
	TagIssuerURL                        = "5F50"
	TagAuthorisedAmount                 = "9F02"
	TagAmountOther                      = "9F03"
	TagApplicationIdentifier            = "9F06"
	TagApplicationUsageControl          = "9F07"
	TagApplicationVersionNumber         = "9F08"
	TagTerminalApplicationVersionNumber = "9F09"
	TagIssuerApplicationData            = "9F10"
	TagMerchantCategoryCode             = "9F15"
	TagTokenRequestorID                 = "9F19"
	TagTerminalCountryCode              = "9F1A"
	TagInterfaceDeviceSerialNumber      = "9F1E"
	TagPaymentAccountReference          = "9F24"
	TagLast4DigitsOfPAN                 = "9F25"
	TagApplicationCryptogram            = "9F26"
	TagCryptogramInformationData        = "9F27"
	TagTerminalCapabilities             = "9F33"
	TagCardholderVerificationMethod     = "9F34"
	TagTerminalType                     = "9F35"
	TagApplicationTransactionCounter    = "9F36"
	TagUnpredictableNumber              = "9F37"
	TagTransactionSequenceCounter       = "9F41"
	TagIssuerScriptResults              = "9F5B"
)

var tagMap = map[string]tag{}

func init() {
	loadTags()
	validateTagNameMap()
}

//load Emv tags available
func loadTags() {
	tagMap[TagApplicationDefinitionFileName] = tag{name: TagApplicationDefinitionFileName, minSize: 0, maxSize: 0}
	tagMap[TagApplicationLabel] = tag{name: TagApplicationLabel, minSize: 0, maxSize: 0}
	tagMap[TagTrack2EquivalentData] = tag{name: TagTrack2EquivalentData, minSize: 0, maxSize: 0}
	tagMap[TagApplicationPAN] = tag{name: TagApplicationPAN, minSize: 0, maxSize: 0}
	tagMap[TagIssuerScriptData1] = tag{name: TagIssuerScriptData1, minSize: 0, maxSize: 0}
	tagMap[TagIssuerScriptData2] = tag{name: TagIssuerScriptData2, minSize: 0, maxSize: 0}
	tagMap[TagApplicationInterchangeProfile] = tag{name: TagApplicationInterchangeProfile, minSize: 0, maxSize: 0}
	tagMap[TagDedicatedFileName] = tag{name: TagDedicatedFileName, minSize: 0, maxSize: 0}
	tagMap[TagIssuerAuthenticationData] = tag{name: TagIssuerAuthenticationData, minSize: 0, maxSize: 0}
	tagMap[TagTerminalVerificationResults] = tag{name: TagTerminalVerificationResults, minSize: 0, maxSize: 0}
	tagMap[TagTransactionDate] = tag{name: TagTransactionDate, minSize: 0, maxSize: 0}
	tagMap[TagTransactionStatusInformation] = tag{name: TagTransactionStatusInformation, minSize: 0, maxSize: 0}
	tagMap[TagTransactionType] = tag{name: TagTransactionType, minSize: 0, maxSize: 0}
	tagMap[TagCardholderName] = tag{name: TagCardholderName, minSize: 0, maxSize: 0}
	tagMap[TagApplicationExpirationDate] = tag{name: TagApplicationExpirationDate, minSize: 0, maxSize: 0}
	tagMap[TagIssuerCountryCode] = tag{name: TagIssuerCountryCode, minSize: 0, maxSize: 0}
	tagMap[TagTransactionCurrencyCode] = tag{name: TagTransactionCurrencyCode, minSize: 0, maxSize: 0}
	tagMap[TagLanguagePreference] = tag{name: TagLanguagePreference, minSize: 0, maxSize: 0}
	tagMap[TagPANSequenceNumber] = tag{name: TagPANSequenceNumber, minSize: 0, maxSize: 0}
	tagMap[TagIssuerURL] = tag{name: TagIssuerURL, minSize: 0, maxSize: 0}
	tagMap[TagAuthorisedAmount] = tag{name: TagAuthorisedAmount, minSize: 0, maxSize: 0}
	tagMap[TagAmountOther] = tag{name: TagAmountOther, minSize: 0, maxSize: 0}
	tagMap[TagApplicationIdentifier] = tag{name: TagApplicationIdentifier, minSize: 0, maxSize: 0}
	tagMap[TagApplicationUsageControl] = tag{name: TagApplicationUsageControl, minSize: 0, maxSize: 0}
	tagMap[TagApplicationVersionNumber] = tag{name: TagApplicationVersionNumber, minSize: 0, maxSize: 0}
	tagMap[TagTerminalApplicationVersionNumber] = tag{name: TagTerminalApplicationVersionNumber, minSize: 0, maxSize: 0}
	tagMap[TagIssuerApplicationData] = tag{name: TagIssuerApplicationData, minSize: 0, maxSize: 0}
	tagMap[TagMerchantCategoryCode] = tag{name: TagMerchantCategoryCode, minSize: 0, maxSize: 0}
	tagMap[TagTokenRequestorID] = tag{name: TagTokenRequestorID, minSize: 0, maxSize: 0}
	tagMap[TagTerminalCountryCode] = tag{name: TagTerminalCountryCode, minSize: 0, maxSize: 0}
	tagMap[TagInterfaceDeviceSerialNumber] = tag{name: TagInterfaceDeviceSerialNumber, minSize: 0, maxSize: 0}
	tagMap[TagPaymentAccountReference] = tag{name: TagPaymentAccountReference, minSize: 0, maxSize: 0}
	tagMap[TagLast4DigitsOfPAN] = tag{name: TagLast4DigitsOfPAN, minSize: 0, maxSize: 0}
	tagMap[TagApplicationCryptogram] = tag{name: TagApplicationCryptogram, minSize: 0, maxSize: 0}
	tagMap[TagCryptogramInformationData] = tag{name: TagCryptogramInformationData, minSize: 0, maxSize: 0}
	tagMap[TagTerminalCapabilities] = tag{name: TagTerminalCapabilities, minSize: 0, maxSize: 0}
	tagMap[TagCardholderVerificationMethod] = tag{name: TagCardholderVerificationMethod, minSize: 0, maxSize: 0}
	tagMap[TagTerminalType] = tag{name: TagTerminalType, minSize: 0, maxSize: 0}
	tagMap[TagApplicationTransactionCounter] = tag{name: TagApplicationTransactionCounter, minSize: 0, maxSize: 0}
	tagMap[TagUnpredictableNumber] = tag{name: TagUnpredictableNumber, minSize: 0, maxSize: 0}
	tagMap[TagTransactionSequenceCounter] = tag{name: TagTransactionSequenceCounter, minSize: 0, maxSize: 0}
	tagMap[TagIssuerScriptResults] = tag{name: TagIssuerScriptResults, minSize: 0, maxSize: 0}
}

func validateTagNameMap() bool {
	for k, v := range tagMap {
		//fmt.Printf("\n\tTagName: %s", v.GetName())
		if k != v.GetName() {
			err := fmt.Errorf("Tag Map key: %s is not equal a tag name: %s", k, v.GetName())
			fmt.Printf("[emv]:%s", err)
			return false
		}
	}
	return true
}
