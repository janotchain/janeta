// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bc.proto

/*
Package bc is a generated protocol buffer package.

It is generated from these files:
	bc.proto

It has these top-level messages:
	Hash
	Program
	AssetID
	AssetAmount
	AssetDefinition
	ValueSource
	ValueDestination
	BlockHeader
	TxHeader
	Mux
	Coinbase
	OriginalOutput
	VoteOutput
	VetoInput
	Retirement
	Issuance
	Spend
*/
package bc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Hash struct {
	V0 uint64 `protobuf:"fixed64,1,opt,name=v0" json:"v0,omitempty"`
	V1 uint64 `protobuf:"fixed64,2,opt,name=v1" json:"v1,omitempty"`
	V2 uint64 `protobuf:"fixed64,3,opt,name=v2" json:"v2,omitempty"`
	V3 uint64 `protobuf:"fixed64,4,opt,name=v3" json:"v3,omitempty"`
}

func (m *Hash) Reset()                    { *m = Hash{} }
func (m *Hash) String() string            { return proto.CompactTextString(m) }
func (*Hash) ProtoMessage()               {}
func (*Hash) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Hash) GetV0() uint64 {
	if m != nil {
		return m.V0
	}
	return 0
}

func (m *Hash) GetV1() uint64 {
	if m != nil {
		return m.V1
	}
	return 0
}

func (m *Hash) GetV2() uint64 {
	if m != nil {
		return m.V2
	}
	return 0
}

func (m *Hash) GetV3() uint64 {
	if m != nil {
		return m.V3
	}
	return 0
}

type Program struct {
	VmVersion uint64 `protobuf:"varint,1,opt,name=vm_version,json=vmVersion" json:"vm_version,omitempty"`
	Code      []byte `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (m *Program) Reset()                    { *m = Program{} }
func (m *Program) String() string            { return proto.CompactTextString(m) }
func (*Program) ProtoMessage()               {}
func (*Program) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Program) GetVmVersion() uint64 {
	if m != nil {
		return m.VmVersion
	}
	return 0
}

func (m *Program) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

// This message type duplicates Hash, above. One alternative is to
// embed a Hash inside an AssetID. But it's useful for AssetID to be
// plain old data (without pointers). Another alternative is use Hash
// in any protobuf types where an AssetID is called for, but it's
// preferable to have type safety.
type AssetID struct {
	V0 uint64 `protobuf:"fixed64,1,opt,name=v0" json:"v0,omitempty"`
	V1 uint64 `protobuf:"fixed64,2,opt,name=v1" json:"v1,omitempty"`
	V2 uint64 `protobuf:"fixed64,3,opt,name=v2" json:"v2,omitempty"`
	V3 uint64 `protobuf:"fixed64,4,opt,name=v3" json:"v3,omitempty"`
}

func (m *AssetID) Reset()                    { *m = AssetID{} }
func (m *AssetID) String() string            { return proto.CompactTextString(m) }
func (*AssetID) ProtoMessage()               {}
func (*AssetID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AssetID) GetV0() uint64 {
	if m != nil {
		return m.V0
	}
	return 0
}

func (m *AssetID) GetV1() uint64 {
	if m != nil {
		return m.V1
	}
	return 0
}

func (m *AssetID) GetV2() uint64 {
	if m != nil {
		return m.V2
	}
	return 0
}

func (m *AssetID) GetV3() uint64 {
	if m != nil {
		return m.V3
	}
	return 0
}

type AssetAmount struct {
	AssetId *AssetID `protobuf:"bytes,1,opt,name=asset_id,json=assetId" json:"asset_id,omitempty"`
	Amount  uint64   `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
}

func (m *AssetAmount) Reset()                    { *m = AssetAmount{} }
func (m *AssetAmount) String() string            { return proto.CompactTextString(m) }
func (*AssetAmount) ProtoMessage()               {}
func (*AssetAmount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AssetAmount) GetAssetId() *AssetID {
	if m != nil {
		return m.AssetId
	}
	return nil
}

func (m *AssetAmount) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type AssetDefinition struct {
	IssuanceProgram *Program `protobuf:"bytes,1,opt,name=issuance_program,json=issuanceProgram" json:"issuance_program,omitempty"`
	Data            *Hash    `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
}

func (m *AssetDefinition) Reset()                    { *m = AssetDefinition{} }
func (m *AssetDefinition) String() string            { return proto.CompactTextString(m) }
func (*AssetDefinition) ProtoMessage()               {}
func (*AssetDefinition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AssetDefinition) GetIssuanceProgram() *Program {
	if m != nil {
		return m.IssuanceProgram
	}
	return nil
}

func (m *AssetDefinition) GetData() *Hash {
	if m != nil {
		return m.Data
	}
	return nil
}

type ValueSource struct {
	Ref      *Hash        `protobuf:"bytes,1,opt,name=ref" json:"ref,omitempty"`
	Value    *AssetAmount `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	Position uint64       `protobuf:"varint,3,opt,name=position" json:"position,omitempty"`
}

func (m *ValueSource) Reset()                    { *m = ValueSource{} }
func (m *ValueSource) String() string            { return proto.CompactTextString(m) }
func (*ValueSource) ProtoMessage()               {}
func (*ValueSource) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ValueSource) GetRef() *Hash {
	if m != nil {
		return m.Ref
	}
	return nil
}

func (m *ValueSource) GetValue() *AssetAmount {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ValueSource) GetPosition() uint64 {
	if m != nil {
		return m.Position
	}
	return 0
}

type ValueDestination struct {
	Ref      *Hash        `protobuf:"bytes,1,opt,name=ref" json:"ref,omitempty"`
	Value    *AssetAmount `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	Position uint64       `protobuf:"varint,3,opt,name=position" json:"position,omitempty"`
}

func (m *ValueDestination) Reset()                    { *m = ValueDestination{} }
func (m *ValueDestination) String() string            { return proto.CompactTextString(m) }
func (*ValueDestination) ProtoMessage()               {}
func (*ValueDestination) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ValueDestination) GetRef() *Hash {
	if m != nil {
		return m.Ref
	}
	return nil
}

func (m *ValueDestination) GetValue() *AssetAmount {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ValueDestination) GetPosition() uint64 {
	if m != nil {
		return m.Position
	}
	return 0
}

type BlockHeader struct {
	Version          uint64 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Height           uint64 `protobuf:"varint,2,opt,name=height" json:"height,omitempty"`
	PreviousBlockId  *Hash  `protobuf:"bytes,3,opt,name=previous_block_id,json=previousBlockId" json:"previous_block_id,omitempty"`
	Timestamp        uint64 `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	TransactionsRoot *Hash  `protobuf:"bytes,5,opt,name=transactions_root,json=transactionsRoot" json:"transactions_root,omitempty"`
}

func (m *BlockHeader) Reset()                    { *m = BlockHeader{} }
func (m *BlockHeader) String() string            { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()               {}
func (*BlockHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *BlockHeader) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *BlockHeader) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockHeader) GetPreviousBlockId() *Hash {
	if m != nil {
		return m.PreviousBlockId
	}
	return nil
}

func (m *BlockHeader) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *BlockHeader) GetTransactionsRoot() *Hash {
	if m != nil {
		return m.TransactionsRoot
	}
	return nil
}

type TxHeader struct {
	Version        uint64  `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	SerializedSize uint64  `protobuf:"varint,2,opt,name=serialized_size,json=serializedSize" json:"serialized_size,omitempty"`
	TimeRange      uint64  `protobuf:"varint,3,opt,name=time_range,json=timeRange" json:"time_range,omitempty"`
	ResultIds      []*Hash `protobuf:"bytes,4,rep,name=result_ids,json=resultIds" json:"result_ids,omitempty"`
}

func (m *TxHeader) Reset()                    { *m = TxHeader{} }
func (m *TxHeader) String() string            { return proto.CompactTextString(m) }
func (*TxHeader) ProtoMessage()               {}
func (*TxHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *TxHeader) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *TxHeader) GetSerializedSize() uint64 {
	if m != nil {
		return m.SerializedSize
	}
	return 0
}

func (m *TxHeader) GetTimeRange() uint64 {
	if m != nil {
		return m.TimeRange
	}
	return 0
}

func (m *TxHeader) GetResultIds() []*Hash {
	if m != nil {
		return m.ResultIds
	}
	return nil
}

type Mux struct {
	Sources             []*ValueSource      `protobuf:"bytes,1,rep,name=sources" json:"sources,omitempty"`
	Program             *Program            `protobuf:"bytes,2,opt,name=program" json:"program,omitempty"`
	WitnessDestinations []*ValueDestination `protobuf:"bytes,3,rep,name=witness_destinations,json=witnessDestinations" json:"witness_destinations,omitempty"`
	WitnessArguments    [][]byte            `protobuf:"bytes,4,rep,name=witness_arguments,json=witnessArguments,proto3" json:"witness_arguments,omitempty"`
}

func (m *Mux) Reset()                    { *m = Mux{} }
func (m *Mux) String() string            { return proto.CompactTextString(m) }
func (*Mux) ProtoMessage()               {}
func (*Mux) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Mux) GetSources() []*ValueSource {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *Mux) GetProgram() *Program {
	if m != nil {
		return m.Program
	}
	return nil
}

func (m *Mux) GetWitnessDestinations() []*ValueDestination {
	if m != nil {
		return m.WitnessDestinations
	}
	return nil
}

func (m *Mux) GetWitnessArguments() [][]byte {
	if m != nil {
		return m.WitnessArguments
	}
	return nil
}

type Coinbase struct {
	WitnessDestination *ValueDestination `protobuf:"bytes,1,opt,name=witness_destination,json=witnessDestination" json:"witness_destination,omitempty"`
	Arbitrary          []byte            `protobuf:"bytes,2,opt,name=arbitrary,proto3" json:"arbitrary,omitempty"`
}

func (m *Coinbase) Reset()                    { *m = Coinbase{} }
func (m *Coinbase) String() string            { return proto.CompactTextString(m) }
func (*Coinbase) ProtoMessage()               {}
func (*Coinbase) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Coinbase) GetWitnessDestination() *ValueDestination {
	if m != nil {
		return m.WitnessDestination
	}
	return nil
}

func (m *Coinbase) GetArbitrary() []byte {
	if m != nil {
		return m.Arbitrary
	}
	return nil
}

type OriginalOutput struct {
	Source         *ValueSource `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	ControlProgram *Program     `protobuf:"bytes,2,opt,name=control_program,json=controlProgram" json:"control_program,omitempty"`
	Ordinal        uint64       `protobuf:"varint,3,opt,name=ordinal" json:"ordinal,omitempty"`
	StateData      [][]byte     `protobuf:"bytes,4,rep,name=state_data,json=stateData,proto3" json:"state_data,omitempty"`
}

func (m *OriginalOutput) Reset()                    { *m = OriginalOutput{} }
func (m *OriginalOutput) String() string            { return proto.CompactTextString(m) }
func (*OriginalOutput) ProtoMessage()               {}
func (*OriginalOutput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *OriginalOutput) GetSource() *ValueSource {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *OriginalOutput) GetControlProgram() *Program {
	if m != nil {
		return m.ControlProgram
	}
	return nil
}

func (m *OriginalOutput) GetOrdinal() uint64 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

func (m *OriginalOutput) GetStateData() [][]byte {
	if m != nil {
		return m.StateData
	}
	return nil
}

type VoteOutput struct {
	Source         *ValueSource `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	ControlProgram *Program     `protobuf:"bytes,2,opt,name=control_program,json=controlProgram" json:"control_program,omitempty"`
	Ordinal        uint64       `protobuf:"varint,3,opt,name=ordinal" json:"ordinal,omitempty"`
	Vote           []byte       `protobuf:"bytes,4,opt,name=vote,proto3" json:"vote,omitempty"`
	StateData      [][]byte     `protobuf:"bytes,5,rep,name=state_data,json=stateData,proto3" json:"state_data,omitempty"`
}

func (m *VoteOutput) Reset()                    { *m = VoteOutput{} }
func (m *VoteOutput) String() string            { return proto.CompactTextString(m) }
func (*VoteOutput) ProtoMessage()               {}
func (*VoteOutput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *VoteOutput) GetSource() *ValueSource {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *VoteOutput) GetControlProgram() *Program {
	if m != nil {
		return m.ControlProgram
	}
	return nil
}

func (m *VoteOutput) GetOrdinal() uint64 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

func (m *VoteOutput) GetVote() []byte {
	if m != nil {
		return m.Vote
	}
	return nil
}

func (m *VoteOutput) GetStateData() [][]byte {
	if m != nil {
		return m.StateData
	}
	return nil
}

type VetoInput struct {
	SpentOutputId      *Hash             `protobuf:"bytes,1,opt,name=spent_output_id,json=spentOutputId" json:"spent_output_id,omitempty"`
	WitnessDestination *ValueDestination `protobuf:"bytes,2,opt,name=witness_destination,json=witnessDestination" json:"witness_destination,omitempty"`
	WitnessArguments   [][]byte          `protobuf:"bytes,3,rep,name=witness_arguments,json=witnessArguments,proto3" json:"witness_arguments,omitempty"`
	Ordinal            uint64            `protobuf:"varint,4,opt,name=ordinal" json:"ordinal,omitempty"`
}

func (m *VetoInput) Reset()                    { *m = VetoInput{} }
func (m *VetoInput) String() string            { return proto.CompactTextString(m) }
func (*VetoInput) ProtoMessage()               {}
func (*VetoInput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *VetoInput) GetSpentOutputId() *Hash {
	if m != nil {
		return m.SpentOutputId
	}
	return nil
}

func (m *VetoInput) GetWitnessDestination() *ValueDestination {
	if m != nil {
		return m.WitnessDestination
	}
	return nil
}

func (m *VetoInput) GetWitnessArguments() [][]byte {
	if m != nil {
		return m.WitnessArguments
	}
	return nil
}

func (m *VetoInput) GetOrdinal() uint64 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

type Retirement struct {
	Source  *ValueSource `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	Ordinal uint64       `protobuf:"varint,2,opt,name=ordinal" json:"ordinal,omitempty"`
}

func (m *Retirement) Reset()                    { *m = Retirement{} }
func (m *Retirement) String() string            { return proto.CompactTextString(m) }
func (*Retirement) ProtoMessage()               {}
func (*Retirement) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *Retirement) GetSource() *ValueSource {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Retirement) GetOrdinal() uint64 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

type Issuance struct {
	NonceHash              *Hash             `protobuf:"bytes,1,opt,name=nonce_hash,json=nonceHash" json:"nonce_hash,omitempty"`
	Value                  *AssetAmount      `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	WitnessDestination     *ValueDestination `protobuf:"bytes,3,opt,name=witness_destination,json=witnessDestination" json:"witness_destination,omitempty"`
	WitnessAssetDefinition *AssetDefinition  `protobuf:"bytes,4,opt,name=witness_asset_definition,json=witnessAssetDefinition" json:"witness_asset_definition,omitempty"`
	WitnessArguments       [][]byte          `protobuf:"bytes,5,rep,name=witness_arguments,json=witnessArguments,proto3" json:"witness_arguments,omitempty"`
	Ordinal                uint64            `protobuf:"varint,6,opt,name=ordinal" json:"ordinal,omitempty"`
}

func (m *Issuance) Reset()                    { *m = Issuance{} }
func (m *Issuance) String() string            { return proto.CompactTextString(m) }
func (*Issuance) ProtoMessage()               {}
func (*Issuance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *Issuance) GetNonceHash() *Hash {
	if m != nil {
		return m.NonceHash
	}
	return nil
}

func (m *Issuance) GetValue() *AssetAmount {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Issuance) GetWitnessDestination() *ValueDestination {
	if m != nil {
		return m.WitnessDestination
	}
	return nil
}

func (m *Issuance) GetWitnessAssetDefinition() *AssetDefinition {
	if m != nil {
		return m.WitnessAssetDefinition
	}
	return nil
}

func (m *Issuance) GetWitnessArguments() [][]byte {
	if m != nil {
		return m.WitnessArguments
	}
	return nil
}

func (m *Issuance) GetOrdinal() uint64 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

type Spend struct {
	SpentOutputId      *Hash             `protobuf:"bytes,1,opt,name=spent_output_id,json=spentOutputId" json:"spent_output_id,omitempty"`
	WitnessDestination *ValueDestination `protobuf:"bytes,2,opt,name=witness_destination,json=witnessDestination" json:"witness_destination,omitempty"`
	WitnessArguments   [][]byte          `protobuf:"bytes,3,rep,name=witness_arguments,json=witnessArguments,proto3" json:"witness_arguments,omitempty"`
	Ordinal            uint64            `protobuf:"varint,4,opt,name=ordinal" json:"ordinal,omitempty"`
}

func (m *Spend) Reset()                    { *m = Spend{} }
func (m *Spend) String() string            { return proto.CompactTextString(m) }
func (*Spend) ProtoMessage()               {}
func (*Spend) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *Spend) GetSpentOutputId() *Hash {
	if m != nil {
		return m.SpentOutputId
	}
	return nil
}

func (m *Spend) GetWitnessDestination() *ValueDestination {
	if m != nil {
		return m.WitnessDestination
	}
	return nil
}

func (m *Spend) GetWitnessArguments() [][]byte {
	if m != nil {
		return m.WitnessArguments
	}
	return nil
}

func (m *Spend) GetOrdinal() uint64 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

func init() {
	proto.RegisterType((*Hash)(nil), "bc.Hash")
	proto.RegisterType((*Program)(nil), "bc.Program")
	proto.RegisterType((*AssetID)(nil), "bc.AssetID")
	proto.RegisterType((*AssetAmount)(nil), "bc.AssetAmount")
	proto.RegisterType((*AssetDefinition)(nil), "bc.AssetDefinition")
	proto.RegisterType((*ValueSource)(nil), "bc.ValueSource")
	proto.RegisterType((*ValueDestination)(nil), "bc.ValueDestination")
	proto.RegisterType((*BlockHeader)(nil), "bc.BlockHeader")
	proto.RegisterType((*TxHeader)(nil), "bc.TxHeader")
	proto.RegisterType((*Mux)(nil), "bc.Mux")
	proto.RegisterType((*Coinbase)(nil), "bc.Coinbase")
	proto.RegisterType((*OriginalOutput)(nil), "bc.OriginalOutput")
	proto.RegisterType((*VoteOutput)(nil), "bc.VoteOutput")
	proto.RegisterType((*VetoInput)(nil), "bc.VetoInput")
	proto.RegisterType((*Retirement)(nil), "bc.Retirement")
	proto.RegisterType((*Issuance)(nil), "bc.Issuance")
	proto.RegisterType((*Spend)(nil), "bc.Spend")
}

func init() { proto.RegisterFile("bc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 852 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xd6, 0xda, 0xeb, 0xbf, 0xe3, 0x10, 0x27, 0x93, 0xaa, 0x5a, 0x55, 0x45, 0xaa, 0x56, 0x2a,
	0x01, 0x21, 0x45, 0xa9, 0x53, 0xb8, 0xe2, 0x26, 0x10, 0xa0, 0xbe, 0x88, 0x82, 0x26, 0x28, 0xb7,
	0xab, 0xf1, 0xee, 0xd4, 0x1e, 0x61, 0xcf, 0x2c, 0x33, 0xb3, 0x4b, 0xc9, 0x7b, 0xf0, 0x0e, 0xbc,
	0x01, 0xbc, 0x01, 0x12, 0x88, 0x77, 0x42, 0x73, 0x76, 0xd6, 0xeb, 0xbf, 0xd2, 0x56, 0x08, 0x21,
	0xee, 0x7c, 0x7e, 0xf6, 0x9c, 0xef, 0x7c, 0xe7, 0x9b, 0x19, 0x43, 0x7f, 0x9a, 0x9e, 0xe5, 0x5a,
	0x59, 0x45, 0x5a, 0xd3, 0x34, 0xfe, 0x0a, 0xc2, 0x17, 0xcc, 0xcc, 0xc9, 0x21, 0xb4, 0xca, 0xf3,
	0x28, 0x78, 0x12, 0x7c, 0xd8, 0xa5, 0xad, 0xf2, 0x1c, 0xed, 0x67, 0x51, 0xcb, 0xdb, 0xcf, 0xd0,
	0x1e, 0x47, 0x6d, 0x6f, 0x8f, 0xd1, 0xbe, 0x88, 0x42, 0x6f, 0x5f, 0xc4, 0x9f, 0x41, 0xef, 0x1b,
	0xad, 0x66, 0x9a, 0x2d, 0xc9, 0xfb, 0x00, 0xe5, 0x32, 0x29, 0xb9, 0x36, 0x42, 0x49, 0x2c, 0x19,
	0xd2, 0x41, 0xb9, 0xbc, 0xab, 0x1c, 0x84, 0x40, 0x98, 0xaa, 0x8c, 0x63, 0xed, 0x03, 0x8a, 0xbf,
	0xe3, 0x09, 0xf4, 0x2e, 0x8d, 0xe1, 0x76, 0x72, 0xf5, 0x8f, 0x81, 0x5c, 0xc3, 0x10, 0x4b, 0x5d,
	0x2e, 0x55, 0x21, 0x2d, 0xf9, 0x00, 0xfa, 0xcc, 0x99, 0x89, 0xc8, 0xb0, 0xe8, 0x70, 0x3c, 0x3c,
	0x9b, 0xa6, 0x67, 0xbe, 0x1b, 0xed, 0x61, 0x70, 0x92, 0x91, 0x87, 0xd0, 0x65, 0xf8, 0x05, 0xb6,
	0x0a, 0xa9, 0xb7, 0xe2, 0x19, 0x8c, 0x30, 0xf7, 0x8a, 0xbf, 0x14, 0x52, 0x58, 0x37, 0xc0, 0xa7,
	0x70, 0x24, 0x8c, 0x29, 0x98, 0x4c, 0x79, 0x92, 0x57, 0x33, 0xaf, 0x97, 0xf6, 0x34, 0xd0, 0x51,
	0x9d, 0x54, 0xf3, 0xf2, 0x18, 0xc2, 0x8c, 0x59, 0x86, 0x0d, 0x86, 0xe3, 0xbe, 0xcb, 0x75, 0xd4,
	0x53, 0xf4, 0xc6, 0x0b, 0x18, 0xde, 0xb1, 0x45, 0xc1, 0x6f, 0x55, 0xa1, 0x53, 0x4e, 0x1e, 0x41,
	0x5b, 0xf3, 0x97, 0xbe, 0x6e, 0x93, 0xeb, 0x9c, 0xe4, 0x29, 0x74, 0x4a, 0x97, 0xea, 0x2b, 0x8d,
	0x56, 0x03, 0x55, 0x33, 0xd3, 0x2a, 0x4a, 0x1e, 0x41, 0x3f, 0x57, 0x06, 0x31, 0x23, 0x5f, 0x21,
	0x5d, 0xd9, 0xf1, 0xf7, 0x70, 0x84, 0xdd, 0xae, 0xb8, 0xb1, 0x42, 0x32, 0x9c, 0xeb, 0x5f, 0x6e,
	0xf9, 0x7b, 0x00, 0xc3, 0xcf, 0x17, 0x2a, 0xfd, 0xee, 0x05, 0x67, 0x19, 0xd7, 0x24, 0x82, 0xde,
	0xa6, 0x46, 0x6a, 0xd3, 0xed, 0x62, 0xce, 0xc5, 0x6c, 0xbe, 0xda, 0x45, 0x65, 0x91, 0xe7, 0x70,
	0x9c, 0x6b, 0x5e, 0x0a, 0x55, 0x98, 0x64, 0xea, 0x2a, 0xb9, 0xa5, 0xb6, 0xb7, 0xe0, 0x8e, 0xea,
	0x14, 0xec, 0x35, 0xc9, 0xc8, 0x63, 0x18, 0x58, 0xb1, 0xe4, 0xc6, 0xb2, 0x65, 0x8e, 0x3a, 0x09,
	0x69, 0xe3, 0x20, 0x9f, 0xc0, 0xb1, 0xd5, 0x4c, 0x1a, 0x96, 0x3a, 0x90, 0x26, 0xd1, 0x4a, 0xd9,
	0xa8, 0xb3, 0x55, 0xf3, 0x68, 0x3d, 0x85, 0x2a, 0x65, 0xe3, 0x9f, 0x02, 0xe8, 0x7f, 0xfb, 0xea,
	0x8d, 0x93, 0x9c, 0xc2, 0xc8, 0x70, 0x2d, 0xd8, 0x42, 0xdc, 0xf3, 0x2c, 0x31, 0xe2, 0x9e, 0xfb,
	0x91, 0x0e, 0x1b, 0xf7, 0xad, 0xb8, 0xe7, 0xee, 0xcc, 0x38, 0x4c, 0x89, 0x66, 0x72, 0xc6, 0x3d,
	0x75, 0x88, 0x92, 0x3a, 0x07, 0x39, 0x05, 0xd0, 0xdc, 0x14, 0x0b, 0x27, 0x63, 0x13, 0x85, 0x4f,
	0xda, 0x1b, 0xf0, 0x06, 0x55, 0x6c, 0x92, 0x99, 0xf8, 0xcf, 0x00, 0xda, 0xd7, 0xc5, 0x2b, 0xf2,
	0x11, 0xf4, 0x0c, 0x0a, 0xc9, 0x44, 0x01, 0x66, 0xe3, 0xc6, 0xd6, 0x04, 0x46, 0xeb, 0x38, 0x79,
	0x0a, 0xbd, 0x5a, 0xc5, 0xad, 0x5d, 0x15, 0xd7, 0x31, 0xf2, 0x35, 0x3c, 0xf8, 0x41, 0x58, 0xc9,
	0x8d, 0x49, 0xb2, 0x46, 0x34, 0x26, 0x6a, 0x63, 0xf9, 0x07, 0xab, 0xf2, 0x6b, 0x8a, 0xa2, 0x27,
	0xfe, 0x8b, 0x35, 0x9f, 0x21, 0x1f, 0xc3, 0x71, 0x5d, 0x88, 0xe9, 0x59, 0xb1, 0xe4, 0xd2, 0x56,
	0x23, 0x1d, 0xd0, 0x23, 0x1f, 0xb8, 0xac, 0xfd, 0xb1, 0x82, 0xfe, 0x17, 0x4a, 0xc8, 0x29, 0x33,
	0x9c, 0x7c, 0x09, 0x27, 0x7b, 0x10, 0x78, 0xbd, 0xee, 0x07, 0x40, 0x76, 0x01, 0x38, 0x3d, 0x30,
	0x3d, 0x15, 0x56, 0x33, 0xfd, 0xa3, 0xbf, 0x84, 0x1a, 0x47, 0xfc, 0x73, 0x00, 0x87, 0x37, 0x5a,
	0xcc, 0x84, 0x64, 0x8b, 0x9b, 0xc2, 0xe6, 0x85, 0x25, 0xa7, 0xd0, 0xad, 0xb8, 0xf2, 0xad, 0x76,
	0xa8, 0xf4, 0x61, 0xf2, 0x1c, 0x46, 0xa9, 0x92, 0x56, 0xab, 0x45, 0xf2, 0x37, 0x8c, 0x1e, 0xfa,
	0x9c, 0xfa, 0x5a, 0x88, 0xa0, 0xa7, 0x74, 0xe6, 0xfa, 0xf9, 0xbd, 0xd7, 0xa6, 0x13, 0x85, 0xb1,
	0xcc, 0xf2, 0x04, 0xaf, 0x8d, 0x8a, 0xa2, 0x01, 0x7a, 0xae, 0xdc, 0x8d, 0xf1, 0x6b, 0x00, 0x70,
	0xa7, 0x2c, 0xff, 0xaf, 0x61, 0x12, 0x08, 0x4b, 0x65, 0x39, 0x9e, 0xad, 0x03, 0x8a, 0xbf, 0xb7,
	0xa0, 0x77, 0xb6, 0xa1, 0xff, 0x11, 0xc0, 0xe0, 0x8e, 0x5b, 0x35, 0x91, 0x0e, 0xf9, 0x39, 0x8c,
	0x4c, 0xce, 0xa5, 0x4d, 0x14, 0x4e, 0xd2, 0x5c, 0xd5, 0x8d, 0xc4, 0xdf, 0xc3, 0x84, 0x6a, 0xd2,
	0x49, 0xf6, 0x3a, 0x29, 0xb4, 0xde, 0x51, 0x0a, 0x7b, 0xa5, 0xd8, 0xde, 0x2f, 0xc5, 0x75, 0x02,
	0xc2, 0x0d, 0x02, 0xe2, 0x1b, 0x00, 0xca, 0xad, 0xd0, 0xdc, 0x25, 0xbe, 0xfd, 0x1e, 0xd6, 0x0a,
	0xb6, 0x36, 0x0b, 0xfe, 0xd2, 0x82, 0xfe, 0xc4, 0xbf, 0x1e, 0xee, 0xec, 0x4b, 0xe5, 0xde, 0x9a,
	0x39, 0x33, 0xf3, 0x1d, 0x62, 0x06, 0x18, 0xc3, 0x27, 0xfc, 0x2d, 0xef, 0xe8, 0xd7, 0x70, 0xd7,
	0x7e, 0x47, 0xee, 0xae, 0x21, 0x5a, 0x71, 0x87, 0x0f, 0x6c, 0xb6, 0x7a, 0x21, 0x91, 0x9f, 0xe1,
	0xf8, 0x64, 0x05, 0xa0, 0x79, 0x3c, 0xe9, 0xc3, 0x9a, 0xd7, 0xad, 0x47, 0x75, 0xef, 0x2a, 0x3a,
	0x6f, 0x5e, 0x45, 0x77, 0x93, 0xb9, 0xdf, 0x02, 0xe8, 0xdc, 0xe6, 0x5c, 0x66, 0xff, 0x77, 0x51,
	0x4d, 0xbb, 0xf8, 0x1f, 0xed, 0xe2, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xe5, 0xe4, 0x4f,
	0xaf, 0x09, 0x00, 0x00,
}
