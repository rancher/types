package client

import (
	"github.com/rancher/norman/types"
)

const (
	TokenType            = "token"
	TokenFieldAPIVersion = "apiVersion"
	TokenFieldIsCLI      = "isCLI"
	TokenFieldKind       = "kind"
	TokenFieldObjectMeta = "objectMeta"
	TokenFieldTokenKey   = "tokenKey"
	TokenFieldTokenValue = "tokenValue"
	TokenFieldUser       = "user"
)

type Token struct {
	types.Resource
	APIVersion string     `json:"apiVersion,omitempty"`
	IsCLI      bool       `json:"isCLI,omitempty"`
	Kind       string     `json:"kind,omitempty"`
	ObjectMeta ObjectMeta `json:"objectMeta,omitempty"`
	TokenKey   string     `json:"tokenKey,omitempty"`
	TokenValue string     `json:"tokenValue,omitempty"`
	User       string     `json:"user,omitempty"`
}
type TokenCollection struct {
	types.Collection
	Data   []Token `json:"data,omitempty"`
	client *TokenClient
}

type TokenClient struct {
	apiClient *Client
}

type TokenOperations interface {
	List(opts *types.ListOpts) (*TokenCollection, error)
	Create(opts *Token) (*Token, error)
	Update(existing *Token, updates interface{}) (*Token, error)
	ByID(id string) (*Token, error)
	Delete(container *Token) error
}

func newTokenClient(apiClient *Client) *TokenClient {
	return &TokenClient{
		apiClient: apiClient,
	}
}

func (c *TokenClient) Create(container *Token) (*Token, error) {
	resp := &Token{}
	err := c.apiClient.Ops.DoCreate(TokenType, container, resp)
	return resp, err
}

func (c *TokenClient) Update(existing *Token, updates interface{}) (*Token, error) {
	resp := &Token{}
	err := c.apiClient.Ops.DoUpdate(TokenType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *TokenClient) List(opts *types.ListOpts) (*TokenCollection, error) {
	resp := &TokenCollection{}
	err := c.apiClient.Ops.DoList(TokenType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *TokenCollection) Next() (*TokenCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &TokenCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *TokenClient) ByID(id string) (*Token, error) {
	resp := &Token{}
	err := c.apiClient.Ops.DoByID(TokenType, id, resp)
	return resp, err
}

func (c *TokenClient) Delete(container *Token) error {
	return c.apiClient.Ops.DoResourceDelete(TokenType, &container.Resource)
}
