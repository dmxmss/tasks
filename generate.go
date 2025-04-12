package main 

//go:generate redocly bundle specification/main.yaml --output openapi.yaml
//go:generate go tool oapi-codegen -config oapi-cfg.yaml openapi.yaml
