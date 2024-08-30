package models
// The policy that controls the organizations in the enterprise that are allowed to run GitHub Actions.
type EnabledOrganizations int

const (
    ALL_ENABLEDORGANIZATIONS EnabledOrganizations = iota
    NONE_ENABLEDORGANIZATIONS
    SELECTED_ENABLEDORGANIZATIONS
)

func (i EnabledOrganizations) String() string {
    return []string{"all", "none", "selected"}[i]
}
func ParseEnabledOrganizations(v string) (any, error) {
    result := ALL_ENABLEDORGANIZATIONS
    switch v {
        case "all":
            result = ALL_ENABLEDORGANIZATIONS
        case "none":
            result = NONE_ENABLEDORGANIZATIONS
        case "selected":
            result = SELECTED_ENABLEDORGANIZATIONS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeEnabledOrganizations(values []EnabledOrganizations) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i EnabledOrganizations) isMultiValue() bool {
    return false
}
