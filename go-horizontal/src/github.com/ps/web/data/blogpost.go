package data

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	"github.com/ps/entity"
)

type BlogPostRepository interface {
	GetRecentTitles(count int) ([]*BlogSummary, error)
	GetRecentPosts(count int) ([]*entity.BlogPost, error)
	CreatePost(post *entity.BlogPost) (*entity.BlogPost, error)
	UpdatePost(post *entity.BlogPost) (*entity.BlogPost, error)
	GetById(postId int) (*entity.BlogPost, error)
}

func NewBlogPostRepository() BlogPostRepository {
	return &blogPostRepository{}
}

type blogPostRepository struct{}

type sortByDate []entity.BlogPost

func (s sortByDate) Len() int {
	return len(s)
}

func (s sortByDate) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortByDate) Less(i, j int) bool {
	return s[i].PublishDate.After(*s[j].PublishDate)
}

func (r *blogPostRepository) GetRecentTitles(count int) ([]*BlogSummary, error) {
	resp, err := http.Get(*dataServiceUrl + "/posts")

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var posts []entity.BlogPost
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&posts)

	if err != nil {
		return nil, err
	}

	sort.Sort(sortByDate(posts))

	result := []*BlogSummary{}
	for i := 0; i < count && i < len(posts); i++ {
		result = append(result, &BlogSummary{
			ID:         posts[i].ID,
			Subject:    posts[i].Subject,
			AuthorName: posts[i].Author.Username,
		})
	}

	return result, nil
}

func (r *blogPostRepository) GetRecentPosts(count int) ([]*entity.BlogPost, error) {
	resp, err := http.Get(*dataServiceUrl + "/posts")

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var posts []entity.BlogPost
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&posts)

	if err != nil {
		return nil, err
	}
	sort.Sort(sortByDate(posts))

	result := []*entity.BlogPost{}
	for i := 0; i < count && i < len(posts); i++ {
		result = append(result, &posts[i])
	}

	return result, nil
}

func (r *blogPostRepository) CreatePost(post *entity.BlogPost) (*entity.BlogPost, error) {
	buf := bytes.Buffer{}
	enc := json.NewEncoder(&buf)

	err := enc.Encode(post)

	if err != nil {
		return nil, err
	}

	resp, err := http.Post(*dataServiceUrl+"/posts", "application/json", &buf)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(post)

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *blogPostRepository) UpdatePost(post *entity.BlogPost) (*entity.BlogPost, error) {
	buf := bytes.Buffer{}
	enc := json.NewEncoder(&buf)

	err := enc.Encode(post)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, *dataServiceUrl+"/posts/"+
		strconv.Itoa(post.ID), &buf)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(post)

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *blogPostRepository) GetById(postId int) (*entity.BlogPost, error) {
	resp, err := http.Get(*dataServiceUrl + "/posts/" + strconv.Itoa(postId))

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var post *entity.BlogPost
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&post)

	if err != nil {
		return nil, err
	}

	return post, nil
}
