package util

type HttpHeaders map[string]string

var JsonContentDefaultHeader = HttpHeaders{
	"Content-Type": "application/json",
}
