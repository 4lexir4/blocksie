// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: proto/types.proto

package proto

import (
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

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header       *Header        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Transactions []*Transaction `protobuf:"bytes,2,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_proto_types_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Block) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   int32  `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Height    int32  `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	PrvHash   []byte `protobuf:"bytes,3,opt,name=prvHash,proto3" json:"prvHash,omitempty"`
	RoorHahs  []byte `protobuf:"bytes,4,opt,name=roorHahs,proto3" json:"roorHahs,omitempty"` // merkle root of txs
	Timestamp int64  `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_proto_types_proto_rawDescGZIP(), []int{1}
}

func (x *Header) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Header) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Header) GetPrvHash() []byte {
	if x != nil {
		return x.PrvHash
	}
	return nil
}

func (x *Header) GetRoorHahs() []byte {
	if x != nil {
		return x.RoorHahs
	}
	return nil
}

func (x *Header) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type TxInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The previous hash of the transaction
	// containing the output we want to spend
	PrvHash []byte `protobuf:"bytes,1,opt,name=prvHash,proto3" json:"prvHash,omitempty"`
	// The index of the output of the previous
	// transaction we want to spend
	PrvOutIndex uint32 `protobuf:"varint,2,opt,name=prvOutIndex,proto3" json:"prvOutIndex,omitempty"`
	PublicKey   []byte `protobuf:"bytes,3,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	Signature   []byte `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *TxInput) Reset() {
	*x = TxInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxInput) ProtoMessage() {}

func (x *TxInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxInput.ProtoReflect.Descriptor instead.
func (*TxInput) Descriptor() ([]byte, []int) {
	return file_proto_types_proto_rawDescGZIP(), []int{2}
}

func (x *TxInput) GetPrvHash() []byte {
	if x != nil {
		return x.PrvHash
	}
	return nil
}

func (x *TxInput) GetPrvOutIndex() uint32 {
	if x != nil {
		return x.PrvOutIndex
	}
	return 0
}

func (x *TxInput) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *TxInput) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type TxOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amout   int64  `protobuf:"varint,1,opt,name=amout,proto3" json:"amout,omitempty"`
	Address []byte `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *TxOutput) Reset() {
	*x = TxOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxOutput) ProtoMessage() {}

func (x *TxOutput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxOutput.ProtoReflect.Descriptor instead.
func (*TxOutput) Descriptor() ([]byte, []int) {
	return file_proto_types_proto_rawDescGZIP(), []int{3}
}

func (x *TxOutput) GetAmout() int64 {
	if x != nil {
		return x.Amout
	}
	return 0
}

func (x *TxOutput) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version int32       `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Inputs  []*TxInput  `protobuf:"bytes,2,rep,name=inputs,proto3" json:"inputs,omitempty"`
	Outputs []*TxOutput `protobuf:"bytes,3,rep,name=outputs,proto3" json:"outputs,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_proto_types_proto_rawDescGZIP(), []int{4}
}

func (x *Transaction) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Transaction) GetInputs() []*TxInput {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *Transaction) GetOutputs() []*TxOutput {
	if x != nil {
		return x.Outputs
	}
	return nil
}

var File_proto_types_proto protoreflect.FileDescriptor

var file_proto_types_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1f, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x30, 0x0a,
	0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x8e, 0x01, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x72, 0x76, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70,
	0x72, 0x76, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x72, 0x48, 0x61,
	0x68, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x72, 0x48, 0x61,
	0x68, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x22, 0x81, 0x01, 0x0a, 0x07, 0x54, 0x78, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x72, 0x76, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70,
	0x72, 0x76, 0x48, 0x61, 0x73, 0x68, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x76, 0x4f, 0x75, 0x74,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x70, 0x72, 0x76,
	0x4f, 0x75, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x22, 0x3a, 0x0a, 0x08, 0x54, 0x78, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x6d, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x61, 0x6d, 0x6f, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x22, 0x6e, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x06, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x54, 0x78, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x52, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x12, 0x23, 0x0a, 0x07, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x54,
	0x78, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73,
	0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x34,
	0x6c, 0x65, 0x78, 0x69, 0x72, 0x34, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x69, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_types_proto_rawDescOnce sync.Once
	file_proto_types_proto_rawDescData = file_proto_types_proto_rawDesc
)

func file_proto_types_proto_rawDescGZIP() []byte {
	file_proto_types_proto_rawDescOnce.Do(func() {
		file_proto_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_types_proto_rawDescData)
	})
	return file_proto_types_proto_rawDescData
}

var file_proto_types_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_types_proto_goTypes = []interface{}{
	(*Block)(nil),       // 0: Block
	(*Header)(nil),      // 1: Header
	(*TxInput)(nil),     // 2: TxInput
	(*TxOutput)(nil),    // 3: TxOutput
	(*Transaction)(nil), // 4: Transaction
}
var file_proto_types_proto_depIdxs = []int32{
	1, // 0: Block.header:type_name -> Header
	4, // 1: Block.transactions:type_name -> Transaction
	2, // 2: Transaction.inputs:type_name -> TxInput
	3, // 3: Transaction.outputs:type_name -> TxOutput
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_types_proto_init() }
func file_proto_types_proto_init() {
	if File_proto_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_proto_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_proto_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxInput); i {
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
		file_proto_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxOutput); i {
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
		file_proto_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
			RawDescriptor: file_proto_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_types_proto_goTypes,
		DependencyIndexes: file_proto_types_proto_depIdxs,
		MessageInfos:      file_proto_types_proto_msgTypes,
	}.Build()
	File_proto_types_proto = out.File
	file_proto_types_proto_rawDesc = nil
	file_proto_types_proto_goTypes = nil
	file_proto_types_proto_depIdxs = nil
}
