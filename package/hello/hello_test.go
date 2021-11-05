package hello

import "testing"

func TestGreeting(t *testing.T) {
	if res := Greeting(); res != "hello world" {
		t.Errorf("erreur : %v", "No greetin")

	}
}
