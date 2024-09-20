package models
type ScimUser_operations_op int

const (
    ADD_SCIMUSER_OPERATIONS_OP ScimUser_operations_op = iota
    REMOVE_SCIMUSER_OPERATIONS_OP
    REPLACE_SCIMUSER_OPERATIONS_OP
)

func (i ScimUser_operations_op) String() string {
    return []string{"add", "remove", "replace"}[i]
}
func ParseScimUser_operations_op(v string) (any, error) {
    result := ADD_SCIMUSER_OPERATIONS_OP
    switch v {
        case "add":
            result = ADD_SCIMUSER_OPERATIONS_OP
        case "remove":
            result = REMOVE_SCIMUSER_OPERATIONS_OP
        case "replace":
            result = REPLACE_SCIMUSER_OPERATIONS_OP
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeScimUser_operations_op(values []ScimUser_operations_op) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ScimUser_operations_op) isMultiValue() bool {
    return false
}
