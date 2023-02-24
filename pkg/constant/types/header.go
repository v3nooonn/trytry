package types

import "strings"

type HeaderKey string

const (
	HeaderFieldContentType     HeaderKey = "Content-Type"
	HeaderFieldContentEncoding HeaderKey = "Content-Encoding"
	HeaderFieldAccept          HeaderKey = "Accept"
	HeaderFieldAcceptEncoding  HeaderKey = "Accept-Encoding"
	HeaderFieldAcceptLanguage  HeaderKey = "Accept-Language"
	HeaderFieldAuthorization   HeaderKey = "Authorization"
)

func (r HeaderKey) String() string {
	return string(r)
}

var AllowedHeaders = strings.Join(
	[]string{
		HeaderFieldContentType.String(),
		HeaderFieldContentEncoding.String(),
		HeaderFieldAccept.String(),
		HeaderFieldAcceptEncoding.String(),
		HeaderFieldAcceptLanguage.String(),
		HeaderFieldAuthorization.String(),
		//KeyTenant.String(),
		//KeyThrough.String(),
		//KeySource.String(),
	},
	",",
)
