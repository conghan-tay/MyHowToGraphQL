package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	db "HowToGraphql/db/sqlc"
	"HowToGraphql/graph/generated"
	"HowToGraphql/graph/model"
	"HowToGraphql/internal/auth"
	"HowToGraphql/internal/middleware"
	"HowToGraphql/pkg/jwt"
	"context"
	"errors"
	"fmt"
	"strconv"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return &model.Link{}, fmt.Errorf("access denied")
	}


	n, err := strconv.ParseInt(user.ID, 10, 64)
	if err != nil {
		return &model.Link{}, fmt.Errorf("error with id")
	}

	result, err := db.DbQueries.CreateLink(ctx, db.CreateLinkParams {
		Title: input.Title,
		Address: input.Address,
		UserID: n,
	})

	if err != nil {
		return nil, err
	}

	link := &model.Link{
		ID:      strconv.FormatInt(result.ID, 10),
		Title:   result.Title,
		Address: result.Address,
		User: &model.User{
			ID:   user.ID,
		},
	}
	
	return link, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	hashedPassword, err := auth.HashPassword(input.Password)
	if err != nil {
		return "", err
	}
	_, err = db.DbQueries.CreateUser(ctx, db.CreateUserParams{
		Username: input.Username,
		Password: hashedPassword,
	})
	if err != nil {
		return  "", err
	}
	token, err := jwt.GenerateToken(input.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	if !input.Authenticate(ctx) {
		return "", errors.New("Fail to authenticate")
	}
	token, err := jwt.GenerateToken(input.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	result, err := db.DbQueries.GetLinks(ctx)
	if err != nil {
		return nil, err
	}

	var links []*model.Link
	for i:= range result {
		link := &model.Link{
			ID:      strconv.FormatInt(result[i].ID, 10),
			Title:   result[i].Title,
			Address: result[i].Address,
			User : &model.User{
				ID:   strconv.FormatInt(result[i].UserID,10),
				Name: result[i].Username,
			},
		}
		links = append(links, link)
	}

	return links, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
