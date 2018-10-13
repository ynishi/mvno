package main

import "testing"

func TestCreateName(t *testing.T) {
	expected := "pre1.txt"
	name := createName("pre", 1, "file1.txt")
	if name != "pre1.txt" {
		t.Errorf("created name not matched,\n want: %v,\n have: %v", expected, name)
	}
}

func TestIsAvailable(t *testing.T) {
	name1 := createName("file", 1, "file.txt")
	if isAvailable("test/data", name1) {
		t.Errorf("%v is not available in test/data but returned available", name1)
	}
	name2 := createName("file", 2, "file.txt")
	if !isAvailable("test/data", name2) {
		t.Errorf("%v is available in test/data but returned not available", name2)
	}
}

func TestCreateMoveToPath(t *testing.T) {
	path := createMoveToPath("pre", "test/data/file1.txt", "test/new/")
	expected := "test/new/pre0.txt"
	if path != expected {
		t.Errorf("created path not matched,\n want: %v,\n have: %v", expected, path)
	}
}
