package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"github.com/spkaeros/rscgo/pkg/server/config"
	"github.com/spkaeros/rscgo/pkg/server/log"
	"io/ioutil"
	"os"
)

//RsaKey The RSA key for use in decoding the login packet
var RsaKey *rsa.PrivateKey

//LoadRsaKey Read the RSA key into memory.
func LoadRsaKey() {
	buf, err := ioutil.ReadFile(config.DataDir() + config.RsaKey())
	if err != nil {
		log.Error.Printf("Could not read RSA key from file:%v", err)
		os.Exit(103)
	}
	key, err := x509.ParsePKCS8PrivateKey(buf)
	if err != nil {
		log.Warning.Printf("Could not parse RSA key:%v", err)
		os.Exit(104)
	}
	RsaKey = key.(*rsa.PrivateKey)
}

//DecryptRSABlock Attempts to decrypt the payload buffer.  Returns the decrypted buffer upon success, otherwise returns nil.
func DecryptRSABlock(payload []byte) []byte {
	buf, err := rsa.DecryptPKCS1v15(rand.Reader, RsaKey, payload)
	if err != nil {
		log.Warning.Println("Could not decrypt RSA block:", err)
		return nil
	}
	return buf
}
