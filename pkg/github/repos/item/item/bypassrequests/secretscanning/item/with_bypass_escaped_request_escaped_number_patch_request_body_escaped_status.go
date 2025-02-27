package item
// The review action to perform on the bypass request.
type WithBypass_request_numberPatchRequestBody_status int

const (
    APPROVE_WITHBYPASS_REQUEST_NUMBERPATCHREQUESTBODY_STATUS WithBypass_request_numberPatchRequestBody_status = iota
    REJECT_WITHBYPASS_REQUEST_NUMBERPATCHREQUESTBODY_STATUS
)

func (i WithBypass_request_numberPatchRequestBody_status) String() string {
    return []string{"approve", "reject"}[i]
}
func ParseWithBypass_request_numberPatchRequestBody_status(v string) (any, error) {
    result := APPROVE_WITHBYPASS_REQUEST_NUMBERPATCHREQUESTBODY_STATUS
    switch v {
        case "approve":
            result = APPROVE_WITHBYPASS_REQUEST_NUMBERPATCHREQUESTBODY_STATUS
        case "reject":
            result = REJECT_WITHBYPASS_REQUEST_NUMBERPATCHREQUESTBODY_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithBypass_request_numberPatchRequestBody_status(values []WithBypass_request_numberPatchRequestBody_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithBypass_request_numberPatchRequestBody_status) isMultiValue() bool {
    return false
}
