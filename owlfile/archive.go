package owlfile
/*
These method depends on system path
needs:
1. uuidgen
2. tar
*/


import(
	"log"
	"os/exec"
	"strings"
)
func TarArchive(dis string, files []string) (filename string, err error) {
	uuid, err := exec.Command("uuidgen").Output()
	if err != nil{
		log.Fatal(err)
	}
	archive := dis + "/" + strings.Replace(string(uuid[:]), "\n", "", -1) + ".tar.gz"

	exec_command := []string{ "zcvf", archive}
	exec_command = append(exec_command, files...)
	exec_result, err := exec.Command( "tar", exec_command... ).Output()
	log.Printf("Starting archiving")
	if err != nil{
		log.Fatal(err)
	}
	log.Printf(string(exec_result[:]))
	return archive, nil
}