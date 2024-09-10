package models
type ScimEnterpriseUserList_schemas int

const (
    URNIETFPARAMSSCIMAPIMESSAGES20LISTRESPONSE_SCIMENTERPRISEUSERLIST_SCHEMAS ScimEnterpriseUserList_schemas = iota
)

func (i ScimEnterpriseUserList_schemas) String() string {
    return []string{"urn:ietf:params:scim:api:messages:2.0:ListResponse"}[i]
}
func ParseScimEnterpriseUserList_schemas(v string) (any, error) {
    result := URNIETFPARAMSSCIMAPIMESSAGES20LISTRESPONSE_SCIMENTERPRISEUSERLIST_SCHEMAS
    switch v {
        case "urn:ietf:params:scim:api:messages:2.0:ListResponse":
            result = URNIETFPARAMSSCIMAPIMESSAGES20LISTRESPONSE_SCIMENTERPRISEUSERLIST_SCHEMAS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeScimEnterpriseUserList_schemas(values []ScimEnterpriseUserList_schemas) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ScimEnterpriseUserList_schemas) isMultiValue() bool {
    return false
}
