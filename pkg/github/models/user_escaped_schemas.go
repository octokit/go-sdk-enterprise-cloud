package models
type User_schemas int

const (
    URNIETFPARAMSSCIMSCHEMASCORE20USER_USER_SCHEMAS User_schemas = iota
)

func (i User_schemas) String() string {
    return []string{"urn:ietf:params:scim:schemas:core:2.0:User"}[i]
}
func ParseUser_schemas(v string) (any, error) {
    result := URNIETFPARAMSSCIMSCHEMASCORE20USER_USER_SCHEMAS
    switch v {
        case "urn:ietf:params:scim:schemas:core:2.0:User":
            result = URNIETFPARAMSSCIMSCHEMASCORE20USER_USER_SCHEMAS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeUser_schemas(values []User_schemas) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i User_schemas) isMultiValue() bool {
    return false
}
