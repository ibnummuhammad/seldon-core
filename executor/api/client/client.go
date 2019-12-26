package client

import (
	"context"
	"fmt"
	"github.com/seldonio/seldon-core/executor/api/payload"
	"golang.org/x/xerrors"
	"io"
)

const (
	SeldonPredictPath         = "/predict"
	SeldonTransformInputPath  = "/transform-input"
	SeldonTransformOutputPath = "/transform-output"
	SeldonCombinePath         = "/aggregate"
	SeldonRoutePath           = "/route"
)

type SeldonApiClient interface {
	Chain(ctx context.Context, msg payload.SeldonPayload) (payload.SeldonPayload, error)
	Predict(ctx context.Context, modelName string, host string, port int32, msg payload.SeldonPayload) (payload.SeldonPayload, error)
	TransformInput(ctx context.Context, modelName string, host string, port int32, msg payload.SeldonPayload) (payload.SeldonPayload, error)
	Route(ctx context.Context, modelName string, host string, port int32, msg payload.SeldonPayload) (int, error)
	Combine(ctx context.Context, modelName string, host string, port int32, msgs []payload.SeldonPayload) (payload.SeldonPayload, error)
	TransformOutput(ctx context.Context, modelName string, host string, port int32, msg payload.SeldonPayload) (payload.SeldonPayload, error)
	Unmarshall(msg []byte) (payload.SeldonPayload, error)
	Marshall(out io.Writer, msg payload.SeldonPayload) error
	CreateErrorPayload(err error) payload.SeldonPayload
}

type SeldonApiError struct {
	Message string
	Code    int
	frame   xerrors.Frame
}

func (se SeldonApiError) FormatError(p xerrors.Printer) error {
	p.Printf("%d %s", se.Code, se.Message)
	se.frame.Format(p)
	return nil
}

func (se SeldonApiError) Format(f fmt.State, c rune) {
	xerrors.FormatError(se, f, c)
}

func (se SeldonApiError) Error() string {
	return fmt.Sprint(se)
}
