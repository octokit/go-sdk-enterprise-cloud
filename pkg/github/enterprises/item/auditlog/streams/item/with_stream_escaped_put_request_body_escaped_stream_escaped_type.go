package item
// The audit log streaming provider. The name is case sensitive.
type WithStream_PutRequestBody_stream_type int

const (
    AZUREBLOBSTORAGE_WITHSTREAM_PUTREQUESTBODY_STREAM_TYPE WithStream_PutRequestBody_stream_type = iota
    AZUREEVENTHUBS_WITHSTREAM_PUTREQUESTBODY_STREAM_TYPE
    AMAZONS3_WITHSTREAM_PUTREQUESTBODY_STREAM_TYPE
    SPLUNK_WITHSTREAM_PUTREQUESTBODY_STREAM_TYPE
    HTTPSEVENTCOLLECTOR_WITHSTREAM_PUTREQUESTBODY_STREAM_TYPE
    GOOGLECLOUDSTORAGE_WITHSTREAM_PUTREQUESTBODY_STREAM_TYPE
    DATADOG_WITHSTREAM_PUTREQUESTBODY_STREAM_TYPE
)

func (i WithStream_PutRequestBody_stream_type) String() string {
    return []string{"Azure Blob Storage", "Azure Event Hubs", "Amazon S3", "Splunk", "HTTPS Event Collector", "Google Cloud Storage", "Datadog"}[i]
}
func ParseWithStream_PutRequestBody_stream_type(v string) (any, error) {
    result := AZUREBLOBSTORAGE_WITHSTREAM_PUTREQUESTBODY_STREAM_TYPE
    switch v {
        case "Azure Blob Storage":
            result = AZUREBLOBSTORAGE_WITHSTREAM_PUTREQUESTBODY_STREAM_TYPE
        case "Azure Event Hubs":
            result = AZUREEVENTHUBS_WITHSTREAM_PUTREQUESTBODY_STREAM_TYPE
        case "Amazon S3":
            result = AMAZONS3_WITHSTREAM_PUTREQUESTBODY_STREAM_TYPE
        case "Splunk":
            result = SPLUNK_WITHSTREAM_PUTREQUESTBODY_STREAM_TYPE
        case "HTTPS Event Collector":
            result = HTTPSEVENTCOLLECTOR_WITHSTREAM_PUTREQUESTBODY_STREAM_TYPE
        case "Google Cloud Storage":
            result = GOOGLECLOUDSTORAGE_WITHSTREAM_PUTREQUESTBODY_STREAM_TYPE
        case "Datadog":
            result = DATADOG_WITHSTREAM_PUTREQUESTBODY_STREAM_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithStream_PutRequestBody_stream_type(values []WithStream_PutRequestBody_stream_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithStream_PutRequestBody_stream_type) isMultiValue() bool {
    return false
}
