package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/saurabh-sikchi/ytask/cmd"
	"github.com/saurabh-sikchi/ytask/db"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "ytasks.db")

	must(db.Init(dbPath))

	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
