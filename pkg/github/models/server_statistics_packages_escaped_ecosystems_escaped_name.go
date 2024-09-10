package models
// The name of the package ecosystem
type ServerStatisticsPackages_ecosystems_name int

const (
    NPM_SERVERSTATISTICSPACKAGES_ECOSYSTEMS_NAME ServerStatisticsPackages_ecosystems_name = iota
    MAVEN_SERVERSTATISTICSPACKAGES_ECOSYSTEMS_NAME
    DOCKER_SERVERSTATISTICSPACKAGES_ECOSYSTEMS_NAME
    NUGET_SERVERSTATISTICSPACKAGES_ECOSYSTEMS_NAME
    RUBYGEMS_SERVERSTATISTICSPACKAGES_ECOSYSTEMS_NAME
    CONTAINERS_SERVERSTATISTICSPACKAGES_ECOSYSTEMS_NAME
)

func (i ServerStatisticsPackages_ecosystems_name) String() string {
    return []string{"npm", "maven", "docker", "nuget", "rubygems", "containers"}[i]
}
func ParseServerStatisticsPackages_ecosystems_name(v string) (any, error) {
    result := NPM_SERVERSTATISTICSPACKAGES_ECOSYSTEMS_NAME
    switch v {
        case "npm":
            result = NPM_SERVERSTATISTICSPACKAGES_ECOSYSTEMS_NAME
        case "maven":
            result = MAVEN_SERVERSTATISTICSPACKAGES_ECOSYSTEMS_NAME
        case "docker":
            result = DOCKER_SERVERSTATISTICSPACKAGES_ECOSYSTEMS_NAME
        case "nuget":
            result = NUGET_SERVERSTATISTICSPACKAGES_ECOSYSTEMS_NAME
        case "rubygems":
            result = RUBYGEMS_SERVERSTATISTICSPACKAGES_ECOSYSTEMS_NAME
        case "containers":
            result = CONTAINERS_SERVERSTATISTICSPACKAGES_ECOSYSTEMS_NAME
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeServerStatisticsPackages_ecosystems_name(values []ServerStatisticsPackages_ecosystems_name) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ServerStatisticsPackages_ecosystems_name) isMultiValue() bool {
    return false
}
