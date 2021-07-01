package implement

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/grpc/company"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/util"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/validator"
)

type implementation struct {
	validator validator.Validator
	repo      util.Repository
	uuid      util.UUID
}

func NewCompanyGrpcService(validator validator.Validator, repo util.Repository, uuid util.UUID) (service company.Service) {
	return &implementation{validator: validator, repo: repo, uuid: uuid}
}
