package main

import (
	"fmt"
	"flag"
	"bufio"
	"os"
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
	"log"
	"strings"
	"github.com/bennyg93/brocade-certificates"
)

func Credentials() (string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	}
	password := string(bytePassword)

	return strings.TrimSpace(username), strings.TrimSpace(password)
}

func main() {
	// Command-line arguments
	OPTION := flag.String("option", "", "Command option. 'show' 'add' 'delete'")
	LOADBALANCER := flag.String("loadbalancer", "", "Enter the URL for the Load Balancer")
	NAME := flag.String("name", "", "The name of the certificate")
	CERTIFICATE := flag.String("cert", "", "Path to the certificate file")
	PRIVATEKEY := flag.String("key", "", "Path to the private key")
	API := flag.String("api", "3.8", "Brocade VTM API version number")

	flag.Parse()

	if *OPTION == "" || *LOADBALANCER == "" || *NAME == "" {
		fmt.Println("Flags 'option', 'loadbalancer' and 'name' require values")
		fmt.Println("Example: brocade.go -option=show -loadbalancer=h1sta01-v00.devops.stg2.ovp.bskyb.com -name=ovp.bskyb.com")
	} else {
		switch *OPTION {
		case "show":

			if *NAME == "all" {
				// getting username and password from user input
				username, password := Credentials()
				// Showing all certs on LB
				fmt.Println("Showing all Certs")
				brocade.Showall(*LOADBALANCER, username, password, *API)
			} else {
				// getting username and password from user input
				username, password := Credentials()
				// Showing specified cert
				brocade.Showcert(*LOADBALANCER, *NAME, username, password, *API)
			}

		case "add":
			if *CERTIFICATE == "" {
				fmt.Println("Flags 'cert' and 'key' require values when adding a new cert")
				fmt.Println("Example: brocade.go -option=add -loadbalancer=h1sta01-v00.devops.stg2.ovp.bskyb.com -name=ovp.bskyb.com -cert=/path/to/cert -key=/path/to/key")
			} else {
				// getting username and password from user input
				username, password := Credentials()
				// Adding a specified certificate
				fmt.Println("Add Cert")
				brocade.Addcert(*LOADBALANCER, *NAME, *CERTIFICATE, *PRIVATEKEY, username, password, *API)
			}

		case "delete":
			// getting username and password from user input
			username, password := Credentials()
			// Deleting a specified certificate
			fmt.Println("Deleting Cert")
			brocade.Delcert(*LOADBALANCER, *NAME, username, password, *API)

		default:
			// Default message if user does not specify an option
			fmt.Println("Please specify option 'show' or 'add' or 'delete'.")
		}
	}
}