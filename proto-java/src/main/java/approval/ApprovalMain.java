package approval;

import example.dummy.ApprovalOuterClass;

import java.util.ArrayList;
import java.util.List;
import java.util.UUID;

public class ApprovalMain {

    public static void main(String[] args) {
        List<ApprovalOuterClass.Identity> testIdentities = new ArrayList<>();
        testIdentities.add(ApprovalOuterClass.Identity.newBuilder()
                .setType("IDENTITY")
                .setId(UUID.randomUUID().toString())
                .setSerialOrder(1)
                .build());
        testIdentities.add(ApprovalOuterClass.Identity.newBuilder()
                .setType("GOVERNANCE_GROUP")
                .setId(UUID.randomUUID().toString())
                .setSerialOrder(2)
                .build());

        ApprovalOuterClass.Approval testApproval = ApprovalOuterClass.Approval.newBuilder()
                .setId(UUID.randomUUID().toString())
                .setTenantId(UUID.randomUUID().toString())
                .setStatus(ApprovalOuterClass.Status.PENDING)
                .addAllApprovers(testIdentities)
                .build();

        System.out.println(testApproval);
    }
}
