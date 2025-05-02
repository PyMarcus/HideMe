package main

import (
	"log"
	"os"

	"github.com/PyMarcus/hideme/common"
	"github.com/PyMarcus/hideme/decrypt"
	"github.com/PyMarcus/hideme/encrypt"
	"github.com/spf13/cobra"
)

var filePath string
var password string
var ecrypt bool
var dcrypt bool

var cmd = &cobra.Command{
	Use:   "hideme",
	Short: "Encrypt your texts files",
	Long:  "A tool to encrypt/decrypt text files based on password",
	Run: func(command *cobra.Command, args []string) {
		
		if filePath == "" {
			log.Println("[-] filePath is not defined (--file)")
			return
		}

		if password == "" {
			log.Println("[-] password is not defined (--p)")
			return
		}

		if ecrypt && dcrypt {
			log.Println("[-] Invalid command")
			os.Exit(1)
		}
		key := common.GenerateKeyFromPassword(password)

		pb := []byte(key)
		
		if ecrypt {
			log.Println("🔒 Encrypting file...:", filePath)
			err := encrypt.EncryptFile(pb, filePath)
			if err != nil{
				log.Println(err)
				os.Exit(1)
			}
			log.Println("OK")
		} else if dcrypt {
			log.Println("🔓 Decrypting file:", filePath)
			err := decrypt.DecryptFile(pb, filePath)
			if err != nil{
				log.Println(err)
				os.Exit(1)
			}
			log.Println("OK")
		} else {
			log.Println("You need to give -e to encrypt or -d to decrypt with -p for password")
			os.Exit(1)
		}

	},
}

func init() {
	cmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "path to file")
	cmd.PersistentFlags().StringVarP(&password, "password", "p", "", "password")
	cmd.PersistentFlags().BoolVarP(&dcrypt, "decrypt", "d", false, "Decrypt file")
	cmd.PersistentFlags().BoolVarP(&ecrypt, "encrypt", "e", false, "Encrypt file")
}

func main() {
	cmd.Execute()
}
