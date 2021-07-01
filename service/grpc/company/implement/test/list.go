package test

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/grpc/company/companyin"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/util"
)

func (suite *PackageTestSuite) TestListCompanyGPRCSuccess() {
	opt := companyin.MakeTestCreatePageOptGRPCInput()
	finalOpt := companyin.MakeTestCreatePageOptGRPCInputToPageOtpDomain(opt)
	records := companyin.MakeTestRecordCompanyDomain()

	suite.repo.On("List", mock.Anything, finalOpt, &domain.Company{}).Once().Return(len(records), records, nil)
	_, err := suite.service.List(suite.ctx, opt)

	suite.Len(records, 2)
	suite.NoError(err)
}

func (suite *PackageTestSuite) TestListCompanyGPRCError() {
	opt := companyin.MakeTestCreatePageOptGRPCInput()
	finalOpt := companyin.MakeTestCreatePageOptGRPCInputToPageOtpDomain(opt)
	_ = companyin.MakeTestRecordCompanyDomain()

	givenListRepoError := errors.New("REPOSITORY")
	expectedError := util.RepoListErr(errors.New("REPOSITORY"))

	suite.repo.On("List", mock.Anything, finalOpt, &domain.Company{}).Once().Return(0, nil, givenListRepoError)
	_, err := suite.service.List(suite.ctx, opt)

	suite.Error(err)
	suite.Equal(err, expectedError)
}
