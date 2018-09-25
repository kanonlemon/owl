package owlfile

import(
	"regexp"
	"os"
	"log"
	"time"
	"strconv"
)


const (
	Usage = `
	PATTERN LIKE: COMPARE OPERATOR NUMBER UNIT
	UNIT(defined by Golang time package): 
	s: second
	i: minute
	h: hour
	d: day
	m: month
	y: year
	Example:
	C>10d: created more than ten days
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
		count, convErr := strconv.Atoi( condition[2: len(condition) - 1] )
		
		if convErr != nil{
			log.Fatal(convErr)
		}
		
		unit := condition[len(condition) - 1: len(condition) ]
		var duration time.Duration

		switch unit{
		case UNIT_SECOND : duration = time.Second    ; break;
		case UNIT_MINITE : duration = time.Minute    ; break;
		case UNIT_HOUR   : duration = time.Hour      ; break;
		case UNIT_DAY    : duration = time.Hour * 12 ; break;
		}

		cmptime, tiError := timeinfo(fileInfo, compare)

		if tiError != nil{
			log.Fatal(tiError)
		}

		jgtime := time.Now().Add(  -1 * time.Duration(  float64(count) * duration.Seconds()) )

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
	}
	return isMeet, nil;
}