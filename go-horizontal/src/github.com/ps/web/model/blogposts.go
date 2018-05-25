package model

import (
	"github.com/ps/entity"
	"github.com/ps/web/data"
)

var (
	blogPostRepository data.BlogPostRepository = data.NewBlogPostRepository()
	commentRepository  data.CommentRepository  = data.NewCommentRepository()
)

func GetLastPosts(count int) ([]*entity.BlogPost, error) {
	return blogPostRepository.GetRecentPosts(count)
}

func GetLastPostTitles(count int) ([]*data.BlogSummary, error) {
	return blogPostRepository.GetRecentTitles(count)
}

func CreateBlogPost(post *entity.BlogPost) (*entity.BlogPost, error) {
	return blogPostRepository.CreatePost(post)
}

func UpdateBlogPost(post *entity.BlogPost) (*entity.BlogPost, error) {
	return blogPostRepository.UpdatePost(post)
}

func GetPostById(postId int) (*entity.BlogPost, error) {
	return blogPostRepository.GetById(postId)
}

func CreateComment(comment *entity.Comment, postId int) (*entity.Comment, error) {
	return commentRepository.CreateComment(comment, postId)
}
