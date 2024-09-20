package models
type PatchSchema_schemas int

const (
    URNIETFPARAMSSCIMAPIMESSAGES20PATCHOP_PATCHSCHEMA_SCHEMAS PatchSchema_schemas = iota
)

func (i PatchSchema_schemas) String() string {
    return []string{"urn:ietf:params:scim:api:messages:2.0:PatchOp"}[i]
}
func ParsePatchSchema_schemas(v string) (any, error) {
    result := URNIETFPARAMSSCIMAPIMESSAGES20PATCHOP_PATCHSCHEMA_SCHEMAS
    switch v {
        case "urn:ietf:params:scim:api:messages:2.0:PatchOp":
            result = URNIETFPARAMSSCIMAPIMESSAGES20PATCHOP_PATCHSCHEMA_SCHEMAS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePatchSchema_schemas(values []PatchSchema_schemas) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PatchSchema_schemas) isMultiValue() bool {
    return false
}
