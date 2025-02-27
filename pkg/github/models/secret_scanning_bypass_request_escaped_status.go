package models
// The status of the bypass request.
type SecretScanningBypassRequest_status int

const (
    PENDING_SECRETSCANNINGBYPASSREQUEST_STATUS SecretScanningBypassRequest_status = iota
    DENIED_SECRETSCANNINGBYPASSREQUEST_STATUS
    APPROVED_SECRETSCANNINGBYPASSREQUEST_STATUS
    CANCELLED_SECRETSCANNINGBYPASSREQUEST_STATUS
    COMPLETED_SECRETSCANNINGBYPASSREQUEST_STATUS
    EXPIRED_SECRETSCANNINGBYPASSREQUEST_STATUS
    OPEN_SECRETSCANNINGBYPASSREQUEST_STATUS
)

func (i SecretScanningBypassRequest_status) String() string {
    return []string{"pending", "denied", "approved", "cancelled", "completed", "expired", "open"}[i]
}
func ParseSecretScanningBypassRequest_status(v string) (any, error) {
    result := PENDING_SECRETSCANNINGBYPASSREQUEST_STATUS
    switch v {
        case "pending":
            result = PENDING_SECRETSCANNINGBYPASSREQUEST_STATUS
        case "denied":
            result = DENIED_SECRETSCANNINGBYPASSREQUEST_STATUS
        case "approved":
            result = APPROVED_SECRETSCANNINGBYPASSREQUEST_STATUS
        case "cancelled":
            result = CANCELLED_SECRETSCANNINGBYPASSREQUEST_STATUS
        case "completed":
            result = COMPLETED_SECRETSCANNINGBYPASSREQUEST_STATUS
        case "expired":
            result = EXPIRED_SECRETSCANNINGBYPASSREQUEST_STATUS
        case "open":
            result = OPEN_SECRETSCANNINGBYPASSREQUEST_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSecretScanningBypassRequest_status(values []SecretScanningBypassRequest_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SecretScanningBypassRequest_status) isMultiValue() bool {
    return false
}
