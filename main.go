package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalln("missing args, try \"withconf <config path> <binary path> <command>\"")
	}

	path, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	var parsedJSON map[string]interface{}

	err = json.Unmarshal(b, &parsedJSON)
	if err != nil {
		log.Fatal(err)
	}

	var command []string
	for i := 3; i < len(os.Args); i++ {
		command = append(command, os.Args[i])
	}

	for k, v := range parsedJSON {
		switch v.(type) {
		case bool:
			command = append(command, "--"+k)
		case string:
			command = append(command, "--"+k+"="+fmt.Sprint(v))
		}
	}

	cmd := exec.Command(os.Args[2], command...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
