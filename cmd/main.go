package main

import (
	"fmt"

	"github.com/google/uuid"
	proto_go "github.com/travis-wisch-sp/proto-test/proto-go"
)

func main() {
	fmt.Println("Testing generated proto files")

	var testIdentities []*proto_go.Identity
	testIdentities = append(testIdentities, &proto_go.Identity{
		Id:          uuid.NewString(),
		Type:        "IDENTITY",
		SerialOrder: 1,
	})

	testApproval := proto_go.Approval{
		Id:        uuid.NewString(),
		TenantId:  uuid.NewString(),
		Approvers: testIdentities,
		Status:    proto_go.Status_PENDING,
	}

	fmt.Println("Test approval struct:")
	fmt.Println(testApproval)
}
