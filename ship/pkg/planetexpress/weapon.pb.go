// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: weapon.proto

package planetexpress

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

type Weapon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AmmoCapacity int64  `protobuf:"varint,2,opt,name=ammo_capacity,json=ammoCapacity,proto3" json:"ammo_capacity,omitempty"`
	Range        int64  `protobuf:"varint,3,opt,name=range,proto3" json:"range,omitempty"` // in miles
}

func (x *Weapon) Reset() {
	*x = Weapon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weapon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Weapon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Weapon) ProtoMessage() {}

func (x *Weapon) ProtoReflect() protoreflect.Message {
	mi := &file_weapon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Weapon.ProtoReflect.Descriptor instead.
func (*Weapon) Descriptor() ([]byte, []int) {
	return file_weapon_proto_rawDescGZIP(), []int{0}
}

func (x *Weapon) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Weapon) GetAmmoCapacity() int64 {
	if x != nil {
		return x.AmmoCapacity
	}
	return 0
}

func (x *Weapon) GetRange() int64 {
	if x != nil {
		return x.Range
	}
	return 0
}

type GetWeaponsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Weapons []*Weapon `protobuf:"bytes,1,rep,name=weapons,proto3" json:"weapons,omitempty"`
}

func (x *GetWeaponsResponse) Reset() {
	*x = GetWeaponsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weapon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWeaponsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWeaponsResponse) ProtoMessage() {}

func (x *GetWeaponsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_weapon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWeaponsResponse.ProtoReflect.Descriptor instead.
func (*GetWeaponsResponse) Descriptor() ([]byte, []int) {
	return file_weapon_proto_rawDescGZIP(), []int{1}
}

func (x *GetWeaponsResponse) GetWeapons() []*Weapon {
	if x != nil {
		return x.Weapons
	}
	return nil
}

var File_weapon_proto protoreflect.FileDescriptor

var file_weapon_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x22, 0x57, 0x0a,
	0x06, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x61,
	0x6d, 0x6d, 0x6f, 0x5f, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x61, 0x6d, 0x6d, 0x6f, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x45, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61,
	0x70, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07,
	0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x57, 0x65,
	0x61, 0x70, 0x6f, 0x6e, 0x52, 0x07, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x73, 0x42, 0x13, 0x5a,
	0x11, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x65, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_weapon_proto_rawDescOnce sync.Once
	file_weapon_proto_rawDescData = file_weapon_proto_rawDesc
)

func file_weapon_proto_rawDescGZIP() []byte {
	file_weapon_proto_rawDescOnce.Do(func() {
		file_weapon_proto_rawDescData = protoimpl.X.CompressGZIP(file_weapon_proto_rawDescData)
	})
	return file_weapon_proto_rawDescData
}

var file_weapon_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_weapon_proto_goTypes = []interface{}{
	(*Weapon)(nil),             // 0: planetexpress.Weapon
	(*GetWeaponsResponse)(nil), // 1: planetexpress.GetWeaponsResponse
}
var file_weapon_proto_depIdxs = []int32{
	0, // 0: planetexpress.GetWeaponsResponse.weapons:type_name -> planetexpress.Weapon
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_weapon_proto_init() }
func file_weapon_proto_init() {
	if File_weapon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_weapon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Weapon); i {
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
		file_weapon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWeaponsResponse); i {
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
			RawDescriptor: file_weapon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_weapon_proto_goTypes,
		DependencyIndexes: file_weapon_proto_depIdxs,
		MessageInfos:      file_weapon_proto_msgTypes,
	}.Build()
	File_weapon_proto = out.File
	file_weapon_proto_rawDesc = nil
	file_weapon_proto_goTypes = nil
	file_weapon_proto_depIdxs = nil
}
