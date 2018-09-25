package main

import(
	"fmt"
	"flag"
	"log"
	"os"
	owlfile "github.com/owl/owlfile"
)


func main(){
	configFilePath := flag.String("f", "owl.yaml", "the config file of yaml")
	flag.Parse()

	if configFilePath == nil{
		log.Fatalf("must have a config file")
		os.Exit(1)	
	}
	config := owlfile.LoadYaml(*configFilePath)
	owlfile.Owlfile(config)
	fmt.Printf("Owl done")
}