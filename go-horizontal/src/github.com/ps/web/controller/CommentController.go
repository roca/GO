package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ps/entity"
	"github.com/ps/web/model"
)

type CommentController struct{}

func (c *CommentController) createComment(w http.ResponseWriter, r *http.Request) {
	matches := commentsPath.FindStringSubmatch(r.URL.Path)

	//no need to check for error since regex guarantees an integer value
	postID, _ := strconv.Atoi(matches[1])

	r.ParseForm()
	now := time.Now()
	comment := &entity.Comment{
		ContentItem: entity.ContentItem{
			Subject:     r.FormValue("subject"),
			Body:        r.FormValue("body"),
			Author:      nil,
			Comments:    []entity.Comment{},
			CreatedDate: &now,
			PublishDate: &now,
			IsPublished: true,
		},
	}

	comment, err := model.CreateComment(comment, postID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("Location", "/posts/"+strconv.Itoa(postID))
	w.WriteHeader(http.StatusSeeOther)
}
