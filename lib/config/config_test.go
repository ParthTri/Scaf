package config

import (
	"fmt"
	"os"
	"testing"
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

func TestReadTesting(t *testing.T) {
	var outcome Template = Template{
		Title: "Go",
		Description: "Template for quickly setting up Go projects",
	};

	path, err := FindConfigDir();

	tmp := ConfigDir{}
	tmp.Path = path;
	tmp.Templates, _ = os.ReadDir(path)

	template, err := tmp.FindTemplate("example.toml")

	recieved, err := ReadTemplate(template);

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if outcome.Title != recieved.Title {
		t.Log("Titles do not match")
		t.Fail()
	}

	if outcome.Description != recieved.Description {
		t.Log("Descriptions do not match")
		t.Fail()
	}
}
