package models
type UserResponse_schemas int

const (
    URNIETFPARAMSSCIMSCHEMASCORE20USER_USERRESPONSE_SCHEMAS UserResponse_schemas = iota
)

func (i UserResponse_schemas) String() string {
    return []string{"urn:ietf:params:scim:schemas:core:2.0:User"}[i]
}
func ParseUserResponse_schemas(v string) (any, error) {
    result := URNIETFPARAMSSCIMSCHEMASCORE20USER_USERRESPONSE_SCHEMAS
    switch v {
        case "urn:ietf:params:scim:schemas:core:2.0:User":
            result = URNIETFPARAMSSCIMSCHEMASCORE20USER_USERRESPONSE_SCHEMAS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeUserResponse_schemas(values []UserResponse_schemas) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i UserResponse_schemas) isMultiValue() bool {
    return false
}
