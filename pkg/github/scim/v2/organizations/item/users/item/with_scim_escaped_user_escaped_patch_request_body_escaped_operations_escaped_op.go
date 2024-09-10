package item
type WithScim_user_PatchRequestBody_Operations_op int

const (
    ADD_WITHSCIM_USER_PATCHREQUESTBODY_OPERATIONS_OP WithScim_user_PatchRequestBody_Operations_op = iota
    REMOVE_WITHSCIM_USER_PATCHREQUESTBODY_OPERATIONS_OP
    REPLACE_WITHSCIM_USER_PATCHREQUESTBODY_OPERATIONS_OP
)

func (i WithScim_user_PatchRequestBody_Operations_op) String() string {
    return []string{"add", "remove", "replace"}[i]
}
func ParseWithScim_user_PatchRequestBody_Operations_op(v string) (any, error) {
    result := ADD_WITHSCIM_USER_PATCHREQUESTBODY_OPERATIONS_OP
    switch v {
        case "add":
            result = ADD_WITHSCIM_USER_PATCHREQUESTBODY_OPERATIONS_OP
        case "remove":
            result = REMOVE_WITHSCIM_USER_PATCHREQUESTBODY_OPERATIONS_OP
        case "replace":
            result = REPLACE_WITHSCIM_USER_PATCHREQUESTBODY_OPERATIONS_OP
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithScim_user_PatchRequestBody_Operations_op(values []WithScim_user_PatchRequestBody_Operations_op) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithScim_user_PatchRequestBody_Operations_op) isMultiValue() bool {
    return false
}
