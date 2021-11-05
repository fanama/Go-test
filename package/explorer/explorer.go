package explorer

import (
	"encoding/json"
	"os"
)

func ReadFile(path string) (conf Configuration, err error) {

	file, _ := os.Open(path)

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil {
		return conf, err
	}

	Conf = conf

	return conf, err
}
