package types

import "strings"

type HeaderKey = string

const (
	HeaderFieldContentType     HeaderKey = "Content-Type"
	HeaderFieldContentEncoding HeaderKey = "Content-Encoding"
	HeaderFieldAccept          HeaderKey = "Accept"
	HeaderFieldAcceptEncoding  HeaderKey = "Accept-Encoding"
	HeaderFieldAcceptLanguage  HeaderKey = "Accept-Language"
	HeaderFieldAuthorization   HeaderKey = "Authorization"
)

var AllowedHeaders = strings.Join(
	[]string{
		HeaderFieldContentType,
		HeaderFieldContentEncoding,
		HeaderFieldAccept,
		HeaderFieldAcceptEncoding,
		HeaderFieldAcceptLanguage,
		HeaderFieldAuthorization,
		//KeyTenant.String(),
		//KeyThrough.String(),
		//KeySource.String(),
	},
	",",
)
