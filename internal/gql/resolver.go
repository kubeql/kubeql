//go:generate gorunpkg github.com/99designs/gqlgen

package gql

import (
	context "context"
	time "time"
)

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateReview(ctx context.Context, episode Episode, review ReviewInput) (*Review, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Hero(ctx context.Context, episode *Episode) (Character, error) {
	panic("not implemented")
}
func (r *queryResolver) Reviews(ctx context.Context, episode Episode, since *time.Time) ([]Review, error) {
	panic("not implemented")
}
func (r *queryResolver) Search(ctx context.Context, text string) ([]SearchResult, error) {
	panic("not implemented")
}
func (r *queryResolver) Character(ctx context.Context, id string) (Character, error) {
	panic("not implemented")
}
func (r *queryResolver) Droid(ctx context.Context, id string) (*Droid, error) {
	panic("not implemented")
}
func (r *queryResolver) Human(ctx context.Context, id string) (*Human, error) {
	panic("not implemented")
}
func (r *queryResolver) Starship(ctx context.Context, id string) (*Starship, error) {
	panic("not implemented")
}
