package owlfile

import(
	"os"
	"io/ioutil"
	"log"
	"strings"
)


const (
	OPERATION_ARCHIVE = "archive"
	OPERATION_DELETE = "delete"
)

func Owlfile(config Config){
	log.Printf("Traverse path: %s; Recursive: %t \n", config.Dir, config.Recursive)
	filepaths, err := traverseDir(config)
	if err != nil {
		log.Fatal(err)
	}
	total := len(filepaths)
	
	if config.Operation == OPERATION_ARCHIVE{
		archive, arerr := TarArchive(config.ArchivePath, filepaths)
		if arerr != nil{
			log.Fatal(arerr)
		}
		log.Printf("Archive file saved:%s", archive)
	}
	for i, filepath := range filepaths{
		log.Printf("%d %% %d of %d: %s", i * 100 / total , i + 1, total, filepath )
		os.Remove(filepath)
	} 
}

func judgeFile(fileinfo os.FileInfo, config *Config)(res bool, err error){
	isPattern, patternErr := IsPattern(fileinfo, config.Pattern)  
	isMeet, meetErr:= MeetConditions(fileinfo, strings.Split(config.Conditions, ","))
	if patternErr != nil || meetErr != nil{
		log.Fatal(patternErr, meetErr)
	}
	return isPattern && isMeet, nil
}


func traverseDir(config Config) (filepaths []string, err error){

	files, err := ioutil.ReadDir(config.Dir)

	if err != nil{
		log.Fatal(err)
	}

	for _, file := range files{
		filepath := config.Dir + "/" + file.Name()
		if file.IsDir()  {
			if config.Recursive {
				curDir := config.Dir
				config.Dir = filepath
				innerFilepaths, innerErr := traverseDir(config)
				config.Dir = curDir
				if innerErr != nil{
					log.Fatal(err)
				}
				for _, innerFilepath := range innerFilepaths{
					filepaths = append(filepaths, innerFilepath)
				}
			}
		} else{
			flag, err := judgeFile(file, &config)
			
			if err != nil{
				log.Fatal(err)
			}else{
				if flag{
					filepaths = append(filepaths , filepath)
				}else{
					continue
				}
			}
		}
	}
	return filepaths, nil
} 
