package pushrules
type GetRequest_statusQueryParameterType int

const (
    COMPLETED_GETREQUEST_STATUSQUERYPARAMETERTYPE GetRequest_statusQueryParameterType = iota
    CANCELLED_GETREQUEST_STATUSQUERYPARAMETERTYPE
    EXPIRED_GETREQUEST_STATUSQUERYPARAMETERTYPE
    DENIED_GETREQUEST_STATUSQUERYPARAMETERTYPE
    OPEN_GETREQUEST_STATUSQUERYPARAMETERTYPE
    ALL_GETREQUEST_STATUSQUERYPARAMETERTYPE
)

func (i GetRequest_statusQueryParameterType) String() string {
    return []string{"completed", "cancelled", "expired", "denied", "open", "all"}[i]
}
func ParseGetRequest_statusQueryParameterType(v string) (any, error) {
    result := COMPLETED_GETREQUEST_STATUSQUERYPARAMETERTYPE
    switch v {
        case "completed":
            result = COMPLETED_GETREQUEST_STATUSQUERYPARAMETERTYPE
        case "cancelled":
            result = CANCELLED_GETREQUEST_STATUSQUERYPARAMETERTYPE
        case "expired":
            result = EXPIRED_GETREQUEST_STATUSQUERYPARAMETERTYPE
        case "denied":
            result = DENIED_GETREQUEST_STATUSQUERYPARAMETERTYPE
        case "open":
            result = OPEN_GETREQUEST_STATUSQUERYPARAMETERTYPE
        case "all":
            result = ALL_GETREQUEST_STATUSQUERYPARAMETERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGetRequest_statusQueryParameterType(values []GetRequest_statusQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetRequest_statusQueryParameterType) isMultiValue() bool {
    return false
}
