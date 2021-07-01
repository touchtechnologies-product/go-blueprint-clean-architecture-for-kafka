package implement

import (
	"context"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/staff/staffin"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/util"
)

func (impl *implementation) Delete(ctx context.Context, input *staffin.DeleteInput) (err error) {
	staff := &domain.Staff{}
	filters := makeStaffIDFilters(input.ID)

	err = impl.repo.Read(ctx, filters, staff)
	if err != nil {
		return util.RepoReadErr(err)
	}

	err = impl.repo.Delete(ctx, filters)
	if err != nil {
		return util.RepoDeleteErr(err)
	}

	return nil
}
