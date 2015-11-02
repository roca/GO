package models

import (
	"testing"
)

func Test_GetMember(t *testing.T) {
	member, err := GetMember("test@pluralsight.com", "secret")

	if err == nil {
		if member.Email() != "test@pluralsight.com" {
			t.Log("Member email is not found or incorrect")
			t.Fail()
		}

	} else {
		t.Log(err)
		t.Fail()
	}

}

func Test_CreateSession(t *testing.T) {
	member, err := GetMember("test@pluralsight.com", "secret")
	session, err := CreateSession(member)

	if err == nil {
		if session.MemberId() != 2 {
			t.Log("Session could not be created")
			t.Fail()
		}

	} else {
		t.Log(err)
		t.Fail()
	}

}
