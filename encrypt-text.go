package aes

import "encoding/base64"

/*
EncryptTextToBytes encrypts the text with the given key and returns the encrypted bytes.
*/
func EncryptTextToBytes(text, key string) []byte {
	src := []byte(text)
	return EncryptBytesToBytes(src, key)
}

/*
EncryptTextToText encrypts the text with the given key and returns the encrypted text.
*/
func EncryptTextToText(text, key string) string {
	return string(EncryptTextToBytes(text, key))
}

/*
EncryptTextToBase64 encrypts the text with the given key and returns the encrypted base64 string.
*/
func EncryptTextToBase64(text, key string) string {
	return base64.StdEncoding.EncodeToString(EncryptTextToBytes(text, key))
}
