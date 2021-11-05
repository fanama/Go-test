package explorer

import (
	"testing"
)

func TestFile(t *testing.T) {
	_, err := ReadFile("./conf.json")

	if err != nil {
		t.Errorf("erreur : %v", err.Error())
	}

}
