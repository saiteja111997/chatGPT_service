package utilities

type intent struct {
	DisplayName string                 `json:"displayName"`
	Parameters  map[string]interface{} `json:"parameters"`
}

type queryResult struct {
	IntentInfo intent `json:"intentInfo"`
}

type IntentParameterValue struct {
	OriginalValue string `json:"originalValue"`
	ResolvedValue string `json:"resolvedValue"`
}

// webhookRequest is used to unmarshal a WebhookRequest JSON object. Note that
// not all members need to be defined--just those that you need to process.
// As an alternative, you could use the types provided by
// the Dialogflow protocol buffers:
// https://godoc.org/google.golang.org/genproto/googleapis/cloud/dialogflow/v2#WebhookRequest
type WebhookRequest struct {
	Session     string      `json:"session"`
	ResponseID  string      `json:"responseId"`
	QueryResult queryResult `json:"queryResult"`
}

// webhookResponse is used to marshal a WebhookResponse JSON object. Note that
// not all members need to be defined--just those that you need to process.
// As an alternative, you could use the types provided by
// the Dialogflow protocol buffers:
// https://godoc.org/google.golang.org/genproto/googleapis/cloud/dialogflow/v2#WebhookResponse

type TextObj struct {
	Text []string `json:"text"`
}

type MessageObject struct {
	Text TextObj `json:"text"`
}

type Message struct {
	Messages []MessageObject `json:"messages"`
}

type WebhookResponse struct {
	FulfillmentResponse Message `json:"fulfillmentResponse"`
}

// type SummaryUpdateMessageBody struct {
// 	Message string `json:"message"`
// }
