package models
// A type of a resource
type Meta_resourceType int

const (
    USER_META_RESOURCETYPE Meta_resourceType = iota
    GROUP_META_RESOURCETYPE
)

func (i Meta_resourceType) String() string {
    return []string{"User", "Group"}[i]
}
func ParseMeta_resourceType(v string) (any, error) {
    result := USER_META_RESOURCETYPE
    switch v {
        case "User":
            result = USER_META_RESOURCETYPE
        case "Group":
            result = GROUP_META_RESOURCETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeMeta_resourceType(values []Meta_resourceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Meta_resourceType) isMultiValue() bool {
    return false
}
