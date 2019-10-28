// Code generated by goa v3.0.6, DO NOT EDIT.
//
// SSAServer client HTTP transport
//
// Command:
// $ goa gen SSAServer/design

package client

import (
	ssaserver "SSAServer/gen/ssa_server"
	"context"
	"mime/multipart"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the SSAServer service endpoint HTTP clients.
type Client struct {
	// Register Doer is the HTTP client used to make requests to the Register
	// endpoint.
	RegisterDoer goahttp.Doer

	// Login Doer is the HTTP client used to make requests to the Login endpoint.
	LoginDoer goahttp.Doer

	// ChangeGroup Doer is the HTTP client used to make requests to the
	// Change_group endpoint.
	ChangeGroupDoer goahttp.Doer

	// DeleteUser Doer is the HTTP client used to make requests to the Delete_user
	// endpoint.
	DeleteUserDoer goahttp.Doer

	// SaveData Doer is the HTTP client used to make requests to the Save_data
	// endpoint.
	SaveDataDoer goahttp.Doer

	// ReturnDataList Doer is the HTTP client used to make requests to the
	// Return_data_list endpoint.
	ReturnDataListDoer goahttp.Doer

	// PickUpData Doer is the HTTP client used to make requests to the Pick_up_data
	// endpoint.
	PickUpDataDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// SSAServerSaveDataEncoderFunc is the type to encode multipart request for the
// "SSAServer" service "Save_data" endpoint.
type SSAServerSaveDataEncoderFunc func(*multipart.Writer, *ssaserver.SaveDataPayload) error

// SSAServerPickUpDataEncoderFunc is the type to encode multipart request for
// the "SSAServer" service "Pick_up_data" endpoint.
type SSAServerPickUpDataEncoderFunc func(*multipart.Writer, *ssaserver.PickUpDataPayload) error

// NewClient instantiates HTTP clients for all the SSAServer service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		RegisterDoer:        doer,
		LoginDoer:           doer,
		ChangeGroupDoer:     doer,
		DeleteUserDoer:      doer,
		SaveDataDoer:        doer,
		ReturnDataListDoer:  doer,
		PickUpDataDoer:      doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Register returns an endpoint that makes HTTP requests to the SSAServer
// service Register server.
func (c *Client) Register() goa.Endpoint {
	var (
		encodeRequest  = EncodeRegisterRequest(c.encoder)
		decodeResponse = DecodeRegisterResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildRegisterRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RegisterDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("SSAServer", "Register", err)
		}
		return decodeResponse(resp)
	}
}

// Login returns an endpoint that makes HTTP requests to the SSAServer service
// Login server.
func (c *Client) Login() goa.Endpoint {
	var (
		encodeRequest  = EncodeLoginRequest(c.encoder)
		decodeResponse = DecodeLoginResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildLoginRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.LoginDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("SSAServer", "Login", err)
		}
		return decodeResponse(resp)
	}
}

// ChangeGroup returns an endpoint that makes HTTP requests to the SSAServer
// service Change_group server.
func (c *Client) ChangeGroup() goa.Endpoint {
	var (
		encodeRequest  = EncodeChangeGroupRequest(c.encoder)
		decodeResponse = DecodeChangeGroupResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildChangeGroupRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ChangeGroupDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("SSAServer", "Change_group", err)
		}
		return decodeResponse(resp)
	}
}

// DeleteUser returns an endpoint that makes HTTP requests to the SSAServer
// service Delete_user server.
func (c *Client) DeleteUser() goa.Endpoint {
	var (
		encodeRequest  = EncodeDeleteUserRequest(c.encoder)
		decodeResponse = DecodeDeleteUserResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeleteUserRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteUserDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("SSAServer", "Delete_user", err)
		}
		return decodeResponse(resp)
	}
}

// SaveData returns an endpoint that makes HTTP requests to the SSAServer
// service Save_data server.
func (c *Client) SaveData(sSAServerSaveDataEncoderFn SSAServerSaveDataEncoderFunc) goa.Endpoint {
	var (
		encodeRequest  = EncodeSaveDataRequest(NewSSAServerSaveDataEncoder(sSAServerSaveDataEncoderFn))
		decodeResponse = DecodeSaveDataResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildSaveDataRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.SaveDataDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("SSAServer", "Save_data", err)
		}
		return decodeResponse(resp)
	}
}

// ReturnDataList returns an endpoint that makes HTTP requests to the SSAServer
// service Return_data_list server.
func (c *Client) ReturnDataList() goa.Endpoint {
	var (
		encodeRequest  = EncodeReturnDataListRequest(c.encoder)
		decodeResponse = DecodeReturnDataListResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildReturnDataListRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ReturnDataListDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("SSAServer", "Return_data_list", err)
		}
		return decodeResponse(resp)
	}
}

// PickUpData returns an endpoint that makes HTTP requests to the SSAServer
// service Pick_up_data server.
func (c *Client) PickUpData(sSAServerPickUpDataEncoderFn SSAServerPickUpDataEncoderFunc) goa.Endpoint {
	var (
		encodeRequest  = EncodePickUpDataRequest(NewSSAServerPickUpDataEncoder(sSAServerPickUpDataEncoderFn))
		decodeResponse = DecodePickUpDataResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildPickUpDataRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.PickUpDataDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("SSAServer", "Pick_up_data", err)
		}
		return decodeResponse(resp)
	}
}
