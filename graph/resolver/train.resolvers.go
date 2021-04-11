package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"graphql-poc/graph/model"
)

func (r *queryResolver) GetTrains(ctx context.Context) ([]*model.Train, error) {
	trainsSlice := []*model.Train{}

	firstTrain := model.Train{
		Number: 12429,
		Name:   "LKO-NDLS AC Exp",
	}

	secondTrain := model.Train{
		Number: 20503,
		Name:   "Dbgr Raj",
	}

	trainsSlice = append(trainsSlice, &firstTrain, &secondTrain)

	return trainsSlice, nil
}
