package implement

import (
	"context"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/staff/out"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/staff/staffin"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/util"
)

func (impl *implementation) Read(ctx context.Context, input *staffin.ReadInput) (view *out.StaffView, err error) {
	staff := &domain.Staff{}
	filters := makeStaffIDFilters(input.StaffID)

	err = impl.repo.Read(ctx, filters, staff)
	if err != nil {
		return nil, util.RepoReadErr(err)
	}

	return out.StaffToView(staff), nil
}
