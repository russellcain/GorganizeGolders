package folderTraversal

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func EmbedPathInMD(path string, filePattern *regexp.Regexp) []string {
	mdCommands := iterateOverFiles(path, filePattern)
	return mdCommands
}

func catchError(e error) {
	if e != nil {
		// my best understanding of how Go tackles error handling
		panic(e)
	}
}

func iterateOverFiles(path string, filePattern *regexp.Regexp) []string {
	var mdSnippets []string

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if filePattern.Match([]byte(path)) {
			fileName := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
			relativePath := strings.Replace(path, "/Users/rcain/IdeaProjects/KeystoneUI/mockups", ".", 1)
			//fmt.Println("Add to your file:\n", generateMdForFiles(path, fileName))
			mdSnippets = append(mdSnippets, generateMdForFiles(relativePath, fileName))
		}

		return nil
	})
	catchError(err)
	return mdSnippets
}

func generateMdForFiles(filePath string, fileName string) string {
	return  `
----------
##### ` + strings.ReplaceAll(fileName, "_", " ") + `
Inspiration: 

![alt text](`+ filePath +`)

ToDo:

- ...
    - ...
	`
}