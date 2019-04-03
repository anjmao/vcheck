package client

import "context"

type GetVersionRequest struct {
}

func (GetVersionRequest) Reset()         {}
func (GetVersionRequest) String() string { return "" }
func (GetVersionRequest) ProtoMessage()  {}

type GetVersionReply struct {
	BuildVersion string `protobuf:"bytes,1,opt,name=buildVersion,proto3" json:"buildVersion,omitempty"`
}

func (GetVersionReply) Reset()         {}
func (GetVersionReply) String() string { return "" }
func (GetVersionReply) ProtoMessage()  {}

type Client interface {
	GetVersion(ctx context.Context) (*GetVersionReply, error)
}
