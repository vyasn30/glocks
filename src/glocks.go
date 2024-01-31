package main

import (
  "fmt"
  "os"
//  "io"
  "log"
  "strings"
)

func run(sauce string)  {
  source := strings.Fields(sauce);
  for i:=0; i<len(source); i++ {
    fmt.Println(source[i]); 
  }
}

func runFile(filePath string) {
  contents, err := os.ReadFile(filePath);
  if err != nil {
    log.Fatal(err);
    return
  }
  sauce := string(contents);
  fmt.Println(sauce);
  run(sauce);
}


func runPrompt() {
  fmt.Println("Prompt for the glocks lang. Enjoy prooompting bois");

  for ;; {
    fmt.Print(">>");
    var line string;
    fmt.Scanf("%s", &line);
    if line == "" {
      break;
    }
    run(line);
  }
}

func main() {

  if (len(os.Args) > 2) {
    fmt.Println("Usage: glocks[script]");
  }else if (len(os.Args) == 2) {
    runFile(os.Args[1]);
  }else {
    runPrompt();
  }
}
