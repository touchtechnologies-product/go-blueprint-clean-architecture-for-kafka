package test

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/grpc/company/companyin"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/util"
)

func (suite *PackageTestSuite) TestCreateCompanyGPRCSuccess() {
	givenInput := companyin.MakeTestCreateCompGRPCInput()
	givenCompany := companyin.MakeTestCreateCompGRPCCToCompanyDomain()

	suite.uuid.On("Generate").Once().Return("")
	suite.validator.On("Validate", givenInput).Once().Return(nil)
	suite.repo.On("Create", mock.Anything, givenCompany).Once().Return(givenCompany.ID, nil)
	_, err := suite.service.Create(suite.ctx, givenInput)

	suite.NoError(err)
}

func (suite *PackageTestSuite) TestCreateCompanyGPRCError() {
	givenInput := companyin.MakeTestCreateCompGRPCInput()
	givenCompany := companyin.MakeTestCreateCompGRPCCToCompanyDomain()

	givenCreateRepoError := errors.New("REPOSITORY")
	expectedError := util.RepoCreateErr(errors.New("REPOSITORY"))

	suite.uuid.On("Generate").Once().Return("")
	suite.validator.On("Validate", givenInput).Once().Return(nil)
	suite.repo.On("Create", mock.Anything, givenCompany).Once().Return("", givenCreateRepoError)
	_, err := suite.service.Create(suite.ctx, givenInput)

	suite.Equal(err, expectedError)
}

func (suite *PackageTestSuite) TestCreateCompanyGPRCValidationError() {
	givenInput := companyin.MakeTestCreateCompGRPCInput()

	givenValidateError := errors.New("VALIDATION")
	expectedError := util.ValidationCreateErr(errors.New("VALIDATION"))

	suite.uuid.On("Generate").Once().Return("")
	suite.validator.On("Validate", givenInput).Once().Return(givenValidateError)
	_, err := suite.service.Create(suite.ctx, givenInput)

	suite.Equal(err, expectedError)
}
