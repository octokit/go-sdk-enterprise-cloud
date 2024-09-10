package models
// The system role from which this role can inherit permissions.
type OrganizationCustomOrganizationRoleUpdateSchema_base_role int

const (
    NONE_ORGANIZATIONCUSTOMORGANIZATIONROLEUPDATESCHEMA_BASE_ROLE OrganizationCustomOrganizationRoleUpdateSchema_base_role = iota
    READ_ORGANIZATIONCUSTOMORGANIZATIONROLEUPDATESCHEMA_BASE_ROLE
    TRIAGE_ORGANIZATIONCUSTOMORGANIZATIONROLEUPDATESCHEMA_BASE_ROLE
    WRITE_ORGANIZATIONCUSTOMORGANIZATIONROLEUPDATESCHEMA_BASE_ROLE
    MAINTAIN_ORGANIZATIONCUSTOMORGANIZATIONROLEUPDATESCHEMA_BASE_ROLE
    ADMIN_ORGANIZATIONCUSTOMORGANIZATIONROLEUPDATESCHEMA_BASE_ROLE
)

func (i OrganizationCustomOrganizationRoleUpdateSchema_base_role) String() string {
    return []string{"none", "read", "triage", "write", "maintain", "admin"}[i]
}
func ParseOrganizationCustomOrganizationRoleUpdateSchema_base_role(v string) (any, error) {
    result := NONE_ORGANIZATIONCUSTOMORGANIZATIONROLEUPDATESCHEMA_BASE_ROLE
    switch v {
        case "none":
            result = NONE_ORGANIZATIONCUSTOMORGANIZATIONROLEUPDATESCHEMA_BASE_ROLE
        case "read":
            result = READ_ORGANIZATIONCUSTOMORGANIZATIONROLEUPDATESCHEMA_BASE_ROLE
        case "triage":
            result = TRIAGE_ORGANIZATIONCUSTOMORGANIZATIONROLEUPDATESCHEMA_BASE_ROLE
        case "write":
            result = WRITE_ORGANIZATIONCUSTOMORGANIZATIONROLEUPDATESCHEMA_BASE_ROLE
        case "maintain":
            result = MAINTAIN_ORGANIZATIONCUSTOMORGANIZATIONROLEUPDATESCHEMA_BASE_ROLE
        case "admin":
            result = ADMIN_ORGANIZATIONCUSTOMORGANIZATIONROLEUPDATESCHEMA_BASE_ROLE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeOrganizationCustomOrganizationRoleUpdateSchema_base_role(values []OrganizationCustomOrganizationRoleUpdateSchema_base_role) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OrganizationCustomOrganizationRoleUpdateSchema_base_role) isMultiValue() bool {
    return false
}
