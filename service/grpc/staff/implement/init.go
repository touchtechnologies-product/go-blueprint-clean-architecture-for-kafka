package implement

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/grpc/staff"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/util"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/validator"
)

type implementation struct {
	validator validator.Validator
	repo      util.Repository
	uuid      util.UUID
}

func NewStaffGrpcService(validator validator.Validator, repo util.Repository, uuid util.UUID) (service staff.Service) {
	return &implementation{validator: validator, repo: repo, uuid: uuid}
}
