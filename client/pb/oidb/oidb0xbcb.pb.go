// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: oidb0xbcb.proto

package oidb

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CheckUrlReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url         []string `protobuf:"bytes,1,rep,name=url" json:"url,omitempty"`
	Refer       *string  `protobuf:"bytes,2,opt,name=refer" json:"refer,omitempty"`
	Plateform   *string  `protobuf:"bytes,3,opt,name=plateform" json:"plateform,omitempty"`
	QqPfTo      *string  `protobuf:"bytes,4,opt,name=qqPfTo" json:"qqPfTo,omitempty"`
	Type        *uint32  `protobuf:"varint,5,opt,name=type" json:"type,omitempty"`
	From        *uint32  `protobuf:"varint,6,opt,name=from" json:"from,omitempty"`
	Chatid      *uint64  `protobuf:"varint,7,opt,name=chatid" json:"chatid,omitempty"`
	ServiceType *uint64  `protobuf:"varint,8,opt,name=serviceType" json:"serviceType,omitempty"`
	SendUin     *uint64  `protobuf:"varint,9,opt,name=sendUin" json:"sendUin,omitempty"`
	ReqType     *string  `protobuf:"bytes,10,opt,name=reqType" json:"reqType,omitempty"`
	OriginalUrl *string  `protobuf:"bytes,11,opt,name=originalUrl" json:"originalUrl,omitempty"`
	IsArk       *bool    `protobuf:"varint,12,opt,name=isArk" json:"isArk,omitempty"`
	ArkName     *string  `protobuf:"bytes,13,opt,name=arkName" json:"arkName,omitempty"`
	IsFinish    *bool    `protobuf:"varint,14,opt,name=isFinish" json:"isFinish,omitempty"`
	SrcUrls     []string `protobuf:"bytes,15,rep,name=srcUrls" json:"srcUrls,omitempty"`
	SrcPlatform *uint32  `protobuf:"varint,16,opt,name=srcPlatform" json:"srcPlatform,omitempty"`
	Qua         *string  `protobuf:"bytes,17,opt,name=qua" json:"qua,omitempty"`
}

func (x *CheckUrlReq) Reset() {
	*x = CheckUrlReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0xbcb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUrlReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUrlReq) ProtoMessage() {}

func (x *CheckUrlReq) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0xbcb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUrlReq.ProtoReflect.Descriptor instead.
func (*CheckUrlReq) Descriptor() ([]byte, []int) {
	return file_oidb0xbcb_proto_rawDescGZIP(), []int{0}
}

func (x *CheckUrlReq) GetUrl() []string {
	if x != nil {
		return x.Url
	}
	return nil
}

func (x *CheckUrlReq) GetRefer() string {
	if x != nil && x.Refer != nil {
		return *x.Refer
	}
	return ""
}

func (x *CheckUrlReq) GetPlateform() string {
	if x != nil && x.Plateform != nil {
		return *x.Plateform
	}
	return ""
}

func (x *CheckUrlReq) GetQqPfTo() string {
	if x != nil && x.QqPfTo != nil {
		return *x.QqPfTo
	}
	return ""
}

func (x *CheckUrlReq) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *CheckUrlReq) GetFrom() uint32 {
	if x != nil && x.From != nil {
		return *x.From
	}
	return 0
}

func (x *CheckUrlReq) GetChatid() uint64 {
	if x != nil && x.Chatid != nil {
		return *x.Chatid
	}
	return 0
}

func (x *CheckUrlReq) GetServiceType() uint64 {
	if x != nil && x.ServiceType != nil {
		return *x.ServiceType
	}
	return 0
}

func (x *CheckUrlReq) GetSendUin() uint64 {
	if x != nil && x.SendUin != nil {
		return *x.SendUin
	}
	return 0
}

func (x *CheckUrlReq) GetReqType() string {
	if x != nil && x.ReqType != nil {
		return *x.ReqType
	}
	return ""
}

func (x *CheckUrlReq) GetOriginalUrl() string {
	if x != nil && x.OriginalUrl != nil {
		return *x.OriginalUrl
	}
	return ""
}

func (x *CheckUrlReq) GetIsArk() bool {
	if x != nil && x.IsArk != nil {
		return *x.IsArk
	}
	return false
}

func (x *CheckUrlReq) GetArkName() string {
	if x != nil && x.ArkName != nil {
		return *x.ArkName
	}
	return ""
}

func (x *CheckUrlReq) GetIsFinish() bool {
	if x != nil && x.IsFinish != nil {
		return *x.IsFinish
	}
	return false
}

func (x *CheckUrlReq) GetSrcUrls() []string {
	if x != nil {
		return x.SrcUrls
	}
	return nil
}

func (x *CheckUrlReq) GetSrcPlatform() uint32 {
	if x != nil && x.SrcPlatform != nil {
		return *x.SrcPlatform
	}
	return 0
}

func (x *CheckUrlReq) GetQua() string {
	if x != nil && x.Qua != nil {
		return *x.Qua
	}
	return ""
}

type CheckUrlReqItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url         *string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Refer       *string `protobuf:"bytes,2,opt,name=refer" json:"refer,omitempty"`
	Plateform   *string `protobuf:"bytes,3,opt,name=plateform" json:"plateform,omitempty"`
	QqPfTo      *string `protobuf:"bytes,4,opt,name=qqPfTo" json:"qqPfTo,omitempty"`
	Type        *uint32 `protobuf:"varint,5,opt,name=type" json:"type,omitempty"`
	From        *uint32 `protobuf:"varint,6,opt,name=from" json:"from,omitempty"`
	Chatid      *uint64 `protobuf:"varint,7,opt,name=chatid" json:"chatid,omitempty"`
	ServiceType *uint64 `protobuf:"varint,8,opt,name=serviceType" json:"serviceType,omitempty"`
	SendUin     *uint64 `protobuf:"varint,9,opt,name=sendUin" json:"sendUin,omitempty"`
	ReqType     *string `protobuf:"bytes,10,opt,name=reqType" json:"reqType,omitempty"`
}

func (x *CheckUrlReqItem) Reset() {
	*x = CheckUrlReqItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0xbcb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUrlReqItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUrlReqItem) ProtoMessage() {}

func (x *CheckUrlReqItem) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0xbcb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUrlReqItem.ProtoReflect.Descriptor instead.
func (*CheckUrlReqItem) Descriptor() ([]byte, []int) {
	return file_oidb0xbcb_proto_rawDescGZIP(), []int{1}
}

func (x *CheckUrlReqItem) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *CheckUrlReqItem) GetRefer() string {
	if x != nil && x.Refer != nil {
		return *x.Refer
	}
	return ""
}

func (x *CheckUrlReqItem) GetPlateform() string {
	if x != nil && x.Plateform != nil {
		return *x.Plateform
	}
	return ""
}

func (x *CheckUrlReqItem) GetQqPfTo() string {
	if x != nil && x.QqPfTo != nil {
		return *x.QqPfTo
	}
	return ""
}

func (x *CheckUrlReqItem) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *CheckUrlReqItem) GetFrom() uint32 {
	if x != nil && x.From != nil {
		return *x.From
	}
	return 0
}

func (x *CheckUrlReqItem) GetChatid() uint64 {
	if x != nil && x.Chatid != nil {
		return *x.Chatid
	}
	return 0
}

func (x *CheckUrlReqItem) GetServiceType() uint64 {
	if x != nil && x.ServiceType != nil {
		return *x.ServiceType
	}
	return 0
}

func (x *CheckUrlReqItem) GetSendUin() uint64 {
	if x != nil && x.SendUin != nil {
		return *x.SendUin
	}
	return 0
}

func (x *CheckUrlReqItem) GetReqType() string {
	if x != nil && x.ReqType != nil {
		return *x.ReqType
	}
	return ""
}

type CheckUrlRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results         []*UrlCheckResult `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
	NextReqDuration *uint32           `protobuf:"varint,2,opt,name=nextReqDuration" json:"nextReqDuration,omitempty"`
}

func (x *CheckUrlRsp) Reset() {
	*x = CheckUrlRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0xbcb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUrlRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUrlRsp) ProtoMessage() {}

func (x *CheckUrlRsp) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0xbcb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUrlRsp.ProtoReflect.Descriptor instead.
func (*CheckUrlRsp) Descriptor() ([]byte, []int) {
	return file_oidb0xbcb_proto_rawDescGZIP(), []int{2}
}

func (x *CheckUrlRsp) GetResults() []*UrlCheckResult {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *CheckUrlRsp) GetNextReqDuration() uint32 {
	if x != nil && x.NextReqDuration != nil {
		return *x.NextReqDuration
	}
	return 0
}

type DBCBReqBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NotUseCache *int32       `protobuf:"varint,9,opt,name=notUseCache" json:"notUseCache,omitempty"`
	CheckUrlReq *CheckUrlReq `protobuf:"bytes,10,opt,name=checkUrlReq" json:"checkUrlReq,omitempty"`
}

func (x *DBCBReqBody) Reset() {
	*x = DBCBReqBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0xbcb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBCBReqBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBCBReqBody) ProtoMessage() {}

func (x *DBCBReqBody) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0xbcb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBCBReqBody.ProtoReflect.Descriptor instead.
func (*DBCBReqBody) Descriptor() ([]byte, []int) {
	return file_oidb0xbcb_proto_rawDescGZIP(), []int{3}
}

func (x *DBCBReqBody) GetNotUseCache() int32 {
	if x != nil && x.NotUseCache != nil {
		return *x.NotUseCache
	}
	return 0
}

func (x *DBCBReqBody) GetCheckUrlReq() *CheckUrlReq {
	if x != nil {
		return x.CheckUrlReq
	}
	return nil
}

type DBCBRspBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Wording     *string      `protobuf:"bytes,1,opt,name=wording" json:"wording,omitempty"`
	CheckUrlRsp *CheckUrlRsp `protobuf:"bytes,10,opt,name=checkUrlRsp" json:"checkUrlRsp,omitempty"`
}

func (x *DBCBRspBody) Reset() {
	*x = DBCBRspBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0xbcb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBCBRspBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBCBRspBody) ProtoMessage() {}

func (x *DBCBRspBody) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0xbcb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBCBRspBody.ProtoReflect.Descriptor instead.
func (*DBCBRspBody) Descriptor() ([]byte, []int) {
	return file_oidb0xbcb_proto_rawDescGZIP(), []int{4}
}

func (x *DBCBRspBody) GetWording() string {
	if x != nil && x.Wording != nil {
		return *x.Wording
	}
	return ""
}

func (x *DBCBRspBody) GetCheckUrlRsp() *CheckUrlRsp {
	if x != nil {
		return x.CheckUrlRsp
	}
	return nil
}

type UrlCheckResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url          *string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Result       *uint32 `protobuf:"varint,2,opt,name=result" json:"result,omitempty"`
	JumpResult   *uint32 `protobuf:"varint,3,opt,name=jumpResult" json:"jumpResult,omitempty"`
	JumpUrl      *string `protobuf:"bytes,4,opt,name=jumpUrl" json:"jumpUrl,omitempty"`
	Level        *uint32 `protobuf:"varint,5,opt,name=level" json:"level,omitempty"`
	SubLevel     *uint32 `protobuf:"varint,6,opt,name=subLevel" json:"subLevel,omitempty"`
	Umrtype      *uint32 `protobuf:"varint,7,opt,name=umrtype" json:"umrtype,omitempty"`
	RetFrom      *uint32 `protobuf:"varint,8,opt,name=retFrom" json:"retFrom,omitempty"`
	OperationBit *uint64 `protobuf:"varint,9,opt,name=operationBit" json:"operationBit,omitempty"`
}

func (x *UrlCheckResult) Reset() {
	*x = UrlCheckResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0xbcb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UrlCheckResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlCheckResult) ProtoMessage() {}

func (x *UrlCheckResult) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0xbcb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlCheckResult.ProtoReflect.Descriptor instead.
func (*UrlCheckResult) Descriptor() ([]byte, []int) {
	return file_oidb0xbcb_proto_rawDescGZIP(), []int{5}
}

func (x *UrlCheckResult) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *UrlCheckResult) GetResult() uint32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *UrlCheckResult) GetJumpResult() uint32 {
	if x != nil && x.JumpResult != nil {
		return *x.JumpResult
	}
	return 0
}

func (x *UrlCheckResult) GetJumpUrl() string {
	if x != nil && x.JumpUrl != nil {
		return *x.JumpUrl
	}
	return ""
}

func (x *UrlCheckResult) GetLevel() uint32 {
	if x != nil && x.Level != nil {
		return *x.Level
	}
	return 0
}

func (x *UrlCheckResult) GetSubLevel() uint32 {
	if x != nil && x.SubLevel != nil {
		return *x.SubLevel
	}
	return 0
}

func (x *UrlCheckResult) GetUmrtype() uint32 {
	if x != nil && x.Umrtype != nil {
		return *x.Umrtype
	}
	return 0
}

func (x *UrlCheckResult) GetRetFrom() uint32 {
	if x != nil && x.RetFrom != nil {
		return *x.RetFrom
	}
	return 0
}

func (x *UrlCheckResult) GetOperationBit() uint64 {
	if x != nil && x.OperationBit != nil {
		return *x.OperationBit
	}
	return 0
}

var File_oidb0xbcb_proto protoreflect.FileDescriptor

var file_oidb0xbcb_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6f, 0x69, 0x64, 0x62, 0x30, 0x78, 0x62, 0x63, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xbd, 0x03, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x72, 0x6c, 0x52, 0x65,
	0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x66, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x66, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x71, 0x71, 0x50, 0x66, 0x54,
	0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x71, 0x71, 0x50, 0x66, 0x54, 0x6f, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x74, 0x69,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x69, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x6e, 0x64, 0x55, 0x69, 0x6e, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x73, 0x65, 0x6e, 0x64, 0x55, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x71, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65,
	0x71, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x55, 0x72, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x41, 0x72, 0x6b,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x41, 0x72, 0x6b, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x72, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x72, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x46, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x72, 0x63, 0x55, 0x72, 0x6c, 0x73, 0x18, 0x0f,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x72, 0x63, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x73, 0x72, 0x63, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x73, 0x72, 0x63, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12,
	0x10, 0x0a, 0x03, 0x71, 0x75, 0x61, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x71, 0x75,
	0x61, 0x22, 0x85, 0x02, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x72, 0x6c, 0x52, 0x65,
	0x71, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x66, 0x65, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x71,
	0x71, 0x50, 0x66, 0x54, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x71, 0x71, 0x50,
	0x66, 0x54, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x68, 0x61, 0x74, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x63, 0x68, 0x61,
	0x74, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x6e, 0x64, 0x55, 0x69, 0x6e,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x65, 0x6e, 0x64, 0x55, 0x69, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x72, 0x65, 0x71, 0x54, 0x79, 0x70, 0x65, 0x22, 0x62, 0x0a, 0x0b, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x55, 0x72, 0x6c, 0x52, 0x73, 0x70, 0x12, 0x29, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x55, 0x72, 0x6c, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6e, 0x65,
	0x78, 0x74, 0x52, 0x65, 0x71, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5f, 0x0a,
	0x0b, 0x44, 0x42, 0x43, 0x42, 0x52, 0x65, 0x71, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x20, 0x0a, 0x0b,
	0x6e, 0x6f, 0x74, 0x55, 0x73, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x6e, 0x6f, 0x74, 0x55, 0x73, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x2e,
	0x0a, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x72, 0x6c, 0x52, 0x65,
	0x71, 0x52, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x22, 0x57,
	0x0a, 0x0b, 0x44, 0x42, 0x43, 0x42, 0x52, 0x73, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x77, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x77, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x2e, 0x0a, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x55, 0x72, 0x6c, 0x52, 0x73, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x55, 0x72, 0x6c, 0x52, 0x73, 0x70, 0x52, 0x0b, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x55, 0x72, 0x6c, 0x52, 0x73, 0x70, 0x22, 0xfe, 0x01, 0x0a, 0x0e, 0x55, 0x72, 0x6c, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6a, 0x75, 0x6d, 0x70, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6a, 0x75, 0x6d, 0x70, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6a, 0x75, 0x6d, 0x70, 0x55, 0x72, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x75, 0x6d, 0x70, 0x55, 0x72, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x75, 0x62, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x18, 0x0a, 0x07, 0x75, 0x6d, 0x72, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x75, 0x6d, 0x72, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x74, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74,
	0x46, 0x72, 0x6f, 0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x69, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x69, 0x74, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x6f, 0x69,
	0x64, 0x62,
}

var (
	file_oidb0xbcb_proto_rawDescOnce sync.Once
	file_oidb0xbcb_proto_rawDescData = file_oidb0xbcb_proto_rawDesc
)

func file_oidb0xbcb_proto_rawDescGZIP() []byte {
	file_oidb0xbcb_proto_rawDescOnce.Do(func() {
		file_oidb0xbcb_proto_rawDescData = protoimpl.X.CompressGZIP(file_oidb0xbcb_proto_rawDescData)
	})
	return file_oidb0xbcb_proto_rawDescData
}

var file_oidb0xbcb_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_oidb0xbcb_proto_goTypes = []interface{}{
	(*CheckUrlReq)(nil),     // 0: CheckUrlReq
	(*CheckUrlReqItem)(nil), // 1: CheckUrlReqItem
	(*CheckUrlRsp)(nil),     // 2: CheckUrlRsp
	(*DBCBReqBody)(nil),     // 3: DBCBReqBody
	(*DBCBRspBody)(nil),     // 4: DBCBRspBody
	(*UrlCheckResult)(nil),  // 5: UrlCheckResult
}
var file_oidb0xbcb_proto_depIdxs = []int32{
	5, // 0: CheckUrlRsp.results:type_name -> UrlCheckResult
	0, // 1: DBCBReqBody.checkUrlReq:type_name -> CheckUrlReq
	2, // 2: DBCBRspBody.checkUrlRsp:type_name -> CheckUrlRsp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_oidb0xbcb_proto_init() }
func file_oidb0xbcb_proto_init() {
	if File_oidb0xbcb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oidb0xbcb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUrlReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_oidb0xbcb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUrlReqItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_oidb0xbcb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUrlRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_oidb0xbcb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBCBReqBody); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_oidb0xbcb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBCBRspBody); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_oidb0xbcb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UrlCheckResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_oidb0xbcb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_oidb0xbcb_proto_goTypes,
		DependencyIndexes: file_oidb0xbcb_proto_depIdxs,
		MessageInfos:      file_oidb0xbcb_proto_msgTypes,
	}.Build()
	File_oidb0xbcb_proto = out.File
	file_oidb0xbcb_proto_rawDesc = nil
	file_oidb0xbcb_proto_goTypes = nil
	file_oidb0xbcb_proto_depIdxs = nil
}
