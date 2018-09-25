package owlfile

import(
	"io/ioutil"
	"os"
	"log"
	"gopkg.in/yaml.v2"
)

type Config struct{
	Dir         string    `yaml:"dir"`
	Pattern     string    `yaml:"pattern"`
	Recursive   bool      `yaml:"recursive"`	
	Conditions  string  `yaml:"conditions"`
	Operation   string    `yaml:"opration"`
	ArchivePath string    `yaml:"archive_path"`
}

func LoadYaml(filepath string) Config {

	content, err := ioutil.ReadFile(filepath)

	if err != nil{
		log.Fatalf("IO failed with yaml file")
		os.Exit(1) 
	}


	var config Config
	
	yamlerr := yaml.Unmarshal(content, &config)

	if yamlerr != nil {
		log.Fatalf("unmarshal failed with yaml")
		os.Exit(1)
	}

	return config
}