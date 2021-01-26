// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: accountsearch.proto

package profilecard

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
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

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  *float64 `protobuf:"fixed64,1,opt,name=latitude" json:"latitude,omitempty"`
	Longitude *float64 `protobuf:"fixed64,2,opt,name=longitude" json:"longitude,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountsearch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_accountsearch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_accountsearch_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetLatitude() float64 {
	if x != nil && x.Latitude != nil {
		return *x.Latitude
	}
	return 0
}

func (x *Location) GetLongitude() float64 {
	if x != nil && x.Longitude != nil {
		return *x.Longitude
	}
	return 0
}

type ResultItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FeedId      []byte   `protobuf:"bytes,1,opt,name=feedId" json:"feedId,omitempty"`
	Name        []byte   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	PicUrl      []byte   `protobuf:"bytes,3,opt,name=picUrl" json:"picUrl,omitempty"`
	JmpUrl      []byte   `protobuf:"bytes,4,opt,name=jmpUrl" json:"jmpUrl,omitempty"`
	FeedType    []byte   `protobuf:"bytes,5,opt,name=feedType" json:"feedType,omitempty"`
	Summary     []byte   `protobuf:"bytes,6,opt,name=summary" json:"summary,omitempty"`
	HasVideo    []byte   `protobuf:"bytes,7,opt,name=hasVideo" json:"hasVideo,omitempty"`
	PhtotUpdate []byte   `protobuf:"bytes,8,opt,name=phtotUpdate" json:"phtotUpdate,omitempty"`
	Uin         *uint64  `protobuf:"varint,9,opt,name=uin" json:"uin,omitempty"`
	ResultId    []byte   `protobuf:"bytes,10,opt,name=resultId" json:"resultId,omitempty"`
	Ftime       *uint32  `protobuf:"varint,11,opt,name=ftime" json:"ftime,omitempty"`
	NickName    []byte   `protobuf:"bytes,12,opt,name=nickName" json:"nickName,omitempty"`
	PicUrlList  [][]byte `protobuf:"bytes,13,rep,name=picUrlList" json:"picUrlList,omitempty"`
	TotalPicNum *uint32  `protobuf:"varint,14,opt,name=totalPicNum" json:"totalPicNum,omitempty"`
}

func (x *ResultItem) Reset() {
	*x = ResultItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountsearch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultItem) ProtoMessage() {}

func (x *ResultItem) ProtoReflect() protoreflect.Message {
	mi := &file_accountsearch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultItem.ProtoReflect.Descriptor instead.
func (*ResultItem) Descriptor() ([]byte, []int) {
	return file_accountsearch_proto_rawDescGZIP(), []int{1}
}

func (x *ResultItem) GetFeedId() []byte {
	if x != nil {
		return x.FeedId
	}
	return nil
}

func (x *ResultItem) GetName() []byte {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *ResultItem) GetPicUrl() []byte {
	if x != nil {
		return x.PicUrl
	}
	return nil
}

func (x *ResultItem) GetJmpUrl() []byte {
	if x != nil {
		return x.JmpUrl
	}
	return nil
}

func (x *ResultItem) GetFeedType() []byte {
	if x != nil {
		return x.FeedType
	}
	return nil
}

func (x *ResultItem) GetSummary() []byte {
	if x != nil {
		return x.Summary
	}
	return nil
}

func (x *ResultItem) GetHasVideo() []byte {
	if x != nil {
		return x.HasVideo
	}
	return nil
}

func (x *ResultItem) GetPhtotUpdate() []byte {
	if x != nil {
		return x.PhtotUpdate
	}
	return nil
}

func (x *ResultItem) GetUin() uint64 {
	if x != nil && x.Uin != nil {
		return *x.Uin
	}
	return 0
}

func (x *ResultItem) GetResultId() []byte {
	if x != nil {
		return x.ResultId
	}
	return nil
}

func (x *ResultItem) GetFtime() uint32 {
	if x != nil && x.Ftime != nil {
		return *x.Ftime
	}
	return 0
}

func (x *ResultItem) GetNickName() []byte {
	if x != nil {
		return x.NickName
	}
	return nil
}

func (x *ResultItem) GetPicUrlList() [][]byte {
	if x != nil {
		return x.PicUrlList
	}
	return nil
}

func (x *ResultItem) GetTotalPicNum() uint32 {
	if x != nil && x.TotalPicNum != nil {
		return *x.TotalPicNum
	}
	return 0
}

type Hotwordrecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hotword            *string `protobuf:"bytes,1,opt,name=hotword" json:"hotword,omitempty"`
	HotwordType        *uint32 `protobuf:"varint,2,opt,name=hotwordType" json:"hotwordType,omitempty"`
	HotwordCoverUrl    *string `protobuf:"bytes,3,opt,name=hotwordCoverUrl" json:"hotwordCoverUrl,omitempty"`
	HotwordTitle       *string `protobuf:"bytes,4,opt,name=hotwordTitle" json:"hotwordTitle,omitempty"`
	HotwordDescription *string `protobuf:"bytes,5,opt,name=hotwordDescription" json:"hotwordDescription,omitempty"`
}

func (x *Hotwordrecord) Reset() {
	*x = Hotwordrecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountsearch_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hotwordrecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hotwordrecord) ProtoMessage() {}

func (x *Hotwordrecord) ProtoReflect() protoreflect.Message {
	mi := &file_accountsearch_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hotwordrecord.ProtoReflect.Descriptor instead.
func (*Hotwordrecord) Descriptor() ([]byte, []int) {
	return file_accountsearch_proto_rawDescGZIP(), []int{2}
}

func (x *Hotwordrecord) GetHotword() string {
	if x != nil && x.Hotword != nil {
		return *x.Hotword
	}
	return ""
}

func (x *Hotwordrecord) GetHotwordType() uint32 {
	if x != nil && x.HotwordType != nil {
		return *x.HotwordType
	}
	return 0
}

func (x *Hotwordrecord) GetHotwordCoverUrl() string {
	if x != nil && x.HotwordCoverUrl != nil {
		return *x.HotwordCoverUrl
	}
	return ""
}

func (x *Hotwordrecord) GetHotwordTitle() string {
	if x != nil && x.HotwordTitle != nil {
		return *x.HotwordTitle
	}
	return ""
}

func (x *Hotwordrecord) GetHotwordDescription() string {
	if x != nil && x.HotwordDescription != nil {
		return *x.HotwordDescription
	}
	return ""
}

type AccountSearchRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uin               *uint64 `protobuf:"varint,1,opt,name=uin" json:"uin,omitempty"`
	Code              *uint64 `protobuf:"varint,2,opt,name=code" json:"code,omitempty"`
	Source            *uint32 `protobuf:"varint,3,opt,name=source" json:"source,omitempty"`
	Name              *string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Sex               *uint32 `protobuf:"varint,5,opt,name=sex" json:"sex,omitempty"`
	Age               *uint32 `protobuf:"varint,6,opt,name=age" json:"age,omitempty"`
	Accout            *string `protobuf:"bytes,7,opt,name=accout" json:"accout,omitempty"`
	Brief             *string `protobuf:"bytes,8,opt,name=brief" json:"brief,omitempty"`
	Number            *uint32 `protobuf:"varint,9,opt,name=number" json:"number,omitempty"`
	Flag              *uint64 `protobuf:"varint,10,opt,name=flag" json:"flag,omitempty"`
	Relation          *uint64 `protobuf:"varint,11,opt,name=relation" json:"relation,omitempty"`
	Mobile            *string `protobuf:"bytes,12,opt,name=mobile" json:"mobile,omitempty"`
	Sign              []byte  `protobuf:"bytes,13,opt,name=sign" json:"sign,omitempty"`
	Country           *uint32 `protobuf:"varint,14,opt,name=country" json:"country,omitempty"`
	Province          *uint32 `protobuf:"varint,15,opt,name=province" json:"province,omitempty"`
	City              *uint32 `protobuf:"varint,16,opt,name=city" json:"city,omitempty"`
	ClassIndex        *uint32 `protobuf:"varint,17,opt,name=classIndex" json:"classIndex,omitempty"`
	ClassName         *string `protobuf:"bytes,18,opt,name=className" json:"className,omitempty"`
	CountryName       *string `protobuf:"bytes,19,opt,name=countryName" json:"countryName,omitempty"`
	ProvinceName      *string `protobuf:"bytes,20,opt,name=provinceName" json:"provinceName,omitempty"`
	CityName          *string `protobuf:"bytes,21,opt,name=cityName" json:"cityName,omitempty"`
	AccountFlag       *uint32 `protobuf:"varint,22,opt,name=accountFlag" json:"accountFlag,omitempty"`
	TitleImage        *string `protobuf:"bytes,23,opt,name=titleImage" json:"titleImage,omitempty"`
	ArticleShortUrl   *string `protobuf:"bytes,24,opt,name=articleShortUrl" json:"articleShortUrl,omitempty"`
	ArticleCreateTime *string `protobuf:"bytes,25,opt,name=articleCreateTime" json:"articleCreateTime,omitempty"`
	ArticleAuthor     *string `protobuf:"bytes,26,opt,name=articleAuthor" json:"articleAuthor,omitempty"`
	AccountId         *uint64 `protobuf:"varint,27,opt,name=accountId" json:"accountId,omitempty"`
	//repeated Label groupLabels = 30;
	VideoAccount  *uint32 `protobuf:"varint,31,opt,name=videoAccount" json:"videoAccount,omitempty"`
	VideoArticle  *uint32 `protobuf:"varint,32,opt,name=videoArticle" json:"videoArticle,omitempty"`
	UinPrivilege  *int32  `protobuf:"varint,33,opt,name=uinPrivilege" json:"uinPrivilege,omitempty"`
	JoinGroupAuth []byte  `protobuf:"bytes,34,opt,name=joinGroupAuth" json:"joinGroupAuth,omitempty"`
	Token         []byte  `protobuf:"bytes,500,opt,name=token" json:"token,omitempty"`
	Richflag1_59  *uint32 `protobuf:"varint,40603,opt,name=richflag1_59,json=richflag159" json:"richflag1_59,omitempty"`
	Richflag4_409 *uint32 `protobuf:"varint,42409,opt,name=richflag4_409,json=richflag4409" json:"richflag4_409,omitempty"`
}

func (x *AccountSearchRecord) Reset() {
	*x = AccountSearchRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountsearch_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountSearchRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountSearchRecord) ProtoMessage() {}

func (x *AccountSearchRecord) ProtoReflect() protoreflect.Message {
	mi := &file_accountsearch_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountSearchRecord.ProtoReflect.Descriptor instead.
func (*AccountSearchRecord) Descriptor() ([]byte, []int) {
	return file_accountsearch_proto_rawDescGZIP(), []int{3}
}

func (x *AccountSearchRecord) GetUin() uint64 {
	if x != nil && x.Uin != nil {
		return *x.Uin
	}
	return 0
}

func (x *AccountSearchRecord) GetCode() uint64 {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return 0
}

func (x *AccountSearchRecord) GetSource() uint32 {
	if x != nil && x.Source != nil {
		return *x.Source
	}
	return 0
}

func (x *AccountSearchRecord) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *AccountSearchRecord) GetSex() uint32 {
	if x != nil && x.Sex != nil {
		return *x.Sex
	}
	return 0
}

func (x *AccountSearchRecord) GetAge() uint32 {
	if x != nil && x.Age != nil {
		return *x.Age
	}
	return 0
}

func (x *AccountSearchRecord) GetAccout() string {
	if x != nil && x.Accout != nil {
		return *x.Accout
	}
	return ""
}

func (x *AccountSearchRecord) GetBrief() string {
	if x != nil && x.Brief != nil {
		return *x.Brief
	}
	return ""
}

func (x *AccountSearchRecord) GetNumber() uint32 {
	if x != nil && x.Number != nil {
		return *x.Number
	}
	return 0
}

func (x *AccountSearchRecord) GetFlag() uint64 {
	if x != nil && x.Flag != nil {
		return *x.Flag
	}
	return 0
}

func (x *AccountSearchRecord) GetRelation() uint64 {
	if x != nil && x.Relation != nil {
		return *x.Relation
	}
	return 0
}

func (x *AccountSearchRecord) GetMobile() string {
	if x != nil && x.Mobile != nil {
		return *x.Mobile
	}
	return ""
}

func (x *AccountSearchRecord) GetSign() []byte {
	if x != nil {
		return x.Sign
	}
	return nil
}

func (x *AccountSearchRecord) GetCountry() uint32 {
	if x != nil && x.Country != nil {
		return *x.Country
	}
	return 0
}

func (x *AccountSearchRecord) GetProvince() uint32 {
	if x != nil && x.Province != nil {
		return *x.Province
	}
	return 0
}

func (x *AccountSearchRecord) GetCity() uint32 {
	if x != nil && x.City != nil {
		return *x.City
	}
	return 0
}

func (x *AccountSearchRecord) GetClassIndex() uint32 {
	if x != nil && x.ClassIndex != nil {
		return *x.ClassIndex
	}
	return 0
}

func (x *AccountSearchRecord) GetClassName() string {
	if x != nil && x.ClassName != nil {
		return *x.ClassName
	}
	return ""
}

func (x *AccountSearchRecord) GetCountryName() string {
	if x != nil && x.CountryName != nil {
		return *x.CountryName
	}
	return ""
}

func (x *AccountSearchRecord) GetProvinceName() string {
	if x != nil && x.ProvinceName != nil {
		return *x.ProvinceName
	}
	return ""
}

func (x *AccountSearchRecord) GetCityName() string {
	if x != nil && x.CityName != nil {
		return *x.CityName
	}
	return ""
}

func (x *AccountSearchRecord) GetAccountFlag() uint32 {
	if x != nil && x.AccountFlag != nil {
		return *x.AccountFlag
	}
	return 0
}

func (x *AccountSearchRecord) GetTitleImage() string {
	if x != nil && x.TitleImage != nil {
		return *x.TitleImage
	}
	return ""
}

func (x *AccountSearchRecord) GetArticleShortUrl() string {
	if x != nil && x.ArticleShortUrl != nil {
		return *x.ArticleShortUrl
	}
	return ""
}

func (x *AccountSearchRecord) GetArticleCreateTime() string {
	if x != nil && x.ArticleCreateTime != nil {
		return *x.ArticleCreateTime
	}
	return ""
}

func (x *AccountSearchRecord) GetArticleAuthor() string {
	if x != nil && x.ArticleAuthor != nil {
		return *x.ArticleAuthor
	}
	return ""
}

func (x *AccountSearchRecord) GetAccountId() uint64 {
	if x != nil && x.AccountId != nil {
		return *x.AccountId
	}
	return 0
}

func (x *AccountSearchRecord) GetVideoAccount() uint32 {
	if x != nil && x.VideoAccount != nil {
		return *x.VideoAccount
	}
	return 0
}

func (x *AccountSearchRecord) GetVideoArticle() uint32 {
	if x != nil && x.VideoArticle != nil {
		return *x.VideoArticle
	}
	return 0
}

func (x *AccountSearchRecord) GetUinPrivilege() int32 {
	if x != nil && x.UinPrivilege != nil {
		return *x.UinPrivilege
	}
	return 0
}

func (x *AccountSearchRecord) GetJoinGroupAuth() []byte {
	if x != nil {
		return x.JoinGroupAuth
	}
	return nil
}

func (x *AccountSearchRecord) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *AccountSearchRecord) GetRichflag1_59() uint32 {
	if x != nil && x.Richflag1_59 != nil {
		return *x.Richflag1_59
	}
	return 0
}

func (x *AccountSearchRecord) GetRichflag4_409() uint32 {
	if x != nil && x.Richflag4_409 != nil {
		return *x.Richflag4_409
	}
	return 0
}

type AccountSearch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start         *int32                 `protobuf:"varint,1,opt,name=start" json:"start,omitempty"`
	Count         *uint32                `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	End           *uint32                `protobuf:"varint,3,opt,name=end" json:"end,omitempty"`
	Keyword       *string                `protobuf:"bytes,4,opt,name=keyword" json:"keyword,omitempty"`
	List          []*AccountSearchRecord `protobuf:"bytes,5,rep,name=list" json:"list,omitempty"`
	Highlight     []string               `protobuf:"bytes,6,rep,name=highlight" json:"highlight,omitempty"`
	UserLocation  *Location              `protobuf:"bytes,10,opt,name=userLocation" json:"userLocation,omitempty"`
	LocationGroup *bool                  `protobuf:"varint,11,opt,name=locationGroup" json:"locationGroup,omitempty"`
	Filtertype    *int32                 `protobuf:"varint,12,opt,name=filtertype" json:"filtertype,omitempty"`
	//repeated C33304record recommendList = 13;
	HotwordRecord  *Hotwordrecord `protobuf:"bytes,14,opt,name=hotwordRecord" json:"hotwordRecord,omitempty"`
	ArticleMoreUrl *string        `protobuf:"bytes,15,opt,name=articleMoreUrl" json:"articleMoreUrl,omitempty"`
	ResultItems    []*ResultItem  `protobuf:"bytes,16,rep,name=resultItems" json:"resultItems,omitempty"`
	KeywordSuicide *bool          `protobuf:"varint,17,opt,name=keywordSuicide" json:"keywordSuicide,omitempty"`
	ExactSearch    *bool          `protobuf:"varint,18,opt,name=exactSearch" json:"exactSearch,omitempty"`
}

func (x *AccountSearch) Reset() {
	*x = AccountSearch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountsearch_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountSearch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountSearch) ProtoMessage() {}

func (x *AccountSearch) ProtoReflect() protoreflect.Message {
	mi := &file_accountsearch_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountSearch.ProtoReflect.Descriptor instead.
func (*AccountSearch) Descriptor() ([]byte, []int) {
	return file_accountsearch_proto_rawDescGZIP(), []int{4}
}

func (x *AccountSearch) GetStart() int32 {
	if x != nil && x.Start != nil {
		return *x.Start
	}
	return 0
}

func (x *AccountSearch) GetCount() uint32 {
	if x != nil && x.Count != nil {
		return *x.Count
	}
	return 0
}

func (x *AccountSearch) GetEnd() uint32 {
	if x != nil && x.End != nil {
		return *x.End
	}
	return 0
}

func (x *AccountSearch) GetKeyword() string {
	if x != nil && x.Keyword != nil {
		return *x.Keyword
	}
	return ""
}

func (x *AccountSearch) GetList() []*AccountSearchRecord {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *AccountSearch) GetHighlight() []string {
	if x != nil {
		return x.Highlight
	}
	return nil
}

func (x *AccountSearch) GetUserLocation() *Location {
	if x != nil {
		return x.UserLocation
	}
	return nil
}

func (x *AccountSearch) GetLocationGroup() bool {
	if x != nil && x.LocationGroup != nil {
		return *x.LocationGroup
	}
	return false
}

func (x *AccountSearch) GetFiltertype() int32 {
	if x != nil && x.Filtertype != nil {
		return *x.Filtertype
	}
	return 0
}

func (x *AccountSearch) GetHotwordRecord() *Hotwordrecord {
	if x != nil {
		return x.HotwordRecord
	}
	return nil
}

func (x *AccountSearch) GetArticleMoreUrl() string {
	if x != nil && x.ArticleMoreUrl != nil {
		return *x.ArticleMoreUrl
	}
	return ""
}

func (x *AccountSearch) GetResultItems() []*ResultItem {
	if x != nil {
		return x.ResultItems
	}
	return nil
}

func (x *AccountSearch) GetKeywordSuicide() bool {
	if x != nil && x.KeywordSuicide != nil {
		return *x.KeywordSuicide
	}
	return false
}

func (x *AccountSearch) GetExactSearch() bool {
	if x != nil && x.ExactSearch != nil {
		return *x.ExactSearch
	}
	return false
}

var File_accountsearch_proto protoreflect.FileDescriptor

var file_accountsearch_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0xfe, 0x02, 0x0a, 0x0a,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x65,
	0x65, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x66, 0x65, 0x65, 0x64,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x69, 0x63, 0x55, 0x72, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x69, 0x63, 0x55, 0x72, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x6a, 0x6d, 0x70, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06,
	0x6a, 0x6d, 0x70, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x65, 0x65, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x66, 0x65, 0x65, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08,
	0x68, 0x61, 0x73, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x68, 0x61, 0x73, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68, 0x74, 0x6f,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x70,
	0x68, 0x74, 0x6f, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69,
	0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x66, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x69,
	0x63, 0x55, 0x72, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0a,
	0x70, 0x69, 0x63, 0x55, 0x72, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x50, 0x69, 0x63, 0x4e, 0x75, 0x6d, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x69, 0x63, 0x4e, 0x75, 0x6d, 0x22, 0xc9, 0x01, 0x0a,
	0x0d, 0x68, 0x6f, 0x74, 0x77, 0x6f, 0x72, 0x64, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x68, 0x6f, 0x74, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x68, 0x6f, 0x74, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x68, 0x6f, 0x74, 0x77,
	0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x68,
	0x6f, 0x74, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x68, 0x6f,
	0x74, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x68, 0x6f, 0x74, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x6f, 0x76, 0x65,
	0x72, 0x55, 0x72, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x68, 0x6f, 0x74, 0x77, 0x6f, 0x72, 0x64, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x68, 0x6f, 0x74, 0x77,
	0x6f, 0x72, 0x64, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x68, 0x6f, 0x74, 0x77,
	0x6f, 0x72, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x68, 0x6f, 0x74, 0x77, 0x6f, 0x72, 0x64, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xea, 0x07, 0x0a, 0x13, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75,
	0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x73, 0x65, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x72, 0x69, 0x65, 0x66, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62,
	0x72, 0x69, 0x65, 0x66, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x6c, 0x61, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x69, 0x74, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x69, 0x74, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x6c, 0x61,
	0x67, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x46, 0x6c, 0x61, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x2c,
	0x0a, 0x11, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x1a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18,
	0x1b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x22, 0x0a, 0x0c, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x1f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x18, 0x20, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x50,
	0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x18, 0x21, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x75, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x6a, 0x6f, 0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x75, 0x74, 0x68, 0x18, 0x22, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0d, 0x6a, 0x6f, 0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x75,
	0x74, 0x68, 0x12, 0x15, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0xf4, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0c, 0x72, 0x69, 0x63,
	0x68, 0x66, 0x6c, 0x61, 0x67, 0x31, 0x5f, 0x35, 0x39, 0x18, 0x9b, 0xbd, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x72, 0x69, 0x63, 0x68, 0x66, 0x6c, 0x61, 0x67, 0x31, 0x35, 0x39, 0x12, 0x25,
	0x0a, 0x0d, 0x72, 0x69, 0x63, 0x68, 0x66, 0x6c, 0x61, 0x67, 0x34, 0x5f, 0x34, 0x30, 0x39, 0x18,
	0xa9, 0xcb, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x72, 0x69, 0x63, 0x68, 0x66, 0x6c, 0x61,
	0x67, 0x34, 0x34, 0x30, 0x39, 0x22, 0xfb, 0x03, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x28, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x69, 0x67,
	0x68, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x68, 0x69,
	0x67, 0x68, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x2d, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1e, 0x0a, 0x0a,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x74, 0x79, 0x70, 0x65, 0x12, 0x34, 0x0a, 0x0d,
	0x68, 0x6f, 0x74, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x68, 0x6f, 0x74, 0x77, 0x6f, 0x72, 0x64, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x0d, 0x68, 0x6f, 0x74, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4d, 0x6f, 0x72,
	0x65, 0x55, 0x72, 0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x4d, 0x6f, 0x72, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x2d, 0x0a, 0x0b, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0b, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x53, 0x75, 0x69, 0x63, 0x69, 0x64, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0e, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x53, 0x75, 0x69, 0x63, 0x69, 0x64,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78, 0x61, 0x63, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x65, 0x78, 0x61, 0x63, 0x74, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x63, 0x61, 0x72, 0x64,
}

var (
	file_accountsearch_proto_rawDescOnce sync.Once
	file_accountsearch_proto_rawDescData = file_accountsearch_proto_rawDesc
)

func file_accountsearch_proto_rawDescGZIP() []byte {
	file_accountsearch_proto_rawDescOnce.Do(func() {
		file_accountsearch_proto_rawDescData = protoimpl.X.CompressGZIP(file_accountsearch_proto_rawDescData)
	})
	return file_accountsearch_proto_rawDescData
}

var file_accountsearch_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_accountsearch_proto_goTypes = []interface{}{
	(*Location)(nil),            // 0: Location
	(*ResultItem)(nil),          // 1: ResultItem
	(*Hotwordrecord)(nil),       // 2: hotwordrecord
	(*AccountSearchRecord)(nil), // 3: AccountSearchRecord
	(*AccountSearch)(nil),       // 4: AccountSearch
}
var file_accountsearch_proto_depIdxs = []int32{
	3, // 0: AccountSearch.list:type_name -> AccountSearchRecord
	0, // 1: AccountSearch.userLocation:type_name -> Location
	2, // 2: AccountSearch.hotwordRecord:type_name -> hotwordrecord
	1, // 3: AccountSearch.resultItems:type_name -> ResultItem
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_accountsearch_proto_init() }
func file_accountsearch_proto_init() {
	if File_accountsearch_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_accountsearch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_accountsearch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultItem); i {
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
		file_accountsearch_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hotwordrecord); i {
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
		file_accountsearch_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountSearchRecord); i {
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
		file_accountsearch_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountSearch); i {
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
			RawDescriptor: file_accountsearch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_accountsearch_proto_goTypes,
		DependencyIndexes: file_accountsearch_proto_depIdxs,
		MessageInfos:      file_accountsearch_proto_msgTypes,
	}.Build()
	File_accountsearch_proto = out.File
	file_accountsearch_proto_rawDesc = nil
	file_accountsearch_proto_goTypes = nil
	file_accountsearch_proto_depIdxs = nil
}
