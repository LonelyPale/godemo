package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/minio/sio"
	"golang.org/x/crypto/hkdf"
	"golang.org/x/sys/cpu"
)

func main() {
	var supportsAES = (cpu.X86.HasAES && cpu.X86.HasPCLMULQDQ) || runtime.GOARCH == "s390x"
	fmt.Println("supportsAES:", supportsAES)

	ExampleEncryptReader()
}

func ExampleEncrypt() {
	// the master key used to derive encryption keys
	// this key must be keep secret
	masterkey, err := hex.DecodeString("000102030405060708090A0B0C0D0E0FF0E0D0C0B0A090807060504030201000") // use your own key here
	if err != nil {
		fmt.Printf("Cannot decode hex key: %v", err) // add error handling
		return
	}

	// generate a random nonce to derive an encryption key from the master key
	// this nonce must be saved to be able to decrypt the data again - it is not
	// required to keep it secret
	var nonce [32]byte
	if _, err = io.ReadFull(rand.Reader, nonce[:]); err != nil {
		fmt.Printf("Failed to read random data: %v", err) // add error handling
		return
	}

	// derive an encryption key from the master key and the nonce
	var key [32]byte
	kdf := hkdf.New(sha256.New, masterkey, nonce[:], nil)
	if _, err = io.ReadFull(kdf, key[:]); err != nil {
		fmt.Printf("Failed to derive encryption key: %v", err) // add error handling
		return
	}

	input := os.Stdin   // customize for your needs - the plaintext
	output := os.Stdout // customize from your needs - the decrypted output

	if _, err = sio.Encrypt(output, input, sio.Config{Key: key[:]}); err != nil {
		fmt.Printf("Failed to encrypt data: %v", err) // add error handling
		return
	}
}

func ExampleDecrypt() {
	// the master key used to derive encryption keys
	masterkey, err := hex.DecodeString("000102030405060708090A0B0C0D0E0FF0E0D0C0B0A090807060504030201000") // use your own key here
	if err != nil {
		fmt.Printf("Cannot decode hex key: %v", err) // add error handling
		return
	}

	// the nonce used to derive the encryption key
	nonce, err := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000001") // use your generated nonce here
	if err != nil {
		fmt.Printf("Cannot decode hex key: %v", err) // add error handling
		return
	}

	// derive the encryption key from the master key and the nonce
	var key [32]byte
	kdf := hkdf.New(sha256.New, masterkey, nonce, nil)
	if _, err = io.ReadFull(kdf, key[:]); err != nil {
		fmt.Printf("Failed to derive encryption key: %v", err) // add error handling
		return
	}

	input := os.Stdin   // customize for your needs - the encrypted data
	output := os.Stdout // customize from your needs - the decrypted output

	if _, err = sio.Decrypt(output, input, sio.Config{Key: key[:]}); err != nil {
		if _, ok := err.(sio.Error); ok {
			fmt.Printf("Malformed encrypted data: %v", err) // add error handling - here we know that the data is malformed/not authentic.
			return
		}
		fmt.Printf("Failed to decrypt data: %v", err) // add error handling
		return
	}
}

func ExampleEncryptReader() {
	// the master key used to derive encryption keys
	// this key must be keep secret
	masterkey, err := hex.DecodeString("000102030405060708090A0B0C0D0E0FF0E0D0C0B0A090807060504030201000") // use your own key here
	if err != nil {
		fmt.Printf("Cannot decode hex key: %v", err) // add error handling
		return
	}

	// generate a random nonce to derive an encryption key from the master key
	// this nonce must be saved to be able to decrypt the data again - it is not
	// required to keep it secret
	var nonce [32]byte
	if _, err = io.ReadFull(rand.Reader, nonce[:]); err != nil {
		fmt.Printf("Failed to read random data: %v", err) // add error handling
		return
	}

	// derive an encryption key from the master key and the nonce
	var key [32]byte
	kdf := hkdf.New(sha256.New, masterkey, nonce[:], nil)
	if _, err = io.ReadFull(kdf, key[:]); err != nil {
		fmt.Printf("Failed to derive encryption key: %v", err) // add error handling
		return
	}

	input := os.Stdin // customize for your needs - the plaintext input
	encrypted, err := sio.EncryptReader(input, sio.Config{Key: key[:]})
	if err != nil {
		fmt.Printf("Failed to encrypted reader: %v", err) // add error handling
		return
	}

	// the encrypted io.Reader can be used like every other reader - e.g. for copying
	if _, err := io.Copy(os.Stdout, encrypted); err != nil {
		fmt.Printf("Failed to copy data: %v", err) // add error handling
		return
	}
}

func ExampleEncryptWriter() {
	// the master key used to derive encryption keys
	// this key must be keep secret
	masterkey, err := hex.DecodeString("000102030405060708090A0B0C0D0E0FF0E0D0C0B0A090807060504030201000") // use your own key here
	if err != nil {
		fmt.Printf("Cannot decode hex key: %v", err) // add error handling
		return
	}

	// generate a random nonce to derive an encryption key from the master key
	// this nonce must be saved to be able to decrypt the data again - it is not
	// required to keep it secret
	var nonce [32]byte
	if _, err = io.ReadFull(rand.Reader, nonce[:]); err != nil {
		fmt.Printf("Failed to read random data: %v", err) // add error handling
		return
	}

	// derive an encryption key from the master key and the nonce
	var key [32]byte
	kdf := hkdf.New(sha256.New, masterkey, nonce[:], nil)
	if _, err = io.ReadFull(kdf, key[:]); err != nil {
		fmt.Printf("Failed to derive encryption key: %v", err) // add error handling
		return
	}

	output := os.Stdout // customize for your needs - the encrypted output
	encrypted, err := sio.EncryptWriter(output, sio.Config{Key: key[:]})
	if err != nil {
		fmt.Printf("Failed to encrypted writer: %v", err) // add error handling
		return
	}

	// the encrypted io.Writer can be used now but it MUST be closed at the end to
	// finalize the encryption.
	if _, err = io.Copy(encrypted, os.Stdin); err != nil {
		fmt.Printf("Failed to copy data: %v", err) // add error handling
		return
	}
	if err = encrypted.Close(); err != nil {
		fmt.Printf("Failed to finalize encryption: %v", err) // add error handling
		return
	}
}
