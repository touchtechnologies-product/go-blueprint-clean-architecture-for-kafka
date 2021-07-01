package view_test

import (
	"testing"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/app/view/test"

	"github.com/stretchr/testify/suite"
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(test.PackageTestSuite))
}
