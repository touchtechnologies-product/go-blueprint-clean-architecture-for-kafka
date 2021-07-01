package company

import (
	"context"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/company/out"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/company/companyin"
)

//go:generate mockery --name=Service
type Service interface {
	List(ctx context.Context, input *domain.PageOption) (total int, items []*out.CompanyView, err error)
	Create(ctx context.Context, input *companyin.CreateInput) (ID string, err error)
	Read(ctx context.Context, input *companyin.ReadInput) (company *out.CompanyView, err error)
	Update(ctx context.Context, input *companyin.UpdateInput) (err error)
	Delete(ctx context.Context, input *companyin.DeleteInput) (err error)
}
