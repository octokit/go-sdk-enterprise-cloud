package models
// The system role from which this role inherits permissions.
type OrganizationCustomRepositoryRoleCreateSchema_base_role int

const (
    READ_ORGANIZATIONCUSTOMREPOSITORYROLECREATESCHEMA_BASE_ROLE OrganizationCustomRepositoryRoleCreateSchema_base_role = iota
    TRIAGE_ORGANIZATIONCUSTOMREPOSITORYROLECREATESCHEMA_BASE_ROLE
    WRITE_ORGANIZATIONCUSTOMREPOSITORYROLECREATESCHEMA_BASE_ROLE
    MAINTAIN_ORGANIZATIONCUSTOMREPOSITORYROLECREATESCHEMA_BASE_ROLE
)

func (i OrganizationCustomRepositoryRoleCreateSchema_base_role) String() string {
    return []string{"read", "triage", "write", "maintain"}[i]
}
func ParseOrganizationCustomRepositoryRoleCreateSchema_base_role(v string) (any, error) {
    result := READ_ORGANIZATIONCUSTOMREPOSITORYROLECREATESCHEMA_BASE_ROLE
    switch v {
        case "read":
            result = READ_ORGANIZATIONCUSTOMREPOSITORYROLECREATESCHEMA_BASE_ROLE
        case "triage":
            result = TRIAGE_ORGANIZATIONCUSTOMREPOSITORYROLECREATESCHEMA_BASE_ROLE
        case "write":
            result = WRITE_ORGANIZATIONCUSTOMREPOSITORYROLECREATESCHEMA_BASE_ROLE
        case "maintain":
            result = MAINTAIN_ORGANIZATIONCUSTOMREPOSITORYROLECREATESCHEMA_BASE_ROLE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeOrganizationCustomRepositoryRoleCreateSchema_base_role(values []OrganizationCustomRepositoryRoleCreateSchema_base_role) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OrganizationCustomRepositoryRoleCreateSchema_base_role) isMultiValue() bool {
    return false
}
