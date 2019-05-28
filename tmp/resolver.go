//go:generate gorunpkg github.com/99designs/gqlgen

package tmp

import (
	context "context"

	main "github.com/contesini/Gest-o-de-clientes-Banco-Uati/server"
	prisma "github.com/prisma/prisma-examples/go/graphql/prisma-client"
)

type Resolver struct{}

func (r *Resolver) Mutation() main.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Post() main.PostResolver {
	return &postResolver{r}
}
func (r *Resolver) Query() main.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) User() main.UserResolver {
	return &userResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) SignupUser(ctx context.Context, email string, name *string) (prisma.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateDraft(ctx context.Context, title string, content *string, authorEmail string) (prisma.Post, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeletePost(ctx context.Context, id string) (*prisma.Post, error) {
	panic("not implemented")
}
func (r *mutationResolver) Publish(ctx context.Context, id string) (*prisma.Post, error) {
	panic("not implemented")
}

type postResolver struct{ *Resolver }

func (r *postResolver) Author(ctx context.Context, obj *prisma.Post) (prisma.User, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Feed(ctx context.Context) ([]prisma.Post, error) {
	panic("not implemented")
}
func (r *queryResolver) FilterPosts(ctx context.Context, searchString *string) ([]prisma.Post, error) {
	panic("not implemented")
}
func (r *queryResolver) Post(ctx context.Context, id string) (*prisma.Post, error) {
	panic("not implemented")
}

type userResolver struct{ *Resolver }

func (r *userResolver) Posts(ctx context.Context, obj *prisma.User) ([]prisma.Post, error) {
	panic("not implemented")
}
