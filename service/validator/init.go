package validator

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/util"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/company/companyin"

	"github.com/go-playground/validator/v10"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/domain"
)

type GoPlayGroundValidator struct {
	validate    *validator.Validate
	companyRepo util.Repository
	staffRepo   util.Repository
}

func New(companyRepo util.Repository, staffRepo util.Repository) (v *GoPlayGroundValidator) {
	v = &GoPlayGroundValidator{
		validate:    validator.New(),
		companyRepo: companyRepo,
		staffRepo:   staffRepo,
	}

	v.validate.RegisterStructValidation(v.CompanyCreateStructLevelValidation, &companyin.CreateInput{})
	v.validate.RegisterStructValidation(v.CompanyUpdateStructLevelValidation, &companyin.UpdateInput{})
	v.validate.RegisterStructValidation(v.PageOptionStructLevelValidation, &domain.PageOption{})

	return v
}

func (v *GoPlayGroundValidator) Validate(item interface{}) (err error) {
	return v.validate.Struct(item)
}
