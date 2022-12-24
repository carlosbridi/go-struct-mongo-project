package configurations_yaml

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

const MSG_YML_BEANS = "Initializing yml properties bean."

var once sync.Once
var YmlConfigs *map[string]Property

type Property struct {
	Value string
}

func GetBeanPropertyByName(propertyName string) string {
	properties := *GetYmlConfigBean()
	property := properties[propertyName]
	return property.Value
}

func GetYmlConfigBean() *map[string]Property {

	once.Do(func() { // <-- atomic, does not allow repeating

		if YmlConfigs == nil {
			YmlConfigs = getYmlConfig()
		} // <-- thread safe

	})
	return YmlConfigs
}

func getYmlConfig() *map[string]Property {

	log.Println(MSG_YML_BEANS)
	yfile, err := os.ReadFile("./application.yaml")
	if err != nil {
		log.Fatal(err)
	}

	data := make(map[string]Property)
	err2 := yaml.Unmarshal(yfile, &data)
	if err2 != nil {
		log.Fatal(err2)
	}

	for k, v := range data {
		fmt.Printf("%s: %s\n", k, v)
	}

	return &data
}
