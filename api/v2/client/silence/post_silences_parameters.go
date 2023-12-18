// Code generated by go-swagger; DO NOT EDIT.

// Copyright Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package silence

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/tyr1k/alertmanager/api/v2/models"
)

// NewPostSilencesParams creates a new PostSilencesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostSilencesParams() *PostSilencesParams {
	return &PostSilencesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostSilencesParamsWithTimeout creates a new PostSilencesParams object
// with the ability to set a timeout on a request.
func NewPostSilencesParamsWithTimeout(timeout time.Duration) *PostSilencesParams {
	return &PostSilencesParams{
		timeout: timeout,
	}
}

// NewPostSilencesParamsWithContext creates a new PostSilencesParams object
// with the ability to set a context for a request.
func NewPostSilencesParamsWithContext(ctx context.Context) *PostSilencesParams {
	return &PostSilencesParams{
		Context: ctx,
	}
}

// NewPostSilencesParamsWithHTTPClient creates a new PostSilencesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostSilencesParamsWithHTTPClient(client *http.Client) *PostSilencesParams {
	return &PostSilencesParams{
		HTTPClient: client,
	}
}

/*
PostSilencesParams contains all the parameters to send to the API endpoint

	for the post silences operation.

	Typically these are written to a http.Request.
*/
type PostSilencesParams struct {

	/* Silence.

	   The silence to create
	*/
	Silence *models.PostableSilence

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post silences params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSilencesParams) WithDefaults() *PostSilencesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post silences params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSilencesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post silences params
func (o *PostSilencesParams) WithTimeout(timeout time.Duration) *PostSilencesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post silences params
func (o *PostSilencesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post silences params
func (o *PostSilencesParams) WithContext(ctx context.Context) *PostSilencesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post silences params
func (o *PostSilencesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post silences params
func (o *PostSilencesParams) WithHTTPClient(client *http.Client) *PostSilencesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post silences params
func (o *PostSilencesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSilence adds the silence to the post silences params
func (o *PostSilencesParams) WithSilence(silence *models.PostableSilence) *PostSilencesParams {
	o.SetSilence(silence)
	return o
}

// SetSilence adds the silence to the post silences params
func (o *PostSilencesParams) SetSilence(silence *models.PostableSilence) {
	o.Silence = silence
}

// WriteToRequest writes these params to a swagger request
func (o *PostSilencesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Silence != nil {
		if err := r.SetBodyParam(o.Silence); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
