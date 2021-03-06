// Code generated by go-swagger; DO NOT EDIT.

package account_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new account manager API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for account manager API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ChangePassword change password API
*/
func (a *Client) ChangePassword(params *ChangePasswordParams, authInfo runtime.ClientAuthInfoWriter) (*ChangePasswordOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangePasswordParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ChangePassword",
		Method:             "POST",
		PathPattern:        "/v1/users/password:change",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangePasswordReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChangePasswordOK), nil

}

/*
CreateGroup groups
*/
func (a *Client) CreateGroup(params *CreateGroupParams, authInfo runtime.ClientAuthInfoWriter) (*CreateGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateGroup",
		Method:             "POST",
		PathPattern:        "/v1/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateGroupOK), nil

}

/*
CreatePasswordReset create password reset API
*/
func (a *Client) CreatePasswordReset(params *CreatePasswordResetParams, authInfo runtime.ClientAuthInfoWriter) (*CreatePasswordResetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePasswordResetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreatePasswordReset",
		Method:             "POST",
		PathPattern:        "/v1/users/password:reset",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreatePasswordResetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreatePasswordResetOK), nil

}

/*
CreateUser admins permission
*/
func (a *Client) CreateUser(params *CreateUserParams, authInfo runtime.ClientAuthInfoWriter) (*CreateUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateUser",
		Method:             "POST",
		PathPattern:        "/v1/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateUserOK), nil

}

/*
DeleteGroups delete groups API
*/
func (a *Client) DeleteGroups(params *DeleteGroupsParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteGroups",
		Method:             "DELETE",
		PathPattern:        "/v1/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteGroupsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteGroupsOK), nil

}

/*
DeleteUsers delete users API
*/
func (a *Client) DeleteUsers(params *DeleteUsersParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUsersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteUsers",
		Method:             "DELETE",
		PathPattern:        "/v1/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteUsersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteUsersOK), nil

}

/*
DescribeGroups describe groups API
*/
func (a *Client) DescribeGroups(params *DescribeGroupsParams, authInfo runtime.ClientAuthInfoWriter) (*DescribeGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DescribeGroups",
		Method:             "GET",
		PathPattern:        "/v1/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DescribeGroupsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DescribeGroupsOK), nil

}

/*
DescribeUsers describe users API
*/
func (a *Client) DescribeUsers(params *DescribeUsersParams, authInfo runtime.ClientAuthInfoWriter) (*DescribeUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeUsersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DescribeUsers",
		Method:             "GET",
		PathPattern:        "/v1/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DescribeUsersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DescribeUsersOK), nil

}

/*
GetPasswordReset get password reset API
*/
func (a *Client) GetPasswordReset(params *GetPasswordResetParams, authInfo runtime.ClientAuthInfoWriter) (*GetPasswordResetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPasswordResetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPasswordReset",
		Method:             "GET",
		PathPattern:        "/v1/users/password/reset/{reset_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPasswordResetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPasswordResetOK), nil

}

/*
JoinGroup join group API
*/
func (a *Client) JoinGroup(params *JoinGroupParams, authInfo runtime.ClientAuthInfoWriter) (*JoinGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewJoinGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "JoinGroup",
		Method:             "POST",
		PathPattern:        "/v1/groups:join",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &JoinGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*JoinGroupOK), nil

}

/*
LeaveGroup leave group API
*/
func (a *Client) LeaveGroup(params *LeaveGroupParams, authInfo runtime.ClientAuthInfoWriter) (*LeaveGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLeaveGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LeaveGroup",
		Method:             "POST",
		PathPattern:        "/v1/groups:leave",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LeaveGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LeaveGroupOK), nil

}

/*
ModifyGroup modify group API
*/
func (a *Client) ModifyGroup(params *ModifyGroupParams, authInfo runtime.ClientAuthInfoWriter) (*ModifyGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ModifyGroup",
		Method:             "PATCH",
		PathPattern:        "/v1/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ModifyGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifyGroupOK), nil

}

/*
ModifyUser modify user API
*/
func (a *Client) ModifyUser(params *ModifyUserParams, authInfo runtime.ClientAuthInfoWriter) (*ModifyUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ModifyUser",
		Method:             "PATCH",
		PathPattern:        "/v1/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ModifyUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifyUserOK), nil

}

/*
ValidateUserPassword validate user password API
*/
func (a *Client) ValidateUserPassword(params *ValidateUserPasswordParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateUserPasswordOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateUserPasswordParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ValidateUserPassword",
		Method:             "POST",
		PathPattern:        "/v1/users/password:validate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ValidateUserPasswordReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ValidateUserPasswordOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
