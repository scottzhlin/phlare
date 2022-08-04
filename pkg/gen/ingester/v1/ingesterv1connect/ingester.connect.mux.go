// Code generated by protoc-gen-connect-go-mux. DO NOT EDIT.
//
// Source: ingester/v1/ingester.proto

package ingesterv1connect

import (
	connect_go "github.com/bufbuild/connect-go"
	mux "github.com/gorilla/mux"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

// RegisterIngesterServiceHandler register an HTTP handler to a mux.Router from the service
// implementation.
func RegisterIngesterServiceHandler(mux *mux.Router, svc IngesterServiceHandler, opts ...connect_go.HandlerOption) {
	mux.Handle("/ingester.v1.IngesterService/Push", connect_go.NewUnaryHandler(
		"/ingester.v1.IngesterService/Push",
		svc.Push,
		opts...,
	))
	mux.Handle("/ingester.v1.IngesterService/LabelValues", connect_go.NewUnaryHandler(
		"/ingester.v1.IngesterService/LabelValues",
		svc.LabelValues,
		opts...,
	))
	mux.Handle("/ingester.v1.IngesterService/ProfileTypes", connect_go.NewUnaryHandler(
		"/ingester.v1.IngesterService/ProfileTypes",
		svc.ProfileTypes,
		opts...,
	))
	mux.Handle("/ingester.v1.IngesterService/Series", connect_go.NewUnaryHandler(
		"/ingester.v1.IngesterService/Series",
		svc.Series,
		opts...,
	))
	mux.Handle("/ingester.v1.IngesterService/Flush", connect_go.NewUnaryHandler(
		"/ingester.v1.IngesterService/Flush",
		svc.Flush,
		opts...,
	))
	mux.Handle("/ingester.v1.IngesterService/SelectProfiles", connect_go.NewUnaryHandler(
		"/ingester.v1.IngesterService/SelectProfiles",
		svc.SelectProfiles,
		opts...,
	))
}
