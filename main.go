package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"syscall"
	"time"

	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/oauth2/clientcredentials"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("teslatoken: ")
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: teslatoken <email>\n")
		os.Exit(2)
	}

	var (
		reader   = bufio.NewReader(os.Stdin)
		email    = os.Args[1]
		password = readPassword(reader)

		config = &clientcredentials.Config{
			ClientID:     "81527cff06843c8634fdc09e8ac0abefb46ac849f38fe1e431c2ef2106796384",
			ClientSecret: "c7257eb71a564034f9419ee651c7d0e5f7aa6bfbd18bafb5c5c033b093bb2fa3",
			TokenURL:     "https://owner-api.teslamotors.com/oauth/token",
			EndpointParams: url.Values{
				"grant_type":    []string{"password"},
				"email":         []string{email},
				"password":      []string{password},
				"client_id":     []string{"81527cff06843c8634fdc09e8ac0abefb46ac849f38fe1e431c2ef2106796384"},
				"client_secret": []string{"c7257eb71a564034f9419ee651c7d0e5f7aa6bfbd18bafb5c5c033b093bb2fa3"},
			},
		}
	)

	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Access token: %s\n", token.AccessToken)
	fmt.Printf("Refresh token: %s\n", token.RefreshToken)
	fmt.Printf("Expiry: %s\n", token.Expiry.Format(time.RFC1123))
}

func readPassword(r *bufio.Reader) string {
	for {
		fmt.Printf("Password: ")
		password, err := terminal.ReadPassword(int(syscall.Stdin))
		fmt.Println("")
		if err != nil {
			log.Fatalln(err)
		}
		if len(password) != 0 {
			return string(password)
		}
	}
}
