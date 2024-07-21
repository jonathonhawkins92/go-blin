package responses

import "net/http"

// Continue returns a Response with status 100 Continue.
// The server has received the request headers and the client should proceed to send the request body.
func (response *Response) Http100Continue() *Response {
	response.status = http.StatusContinue
	return response
}

// Continue returns a Response with status 100 Continue.
// The server has received the request headers and the client should proceed to send the request body.
func (response *Response) Continue() *Response {
	return response.Http100Continue()
}

// Continue returns a Response with status 100 Continue.
// The server has received the request headers and the client should proceed to send the request body.
func (response *Response) Http100() *Response {
	return response.Http100Continue()
}

// SwitchingProtocols returns a Response with status 101 Switching Protocols.
// The server is switching protocols according to the Upgrade header sent by the client.
func (response *Response) Http101SwitchingProtocols() *Response {
	response.status = http.StatusSwitchingProtocols
	return response
}

// SwitchingProtocols returns a Response with status 101 Switching Protocols.
// The server is switching protocols according to the Upgrade header sent by the client.
func (response *Response) SwitchingProtocols() *Response {
	return response.Http101SwitchingProtocols()
}

// SwitchingProtocols returns a Response with status 101 Switching Protocols.
// The server is switching protocols according to the Upgrade header sent by the client.
func (response *Response) Http101() *Response {
	return response.Http101SwitchingProtocols()
}

// Processing returns a Response with status 102 Processing.
// The server has received and is processing the request, but no response is available yet.
func (response *Response) Http102Processing() *Response {
	response.status = http.StatusProcessing
	return response
}

// Processing returns a Response with status 102 Processing.
// The server has received and is processing the request, but no response is available yet.
func (response *Response) Processing() *Response {
	return response.Http102Processing()
}

// Processing returns a Response with status 102 Processing.
// The server has received and is processing the request, but no response is available yet.
func (response *Response) Http102() *Response {
	return response.Http102Processing()
}

// EarlyHints returns a Response with status 103 Early Hints.
// Used to return some response headers before final HTTP message.
func (response *Response) Http103EarlyHints() *Response {
	response.status = http.StatusEarlyHints
	return response
}

// EarlyHints returns a Response with status 103 Early Hints.
// Used to return some response headers before final HTTP message.
func (response *Response) EarlyHints() *Response {
	return response.Http103EarlyHints()
}

// EarlyHints returns a Response with status 103 Early Hints.
// Used to return some response headers before final HTTP message.
func (response *Response) Http103() *Response {
	return response.Http103EarlyHints()
}

// OK returns a Response with status 200 OK.
// The request has succeeded.
func (response *Response) Http200OK() *Response {
	response.status = http.StatusOK
	return response
}

// OK returns a Response with status 200 OK.
// The request has succeeded.
func (response *Response) OK() *Response {
	return response.Http200OK()
}

// OK returns a Response with status 200 OK.
// The request has succeeded.
func (response *Response) Http200() *Response {
	return response.Http200OK()
}

// Created returns a Response with status 201 Created.
// The request has been fulfilled and a new resource has been created.
func (response *Response) Http201Created() *Response {
	response.status = http.StatusCreated
	return response
}

// Created returns a Response with status 201 Created.
// The request has been fulfilled and a new resource has been created.
func (response *Response) Created() *Response {
	return response.Http201Created()
}

// Created returns a Response with status 201 Created.
// The request has been fulfilled and a new resource has been created.
func (response *Response) Http201() *Response {
	return response.Http201Created()
}

// Accepted returns a Response with status 202 Accepted.
// The request has been accepted for processing, but the processing has not been completed.
func (response *Response) Http202Accepted() *Response {
	response.status = http.StatusAccepted
	return response
}

// Accepted returns a Response with status 202 Accepted.
// The request has been accepted for processing, but the processing has not been completed.
func (response *Response) Accepted() *Response {
	return response.Http202Accepted()
}

// Accepted returns a Response with status 202 Accepted.
// The request has been accepted for processing, but the processing has not been completed.
func (response *Response) Http202() *Response {
	return response.Http202Accepted()
}

// NonAuthoritativeInfo returns a Response with status 203 Non-Authoritative Information.
// The server is a transforming proxy that received a 200 OK from its origin but is returning a modified version of the origin's response.
func (response *Response) Http203NonAuthoritativeInfo() *Response {
	response.status = http.StatusNonAuthoritativeInfo
	return response
}

// NonAuthoritativeInfo returns a Response with status 203 Non-Authoritative Information.
// The server is a transforming proxy that received a 200 OK from its origin but is returning a modified version of the origin's response.
func (response *Response) NonAuthoritativeInfo() *Response {
	return response.Http203NonAuthoritativeInfo()
}

// NonAuthoritativeInfo returns a Response with status 203 Non-Authoritative Information.
// The server is a transforming proxy that received a 200 OK from its origin but is returning a modified version of the origin's response.
func (response *Response) Http203() *Response {
	return response.Http203NonAuthoritativeInfo()
}

// NoContent returns a Response with status 204 No Content.
// The server successfully processed the request and is not returning any content.
func (response *Response) Http204NoContent() *Response {
	response.status = http.StatusNoContent
	return response
}

// NoContent returns a Response with status 204 No Content.
// The server successfully processed the request and is not returning any content.
func (response *Response) NoContent() *Response {
	return response.Http204NoContent()
}

// NoContent returns a Response with status 204 No Content.
// The server successfully processed the request and is not returning any content.
func (response *Response) Http204() *Response {
	return response.Http204NoContent()
}

// ResetContent returns a Response with status 205 Reset Content.
// The server successfully processed the request, but is not returning any content. The client should reset the document view.
func (response *Response) Http205ResetContent() *Response {
	response.status = http.StatusResetContent
	return response
}

// ResetContent returns a Response with status 205 Reset Content.
// The server successfully processed the request, but is not returning any content. The client should reset the document view.
func (response *Response) ResetContent() *Response {
	return response.Http205ResetContent()
}

// ResetContent returns a Response with status 205 Reset Content.
// The server successfully processed the request, but is not returning any content. The client should reset the document view.
func (response *Response) Http205() *Response {
	return response.Http205ResetContent()
}

// PartialContent returns a Response with status 206 Partial Content.
// The server is delivering only part of the resource due to a range header sent by the client.
func (response *Response) Http206PartialContent() *Response {
	response.status = http.StatusPartialContent
	return response
}

// PartialContent returns a Response with status 206 Partial Content.
// The server is delivering only part of the resource due to a range header sent by the client.
func (response *Response) PartialContent() *Response {
	return response.Http206PartialContent()
}

// PartialContent returns a Response with status 206 Partial Content.
// The server is delivering only part of the resource due to a range header sent by the client.
func (response *Response) Http206() *Response {
	return response.Http206PartialContent()
}

// MultiStatus returns a Response with status 207 Multi-Status.
// The message body that follows is an XML message and can contain a number of separate response codes.
func (response *Response) Http207MultiStatus() *Response {
	response.status = http.StatusMultiStatus
	return response
}

// MultiStatus returns a Response with status 207 Multi-Status.
// The message body that follows is an XML message and can contain a number of separate response codes.
func (response *Response) MultiStatus() *Response {
	return response.Http207MultiStatus()
}

// MultiStatus returns a Response with status 207 Multi-Status.
// The message body that follows is an XML message and can contain a number of separate response codes.
func (response *Response) Http207() *Response {
	return response.Http207MultiStatus()
}

// AlreadyReported returns a Response with status 208 Already Reported.
// The members of a DAV binding have already been enumerated in a previous reply to this request, and are not being included again.
func (response *Response) Http208AlreadyReported() *Response {
	response.status = http.StatusAlreadyReported
	return response
}

// AlreadyReported returns a Response with status 208 Already Reported.
// The members of a DAV binding have already been enumerated in a previous reply to this request, and are not being included again.
func (response *Response) AlreadyReported() *Response {
	return response.Http208AlreadyReported()
}

// AlreadyReported returns a Response with status 208 Already Reported.
// The members of a DAV binding have already been enumerated in a previous reply to this request, and are not being included again.
func (response *Response) Http208() *Response {
	return response.Http208AlreadyReported()
}

// IMUsed returns a Response with status 226 IM Used.
// The server has fulfilled a request for the resource, and the response is a representation of the result of one or more instance-manipulations applied to the current instance.
func (response *Response) Http226IMUsed() *Response {
	response.status = http.StatusIMUsed
	return response
}

// IMUsed returns a Response with status 226 IM Used.
// The server has fulfilled a request for the resource, and the response is a representation of the result of one or more instance-manipulations applied to the current instance.
func (response *Response) IMUsed() *Response {
	return response.Http226IMUsed()
}

// IMUsed returns a Response with status 226 IM Used.
// The server has fulfilled a request for the resource, and the response is a representation of the result of one or more instance-manipulations applied to the current instance.
func (response *Response) Http226() *Response {
	return response.Http226IMUsed()
}

// MultipleChoices returns a Response with status 300 Multiple Choices.
// The requested resource corresponds to any one of a set of representations, each with its own specific location.
func (response *Response) Http300MultipleChoices() *Response {
	response.status = http.StatusMultipleChoices
	return response
}

// MultipleChoices returns a Response with status 300 Multiple Choices.
// The requested resource corresponds to any one of a set of representations, each with its own specific location.
func (response *Response) MultipleChoices() *Response {
	return response.Http300MultipleChoices()
}

// MultipleChoices returns a Response with status 300 Multiple Choices.
// The requested resource corresponds to any one of a set of representations, each with its own specific location.
func (response *Response) Http300() *Response {
	return response.Http300MultipleChoices()
}

// MovedPermanently returns a Response with status 301 Moved Permanently.
// The requested resource has been assigned a new permanent URI and any future references to this resource should use one of the returned URIs.
func (response *Response) Http301MovedPermanently() *Response {
	response.status = http.StatusMovedPermanently
	return response
}

// MovedPermanently returns a Response with status 301 Moved Permanently.
// The requested resource has been assigned a new permanent URI and any future references to this resource should use one of the returned URIs.
func (response *Response) MovedPermanently() *Response {
	return response.Http301MovedPermanently()
}

// MovedPermanently returns a Response with status 301 Moved Permanently.
// The requested resource has been assigned a new permanent URI and any future references to this resource should use one of the returned URIs.
func (response *Response) Http301() *Response {
	return response.Http301MovedPermanently()
}

// Found returns a Response with status 302 Found.
// The requested resource resides temporarily under a different URI.
func (response *Response) Http302Found() *Response {
	response.status = http.StatusFound
	return response
}

// Found returns a Response with status 302 Found.
// The requested resource resides temporarily under a different URI.
func (response *Response) Found() *Response {
	return response.Http302Found()
}

// Found returns a Response with status 302 Found.
// The requested resource resides temporarily under a different URI.
func (response *Response) Http302() *Response {
	return response.Http302Found()
}

// SeeOther returns a Response with status 303 See Other.
// The response to the request can be found under a different URI and should be retrieved using a GET method on that resource.
func (response *Response) Http303SeeOther() *Response {
	response.status = http.StatusSeeOther
	return response
}

// SeeOther returns a Response with status 303 See Other.
// The response to the request can be found under a different URI and should be retrieved using a GET method on that resource.
func (response *Response) SeeOther() *Response {
	return response.Http303SeeOther()
}

// SeeOther returns a Response with status 303 See Other.
// The response to the request can be found under a different URI and should be retrieved using a GET method on that resource.
func (response *Response) Http303() *Response {
	return response.Http303SeeOther()
}

// NotModified returns a Response with status 304 Not Modified.
// Indicates that the resource has not been modified since the version specified by the request headers If-Modified-Since or If-None-Match.
func (response *Response) Http304NotModified() *Response {
	response.status = http.StatusNotModified
	return response
}

// NotModified returns a Response with status 304 Not Modified.
// Indicates that the resource has not been modified since the version specified by the request headers If-Modified-Since or If-None-Match.
func (response *Response) NotModified() *Response {
	return response.Http304NotModified()
}

// NotModified returns a Response with status 304 Not Modified.
// Indicates that the resource has not been modified since the version specified by the request headers If-Modified-Since or If-None-Match.
func (response *Response) Http304() *Response {
	return response.Http304NotModified()
}

// UseProxy returns a Response with status 305 Use Proxy.
// The requested resource must be accessed through the proxy given by the Location field.
func (response *Response) Http305UseProxy() *Response {
	response.status = http.StatusUseProxy
	return response
}

// UseProxy returns a Response with status 305 Use Proxy.
// The requested resource must be accessed through the proxy given by the Location field.
func (response *Response) UseProxy() *Response {
	return response.Http305UseProxy()
}

// UseProxy returns a Response with status 305 Use Proxy.
// The requested resource must be accessed through the proxy given by the Location field.
func (response *Response) Http305() *Response {
	return response.Http305UseProxy()
}

// TemporaryRedirect returns a Response with status 307 Temporary Redirect.
// The requested resource resides temporarily under a different URI.
func (response *Response) Http307TemporaryRedirect() *Response {
	response.status = http.StatusTemporaryRedirect
	return response
}

// TemporaryRedirect returns a Response with status 307 Temporary Redirect.
// The requested resource resides temporarily under a different URI.
func (response *Response) TemporaryRedirect() *Response {
	return response.Http307TemporaryRedirect()
}

// TemporaryRedirect returns a Response with status 307 Temporary Redirect.
// The requested resource resides temporarily under a different URI.
func (response *Response) Http307() *Response {
	return response.Http307TemporaryRedirect()
}

// PermanentRedirect returns a Response with status 308 Permanent Redirect.
// The requested resource has been assigned a new permanent URI and any future references to this resource should use one of the returned URIs.
func (response *Response) Http308PermanentRedirect() *Response {
	response.status = http.StatusPermanentRedirect
	return response
}

// PermanentRedirect returns a Response with status 308 Permanent Redirect.
// The requested resource has been assigned a new permanent URI and any future references to this resource should use one of the returned URIs.
func (response *Response) PermanentRedirect() *Response {
	return response.Http308PermanentRedirect()
}

// PermanentRedirect returns a Response with status 308 Permanent Redirect.
// The requested resource has been assigned a new permanent URI and any future references to this resource should use one of the returned URIs.
func (response *Response) Http308() *Response {
	return response.Http308PermanentRedirect()
}

// BadRequest returns a Response with status 400 Bad Request.
// The server cannot or will not process the request due to an apparent client error.
func (response *Response) Http400BadRequest() *Response {
	response.status = http.StatusBadRequest
	return response
}

// BadRequest returns a Response with status 400 Bad Request.
// The server cannot or will not process the request due to an apparent client error.
func (response *Response) BadRequest() *Response {
	return response.Http400BadRequest()
}

// BadRequest returns a Response with status 400 Bad Request.
// The server cannot or will not process the request due to an apparent client error.
func (response *Response) Http400() *Response {
	return response.Http400BadRequest()
}

// Unauthorized returns a Response with status 401 Unauthorized.
// Similar to 403 Forbidden, but specifically for use when authentication is required and has failed or has not yet been provided.
func (response *Response) Http401Unauthorized() *Response {
	response.status = http.StatusUnauthorized
	return response
}

// Unauthorized returns a Response with status 401 Unauthorized.
// Similar to 403 Forbidden, but specifically for use when authentication is required and has failed or has not yet been provided.
func (response *Response) Unauthorized() *Response {
	return response.Http401Unauthorized()
}

// Unauthorized returns a Response with status 401 Unauthorized.
// Similar to 403 Forbidden, but specifically for use when authentication is required and has failed or has not yet been provided.
func (response *Response) Http401() *Response {
	return response.Http401Unauthorized()
}

// PaymentRequired returns a Response with status 402 Payment Required.
// Reserved for future use.
func (response *Response) Http402PaymentRequired() *Response {
	response.status = http.StatusPaymentRequired
	return response
}

// PaymentRequired returns a Response with status 402 Payment Required.
// Reserved for future use.
func (response *Response) PaymentRequired() *Response {
	return response.Http402PaymentRequired()
}

// PaymentRequired returns a Response with status 402 Payment Required.
// Reserved for future use.
func (response *Response) Http402() *Response {
	return response.Http402PaymentRequired()
}

// Forbidden returns a Response with status 403 Forbidden.
// The request was valid, but the server is refusing action. The user might not have the necessary permissions for a resource.
func (response *Response) Http403Forbidden() *Response {
	response.status = http.StatusForbidden
	return response
}

// Forbidden returns a Response with status 403 Forbidden.
// The request was valid, but the server is refusing action. The user might not have the necessary permissions for a resource.
func (response *Response) Forbidden() *Response {
	return response.Http403Forbidden()
}

// Forbidden returns a Response with status 403 Forbidden.
// The request was valid, but the server is refusing action. The user might not have the necessary permissions for a resource.
func (response *Response) Http403() *Response {
	return response.Http403Forbidden()
}

// NotFound returns a Response with status 404 Not Found.
// The requested resource could not be found but may be available in the future.
func (response *Response) Http404NotFound() *Response {
	response.status = http.StatusNotFound
	return response
}

// NotFound returns a Response with status 404 Not Found.
// The requested resource could not be found but may be available in the future.
func (response *Response) NotFound() *Response {
	return response.Http404NotFound()
}

// NotFound returns a Response with status 404 Not Found.
// The requested resource could not be found but may be available in the future.
func (response *Response) Http404() *Response {
	return response.Http404NotFound()
}

// MethodNotAllowed returns a Response with status 405 Method Not Allowed.
// A request method is not supported for the requested resource.
func (response *Response) Http405MethodNotAllowed() *Response {
	response.status = http.StatusMethodNotAllowed
	return response
}

// MethodNotAllowed returns a Response with status 405 Method Not Allowed.
// A request method is not supported for the requested resource.
func (response *Response) MethodNotAllowed() *Response {
	return response.Http405MethodNotAllowed()
}

// MethodNotAllowed returns a Response with status 405 Method Not Allowed.
// A request method is not supported for the requested resource.
func (response *Response) Http405() *Response {
	return response.Http405MethodNotAllowed()
}

// NotAcceptable returns a Response with status 406 Not Acceptable.
// The requested resource is capable of generating only content not acceptable according to the Accept headers sent in the request.
func (response *Response) Http406NotAcceptable() *Response {
	response.status = http.StatusNotAcceptable
	return response
}

// NotAcceptable returns a Response with status 406 Not Acceptable.
// The requested resource is capable of generating only content not acceptable according to the Accept headers sent in the request.
func (response *Response) NotAcceptable() *Response {
	return response.Http406NotAcceptable()
}

// NotAcceptable returns a Response with status 406 Not Acceptable.
// The requested resource is capable of generating only content not acceptable according to the Accept headers sent in the request.
func (response *Response) Http406() *Response {
	return response.Http406NotAcceptable()
}

// ProxyAuthRequired returns a Response with status 407 Proxy Authentication Required.
// The client must first authenticate itself with the proxy.
func (response *Response) Http407ProxyAuthRequired() *Response {
	response.status = http.StatusProxyAuthRequired
	return response
}

// ProxyAuthRequired returns a Response with status 407 Proxy Authentication Required.
// The client must first authenticate itself with the proxy.
func (response *Response) ProxyAuthRequired() *Response {
	return response.Http407ProxyAuthRequired()
}

// ProxyAuthRequired returns a Response with status 407 Proxy Authentication Required.
// The client must first authenticate itself with the proxy.
func (response *Response) Http407() *Response {
	return response.Http407ProxyAuthRequired()
}

// RequestTimeout returns a Response with status 408 Request Timeout.
// The server timed out waiting for the request.
func (response *Response) Http408RequestTimeout() *Response {
	response.status = http.StatusRequestTimeout
	return response
}

// RequestTimeout returns a Response with status 408 Request Timeout.
// The server timed out waiting for the request.
func (response *Response) RequestTimeout() *Response {
	return response.Http408RequestTimeout()
}

// RequestTimeout returns a Response with status 408 Request Timeout.
// The server timed out waiting for the request.
func (response *Response) Http408() *Response {
	return response.Http408RequestTimeout()
}

// Conflict returns a Response with status 409 Conflict.
// Indicates that the request could not be processed because of conflict in the current state of the resource.
func (response *Response) Http409Conflict() *Response {
	response.status = http.StatusConflict
	return response
}

// Conflict returns a Response with status 409 Conflict.
// Indicates that the request could not be processed because of conflict in the current state of the resource.
func (response *Response) Conflict() *Response {
	return response.Http409Conflict()
}

// Conflict returns a Response with status 409 Conflict.
// Indicates that the request could not be processed because of conflict in the current state of the resource.
func (response *Response) Http409() *Response {
	return response.Http409Conflict()
}

// Gone returns a Response with status 410 Gone.
// Indicates that the resource requested is no longer available and will not be available again.
func (response *Response) Http410Gone() *Response {
	response.status = http.StatusGone
	return response
}

// Gone returns a Response with status 410 Gone.
// Indicates that the resource requested is no longer available and will not be available again.
func (response *Response) Gone() *Response {
	return response.Http410Gone()
}

// Gone returns a Response with status 410 Gone.
// Indicates that the resource requested is no longer available and will not be available again.
func (response *Response) Http410() *Response {
	return response.Http410Gone()
}

// LengthRequired returns a Response with status 411 Length Required.
// The request did not specify the length of its content, which is required by the requested resource.
func (response *Response) Http411LengthRequired() *Response {
	response.status = http.StatusLengthRequired
	return response
}

// LengthRequired returns a Response with status 411 Length Required.
// The request did not specify the length of its content, which is required by the requested resource.
func (response *Response) LengthRequired() *Response {
	return response.Http411LengthRequired()
}

// LengthRequired returns a Response with status 411 Length Required.
// The request did not specify the length of its content, which is required by the requested resource.
func (response *Response) Http411() *Response {
	return response.Http411LengthRequired()
}

// PreconditionFailed returns a Response with status 412 Precondition Failed.
// The server does not meet one of the preconditions that the requester put on the request.
func (response *Response) Http412PreconditionFailed() *Response {
	response.status = http.StatusPreconditionFailed
	return response
}

// PreconditionFailed returns a Response with status 412 Precondition Failed.
// The server does not meet one of the preconditions that the requester put on the request.
func (response *Response) PreconditionFailed() *Response {
	return response.Http412PreconditionFailed()
}

// PreconditionFailed returns a Response with status 412 Precondition Failed.
// The server does not meet one of the preconditions that the requester put on the request.
func (response *Response) Http412() *Response {
	return response.Http412PreconditionFailed()
}

// RequestEntityTooLarge returns a Response with status 413 Request Entity Too Large.
// The request is larger than the server is willing or able to process.
func (response *Response) Http413RequestEntityTooLarge() *Response {
	response.status = http.StatusRequestEntityTooLarge
	return response
}

// RequestEntityTooLarge returns a Response with status 413 Request Entity Too Large.
// The request is larger than the server is willing or able to process.
func (response *Response) RequestEntityTooLarge() *Response {
	return response.Http413RequestEntityTooLarge()
}

// RequestEntityTooLarge returns a Response with status 413 Request Entity Too Large.
// The request is larger than the server is willing or able to process.
func (response *Response) Http413() *Response {
	return response.Http413RequestEntityTooLarge()
}

// RequestURITooLong returns a Response with status 414 Request-URI Too Long.
// The URI provided was too long for the server to process.
func (response *Response) Http414RequestURITooLong() *Response {
	response.status = http.StatusRequestURITooLong
	return response
}

// RequestURITooLong returns a Response with status 414 Request-URI Too Long.
// The URI provided was too long for the server to process.
func (response *Response) RequestURITooLong() *Response {
	return response.Http414RequestURITooLong()
}

// RequestURITooLong returns a Response with status 414 Request-URI Too Long.
// The URI provided was too long for the server to process.
func (response *Response) Http414() *Response {
	return response.Http414RequestURITooLong()
}

// UnsupportedMediaType returns a Response with status 415 Unsupported Media Type.
// The request entity has a media type which the server or resource does not support.
func (response *Response) Http415UnsupportedMediaType() *Response {
	response.status = http.StatusUnsupportedMediaType
	return response
}

// UnsupportedMediaType returns a Response with status 415 Unsupported Media Type.
// The request entity has a media type which the server or resource does not support.
func (response *Response) UnsupportedMediaType() *Response {
	return response.Http415UnsupportedMediaType()
}

// UnsupportedMediaType returns a Response with status 415 Unsupported Media Type.
// The request entity has a media type which the server or resource does not support.
func (response *Response) Http415() *Response {
	return response.Http415UnsupportedMediaType()
}

// RequestedRangeNotSatisfiable returns a Response with status 416 Requested Range Not Satisfiable.
// The client has asked for a portion of the file, but the server cannot supply that portion.
func (response *Response) Http416RequestedRangeNotSatisfiable() *Response {
	response.status = http.StatusRequestedRangeNotSatisfiable
	return response
}

// RequestedRangeNotSatisfiable returns a Response with status 416 Requested Range Not Satisfiable.
// The client has asked for a portion of the file, but the server cannot supply that portion.
func (response *Response) RequestedRangeNotSatisfiable() *Response {
	return response.Http416RequestedRangeNotSatisfiable()
}

// RequestedRangeNotSatisfiable returns a Response with status 416 Requested Range Not Satisfiable.
// The client has asked for a portion of the file, but the server cannot supply that portion.
func (response *Response) Http416() *Response {
	return response.Http416RequestedRangeNotSatisfiable()
}

// ExpectationFailed returns a Response with status 417 Expectation Failed.
// The server cannot meet the requirements of the Expect request-header field.
func (response *Response) Http417ExpectationFailed() *Response {
	response.status = http.StatusExpectationFailed
	return response
}

// ExpectationFailed returns a Response with status 417 Expectation Failed.
// The server cannot meet the requirements of the Expect request-header field.
func (response *Response) ExpectationFailed() *Response {
	return response.Http417ExpectationFailed()
}

// ExpectationFailed returns a Response with status 417 Expectation Failed.
// The server cannot meet the requirements of the Expect request-header field.
func (response *Response) Http417() *Response {
	return response.Http417ExpectationFailed()
}

// Teapot returns a Response with status 418 I'm a teapot.
// This code was defined in 1998 as one of the traditional IETF April Fools' jokes, in RFC 2324, Hyper Text Coffee Pot Control Protocol.
func (response *Response) Http418Teapot() *Response {
	response.status = http.StatusTeapot
	return response
}

// Teapot returns a Response with status 418 I'm a teapot.
// This code was defined in 1998 as one of the traditional IETF April Fools' jokes, in RFC 2324, Hyper Text Coffee Pot Control Protocol.
func (response *Response) Teapot() *Response {
	return response.Http418Teapot()
}

// Teapot returns a Response with status 418 I'm a teapot.
// This code was defined in 1998 as one of the traditional IETF April Fools' jokes, in RFC 2324, Hyper Text Coffee Pot Control Protocol.
func (response *Response) Http418() *Response {
	return response.Http418Teapot()
}

// MisdirectedRequest returns a Response with status 421 Misdirected Request.
// The request was directed at a server that is not able to produce a response.
func (response *Response) Http421MisdirectedRequest() *Response {
	response.status = http.StatusMisdirectedRequest
	return response
}

// MisdirectedRequest returns a Response with status 421 Misdirected Request.
// The request was directed at a server that is not able to produce a response.
func (response *Response) MisdirectedRequest() *Response {
	return response.Http421MisdirectedRequest()
}

// MisdirectedRequest returns a Response with status 421 Misdirected Request.
// The request was directed at a server that is not able to produce a response.
func (response *Response) Http421() *Response {
	return response.Http421MisdirectedRequest()
}

// UnprocessableEntity returns a Response with status 422 Unprocessable Entity.
// The request was well-formed but was unable to be followed due to semantic errors.
func (response *Response) Http422UnprocessableEntity() *Response {
	response.status = http.StatusUnprocessableEntity
	return response
}

// UnprocessableEntity returns a Response with status 422 Unprocessable Entity.
// The request was well-formed but was unable to be followed due to semantic errors.
func (response *Response) UnprocessableEntity() *Response {
	return response.Http422UnprocessableEntity()
}

// UnprocessableEntity returns a Response with status 422 Unprocessable Entity.
// The request was well-formed but was unable to be followed due to semantic errors.
func (response *Response) Http422() *Response {
	return response.Http422UnprocessableEntity()
}

// Locked returns a Response with status 423 Locked.
// The resource that is being accessed is locked.
func (response *Response) Http423Locked() *Response {
	response.status = http.StatusLocked
	return response
}

// Locked returns a Response with status 423 Locked.
// The resource that is being accessed is locked.
func (response *Response) Locked() *Response {
	return response.Http423Locked()
}

// Locked returns a Response with status 423 Locked.
// The resource that is being accessed is locked.
func (response *Response) Http423() *Response {
	return response.Http423Locked()
}

// FailedDependency returns a Response with status 424 Failed Dependency.
// The request failed because it depended on another request and that request failed.
func (response *Response) Http424FailedDependency() *Response {
	response.status = http.StatusFailedDependency
	return response
}

// FailedDependency returns a Response with status 424 Failed Dependency.
// The request failed because it depended on another request and that request failed.
func (response *Response) FailedDependency() *Response {
	return response.Http424FailedDependency()
}

// FailedDependency returns a Response with status 424 Failed Dependency.
// The request failed because it depended on another request and that request failed.
func (response *Response) Http424() *Response {
	return response.Http424FailedDependency()
}

// TooEarly returns a Response with status 425 Too Early.
// Indicates that the server is unwilling to risk processing a request that might be replayed.
func (response *Response) Http425TooEarly() *Response {
	response.status = http.StatusTooEarly
	return response
}

// TooEarly returns a Response with status 425 Too Early.
// Indicates that the server is unwilling to risk processing a request that might be replayed.
func (response *Response) TooEarly() *Response {
	return response.Http425TooEarly()
}

// TooEarly returns a Response with status 425 Too Early.
// Indicates that the server is unwilling to risk processing a request that might be replayed.
func (response *Response) Http425() *Response {
	return response.Http425TooEarly()
}

// UpgradeRequired returns a Response with status 426 Upgrade Required.
// The client should switch to a different protocol such as TLS/1.0, given in the Upgrade header field.
func (response *Response) Http426UpgradeRequired() *Response {
	response.status = http.StatusUpgradeRequired
	return response
}

// UpgradeRequired returns a Response with status 426 Upgrade Required.
// The client should switch to a different protocol such as TLS/1.0, given in the Upgrade header field.
func (response *Response) UpgradeRequired() *Response {
	return response.Http426UpgradeRequired()
}

// UpgradeRequired returns a Response with status 426 Upgrade Required.
// The client should switch to a different protocol such as TLS/1.0, given in the Upgrade header field.
func (response *Response) Http426() *Response {
	return response.Http426UpgradeRequired()
}

// PreconditionRequired returns a Response with status 428 Precondition Required.
// The origin server requires the request to be conditional.
func (response *Response) Http428PreconditionRequired() *Response {
	response.status = http.StatusPreconditionRequired
	return response
}

// PreconditionRequired returns a Response with status 428 Precondition Required.
// The origin server requires the request to be conditional.
func (response *Response) PreconditionRequired() *Response {
	return response.Http428PreconditionRequired()
}

// PreconditionRequired returns a Response with status 428 Precondition Required.
// The origin server requires the request to be conditional.
func (response *Response) Http428() *Response {
	return response.Http428PreconditionRequired()
}

// TooManyRequests returns a Response with status 429 Too Many Requests.
// The user has sent too many requests in a given amount of time ("rate limiting").
func (response *Response) Http429TooManyRequests() *Response {
	response.status = http.StatusTooManyRequests
	return response
}

// TooManyRequests returns a Response with status 429 Too Many Requests.
// The user has sent too many requests in a given amount of time ("rate limiting").
func (response *Response) TooManyRequests() *Response {
	return response.Http429TooManyRequests()
}

// TooManyRequests returns a Response with status 429 Too Many Requests.
// The user has sent too many requests in a given amount of time ("rate limiting").
func (response *Response) Http429() *Response {
	return response.Http429TooManyRequests()
}

// RequestHeaderFieldsTooLarge returns a Response with status 431 Request Header Fields Too Large.
// The server is unwilling to process the request because either an individual header field, or all the header fields collectively, are too large.
func (response *Response) Http431RequestHeaderFieldsTooLarge() *Response {
	response.status = http.StatusRequestHeaderFieldsTooLarge
	return response
}

// RequestHeaderFieldsTooLarge returns a Response with status 431 Request Header Fields Too Large.
// The server is unwilling to process the request because either an individual header field, or all the header fields collectively, are too large.
func (response *Response) RequestHeaderFieldsTooLarge() *Response {
	return response.Http431RequestHeaderFieldsTooLarge()
}

// RequestHeaderFieldsTooLarge returns a Response with status 431 Request Header Fields Too Large.
// The server is unwilling to process the request because either an individual header field, or all the header fields collectively, are too large.
func (response *Response) Http431() *Response {
	return response.Http431RequestHeaderFieldsTooLarge()
}

// UnavailableForLegalReasons returns a Response with status 451 Unavailable For Legal Reasons.
// A server operator has received a legal demand to deny access to a resource or to a set of resources that includes the requested resource.
func (response *Response) Http451UnavailableForLegalReasons() *Response {
	response.status = http.StatusUnavailableForLegalReasons
	return response
}

// UnavailableForLegalReasons returns a Response with status 451 Unavailable For Legal Reasons.
// A server operator has received a legal demand to deny access to a resource or to a set of resources that includes the requested resource.
func (response *Response) UnavailableForLegalReasons() *Response {
	return response.Http451UnavailableForLegalReasons()
}

// UnavailableForLegalReasons returns a Response with status 451 Unavailable For Legal Reasons.
// A server operator has received a legal demand to deny access to a resource or to a set of resources that includes the requested resource.
func (response *Response) Http451() *Response {
	return response.Http451UnavailableForLegalReasons()
}

// InternalServerError returns a Response with status 500 Internal Server Error.
// A generic error message, given when an unexpected condition was encountered and no more specific message is suitable.
func (response *Response) Http500InternalServerError() *Response {
	response.status = http.StatusInternalServerError
	return response
}

// InternalServerError returns a Response with status 500 Internal Server Error.
// A generic error message, given when an unexpected condition was encountered and no more specific message is suitable.
func (response *Response) InternalServerError() *Response {
	return response.Http500InternalServerError()
}

// InternalServerError returns a Response with status 500 Internal Server Error.
// A generic error message, given when an unexpected condition was encountered and no more specific message is suitable.
func (response *Response) Http500() *Response {
	return response.Http500InternalServerError()
}

// NotImplemented returns a Response with status 501 Not Implemented.
// The server either does not recognize the request method, or it lacks the ability to fulfill the request.
func (response *Response) Http501NotImplemented() *Response {
	response.status = http.StatusNotImplemented
	return response
}

// NotImplemented returns a Response with status 501 Not Implemented.
// The server either does not recognize the request method, or it lacks the ability to fulfill the request.
func (response *Response) NotImplemented() *Response {
	return response.Http501NotImplemented()
}

// NotImplemented returns a Response with status 501 Not Implemented.
// The server either does not recognize the request method, or it lacks the ability to fulfill the request.
func (response *Response) Http501() *Response {
	return response.Http501NotImplemented()
}

// BadGateway returns a Response with status 502 Bad Gateway.
// The server was acting as a gateway or proxy and received an invalid response from the upstream server.
func (response *Response) Http502BadGateway() *Response {
	response.status = http.StatusBadGateway
	return response
}

// BadGateway returns a Response with status 502 Bad Gateway.
// The server was acting as a gateway or proxy and received an invalid response from the upstream server.
func (response *Response) BadGateway() *Response {
	return response.Http502BadGateway()
}

// BadGateway returns a Response with status 502 Bad Gateway.
// The server was acting as a gateway or proxy and received an invalid response from the upstream server.
func (response *Response) Http502() *Response {
	return response.Http502BadGateway()
}

// ServiceUnavailable returns a Response with status 503 Service Unavailable.
// The server is currently unavailable (because it is overloaded or down for maintenance).
func (response *Response) Http503ServiceUnavailable() *Response {
	response.status = http.StatusServiceUnavailable
	return response
}

// ServiceUnavailable returns a Response with status 503 Service Unavailable.
// The server is currently unavailable (because it is overloaded or down for maintenance).
func (response *Response) ServiceUnavailable() *Response {
	return response.Http503ServiceUnavailable()
}

// ServiceUnavailable returns a Response with status 503 Service Unavailable.
// The server is currently unavailable (because it is overloaded or down for maintenance).
func (response *Response) Http503() *Response {
	return response.Http503ServiceUnavailable()
}

// GatewayTimeout returns a Response with status 504 Gateway Timeout.
// The server was acting as a gateway or proxy and did not receive a timely response from the upstream server.
func (response *Response) Http504GatewayTimeout() *Response {
	response.status = http.StatusGatewayTimeout
	return response
}

// GatewayTimeout returns a Response with status 504 Gateway Timeout.
// The server was acting as a gateway or proxy and did not receive a timely response from the upstream server.
func (response *Response) GatewayTimeout() *Response {
	return response.Http504GatewayTimeout()
}

// GatewayTimeout returns a Response with status 504 Gateway Timeout.
// The server was acting as a gateway or proxy and did not receive a timely response from the upstream server.
func (response *Response) Http504() *Response {
	return response.Http504GatewayTimeout()
}

// HTTPVersionNotSupported returns a Response with status 505 HTTP Version Not Supported.
// The server does not support the HTTP protocol version used in the request.
func (response *Response) Http505HTTPVersionNotSupported() *Response {
	response.status = http.StatusHTTPVersionNotSupported
	return response
}

// HTTPVersionNotSupported returns a Response with status 505 HTTP Version Not Supported.
// The server does not support the HTTP protocol version used in the request.
func (response *Response) HTTPVersionNotSupported() *Response {
	return response.Http505HTTPVersionNotSupported()
}

// HTTPVersionNotSupported returns a Response with status 505 HTTP Version Not Supported.
// The server does not support the HTTP protocol version used in the request.
func (response *Response) Http505() *Response {
	return response.Http505HTTPVersionNotSupported()
}

// VariantAlsoNegotiates returns a Response with status 506 Variant Also Negotiates.
// Transparent content negotiation for the request results in a circular reference.
func (response *Response) Http506VariantAlsoNegotiates() *Response {
	response.status = http.StatusVariantAlsoNegotiates
	return response
}

// VariantAlsoNegotiates returns a Response with status 506 Variant Also Negotiates.
// Transparent content negotiation for the request results in a circular reference.
func (response *Response) VariantAlsoNegotiates() *Response {
	return response.Http506VariantAlsoNegotiates()
}

// VariantAlsoNegotiates returns a Response with status 506 Variant Also Negotiates.
// Transparent content negotiation for the request results in a circular reference.
func (response *Response) Http506() *Response {
	return response.Http506VariantAlsoNegotiates()
}

// InsufficientStorage returns a Response with status 507 Insufficient Storage.
// The server is unable to store the representation needed to complete the request.
func (response *Response) Http507InsufficientStorage() *Response {
	response.status = http.StatusInsufficientStorage
	return response
}

// InsufficientStorage returns a Response with status 507 Insufficient Storage.
// The server is unable to store the representation needed to complete the request.
func (response *Response) InsufficientStorage() *Response {
	return response.Http507InsufficientStorage()
}

// InsufficientStorage returns a Response with status 507 Insufficient Storage.
// The server is unable to store the representation needed to complete the request.
func (response *Response) Http507() *Response {
	return response.Http507InsufficientStorage()
}

// LoopDetected returns a Response with status 508 Loop Detected.
// The server detected an infinite loop while processing the request.
func (response *Response) Http508LoopDetected() *Response {
	response.status = http.StatusLoopDetected
	return response
}

// LoopDetected returns a Response with status 508 Loop Detected.
// The server detected an infinite loop while processing the request.
func (response *Response) LoopDetected() *Response {
	return response.Http508LoopDetected()
}

// LoopDetected returns a Response with status 508 Loop Detected.
// The server detected an infinite loop while processing the request.
func (response *Response) Http508() *Response {
	return response.Http508LoopDetected()
}

// NotExtended returns a Response with status 510 Not Extended.
// Further extensions to the request are required for the server to fulfill it.
func (response *Response) Http510NotExtended() *Response {
	response.status = http.StatusNotExtended
	return response
}

// NotExtended returns a Response with status 510 Not Extended.
// Further extensions to the request are required for the server to fulfill it.
func (response *Response) NotExtended() *Response {
	return response.Http510NotExtended()
}

// NotExtended returns a Response with status 510 Not Extended.
// Further extensions to the request are required for the server to fulfill it.
func (response *Response) Http510() *Response {
	return response.Http510NotExtended()
}

// NetworkAuthenticationRequired returns a Response with status 511 Network Authentication Required.
// The client needs to authenticate to gain network access.
func (response *Response) Http511NetworkAuthenticationRequired() *Response {
	response.status = http.StatusNetworkAuthenticationRequired
	return response
}

// NetworkAuthenticationRequired returns a Response with status 511 Network Authentication Required.
// The client needs to authenticate to gain network access.
func (response *Response) NetworkAuthenticationRequired() *Response {
	return response.Http511NetworkAuthenticationRequired()
}

// NetworkAuthenticationRequired returns a Response with status 511 Network Authentication Required.
// The client needs to authenticate to gain network access.
func (response *Response) Http511() *Response {
	return response.Http511NetworkAuthenticationRequired()
}
