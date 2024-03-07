package aes

/*
DecryptTextToBytes decrypts the encrypted bytes with the given key and returns the decrypted bytes.
*/
func DecryptTextToBytes(encrypted string, key string) []byte {
	src := []byte(encrypted)
	return DecryptBytesToBytes(src, key)
}

/*
DecryptTextToText decrypts the encrypted bytes with the given key and returns the decrypted text.
*/
func DecryptTextToText(encrypted string, key string) string {
	return string(DecryptTextToBytes(encrypted, key))
}
