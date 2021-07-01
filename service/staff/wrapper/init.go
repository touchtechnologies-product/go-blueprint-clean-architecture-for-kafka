package implement

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/staff"
)

type wrapper struct {
	service staff.Service
}

func _(service staff.Service) staff.Service {
	return &wrapper{
		service: service,
	}
}
