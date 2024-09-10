package models
type ScimEnterpriseGroupList_schemas int

const (
    URNIETFPARAMSSCIMAPIMESSAGES20LISTRESPONSE_SCIMENTERPRISEGROUPLIST_SCHEMAS ScimEnterpriseGroupList_schemas = iota
)

func (i ScimEnterpriseGroupList_schemas) String() string {
    return []string{"urn:ietf:params:scim:api:messages:2.0:ListResponse"}[i]
}
func ParseScimEnterpriseGroupList_schemas(v string) (any, error) {
    result := URNIETFPARAMSSCIMAPIMESSAGES20LISTRESPONSE_SCIMENTERPRISEGROUPLIST_SCHEMAS
    switch v {
        case "urn:ietf:params:scim:api:messages:2.0:ListResponse":
            result = URNIETFPARAMSSCIMAPIMESSAGES20LISTRESPONSE_SCIMENTERPRISEGROUPLIST_SCHEMAS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeScimEnterpriseGroupList_schemas(values []ScimEnterpriseGroupList_schemas) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ScimEnterpriseGroupList_schemas) isMultiValue() bool {
    return false
}
