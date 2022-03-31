//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package filegroup

import "io"

// FilesClientGetEmptyFileResponse contains the response from method FilesClient.GetEmptyFile.
type FilesClientGetEmptyFileResponse struct {
	// Body contains the streaming response.
	Body io.ReadCloser
}

// FilesClientGetFileLargeResponse contains the response from method FilesClient.GetFileLarge.
type FilesClientGetFileLargeResponse struct {
	// Body contains the streaming response.
	Body io.ReadCloser
}

// FilesClientGetFileResponse contains the response from method FilesClient.GetFile.
type FilesClientGetFileResponse struct {
	// Body contains the streaming response.
	Body io.ReadCloser
}
