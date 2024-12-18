package repos
// The visibility of the repository.
type ReposPostRequestBody_visibility int

const (
    PUBLIC_REPOSPOSTREQUESTBODY_VISIBILITY ReposPostRequestBody_visibility = iota
    PRIVATE_REPOSPOSTREQUESTBODY_VISIBILITY
    INTERNAL_REPOSPOSTREQUESTBODY_VISIBILITY
)

func (i ReposPostRequestBody_visibility) String() string {
    return []string{"public", "private", "internal"}[i]
}
func ParseReposPostRequestBody_visibility(v string) (any, error) {
    result := PUBLIC_REPOSPOSTREQUESTBODY_VISIBILITY
    switch v {
        case "public":
            result = PUBLIC_REPOSPOSTREQUESTBODY_VISIBILITY
        case "private":
            result = PRIVATE_REPOSPOSTREQUESTBODY_VISIBILITY
        case "internal":
            result = INTERNAL_REPOSPOSTREQUESTBODY_VISIBILITY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeReposPostRequestBody_visibility(values []ReposPostRequestBody_visibility) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ReposPostRequestBody_visibility) isMultiValue() bool {
    return false
}
