package test

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/company/companyin"
	"net/http"
	"reflect"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/app/view"
)

type ItemStruct struct {
	Title string
	Body  string
}

func (suite *PackageTestSuite) TestMakeSuccessResp() {
	st := struct {
		Title string `validate:"required"`
		Body  string `validate:"required"`
	}{
		Title: "",
		Body:  "",
	}
	view.MakeSuccessResp(suite.ctx, http.StatusOK, st)
	suite.Equal(http.StatusOK, suite.ctx.Writer.Status())
}

func (suite *PackageTestSuite) TestMakePaginatorResp() {
	items := []ItemStruct{
		{
			Title: "Test",
			Body:  "Test",
		},
	}
	view.MakePaginatorResp(suite.ctx, 1, items)
	suite.Equal(http.StatusOK, suite.ctx.Writer.Status())
}

func (suite *PackageTestSuite) TestMakePaginatorRespNoContent() {
	var items []ItemStruct
	view.MakePaginatorResp(suite.ctx, 0, items)
	suite.Equal(http.StatusNoContent, suite.ctx.Writer.Status())
}

func (suite *PackageTestSuite) TestMakeCreatedResp() {
	newID := "new ID"
	company := &companyin.CreateInput{}
	view.MakeCreatedResp(suite.ctx, newID, reflect.TypeOf(company))
	suite.Equal(http.StatusCreated, suite.ctx.Writer.Status())
	suite.Equal(newID, suite.ctx.Writer.Header().Get("Content-Location"))
}
