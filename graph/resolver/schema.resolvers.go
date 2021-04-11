package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphql-poc/graph/generated"
	"graphql-poc/graph/model"
	"graphql-poc/graph/myModels"
)

func (r *queryResolver) GetUsers(ctx context.Context) ([]*myModels.User, error) {
	userSlice := []*myModels.User{}

	for i := 0; i < 1; i++ {
		myUser := myModels.User{
			Firstname: "Abhinav",
			Lastname:  "Singh",
		}
		userSlice = append(userSlice, &myUser)
	}

	return userSlice, nil
}

func (r *queryResolver) GetName(ctx context.Context) (*model.Name, error) {
	nameObj := model.Name{
		Firstname: "Abhinav",
		Lastname:  "Singh",
	}
	fmt.Println("Go Static Resolver running")
	return &nameObj, nil
}

func (r *userResolver) Address(ctx context.Context, obj *myModels.User) (*model.Address, error) {
	// ========= PARENT ============ //
	fmt.Println("User firstname ", obj.Firstname)
	fmt.Println("User Lastname ", obj.Lastname)
	// ============================= //
	addressObj := model.Address{
		Road:     "Delhi Road",
		Locality: "Mansarover, Moradabad",
	}

	return &addressObj, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
