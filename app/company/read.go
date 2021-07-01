package company

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/company/companyin"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/app/view"
)

// Read godoc
// @Tags Companies
// @Summary Get a company by company ID
// @Description Response a company data with a given company ID
// @param company_id path string true "Company ID"
// @Produce  json
// @Success 200 {object} view.SuccessResp{data=out.CompanyView}
// @Success 204 {object} view.SuccessResp
// @Success 401 {object} view.ErrResp
// @Success 403 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Success 503 {object} view.ErrResp
// @Router /companies/{company_id} [get]
func (ctrl *Controller) Read(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.company.Read",
	)
	defer span.Finish()

	input := &companyin.ReadInput{CompanyID: c.Param("id")}
	company, err := ctrl.service.Read(ctx, input)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, company)
}

//func (ctrl *Controller) Read(c *gin.Context) {
//	span, ctx := opentracing.StartSpanFromContextWithTracer(
//		c.Request.Context(),
//		opentracing.GlobalTracer(),
//		"handler.company.Read",
//	)
//	defer span.Finish()
//
//	input := &companyin.ReadInput{CompanyID: c.Param("id")}
//
//	company, err := ctrl.service.Read(ctx, input)
//	if err != nil {
//		view.MakeErrResp(c, err)
//		return
//	}
//
//	view.MakeSuccessResp(c, http.StatusOK, company)
//}
