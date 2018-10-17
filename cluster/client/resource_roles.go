// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "auth": resource_roles Resource Client
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

// AssignRoleResourceRolesPayload is the resource_roles assignRole action payload.
type AssignRoleResourceRolesPayload struct {
	Data []*AssignRoleData `form:"data" json:"data" xml:"data"`
}

// AssignRoleResourceRolesPath computes a request path to the assignRole action of resource_roles.
func AssignRoleResourceRolesPath(resourceID string) string {
	param0 := resourceID

	return fmt.Sprintf("/api/resources/%s/roles", param0)
}

// Assigns roles to one or more identities, for a specific resource
func (c *Client) AssignRoleResourceRoles(ctx context.Context, path string, payload *AssignRoleResourceRolesPayload, contentType string) (*http.Response, error) {
	req, err := c.NewAssignRoleResourceRolesRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAssignRoleResourceRolesRequest create the request corresponding to the assignRole action endpoint of the resource_roles resource.
func (c *Client) NewAssignRoleResourceRolesRequest(ctx context.Context, path string, payload *AssignRoleResourceRolesPayload, contentType string) (*http.Request, error) {
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
	req, err := http.NewRequest("PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// ListAssignedResourceRolesPath computes a request path to the listAssigned action of resource_roles.
func ListAssignedResourceRolesPath(resourceID string) string {
	param0 := resourceID

	return fmt.Sprintf("/api/resources/%s/roles", param0)
}

// List assigned roles by resource
func (c *Client) ListAssignedResourceRoles(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListAssignedResourceRolesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListAssignedResourceRolesRequest create the request corresponding to the listAssigned action endpoint of the resource_roles resource.
func (c *Client) NewListAssignedResourceRolesRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// ListAssignedByRoleNameResourceRolesPath computes a request path to the listAssignedByRoleName action of resource_roles.
func ListAssignedByRoleNameResourceRolesPath(resourceID string, roleName string) string {
	param0 := resourceID
	param1 := roleName

	return fmt.Sprintf("/api/resources/%s/roles/%s", param0, param1)
}

// List assigned roles for a specific role name, for a specific resource
func (c *Client) ListAssignedByRoleNameResourceRoles(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListAssignedByRoleNameResourceRolesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListAssignedByRoleNameResourceRolesRequest create the request corresponding to the listAssignedByRoleName action endpoint of the resource_roles resource.
func (c *Client) NewListAssignedByRoleNameResourceRolesRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}
