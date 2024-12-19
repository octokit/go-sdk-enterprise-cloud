package models
// Authentication Type for Amazon S3.
type AmazonS3AccessKeysConfig_authentication_type int

const (
    ACCESS_KEYS_AMAZONS3ACCESSKEYSCONFIG_AUTHENTICATION_TYPE AmazonS3AccessKeysConfig_authentication_type = iota
)

func (i AmazonS3AccessKeysConfig_authentication_type) String() string {
    return []string{"access_keys"}[i]
}
func ParseAmazonS3AccessKeysConfig_authentication_type(v string) (any, error) {
    result := ACCESS_KEYS_AMAZONS3ACCESSKEYSCONFIG_AUTHENTICATION_TYPE
    switch v {
        case "access_keys":
            result = ACCESS_KEYS_AMAZONS3ACCESSKEYSCONFIG_AUTHENTICATION_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAmazonS3AccessKeysConfig_authentication_type(values []AmazonS3AccessKeysConfig_authentication_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AmazonS3AccessKeysConfig_authentication_type) isMultiValue() bool {
    return false
}
