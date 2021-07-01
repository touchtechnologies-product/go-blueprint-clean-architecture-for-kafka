package implement

import (
	"context"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/domain"
	pb "github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/grpc/company/protobuf"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/util"
)

func (impl *implementation) Delete(ctx context.Context, input *pb.DeleteCompanyRequest) (item *pb.DeleteCompanyResponse, err error) {
	company := &domain.Company{}
	filters := makeCompanyIDFilters(input.Id)

	err = impl.repo.Read(ctx, filters, company)
	if err != nil {
		return nil, util.RepoReadErr(err)
	}

	err = impl.repo.Delete(ctx, filters)
	if err != nil {
		return nil, util.RepoDeleteErr(err)
	}

	return new(pb.DeleteCompanyResponse), err
}
