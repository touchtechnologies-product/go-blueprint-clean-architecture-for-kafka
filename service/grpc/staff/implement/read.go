package implement

import (
	"context"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/domain"
	pb "github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/grpc/staff/protobuf"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/staff/out"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/util"
)

func (impl *implementation) Read(ctx context.Context, input *pb.ReadStaffRequest) (view *pb.ReadStaffResponse, err error) {
	staff := &domain.Staff{}
	filters := makeStaffIDFilters(input.StaffId)
	err = impl.repo.Read(ctx, filters, staff)
	if err != nil {
		return nil, util.RepoReadErr(err)
	}

	view = out.SingleItemSingleItemGrpcView(staff)
	return
}
