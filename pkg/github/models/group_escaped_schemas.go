package models
type Group_schemas int

const (
    URNIETFPARAMSSCIMSCHEMASCORE20GROUP_GROUP_SCHEMAS Group_schemas = iota
)

func (i Group_schemas) String() string {
    return []string{"urn:ietf:params:scim:schemas:core:2.0:Group"}[i]
}
func ParseGroup_schemas(v string) (any, error) {
    result := URNIETFPARAMSSCIMSCHEMASCORE20GROUP_GROUP_SCHEMAS
    switch v {
        case "urn:ietf:params:scim:schemas:core:2.0:Group":
            result = URNIETFPARAMSSCIMSCHEMASCORE20GROUP_GROUP_SCHEMAS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGroup_schemas(values []Group_schemas) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Group_schemas) isMultiValue() bool {
    return false
}
