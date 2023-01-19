// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func encodeDisjointSecurityResponse(response *DisjointSecurityOK, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	return nil
}

func encodeIntersectSecurityResponse(response *IntersectSecurityOK, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	return nil
}

func encodeOptionalSecurityResponse(response *OptionalSecurityOK, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	return nil
}