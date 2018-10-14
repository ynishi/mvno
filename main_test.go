package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestCreateName(t *testing.T) {
	expected := "pre1.txt"
	name := createName("pre", 1, "%d", "file1.txt")
	if name != "pre1.txt" {
		t.Errorf("created name not matched,\n want: %v,\n have: %v", expected, name)
	}
}

func TestIsAvailable(t *testing.T) {
	name1 := createName("file", 1, "%d", "file.txt")
	if isAvailable("test/data", name1) {
		t.Errorf("%v is not available in test/data but returned available", name1)
	}
	name2 := createName("file", 2, "%d", "file.txt")
	if !isAvailable("test/data", name2) {
		t.Errorf("%v is available in test/data but returned not available", name2)
	}
}

func TestCreateMoveToPath(t *testing.T) {
	path, latest := createMoveToPath("pre", 0, "%d", "test/data/file1.txt", "test/new/")
	if latest != 0 {
		t.Errorf("latest number not matched,\n want: %v,\n have: %v", 0, latest)
	}
	expected := "test/new/pre0.txt"
	if path != expected {
		t.Errorf("created path not matched,\n want: %v,\n have: %v", expected, path)
	}
}

func TestMoveByCopy(t *testing.T) {
	content := []byte("this is testfile")
	tmpfile, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	defer tmpfile.Close()

	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}
	destfilename := fmt.Sprintf("%s_dest", tmpfile.Name())
	err = moveByCopy(tmpfile.Name(), destfilename)
	if err != nil {
		t.Fatal(err)
	}
	if _, err = os.Stat(tmpfile.Name()); os.IsExist(err) {
		t.Errorf("not deleted oldfile:%v", tmpfile.Name())
	}
	movedContent, err := ioutil.ReadFile(destfilename)
	if err != nil {
		t.Fatal(err)
	}
	if string(movedContent) != string(content) {
		t.Errorf("not matched content,\n want:%v,\n have:%v", content, movedContent)
	}

}
