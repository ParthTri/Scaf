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

func TestFindTemplate(t *testing.T) {
	tmp := ConfigDir{}

	file_path, _ := FindConfigDir()

	tmp.Path = file_path
	tmp.Templates, _ = os.ReadDir(file_path)
	
	_, err := tmp.FindTemplate("example.toml")

	if err != nil {
		t.Fail()
	}
}
