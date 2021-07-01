package staff

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	pb "github.com/touchtechnologies-product/go-blueprint-clean-architecture-for-kafka/service/grpc/staff/protobuf"
)

func (impl *implement) CreateStaff(c *gin.Context) {
	ctx := context.Background()
	client := impl.conn
	grpcStaffClient := pb.NewStaffGrpcServiceClient(client)

	input := &pb.CreateStaffRequest{}
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(400, gin.H{"success": false, "code": 400, "msg": err.Error(), "data": new(pb.CreateStaffResponse)})
		return
	}

	output, err := grpcStaffClient.Create(ctx, input)
	fmt.Println(output)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(500, gin.H{"success": false, "code": 500, "msg": err.Error(), "data": new(pb.CreateStaffResponse)})
		return
	}

	if output != nil {
		c.JSON(201, gin.H{"success": true, "code": 201, "msg": "Created", "data": output})
		return
	}

	c.JSON(500, gin.H{"success": true, "code": 500, "msg": "Something went wrong!", "data": new(pb.CreateStaffResponse)})
	return
}

func (impl *implement) ListStaff(c *gin.Context) {
	ctx := context.Background()
	client := impl.conn
	grpcStaffClient := pb.NewStaffGrpcServiceClient(client)

	input := &pb.ListStaffRequest{}
	if err := c.ShouldBindQuery(input); err != nil {
		c.JSON(400, gin.H{"success": false, "code": 400, "msg": err.Error(), "data": new(pb.ListStaffRequest)})
		return
	}

	output, err := grpcStaffClient.List(ctx, input)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "code": 500, "msg": err.Error(), "data": new(pb.ListStaffRequest)})
		return
	} else {

	}

	if output != nil {
		c.JSON(200, gin.H{"success": true, "code": 200, "msg": "Successfully", "data": output})
		return
	}

	c.JSON(200, gin.H{"success": true, "code": 404, "msg": "Data not found", "data": new(pb.ListStaffRequest)})
	return
}

func (impl *implement) ReadStaff(c *gin.Context) {
	ctx := context.Background()
	client := impl.conn
	grpcStaffClient := pb.NewStaffGrpcServiceClient(client)

	input := &pb.ReadStaffRequest{StaffId: c.Param("id")}
	output, err := grpcStaffClient.Read(ctx, input)
	if err != nil {
		c.JSON(200, gin.H{"success": true, "code": 404, "msg": "Data not found", "data": new(pb.UpdateStaffResponse)})
		return
	}

	c.JSON(200, gin.H{"success": true, "code": 200, "msg": "Successfully", "data": output})
	return
}

func (impl *implement) UpdateStaff(c *gin.Context) {
	ctx := context.Background()
	client := impl.conn
	grpcStaffClient := pb.NewStaffGrpcServiceClient(client)

	input := &pb.UpdateStaffRequest{Id: c.Param("id")}

	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(400, gin.H{"success": false, "code": 400, "msg": err.Error(), "data": new(pb.UpdateStaffResponse)})
		return
	}

	output, err := grpcStaffClient.Update(ctx, input)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "code": 500, "msg": err.Error(), "data": new(pb.UpdateStaffResponse)})
		return
	}

	c.JSON(200, gin.H{"success": true, "code": 200, "msg": "Updated", "data": output})
	return
}

func (impl *implement) DeleteStaff(c *gin.Context) {
	ctx := context.Background()
	client := impl.conn
	grpcStaffClient := pb.NewStaffGrpcServiceClient(client)

	input := &pb.DeleteStaffRequest{Id: c.Param("id")}

	_, err := grpcStaffClient.Delete(ctx, input)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "code": 500, "msg": err.Error(), "data": new(pb.DeleteStaffResponse)})
		return
	}

	c.JSON(200, gin.H{"success": true, "code": 200, "msg": "Deleted", "data": new(pb.DeleteStaffResponse)})
	return
}
