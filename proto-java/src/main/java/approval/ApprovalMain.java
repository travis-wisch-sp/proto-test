package approval;

import approval.request.Approval;
import approval.request.Approval.ApprovalRequest;
import approval.request.Approval.Identity;

import com.google.protobuf.InvalidProtocolBufferException;
import com.google.protobuf.Timestamp;

import java.time.OffsetDateTime;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.UUID;

public class ApprovalMain {

    public static void main(String[] args) throws InvalidProtocolBufferException {
        List<approval.request.Approval.Identity> testIdentities = new ArrayList<>();
        testIdentities.add(approval.request.Approval.Identity.newBuilder()
                .setType("IDENTITY")
                .setIdentityId(UUID.randomUUID().toString())
                .setSerialOrder(1)
                .build());
        testIdentities.add(Identity.newBuilder()
                .setType("GOVERNANCE_GROUP")
                .setIdentityId(UUID.randomUUID().toString())
                .setSerialOrder(2)
                .build());

        OffsetDateTime currentTime = OffsetDateTime.now();

        ApprovalRequest testApproval = ApprovalRequest.newBuilder()
                .setId(UUID.randomUUID().toString())
                .setTenantId(UUID.randomUUID().toString())
                .setStatus(Approval.Status.PENDING)
                .addAllApprovers(testIdentities)
                .setCreatedDate(Timestamp.newBuilder()
                        .setSeconds(currentTime.toInstant().getEpochSecond())
                        .setNanos(currentTime.toInstant().getNano())
                        .build())
                .build();

        System.out.println("Test approval object: " + testApproval);

        byte[] encodedApproval = testApproval.toByteArray();
        System.out.println("Encoded approval object: " + Arrays.toString(encodedApproval));

        ApprovalRequest decodedApproval = ApprovalRequest.parseFrom(encodedApproval);
        System.out.println("\nDecoded approval object: " + decodedApproval);
    }
}
