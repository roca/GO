package data

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ps/entity"
)

type CommentRepository interface {
	CreateComment(comment *entity.Comment, postId int) (*entity.Comment, error)
}

func NewCommentRepository() CommentRepository {
	return &commentRepository{}
}

type commentRepository struct{}

func (r *commentRepository) CreateComment(comment *entity.Comment, postId int) (*entity.Comment, error) {
	buf := bytes.Buffer{}
	enc := json.NewEncoder(&buf)
	enc.Encode(comment)

	resp, err := http.Post(*dataServiceUrl+"/posts/"+strconv.Itoa(postId)+"/comments",
		"application/json", &buf)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {

		errorText, _ := ioutil.ReadAll(resp.Body)
		return nil, errors.New(string(errorText))
	}

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&comment)

	if err != nil {
		return nil, err
	}
	return comment, nil
}
