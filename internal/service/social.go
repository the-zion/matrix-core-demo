package service

import (
	"context"
	v1 "cube-core/api/realworld/v1"
)

func (r *RealWorldService) FollowUser(ctx context.Context, request *v1.FollowUserRequest) (*v1.ProfileReply, error) {
	return &v1.ProfileReply{Profile: &v1.ProfileReply_Profile{
		Username: "",
	}}, nil
}

func (r *RealWorldService) UnfollowUser(ctx context.Context, request *v1.UnfollowUserRequest) (*v1.ProfileReply, error) {
	return &v1.ProfileReply{Profile: &v1.ProfileReply_Profile{
		Username: "",
	}}, nil
}

func (r *RealWorldService) GetProfile(ctx context.Context, request *v1.GetProfileRequest) (*v1.ProfileReply, error) {
	return &v1.ProfileReply{Profile: &v1.ProfileReply_Profile{
		Username: "",
	}}, nil
}

func (r *RealWorldService) GetArticle(cxt context.Context, request *v1.GetArticleRequest) (*v1.SingleArticleReply, error) {
	return &v1.SingleArticleReply{Article: &v1.SingleArticleReply_Article{Title: ""}}, nil
}

func (r *RealWorldService) CreateArticle(cxt context.Context, request *v1.CreateArticleRequest) (*v1.SingleArticleReply, error) {
	return &v1.SingleArticleReply{Article: &v1.SingleArticleReply_Article{Title: ""}}, nil
}

func (r *RealWorldService) UpdateArticle(ctx context.Context, request *v1.UpdateArticleRequest) (*v1.SingleArticleReply, error) {
	return &v1.SingleArticleReply{Article: &v1.SingleArticleReply_Article{Title: ""}}, nil
}

func (r *RealWorldService) DeleteArticle(ctx context.Context, request *v1.DeleteArticleRequest) (*v1.SingleArticleReply, error) {
	return &v1.SingleArticleReply{Article: &v1.SingleArticleReply_Article{Title: ""}}, nil
}

func (r *RealWorldService) ListArticles(ctx context.Context, request *v1.ListArticlesRequest) (*v1.MultipleArticlesReply, error) {
	return &v1.MultipleArticlesReply{Articles: []*v1.MultipleArticlesReply_Articles{}}, nil
}

func (r *RealWorldService) FeedArticles(ctx context.Context, request *v1.FeedArticlesRequest) (*v1.MultipleArticlesReply, error) {
	return &v1.MultipleArticlesReply{Articles: []*v1.MultipleArticlesReply_Articles{}}, nil
}

func (r *RealWorldService) FavoriteArticle(ctx context.Context, request *v1.FavoriteArticleRequest) (*v1.SingleArticleReply, error) {
	return &v1.SingleArticleReply{Article: &v1.SingleArticleReply_Article{Title: ""}}, nil
}

func (r *RealWorldService) UnfavoriteArticle(ctx context.Context, request *v1.UnfavoriteArticleRequest) (*v1.SingleArticleReply, error) {
	return &v1.SingleArticleReply{Article: &v1.SingleArticleReply_Article{Title: ""}}, nil
}

func (r *RealWorldService) AddComment(ctx context.Context, request *v1.AddCommentRequest) (*v1.SingleCommentReply, error) {
	return &v1.SingleCommentReply{Comment: &v1.SingleCommentReply_Comment{Id: 1}}, nil
}

func (r *RealWorldService) DeleteComment(ctx context.Context, request *v1.DeleteCommentRequest) (*v1.SingleCommentReply, error) {
	return &v1.SingleCommentReply{Comment: &v1.SingleCommentReply_Comment{Id: 1}}, nil
}

func (r *RealWorldService) GetComment(ctx context.Context, request *v1.GetCommentRequest) (*v1.MultipleCommentsReply, error) {
	return &v1.MultipleCommentsReply{Comments: []*v1.MultipleCommentsReply_Comments{}}, nil
}

func (r *RealWorldService) GetTags(ctx context.Context, request *v1.GetTagsRequest) (*v1.TagListReply, error) {
	return &v1.TagListReply{Tags: []string{""}}, nil
}
