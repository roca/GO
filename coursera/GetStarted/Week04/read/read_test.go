package main

import "testing"

func TestReadPeopleFromFile(t *testing.T) {
	var people People
	people, err := ReadPeopleFromFile("./test.txt")

	if err != nil {
		t.Error(err)
	}

	if len(people) == 0 {
		t.Error("No text read from file")
	}

	if people[0].fname != makeByteArrayFromString("John") {
		t.Errorf("%s not equal %s", people[0].fname, "John")
	}

	if len(people[0].fname) != 20 {
		t.Errorf("Field length of '%s' is not equal 20", people[0].fname)
	}

	if people[0].lname != makeByteArrayFromString("Smith") {
		t.Errorf("%s not equal %s", people[0].lname, "Smith")
	}

	if len(people[0].lname) != 20 {
		t.Errorf("Field length of '%s' is not equal 20", people[0].lname)
	}

}
