package stdlib

// Generated by 'goexports net/http/httptrace'. Do not edit!

import (
	"net/http/httptrace"
	"reflect"
)

func init() {
	Value["net/http/httptrace"] = map[string]reflect.Value{
		"ContextClientTrace": reflect.ValueOf(httptrace.ContextClientTrace),
		"WithClientTrace":    reflect.ValueOf(httptrace.WithClientTrace),
	}
	Type["net/http/httptrace"] = map[string]reflect.Type{
		"ClientTrace":      reflect.TypeOf((*httptrace.ClientTrace)(nil)).Elem(),
		"DNSDoneInfo":      reflect.TypeOf((*httptrace.DNSDoneInfo)(nil)).Elem(),
		"DNSStartInfo":     reflect.TypeOf((*httptrace.DNSStartInfo)(nil)).Elem(),
		"GotConnInfo":      reflect.TypeOf((*httptrace.GotConnInfo)(nil)).Elem(),
		"WroteRequestInfo": reflect.TypeOf((*httptrace.WroteRequestInfo)(nil)).Elem(),
	}
}