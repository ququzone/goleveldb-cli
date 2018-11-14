package main

import (
	"fmt"
	"os"
	"strings"

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
	cmds := strings.Split(cmd, " ")
	switch cmds[0] {
	case "":
	case "help":
		printHelp()
	case "exit":
		s.DB.Close()
		os.Exit(0)
	case "keys":
		s.keys()
	case "get":
		s.get(cmds)
	case "put":
		s.put(cmds)
	case "del":
		s.del(cmds)
	default:
		fmt.Printf("Unknow cmd %s.\n", cmd)
		printHelp()
	}
}

func (s *Session) keys() {
	iter := s.DB.NewIterator(nil, nil)
	keys := ""
	for iter.Next() {
		key := iter.Key()
		keys = string(key) + "\n"
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Print(keys)
	}
}

func (s *Session) get(cmds []string) {
	if len(cmds) != 2 {
		fmt.Println("error get cmd")
		return
	}
	v, err := s.DB.Get([]byte(cmds[1]), nil)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Println(string(v))
	}
}

func (s *Session) put(cmds []string) {
	if len(cmds) != 3 {
		fmt.Println("error put cmd")
		return
	}
	err := s.DB.Put([]byte(cmds[1]), []byte(cmds[2]), nil)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}

func (s *Session) del(cmds []string) {

}

func printHelp() {
	fmt.Println(`Commands:
	keys     list the keys in the current range
	get      get a key from the database <str>
	put      put a key/value into the database <str>
	del      delete a key/value from the database <str>
	exit     quit cli
	help     print this list of REPL commands
	`)
}
