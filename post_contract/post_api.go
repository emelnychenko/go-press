package post_contract

import (
	"../common"
	"../post_domain"
)

type (
	PostApi interface {
		ListPosts() ([]*post_domain.Post, common.Error)
		GetPost(postId post_domain.PostId) (*post_domain.Post, common.Error)
	}
)
