package util

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"errors"
)

const (
	NOPADDING = iota
	PKCS5PADDING
)

func PaddingPKCS5(text []byte, blockSize int) []byte {
	padding := blockSize - len(text)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(text, padText...)
}

func UnPaddingPKCS5(text []byte) []byte {
	length := len(text)
	padText := int(text[length-1])
	return text[:(length - padText)]
}

func Padding(text []byte, padding int) ([]byte, error) {
	switch padding {
	case NOPADDING:
		if len(text)%8 != 0 {
			return nil, errors.New("text length invalid")
		}
	case PKCS5PADDING:
		return PaddingPKCS5(text, 8), nil
	default:
		return nil, errors.New("padding type error")
	}
	return text, nil
}

func UnPadding(text []byte, padding int) ([]byte, error) {
	switch padding {
	case NOPADDING:
		if len(text)%8 != 0 {
			return nil, errors.New("text length invalid")
		}
	case PKCS5PADDING:
		return UnPaddingPKCS5(text), nil
	default:
		return nil, errors.New("padding type error")
	}
	return text, nil
}

// 加密
func DesEncryptCBC(plainText []byte, key []byte, iv []byte, padding int) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(iv) != block.BlockSize() {
		return nil, errors.New("iv length invalid")
	}
	text, err := Padding(plainText, padding)
	if err != nil {
		return nil, err
	}
	cipherText := make([]byte, len(text))
	encrypter := cipher.NewCBCEncrypter(block, iv)
	encrypter.CryptBlocks(cipherText, text)
	return cipherText, nil
}

// 解密
func DesDecryptCBC(cipherText []byte, key []byte, iv []byte, padding int) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(iv) != block.BlockSize() {
		return nil, errors.New("iv length invalid")
	}
	text := make([]byte, len(cipherText))
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypter.CryptBlocks(text, cipherText)
	plainText, err := UnPadding(text, padding)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}
