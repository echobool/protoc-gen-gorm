// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example/user/user.proto

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	example/user/user.proto

It has these top-level messages:
	User
	Email
	Address
	Language
	CreditCard
	Task
*/
package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/infobloxopen/protoc-gen-gorm/options"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id         int32                       `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	CreatedAt  *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt  *google_protobuf1.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
	Birthday   *google_protobuf1.Timestamp `protobuf:"bytes,4,opt,name=birthday" json:"birthday,omitempty"`
	Age        uint32                      `protobuf:"varint,5,opt,name=age" json:"age,omitempty"`
	Num        uint32                      `protobuf:"varint,6,opt,name=num" json:"num,omitempty"`
	CreditCard *CreditCard                 `protobuf:"bytes,7,opt,name=credit_card,json=creditCard" json:"credit_card,omitempty"`
	Emails     []*Email                    `protobuf:"bytes,8,rep,name=emails" json:"emails,omitempty"`
	Tasks      []*Task                     `protobuf:"bytes,9,rep,name=tasks" json:"tasks,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetCreatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *User) GetUpdatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *User) GetBirthday() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Birthday
	}
	return nil
}

func (m *User) GetAge() uint32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *User) GetNum() uint32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *User) GetCreditCard() *CreditCard {
	if m != nil {
		return m.CreditCard
	}
	return nil
}

func (m *User) GetEmails() []*Email {
	if m != nil {
		return m.Emails
	}
	return nil
}

func (m *User) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type Email struct {
	Id         int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Email      string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Subscribed bool   `protobuf:"varint,3,opt,name=subscribed" json:"subscribed,omitempty"`
}

func (m *Email) Reset()                    { *m = Email{} }
func (m *Email) String() string            { return proto.CompactTextString(m) }
func (*Email) ProtoMessage()               {}
func (*Email) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Email) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Email) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Email) GetSubscribed() bool {
	if m != nil {
		return m.Subscribed
	}
	return false
}

type Address struct {
	Id        int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Address_1 string `protobuf:"bytes,2,opt,name=address_1,json=address1" json:"address_1,omitempty"`
	Address_2 string `protobuf:"bytes,3,opt,name=address_2,json=address2" json:"address_2,omitempty"`
	Post      string `protobuf:"bytes,4,opt,name=post" json:"post,omitempty"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Address) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Address) GetAddress_1() string {
	if m != nil {
		return m.Address_1
	}
	return ""
}

func (m *Address) GetAddress_2() string {
	if m != nil {
		return m.Address_2
	}
	return ""
}

func (m *Address) GetPost() string {
	if m != nil {
		return m.Post
	}
	return ""
}

type Language struct {
	Id   int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Code string `protobuf:"bytes,3,opt,name=code" json:"code,omitempty"`
}

func (m *Language) Reset()                    { *m = Language{} }
func (m *Language) String() string            { return proto.CompactTextString(m) }
func (*Language) ProtoMessage()               {}
func (*Language) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Language) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Language) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Language) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type CreditCard struct {
	Id        int32                       `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	CreatedAt *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt *google_protobuf1.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
	Number    string                      `protobuf:"bytes,4,opt,name=number" json:"number,omitempty"`
}

func (m *CreditCard) Reset()                    { *m = CreditCard{} }
func (m *CreditCard) String() string            { return proto.CompactTextString(m) }
func (*CreditCard) ProtoMessage()               {}
func (*CreditCard) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreditCard) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CreditCard) GetCreatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *CreditCard) GetUpdatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *CreditCard) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

type Task struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Task) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*Email)(nil), "user.Email")
	proto.RegisterType((*Address)(nil), "user.Address")
	proto.RegisterType((*Language)(nil), "user.Language")
	proto.RegisterType((*CreditCard)(nil), "user.CreditCard")
	proto.RegisterType((*Task)(nil), "user.Task")
}

func init() { proto.RegisterFile("example/user/user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x94, 0xb1, 0x6e, 0xdb, 0x3c,
	0x10, 0xc7, 0x21, 0x5b, 0x76, 0xe4, 0xd3, 0xf7, 0x15, 0x06, 0x51, 0xb4, 0x8a, 0x0b, 0xb4, 0x82,
	0xbb, 0x78, 0x89, 0x04, 0xab, 0x40, 0x81, 0x24, 0x93, 0x9d, 0x76, 0x29, 0x3a, 0x11, 0xc9, 0xd2,
	0xc5, 0xa0, 0x44, 0x46, 0x21, 0x62, 0x89, 0x02, 0x49, 0x01, 0xc9, 0xde, 0x27, 0xea, 0x96, 0x3c,
	0x5d, 0x41, 0x8a, 0x56, 0x84, 0x66, 0x48, 0xc7, 0x2e, 0xc4, 0xf1, 0xee, 0x7f, 0xbf, 0x23, 0xef,
	0x44, 0xc1, 0x5b, 0x76, 0x47, 0xaa, 0x66, 0xcf, 0xd2, 0x56, 0x31, 0x69, 0x97, 0xa4, 0x91, 0x42,
	0x0b, 0xe4, 0x1b, 0x7b, 0x71, 0x56, 0x72, 0x7d, 0xd3, 0xe6, 0x49, 0x21, 0xaa, 0x94, 0xd7, 0xd7,
	0x22, 0xdf, 0x8b, 0x3b, 0xd1, 0xb0, 0x3a, 0xb5, 0xa2, 0xe2, 0xa4, 0x64, 0xf5, 0x49, 0x29, 0x64,
	0x95, 0x8a, 0x46, 0x73, 0x51, 0xab, 0xd4, 0x6c, 0x3a, 0xc2, 0xe2, 0x43, 0x29, 0x44, 0xb9, 0x67,
	0x9d, 0x34, 0x6f, 0xaf, 0x53, 0xcd, 0x2b, 0xa6, 0x34, 0xa9, 0x9a, 0x4e, 0xb0, 0xfc, 0x39, 0x06,
	0xff, 0x4a, 0x31, 0x89, 0x5e, 0xc1, 0x88, 0xd3, 0xc8, 0x8b, 0xbd, 0xd5, 0x04, 0x8f, 0x38, 0x45,
	0xa7, 0x00, 0x85, 0x64, 0x44, 0x33, 0xba, 0x23, 0x3a, 0x1a, 0xc5, 0xde, 0x2a, 0xcc, 0x16, 0x49,
	0x87, 0x4b, 0x0e, 0xb8, 0xe4, 0xf2, 0x80, 0xc3, 0x33, 0xa7, 0xde, 0x68, 0x93, 0xda, 0x36, 0xf4,
	0x90, 0x3a, 0x7e, 0x39, 0xd5, 0xa9, 0x37, 0x1a, 0x7d, 0x86, 0x20, 0xe7, 0x52, 0xdf, 0x50, 0x72,
	0x1f, 0xf9, 0x2f, 0x26, 0xf6, 0x5a, 0x14, 0xc1, 0x98, 0x94, 0x2c, 0x9a, 0xc4, 0xde, 0xea, 0xff,
	0xed, 0xf4, 0xf1, 0xe1, 0x78, 0x34, 0xf7, 0xb0, 0x71, 0xa1, 0x39, 0x8c, 0xeb, 0xb6, 0x8a, 0xa6,
	0x26, 0x82, 0x8d, 0x89, 0xd6, 0x10, 0x16, 0x92, 0x51, 0xae, 0x77, 0x05, 0x91, 0x34, 0x3a, 0xb2,
	0x65, 0xe6, 0x89, 0xed, 0xfb, 0x85, 0x0d, 0x5c, 0x10, 0x49, 0x31, 0x14, 0xbd, 0x8d, 0x3e, 0xc2,
	0x94, 0x55, 0x84, 0xef, 0x55, 0x14, 0xc4, 0xe3, 0x55, 0x98, 0x85, 0x9d, 0xfa, 0xab, 0xf1, 0x61,
	0x17, 0x42, 0x19, 0x4c, 0x34, 0x51, 0xb7, 0x2a, 0x9a, 0x59, 0x0d, 0x74, 0x9a, 0x4b, 0xa2, 0x6e,
	0xb7, 0xf3, 0xc7, 0x87, 0xe3, 0xff, 0x96, 0xb0, 0x0c, 0x1a, 0xc9, 0x85, 0xe4, 0xfa, 0x1e, 0x77,
	0xd2, 0x33, 0x7b, 0xd4, 0xc0, 0x5b, 0x5e, 0xc1, 0xc4, 0xc2, 0x9e, 0x8d, 0xe1, 0x35, 0x4c, 0x2c,
	0xde, 0x4e, 0x60, 0x86, 0xbb, 0x0d, 0x7a, 0x0f, 0xa0, 0xda, 0x5c, 0x15, 0x92, 0xe7, 0x8c, 0xda,
	0x0e, 0x07, 0x78, 0xe0, 0xe9, 0xb1, 0x02, 0x8e, 0x36, 0x94, 0x4a, 0xa6, 0xd4, 0x33, 0xf0, 0x3b,
	0x98, 0x91, 0x2e, 0xb4, 0x5b, 0x3b, 0x78, 0xe0, 0x1c, 0xeb, 0x61, 0x30, 0xb3, 0xf8, 0xa7, 0x60,
	0x86, 0x10, 0xf8, 0x8d, 0x50, 0xda, 0xce, 0x67, 0x86, 0xad, 0xdd, 0x17, 0xfc, 0x06, 0xc1, 0x77,
	0x52, 0x97, 0xad, 0xe9, 0xfc, 0x9f, 0x15, 0x11, 0xf8, 0x35, 0xa9, 0x98, 0x2b, 0x66, 0x6d, 0xe3,
	0x2b, 0x04, 0x65, 0xae, 0x86, 0xb5, 0x7b, 0xd6, 0x2f, 0x0f, 0xe0, 0x69, 0x1e, 0xff, 0xc8, 0x07,
	0xfa, 0x06, 0xa6, 0x75, 0x5b, 0xe5, 0x4c, 0xba, 0xeb, 0xbb, 0x5d, 0x7f, 0xe8, 0x2f, 0xe0, 0x9b,
	0x89, 0xf7, 0x97, 0xf5, 0x06, 0x97, 0x8d, 0x21, 0xa4, 0xcc, 0x8c, 0xc8, 0xbe, 0x53, 0xd7, 0x87,
	0xa1, 0xeb, 0x40, 0xd9, 0x9e, 0xff, 0x38, 0xfd, 0xdb, 0x47, 0x3f, 0xfc, 0x77, 0x9c, 0x9b, 0x25,
	0x9f, 0x5a, 0xc9, 0xa7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x23, 0xf9, 0x61, 0xfc, 0x57, 0x04,
	0x00, 0x00,
}