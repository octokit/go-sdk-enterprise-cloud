package models
// The system role from which this role can inherit permissions.
type OrganizationCustomOrganizationRoleCreateSchema_base_role int

const (
    READ_ORGANIZATIONCUSTOMORGANIZATIONROLECREATESCHEMA_BASE_ROLE OrganizationCustomOrganizationRoleCreateSchema_base_role = iota
    TRIAGE_ORGANIZATIONCUSTOMORGANIZATIONROLECREATESCHEMA_BASE_ROLE
    WRITE_ORGANIZATIONCUSTOMORGANIZATIONROLECREATESCHEMA_BASE_ROLE
    MAINTAIN_ORGANIZATIONCUSTOMORGANIZATIONROLECREATESCHEMA_BASE_ROLE
    ADMIN_ORGANIZATIONCUSTOMORGANIZATIONROLECREATESCHEMA_BASE_ROLE
)

func (i OrganizationCustomOrganizationRoleCreateSchema_base_role) String() string {
    return []string{"read", "triage", "write", "maintain", "admin"}[i]
}
func ParseOrganizationCustomOrganizationRoleCreateSchema_base_role(v string) (any, error) {
    result := READ_ORGANIZATIONCUSTOMORGANIZATIONROLECREATESCHEMA_BASE_ROLE
    switch v {
        case "read":
            result = READ_ORGANIZATIONCUSTOMORGANIZATIONROLECREATESCHEMA_BASE_ROLE
        case "triage":
            result = TRIAGE_ORGANIZATIONCUSTOMORGANIZATIONROLECREATESCHEMA_BASE_ROLE
        case "write":
            result = WRITE_ORGANIZATIONCUSTOMORGANIZATIONROLECREATESCHEMA_BASE_ROLE
        case "maintain":
            result = MAINTAIN_ORGANIZATIONCUSTOMORGANIZATIONROLECREATESCHEMA_BASE_ROLE
        case "admin":
            result = ADMIN_ORGANIZATIONCUSTOMORGANIZATIONROLECREATESCHEMA_BASE_ROLE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeOrganizationCustomOrganizationRoleCreateSchema_base_role(values []OrganizationCustomOrganizationRoleCreateSchema_base_role) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OrganizationCustomOrganizationRoleCreateSchema_base_role) isMultiValue() bool {
    return false
}
