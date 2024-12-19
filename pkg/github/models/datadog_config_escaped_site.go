package models
// Datadog Site to use.
type DatadogConfig_site int

const (
    US_DATADOGCONFIG_SITE DatadogConfig_site = iota
    US3_DATADOGCONFIG_SITE
    US5_DATADOGCONFIG_SITE
    EU1_DATADOGCONFIG_SITE
    US1FED_DATADOGCONFIG_SITE
    AP1_DATADOGCONFIG_SITE
)

func (i DatadogConfig_site) String() string {
    return []string{"US", "US3", "US5", "EU1", "US1-FED", "AP1"}[i]
}
func ParseDatadogConfig_site(v string) (any, error) {
    result := US_DATADOGCONFIG_SITE
    switch v {
        case "US":
            result = US_DATADOGCONFIG_SITE
        case "US3":
            result = US3_DATADOGCONFIG_SITE
        case "US5":
            result = US5_DATADOGCONFIG_SITE
        case "EU1":
            result = EU1_DATADOGCONFIG_SITE
        case "US1-FED":
            result = US1FED_DATADOGCONFIG_SITE
        case "AP1":
            result = AP1_DATADOGCONFIG_SITE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDatadogConfig_site(values []DatadogConfig_site) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DatadogConfig_site) isMultiValue() bool {
    return false
}
