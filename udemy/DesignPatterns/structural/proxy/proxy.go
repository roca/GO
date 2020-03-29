package proxy

import (
	"fmt"
)

type User struct {
	ID int32
}

type UserFinder interface {
	FindUser(id int32) (User, error)
}

type UserList []User

func (t *UserList) FindUser(id int32) (User, error) {
	for i := 0; i < len(*t); i++ {
		if (*t)[i].ID == id {
			return (*t)[i], nil
		}
	}
	return User{}, fmt.Errorf("User %d could not be found\n", id)
}
func (t *UserList) addUser(newUser User) {
	*t = append(*t, newUser)
}

type UserListProxy struct {
	MockedDatabase      *UserList
	StackCache          UserList
	StackSize           int
	LastSearchUsedCache bool
}

func (u *UserListProxy) FindUser(id int32) (User, error) {
	//Search for the object in the cache list first
	user, err := u.StackCache.FindUser(id)
	if err == nil {
		fmt.Println("Returning user from cache")
		u.LastSearchUsedCache = true
		return user, nil
	}
	user, err = u.MockedDatabase.FindUser(id)
	if err != nil {
		return User{}, err
	}
	u.addUserToStack(user)
	fmt.Println("Returning user from database")
	u.LastSearchUsedCache = false
	return user, nil
}

func (u *UserListProxy) addUserToStack(user User) {
	if len(u.StackCache) >= u.StackSize {
		u.StackCache = append(u.StackCache[1:], user)
	} else {
		u.StackCache.addUser(user)
	}
}
