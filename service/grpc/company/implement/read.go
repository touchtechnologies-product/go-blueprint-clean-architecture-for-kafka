package implement

import (
	"context"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/grpc/company/out"
	pb "github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/grpc/company/protobuf"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/util"
)

func (impl *implementation) Read(ctx context.Context, input *pb.ReadCompanyRequest) (view *pb.ReadCompanyResponse, err error) {
	company := &domain.Company{}
	filters := makeCompanyIDFilters(input.CompanyId)
	err = impl.repo.Read(ctx, filters, company)
	if err != nil {
		return nil, util.RepoReadErr(err)
	}

	return out.SingleItemSingleItemGrpcView(company), nil
}
