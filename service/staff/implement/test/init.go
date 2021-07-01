package test

import (
	"context"
	"fmt"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/implement"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util/mocks"
	validatorMocks "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/validator/mocks"

	"github.com/stretchr/testify/suite"
)

// PackageTestSuite struct
type PackageTestSuite struct {
	suite.Suite
	ctx       context.Context
	validator *validatorMocks.Validator
	repo      *mocks.Repository
	uuid      *mocks.UUID
	service   staff.Service
}

func (suite *PackageTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.validator = &validatorMocks.Validator{}
	suite.repo = &mocks.Repository{}
	suite.uuid = &mocks.UUID{}
	suite.service = implement.New(suite.validator, suite.repo, suite.uuid)
}

func (suite *PackageTestSuite) makeStaffIDFilter(staffID string) (filters []string) {
	return []string{
		fmt.Sprintf("id:eq:%s", staffID),
	}
}
