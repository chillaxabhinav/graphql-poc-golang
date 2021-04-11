package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"graphql-poc/graph/model"
	"io/ioutil"
	"net/http"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *queryResolver) GetBitcoinPrice(ctx context.Context, currency *string) (*model.DataBitcoinPrice, error) {
	bitcoinEndpoint := "https://api.coinbase.com/v2/prices/spot?currency=" + *currency

	response, err := http.Get(bitcoinEndpoint)

	if err != nil {
		return nil, gqlerror.Errorf("Internal Server error, Could not get bitcoin rate")
	}

	data, _ := ioutil.ReadAll(response.Body) // => Parsing response.body to []byte

	// bitcoinData := new(model.DataBitcoinPrice) => directly gives the memory address as used new keyword

	bitcoinData := model.DataBitcoinPrice{} // => instantiates DataBitcoinPrice struct, hence provides the value not address

	err2 := json.Unmarshal(data, &bitcoinData)

	if err2 != nil {
		return nil, gqlerror.Errorf("Internal Server error")
	}

	return &bitcoinData, nil
}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
type dataBitcoinPriceResolver struct{ *Resolver }
