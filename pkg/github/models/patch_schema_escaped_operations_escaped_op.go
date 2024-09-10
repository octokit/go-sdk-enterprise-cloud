package models
type PatchSchema_Operations_op int

const (
    ADD_PATCHSCHEMA_OPERATIONS_OP PatchSchema_Operations_op = iota
    REPLACE_PATCHSCHEMA_OPERATIONS_OP
    REMOVE_PATCHSCHEMA_OPERATIONS_OP
)

func (i PatchSchema_Operations_op) String() string {
    return []string{"add", "replace", "remove"}[i]
}
func ParsePatchSchema_Operations_op(v string) (any, error) {
    result := ADD_PATCHSCHEMA_OPERATIONS_OP
    switch v {
        case "add":
            result = ADD_PATCHSCHEMA_OPERATIONS_OP
        case "replace":
            result = REPLACE_PATCHSCHEMA_OPERATIONS_OP
        case "remove":
            result = REMOVE_PATCHSCHEMA_OPERATIONS_OP
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePatchSchema_Operations_op(values []PatchSchema_Operations_op) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PatchSchema_Operations_op) isMultiValue() bool {
    return false
}
