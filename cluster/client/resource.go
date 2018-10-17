// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "auth": resource Resource Client
//
// Command:
// $ goagen
// --design=github.com/fabric8-services/fabric8-cluster/design
// --out=$(GOPATH)/src/github.com/fabric8-services/fabric8-cluster-client/cluster
// --pkg=client
// --version=v1.3.0

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// DeleteResourcePath computes a request path to the delete action of resource.
func DeleteResourcePath(resourceID string) string {
	param0 := resourceID

	return fmt.Sprintf("/api/resource/%s", param0)
}

// Delete a resource
func (c *Client) DeleteResource(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteResourceRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteResourceRequest create the request corresponding to the delete action endpoint of the resource resource.
func (c *Client) NewDeleteResourceRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ReadResourcePath computes a request path to the read action of resource.
func ReadResourcePath(resourceID string) string {
	param0 := resourceID

	return fmt.Sprintf("/api/resource/%s", param0)
}

// Read a specific resource
func (c *Client) ReadResource(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewReadResourceRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewReadResourceRequest create the request corresponding to the read action endpoint of the resource resource.
func (c *Client) NewReadResourceRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// RegisterResourcePayload is the resource register action payload.
type RegisterResourcePayload struct {
	// The parent resource (of the same type) to which this resource belongs
	ParentResourceID *string `form:"parent_resource_id,omitempty" json:"parent_resource_id,omitempty" xml:"parent_resource_id,omitempty"`
	// The identifier for this resource. If left blank, one will be generated
	ResourceID *string `form:"resource_id,omitempty" json:"resource_id,omitempty" xml:"resource_id,omitempty"`
	// The type of resource
	Type string `form:"type" json:"type" xml:"type"`
}

// RegisterResourcePath computes a request path to the register action of resource.
func RegisterResourcePath() string {

	return fmt.Sprintf("/api/resource")
}

// Register a new resource
func (c *Client) RegisterResource(ctx context.Context, path string, payload *RegisterResourcePayload, contentType string) (*http.Response, error) {
	req, err := c.NewRegisterResourceRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewRegisterResourceRequest create the request corresponding to the register action endpoint of the resource resource.
func (c *Client) NewRegisterResourceRequest(ctx context.Context, path string, payload *RegisterResourcePayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}
