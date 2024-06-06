package somniumsystem

import (
	"somnium/internal/module"
	desc "somnium/pkg/api/somnium/v1"
)

type service struct {
	desc.UnimplementedSomniumServiceServer
	module *module.Model
}

func NewSomniumSystem(module *module.Model) *service {
	return &service{
		module: module,
	}
}
