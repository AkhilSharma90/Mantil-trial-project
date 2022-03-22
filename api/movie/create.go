package movie

import (
	"context"
	auth "github.com/akhil/auth-project/api/auth"
	"github.com/google/uuid"

)

type AddRequest struct {
	Title string
	Rating string
	Token string
}

type AddResponse struct {
	Message string
}

func (t *Movie) Create(ctx context.Context, req *AddRequest) (*AddResponse) {
	value := auth.Validate(req.Token)
	if value == "valid"{
	id := uuid.NewString()
	t.kv.Put(id, &MovieItem{
		ID:        id,
		Title:     req.Title,
		Rating:    req.Rating,
	})
	
	return &AddResponse{
		Message: "data added",
		}
	} else{
		return &AddResponse{
			Message: "token invalid",
		}
	}
}
