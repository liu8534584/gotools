package resputil

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liu8534584/gotools/resputil/bizcode"
	"github.com/valyala/fasthttp"
	"net/http"
)

const (
	JsonHeader = "application/json"
)

type Body struct {
	Code    bizcode.Code `json:"code"`
	Message string       `json:"msg"`
	Data    interface{}  `json:"result"`
}

type Response struct {
	ctx        context.Context
	httpStatus int
	body       *Body
}

func New(ctx context.Context, code bizcode.Code) *Response {
	body := &Body{
		Code:    code,
		Message: code.String(),
	}

	return &Response{
		ctx:        ctx,
		httpStatus: http.StatusOK,
		body:       body,
	}
}

func Success(ctx context.Context) *Response {
	return New(ctx, bizcode.OK)
}

func Fail(ctx context.Context) *Response {
	return New(ctx, bizcode.InternalServerError)
}

func NeedLogin(ctx context.Context) *Response {
	return New(ctx, bizcode.NeedLogin)
}

func (r *Response) Message(msg string) *Response {
	r.body.Message = msg
	return r
}

func (r *Response) HTTPStatus(httpStatus int) *Response {
	r.httpStatus = httpStatus
	return r
}

func (r *Response) Data(data interface{}) *Response {
	r.body.Data = data
	return r
}

func (r *Response) Return() {
	switch r.ctx.(type) {
	case *gin.Context:
		r.ctx.(*gin.Context).PureJSON(r.httpStatus, r.body)
	case *fasthttp.RequestCtx:
		body, _ := json.Marshal(r.body)
		c := r.ctx.(*fasthttp.RequestCtx)
		c.Response.Header.SetContentType(JsonHeader)
		c.Response.SetBodyRaw(body)
	default:
		fmt.Println(r.httpStatus, r.body)
	}
}
