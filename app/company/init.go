package company

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/company"
)

type Controller struct {
	service company.Service
}

func New(service company.Service) (ctrl *Controller) {
	return &Controller{service}
}
