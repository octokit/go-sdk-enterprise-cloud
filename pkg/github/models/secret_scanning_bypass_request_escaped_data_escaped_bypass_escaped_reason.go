package models
// The reason the bypass was requested.
type SecretScanningBypassRequest_data_bypass_reason int

const (
    USED_IN_TESTS_SECRETSCANNINGBYPASSREQUEST_DATA_BYPASS_REASON SecretScanningBypassRequest_data_bypass_reason = iota
    FALSE_POSITIVE_SECRETSCANNINGBYPASSREQUEST_DATA_BYPASS_REASON
    FIX_LATER_SECRETSCANNINGBYPASSREQUEST_DATA_BYPASS_REASON
)

func (i SecretScanningBypassRequest_data_bypass_reason) String() string {
    return []string{"used_in_tests", "false_positive", "fix_later"}[i]
}
func ParseSecretScanningBypassRequest_data_bypass_reason(v string) (any, error) {
    result := USED_IN_TESTS_SECRETSCANNINGBYPASSREQUEST_DATA_BYPASS_REASON
    switch v {
        case "used_in_tests":
            result = USED_IN_TESTS_SECRETSCANNINGBYPASSREQUEST_DATA_BYPASS_REASON
        case "false_positive":
            result = FALSE_POSITIVE_SECRETSCANNINGBYPASSREQUEST_DATA_BYPASS_REASON
        case "fix_later":
            result = FIX_LATER_SECRETSCANNINGBYPASSREQUEST_DATA_BYPASS_REASON
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSecretScanningBypassRequest_data_bypass_reason(values []SecretScanningBypassRequest_data_bypass_reason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SecretScanningBypassRequest_data_bypass_reason) isMultiValue() bool {
    return false
}
