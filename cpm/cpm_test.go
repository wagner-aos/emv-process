package cpm

import (
	"log"

	"github.com/wagner-aos/emv-qrcode/emv/cpm"
)

func TestCPM() {
	emv := new(cpm.EMVQR)

	cdt := new(cpm.CommonDataTemplate)
	cdt.DataApplicationPAN = "1234567890123458"
	cdt.DataCardholderName = "WAGNER ALVES"
	cdt.DataLanguagePreference = "ruesdeen"

	cdtt := new(cpm.CommonDataTransparentTemplate)
	cdtt.DataIssuerApplicationData = "06010A03000000"
	cdtt.DataApplicationCryptogram = "584FD385FA234BCC"
	cdtt.DataApplicationTransactionCounter = "0001"
	cdtt.DataUnpredictableNumber = "6D58EF13"
	cdt.CommonDataTransparentTemplates = append(cdt.CommonDataTransparentTemplates, *cdtt)

	emv.CommonDataTemplates = append(emv.CommonDataTemplates, *cdt)

	comemvCode, err := emv.GeneratePayload()
	if err != nil {
		log.Println(err)
	}
	log.Println(comemvCode)
	// hQVDUFYwMWETTwegAAAAVVVVUAhQcm9kdWN0MWETTwegAAAAZmZmUAhQcm9kdWN0MmJJWggSNFZ4kBI0WF8gDkNBUkRIT0xERVIvRU1WXy0IcnVlc2RlZW5kIZ8QBwYBCgMAAACfJghYT9OF+iNLzJ82AgABnzcEbVjvEw==
}
