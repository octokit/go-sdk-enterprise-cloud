package models
// Shows if a package system is enabled, disabled, or read-only in a GHES installation
type ServerStatisticsPackages_ecosystems_enabled int

const (
    TRUE_SERVERSTATISTICSPACKAGES_ECOSYSTEMS_ENABLED ServerStatisticsPackages_ecosystems_enabled = iota
    FALSE_SERVERSTATISTICSPACKAGES_ECOSYSTEMS_ENABLED
    READONLY_SERVERSTATISTICSPACKAGES_ECOSYSTEMS_ENABLED
)

func (i ServerStatisticsPackages_ecosystems_enabled) String() string {
    return []string{"TRUE", "FALSE", "READONLY"}[i]
}
func ParseServerStatisticsPackages_ecosystems_enabled(v string) (any, error) {
    result := TRUE_SERVERSTATISTICSPACKAGES_ECOSYSTEMS_ENABLED
    switch v {
        case "TRUE":
            result = TRUE_SERVERSTATISTICSPACKAGES_ECOSYSTEMS_ENABLED
        case "FALSE":
            result = FALSE_SERVERSTATISTICSPACKAGES_ECOSYSTEMS_ENABLED
        case "READONLY":
            result = READONLY_SERVERSTATISTICSPACKAGES_ECOSYSTEMS_ENABLED
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeServerStatisticsPackages_ecosystems_enabled(values []ServerStatisticsPackages_ecosystems_enabled) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ServerStatisticsPackages_ecosystems_enabled) isMultiValue() bool {
    return false
}
