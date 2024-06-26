package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

func encrypt(key, plaintext []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(plaintext) % block.BlockSize() != 0 {
		l := len(plaintext) + (block.BlockSize() - (len(plaintext) % block.BlockSize()))
		plaintext = padMsg(plaintext, l)
	}
	cipherText := make([]byte, len(plaintext) + block.BlockSize())
	iv := cipherText[:block.BlockSize()]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[des.BlockSize:], plaintext)
	return cipherText, nil
}

func padMsg(plaintext []byte, blockSize int) []byte {
	result := make([]byte, blockSize)
	copy(result, plaintext)
	return result
}

func decrypt(key, ciphertext []byte) (plaintext []byte, err error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(ciphertext) < des.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := ciphertext[:des.BlockSize]
	ciphertext = ciphertext[des.BlockSize:]
	if len(ciphertext)%des.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	return ciphertext, nil
}

func padWithZeros(block []byte, desiredSize int) []byte {
	for len(block) < desiredSize {
		block = append(block, 0)
	}
	return block
}

func test(key, plaintext []byte) {
	ciphertext, err := encrypt(key, plaintext)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Encrypting '%v' with key '%v'...\n", string(plaintext), string(key))
	decryptedText, err := decrypt(key, ciphertext)
	if err != nil {
		fmt.Println(err)
		return
	}
	decryptedText = bytes.Trim(decryptedText, "\x00")
	fmt.Printf("Decrypted: '%v'\n", string(decryptedText))
	fmt.Println("========")
}

func main() {
	test(
		[]byte("12344321"),
		[]byte("Today I met my crush, what a hunk"),
	)

	test(
		[]byte("p@$$w0rd"),
		[]byte("I hope my boyfriend never finds out about this"),
	)
}
