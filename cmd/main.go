package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	proto_go "github.com/travis-wisch-sp/proto-test/proto-go"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	fmt.Println("Testing generated proto files")

	var testIdentities []*proto_go.Identity
	testIdentities = append(testIdentities, &proto_go.Identity{
		Id:          uuid.NewString(),
		Type:        "IDENTITY",
		SerialOrder: 1,
	})

	testApproval := &proto_go.ApprovalRequest{
		Id:          uuid.NewString(),
		TenantId:    uuid.NewString(),
		Approvers:   testIdentities,
		Status:      proto_go.Status_PENDING,
		CreatedDate: timestamppb.New(time.Now()),
	}
	fmt.Println("Test approval struct:")
	fmt.Println(testApproval)

	encodedApproval, err := proto.Marshal(testApproval)
	if err != nil {
		return
	}
	fmt.Println("Encoded approval struct:")
	fmt.Println(encodedApproval)

	decodedApproval := &proto_go.ApprovalRequest{}
	if err := proto.Unmarshal(encodedApproval, decodedApproval); err != nil {
		return
	}

	fmt.Println("Decoded approval struct:")
	fmt.Println(decodedApproval.String())
}
