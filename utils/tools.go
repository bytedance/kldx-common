package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

func paddingN(text []byte, size int) []byte {
	if len(text) > size {
		return text[:size]
	}

	return append(text, bytes.Repeat([]byte("0"), size-len(text))...)
}

func AesDecryptText(fieldID int64, realEncryptKey []byte, encryptedText string) (originText string, err error) {
	if len(realEncryptKey) != 32 {
		return "", fmt.Errorf("ilegal length")
	}
	iv, err := getInitialVector(fmt.Sprintf("%d", fieldID))
	if err != nil {
		return "", err
	}
	bs, err := hex2bin(encryptedText)

	if err != nil {

		return "", err
	}
	originBytes, err := aesCbsDecrypt(bs, iv, realEncryptKey)
	if err != nil {
		return "", err
	}

	return string(originBytes), nil
}

func getInitialVector(str string) ([]byte, error) {
	h := md5.New()
	if _, err := io.WriteString(h, str); err != nil {
		return nil, err
	}
	return h.Sum(nil)[:], nil
}

func hex2bin(bs string) ([]byte, error) {
	if len(bs)%2 != 0 {
		return nil, hex.ErrLength
	}
	src := []byte(bs)
	dst := make([]byte, hex.DecodedLen(len(bs)))

	_, err := hex.Decode(dst, src)
	if err != nil {
		return nil, err

	}
	return dst, nil
}

func aesCbsDecrypt(cipherText, iv []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	plainText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(plainText, cipherText)
	plainText = unPaddingN(plainText)
	return plainText, nil
}

func unPaddingN(cipherText []byte) []byte {
	end := cipherText[len(cipherText)-1]
	cipherText = cipherText[:len(cipherText)-int(end)]
	return cipherText
}