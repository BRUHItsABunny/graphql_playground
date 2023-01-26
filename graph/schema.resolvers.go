package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"graphqlPlayground/graph/generated"
	"graphqlPlayground/graph/model"
	"graphqlPlayground/utils"
)

// CreateMovie is the resolver for the createMovie field.
func (r *mutationResolver) CreateMovie(ctx context.Context, input model.NewMovie) (*model.Movie, error) {
	movie := model.Movie{
		Title: input.Title,
		URL:   input.URL,
	}

	_, err := r.DB.NewInsert().Model(&movie).
		// Needed to actually auto-generate values for these columns for Postgres
		ExcludeColumn("id", "release_date").
		// Needed on Postgres to get the row ID after insert
		Returning("id").Exec(ctx)
	if err != nil {
		return nil, utils.WrapError(err, "error inserting new movie")
	}

	return &movie, nil
}

// Movies is the resolver for the movies field.
func (r *queryResolver) Movies(ctx context.Context) ([]*model.Movie, error) {
	var movies []*model.Movie

	err := r.DB.NewSelect().Model(&model.Movie{}).Scan(ctx, &movies)
	if err != nil {
		return nil, utils.WrapError(err, "error selecting movies")
	}

	return movies, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
