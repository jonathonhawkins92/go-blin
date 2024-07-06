package responses

import "net/http"

// Continue returns a Response with status 100 Continue.
// The server has received the request headers and the client should proceed to send the request body.
func Http100Continue() *Response {
	return &Response{status: http.StatusContinue}
}

// Continue returns a Response with status 100 Continue.
// The server has received the request headers and the client should proceed to send the request body.
func Continue() *Response {
	return Http100Continue()
}

// Continue returns a Response with status 100 Continue.
// The server has received the request headers and the client should proceed to send the request body.
func Http100() *Response {
	return Http100Continue()
}

// SwitchingProtocols returns a Response with status 101 Switching Protocols.
// The server is switching protocols according to the Upgrade header sent by the client.
func Http101SwitchingProtocols() *Response {
	return &Response{status: http.StatusSwitchingProtocols}
}

// SwitchingProtocols returns a Response with status 101 Switching Protocols.
// The server is switching protocols according to the Upgrade header sent by the client.
func SwitchingProtocols() *Response {
	return Http101SwitchingProtocols()
}

// SwitchingProtocols returns a Response with status 101 Switching Protocols.
// The server is switching protocols according to the Upgrade header sent by the client.
func Http101() *Response {
	return Http101SwitchingProtocols()
}

// Processing returns a Response with status 102 Processing.
// The server has received and is processing the request, but no response is available yet.
func Http102Processing() *Response {
	return &Response{status: http.StatusProcessing}
}

// Processing returns a Response with status 102 Processing.
// The server has received and is processing the request, but no response is available yet.
func Processing() *Response {
	return Http102Processing()
}

// Processing returns a Response with status 102 Processing.
// The server has received and is processing the request, but no response is available yet.
func Http102() *Response {
	return Http102Processing()
}

// EarlyHints returns a Response with status 103 Early Hints.
// Used to return some response headers before final HTTP message.
func Http103EarlyHints() *Response {
	return &Response{status: http.StatusEarlyHints}
}

// EarlyHints returns a Response with status 103 Early Hints.
// Used to return some response headers before final HTTP message.
func EarlyHints() *Response {
	return Http103EarlyHints()
}

// EarlyHints returns a Response with status 103 Early Hints.
// Used to return some response headers before final HTTP message.
func Http103() *Response {
	return Http103EarlyHints()
}

// OK returns a Response with status 200 OK.
// The request has succeeded.
func Http200OK() *Response {
	return &Response{status: http.StatusOK}
}

// OK returns a Response with status 200 OK.
// The request has succeeded.
func OK() *Response {
	return Http200OK()
}

// OK returns a Response with status 200 OK.
// The request has succeeded.
func Http200() *Response {
	return Http200OK()
}

// Created returns a Response with status 201 Created.
// The request has been fulfilled and a new resource has been created.
func Http201Created() *Response {
	return &Response{status: http.StatusCreated}
}

// Created returns a Response with status 201 Created.
// The request has been fulfilled and a new resource has been created.
func Created() *Response {
	return Http201Created()
}

// Created returns a Response with status 201 Created.
// The request has been fulfilled and a new resource has been created.
func Http201() *Response {
	return Http201Created()
}

// Accepted returns a Response with status 202 Accepted.
// The request has been accepted for processing, but the processing has not been completed.
func Http202Accepted() *Response {
	return &Response{status: http.StatusAccepted}
}

// Accepted returns a Response with status 202 Accepted.
// The request has been accepted for processing, but the processing has not been completed.
func Accepted() *Response {
	return Http202Accepted()
}

// Accepted returns a Response with status 202 Accepted.
// The request has been accepted for processing, but the processing has not been completed.
func Http202() *Response {
	return Http202Accepted()
}

// NonAuthoritativeInfo returns a Response with status 203 Non-Authoritative Information.
// The server is a transforming proxy that received a 200 OK from its origin but is returning a modified version of the origin's response.
func Http203NonAuthoritativeInfo() *Response {
	return &Response{status: http.StatusNonAuthoritativeInfo}
}

// NonAuthoritativeInfo returns a Response with status 203 Non-Authoritative Information.
// The server is a transforming proxy that received a 200 OK from its origin but is returning a modified version of the origin's response.
func NonAuthoritativeInfo() *Response {
	return Http203NonAuthoritativeInfo()
}

// NonAuthoritativeInfo returns a Response with status 203 Non-Authoritative Information.
// The server is a transforming proxy that received a 200 OK from its origin but is returning a modified version of the origin's response.
func Http203() *Response {
	return Http203NonAuthoritativeInfo()
}

// NoContent returns a Response with status 204 No Content.
// The server successfully processed the request and is not returning any content.
func Http204NoContent() *Response {
	return &Response{status: http.StatusNoContent}
}

// NoContent returns a Response with status 204 No Content.
// The server successfully processed the request and is not returning any content.
func NoContent() *Response {
	return Http204NoContent()
}

// NoContent returns a Response with status 204 No Content.
// The server successfully processed the request and is not returning any content.
func Http204() *Response {
	return Http204NoContent()
}

// ResetContent returns a Response with status 205 Reset Content.
// The server successfully processed the request, but is not returning any content. The client should reset the document view.
func Http205ResetContent() *Response {
	return &Response{status: http.StatusResetContent}
}

// ResetContent returns a Response with status 205 Reset Content.
// The server successfully processed the request, but is not returning any content. The client should reset the document view.
func ResetContent() *Response {
	return Http205ResetContent()
}

// ResetContent returns a Response with status 205 Reset Content.
// The server successfully processed the request, but is not returning any content. The client should reset the document view.
func Http205() *Response {
	return Http205ResetContent()
}

// PartialContent returns a Response with status 206 Partial Content.
// The server is delivering only part of the resource due to a range header sent by the client.
func Http206PartialContent() *Response {
	return &Response{status: http.StatusPartialContent}
}

// PartialContent returns a Response with status 206 Partial Content.
// The server is delivering only part of the resource due to a range header sent by the client.
func PartialContent() *Response {
	return Http206PartialContent()
}

// PartialContent returns a Response with status 206 Partial Content.
// The server is delivering only part of the resource due to a range header sent by the client.
func Http206() *Response {
	return Http206PartialContent()
}

// MultiStatus returns a Response with status 207 Multi-Status.
// The message body that follows is an XML message and can contain a number of separate response codes.
func Http207MultiStatus() *Response {
	return &Response{status: http.StatusMultiStatus}
}

// MultiStatus returns a Response with status 207 Multi-Status.
// The message body that follows is an XML message and can contain a number of separate response codes.
func MultiStatus() *Response {
	return Http207MultiStatus()
}

// MultiStatus returns a Response with status 207 Multi-Status.
// The message body that follows is an XML message and can contain a number of separate response codes.
func Http207() *Response {
	return Http207MultiStatus()
}

// AlreadyReported returns a Response with status 208 Already Reported.
// The members of a DAV binding have already been enumerated in a previous reply to this request, and are not being included again.
func Http208AlreadyReported() *Response {
	return &Response{status: http.StatusAlreadyReported}
}

// AlreadyReported returns a Response with status 208 Already Reported.
// The members of a DAV binding have already been enumerated in a previous reply to this request, and are not being included again.
func AlreadyReported() *Response {
	return Http208AlreadyReported()
}

// AlreadyReported returns a Response with status 208 Already Reported.
// The members of a DAV binding have already been enumerated in a previous reply to this request, and are not being included again.
func Http208() *Response {
	return Http208AlreadyReported()
}

// IMUsed returns a Response with status 226 IM Used.
// The server has fulfilled a request for the resource, and the response is a representation of the result of one or more instance-manipulations applied to the current instance.
func Http226IMUsed() *Response {
	return &Response{status: http.StatusIMUsed}
}

// IMUsed returns a Response with status 226 IM Used.
// The server has fulfilled a request for the resource, and the response is a representation of the result of one or more instance-manipulations applied to the current instance.
func IMUsed() *Response {
	return Http226IMUsed()
}

// IMUsed returns a Response with status 226 IM Used.
// The server has fulfilled a request for the resource, and the response is a representation of the result of one or more instance-manipulations applied to the current instance.
func Http226() *Response {
	return Http226IMUsed()
}

// MultipleChoices returns a Response with status 300 Multiple Choices.
// The requested resource corresponds to any one of a set of representations, each with its own specific location.
func Http300MultipleChoices() *Response {
	return &Response{status: http.StatusMultipleChoices}
}

// MultipleChoices returns a Response with status 300 Multiple Choices.
// The requested resource corresponds to any one of a set of representations, each with its own specific location.
func MultipleChoices() *Response {
	return Http300MultipleChoices()
}

// MultipleChoices returns a Response with status 300 Multiple Choices.
// The requested resource corresponds to any one of a set of representations, each with its own specific location.
func Http300() *Response {
	return Http300MultipleChoices()
}

// MovedPermanently returns a Response with status 301 Moved Permanently.
// The requested resource has been assigned a new permanent URI and any future references to this resource should use one of the returned URIs.
func Http301MovedPermanently() *Response {
	return &Response{status: http.StatusMovedPermanently}
}

// MovedPermanently returns a Response with status 301 Moved Permanently.
// The requested resource has been assigned a new permanent URI and any future references to this resource should use one of the returned URIs.
func MovedPermanently() *Response {
	return Http301MovedPermanently()
}

// MovedPermanently returns a Response with status 301 Moved Permanently.
// The requested resource has been assigned a new permanent URI and any future references to this resource should use one of the returned URIs.
func Http301() *Response {
	return Http301MovedPermanently()
}

// Found returns a Response with status 302 Found.
// The requested resource resides temporarily under a different URI.
func Http302Found() *Response {
	return &Response{status: http.StatusFound}
}

// Found returns a Response with status 302 Found.
// The requested resource resides temporarily under a different URI.
func Found() *Response {
	return Http302Found()
}

// Found returns a Response with status 302 Found.
// The requested resource resides temporarily under a different URI.
func Http302() *Response {
	return Http302Found()
}

// SeeOther returns a Response with status 303 See Other.
// The response to the request can be found under a different URI and should be retrieved using a GET method on that resource.
func Http303SeeOther() *Response {
	return &Response{status: http.StatusSeeOther}
}

// SeeOther returns a Response with status 303 See Other.
// The response to the request can be found under a different URI and should be retrieved using a GET method on that resource.
func SeeOther() *Response {
	return Http303SeeOther()
}

// SeeOther returns a Response with status 303 See Other.
// The response to the request can be found under a different URI and should be retrieved using a GET method on that resource.
func Http303() *Response {
	return Http303SeeOther()
}

// NotModified returns a Response with status 304 Not Modified.
// Indicates that the resource has not been modified since the version specified by the request headers If-Modified-Since or If-None-Match.
func Http304NotModified() *Response {
	return &Response{status: http.StatusNotModified}
}

// NotModified returns a Response with status 304 Not Modified.
// Indicates that the resource has not been modified since the version specified by the request headers If-Modified-Since or If-None-Match.
func NotModified() *Response {
	return Http304NotModified()
}

// NotModified returns a Response with status 304 Not Modified.
// Indicates that the resource has not been modified since the version specified by the request headers If-Modified-Since or If-None-Match.
func Http304() *Response {
	return Http304NotModified()
}

// UseProxy returns a Response with status 305 Use Proxy.
// The requested resource must be accessed through the proxy given by the Location field.
func Http305UseProxy() *Response {
	return &Response{status: http.StatusUseProxy}
}

// UseProxy returns a Response with status 305 Use Proxy.
// The requested resource must be accessed through the proxy given by the Location field.
func UseProxy() *Response {
	return Http305UseProxy()
}

// UseProxy returns a Response with status 305 Use Proxy.
// The requested resource must be accessed through the proxy given by the Location field.
func Http305() *Response {
	return Http305UseProxy()
}

// TemporaryRedirect returns a Response with status 307 Temporary Redirect.
// The requested resource resides temporarily under a different URI.
func Http307TemporaryRedirect() *Response {
	return &Response{status: http.StatusTemporaryRedirect}
}

// TemporaryRedirect returns a Response with status 307 Temporary Redirect.
// The requested resource resides temporarily under a different URI.
func TemporaryRedirect() *Response {
	return Http307TemporaryRedirect()
}

// TemporaryRedirect returns a Response with status 307 Temporary Redirect.
// The requested resource resides temporarily under a different URI.
func Http307() *Response {
	return Http307TemporaryRedirect()
}

// PermanentRedirect returns a Response with status 308 Permanent Redirect.
// The requested resource has been assigned a new permanent URI and any future references to this resource should use one of the returned URIs.
func Http308PermanentRedirect() *Response {
	return &Response{status: http.StatusPermanentRedirect}
}

// PermanentRedirect returns a Response with status 308 Permanent Redirect.
// The requested resource has been assigned a new permanent URI and any future references to this resource should use one of the returned URIs.
func PermanentRedirect() *Response {
	return Http308PermanentRedirect()
}

// PermanentRedirect returns a Response with status 308 Permanent Redirect.
// The requested resource has been assigned a new permanent URI and any future references to this resource should use one of the returned URIs.
func Http308() *Response {
	return Http308PermanentRedirect()
}

// BadRequest returns a Response with status 400 Bad Request.
// The server cannot or will not process the request due to an apparent client error.
func Http400BadRequest() *Response {
	return &Response{status: http.StatusBadRequest}
}

// BadRequest returns a Response with status 400 Bad Request.
// The server cannot or will not process the request due to an apparent client error.
func BadRequest() *Response {
	return Http400BadRequest()
}

// BadRequest returns a Response with status 400 Bad Request.
// The server cannot or will not process the request due to an apparent client error.
func Http400() *Response {
	return Http400BadRequest()
}

// Unauthorized returns a Response with status 401 Unauthorized.
// Similar to 403 Forbidden, but specifically for use when authentication is required and has failed or has not yet been provided.
func Http401Unauthorized() *Response {
	return &Response{status: http.StatusUnauthorized}
}

// Unauthorized returns a Response with status 401 Unauthorized.
// Similar to 403 Forbidden, but specifically for use when authentication is required and has failed or has not yet been provided.
func Unauthorized() *Response {
	return Http401Unauthorized()
}

// Unauthorized returns a Response with status 401 Unauthorized.
// Similar to 403 Forbidden, but specifically for use when authentication is required and has failed or has not yet been provided.
func Http401() *Response {
	return Http401Unauthorized()
}

// PaymentRequired returns a Response with status 402 Payment Required.
// Reserved for future use.
func Http402PaymentRequired() *Response {
	return &Response{status: http.StatusPaymentRequired}
}

// PaymentRequired returns a Response with status 402 Payment Required.
// Reserved for future use.
func PaymentRequired() *Response {
	return Http402PaymentRequired()
}

// PaymentRequired returns a Response with status 402 Payment Required.
// Reserved for future use.
func Http402() *Response {
	return Http402PaymentRequired()
}

// Forbidden returns a Response with status 403 Forbidden.
// The request was valid, but the server is refusing action. The user might not have the necessary permissions for a resource.
func Http403Forbidden() *Response {
	return &Response{status: http.StatusForbidden}
}

// Forbidden returns a Response with status 403 Forbidden.
// The request was valid, but the server is refusing action. The user might not have the necessary permissions for a resource.
func Forbidden() *Response {
	return Http403Forbidden()
}

// Forbidden returns a Response with status 403 Forbidden.
// The request was valid, but the server is refusing action. The user might not have the necessary permissions for a resource.
func Http403() *Response {
	return Http403Forbidden()
}

// NotFound returns a Response with status 404 Not Found.
// The requested resource could not be found but may be available in the future.
func Http404NotFound() *Response {
	return &Response{status: http.StatusNotFound}
}

// NotFound returns a Response with status 404 Not Found.
// The requested resource could not be found but may be available in the future.
func NotFound() *Response {
	return Http404NotFound()
}

// NotFound returns a Response with status 404 Not Found.
// The requested resource could not be found but may be available in the future.
func Http404() *Response {
	return Http404NotFound()
}

// MethodNotAllowed returns a Response with status 405 Method Not Allowed.
// A request method is not supported for the requested resource.
func Http405MethodNotAllowed() *Response {
	return &Response{status: http.StatusMethodNotAllowed}
}

// MethodNotAllowed returns a Response with status 405 Method Not Allowed.
// A request method is not supported for the requested resource.
func MethodNotAllowed() *Response {
	return Http405MethodNotAllowed()
}

// MethodNotAllowed returns a Response with status 405 Method Not Allowed.
// A request method is not supported for the requested resource.
func Http405() *Response {
	return Http405MethodNotAllowed()
}

// NotAcceptable returns a Response with status 406 Not Acceptable.
// The requested resource is capable of generating only content not acceptable according to the Accept headers sent in the request.
func Http406NotAcceptable() *Response {
	return &Response{status: http.StatusNotAcceptable}
}

// NotAcceptable returns a Response with status 406 Not Acceptable.
// The requested resource is capable of generating only content not acceptable according to the Accept headers sent in the request.
func NotAcceptable() *Response {
	return Http406NotAcceptable()
}

// NotAcceptable returns a Response with status 406 Not Acceptable.
// The requested resource is capable of generating only content not acceptable according to the Accept headers sent in the request.
func Http406() *Response {
	return Http406NotAcceptable()
}

// ProxyAuthRequired returns a Response with status 407 Proxy Authentication Required.
// The client must first authenticate itself with the proxy.
func Http407ProxyAuthRequired() *Response {
	return &Response{status: http.StatusProxyAuthRequired}
}

// ProxyAuthRequired returns a Response with status 407 Proxy Authentication Required.
// The client must first authenticate itself with the proxy.
func ProxyAuthRequired() *Response {
	return Http407ProxyAuthRequired()
}

// ProxyAuthRequired returns a Response with status 407 Proxy Authentication Required.
// The client must first authenticate itself with the proxy.
func Http407() *Response {
	return Http407ProxyAuthRequired()
}

// RequestTimeout returns a Response with status 408 Request Timeout.
// The server timed out waiting for the request.
func Http408RequestTimeout() *Response {
	return &Response{status: http.StatusRequestTimeout}
}

// RequestTimeout returns a Response with status 408 Request Timeout.
// The server timed out waiting for the request.
func RequestTimeout() *Response {
	return Http408RequestTimeout()
}

// RequestTimeout returns a Response with status 408 Request Timeout.
// The server timed out waiting for the request.
func Http408() *Response {
	return Http408RequestTimeout()
}

// Conflict returns a Response with status 409 Conflict.
// Indicates that the request could not be processed because of conflict in the current state of the resource.
func Http409Conflict() *Response {
	return &Response{status: http.StatusConflict}
}

// Conflict returns a Response with status 409 Conflict.
// Indicates that the request could not be processed because of conflict in the current state of the resource.
func Conflict() *Response {
	return Http409Conflict()
}

// Conflict returns a Response with status 409 Conflict.
// Indicates that the request could not be processed because of conflict in the current state of the resource.
func Http409() *Response {
	return Http409Conflict()
}

// Gone returns a Response with status 410 Gone.
// Indicates that the resource requested is no longer available and will not be available again.
func Http410Gone() *Response {
	return &Response{status: http.StatusGone}
}

// Gone returns a Response with status 410 Gone.
// Indicates that the resource requested is no longer available and will not be available again.
func Gone() *Response {
	return Http410Gone()
}

// Gone returns a Response with status 410 Gone.
// Indicates that the resource requested is no longer available and will not be available again.
func Http410() *Response {
	return Http410Gone()
}

// LengthRequired returns a Response with status 411 Length Required.
// The request did not specify the length of its content, which is required by the requested resource.
func Http411LengthRequired() *Response {
	return &Response{status: http.StatusLengthRequired}
}

// LengthRequired returns a Response with status 411 Length Required.
// The request did not specify the length of its content, which is required by the requested resource.
func LengthRequired() *Response {
	return Http411LengthRequired()
}

// LengthRequired returns a Response with status 411 Length Required.
// The request did not specify the length of its content, which is required by the requested resource.
func Http411() *Response {
	return Http411LengthRequired()
}

// PreconditionFailed returns a Response with status 412 Precondition Failed.
// The server does not meet one of the preconditions that the requester put on the request.
func Http412PreconditionFailed() *Response {
	return &Response{status: http.StatusPreconditionFailed}
}

// PreconditionFailed returns a Response with status 412 Precondition Failed.
// The server does not meet one of the preconditions that the requester put on the request.
func PreconditionFailed() *Response {
	return Http412PreconditionFailed()
}

// PreconditionFailed returns a Response with status 412 Precondition Failed.
// The server does not meet one of the preconditions that the requester put on the request.
func Http412() *Response {
	return Http412PreconditionFailed()
}

// RequestEntityTooLarge returns a Response with status 413 Request Entity Too Large.
// The request is larger than the server is willing or able to process.
func Http413RequestEntityTooLarge() *Response {
	return &Response{status: http.StatusRequestEntityTooLarge}
}

// RequestEntityTooLarge returns a Response with status 413 Request Entity Too Large.
// The request is larger than the server is willing or able to process.
func RequestEntityTooLarge() *Response {
	return Http413RequestEntityTooLarge()
}

// RequestEntityTooLarge returns a Response with status 413 Request Entity Too Large.
// The request is larger than the server is willing or able to process.
func Http413() *Response {
	return Http413RequestEntityTooLarge()
}

// RequestURITooLong returns a Response with status 414 Request-URI Too Long.
// The URI provided was too long for the server to process.
func Http414RequestURITooLong() *Response {
	return &Response{status: http.StatusRequestURITooLong}
}

// RequestURITooLong returns a Response with status 414 Request-URI Too Long.
// The URI provided was too long for the server to process.
func RequestURITooLong() *Response {
	return Http414RequestURITooLong()
}

// RequestURITooLong returns a Response with status 414 Request-URI Too Long.
// The URI provided was too long for the server to process.
func Http414() *Response {
	return Http414RequestURITooLong()
}

// UnsupportedMediaType returns a Response with status 415 Unsupported Media Type.
// The request entity has a media type which the server or resource does not support.
func Http415UnsupportedMediaType() *Response {
	return &Response{status: http.StatusUnsupportedMediaType}
}

// UnsupportedMediaType returns a Response with status 415 Unsupported Media Type.
// The request entity has a media type which the server or resource does not support.
func UnsupportedMediaType() *Response {
	return Http415UnsupportedMediaType()
}

// UnsupportedMediaType returns a Response with status 415 Unsupported Media Type.
// The request entity has a media type which the server or resource does not support.
func Http415() *Response {
	return Http415UnsupportedMediaType()
}

// RequestedRangeNotSatisfiable returns a Response with status 416 Requested Range Not Satisfiable.
// The client has asked for a portion of the file, but the server cannot supply that portion.
func Http416RequestedRangeNotSatisfiable() *Response {
	return &Response{status: http.StatusRequestedRangeNotSatisfiable}
}

// RequestedRangeNotSatisfiable returns a Response with status 416 Requested Range Not Satisfiable.
// The client has asked for a portion of the file, but the server cannot supply that portion.
func RequestedRangeNotSatisfiable() *Response {
	return Http416RequestedRangeNotSatisfiable()
}

// RequestedRangeNotSatisfiable returns a Response with status 416 Requested Range Not Satisfiable.
// The client has asked for a portion of the file, but the server cannot supply that portion.
func Http416() *Response {
	return Http416RequestedRangeNotSatisfiable()
}

// ExpectationFailed returns a Response with status 417 Expectation Failed.
// The server cannot meet the requirements of the Expect request-header field.
func Http417ExpectationFailed() *Response {
	return &Response{status: http.StatusExpectationFailed}
}

// ExpectationFailed returns a Response with status 417 Expectation Failed.
// The server cannot meet the requirements of the Expect request-header field.
func ExpectationFailed() *Response {
	return Http417ExpectationFailed()
}

// ExpectationFailed returns a Response with status 417 Expectation Failed.
// The server cannot meet the requirements of the Expect request-header field.
func Http417() *Response {
	return Http417ExpectationFailed()
}

// Teapot returns a Response with status 418 I'm a teapot.
// This code was defined in 1998 as one of the traditional IETF April Fools' jokes, in RFC 2324, Hyper Text Coffee Pot Control Protocol.
func Http418Teapot() *Response {
	return &Response{status: http.StatusTeapot}
}

// Teapot returns a Response with status 418 I'm a teapot.
// This code was defined in 1998 as one of the traditional IETF April Fools' jokes, in RFC 2324, Hyper Text Coffee Pot Control Protocol.
func Teapot() *Response {
	return Http418Teapot()
}

// Teapot returns a Response with status 418 I'm a teapot.
// This code was defined in 1998 as one of the traditional IETF April Fools' jokes, in RFC 2324, Hyper Text Coffee Pot Control Protocol.
func Http418() *Response {
	return Http418Teapot()
}

// MisdirectedRequest returns a Response with status 421 Misdirected Request.
// The request was directed at a server that is not able to produce a response.
func Http421MisdirectedRequest() *Response {
	return &Response{status: http.StatusMisdirectedRequest}
}

// MisdirectedRequest returns a Response with status 421 Misdirected Request.
// The request was directed at a server that is not able to produce a response.
func MisdirectedRequest() *Response {
	return Http421MisdirectedRequest()
}

// MisdirectedRequest returns a Response with status 421 Misdirected Request.
// The request was directed at a server that is not able to produce a response.
func Http421() *Response {
	return Http421MisdirectedRequest()
}

// UnprocessableEntity returns a Response with status 422 Unprocessable Entity.
// The request was well-formed but was unable to be followed due to semantic errors.
func Http422UnprocessableEntity() *Response {
	return &Response{status: http.StatusUnprocessableEntity}
}

// UnprocessableEntity returns a Response with status 422 Unprocessable Entity.
// The request was well-formed but was unable to be followed due to semantic errors.
func UnprocessableEntity() *Response {
	return Http422UnprocessableEntity()
}

// UnprocessableEntity returns a Response with status 422 Unprocessable Entity.
// The request was well-formed but was unable to be followed due to semantic errors.
func Http422() *Response {
	return Http422UnprocessableEntity()
}

// Locked returns a Response with status 423 Locked.
// The resource that is being accessed is locked.
func Http423Locked() *Response {
	return &Response{status: http.StatusLocked}
}

// Locked returns a Response with status 423 Locked.
// The resource that is being accessed is locked.
func Locked() *Response {
	return Http423Locked()
}

// Locked returns a Response with status 423 Locked.
// The resource that is being accessed is locked.
func Http423() *Response {
	return Http423Locked()
}

// FailedDependency returns a Response with status 424 Failed Dependency.
// The request failed because it depended on another request and that request failed.
func Http424FailedDependency() *Response {
	return &Response{status: http.StatusFailedDependency}
}

// FailedDependency returns a Response with status 424 Failed Dependency.
// The request failed because it depended on another request and that request failed.
func FailedDependency() *Response {
	return Http424FailedDependency()
}

// FailedDependency returns a Response with status 424 Failed Dependency.
// The request failed because it depended on another request and that request failed.
func Http424() *Response {
	return Http424FailedDependency()
}

// TooEarly returns a Response with status 425 Too Early.
// Indicates that the server is unwilling to risk processing a request that might be replayed.
func Http425TooEarly() *Response {
	return &Response{status: http.StatusTooEarly}
}

// TooEarly returns a Response with status 425 Too Early.
// Indicates that the server is unwilling to risk processing a request that might be replayed.
func TooEarly() *Response {
	return Http425TooEarly()
}

// TooEarly returns a Response with status 425 Too Early.
// Indicates that the server is unwilling to risk processing a request that might be replayed.
func Http425() *Response {
	return Http425TooEarly()
}

// UpgradeRequired returns a Response with status 426 Upgrade Required.
// The client should switch to a different protocol such as TLS/1.0, given in the Upgrade header field.
func Http426UpgradeRequired() *Response {
	return &Response{status: http.StatusUpgradeRequired}
}

// UpgradeRequired returns a Response with status 426 Upgrade Required.
// The client should switch to a different protocol such as TLS/1.0, given in the Upgrade header field.
func UpgradeRequired() *Response {
	return Http426UpgradeRequired()
}

// UpgradeRequired returns a Response with status 426 Upgrade Required.
// The client should switch to a different protocol such as TLS/1.0, given in the Upgrade header field.
func Http426() *Response {
	return Http426UpgradeRequired()
}

// PreconditionRequired returns a Response with status 428 Precondition Required.
// The origin server requires the request to be conditional.
func Http428PreconditionRequired() *Response {
	return &Response{status: http.StatusPreconditionRequired}
}

// PreconditionRequired returns a Response with status 428 Precondition Required.
// The origin server requires the request to be conditional.
func PreconditionRequired() *Response {
	return Http428PreconditionRequired()
}

// PreconditionRequired returns a Response with status 428 Precondition Required.
// The origin server requires the request to be conditional.
func Http428() *Response {
	return Http428PreconditionRequired()
}

// TooManyRequests returns a Response with status 429 Too Many Requests.
// The user has sent too many requests in a given amount of time ("rate limiting").
func Http429TooManyRequests() *Response {
	return &Response{status: http.StatusTooManyRequests}
}

// TooManyRequests returns a Response with status 429 Too Many Requests.
// The user has sent too many requests in a given amount of time ("rate limiting").
func TooManyRequests() *Response {
	return Http429TooManyRequests()
}

// TooManyRequests returns a Response with status 429 Too Many Requests.
// The user has sent too many requests in a given amount of time ("rate limiting").
func Http429() *Response {
	return Http429TooManyRequests()
}

// RequestHeaderFieldsTooLarge returns a Response with status 431 Request Header Fields Too Large.
// The server is unwilling to process the request because either an individual header field, or all the header fields collectively, are too large.
func Http431RequestHeaderFieldsTooLarge() *Response {
	return &Response{status: http.StatusRequestHeaderFieldsTooLarge}
}

// RequestHeaderFieldsTooLarge returns a Response with status 431 Request Header Fields Too Large.
// The server is unwilling to process the request because either an individual header field, or all the header fields collectively, are too large.
func RequestHeaderFieldsTooLarge() *Response {
	return Http431RequestHeaderFieldsTooLarge()
}

// RequestHeaderFieldsTooLarge returns a Response with status 431 Request Header Fields Too Large.
// The server is unwilling to process the request because either an individual header field, or all the header fields collectively, are too large.
func Http431() *Response {
	return Http431RequestHeaderFieldsTooLarge()
}

// UnavailableForLegalReasons returns a Response with status 451 Unavailable For Legal Reasons.
// A server operator has received a legal demand to deny access to a resource or to a set of resources that includes the requested resource.
func Http451UnavailableForLegalReasons() *Response {
	return &Response{status: http.StatusUnavailableForLegalReasons}
}

// UnavailableForLegalReasons returns a Response with status 451 Unavailable For Legal Reasons.
// A server operator has received a legal demand to deny access to a resource or to a set of resources that includes the requested resource.
func UnavailableForLegalReasons() *Response {
	return Http451UnavailableForLegalReasons()
}

// UnavailableForLegalReasons returns a Response with status 451 Unavailable For Legal Reasons.
// A server operator has received a legal demand to deny access to a resource or to a set of resources that includes the requested resource.
func Http451() *Response {
	return Http451UnavailableForLegalReasons()
}

// InternalServerError returns a Response with status 500 Internal Server Error.
// A generic error message, given when an unexpected condition was encountered and no more specific message is suitable.
func Http500InternalServerError() *Response {
	return &Response{status: http.StatusInternalServerError}
}

// InternalServerError returns a Response with status 500 Internal Server Error.
// A generic error message, given when an unexpected condition was encountered and no more specific message is suitable.
func InternalServerError() *Response {
	return Http500InternalServerError()
}

// InternalServerError returns a Response with status 500 Internal Server Error.
// A generic error message, given when an unexpected condition was encountered and no more specific message is suitable.
func Http500() *Response {
	return Http500InternalServerError()
}

// NotImplemented returns a Response with status 501 Not Implemented.
// The server either does not recognize the request method, or it lacks the ability to fulfill the request.
func Http501NotImplemented() *Response {
	return &Response{status: http.StatusNotImplemented}
}

// NotImplemented returns a Response with status 501 Not Implemented.
// The server either does not recognize the request method, or it lacks the ability to fulfill the request.
func NotImplemented() *Response {
	return Http501NotImplemented()
}

// NotImplemented returns a Response with status 501 Not Implemented.
// The server either does not recognize the request method, or it lacks the ability to fulfill the request.
func Http501() *Response {
	return Http501NotImplemented()
}

// BadGateway returns a Response with status 502 Bad Gateway.
// The server was acting as a gateway or proxy and received an invalid response from the upstream server.
func Http502BadGateway() *Response {
	return &Response{status: http.StatusBadGateway}
}

// BadGateway returns a Response with status 502 Bad Gateway.
// The server was acting as a gateway or proxy and received an invalid response from the upstream server.
func BadGateway() *Response {
	return Http502BadGateway()
}

// BadGateway returns a Response with status 502 Bad Gateway.
// The server was acting as a gateway or proxy and received an invalid response from the upstream server.
func Http502() *Response {
	return Http502BadGateway()
}

// ServiceUnavailable returns a Response with status 503 Service Unavailable.
// The server is currently unavailable (because it is overloaded or down for maintenance).
func Http503ServiceUnavailable() *Response {
	return &Response{status: http.StatusServiceUnavailable}
}

// ServiceUnavailable returns a Response with status 503 Service Unavailable.
// The server is currently unavailable (because it is overloaded or down for maintenance).
func ServiceUnavailable() *Response {
	return Http503ServiceUnavailable()
}

// ServiceUnavailable returns a Response with status 503 Service Unavailable.
// The server is currently unavailable (because it is overloaded or down for maintenance).
func Http503() *Response {
	return Http503ServiceUnavailable()
}

// GatewayTimeout returns a Response with status 504 Gateway Timeout.
// The server was acting as a gateway or proxy and did not receive a timely response from the upstream server.
func Http504GatewayTimeout() *Response {
	return &Response{status: http.StatusGatewayTimeout}
}

// GatewayTimeout returns a Response with status 504 Gateway Timeout.
// The server was acting as a gateway or proxy and did not receive a timely response from the upstream server.
func GatewayTimeout() *Response {
	return Http504GatewayTimeout()
}

// GatewayTimeout returns a Response with status 504 Gateway Timeout.
// The server was acting as a gateway or proxy and did not receive a timely response from the upstream server.
func Http504() *Response {
	return Http504GatewayTimeout()
}

// HTTPVersionNotSupported returns a Response with status 505 HTTP Version Not Supported.
// The server does not support the HTTP protocol version used in the request.
func Http505HTTPVersionNotSupported() *Response {
	return &Response{status: http.StatusHTTPVersionNotSupported}
}

// HTTPVersionNotSupported returns a Response with status 505 HTTP Version Not Supported.
// The server does not support the HTTP protocol version used in the request.
func HTTPVersionNotSupported() *Response {
	return Http505HTTPVersionNotSupported()
}

// HTTPVersionNotSupported returns a Response with status 505 HTTP Version Not Supported.
// The server does not support the HTTP protocol version used in the request.
func Http505() *Response {
	return Http505HTTPVersionNotSupported()
}

// VariantAlsoNegotiates returns a Response with status 506 Variant Also Negotiates.
// Transparent content negotiation for the request results in a circular reference.
func Http506VariantAlsoNegotiates() *Response {
	return &Response{status: http.StatusVariantAlsoNegotiates}
}

// VariantAlsoNegotiates returns a Response with status 506 Variant Also Negotiates.
// Transparent content negotiation for the request results in a circular reference.
func VariantAlsoNegotiates() *Response {
	return Http506VariantAlsoNegotiates()
}

// VariantAlsoNegotiates returns a Response with status 506 Variant Also Negotiates.
// Transparent content negotiation for the request results in a circular reference.
func Http506() *Response {
	return Http506VariantAlsoNegotiates()
}

// InsufficientStorage returns a Response with status 507 Insufficient Storage.
// The server is unable to store the representation needed to complete the request.
func Http507InsufficientStorage() *Response {
	return &Response{status: http.StatusInsufficientStorage}
}

// InsufficientStorage returns a Response with status 507 Insufficient Storage.
// The server is unable to store the representation needed to complete the request.
func InsufficientStorage() *Response {
	return Http507InsufficientStorage()
}

// InsufficientStorage returns a Response with status 507 Insufficient Storage.
// The server is unable to store the representation needed to complete the request.
func Http507() *Response {
	return Http507InsufficientStorage()
}

// LoopDetected returns a Response with status 508 Loop Detected.
// The server detected an infinite loop while processing the request.
func Http508LoopDetected() *Response {
	return &Response{status: http.StatusLoopDetected}
}

// LoopDetected returns a Response with status 508 Loop Detected.
// The server detected an infinite loop while processing the request.
func LoopDetected() *Response {
	return Http508LoopDetected()
}

// LoopDetected returns a Response with status 508 Loop Detected.
// The server detected an infinite loop while processing the request.
func Http508() *Response {
	return Http508LoopDetected()
}

// NotExtended returns a Response with status 510 Not Extended.
// Further extensions to the request are required for the server to fulfill it.
func Http510NotExtended() *Response {
	return &Response{status: http.StatusNotExtended}
}

// NotExtended returns a Response with status 510 Not Extended.
// Further extensions to the request are required for the server to fulfill it.
func NotExtended() *Response {
	return Http510NotExtended()
}

// NotExtended returns a Response with status 510 Not Extended.
// Further extensions to the request are required for the server to fulfill it.
func Http510() *Response {
	return Http510NotExtended()
}

// NetworkAuthenticationRequired returns a Response with status 511 Network Authentication Required.
// The client needs to authenticate to gain network access.
func Http511NetworkAuthenticationRequired() *Response {
	return &Response{status: http.StatusNetworkAuthenticationRequired}
}

// NetworkAuthenticationRequired returns a Response with status 511 Network Authentication Required.
// The client needs to authenticate to gain network access.
func NetworkAuthenticationRequired() *Response {
	return Http511NetworkAuthenticationRequired()
}

// NetworkAuthenticationRequired returns a Response with status 511 Network Authentication Required.
// The client needs to authenticate to gain network access.
func Http511() *Response {
	return Http511NetworkAuthenticationRequired()
}
