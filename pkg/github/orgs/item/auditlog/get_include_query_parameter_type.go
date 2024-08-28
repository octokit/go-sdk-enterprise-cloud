package auditlog
type GetIncludeQueryParameterType int

const (
    WEB_GETINCLUDEQUERYPARAMETERTYPE GetIncludeQueryParameterType = iota
    GIT_GETINCLUDEQUERYPARAMETERTYPE
    ALL_GETINCLUDEQUERYPARAMETERTYPE
)

func (i GetIncludeQueryParameterType) String() string {
    return []string{"web", "git", "all"}[i]
}
func ParseGetIncludeQueryParameterType(v string) (any, error) {
    result := WEB_GETINCLUDEQUERYPARAMETERTYPE
    switch v {
        case "web":
            result = WEB_GETINCLUDEQUERYPARAMETERTYPE
        case "git":
            result = GIT_GETINCLUDEQUERYPARAMETERTYPE
        case "all":
            result = ALL_GETINCLUDEQUERYPARAMETERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGetIncludeQueryParameterType(values []GetIncludeQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetIncludeQueryParameterType) isMultiValue() bool {
    return false
}
