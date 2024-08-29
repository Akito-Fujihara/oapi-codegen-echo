.PHONY: oapi-codegen

oapi-codegen = go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen

oapi-codegen:
	$(oapi-codegen) -generate types,server,spec -package api -o ./oapi/api_gen.go ./oapi/openapi.yaml
