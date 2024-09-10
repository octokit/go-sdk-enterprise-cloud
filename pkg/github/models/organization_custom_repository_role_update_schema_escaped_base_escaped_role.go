package models
// The system role from which this role inherits permissions.
type OrganizationCustomRepositoryRoleUpdateSchema_base_role int

const (
    READ_ORGANIZATIONCUSTOMREPOSITORYROLEUPDATESCHEMA_BASE_ROLE OrganizationCustomRepositoryRoleUpdateSchema_base_role = iota
    TRIAGE_ORGANIZATIONCUSTOMREPOSITORYROLEUPDATESCHEMA_BASE_ROLE
    WRITE_ORGANIZATIONCUSTOMREPOSITORYROLEUPDATESCHEMA_BASE_ROLE
    MAINTAIN_ORGANIZATIONCUSTOMREPOSITORYROLEUPDATESCHEMA_BASE_ROLE
)

func (i OrganizationCustomRepositoryRoleUpdateSchema_base_role) String() string {
    return []string{"read", "triage", "write", "maintain"}[i]
}
func ParseOrganizationCustomRepositoryRoleUpdateSchema_base_role(v string) (any, error) {
    result := READ_ORGANIZATIONCUSTOMREPOSITORYROLEUPDATESCHEMA_BASE_ROLE
    switch v {
        case "read":
            result = READ_ORGANIZATIONCUSTOMREPOSITORYROLEUPDATESCHEMA_BASE_ROLE
        case "triage":
            result = TRIAGE_ORGANIZATIONCUSTOMREPOSITORYROLEUPDATESCHEMA_BASE_ROLE
        case "write":
            result = WRITE_ORGANIZATIONCUSTOMREPOSITORYROLEUPDATESCHEMA_BASE_ROLE
        case "maintain":
            result = MAINTAIN_ORGANIZATIONCUSTOMREPOSITORYROLEUPDATESCHEMA_BASE_ROLE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeOrganizationCustomRepositoryRoleUpdateSchema_base_role(values []OrganizationCustomRepositoryRoleUpdateSchema_base_role) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OrganizationCustomRepositoryRoleUpdateSchema_base_role) isMultiValue() bool {
    return false
}
