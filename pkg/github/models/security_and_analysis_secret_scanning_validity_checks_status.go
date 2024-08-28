package models
type SecurityAndAnalysis_secret_scanning_validity_checks_status int

const (
    ENABLED_SECURITYANDANALYSIS_SECRET_SCANNING_VALIDITY_CHECKS_STATUS SecurityAndAnalysis_secret_scanning_validity_checks_status = iota
    DISABLED_SECURITYANDANALYSIS_SECRET_SCANNING_VALIDITY_CHECKS_STATUS
)

func (i SecurityAndAnalysis_secret_scanning_validity_checks_status) String() string {
    return []string{"enabled", "disabled"}[i]
}
func ParseSecurityAndAnalysis_secret_scanning_validity_checks_status(v string) (any, error) {
    result := ENABLED_SECURITYANDANALYSIS_SECRET_SCANNING_VALIDITY_CHECKS_STATUS
    switch v {
        case "enabled":
            result = ENABLED_SECURITYANDANALYSIS_SECRET_SCANNING_VALIDITY_CHECKS_STATUS
        case "disabled":
            result = DISABLED_SECURITYANDANALYSIS_SECRET_SCANNING_VALIDITY_CHECKS_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSecurityAndAnalysis_secret_scanning_validity_checks_status(values []SecurityAndAnalysis_secret_scanning_validity_checks_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SecurityAndAnalysis_secret_scanning_validity_checks_status) isMultiValue() bool {
    return false
}
