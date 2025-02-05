package domain

// ApprovalConfig contains the configurations for an approval request
type ApprovalConfig struct {
	// Unique ID signifying org/tenant ownership of approval request
	TenantID string `json:"tenantId"`
	// the ID defined by the scope field, this could the approval ID (uuid), specific domain object ID (uuid), approval type (role/source/access_request/entitlement, tenant ID (uuid)
	ID string `json:"id"`
	// the type/scope of the above ID: "APPROVAL_REQUEST", "DOMAIN_OBJECT", "APPROVAL_TYPE", "TENANT"
	Scope string `json:"scope"`
	// configurations around when and how often to send reminders
	ReminderConfig ReminderConfig `json:"reminderConfig"`
	// configurations around when and how often to escalate
	EscalationConfig EscalationConfig `json:"escalationConfig"`
	// configurations around when the approval request should expire
	TimeoutConfig TimeoutConfig `json:"timeoutConfig"`
	// cron timezone settings for all above cron expressions
	CronTimezone CronTimezone `json:"cronTimezone"`
	// if the approval request has an approvalCriteria of SERIAL this chain will be used to determine the assignment order
	SerialChain []ChainTier `json:"serialChain"`
	// determines whether a comment is required when approving or rejecting the approval request
	RequireComment string `json:"requiresComment"`
	// used as a fallback in the case of a failure to escalate
	FallbackApprover Identity `json:"fallbackApprover"`
}

// ReminderConfig contains configurations around when and how often to send reminders
type ReminderConfig struct {
	Enabled                bool `json:"enabled"`
	DaysUntilFirstReminder int  `json:"daysUntilFirstReminder"`
	// cron expression determining when approvers will start getting reminder emails starting after DaysUntilFirstReminder
	ReminderCronSchedule string `json:"reminderCronSchedule"`
	// the max amount of reminder emails to send out until it stops. 0 means unlimited
	MaxReminders int `json:"maxReminders"`
}

// EscalationConfig contains configurations around when and how often to escalate
type EscalationConfig struct {
	Enabled                  bool `json:"enabled"`
	DaysUntilFirstEscalation int  `json:"daysUntilFirstEscalation"`
	// cron expression determining when the approval request will get escalated to the first person in the EscalationChain starting after DaysUntilFirstEscalation
	EscalationCronSchedule string      `json:"escalationCronSchedule"`
	EscalationChain        []ChainTier `json:"escalationChain"`
}

// ChainTier if the approval request has an approvalCriteria.type of SERIAL this chain will be used to determine the assignment order
type ChainTier struct {
	ChainID string `json:"chainId,omitempty"`
	// starting at 1 defines the order in which the identities will get assigned
	Tier int `json:"tier"`
	// the ID the corresponds to the IdentityType
	IdentityID string `json:"identityId"`
	// IdentityType: "IDENTITY", "MANAGER_OF", "GOVERNANCE_GROUP", "SOURCE_OWNER", "ROLE_OWNER", "ACCESS_PROFILE_OWNER", "ENTITLEMENT_OWNER"
	IdentityType string `json:"identityType"`
}

// TimeoutConfig contains configurations around when the approval request should expire
type TimeoutConfig struct {
	Enabled          bool   `json:"enabled"`
	DaysUntilTimeout int    `json:"daysUntilTimeout"`
	TimeoutResult    string `json:"timeoutResult,omitempty"`
}

// CronTimezone cron timezone settings for all above cron expressions
type CronTimezone struct {
	// the cron location
	Location string `json:"location"`
	// the cron offset
	Offset string `json:"offset"`
}
