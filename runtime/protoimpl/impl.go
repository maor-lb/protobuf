// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package protoimpl contains the default implementation for messages
// generated by protoc-gen-go.
//
// WARNING: This package should only ever be imported by generated messages.
// The compatibility agreement covers nothing except for functionality needed
// to keep existing generated messages operational. Breakages that occur due
// to unauthorized usages of this package are not the author's responsibility.
package protoimpl

import (
	"google.golang.org/protobuf/internal/fileinit"
	"google.golang.org/protobuf/internal/impl"
)

// Version is the current minor version of the package.
// This is incremented every time the API of this package expands.
const Version = 0 // v2.{Version}.x

var X impl.Export

type (
	// EnforceVersion is used by code generated by protoc-gen-go
	// to statically enforce a minimum version of this package.
	// A compilation failure implies that this package is too old and
	// needs to be updated to a more recent version.
	//
	// This package can be upgraded by running:
	//	go get -u google.golang.org/protobuf/...
	//
	// Example usage by generated code:
	//	const _ = protoimpl.EnforceVersion(protoimpl.Version - genVersion)
	//
	// If genVersion is lower than Version, then this results in a negative
	// integer overflow failure when evaluating the uint constant.
	EnforceVersion uint

	MessageType = impl.MessageType
	FileBuilder = fileinit.FileBuilder

	// TODO: Change these to more efficient data structures.
	ExtensionFields = map[int32]impl.ExtensionFieldV1
	UnknownFields   = []byte
	SizeCache       = int32

	ExtensionFieldV1 = impl.ExtensionFieldV1
)
