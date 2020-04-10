// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package binacs_api_user_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type UserTestReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserTestReq) Reset()         { *m = UserTestReq{} }
func (m *UserTestReq) String() string { return proto.CompactTextString(m) }
func (*UserTestReq) ProtoMessage()    {}
func (*UserTestReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *UserTestReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserTestReq.Unmarshal(m, b)
}
func (m *UserTestReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserTestReq.Marshal(b, m, deterministic)
}
func (m *UserTestReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserTestReq.Merge(m, src)
}
func (m *UserTestReq) XXX_Size() int {
	return xxx_messageInfo_UserTestReq.Size(m)
}
func (m *UserTestReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserTestReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserTestReq proto.InternalMessageInfo

type UserTestResp struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserTestResp) Reset()         { *m = UserTestResp{} }
func (m *UserTestResp) String() string { return proto.CompactTextString(m) }
func (*UserTestResp) ProtoMessage()    {}
func (*UserTestResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserTestResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserTestResp.Unmarshal(m, b)
}
func (m *UserTestResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserTestResp.Marshal(b, m, deterministic)
}
func (m *UserTestResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserTestResp.Merge(m, src)
}
func (m *UserTestResp) XXX_Size() int {
	return xxx_messageInfo_UserTestResp.Size(m)
}
func (m *UserTestResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserTestResp.DiscardUnknown(m)
}

var xxx_messageInfo_UserTestResp proto.InternalMessageInfo

func (m *UserTestResp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UserTestResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

// UserRegister
type UserRegisterDataObj struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Pwd                  string   `protobuf:"bytes,2,opt,name=pwd,proto3" json:"pwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRegisterDataObj) Reset()         { *m = UserRegisterDataObj{} }
func (m *UserRegisterDataObj) String() string { return proto.CompactTextString(m) }
func (*UserRegisterDataObj) ProtoMessage()    {}
func (*UserRegisterDataObj) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *UserRegisterDataObj) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRegisterDataObj.Unmarshal(m, b)
}
func (m *UserRegisterDataObj) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRegisterDataObj.Marshal(b, m, deterministic)
}
func (m *UserRegisterDataObj) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRegisterDataObj.Merge(m, src)
}
func (m *UserRegisterDataObj) XXX_Size() int {
	return xxx_messageInfo_UserRegisterDataObj.Size(m)
}
func (m *UserRegisterDataObj) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRegisterDataObj.DiscardUnknown(m)
}

var xxx_messageInfo_UserRegisterDataObj proto.InternalMessageInfo

func (m *UserRegisterDataObj) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserRegisterDataObj) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type UserRegisterReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Pwd                  string   `protobuf:"bytes,2,opt,name=pwd,proto3" json:"pwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRegisterReq) Reset()         { *m = UserRegisterReq{} }
func (m *UserRegisterReq) String() string { return proto.CompactTextString(m) }
func (*UserRegisterReq) ProtoMessage()    {}
func (*UserRegisterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *UserRegisterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRegisterReq.Unmarshal(m, b)
}
func (m *UserRegisterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRegisterReq.Marshal(b, m, deterministic)
}
func (m *UserRegisterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRegisterReq.Merge(m, src)
}
func (m *UserRegisterReq) XXX_Size() int {
	return xxx_messageInfo_UserRegisterReq.Size(m)
}
func (m *UserRegisterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRegisterReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserRegisterReq proto.InternalMessageInfo

func (m *UserRegisterReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserRegisterReq) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type UserRegisterResp struct {
	Code                 int64                `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string               `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data                 *UserRegisterDataObj `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UserRegisterResp) Reset()         { *m = UserRegisterResp{} }
func (m *UserRegisterResp) String() string { return proto.CompactTextString(m) }
func (*UserRegisterResp) ProtoMessage()    {}
func (*UserRegisterResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *UserRegisterResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRegisterResp.Unmarshal(m, b)
}
func (m *UserRegisterResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRegisterResp.Marshal(b, m, deterministic)
}
func (m *UserRegisterResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRegisterResp.Merge(m, src)
}
func (m *UserRegisterResp) XXX_Size() int {
	return xxx_messageInfo_UserRegisterResp.Size(m)
}
func (m *UserRegisterResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRegisterResp.DiscardUnknown(m)
}

var xxx_messageInfo_UserRegisterResp proto.InternalMessageInfo

func (m *UserRegisterResp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UserRegisterResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *UserRegisterResp) GetData() *UserRegisterDataObj {
	if m != nil {
		return m.Data
	}
	return nil
}

// UserAuth
type UserTokenObj struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	ExpireTime           int64    `protobuf:"varint,2,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	RefreshToken         string   `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserTokenObj) Reset()         { *m = UserTokenObj{} }
func (m *UserTokenObj) String() string { return proto.CompactTextString(m) }
func (*UserTokenObj) ProtoMessage()    {}
func (*UserTokenObj) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *UserTokenObj) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserTokenObj.Unmarshal(m, b)
}
func (m *UserTokenObj) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserTokenObj.Marshal(b, m, deterministic)
}
func (m *UserTokenObj) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserTokenObj.Merge(m, src)
}
func (m *UserTokenObj) XXX_Size() int {
	return xxx_messageInfo_UserTokenObj.Size(m)
}
func (m *UserTokenObj) XXX_DiscardUnknown() {
	xxx_messageInfo_UserTokenObj.DiscardUnknown(m)
}

var xxx_messageInfo_UserTokenObj proto.InternalMessageInfo

func (m *UserTokenObj) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *UserTokenObj) GetExpireTime() int64 {
	if m != nil {
		return m.ExpireTime
	}
	return 0
}

func (m *UserTokenObj) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type UserAuthReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Pwd                  string   `protobuf:"bytes,2,opt,name=pwd,proto3" json:"pwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserAuthReq) Reset()         { *m = UserAuthReq{} }
func (m *UserAuthReq) String() string { return proto.CompactTextString(m) }
func (*UserAuthReq) ProtoMessage()    {}
func (*UserAuthReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *UserAuthReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAuthReq.Unmarshal(m, b)
}
func (m *UserAuthReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAuthReq.Marshal(b, m, deterministic)
}
func (m *UserAuthReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAuthReq.Merge(m, src)
}
func (m *UserAuthReq) XXX_Size() int {
	return xxx_messageInfo_UserAuthReq.Size(m)
}
func (m *UserAuthReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAuthReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserAuthReq proto.InternalMessageInfo

func (m *UserAuthReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserAuthReq) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type UserAuthResp struct {
	Code                 int64         `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string        `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data                 *UserTokenObj `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UserAuthResp) Reset()         { *m = UserAuthResp{} }
func (m *UserAuthResp) String() string { return proto.CompactTextString(m) }
func (*UserAuthResp) ProtoMessage()    {}
func (*UserAuthResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}

func (m *UserAuthResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAuthResp.Unmarshal(m, b)
}
func (m *UserAuthResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAuthResp.Marshal(b, m, deterministic)
}
func (m *UserAuthResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAuthResp.Merge(m, src)
}
func (m *UserAuthResp) XXX_Size() int {
	return xxx_messageInfo_UserAuthResp.Size(m)
}
func (m *UserAuthResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAuthResp.DiscardUnknown(m)
}

var xxx_messageInfo_UserAuthResp proto.InternalMessageInfo

func (m *UserAuthResp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UserAuthResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *UserAuthResp) GetData() *UserTokenObj {
	if m != nil {
		return m.Data
	}
	return nil
}

// UserRefresh
type UserRefreshReq struct {
	RefreshToken         string   `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRefreshReq) Reset()         { *m = UserRefreshReq{} }
func (m *UserRefreshReq) String() string { return proto.CompactTextString(m) }
func (*UserRefreshReq) ProtoMessage()    {}
func (*UserRefreshReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{8}
}

func (m *UserRefreshReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRefreshReq.Unmarshal(m, b)
}
func (m *UserRefreshReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRefreshReq.Marshal(b, m, deterministic)
}
func (m *UserRefreshReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRefreshReq.Merge(m, src)
}
func (m *UserRefreshReq) XXX_Size() int {
	return xxx_messageInfo_UserRefreshReq.Size(m)
}
func (m *UserRefreshReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRefreshReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserRefreshReq proto.InternalMessageInfo

func (m *UserRefreshReq) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type UserRefreshResp struct {
	Code                 int64         `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string        `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data                 *UserTokenObj `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UserRefreshResp) Reset()         { *m = UserRefreshResp{} }
func (m *UserRefreshResp) String() string { return proto.CompactTextString(m) }
func (*UserRefreshResp) ProtoMessage()    {}
func (*UserRefreshResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{9}
}

func (m *UserRefreshResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRefreshResp.Unmarshal(m, b)
}
func (m *UserRefreshResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRefreshResp.Marshal(b, m, deterministic)
}
func (m *UserRefreshResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRefreshResp.Merge(m, src)
}
func (m *UserRefreshResp) XXX_Size() int {
	return xxx_messageInfo_UserRefreshResp.Size(m)
}
func (m *UserRefreshResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRefreshResp.DiscardUnknown(m)
}

var xxx_messageInfo_UserRefreshResp proto.InternalMessageInfo

func (m *UserRefreshResp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UserRefreshResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *UserRefreshResp) GetData() *UserTokenObj {
	if m != nil {
		return m.Data
	}
	return nil
}

// UserInfo
type UserInfoObj struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Role                 string   `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	Desc                 string   `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	Ctime                string   `protobuf:"bytes,4,opt,name=ctime,proto3" json:"ctime,omitempty"`
	More                 []string `protobuf:"bytes,5,rep,name=more,proto3" json:"more,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoObj) Reset()         { *m = UserInfoObj{} }
func (m *UserInfoObj) String() string { return proto.CompactTextString(m) }
func (*UserInfoObj) ProtoMessage()    {}
func (*UserInfoObj) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{10}
}

func (m *UserInfoObj) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoObj.Unmarshal(m, b)
}
func (m *UserInfoObj) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoObj.Marshal(b, m, deterministic)
}
func (m *UserInfoObj) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoObj.Merge(m, src)
}
func (m *UserInfoObj) XXX_Size() int {
	return xxx_messageInfo_UserInfoObj.Size(m)
}
func (m *UserInfoObj) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoObj.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoObj proto.InternalMessageInfo

func (m *UserInfoObj) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserInfoObj) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *UserInfoObj) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *UserInfoObj) GetCtime() string {
	if m != nil {
		return m.Ctime
	}
	return ""
}

func (m *UserInfoObj) GetMore() []string {
	if m != nil {
		return m.More
	}
	return nil
}

type UserInfoReq struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoReq) Reset()         { *m = UserInfoReq{} }
func (m *UserInfoReq) String() string { return proto.CompactTextString(m) }
func (*UserInfoReq) ProtoMessage()    {}
func (*UserInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{11}
}

func (m *UserInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoReq.Unmarshal(m, b)
}
func (m *UserInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoReq.Marshal(b, m, deterministic)
}
func (m *UserInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoReq.Merge(m, src)
}
func (m *UserInfoReq) XXX_Size() int {
	return xxx_messageInfo_UserInfoReq.Size(m)
}
func (m *UserInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoReq proto.InternalMessageInfo

func (m *UserInfoReq) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type UserInfoResp struct {
	Code                 int64        `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string       `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data                 *UserInfoObj `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UserInfoResp) Reset()         { *m = UserInfoResp{} }
func (m *UserInfoResp) String() string { return proto.CompactTextString(m) }
func (*UserInfoResp) ProtoMessage()    {}
func (*UserInfoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{12}
}

func (m *UserInfoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoResp.Unmarshal(m, b)
}
func (m *UserInfoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoResp.Marshal(b, m, deterministic)
}
func (m *UserInfoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoResp.Merge(m, src)
}
func (m *UserInfoResp) XXX_Size() int {
	return xxx_messageInfo_UserInfoResp.Size(m)
}
func (m *UserInfoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoResp.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoResp proto.InternalMessageInfo

func (m *UserInfoResp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UserInfoResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *UserInfoResp) GetData() *UserInfoObj {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*UserTestReq)(nil), "binacs_api_user_v1.UserTestReq")
	proto.RegisterType((*UserTestResp)(nil), "binacs_api_user_v1.UserTestResp")
	proto.RegisterType((*UserRegisterDataObj)(nil), "binacs_api_user_v1.UserRegisterDataObj")
	proto.RegisterType((*UserRegisterReq)(nil), "binacs_api_user_v1.UserRegisterReq")
	proto.RegisterType((*UserRegisterResp)(nil), "binacs_api_user_v1.UserRegisterResp")
	proto.RegisterType((*UserTokenObj)(nil), "binacs_api_user_v1.UserTokenObj")
	proto.RegisterType((*UserAuthReq)(nil), "binacs_api_user_v1.UserAuthReq")
	proto.RegisterType((*UserAuthResp)(nil), "binacs_api_user_v1.UserAuthResp")
	proto.RegisterType((*UserRefreshReq)(nil), "binacs_api_user_v1.UserRefreshReq")
	proto.RegisterType((*UserRefreshResp)(nil), "binacs_api_user_v1.UserRefreshResp")
	proto.RegisterType((*UserInfoObj)(nil), "binacs_api_user_v1.UserInfoObj")
	proto.RegisterType((*UserInfoReq)(nil), "binacs_api_user_v1.UserInfoReq")
	proto.RegisterType((*UserInfoResp)(nil), "binacs_api_user_v1.UserInfoResp")
}

func init() {
	proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf)
}

var fileDescriptor_116e343673f7ffaf = []byte{
	// 568 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0x63, 0x17, 0xd1, 0x49, 0x5a, 0xca, 0x52, 0x50, 0x88, 0x2a, 0xc5, 0x6c, 0x90, 0x88,
	0x7a, 0x48, 0x68, 0x53, 0x84, 0x54, 0x4e, 0x48, 0x5c, 0x38, 0x21, 0x59, 0xe5, 0x1c, 0x6d, 0xec,
	0x49, 0xb2, 0x6d, 0xe3, 0x75, 0x77, 0x37, 0x0d, 0x12, 0x37, 0x5e, 0x81, 0x47, 0xe2, 0x11, 0x78,
	0x05, 0x1e, 0x04, 0xed, 0x8f, 0xd5, 0x86, 0xc6, 0x49, 0x2e, 0xdc, 0x66, 0xc7, 0xf3, 0xed, 0xcc,
	0xf7, 0xcd, 0xcc, 0x1a, 0x60, 0xae, 0x50, 0xf6, 0x0a, 0x29, 0xb4, 0x20, 0x64, 0xc4, 0x73, 0x96,
	0xaa, 0x21, 0x2b, 0xf8, 0xd0, 0xb8, 0x87, 0xb7, 0x27, 0xad, 0xa3, 0x89, 0x10, 0x93, 0x6b, 0xec,
	0xb3, 0x82, 0xf7, 0x59, 0x9e, 0x0b, 0xcd, 0x34, 0x17, 0xb9, 0x72, 0x08, 0xba, 0x07, 0xf5, 0xaf,
	0x0a, 0xe5, 0x05, 0x2a, 0x9d, 0xe0, 0x0d, 0x3d, 0x83, 0xc6, 0xdd, 0x51, 0x15, 0x84, 0x40, 0x94,
	0x8a, 0x0c, 0x9b, 0x41, 0x1c, 0x74, 0xc3, 0xc4, 0xda, 0xe4, 0x00, 0xc2, 0x99, 0x9a, 0x34, 0x6b,
	0x71, 0xd0, 0xdd, 0x4d, 0x8c, 0x49, 0xdf, 0xc3, 0x33, 0x83, 0x4a, 0x70, 0xc2, 0x95, 0x46, 0xf9,
	0x89, 0x69, 0xf6, 0x65, 0x74, 0x49, 0xf6, 0xa1, 0xc6, 0x33, 0x0b, 0xdd, 0x4d, 0x6a, 0x3c, 0x33,
	0xc0, 0x62, 0x91, 0x95, 0xc0, 0x62, 0x91, 0xd1, 0x01, 0x3c, 0xb9, 0x0f, 0x4c, 0xf0, 0x66, 0x0b,
	0xd0, 0x1c, 0x0e, 0x96, 0x41, 0xdb, 0xd6, 0x49, 0x3e, 0x40, 0x94, 0x31, 0xcd, 0x9a, 0x61, 0x1c,
	0x74, 0xeb, 0xa7, 0x6f, 0x7a, 0x0f, 0xd5, 0xea, 0xad, 0xe0, 0x91, 0x58, 0x10, 0x5d, 0x78, 0x69,
	0xc4, 0x15, 0xe6, 0x86, 0xdd, 0x2b, 0x68, 0xb0, 0x34, 0x45, 0xa5, 0x86, 0xda, 0xb8, 0x7c, 0xc9,
	0x75, 0xe7, 0xb3, 0x51, 0xa4, 0x0d, 0x75, 0xfc, 0x56, 0x70, 0x89, 0x43, 0xcd, 0x67, 0x68, 0x2b,
	0x09, 0x13, 0x70, 0xae, 0x0b, 0x3e, 0x43, 0xd2, 0x81, 0x3d, 0x89, 0x63, 0x89, 0x6a, 0xea, 0x2f,
	0x09, 0xed, 0x25, 0x0d, 0xef, 0xb4, 0xb7, 0xd0, 0xbe, 0x6b, 0xd1, 0xc7, 0xb9, 0x9e, 0x6e, 0x27,
	0xd0, 0xa5, 0xab, 0xd4, 0x01, 0xb6, 0x16, 0xe7, 0x6c, 0x49, 0x9c, 0xb8, 0x4a, 0x9c, 0x92, 0xbf,
	0x57, 0xe5, 0x1d, 0xec, 0x3b, 0xc9, 0x6c, 0xc1, 0xa6, 0xbe, 0x07, 0x9c, 0x82, 0x15, 0x9c, 0x66,
	0x65, 0xe3, 0x3d, 0xec, 0x3f, 0x57, 0x29, 0x9c, 0x84, 0x9f, 0xf3, 0xb1, 0x58, 0x35, 0x98, 0x04,
	0x22, 0x29, 0xae, 0xd1, 0xe7, 0xb1, 0xb6, 0xf1, 0x65, 0xa8, 0x52, 0xdf, 0x11, 0x6b, 0x93, 0x43,
	0xd8, 0x49, 0x6d, 0x27, 0x23, 0xeb, 0x74, 0x07, 0x13, 0x39, 0x13, 0x12, 0x9b, 0x3b, 0x71, 0x68,
	0x22, 0x8d, 0x4d, 0xdf, 0xde, 0x25, 0x34, 0x9a, 0x6c, 0x9e, 0x15, 0xca, 0x5d, 0xd3, 0x1c, 0x62,
	0x6b, 0x39, 0x06, 0x4b, 0x72, 0xb4, 0xab, 0xe4, 0xf0, 0xc4, 0x9d, 0x1a, 0xa7, 0xbf, 0x22, 0x88,
	0x8c, 0x97, 0x5c, 0xc1, 0xe3, 0x72, 0xdb, 0x49, 0x25, 0xd6, 0x3f, 0x0d, 0xad, 0x78, 0x7d, 0x80,
	0x2a, 0xe8, 0xd1, 0x8f, 0xdf, 0x7f, 0x7e, 0xd6, 0x5e, 0xd0, 0xa7, 0xf6, 0xad, 0xb9, 0x3d, 0xe9,
	0x9b, 0xaf, 0x7d, 0x8d, 0x4a, 0x9f, 0x07, 0xc7, 0xe4, 0xbb, 0x23, 0x58, 0x2e, 0x17, 0xe9, 0x6c,
	0x5a, 0x3f, 0x93, 0xf4, 0xf5, 0xe6, 0x20, 0x55, 0xd0, 0xd8, 0x26, 0x6e, 0xd1, 0xe7, 0x4b, 0x89,
	0xa5, 0x0f, 0x31, 0xc9, 0x3d, 0x53, 0xb3, 0x12, 0xd5, 0x4c, 0xfd, 0x86, 0x55, 0x33, 0x2d, 0x37,
	0xaa, 0x82, 0x29, 0x9b, 0xeb, 0xa9, 0x49, 0xb6, 0x70, 0xcd, 0xf7, 0xc3, 0x4d, 0x68, 0x35, 0x87,
	0x72, 0x69, 0x5a, 0x9d, 0x8d, 0x31, 0xaa, 0xa0, 0x6d, 0x9b, 0xf5, 0x25, 0x3d, 0xfc, 0x87, 0xa6,
	0x8d, 0xb8, 0xc7, 0xd2, 0x74, 0x9b, 0xac, 0x9d, 0x85, 0xb5, 0x2c, 0xcb, 0x11, 0xac, 0x60, 0xc9,
	0xf3, 0xb1, 0x38, 0x0f, 0x8e, 0x47, 0x8f, 0xec, 0x0f, 0x64, 0xf0, 0x37, 0x00, 0x00, 0xff, 0xff,
	0xb5, 0xc4, 0x66, 0xaf, 0x80, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	UserTest(ctx context.Context, in *UserTestReq, opts ...grpc.CallOption) (*UserTestResp, error)
	UserRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterResp, error)
	UserAuth(ctx context.Context, in *UserAuthReq, opts ...grpc.CallOption) (*UserAuthResp, error)
	UserRefresh(ctx context.Context, in *UserRefreshReq, opts ...grpc.CallOption) (*UserRefreshResp, error)
	UserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) UserTest(ctx context.Context, in *UserTestReq, opts ...grpc.CallOption) (*UserTestResp, error) {
	out := new(UserTestResp)
	err := c.cc.Invoke(ctx, "/binacs_api_user_v1.User/UserTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterResp, error) {
	out := new(UserRegisterResp)
	err := c.cc.Invoke(ctx, "/binacs_api_user_v1.User/UserRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserAuth(ctx context.Context, in *UserAuthReq, opts ...grpc.CallOption) (*UserAuthResp, error) {
	out := new(UserAuthResp)
	err := c.cc.Invoke(ctx, "/binacs_api_user_v1.User/UserAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserRefresh(ctx context.Context, in *UserRefreshReq, opts ...grpc.CallOption) (*UserRefreshResp, error) {
	out := new(UserRefreshResp)
	err := c.cc.Invoke(ctx, "/binacs_api_user_v1.User/UserRefresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error) {
	out := new(UserInfoResp)
	err := c.cc.Invoke(ctx, "/binacs_api_user_v1.User/UserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	UserTest(context.Context, *UserTestReq) (*UserTestResp, error)
	UserRegister(context.Context, *UserRegisterReq) (*UserRegisterResp, error)
	UserAuth(context.Context, *UserAuthReq) (*UserAuthResp, error)
	UserRefresh(context.Context, *UserRefreshReq) (*UserRefreshResp, error)
	UserInfo(context.Context, *UserInfoReq) (*UserInfoResp, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) UserTest(ctx context.Context, req *UserTestReq) (*UserTestResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserTest not implemented")
}
func (*UnimplementedUserServer) UserRegister(ctx context.Context, req *UserRegisterReq) (*UserRegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}
func (*UnimplementedUserServer) UserAuth(ctx context.Context, req *UserAuthReq) (*UserAuthResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAuth not implemented")
}
func (*UnimplementedUserServer) UserRefresh(ctx context.Context, req *UserRefreshReq) (*UserRefreshResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRefresh not implemented")
}
func (*UnimplementedUserServer) UserInfo(ctx context.Context, req *UserInfoReq) (*UserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_UserTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserTestReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/binacs_api_user_v1.User/UserTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserTest(ctx, req.(*UserTestReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/binacs_api_user_v1.User/UserRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserRegister(ctx, req.(*UserRegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/binacs_api_user_v1.User/UserAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserAuth(ctx, req.(*UserAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRefreshReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/binacs_api_user_v1.User/UserRefresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserRefresh(ctx, req.(*UserRefreshReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/binacs_api_user_v1.User/UserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserInfo(ctx, req.(*UserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "binacs_api_user_v1.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserTest",
			Handler:    _User_UserTest_Handler,
		},
		{
			MethodName: "UserRegister",
			Handler:    _User_UserRegister_Handler,
		},
		{
			MethodName: "UserAuth",
			Handler:    _User_UserAuth_Handler,
		},
		{
			MethodName: "UserRefresh",
			Handler:    _User_UserRefresh_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _User_UserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
