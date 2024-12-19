package models
// The status of the bypass request.
type PushRuleBypassRequest_status int

const (
    PENDING_PUSHRULEBYPASSREQUEST_STATUS PushRuleBypassRequest_status = iota
    DENIED_PUSHRULEBYPASSREQUEST_STATUS
    APPROVED_PUSHRULEBYPASSREQUEST_STATUS
    CANCELLED_PUSHRULEBYPASSREQUEST_STATUS
    COMPLETED_PUSHRULEBYPASSREQUEST_STATUS
    EXPIRED_PUSHRULEBYPASSREQUEST_STATUS
    OPEN_PUSHRULEBYPASSREQUEST_STATUS
)

func (i PushRuleBypassRequest_status) String() string {
    return []string{"pending", "denied", "approved", "cancelled", "completed", "expired", "open"}[i]
}
func ParsePushRuleBypassRequest_status(v string) (any, error) {
    result := PENDING_PUSHRULEBYPASSREQUEST_STATUS
    switch v {
        case "pending":
            result = PENDING_PUSHRULEBYPASSREQUEST_STATUS
        case "denied":
            result = DENIED_PUSHRULEBYPASSREQUEST_STATUS
        case "approved":
            result = APPROVED_PUSHRULEBYPASSREQUEST_STATUS
        case "cancelled":
            result = CANCELLED_PUSHRULEBYPASSREQUEST_STATUS
        case "completed":
            result = COMPLETED_PUSHRULEBYPASSREQUEST_STATUS
        case "expired":
            result = EXPIRED_PUSHRULEBYPASSREQUEST_STATUS
        case "open":
            result = OPEN_PUSHRULEBYPASSREQUEST_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePushRuleBypassRequest_status(values []PushRuleBypassRequest_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PushRuleBypassRequest_status) isMultiValue() bool {
    return false
}
