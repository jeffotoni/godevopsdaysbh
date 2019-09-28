// Go Api server
// @jeffotoni

package cert

// openssl genrsa -out private.rsa 1024
// openssl rsa -in private.rsa -pubout > public.rsa.pub

const (
	RSA_PRIVATE = `-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDZkyyo0qZjDKqcmBy0dS0pezYAfI1JSt+pYlYlvEOSnIy3Tcog
ZNjm/jyIZwGIqM3ZuiWGYMEqlv+/gSMl1/5OmrN9IRjEKouKD+0rl6q45H7v1bPm
fEKmK7u4FjI3sGB1aaLE/L7l4/kj7r0+SqWRRG+tstIkvD58+gOsPJvvuQIDAQAB
AoGAVmt1zqy0+By826wzop30TUiqP3E5cUYzJoVvsFdiA6b44YodYdaMrNuI0akQ
W+aotPWSGBjwQOlvNfKmMk7jrDtA8k9BtW8hZB3BGsCmlu/yAUpO3aGTLu7Z0bMu
z6Mtuzc4nQk4Fp2vmH+TLuJ1MyYME9JW/JVz7UwL0T7FZE0CQQD7BmuPTeml2+To
qvKjNFLosrNwrlvAHobvLoDxC9Oupco0SdaGNCxwlwA+SThMiXywjlUb74WZsq4Y
gQzw61v/AkEA3eMLVR0hpv2Fg7rJnyzKrtCAch6cSPjQ56xQ+GX2UEr0pHzpw4tT
cL7i1uEkRr72SQIvO0UEDTDIepyK/dqURwJBAN+nZbfljDIpjHc8sDh3CxOm/Dd+
MeMnj8OVJG8fwXfO4SzbSNaDr8CU2TOlmxKeQMkx12SvcNJzH6hXesdAe9MCQQCX
keDL43p5d/UcZxwTf6V0hTk4u3BG6LDLkQo+pGf9sTYspqeIzEPTYJHT9zewnAHh
HXzmH98Eo65PE2BeyFFBAkAQzak2TdJP8yoFXnD2wUUFa8xuFwJTzNuuZod8QFuY
CoDftHE1yu1fC6Lu1+euMLFNgplMiwa9L9/bv9iSSndf
-----END RSA PRIVATE KEY-----
`

	RSA_PUBLIC = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDZkyyo0qZjDKqcmBy0dS0pezYA
fI1JSt+pYlYlvEOSnIy3TcogZNjm/jyIZwGIqM3ZuiWGYMEqlv+/gSMl1/5OmrN9
IRjEKouKD+0rl6q45H7v1bPmfEKmK7u4FjI3sGB1aaLE/L7l4/kj7r0+SqWRRG+t
stIkvD58+gOsPJvvuQIDAQAB
-----END PUBLIC KEY-----
`
)
