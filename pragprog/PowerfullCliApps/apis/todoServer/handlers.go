package main

import (
	"errors"
	"net/http"
	"sync"

	"github.com/roca/GO/tree/staging/pragprog/PowerfullCliApps/todo"
)

var (
	ErrNotFound    = errors.New("not found")
	ErrInvalidData = errors.New("invalid data")
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	content := "There's an API here"
	replyTextContent(w, r, http.StatusOK, content)
}

func todoRouter(todoFile string, l sync.Locker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		list := &todo.List{}

		l.Lock()
		defer l.Unlock()
		if err := list.Get(todoFile); err != nil {
			replyError(w, r, http.StatusInternalServerError, err.Error())
			return
		}

		if r.URL.Path == "" {
			switch r.Method {
			case http.MethodGet:
				getAllHandler(w, r, list)
			case http.MethodPost:
				addHandler(w, r, list, todoFile)
			default:
				message := "Method not support"
				replyError(w, r, http.StatusMethodNotAllowed, message)
			}
			return
		}

		id, err := validateID(r.URL.Path, list)
		if err != nil {
			if errors.Is(err, ErrNotFound) {
				replyError(w, r, http.StatusNotFound, err.Error())
				return
			}
			replyError(w, r, http.StatusBadRequest, err.Error())
		}

		switch  r.Method {
		case http.MethodGet:
			getOneHandler(w, r, list, id)
		case http.MethodDelete:
			deleteHandler(w, r, list, id, todoFile)
		case http.MethodPatch:
			patchHandler(w, r, list, id, todoFile)
		default:
			message := "Method not support"
			replyError(w, r, http.StatusMethodNotAllowed, message)
		}
	}
}

func getAllHandler(w http.ResponseWriter, r *http.Request, list *todo.List)               {}
func addHandler(w http.ResponseWriter, r *http.Request, list *todo.List, todoFile string) {}
func validateID(urlPath string, list *todo.List) (int, error) {
	return -1, nil
}
func getOneHandler(w http.ResponseWriter, r *http.Request, list *todo.List, id int) {}
func deleteHandler(w http.ResponseWriter, r *http.Request, list *todo.List, id int, todoFile string) {}
func patchHandler(w http.ResponseWriter, r *http.Request, list *todo.List, id int, todoFile string) {}
