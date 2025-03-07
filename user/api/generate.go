package generate

//go:generate go run github.com/ogen-go/ogen/cmd/ogen@latest --package api_desc --target ../internal/pkg/api --clean api.yaml
//go:generate go run github.com/ogen-go/ogen/cmd/ogen@latest --package user_desc --target ../../api/internal/pkg/user --clean api.yaml
