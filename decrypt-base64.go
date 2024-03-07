package aes

import "encoding/base64"

/*
DecryptBase64ToBytes decrypts the encrypted base64 string with the given key and returns the decrypted bytes.
*/
func DecryptBase64ToBytes(encrypted string, key string) []byte {
	decrypted, _ := base64.StdEncoding.DecodeString(encrypted)
	return DecryptBytesToBytes(decrypted, key)
}

/*
DecryptBase64ToText decrypts the encrypted base64 string with the given key and returns the decrypted text.
*/
func DecryptBase64ToText(encrypted string, key string) string {
	return string(DecryptBase64ToBytes(encrypted, key))
}
