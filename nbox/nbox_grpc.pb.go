// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: nbox/nbox.proto

package nbox

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NanoboxClient is the client API for Nanobox service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NanoboxClient interface {
	// DiscoverMaster returns the info of the current master
	DiscoverMaster(ctx context.Context, in *MasterInfoRequest, opts ...grpc.CallOption) (*MasterInfoResponse, error)
	Get(ctx context.Context, in *GetOrPeakRequest, opts ...grpc.CallOption) (*GetOrPeakResponse, error)
	Peek(ctx context.Context, in *GetOrPeakRequest, opts ...grpc.CallOption) (*GetOrPeakResponse, error)
	Set(ctx context.Context, in *SetOrUpdateRequest, opts ...grpc.CallOption) (*SetOrUpdateResponse, error)
	Update(ctx context.Context, in *SetOrUpdateRequest, opts ...grpc.CallOption) (*SetOrUpdateResponse, error)
	Delete(ctx context.Context, in *DeleteOrPurgeRequest, opts ...grpc.CallOption) (*DeleteOrPurgeResponse, error)
	Purge(ctx context.Context, in *DeleteOrPurgeRequest, opts ...grpc.CallOption) (*DeleteOrPurgeResponse, error)
	Keys(ctx context.Context, in *KeysRequest, opts ...grpc.CallOption) (*KeysResponse, error)
	Entries(ctx context.Context, in *EntriesRequest, opts ...grpc.CallOption) (*EntriesResponse, error)
	Size(ctx context.Context, in *SizeOrCapRequest, opts ...grpc.CallOption) (*SizeOrCapResponse, error)
	Cap(ctx context.Context, in *SizeOrCapRequest, opts ...grpc.CallOption) (*SizeOrCapResponse, error)
	DefaultTTL(ctx context.Context, in *DefaultTTLRequest, opts ...grpc.CallOption) (*DefaultTTLResponse, error)
	Resize(ctx context.Context, in *ResizeRequest, opts ...grpc.CallOption) (*ResizeResponse, error)
	UpdateDefaultTTL(ctx context.Context, in *UpdateDefaultTTLResponse, opts ...grpc.CallOption) (*UpdateDefaultTTLResponse, error)
}

type nanoboxClient struct {
	cc grpc.ClientConnInterface
}

func NewNanoboxClient(cc grpc.ClientConnInterface) NanoboxClient {
	return &nanoboxClient{cc}
}

func (c *nanoboxClient) DiscoverMaster(ctx context.Context, in *MasterInfoRequest, opts ...grpc.CallOption) (*MasterInfoResponse, error) {
	out := new(MasterInfoResponse)
	err := c.cc.Invoke(ctx, "/nbox.Nanobox/DiscoverMaster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nanoboxClient) Get(ctx context.Context, in *GetOrPeakRequest, opts ...grpc.CallOption) (*GetOrPeakResponse, error) {
	out := new(GetOrPeakResponse)
	err := c.cc.Invoke(ctx, "/nbox.Nanobox/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nanoboxClient) Peek(ctx context.Context, in *GetOrPeakRequest, opts ...grpc.CallOption) (*GetOrPeakResponse, error) {
	out := new(GetOrPeakResponse)
	err := c.cc.Invoke(ctx, "/nbox.Nanobox/Peek", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nanoboxClient) Set(ctx context.Context, in *SetOrUpdateRequest, opts ...grpc.CallOption) (*SetOrUpdateResponse, error) {
	out := new(SetOrUpdateResponse)
	err := c.cc.Invoke(ctx, "/nbox.Nanobox/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nanoboxClient) Update(ctx context.Context, in *SetOrUpdateRequest, opts ...grpc.CallOption) (*SetOrUpdateResponse, error) {
	out := new(SetOrUpdateResponse)
	err := c.cc.Invoke(ctx, "/nbox.Nanobox/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nanoboxClient) Delete(ctx context.Context, in *DeleteOrPurgeRequest, opts ...grpc.CallOption) (*DeleteOrPurgeResponse, error) {
	out := new(DeleteOrPurgeResponse)
	err := c.cc.Invoke(ctx, "/nbox.Nanobox/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nanoboxClient) Purge(ctx context.Context, in *DeleteOrPurgeRequest, opts ...grpc.CallOption) (*DeleteOrPurgeResponse, error) {
	out := new(DeleteOrPurgeResponse)
	err := c.cc.Invoke(ctx, "/nbox.Nanobox/Purge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nanoboxClient) Keys(ctx context.Context, in *KeysRequest, opts ...grpc.CallOption) (*KeysResponse, error) {
	out := new(KeysResponse)
	err := c.cc.Invoke(ctx, "/nbox.Nanobox/Keys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nanoboxClient) Entries(ctx context.Context, in *EntriesRequest, opts ...grpc.CallOption) (*EntriesResponse, error) {
	out := new(EntriesResponse)
	err := c.cc.Invoke(ctx, "/nbox.Nanobox/Entries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nanoboxClient) Size(ctx context.Context, in *SizeOrCapRequest, opts ...grpc.CallOption) (*SizeOrCapResponse, error) {
	out := new(SizeOrCapResponse)
	err := c.cc.Invoke(ctx, "/nbox.Nanobox/Size", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nanoboxClient) Cap(ctx context.Context, in *SizeOrCapRequest, opts ...grpc.CallOption) (*SizeOrCapResponse, error) {
	out := new(SizeOrCapResponse)
	err := c.cc.Invoke(ctx, "/nbox.Nanobox/Cap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nanoboxClient) DefaultTTL(ctx context.Context, in *DefaultTTLRequest, opts ...grpc.CallOption) (*DefaultTTLResponse, error) {
	out := new(DefaultTTLResponse)
	err := c.cc.Invoke(ctx, "/nbox.Nanobox/DefaultTTL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nanoboxClient) Resize(ctx context.Context, in *ResizeRequest, opts ...grpc.CallOption) (*ResizeResponse, error) {
	out := new(ResizeResponse)
	err := c.cc.Invoke(ctx, "/nbox.Nanobox/Resize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nanoboxClient) UpdateDefaultTTL(ctx context.Context, in *UpdateDefaultTTLResponse, opts ...grpc.CallOption) (*UpdateDefaultTTLResponse, error) {
	out := new(UpdateDefaultTTLResponse)
	err := c.cc.Invoke(ctx, "/nbox.Nanobox/UpdateDefaultTTL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NanoboxServer is the server API for Nanobox service.
// All implementations must embed UnimplementedNanoboxServer
// for forward compatibility
type NanoboxServer interface {
	// DiscoverMaster returns the info of the current master
	DiscoverMaster(context.Context, *MasterInfoRequest) (*MasterInfoResponse, error)
	Get(context.Context, *GetOrPeakRequest) (*GetOrPeakResponse, error)
	Peek(context.Context, *GetOrPeakRequest) (*GetOrPeakResponse, error)
	Set(context.Context, *SetOrUpdateRequest) (*SetOrUpdateResponse, error)
	Update(context.Context, *SetOrUpdateRequest) (*SetOrUpdateResponse, error)
	Delete(context.Context, *DeleteOrPurgeRequest) (*DeleteOrPurgeResponse, error)
	Purge(context.Context, *DeleteOrPurgeRequest) (*DeleteOrPurgeResponse, error)
	Keys(context.Context, *KeysRequest) (*KeysResponse, error)
	Entries(context.Context, *EntriesRequest) (*EntriesResponse, error)
	Size(context.Context, *SizeOrCapRequest) (*SizeOrCapResponse, error)
	Cap(context.Context, *SizeOrCapRequest) (*SizeOrCapResponse, error)
	DefaultTTL(context.Context, *DefaultTTLRequest) (*DefaultTTLResponse, error)
	Resize(context.Context, *ResizeRequest) (*ResizeResponse, error)
	UpdateDefaultTTL(context.Context, *UpdateDefaultTTLResponse) (*UpdateDefaultTTLResponse, error)
	mustEmbedUnimplementedNanoboxServer()
}

// UnimplementedNanoboxServer must be embedded to have forward compatible implementations.
type UnimplementedNanoboxServer struct {
}

func (UnimplementedNanoboxServer) DiscoverMaster(context.Context, *MasterInfoRequest) (*MasterInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiscoverMaster not implemented")
}
func (UnimplementedNanoboxServer) Get(context.Context, *GetOrPeakRequest) (*GetOrPeakResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedNanoboxServer) Peek(context.Context, *GetOrPeakRequest) (*GetOrPeakResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Peek not implemented")
}
func (UnimplementedNanoboxServer) Set(context.Context, *SetOrUpdateRequest) (*SetOrUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedNanoboxServer) Update(context.Context, *SetOrUpdateRequest) (*SetOrUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedNanoboxServer) Delete(context.Context, *DeleteOrPurgeRequest) (*DeleteOrPurgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedNanoboxServer) Purge(context.Context, *DeleteOrPurgeRequest) (*DeleteOrPurgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Purge not implemented")
}
func (UnimplementedNanoboxServer) Keys(context.Context, *KeysRequest) (*KeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Keys not implemented")
}
func (UnimplementedNanoboxServer) Entries(context.Context, *EntriesRequest) (*EntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Entries not implemented")
}
func (UnimplementedNanoboxServer) Size(context.Context, *SizeOrCapRequest) (*SizeOrCapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Size not implemented")
}
func (UnimplementedNanoboxServer) Cap(context.Context, *SizeOrCapRequest) (*SizeOrCapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cap not implemented")
}
func (UnimplementedNanoboxServer) DefaultTTL(context.Context, *DefaultTTLRequest) (*DefaultTTLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DefaultTTL not implemented")
}
func (UnimplementedNanoboxServer) Resize(context.Context, *ResizeRequest) (*ResizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Resize not implemented")
}
func (UnimplementedNanoboxServer) UpdateDefaultTTL(context.Context, *UpdateDefaultTTLResponse) (*UpdateDefaultTTLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDefaultTTL not implemented")
}
func (UnimplementedNanoboxServer) mustEmbedUnimplementedNanoboxServer() {}

// UnsafeNanoboxServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NanoboxServer will
// result in compilation errors.
type UnsafeNanoboxServer interface {
	mustEmbedUnimplementedNanoboxServer()
}

func RegisterNanoboxServer(s grpc.ServiceRegistrar, srv NanoboxServer) {
	s.RegisterService(&Nanobox_ServiceDesc, srv)
}

func _Nanobox_DiscoverMaster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MasterInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NanoboxServer).DiscoverMaster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nbox.Nanobox/DiscoverMaster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NanoboxServer).DiscoverMaster(ctx, req.(*MasterInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nanobox_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrPeakRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NanoboxServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nbox.Nanobox/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NanoboxServer).Get(ctx, req.(*GetOrPeakRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nanobox_Peek_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrPeakRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NanoboxServer).Peek(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nbox.Nanobox/Peek",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NanoboxServer).Peek(ctx, req.(*GetOrPeakRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nanobox_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetOrUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NanoboxServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nbox.Nanobox/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NanoboxServer).Set(ctx, req.(*SetOrUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nanobox_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetOrUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NanoboxServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nbox.Nanobox/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NanoboxServer).Update(ctx, req.(*SetOrUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nanobox_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrPurgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NanoboxServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nbox.Nanobox/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NanoboxServer).Delete(ctx, req.(*DeleteOrPurgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nanobox_Purge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrPurgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NanoboxServer).Purge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nbox.Nanobox/Purge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NanoboxServer).Purge(ctx, req.(*DeleteOrPurgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nanobox_Keys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NanoboxServer).Keys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nbox.Nanobox/Keys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NanoboxServer).Keys(ctx, req.(*KeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nanobox_Entries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NanoboxServer).Entries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nbox.Nanobox/Entries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NanoboxServer).Entries(ctx, req.(*EntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nanobox_Size_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SizeOrCapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NanoboxServer).Size(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nbox.Nanobox/Size",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NanoboxServer).Size(ctx, req.(*SizeOrCapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nanobox_Cap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SizeOrCapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NanoboxServer).Cap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nbox.Nanobox/Cap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NanoboxServer).Cap(ctx, req.(*SizeOrCapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nanobox_DefaultTTL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DefaultTTLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NanoboxServer).DefaultTTL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nbox.Nanobox/DefaultTTL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NanoboxServer).DefaultTTL(ctx, req.(*DefaultTTLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nanobox_Resize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NanoboxServer).Resize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nbox.Nanobox/Resize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NanoboxServer).Resize(ctx, req.(*ResizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nanobox_UpdateDefaultTTL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDefaultTTLResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NanoboxServer).UpdateDefaultTTL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nbox.Nanobox/UpdateDefaultTTL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NanoboxServer).UpdateDefaultTTL(ctx, req.(*UpdateDefaultTTLResponse))
	}
	return interceptor(ctx, in, info, handler)
}

// Nanobox_ServiceDesc is the grpc.ServiceDesc for Nanobox service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Nanobox_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nbox.Nanobox",
	HandlerType: (*NanoboxServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DiscoverMaster",
			Handler:    _Nanobox_DiscoverMaster_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Nanobox_Get_Handler,
		},
		{
			MethodName: "Peek",
			Handler:    _Nanobox_Peek_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _Nanobox_Set_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Nanobox_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Nanobox_Delete_Handler,
		},
		{
			MethodName: "Purge",
			Handler:    _Nanobox_Purge_Handler,
		},
		{
			MethodName: "Keys",
			Handler:    _Nanobox_Keys_Handler,
		},
		{
			MethodName: "Entries",
			Handler:    _Nanobox_Entries_Handler,
		},
		{
			MethodName: "Size",
			Handler:    _Nanobox_Size_Handler,
		},
		{
			MethodName: "Cap",
			Handler:    _Nanobox_Cap_Handler,
		},
		{
			MethodName: "DefaultTTL",
			Handler:    _Nanobox_DefaultTTL_Handler,
		},
		{
			MethodName: "Resize",
			Handler:    _Nanobox_Resize_Handler,
		},
		{
			MethodName: "UpdateDefaultTTL",
			Handler:    _Nanobox_UpdateDefaultTTL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nbox/nbox.proto",
}
