package main

import (
	"fmt"
	"net/http"
	"crypto/tls"
	"log"
	"io/ioutil"
	"bytes"
	"flag"
	"golang.org/x/crypto/ssh/terminal"
	"bufio"
	"os"
	"syscall"
	"strings"
)

func showCert(url, cert string) string {
	// getting username and password from user input
	username, password := credentials()

	// Building the URL using the username, pass, LB url
	urlBuild := fmt.Sprintf("https://%s:%s@%s:9070/api/tm/3.8/config/active/ssl/server_keys/%s/", username, password, url, cert)

	// Bypassing secure connection
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
	}
	client := &http.Client{Transport: tr}

	// Building the request using the URL build above
	request, err := http.NewRequest("GET", urlBuild, nil)

	// Returning the request and catching end errors
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	// Putting whole response in string
	bodyText, err := ioutil.ReadAll(response.Body)

	// Returning the response status and body
	fmt.Println("response Status", response.Status)
	s := string(bodyText)
	fmt.Println(s)
	return s
}

func showCertAll(url string) string {
	// getting username and password from user input
	username, password := credentials()

	// Building the URL using the username, pass, LB url
	urlBuild := fmt.Sprintf("https://%s:%s@%s:9070/api/tm/3.8/config/active/ssl/server_keys/", username, password, url)

	// Bypassing secure connection
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
	}
	client := &http.Client{Transport: tr}

	// Building the request using the URL build above
	request, err := http.NewRequest("GET", urlBuild, nil)

	// Returning the request and catching end errors
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	// Putting whole response in string
	bodyText, err := ioutil.ReadAll(response.Body)

	// Returning the response status and body
	fmt.Println("response Status", response.Status)
	s := string(bodyText)
	fmt.Println(s)
	return s
}

func addCert(url, cert, cert_path, key_path string) string {
	// getting the contents of each file to build JSON pay load
	certificate := readFile(cert_path)
	privatekey := readFile(key_path)

	data := []byte(`{"properties":{"basic":{"note":"","private":"`+ privatekey +`","public":"`+ certificate +`"}}}`)

	// getting username and password from user input
	username, password := credentials()

	// Building the URL using the username, pass, LB url
	urlBuild := fmt.Sprintf("https://%s:%s@%s:9070/api/tm/3.8/config/active/ssl/server_keys/%s/", username, password, url, cert)

	// Bypassing secure connection
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
	}
	client := &http.Client{Transport: tr}

	// Building the request using the URL build above
	request, err := http.NewRequest("PUT", urlBuild, bytes.NewBuffer(data))
	request.Header.Set("Content-Type", "application/json")

	// Returning the request and catching end errors
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	// Putting whole response in string
	bodyText, err := ioutil.ReadAll(response.Body)

	// Returning the response status and body
	fmt.Println("response Status", response.Status)
	s := string(bodyText)
	fmt.Println(s)
	return s
}

func delCert(url, cert string) string {
	// getting username and password from user input
	username, password := credentials()

	// Building the URL using the username, pass, LB url
	urlBuild := fmt.Sprintf("https://%s:%s@%s:9070/api/tm/3.8/config/active/ssl/server_keys/%s/", username, password, url, cert)

	// Bypassing secure connection
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
	}
	client := &http.Client{Transport: tr}

	// Building the request using the URL build above
	request, err := http.NewRequest("DELETE", urlBuild, nil)

	// Returning the request and catching end errors
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	// Putting whole response in string
	bodyText, err := ioutil.ReadAll(response.Body)

	// Returning the response status and body
	fmt.Println("response Status", response.Status)
	s := string(bodyText)
	fmt.Println(s)
	return s
}

func readFile(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	cert := string(data)
	//fmt.Println(cert)
	return cert
}

func credentials() (string, string) {
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

	flag.Parse()

	if *OPTION == "" || *LOADBALANCER == "" || *NAME == "" {
		fmt.Println("Flags 'option', 'loadbalancer' and 'name' require values")
		fmt.Println("Example: cert-manager.go -option=show -loadbalancer=example-vtm-node-01.com -name=example.cert.com")
	} else {
		switch *OPTION {
			case "show":

				if *NAME == "all" {
					// Showing all certs on LB
					fmt.Println("Showing all Certs")
					showCertAll(*LOADBALANCER)
				} else {
					// Showing specified cert
					showCert(*LOADBALANCER, *NAME)
				}

			case "add":
				if *CERTIFICATE == "" {
					fmt.Println("Flags 'cert' and 'key' require values when adding a new cert")
					fmt.Println("Example: cert-manager.go -option=add -loadbalancer=example-vtm-node-01.com -name=example.cert.com -cert=/path/to/cert -key=/path/to/key")
				} else {
					// Adding a specified certificate
					fmt.Println("Add Cert")
					addCert(*LOADBALANCER, *NAME, *CERTIFICATE, *PRIVATEKEY)
				}

			case "delete":
				// Deleting a specified certificate
				fmt.Println("Deleting Cert")
				delCert(*LOADBALANCER, *NAME)

			default:
				// Default message if user does not specify an option
				fmt.Println("Please specify option 'show' or 'add' or 'delete'.")
			}
	}
}