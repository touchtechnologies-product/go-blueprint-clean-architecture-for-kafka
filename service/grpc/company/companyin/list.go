package companyin

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/domain"
	pb "github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/grpc/company/protobuf"
	"strconv"
)

func CreateInputPageOptGRPCToPageOtpDomain(opt *pb.ListCompanyRequest) *domain.PageOption {
	page, err := strconv.Atoi(opt.Page)
	if err != nil {
		page = 1
	}

	perPage, err := strconv.Atoi(opt.PerPage)
	if err != nil {
		perPage = 20
	}

	return &domain.PageOption{
		Page:    page,
		PerPage: perPage,
		Filters: opt.Filters,
		Sorts:   opt.Sorts,
	}
}

func MakeTestCreatePageOptGRPCInput() *pb.ListCompanyRequest {
	return &pb.ListCompanyRequest{
		Page:    "1",
		PerPage: "20",
		Filters: []string{"id:eq:test", "name:eq:test"},
		Sorts:   []string{"id:desc"},
	}
}

func MakeTestCreatePageOptGRPCInputToPageOtpDomain(opt *pb.ListCompanyRequest) *domain.PageOption {
	page, err := strconv.Atoi(opt.Page)
	if err != nil {
		page = 1
	}

	perPage, err := strconv.Atoi(opt.PerPage)
	if err != nil {
		perPage = 20
	}

	return &domain.PageOption{
		Page:    page,
		PerPage: perPage,
		Filters: opt.Filters,
		Sorts:   opt.Sorts,
	}
}

func MakeTestRecordCompanyDomain() (records []interface{}) {
	records = []interface{}{
		&domain.Company{Name: "test1", ID: "test1"},
		&domain.Company{Name: "test2", ID: "test2"}}
	return
}
