package implement

import (
	"context"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/domain"
	pb "github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/grpc/staff/protobuf"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/staff/out"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/staff/staffin"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/util"
)

func (impl *implementation) List(ctx context.Context, optGrpc *pb.ListStaffRequest) (items *pb.ListStaffResponse, err error) {

	opt := staffin.CreateInputPageOptGrpcToPageOtpDomain(optGrpc)
	total, records, err := impl.repo.List(ctx, opt, &domain.Staff{})
	if err != nil {
		return nil, util.RepoListErr(err)
	}

	output := make([]*pb.ListStaffResponse_Output, len(records))
	for i, record := range records {
		output[i] = out.ListItemToListItemGrpcView(record.(*domain.Staff))
	}

	return out.ListStaffToListStaffGrpcView(total, output), nil
}
