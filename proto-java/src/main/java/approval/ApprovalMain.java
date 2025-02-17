package approval;

import approval.request.ApprovalOuterClass;
import com.google.protobuf.InvalidProtocolBufferException;
import com.google.protobuf.Timestamp;

import java.time.OffsetDateTime;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.UUID;

public class ApprovalMain {

    public static void main(String[] args) throws InvalidProtocolBufferException {
        List<approval.request.ApprovalOuterClass.Identity> testIdentities = new ArrayList<>();
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

        OffsetDateTime currentTime = OffsetDateTime.now();

        ApprovalOuterClass.Approval testApproval = ApprovalOuterClass.Approval.newBuilder()
                .setId(UUID.randomUUID().toString())
                .setTenantId(UUID.randomUUID().toString())
                .setStatus(ApprovalOuterClass.Status.PENDING)
                .addAllApprovers(testIdentities)
                .setCreatedDate(Timestamp.newBuilder()
                        .setSeconds(currentTime.toInstant().getEpochSecond())
                        .setNanos(currentTime.toInstant().getNano())
                        .build())
                .build();

        System.out.println("Test approval object: " + testApproval);

        byte[] encodedApproval = testApproval.toByteArray();
        System.out.println("Encoded approval object: " + Arrays.toString(encodedApproval));

        ApprovalOuterClass.Approval decodedApproval = ApprovalOuterClass.Approval.parseFrom(encodedApproval);
        System.out.println("\nDecoded approval object: " + decodedApproval);
    }
}
