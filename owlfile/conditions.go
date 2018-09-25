package owlfile

import(
	"regexp"
	"os"
	"log"
	"time"
)


const (
	Usage = `
	PATTERN LIKE: COMPARE OPERATOR NUMBER UNIT
	UNIT(defined by Golang time package): 
	s: second
	m: minute
	h: hour
	Example:
	C>240h: created more than ten days
	M<2h:  moditied in last 2 hour 
	`

	VALID_PATTERN   = `[A|C|M][>|<][0-9]+[s|i|h|d]`
	
	ACCESS_COMPARE  = "A"
	CREATE_COMPARE  = "C"
	UPDATE_COMPARE  = "M"

	GREATE_OPERATOR = ">"
	LESS_OPERATOR   = "<"

	UNIT_SECOND     = "s"
	UNIT_MINITE     = "i"
	UNIT_HOUR       = "h"
	UNIT_DAY        = "d"
)

func isValidPattern(pattern string)(isValid bool, err error ){
	isValid, err = regexp.MatchString(VALID_PATTERN, pattern)
	if err != nil{
		log.Fatal(err)
	}
	return isValid, nil
}

func meetCondition(fileInfo os.FileInfo, condition string)( isMeet bool, err error){
	isMeet = false
	isValid, validError := isValidPattern(condition)
	if validError != nil{
		log.Fatal(validError)
	}
	if !isValid {
		log.Fatalf("Wrong express of condition: %s \nUsage: %s", condition, Usage)
	}else{
		compare := condition[0:1]
		operation := condition[1:2]
		durationStr := condition[2: len(condition)]
		
		duration, parseError := time.ParseDuration(durationStr)

		if parseError != nil{
			log.Fatal(parseError)
		}

		cmptime, tiError := timeinfo(fileInfo, compare)

		if tiError != nil{
			log.Fatal(tiError)
		}

		jgtime := time.Now().Add(  -1 *  duration )

		//log.Printf("%s  [%s]  %s", jgtime.String(),operation, cmptime.String())

		if operation == GREATE_OPERATOR{
			isMeet = jgtime.After(cmptime)
		} else{
			isMeet = cmptime.After(jgtime)
		}
	}
	return isMeet , nil
}

func MeetConditions(fileInfo os.FileInfo, conditions []string)( isMeet bool, err error) {
	//init
	isMeet = false
	for _, condition := range conditions{
		curMeet, meetError := meetCondition(fileInfo, condition)
		if meetError != nil{
			log.Fatal(meetError)
		}
		isMeet = isMeet || curMeet
		if isMeet{
			log.Printf("%s %s",  condition, fileInfo.Name())
			break
		}
	}
	return isMeet, nil;
}