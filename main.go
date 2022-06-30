package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	osArgs := os.Args
	if len(osArgs) < 2 {
		fmt.Printf("missing args\n")
		return
	}
	fileName := os.Args[1]
	if ext := filepath.Ext(fileName); ext != ".mar" {
		fmt.Printf("file name must end with .mar\n")
		return
	}
	bx, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("error occured while reading '%s': %s\n", fileName, err.Error())
	}
	source := string(bx)
	l := NewLexer(source)
	p := NewParser(l)
	program := p.Parse()
	generatedProgram := generatePy(program, 0)
	log.Printf("[COMPILED] \n%s\n", generatedProgram)
	fmt.Println("=======================================")
	generatedFileName := fileName + "_gen.py"
	os.WriteFile(generatedFileName, []byte(generatedProgram), 0777)
	log.Printf("./%s \n", generatedFileName)

	cmd := exec.Command("./" + generatedFileName)
	bx, err = cmd.Output()
	if err != nil {
		panic("error while executing '" + cmd.String() + "': " + err.Error() + "\n")
	}
	fmt.Println(string(bx))
}
