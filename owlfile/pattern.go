package owlfile

import (
	"regexp"
	"os"
	"log"
)

func IsPattern(fileInfo os.FileInfo, pattern string)(isPattern bool, err error){
	log.Printf("%s %s", pattern, fileInfo.Name())
	isPattern, err = regexp.MatchString(pattern, fileInfo.Name())
	if err != nil {
		log.Fatal(err)
	}
	return isPattern, nil
}