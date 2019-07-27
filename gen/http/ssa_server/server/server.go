// Code generated by goa v3.0.3, DO NOT EDIT.
//
// SSAServer HTTP server
//
// Command:
// $ goa gen SSAServer/design

package server

import (
	ssaserver "SSAServer/gen/ssa_server"
	"context"
	"mime/multipart"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the SSAServer service endpoint HTTP handlers.
type Server struct {
	Mounts         []*MountPoint
	Register       http.Handler
	Login          http.Handler
	ChangeGroup    http.Handler
	DeleteUser     http.Handler
	SaveData       http.Handler
	ReturnDataList http.Handler
	PickUpData     http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// SSAServerSaveDataDecoderFunc is the type to decode multipart request for the
// "SSAServer" service "Save_data" endpoint.
type SSAServerSaveDataDecoderFunc func(*multipart.Reader, **ssaserver.SaveDataPayload) error

// SSAServerPickUpDataDecoderFunc is the type to decode multipart request for
// the "SSAServer" service "Pick_up_data" endpoint.
type SSAServerPickUpDataDecoderFunc func(*multipart.Reader, **ssaserver.PickUpDataPayload) error

// New instantiates HTTP handlers for all the SSAServer service endpoints.
func New(
	e *ssaserver.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
	sSAServerSaveDataDecoderFn SSAServerSaveDataDecoderFunc,
	sSAServerPickUpDataDecoderFn SSAServerPickUpDataDecoderFunc,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Register", "POST", "/Registration"},
			{"Login", "POST", "/Login"},
			{"ChangeGroup", "POST", "/users/{user_id}"},
			{"DeleteUser", "DELETE", "/users/{user_id}"},
			{"SaveData", "POST", "/group/{group_id}"},
			{"ReturnDataList", "GET", "/group/{group_id}"},
			{"PickUpData", "GET", "/group/{group_id}/{data_type}"},
			{"gen/http/openapi.json", "GET", "/openapi.json"},
		},
		Register:       NewRegisterHandler(e.Register, mux, dec, enc, eh),
		Login:          NewLoginHandler(e.Login, mux, dec, enc, eh),
		ChangeGroup:    NewChangeGroupHandler(e.ChangeGroup, mux, dec, enc, eh),
		DeleteUser:     NewDeleteUserHandler(e.DeleteUser, mux, dec, enc, eh),
		SaveData:       NewSaveDataHandler(e.SaveData, mux, NewSSAServerSaveDataDecoder(mux, sSAServerSaveDataDecoderFn), enc, eh),
		ReturnDataList: NewReturnDataListHandler(e.ReturnDataList, mux, dec, enc, eh),
		PickUpData:     NewPickUpDataHandler(e.PickUpData, mux, NewSSAServerPickUpDataDecoder(mux, sSAServerPickUpDataDecoderFn), enc, eh),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "SSAServer" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Register = m(s.Register)
	s.Login = m(s.Login)
	s.ChangeGroup = m(s.ChangeGroup)
	s.DeleteUser = m(s.DeleteUser)
	s.SaveData = m(s.SaveData)
	s.ReturnDataList = m(s.ReturnDataList)
	s.PickUpData = m(s.PickUpData)
}

// Mount configures the mux to serve the SSAServer endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountRegisterHandler(mux, h.Register)
	MountLoginHandler(mux, h.Login)
	MountChangeGroupHandler(mux, h.ChangeGroup)
	MountDeleteUserHandler(mux, h.DeleteUser)
	MountSaveDataHandler(mux, h.SaveData)
	MountReturnDataListHandler(mux, h.ReturnDataList)
	MountPickUpDataHandler(mux, h.PickUpData)
	MountGenHTTPOpenapiJSON(mux, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "gen/http/openapi.json")
	}))
}

// MountRegisterHandler configures the mux to serve the "SSAServer" service
// "Register" endpoint.
func MountRegisterHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/Registration", f)
}

// NewRegisterHandler creates a HTTP handler which loads the HTTP request and
// calls the "SSAServer" service "Register" endpoint.
func NewRegisterHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeRegisterRequest(mux, dec)
		encodeResponse = EncodeRegisterResponse(enc)
		encodeError    = EncodeRegisterError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Register")
		ctx = context.WithValue(ctx, goa.ServiceKey, "SSAServer")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountLoginHandler configures the mux to serve the "SSAServer" service
// "Login" endpoint.
func MountLoginHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/Login", f)
}

// NewLoginHandler creates a HTTP handler which loads the HTTP request and
// calls the "SSAServer" service "Login" endpoint.
func NewLoginHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeLoginRequest(mux, dec)
		encodeResponse = EncodeLoginResponse(enc)
		encodeError    = EncodeLoginError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Login")
		ctx = context.WithValue(ctx, goa.ServiceKey, "SSAServer")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountChangeGroupHandler configures the mux to serve the "SSAServer" service
// "Change_group" endpoint.
func MountChangeGroupHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/users/{user_id}", f)
}

// NewChangeGroupHandler creates a HTTP handler which loads the HTTP request
// and calls the "SSAServer" service "Change_group" endpoint.
func NewChangeGroupHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeChangeGroupRequest(mux, dec)
		encodeResponse = EncodeChangeGroupResponse(enc)
		encodeError    = EncodeChangeGroupError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Change_group")
		ctx = context.WithValue(ctx, goa.ServiceKey, "SSAServer")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountDeleteUserHandler configures the mux to serve the "SSAServer" service
// "Delete_user" endpoint.
func MountDeleteUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/users/{user_id}", f)
}

// NewDeleteUserHandler creates a HTTP handler which loads the HTTP request and
// calls the "SSAServer" service "Delete_user" endpoint.
func NewDeleteUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeDeleteUserRequest(mux, dec)
		encodeResponse = EncodeDeleteUserResponse(enc)
		encodeError    = EncodeDeleteUserError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Delete_user")
		ctx = context.WithValue(ctx, goa.ServiceKey, "SSAServer")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountSaveDataHandler configures the mux to serve the "SSAServer" service
// "Save_data" endpoint.
func MountSaveDataHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/group/{group_id}", f)
}

// NewSaveDataHandler creates a HTTP handler which loads the HTTP request and
// calls the "SSAServer" service "Save_data" endpoint.
func NewSaveDataHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeSaveDataRequest(mux, dec)
		encodeResponse = EncodeSaveDataResponse(enc)
		encodeError    = EncodeSaveDataError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Save_data")
		ctx = context.WithValue(ctx, goa.ServiceKey, "SSAServer")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountReturnDataListHandler configures the mux to serve the "SSAServer"
// service "Return_data_list" endpoint.
func MountReturnDataListHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/group/{group_id}", f)
}

// NewReturnDataListHandler creates a HTTP handler which loads the HTTP request
// and calls the "SSAServer" service "Return_data_list" endpoint.
func NewReturnDataListHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeReturnDataListRequest(mux, dec)
		encodeResponse = EncodeReturnDataListResponse(enc)
		encodeError    = EncodeReturnDataListError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Return_data_list")
		ctx = context.WithValue(ctx, goa.ServiceKey, "SSAServer")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountPickUpDataHandler configures the mux to serve the "SSAServer" service
// "Pick_up_data" endpoint.
func MountPickUpDataHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/group/{group_id}/{data_type}", f)
}

// NewPickUpDataHandler creates a HTTP handler which loads the HTTP request and
// calls the "SSAServer" service "Pick_up_data" endpoint.
func NewPickUpDataHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodePickUpDataRequest(mux, dec)
		encodeResponse = EncodePickUpDataResponse(enc)
		encodeError    = EncodePickUpDataError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Pick_up_data")
		ctx = context.WithValue(ctx, goa.ServiceKey, "SSAServer")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountGenHTTPOpenapiJSON configures the mux to serve GET request made to
// "/openapi.json".
func MountGenHTTPOpenapiJSON(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/openapi.json", h.ServeHTTP)
}
