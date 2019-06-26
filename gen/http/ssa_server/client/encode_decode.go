// Code generated by goa v3.0.2, DO NOT EDIT.
//
// SSAServer HTTP client encoders and decoders
//
// Command:
// $ goa gen SSAServer/design

package client

import (
	ssaserver "SSAServer/gen/ssa_server"
	ssaserverviews "SSAServer/gen/ssa_server/views"
	"bytes"
	"context"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildRegisterRequest instantiates a HTTP request object with method and path
// set to call the "SSAServer" service "Register" endpoint
func (c *Client) BuildRegisterRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RegisterSSAServerPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("SSAServer", "Register", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeRegisterRequest returns an encoder for requests sent to the SSAServer
// Register server.
func EncodeRegisterRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*ssaserver.RegisterPayload)
		if !ok {
			return goahttp.ErrInvalidType("SSAServer", "Register", "*ssaserver.RegisterPayload", v)
		}
		body := NewRegisterRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("SSAServer", "Register", err)
		}
		return nil
	}
}

// DecodeRegisterResponse returns a decoder for responses returned by the
// SSAServer Register endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeRegisterResponse may return the following errors:
//	- "Invalid_Group_ID" (type *ssaserver.SsaError): http.StatusNotFound
//	- "The_user_already_exists" (type *ssaserver.SsaError): http.StatusBadRequest
//	- "Invalid_Request" (type *ssaserver.SsaError): http.StatusBadRequest
//	- error: internal error
func DecodeRegisterResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body RegisterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("SSAServer", "Register", err)
			}
			p := NewRegisterSsaResultOK(&body)
			view := "extended"
			vres := &ssaserverviews.SsaResult{p, view}
			if err = ssaserverviews.ValidateSsaResult(vres); err != nil {
				return nil, goahttp.ErrValidationError("SSAServer", "Register", err)
			}
			res := ssaserver.NewSsaResult(vres)
			return res, nil
		case http.StatusNotFound:
			var (
				body RegisterInvalidGroupIDResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("SSAServer", "Register", err)
			}
			err = ValidateRegisterInvalidGroupIDResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("SSAServer", "Register", err)
			}
			return nil, NewRegisterInvalidGroupID(&body)
		case http.StatusBadRequest:
			en := resp.Header.Get("goa-error")
			switch en {
			case "The_user_already_exists":
				var (
					body RegisterTheUserAlreadyExistsResponseBody
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("SSAServer", "Register", err)
				}
				err = ValidateRegisterTheUserAlreadyExistsResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("SSAServer", "Register", err)
				}
				return nil, NewRegisterTheUserAlreadyExists(&body)
			case "Invalid_Request":
				var (
					body RegisterInvalidRequestResponseBody
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("SSAServer", "Register", err)
				}
				err = ValidateRegisterInvalidRequestResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("SSAServer", "Register", err)
				}
				return nil, NewRegisterInvalidRequest(&body)
			default:
				body, _ := ioutil.ReadAll(resp.Body)
				return nil, goahttp.ErrInvalidResponse("SSAServer", "Register", resp.StatusCode, string(body))
			}
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("SSAServer", "Register", resp.StatusCode, string(body))
		}
	}
}

// BuildLoginRequest instantiates a HTTP request object with method and path
// set to call the "SSAServer" service "Login" endpoint
func (c *Client) BuildLoginRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: LoginSSAServerPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("SSAServer", "Login", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeLoginRequest returns an encoder for requests sent to the SSAServer
// Login server.
func EncodeLoginRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*ssaserver.LoginPayload)
		if !ok {
			return goahttp.ErrInvalidType("SSAServer", "Login", "*ssaserver.LoginPayload", v)
		}
		body := NewLoginRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("SSAServer", "Login", err)
		}
		return nil
	}
}

// DecodeLoginResponse returns a decoder for responses returned by the
// SSAServer Login endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeLoginResponse may return the following errors:
//	- "Invalid_Request" (type *ssaserver.SsaError): http.StatusBadRequest
//	- "Invalid_Group_ID" (type *ssaserver.SsaError): http.StatusNotFound
//	- error: internal error
func DecodeLoginResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body bool
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("SSAServer", "Login", err)
			}
			return body, nil
		case http.StatusBadRequest:
			var (
				body LoginInvalidRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("SSAServer", "Login", err)
			}
			err = ValidateLoginInvalidRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("SSAServer", "Login", err)
			}
			return nil, NewLoginInvalidRequest(&body)
		case http.StatusNotFound:
			var (
				body LoginInvalidGroupIDResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("SSAServer", "Login", err)
			}
			err = ValidateLoginInvalidGroupIDResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("SSAServer", "Login", err)
			}
			return nil, NewLoginInvalidGroupID(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("SSAServer", "Login", resp.StatusCode, string(body))
		}
	}
}

// BuildChangeGroupRequest instantiates a HTTP request object with method and
// path set to call the "SSAServer" service "Change_group" endpoint
func (c *Client) BuildChangeGroupRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		userID int
	)
	{
		p, ok := v.(*ssaserver.ChangeGroupPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("SSAServer", "Change_group", "*ssaserver.ChangeGroupPayload", v)
		}
		userID = p.UserID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ChangeGroupSSAServerPath(userID)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("SSAServer", "Change_group", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeChangeGroupRequest returns an encoder for requests sent to the
// SSAServer Change_group server.
func EncodeChangeGroupRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*ssaserver.ChangeGroupPayload)
		if !ok {
			return goahttp.ErrInvalidType("SSAServer", "Change_group", "*ssaserver.ChangeGroupPayload", v)
		}
		body := NewChangeGroupRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("SSAServer", "Change_group", err)
		}
		return nil
	}
}

// DecodeChangeGroupResponse returns a decoder for responses returned by the
// SSAServer Change_group endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeChangeGroupResponse may return the following errors:
//	- "Invalid_Group_ID" (type *ssaserver.SsaError): http.StatusBadRequest
//	- "Invalid_Request" (type *ssaserver.SsaError): http.StatusBadRequest
//	- error: internal error
func DecodeChangeGroupResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body bool
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("SSAServer", "Change_group", err)
			}
			return body, nil
		case http.StatusBadRequest:
			en := resp.Header.Get("goa-error")
			switch en {
			case "Invalid_Group_ID":
				var (
					body ChangeGroupInvalidGroupIDResponseBody
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("SSAServer", "Change_group", err)
				}
				err = ValidateChangeGroupInvalidGroupIDResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("SSAServer", "Change_group", err)
				}
				return nil, NewChangeGroupInvalidGroupID(&body)
			case "Invalid_Request":
				var (
					body ChangeGroupInvalidRequestResponseBody
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("SSAServer", "Change_group", err)
				}
				err = ValidateChangeGroupInvalidRequestResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("SSAServer", "Change_group", err)
				}
				return nil, NewChangeGroupInvalidRequest(&body)
			default:
				body, _ := ioutil.ReadAll(resp.Body)
				return nil, goahttp.ErrInvalidResponse("SSAServer", "Change_group", resp.StatusCode, string(body))
			}
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("SSAServer", "Change_group", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteUserRequest instantiates a HTTP request object with method and
// path set to call the "SSAServer" service "Delete_user" endpoint
func (c *Client) BuildDeleteUserRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		userID int
	)
	{
		p, ok := v.(*ssaserver.DeleteUserPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("SSAServer", "Delete_user", "*ssaserver.DeleteUserPayload", v)
		}
		userID = p.UserID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteUserSSAServerPath(userID)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("SSAServer", "Delete_user", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteUserRequest returns an encoder for requests sent to the
// SSAServer Delete_user server.
func EncodeDeleteUserRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*ssaserver.DeleteUserPayload)
		if !ok {
			return goahttp.ErrInvalidType("SSAServer", "Delete_user", "*ssaserver.DeleteUserPayload", v)
		}
		body := NewDeleteUserRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("SSAServer", "Delete_user", err)
		}
		return nil
	}
}

// DecodeDeleteUserResponse returns a decoder for responses returned by the
// SSAServer Delete_user endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeDeleteUserResponse may return the following errors:
//	- "Invalid_Request" (type *ssaserver.SsaError): http.StatusBadRequest
//	- error: internal error
func DecodeDeleteUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body bool
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("SSAServer", "Delete_user", err)
			}
			return body, nil
		case http.StatusBadRequest:
			var (
				body DeleteUserInvalidRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("SSAServer", "Delete_user", err)
			}
			err = ValidateDeleteUserInvalidRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("SSAServer", "Delete_user", err)
			}
			return nil, NewDeleteUserInvalidRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("SSAServer", "Delete_user", resp.StatusCode, string(body))
		}
	}
}

// BuildSaveDataRequest instantiates a HTTP request object with method and path
// set to call the "SSAServer" service "Save_data" endpoint
func (c *Client) BuildSaveDataRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		groupID string
	)
	{
		p, ok := v.(*ssaserver.SaveDataPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("SSAServer", "Save_data", "*ssaserver.SaveDataPayload", v)
		}
		groupID = p.GroupID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SaveDataSSAServerPath(groupID)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("SSAServer", "Save_data", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSaveDataRequest returns an encoder for requests sent to the SSAServer
// Save_data server.
func EncodeSaveDataRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*ssaserver.SaveDataPayload)
		if !ok {
			return goahttp.ErrInvalidType("SSAServer", "Save_data", "*ssaserver.SaveDataPayload", v)
		}
		if err := encoder(req).Encode(p); err != nil {
			return goahttp.ErrEncodingError("SSAServer", "Save_data", err)
		}
		return nil
	}
}

// NewSSAServerSaveDataEncoder returns an encoder to encode the multipart
// request for the "SSAServer" service "Save_data" endpoint.
func NewSSAServerSaveDataEncoder(encoderFn SSAServerSaveDataEncoderFunc) func(r *http.Request) goahttp.Encoder {
	return func(r *http.Request) goahttp.Encoder {
		body := &bytes.Buffer{}
		mw := multipart.NewWriter(body)
		return goahttp.EncodingFunc(func(v interface{}) error {
			p := v.(*ssaserver.SaveDataPayload)
			if err := encoderFn(mw, p); err != nil {
				return err
			}
			r.Body = ioutil.NopCloser(body)
			r.Header.Set("Content-Type", mw.FormDataContentType())
			return mw.Close()
		})
	}
}

// DecodeSaveDataResponse returns a decoder for responses returned by the
// SSAServer Save_data endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeSaveDataResponse may return the following errors:
//	- "Invalid_Group_ID" (type *ssaserver.SsaError): http.StatusNotFound
//	- "Invalid_Request" (type *ssaserver.SsaError): http.StatusBadRequest
//	- "Invalid_Data" (type *ssaserver.SsaError): http.StatusBadRequest
//	- error: internal error
func DecodeSaveDataResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body bool
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("SSAServer", "Save_data", err)
			}
			return body, nil
		case http.StatusNotFound:
			var (
				body SaveDataInvalidGroupIDResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("SSAServer", "Save_data", err)
			}
			err = ValidateSaveDataInvalidGroupIDResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("SSAServer", "Save_data", err)
			}
			return nil, NewSaveDataInvalidGroupID(&body)
		case http.StatusBadRequest:
			en := resp.Header.Get("goa-error")
			switch en {
			case "Invalid_Request":
				var (
					body SaveDataInvalidRequestResponseBody
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("SSAServer", "Save_data", err)
				}
				err = ValidateSaveDataInvalidRequestResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("SSAServer", "Save_data", err)
				}
				return nil, NewSaveDataInvalidRequest(&body)
			case "Invalid_Data":
				var (
					body SaveDataInvalidDataResponseBody
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("SSAServer", "Save_data", err)
				}
				err = ValidateSaveDataInvalidDataResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("SSAServer", "Save_data", err)
				}
				return nil, NewSaveDataInvalidData(&body)
			default:
				body, _ := ioutil.ReadAll(resp.Body)
				return nil, goahttp.ErrInvalidResponse("SSAServer", "Save_data", resp.StatusCode, string(body))
			}
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("SSAServer", "Save_data", resp.StatusCode, string(body))
		}
	}
}

// BuildReturnDataListRequest instantiates a HTTP request object with method
// and path set to call the "SSAServer" service "Return_data_list" endpoint
func (c *Client) BuildReturnDataListRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		groupID string
	)
	{
		p, ok := v.(*ssaserver.ReturnDataListPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("SSAServer", "Return_data_list", "*ssaserver.ReturnDataListPayload", v)
		}
		groupID = p.GroupID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ReturnDataListSSAServerPath(groupID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("SSAServer", "Return_data_list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeReturnDataListRequest returns an encoder for requests sent to the
// SSAServer Return_data_list server.
func EncodeReturnDataListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*ssaserver.ReturnDataListPayload)
		if !ok {
			return goahttp.ErrInvalidType("SSAServer", "Return_data_list", "*ssaserver.ReturnDataListPayload", v)
		}
		body := NewReturnDataListRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("SSAServer", "Return_data_list", err)
		}
		return nil
	}
}

// DecodeReturnDataListResponse returns a decoder for responses returned by the
// SSAServer Return_data_list endpoint. restoreBody controls whether the
// response body should be restored after having been read.
// DecodeReturnDataListResponse may return the following errors:
//	- "Invalid_Group_ID" (type *ssaserver.SsaError): http.StatusNotFound
//	- "Invalid_Request" (type *ssaserver.SsaError): http.StatusBadRequest
//	- error: internal error
func DecodeReturnDataListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ReturnDataListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("SSAServer", "Return_data_list", err)
			}
			p := NewReturnDataListSsaResultCollectionOK(body)
			view := resp.Header.Get("goa-view")
			vres := ssaserverviews.SsaResultCollection{p, view}
			if err = ssaserverviews.ValidateSsaResultCollection(vres); err != nil {
				return nil, goahttp.ErrValidationError("SSAServer", "Return_data_list", err)
			}
			res := ssaserver.NewSsaResultCollection(vres)
			return res, nil
		case http.StatusNotFound:
			var (
				body ReturnDataListInvalidGroupIDResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("SSAServer", "Return_data_list", err)
			}
			err = ValidateReturnDataListInvalidGroupIDResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("SSAServer", "Return_data_list", err)
			}
			return nil, NewReturnDataListInvalidGroupID(&body)
		case http.StatusBadRequest:
			var (
				body ReturnDataListInvalidRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("SSAServer", "Return_data_list", err)
			}
			err = ValidateReturnDataListInvalidRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("SSAServer", "Return_data_list", err)
			}
			return nil, NewReturnDataListInvalidRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("SSAServer", "Return_data_list", resp.StatusCode, string(body))
		}
	}
}

// BuildPickUpDataRequest instantiates a HTTP request object with method and
// path set to call the "SSAServer" service "Pick_up_data" endpoint
func (c *Client) BuildPickUpDataRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		groupID  string
		dataType string
	)
	{
		p, ok := v.(*ssaserver.PickUpDataPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("SSAServer", "Pick_up_data", "*ssaserver.PickUpDataPayload", v)
		}
		groupID = p.GroupID
		if p.DataType != nil {
			dataType = *p.DataType
		}
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: PickUpDataSSAServerPath(groupID, dataType)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("SSAServer", "Pick_up_data", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodePickUpDataRequest returns an encoder for requests sent to the
// SSAServer Pick_up_data server.
func EncodePickUpDataRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*ssaserver.PickUpDataPayload)
		if !ok {
			return goahttp.ErrInvalidType("SSAServer", "Pick_up_data", "*ssaserver.PickUpDataPayload", v)
		}
		if err := encoder(req).Encode(p); err != nil {
			return goahttp.ErrEncodingError("SSAServer", "Pick_up_data", err)
		}
		return nil
	}
}

// NewSSAServerPickUpDataEncoder returns an encoder to encode the multipart
// request for the "SSAServer" service "Pick_up_data" endpoint.
func NewSSAServerPickUpDataEncoder(encoderFn SSAServerPickUpDataEncoderFunc) func(r *http.Request) goahttp.Encoder {
	return func(r *http.Request) goahttp.Encoder {
		body := &bytes.Buffer{}
		mw := multipart.NewWriter(body)
		return goahttp.EncodingFunc(func(v interface{}) error {
			p := v.(*ssaserver.PickUpDataPayload)
			if err := encoderFn(mw, p); err != nil {
				return err
			}
			r.Body = ioutil.NopCloser(body)
			r.Header.Set("Content-Type", mw.FormDataContentType())
			return mw.Close()
		})
	}
}

// DecodePickUpDataResponse returns a decoder for responses returned by the
// SSAServer Pick_up_data endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodePickUpDataResponse may return the following errors:
//	- "Invalid_Group_ID" (type *ssaserver.SsaError): http.StatusNotFound
//	- "Invalid_Request" (type *ssaserver.SsaError): http.StatusBadRequest
//	- "Invalid_data_name" (type *ssaserver.SsaError): http.StatusBadRequest
//	- error: internal error
func DecodePickUpDataResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body PickUpDataResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("SSAServer", "Pick_up_data", err)
			}
			p := NewPickUpDataSsaResultOK(&body)
			view := "extended"
			vres := &ssaserverviews.SsaResult{p, view}
			if err = ssaserverviews.ValidateSsaResult(vres); err != nil {
				return nil, goahttp.ErrValidationError("SSAServer", "Pick_up_data", err)
			}
			res := ssaserver.NewSsaResult(vres)
			return res, nil
		case http.StatusNotFound:
			var (
				body PickUpDataInvalidGroupIDResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("SSAServer", "Pick_up_data", err)
			}
			err = ValidatePickUpDataInvalidGroupIDResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("SSAServer", "Pick_up_data", err)
			}
			return nil, NewPickUpDataInvalidGroupID(&body)
		case http.StatusBadRequest:
			en := resp.Header.Get("goa-error")
			switch en {
			case "Invalid_Request":
				var (
					body PickUpDataInvalidRequestResponseBody
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("SSAServer", "Pick_up_data", err)
				}
				err = ValidatePickUpDataInvalidRequestResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("SSAServer", "Pick_up_data", err)
				}
				return nil, NewPickUpDataInvalidRequest(&body)
			case "Invalid_data_name":
				var (
					body PickUpDataInvalidDataNameResponseBody
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("SSAServer", "Pick_up_data", err)
				}
				err = ValidatePickUpDataInvalidDataNameResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("SSAServer", "Pick_up_data", err)
				}
				return nil, NewPickUpDataInvalidDataName(&body)
			default:
				body, _ := ioutil.ReadAll(resp.Body)
				return nil, goahttp.ErrInvalidResponse("SSAServer", "Pick_up_data", resp.StatusCode, string(body))
			}
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("SSAServer", "Pick_up_data", resp.StatusCode, string(body))
		}
	}
}
