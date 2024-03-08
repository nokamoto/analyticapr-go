package command

import (
	"strings"

	v1 "github.com/nokamoto/analyticapr-go/pkg/api/v1"
)

type user struct {
	Login string
	Name  string
}

type comment struct {
	Author user
}

type comments []comment

func (c *comments) toComments() []*v1.Comment {
	if c == nil {
		return nil
	}
	var res []*v1.Comment
	for _, comment := range *c {
		res = append(res, &v1.Comment{
			Author: comment.Author.Login,
		})
	}
	return res
}

type review struct {
	Author user
	State  string
}

type reviews []review

func (r *reviews) toReviews() []*v1.Review {
	if r == nil {
		return nil
	}
	var res []*v1.Review
	for _, review := range *r {
		var state v1.Review_State
		switch strings.ToLower(review.State) {
		case "approved":
			state = v1.Review_STATE_APPROVED
		case "commented":
			state = v1.Review_STATE_COMMENTED
		}
		res = append(res, &v1.Review{
			Author: review.Author.Login,
			State:  state,
		})
	}
	return res
}

type pull struct {
	Author   user
	MergedAt string
	Number   int32
	Title    string
	Comments comments
	Reviews  reviews
}

type pulls []pull

func (p *pull) toPullRequest() *v1.PullRequest {
	return &v1.PullRequest{
		Number:   p.Number,
		Title:    p.Title,
		Author:   p.Author.Login,
		Comments: p.Comments.toComments(),
		Reviews:  p.Reviews.toReviews(),
	}
}

func (p *pulls) toPullRequests() []*v1.PullRequest {
	if p == nil {
		return nil
	}
	var res []*v1.PullRequest
	for _, pull := range *p {
		res = append(res, pull.toPullRequest())
	}
	return res
}
