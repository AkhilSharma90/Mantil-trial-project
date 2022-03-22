package movie

import (
	auth "github.com/akhil/auth-project/api/auth"
	"context"

)

type GetResponse struct {
	Movies []MovieItem
	ErrorMsg string
}

type TokenRequest struct{
	Token string
}


func (t *Movie) Get(ctx context.Context, req *TokenRequest) (*GetResponse) {

	value := auth.Validate(req.Token)
	var movies []MovieItem
	if value == "valid"{
	
	_, err := t.kv.FindAll(&movies)
	if err != nil {
		return &GetResponse{
			nil,
			"not ok",
		}
	}

	return &GetResponse{
		movies,
		"",
	}

	}else{
		return &GetResponse{
			nil,
			"token invalid",
		}
	}
}
