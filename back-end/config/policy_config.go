package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type PolicyConfig struct {
	MaxAmount  int `yaml:"max_amount"`
	MeanAmount int `yaml:"min_amount"`
}

var policy PolicyConfig = PolicyConfig{
	MaxAmount:  10000,
	MeanAmount: 5000,
}

func LoadPolicy(path string) {
	policyF, err := os.ReadFile(path)
	if err != nil {
		log.Println("Unable to load config proceeding with defaul values")
		return
	}

	err = yaml.Unmarshal(policyF, &policy)
	if err != nil {
		log.Printf("Unmarshal: %v\n", err)
	}

	log.Println("Config loaded succesfully")

}

func Get() PolicyConfig {
	return policy
}
