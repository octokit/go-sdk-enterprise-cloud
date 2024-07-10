package item
import (
    "errors"
)
// Visibility of a runner group. You can select all organizations or select individual organizations.
type WithRunner_group_PatchRequestBody_visibility int

const (
    SELECTED_WITHRUNNER_GROUP_PATCHREQUESTBODY_VISIBILITY WithRunner_group_PatchRequestBody_visibility = iota
    ALL_WITHRUNNER_GROUP_PATCHREQUESTBODY_VISIBILITY
)

func (i WithRunner_group_PatchRequestBody_visibility) String() string {
    return []string{"selected", "all"}[i]
}
func ParseWithRunner_group_PatchRequestBody_visibility(v string) (any, error) {
    result := SELECTED_WITHRUNNER_GROUP_PATCHREQUESTBODY_VISIBILITY
    switch v {
        case "selected":
            result = SELECTED_WITHRUNNER_GROUP_PATCHREQUESTBODY_VISIBILITY
        case "all":
            result = ALL_WITHRUNNER_GROUP_PATCHREQUESTBODY_VISIBILITY
        default:
            return 0, errors.New("Unknown WithRunner_group_PatchRequestBody_visibility value: " + v)
    }
    return &result, nil
}
func SerializeWithRunner_group_PatchRequestBody_visibility(values []WithRunner_group_PatchRequestBody_visibility) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithRunner_group_PatchRequestBody_visibility) isMultiValue() bool {
    return false
}
