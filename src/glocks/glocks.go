package glocks

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

var InValidFileName = errors.New("Invalid file name")

func isFileNameValid(fileName string) bool {
	var chunks = strings.Split(fileName, ".")
	if chunks[len(chunks)-1] != "glx" {
		return false
	} else {
		return true
	}
}

func getFileName(filePath string) string {
	var chunks = strings.Split(filePath, "/")
	return chunks[len(chunks)-1]
}

type GlocksInterpreter struct {
	HadError   bool
	sauce      string
	error_list []error
}

func (glx *GlocksInterpreter) RunFile(filePath string) error {
	var fileName string = getFileName(filePath)
	if !isFileNameValid(fileName) {
		glx.HadError = true
		glx.error_list = append(glx.error_list, InValidFileName)
		log.Fatal(InValidFileName)
		return InValidFileName
	}

	contents, err := os.ReadFile(filePath)
	if err != nil {
		glx.HadError = true
		glx.error_list = append(glx.error_list, err)
		log.Fatal(err)
		return err
	}
	sauce := string(contents)
	fmt.Println(sauce)
	glx.sauce = sauce
	glx.run()
	return nil
}

func (glx *GlocksInterpreter) RunPrompt() {
	fmt.Println("Prompt for the glocks lang. Enjoy prooompting bois")

	for {
		fmt.Print(">>")
		var line string
		fmt.Scanf("%s", &line)
		if line == "" {
			break
		}
		glx.runLine(line)
	}
}
func (glx *GlocksInterpreter) runLine(line string) {
	glx.sauce = line
	source := strings.Fields(glx.sauce)
	for i := 0; i < len(source); i++ {
		fmt.Println(source[i])
	}
}

func (glx GlocksInterpreter) run() {
	source := strings.Fields(glx.sauce)
	for i := 0; i < len(source); i++ {
		fmt.Println(source[i])
	}
}

func NewGlocksInterpreter(sauce string) *GlocksInterpreter {
	return &GlocksInterpreter{
		sauce:      sauce,
		HadError:   false,
		error_list: []error{},
	}
}
