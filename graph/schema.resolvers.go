package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"smart_intercom_api/graph/generated"
	"smart_intercom_api/graph/model"
	"smart_intercom_api/internal/login"
	"smart_intercom_api/internal/videos"
)

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	return login.LoginMutation(ctx, input)
}

func (r *mutationResolver) ChangePassword(ctx context.Context, input model.NewPassword) (string, error) {
	return login.ChangePasswordMutation(ctx, input)
}

func (r *queryResolver) RefreshToken(ctx context.Context) (string, error) {
	return login.RefreshTokenQuery(ctx)
}

func (r *mutationResolver) CreateVideo(ctx context.Context, input model.NewVideo) (*model.Video, error) {
	return videos.CreateVideoMutation(ctx, input)
}

func (r *queryResolver) Videos(ctx context.Context) ([]*model.Video, error) {
	return videos.VideosQuery(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
