package models

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"
)

type Member struct {
	email     string
	id        int
	password  string
	firstName string
}

func (this *Member) Email() string {
	return this.email
}
func (this *Member) Id() int {
	return this.id
}
func (this *Member) Password() string {
	return this.password
}
func (this *Member) FirstName() string {
	return this.firstName
}

func (this *Member) SetEmail(value string) {
	this.email = value
}
func (this *Member) SetId(value int) {
	this.id = value
}
func (this *Member) SetPassword(value string) {
	this.password = value
}
func (this *Member) SetFirstName(value string) {
	this.firstName = value
}

type Session struct {
	id        int
	memberId  int
	sessionId string
}

func (this *Session) Id() int {
	return this.id
}
func (this *Session) MemberId() int {
	return this.memberId
}
func (this *Session) SessionId() string {
	return this.sessionId
}
func (this *Session) SetId(value int) {
	this.id = value
}
func (this *Session) SetMemberId(value int) {
	this.memberId = value
}
func (this *Session) SetSesionId(value string) {
	this.sessionId = value
}

func GetMember(email string, password string) (Member, error) {
	db, err := getDBConnection()

	if err == nil {
		defer db.Close()
		pwd := sha256.Sum256([]byte(password))
		query := fmt.Sprintf("SELECT id, email, first_name FROM MEMBER WHERE email LIKE '%s' AND password LIKE '%s'", email, hex.EncodeToString(pwd[:]))
		row := db.QueryRow(query)
		result := Member{}
		err = row.Scan(&result.id, &result.email, &result.firstName)
		if err == nil {
			return result, nil
		} else {
			return result, errors.New("Unable to find Member with email: " + email)
		}
	} else {
		return Member{}, errors.New("Unable to get database connection")
	}
}

func CreateSession(member Member) (Session, error) {
	result := Session{}
	result.memberId = member.Id()
	sessionId := sha256.Sum256([]byte(member.Email() + time.Now().Format("12:00:00")))
	result.sessionId = hex.EncodeToString(sessionId[:])

	db, err := getDBConnection()
	if err == nil {
		defer db.Close()
		query := fmt.Sprintf("INSERT INTO WEB_SESSION (MEMBER_ID, SESSION_ID) VALUES (%d, '%s')", member.Id(), result.sessionId)
		record, err := db.Exec(query)
		if err == nil {
			id, _ := record.LastInsertId()
			fmt.Println(id)
			result.SetId(int(id)) //tring to get .Scan(&result.id)
			return result, nil
		} else {
			return Session{}, errors.New("Unable to save session to database: " + err.Error())
		}
	} else {
		return result, errors.New("Unable to get database connection")
	}
}

func GetMemberBySessionId(sessionId string) (Member, error) {
	result := Member{}

	db, err := getDBConnection()
	if err == nil {
		err := db.QueryRow(`
			SELECT member.first_name
			FROM web_session
			JOIN member
			  ON member.id = web_session.member_id
			WHERE web_session.session_id = $1`, sessionId).Scan(&result.firstName)
		if err == nil {
			return result, nil
		} else {
			return Member{}, errors.New("Unable to get member for web_session")
		}
	} else {
		return result, errors.New("Unable to getdatabase connection")
	}
}
