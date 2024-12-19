package models
// The response status to the bypass request until dismissed.
type BypassResponse_status int

const (
    APPROVED_BYPASSRESPONSE_STATUS BypassResponse_status = iota
    DENIED_BYPASSRESPONSE_STATUS
    DISMISSED_BYPASSRESPONSE_STATUS
)

func (i BypassResponse_status) String() string {
    return []string{"approved", "denied", "dismissed"}[i]
}
func ParseBypassResponse_status(v string) (any, error) {
    result := APPROVED_BYPASSRESPONSE_STATUS
    switch v {
        case "approved":
            result = APPROVED_BYPASSRESPONSE_STATUS
        case "denied":
            result = DENIED_BYPASSRESPONSE_STATUS
        case "dismissed":
            result = DISMISSED_BYPASSRESPONSE_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeBypassResponse_status(values []BypassResponse_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i BypassResponse_status) isMultiValue() bool {
    return false
}
