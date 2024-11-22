package models
// Authentication Type for Amazon S3.
type AmazonS3OidcConfig_authentication_type int

const (
    OIDC_AMAZONS3OIDCCONFIG_AUTHENTICATION_TYPE AmazonS3OidcConfig_authentication_type = iota
)

func (i AmazonS3OidcConfig_authentication_type) String() string {
    return []string{"oidc"}[i]
}
func ParseAmazonS3OidcConfig_authentication_type(v string) (any, error) {
    result := OIDC_AMAZONS3OIDCCONFIG_AUTHENTICATION_TYPE
    switch v {
        case "oidc":
            result = OIDC_AMAZONS3OIDCCONFIG_AUTHENTICATION_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAmazonS3OidcConfig_authentication_type(values []AmazonS3OidcConfig_authentication_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AmazonS3OidcConfig_authentication_type) isMultiValue() bool {
    return false
}
