package owlfile

import (
	"regexp"
	"os"
	"log"
)

func IsPattern(fileInfo os.FileInfo, pattern string)(isPattern bool, err error){
	isPattern, err = regexp.MatchString(pattern, fileInfo.Name())
	if err != nil {
		log.Fatal(err)
	}
	if isPattern{
		log.Printf("%s %s", pattern, fileInfo.Name())
	}
	return isPattern, nil
}