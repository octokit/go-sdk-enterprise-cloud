package streams
// The audit log streaming provider. The name is case sensitive.
type StreamsPostRequestBody_stream_type int

const (
    AZUREBLOBSTORAGE_STREAMSPOSTREQUESTBODY_STREAM_TYPE StreamsPostRequestBody_stream_type = iota
    AZUREEVENTHUBS_STREAMSPOSTREQUESTBODY_STREAM_TYPE
    AMAZONS3_STREAMSPOSTREQUESTBODY_STREAM_TYPE
    SPLUNK_STREAMSPOSTREQUESTBODY_STREAM_TYPE
    HTTPSEVENTCOLLECTOR_STREAMSPOSTREQUESTBODY_STREAM_TYPE
    GOOGLECLOUDSTORAGE_STREAMSPOSTREQUESTBODY_STREAM_TYPE
    DATADOG_STREAMSPOSTREQUESTBODY_STREAM_TYPE
)

func (i StreamsPostRequestBody_stream_type) String() string {
    return []string{"Azure Blob Storage", "Azure Event Hubs", "Amazon S3", "Splunk", "HTTPS Event Collector", "Google Cloud Storage", "Datadog"}[i]
}
func ParseStreamsPostRequestBody_stream_type(v string) (any, error) {
    result := AZUREBLOBSTORAGE_STREAMSPOSTREQUESTBODY_STREAM_TYPE
    switch v {
        case "Azure Blob Storage":
            result = AZUREBLOBSTORAGE_STREAMSPOSTREQUESTBODY_STREAM_TYPE
        case "Azure Event Hubs":
            result = AZUREEVENTHUBS_STREAMSPOSTREQUESTBODY_STREAM_TYPE
        case "Amazon S3":
            result = AMAZONS3_STREAMSPOSTREQUESTBODY_STREAM_TYPE
        case "Splunk":
            result = SPLUNK_STREAMSPOSTREQUESTBODY_STREAM_TYPE
        case "HTTPS Event Collector":
            result = HTTPSEVENTCOLLECTOR_STREAMSPOSTREQUESTBODY_STREAM_TYPE
        case "Google Cloud Storage":
            result = GOOGLECLOUDSTORAGE_STREAMSPOSTREQUESTBODY_STREAM_TYPE
        case "Datadog":
            result = DATADOG_STREAMSPOSTREQUESTBODY_STREAM_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeStreamsPostRequestBody_stream_type(values []StreamsPostRequestBody_stream_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i StreamsPostRequestBody_stream_type) isMultiValue() bool {
    return false
}
