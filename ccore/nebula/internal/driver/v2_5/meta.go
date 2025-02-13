package v2_5

import (
	"github.com/facebook/fbthrift/thrift/lib/go/thrift"

	nerrors "github.com/vesoft-inc/nebula-http-gateway/ccore/nebula/errors"
	"github.com/vesoft-inc/nebula-http-gateway/ccore/nebula/internal/thrift/v2_5/meta"
	"github.com/vesoft-inc/nebula-http-gateway/ccore/nebula/types"
)

var (
	_ types.MetaClientDriver = (*defaultMetaClient)(nil)
)

type (
	defaultMetaClient struct {
		meta *meta.MetaServiceClient
	}
)

func newMetaClient(transport thrift.Transport, pf thrift.ProtocolFactory) types.MetaClientDriver {
	return &defaultMetaClient{
		meta: meta.NewMetaServiceClientFactory(transport, pf),
	}
}

func (c *defaultMetaClient) Open() error {
	return c.meta.Open()
}

func (c *defaultMetaClient) VerifyClientVersion() error {
	// v2.5 is not support verify client version, and it's the lowest version, so return not error.
	return nil
}

func (c *defaultMetaClient) Close() error {
	if c.meta != nil {
		if err := c.meta.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (c *defaultMetaClient) AddHosts(endpoints []string) error {
	return nerrors.ErrUnsupported
}

func (c *defaultMetaClient) DropHosts(endpoints []string) error {
	return nerrors.ErrUnsupported
}

func (c *defaultMetaClient) ListSpaces() (types.Spaces, error) {
	req := meta.NewListSpacesReq()

	resp, err := c.meta.ListSpaces(req)
	if err != nil {
		return nil, err
	}

	if err := codeErrorIfHappened(resp.Code, nil); err != nil {
		return nil, err
	}

	return newSpacesWrapper(resp.Spaces), nil
}

func (c *defaultMetaClient) Balance(req types.BalanceReq) (types.Balancer, error) {
	// TODO: add 2.5 Balance logic
	return nil, nerrors.ErrUnsupported
}
