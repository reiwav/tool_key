package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/hyperboloide/lk"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}

func writeLicenseFile(licenseStr string) {
	// Write license file
	err_license := ioutil.WriteFile("license.key", []byte(licenseStr), 0644)
	check(err_license)

}

func (doc *License) GenerateLicense() error {
	// Unmarshal the private key:
	privateKey, err := lk.PrivateKeyFromB64String(privateKeyStr)
	if err != nil {
		return err
	}

	// marshall the document to json bytes:
	docBytes, err := json.Marshal(&doc)
	if err != nil {
		return err

	}

	// generate your license with the private key and the document:
	license, err := lk.NewLicense(privateKey, docBytes)
	if err != nil {
		return err
	}

	// encode the new license to b64, this is what you give to your customer.
	str64, err := license.ToB64String()
	if err != nil {
		return err
	}
	fmt.Println(str64)

	writeLicenseFile(str64)
	return nil
}
