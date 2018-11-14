package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	database := flag.String("d", "database", "database dir")
	flag.Parse()

	session, err := NewSession(*database)
	if err != nil {
		fmt.Printf("Create/Open database %s error: %s\n", *database, err.Error())
		os.Exit(1)
	}
	fmt.Printf("Create/Open database %s successful.\n", *database)
	fmt.Print("> ")

	reader := bufio.NewReader(os.Stdin)

	for {
		cmd, _ := reader.ReadString('\n')
		cmd = cmd[:len(cmd)-1]
		session.Exec(cmd)
		fmt.Print("> ")
	}
}
