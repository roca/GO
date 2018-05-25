package main

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/ps/entity"
)

func main() {
	http.HandleFunc("/posts/", HandleRequest)
	http.HandleFunc("/posts", HandleRequest)

	http.ListenAndServe(":4000", nil)
}

var (
	postsPath    = regexp.MustCompile(`^/posts\?*`)
	postPath     = regexp.MustCompile(`^/posts/(\d+)`)
	commentsPath = regexp.MustCompile(`^/posts/(\d+)/comments`)
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	switch r.Method {
	case http.MethodPost:
		if commentsPath.MatchString(r.URL.Path) {
			HandleCreateComment(w, r)
		} else if postsPath.MatchString(r.URL.Path) {
			HandleCreatePost(w, r)
		}
	case http.MethodPut:
		if postPath.MatchString(r.URL.Path) {
			HandleUpdatePost(w, r)
		}
	case http.MethodGet:
		if postPath.MatchString(r.URL.Path) {
			HandleGetPost(w, r)
		} else if postsPath.MatchString(r.URL.Path) {
			HandleGetPosts(w, r)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Unknown request"}`))
	}
}

func HandleCreateComment(w http.ResponseWriter, r *http.Request) {
	matches := commentsPath.FindStringSubmatch(r.URL.Path)

	//no need to check for error since regex guarantees an integer value
	postId, _ := strconv.Atoi(matches[1])
	var blogPost *entity.BlogPost
	for i := 0; i < len(blogPosts); i++ {
		if blogPosts[i].ID == postId {
			blogPost = &blogPosts[i]
			break
		}
	}

	if blogPost == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"Invalid blog post"}`))
		return
	}

	dec := json.NewDecoder(r.Body)
	var comment entity.Comment
	err := dec.Decode(&comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}

	comment.ID = nextCommentId
	nextCommentId++

	blogPost.Comments = append(blogPost.Comments, comment)

	enc := json.NewEncoder(w)
	enc.Encode(comment)

	w.WriteHeader(http.StatusCreated)
}

func HandleCreatePost(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var blogPost entity.BlogPost
	err := dec.Decode(&blogPost)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}

	blogPost.ID = nextBlogPostId
	nextBlogPostId++

	blogPosts = append(blogPosts, blogPost)

	enc := json.NewEncoder(w)
	enc.Encode(blogPost)

	w.WriteHeader(http.StatusCreated)
}

func HandleUpdatePost(w http.ResponseWriter, r *http.Request) {
	matches := postPath.FindStringSubmatch(r.URL.Path)

	//no need to check for error since regex guarantees an integer value
	postId, _ := strconv.Atoi(matches[1])
	wasFound := false
	for i := 0; i < len(blogPosts); i++ {
		if blogPosts[i].ID == postId {
			wasFound = true

			//remove the old record
			blogPosts = append(blogPosts[:i], blogPosts[i+1:]...)
			break
		}
	}

	if !wasFound {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error":"Could not find specified post to update it"}`))
		return
	}

	var blogPost entity.BlogPost
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&blogPost)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}

	blogPost.ID = postId
	blogPosts = append(blogPosts, blogPost)

	enc := json.NewEncoder(w)
	enc.Encode(&blogPost)
}

func HandleGetPost(w http.ResponseWriter, r *http.Request) {
	matches := postPath.FindStringSubmatch(r.URL.Path)

	//no need to check for error since regex guarantees an integer value
	postId, _ := strconv.Atoi(matches[1])
	var blogPost *entity.BlogPost
	for i := 0; i < len(blogPosts); i++ {
		if blogPosts[i].ID == postId {
			blogPost = &blogPosts[i]
			break
		}
	}

	if blogPost == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"Invalid blog post"}`))
		return
	}

	enc := json.NewEncoder(w)
	enc.Encode(blogPost)

	w.WriteHeader(http.StatusOK)
}

func HandleGetPosts(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	enc.Encode(blogPosts)
}
