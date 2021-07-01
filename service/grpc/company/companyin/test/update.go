package test

import "github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/company/companyin"

func (suite *PackageTestSuite) TestUpdateInputToCompanyDomain() {
	given := companyin.MakeTestUpdateInput()

	actual := companyin.UpdateInputToCompanyDomain(given)

	suite.Equal(given.ID, actual.ID)
	suite.Equal(given.Name, actual.Name)
}
