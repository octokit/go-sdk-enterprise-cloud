package models
type GroupResponse_schemas int

const (
    URNIETFPARAMSSCIMSCHEMASCORE20GROUP_GROUPRESPONSE_SCHEMAS GroupResponse_schemas = iota
    URNIETFPARAMSSCIMAPIMESSAGES20LISTRESPONSE_GROUPRESPONSE_SCHEMAS
)

func (i GroupResponse_schemas) String() string {
    return []string{"urn:ietf:params:scim:schemas:core:2.0:Group", "urn:ietf:params:scim:api:messages:2.0:ListResponse"}[i]
}
func ParseGroupResponse_schemas(v string) (any, error) {
    result := URNIETFPARAMSSCIMSCHEMASCORE20GROUP_GROUPRESPONSE_SCHEMAS
    switch v {
        case "urn:ietf:params:scim:schemas:core:2.0:Group":
            result = URNIETFPARAMSSCIMSCHEMASCORE20GROUP_GROUPRESPONSE_SCHEMAS
        case "urn:ietf:params:scim:api:messages:2.0:ListResponse":
            result = URNIETFPARAMSSCIMAPIMESSAGES20LISTRESPONSE_GROUPRESPONSE_SCHEMAS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGroupResponse_schemas(values []GroupResponse_schemas) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GroupResponse_schemas) isMultiValue() bool {
    return false
}
