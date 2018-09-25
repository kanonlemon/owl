package owlfile


import (
	"syscall"
	"os"
	"time"
	"log"
)

const(
	ACCESS_TIME   = "A"
	CREATE_TIME   = "C"
	MODIFIED_TIME = "M"
)

func timeinfo(fileinfo os.FileInfo, time_type string)(rtime time.Time, err error){
	stat_t := fileinfo.Sys().(*syscall.Stat_t)
	switch time_type{
	case ACCESS_TIME: rtime = time.Unix( stat_t.Atim.Sec, stat_t.Atim.Nsec)
	case CREATE_TIME: rtime = time.Unix( stat_t.Ctim.Sec, stat_t.Ctim.Nsec)
	case MODIFIED_TIME: rtime = time.Unix( stat_t.Mtim.Sec, stat_t.Mtim.Nsec)
	default: log.Fatal("unsupport time") 
	}
	return rtime, nil
}