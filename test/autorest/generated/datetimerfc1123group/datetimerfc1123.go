// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package datetimerfc1123group

import (
	"context"
	azinternal "generatortests/autorest/generated/datetimerfc1123group/internal/datetimerfc1123group"
	"time"
)

// Datetimerfc1123Operations contains the methods for the Datetimerfc1123 group.
type Datetimerfc1123Operations interface {
	// GetInvalid - Get invalid datetime value
	GetInvalid(ctx context.Context) (*Datetimerfc1123GetInvalidResponse, error)
	// GetNull - Get null datetime value
	GetNull(ctx context.Context) (*Datetimerfc1123GetNullResponse, error)
	// GetOverflow - Get overflow datetime value
	GetOverflow(ctx context.Context) (*Datetimerfc1123GetOverflowResponse, error)
	// GetUTCLowercaseMaxDateTime - Get max datetime value fri, 31 dec 9999 23:59:59 gmt
	GetUTCLowercaseMaxDateTime(ctx context.Context) (*Datetimerfc1123GetUTCLowercaseMaxDateTimeResponse, error)
	// GetUTCMinDateTime - Get min datetime value Mon, 1 Jan 0001 00:00:00 GMT
	GetUTCMinDateTime(ctx context.Context) (*Datetimerfc1123GetUTCMinDateTimeResponse, error)
	// GetUTCUppercaseMaxDateTime - Get max datetime value FRI, 31 DEC 9999 23:59:59 GMT
	GetUTCUppercaseMaxDateTime(ctx context.Context) (*Datetimerfc1123GetUTCUppercaseMaxDateTimeResponse, error)
	// GetUnderflow - Get underflow datetime value
	GetUnderflow(ctx context.Context) (*Datetimerfc1123GetUnderflowResponse, error)
	// PutUTCMaxDateTime - Put max datetime value Fri, 31 Dec 9999 23:59:59 GMT
	PutUTCMaxDateTime(ctx context.Context, datetimeBody time.Time) (*Datetimerfc1123PutUTCMaxDateTimeResponse, error)
	// PutUTCMinDateTime - Put min datetime value Mon, 1 Jan 0001 00:00:00 GMT
	PutUTCMinDateTime(ctx context.Context, datetimeBody time.Time) (*Datetimerfc1123PutUTCMinDateTimeResponse, error)
}

type datetimerfc1123Operations struct {
	*Client
	azinternal.Datetimerfc1123Operations
}

// GetInvalid - Get invalid datetime value
func (client *datetimerfc1123Operations) GetInvalid(ctx context.Context) (*Datetimerfc1123GetInvalidResponse, error) {
	req, err := client.GetInvalidCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetInvalidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetNull - Get null datetime value
func (client *datetimerfc1123Operations) GetNull(ctx context.Context) (*Datetimerfc1123GetNullResponse, error) {
	req, err := client.GetNullCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetOverflow - Get overflow datetime value
func (client *datetimerfc1123Operations) GetOverflow(ctx context.Context) (*Datetimerfc1123GetOverflowResponse, error) {
	req, err := client.GetOverflowCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetOverflowHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUTCLowercaseMaxDateTime - Get max datetime value fri, 31 dec 9999 23:59:59 gmt
func (client *datetimerfc1123Operations) GetUTCLowercaseMaxDateTime(ctx context.Context) (*Datetimerfc1123GetUTCLowercaseMaxDateTimeResponse, error) {
	req, err := client.GetUTCLowercaseMaxDateTimeCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetUTCLowercaseMaxDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUTCMinDateTime - Get min datetime value Mon, 1 Jan 0001 00:00:00 GMT
func (client *datetimerfc1123Operations) GetUTCMinDateTime(ctx context.Context) (*Datetimerfc1123GetUTCMinDateTimeResponse, error) {
	req, err := client.GetUTCMinDateTimeCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetUTCMinDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUTCUppercaseMaxDateTime - Get max datetime value FRI, 31 DEC 9999 23:59:59 GMT
func (client *datetimerfc1123Operations) GetUTCUppercaseMaxDateTime(ctx context.Context) (*Datetimerfc1123GetUTCUppercaseMaxDateTimeResponse, error) {
	req, err := client.GetUTCUppercaseMaxDateTimeCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetUTCUppercaseMaxDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUnderflow - Get underflow datetime value
func (client *datetimerfc1123Operations) GetUnderflow(ctx context.Context) (*Datetimerfc1123GetUnderflowResponse, error) {
	req, err := client.GetUnderflowCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetUnderflowHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PutUTCMaxDateTime - Put max datetime value Fri, 31 Dec 9999 23:59:59 GMT
func (client *datetimerfc1123Operations) PutUTCMaxDateTime(ctx context.Context, datetimeBody time.Time) (*Datetimerfc1123PutUTCMaxDateTimeResponse, error) {
	req, err := client.PutUTCMaxDateTimeCreateRequest(*client.u, datetimeBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.PutUTCMaxDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PutUTCMinDateTime - Put min datetime value Mon, 1 Jan 0001 00:00:00 GMT
func (client *datetimerfc1123Operations) PutUTCMinDateTime(ctx context.Context, datetimeBody time.Time) (*Datetimerfc1123PutUTCMinDateTimeResponse, error) {
	req, err := client.PutUTCMinDateTimeCreateRequest(*client.u, datetimeBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.PutUTCMinDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

var _ Datetimerfc1123Operations = (*datetimerfc1123Operations)(nil)
