package domain

import "time"

// ApprovalRequest has all supported fields expected in an approval request the main bread and butter of this service
type ApprovalRequest struct {
	// Unique ID to distinguish the approval request
	ApprovalID string `json:"id"`
	// Unique ID signifying org/tenant ownership of approval request
	TenantID  string     `json:"tenantId"`
	Approvers []Identity `json:"approvers"`
	// Date/time of approval request creation
	CreatedDate  time.Time    `json:"createdDate"`
	BatchRequest BatchRequest `json:"batchRequest"`
	// Date/time that an approval request completion is expected
	DueDate *time.Time `json:"dueDate"`
	// Signifies the approval request type and what event type to use when publishing the completion APPROVAL_V1 event
	Type string `json:"type"`
	// Unique name to ideally identify what the approval request is for
	Name []LocaleField `json:"name"`
	// Approval request description
	Description []LocaleField `json:"description"`
	// Signifies the priority level of the approval request
	// enum: ["LOW","MEDIUM","HIGH"]
	Priority string `json:"priority"`
	// Signifies what medium to use when sending notifications (currently only email is utilized)
	// enum: ["EMAIL","SLACK","TEAMS"]
	Medium []string `json:"medium,omitempty"`
	// The owner/originator of the approval request
	Requester Identity `json:"requester"`
	// Who the approval request was submitted for (on the behalf of)
	Requestee Identity `json:"requestee,omitempty"`
	// Slice of requester/approver made comments on the approval request
	Comments []Comment `json:"comments,omitempty"`
	// Determines how many of the listed approvers are required to consider the approval request "APPROVED" or "REJECTED"
	ApprovalCriteria ApprovalCriteria `json:"approvalCriteria"`
	// A slice of approvers who have approved the approval request
	ApprovedBy []Identity `json:"approvedBy"`
	// A slice of approvers who have rejected the approval request (should never contain more than 1 Identity)
	RejectedBy []Identity `json:"rejectedBy"`
	// Date of finalization of the approval request (acquired all necessary approvals)
	CompletedDate *time.Time `json:"completedDate"`
	// Approval request status
	// enum: ["PENDING","APPROVED","REJECTED", "EXPIRED"]
	Status string `json:"status"`
	// Field used to link an approval request back to a source/role/machine account/entitlement/etc. that the approval request was created for
	ReferenceData []Reference `json:"referenceData"`
	// Field that can include any additional info that may be needed by the service that the approval request originated from
	AdditionalAttributes map[string]any `json:"additionalAttributes"`
	// Signifies whether the approval request can be auto approved if the requester is listed as the sole approver
	// (DIRECT), is one of multiple listed approvers (INDIRECT), or not at all (OFF)
	// enum: ["OFF","DIRECT","INDIRECT"]
	AutoApprove string `json:"autoApprove"` // instantly auto approves if the approvers contains the requester
	// A map of approval configuration for the approval in question.
	ApprovalConfig *ApprovalConfig `json:"approvalConfig,omitempty"`
	// Indicates what tier/set of approvers an approval request is on if the approval criteria is set to "SERIAL"
	SerialStep int `json:"serialStep"`
	// Indicates what tier/set of configured escalation approvers an approval request has been escalated to
	EscalationStep int `json:"escalationStep"`
}

// ApprovalCriteria contains configurations around how the approval request will be determined to be "APPROVED" or "REJECTED"
type ApprovalCriteria struct {
	// SERIAL Uses the serialChain defined in the approval configurations to one by one assign to each approver until the approval criteria is satisfied
	// PARALLEL specifies that all approvers will be sent the approval at the same time
	Type      string              `json:"type"` // possible values: SERIAL, PARALLEL
	Rejection ApprovalCalculation `json:"rejection"`
	Approval  ApprovalCalculation `json:"approval"`
}

// ApprovalCalculation determines how "APPROVED" or "REJECTED" will be calculated
type ApprovalCalculation struct {
	// calculationType: this defines what the field "value" will be used as, either a count or percentage of the total approvers that need to approve or reject
	CalculationType string `json:"calculationType"` // possible values: COUNT, PERCENT
	Value           int    `json:"value"`           // how many or what percent of approvals/rejections are required to complete the approval request
}

// Comment holds approval request comment details
type Comment struct {
	CommentID   string    `json:"commentID,omitempty"`
	Author      Identity  `json:"author"`
	Comment     string    `json:"comment"`
	CreatedDate time.Time `json:"createdDate,omitempty"`
}

// Identity approval request identity that it is assigned to or requested by etc.
// Type can only specify IDENTITY, GOVERNANCE_GROUP, SOURCE_OWNER for now, so identityID should probably get changed to ID
type Identity struct {
	IdentityID  string      `json:"identityID,omitempty"`
	Type        string      `json:"type,omitempty"`
	Name        string      `json:"name,omitempty"`
	OwnerOf     []Reference `json:"ownerOf,omitempty"`
	SerialOrder int         `json:"serialOrder,omitempty"`
}

// BatchRequest used to keep track of how many approval requests are in a batch to only notify approvers once for the entire batch
// required field that defaults to a random UUID and BatchSize of 1
type BatchRequest struct {
	BatchID   string `json:"batchId,omitempty"`
	BatchSize int    `json:"batchSize,omitempty"`
}

// Reference used to link other IDs to the approval request
// this should be a required field so that the requesting service knows what the approval request was sent for
// these references can also be searched on through the UI
type Reference struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
}

// LocaleField for fields that may have more than one language needed to be displayed in the UI
// Locale defaults to en_US if not specified
type LocaleField struct {
	Value  string `json:"value,omitempty"`
	Locale string `json:"locale,omitempty"`
}
