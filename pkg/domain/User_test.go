package domain

import (
	"testing"
)

func TestPasswordBad(t *testing.T) {
	user := NewUser("test", "test")
	if user.CheckPassword("xD") == true {
		t.Errorf("Password check succeded with bad password")
	}
}
func TestPasswordGood(t *testing.T) {
	user := NewUser("test", "test")
	if user.CheckPassword("test") == false {
		t.Errorf("Password check failed with good password")
	}
}
