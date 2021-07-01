package implement

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/company"
)

type wrapper struct {
	service company.Service
}

func _(service company.Service) company.Service {
	return &wrapper{
		service: service,
	}
}
