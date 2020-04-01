package main

import (
  "fmt"
  rcFolder "github.com/russellcain/GorganizeGolders/folderTraversal"
  "regexp"
  "strings"
)

func main() {
  const chosenPath = "."
  chosenFilePattern := regexp.MustCompile(`.*`)
  fmt.Println("Going to run our markdowning function on", chosenPath)
  foundFiles := rcFolder.EmbedPathInMD(chosenPath, chosenFilePattern)
  println(strings.Join(foundFiles, "\n"))
}
