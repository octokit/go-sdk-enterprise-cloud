package models
// The system role from which this role inherits permissions.
type OrganizationCustomRepositoryRole_base_role int

const (
    READ_ORGANIZATIONCUSTOMREPOSITORYROLE_BASE_ROLE OrganizationCustomRepositoryRole_base_role = iota
    TRIAGE_ORGANIZATIONCUSTOMREPOSITORYROLE_BASE_ROLE
    WRITE_ORGANIZATIONCUSTOMREPOSITORYROLE_BASE_ROLE
    MAINTAIN_ORGANIZATIONCUSTOMREPOSITORYROLE_BASE_ROLE
)

func (i OrganizationCustomRepositoryRole_base_role) String() string {
    return []string{"read", "triage", "write", "maintain"}[i]
}
func ParseOrganizationCustomRepositoryRole_base_role(v string) (any, error) {
    result := READ_ORGANIZATIONCUSTOMREPOSITORYROLE_BASE_ROLE
    switch v {
        case "read":
            result = READ_ORGANIZATIONCUSTOMREPOSITORYROLE_BASE_ROLE
        case "triage":
            result = TRIAGE_ORGANIZATIONCUSTOMREPOSITORYROLE_BASE_ROLE
        case "write":
            result = WRITE_ORGANIZATIONCUSTOMREPOSITORYROLE_BASE_ROLE
        case "maintain":
            result = MAINTAIN_ORGANIZATIONCUSTOMREPOSITORYROLE_BASE_ROLE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeOrganizationCustomRepositoryRole_base_role(values []OrganizationCustomRepositoryRole_base_role) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OrganizationCustomRepositoryRole_base_role) isMultiValue() bool {
    return false
}
