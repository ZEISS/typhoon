package cloudevents

// ReplierOption modifies a newly created replier.
type ReplierOption func(*Replier) error

// ReplierWithPayloadPolicy option avoids returning events.
func ReplierWithPayloadPolicy(policy PayloadPolicy) ReplierOption {
	return func(c *Replier) error {
		c.payloadPolicy = policy
		return nil
	}
}

// ReplierWithStaticResponseType option uses a static string for response type.
func ReplierWithStaticResponseType(resType string) ReplierOption {
	return func(c *Replier) error {
		c.responseType = StaticResponse(resType)
		return nil
	}
}

// ReplierWithMappedResponseType option uses a map string to look up response type.
func ReplierWithMappedResponseType(resTypes map[string]string) ReplierOption {
	return func(c *Replier) error {
		c.responseType = MappedResponseType(resTypes)
		return nil
	}
}

// ReplierWithStaticErrorResponseType option uses a static string for error response type.
func ReplierWithStaticErrorResponseType(resType string) ReplierOption {
	return func(c *Replier) error {
		c.responseErrorType = StaticResponse(resType)
		return nil
	}
}

// ReplierWithMappedErrorResponseType option uses a map string to look up for error response type.
func ReplierWithMappedErrorResponseType(resTypes map[string]string) ReplierOption {
	return func(c *Replier) error {
		c.responseErrorType = MappedResponseType(resTypes)
		return nil
	}
}

// ReplierWithStatefulHeaders adds response option to create stateful headers if not present.
func ReplierWithStatefulHeaders(bridge string) ReplierOption {
	return func(c *Replier) error {
		if bridge != "" {
			c.responseOptions = append(c.responseOptions, ResponseWithStatefulHeaders(bridge))
		}
		return nil
	}
}

// ReplierWithProcessedHeaders adds response option to create processed headers.
func ReplierWithProcessedHeaders() ReplierOption {
	return func(c *Replier) error {
		c.responseOptions = append(c.responseOptions, ResponseWithProcessedHeaders())
		return nil
	}
}

// ReplierWithStaticDataContentType sets the response content type for all replies.
func ReplierWithStaticDataContentType(contentType string) ReplierOption {
	return func(c *Replier) error {
		c.responseContentType = StaticResponse(contentType)
		return nil
	}
}

// ReplierWithStaticErrorDataContentType sets the response content type for error replies.
func ReplierWithStaticErrorDataContentType(contentType string) ReplierOption {
	return func(c *Replier) error {
		c.responseErrorContentType = StaticResponse(contentType)
		return nil
	}
}
