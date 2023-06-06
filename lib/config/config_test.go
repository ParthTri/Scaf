package config

import (
	"os"
	"testing"
	"fmt"
)


func TestFindConfigDir(t *testing.T) {
	path, _:= os.UserConfigDir() 
	path = fmt.Sprintf("%v/scaf/", path)

	_, err := FindConfigDir()
	
	if err != nil {
		t.Fatal(err)
	}
}
