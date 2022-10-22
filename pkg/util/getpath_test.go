package util

import "testing"

func TestGetPath(t *testing.T) {
	result := GetPath()

	if result != "" {
		t.Logf("%s - Success !", t.Name())
	} else {
		t.Errorf("Failed! %s - expected nil, receieved %s",
			t.Name(),
			result)
	}

}
