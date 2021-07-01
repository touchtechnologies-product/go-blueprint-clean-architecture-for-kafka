package test

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/company/companyin"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestUpdate() {
	givenInput := companyin.MakeTestUpdateInput()
	givenCompany := domain.MakeTestCompany()
	givenCompanyIDFilter := suite.makeCompanyIDFilter(givenInput.ID)

	suite.repo.On("Read", mock.Anything, givenCompanyIDFilter, &domain.Company{}).Once().Return(nil)
	suite.validator.On("Validate", givenInput).Once().Return(nil)
	givenCompany.ID = ""
	suite.repo.On("Update", mock.Anything, givenCompanyIDFilter, givenCompany).Once().Return(nil)
	err := suite.service.Update(suite.ctx, givenInput)

	suite.NoError(err)
}
