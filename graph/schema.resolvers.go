package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"math/rand"
	"strconv"
	"time"
	"github.com/hadi-ilies/graphql-server/graph/generated"
	"github.com/hadi-ilies/graphql-server/graph/model"
	"github.com/hadi-ilies/graphql-server/graph/db"
)

var videosRepository db.VideoRepository = db.NewVideoRepository()

// CreateVideo is the resolver for the createVideo field.
func (r *mutationResolver) CreateVideo(ctx context.Context, input model.NewVideo) (*model.Video, error) {
	rand.Seed(time.Now().UnixNano())
	video := &model.Video{
		ID:     strconv.Itoa(rand.Int()),
		Title:  input.Title,
		URL:    input.URL,
		Author: &model.User{ID: input.UserID, Name: "User: " + input.UserID},
	}
	videosRepository.Save(video)
	return video, nil
}

// Videos is the resolver for the videos field.
func (r *queryResolver) Videos(ctx context.Context) ([]*model.Video, error) {
	return videosRepository.FindAll(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
