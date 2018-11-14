package main

import (
	"fmt"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
)

// Session ...
type Session struct {
	DB *leveldb.DB
}

// NewSession create session
func NewSession(path string) (*Session, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	return &Session{db}, nil
}

// Exec cmd
func (s *Session) Exec(cmd string) {
	switch cmd {
	case "":
	case "help":
		printHelp()
	case "exit":
		s.DB.Close()
		os.Exit(0)
	default:
		fmt.Printf("Unknow cmd %s.\n", cmd)
		printHelp()
	}
}

func printHelp() {
	fmt.Println(`Commands:
	get      g get a key from the database <str>
	put      p put a key/value into the database <str>
	del      rm delete a key/value from the database <str>
	exit     quit cli
	help     print this list of REPL commands
	`)
}
