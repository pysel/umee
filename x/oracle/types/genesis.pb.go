// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: umee/oracle/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the oracle module's genesis state.
type GenesisState struct {
	Params                        Params                         `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	FeederDelegations             []FeederDelegation             `protobuf:"bytes,2,rep,name=feeder_delegations,json=feederDelegations,proto3" json:"feeder_delegations"`
	ExchangeRates                 ExchangeRateTuples             `protobuf:"bytes,3,rep,name=exchange_rates,json=exchangeRates,proto3,castrepeated=ExchangeRateTuples" json:"exchange_rates"`
	MissCounters                  []MissCounter                  `protobuf:"bytes,4,rep,name=miss_counters,json=missCounters,proto3" json:"miss_counters"`
	AggregateExchangeRatePrevotes []AggregateExchangeRatePrevote `protobuf:"bytes,5,rep,name=aggregate_exchange_rate_prevotes,json=aggregateExchangeRatePrevotes,proto3" json:"aggregate_exchange_rate_prevotes"`
	AggregateExchangeRateVotes    []AggregateExchangeRateVote    `protobuf:"bytes,6,rep,name=aggregate_exchange_rate_votes,json=aggregateExchangeRateVotes,proto3" json:"aggregate_exchange_rate_votes"`
	Medians                       []Price                        `protobuf:"bytes,7,rep,name=medians,proto3" json:"medians"`
	HistoricPrices                []Price                        `protobuf:"bytes,8,rep,name=historic_prices,json=historicPrices,proto3" json:"historic_prices"`
	MedianDeviations              []Price                        `protobuf:"bytes,9,rep,name=medianDeviations,proto3" json:"medianDeviations"`
	// Historic Avg Counter params
	AvgCounterParams AvgCounterParams `protobuf:"bytes,10,opt,name=avg_counter_params,json=avgCounterParams,proto3" json:"avg_counter_params" yaml:"avg_counter_params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_c99b4af40468acc1, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

// FeederDelegation is the address for where oracle feeder authority are
// delegated to. By default this struct is only used at genesis to feed in
// default feeder addresses.
type FeederDelegation struct {
	FeederAddress    string `protobuf:"bytes,1,opt,name=feeder_address,json=feederAddress,proto3" json:"feeder_address,omitempty"`
	ValidatorAddress string `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
}

func (m *FeederDelegation) Reset()         { *m = FeederDelegation{} }
func (m *FeederDelegation) String() string { return proto.CompactTextString(m) }
func (*FeederDelegation) ProtoMessage()    {}
func (*FeederDelegation) Descriptor() ([]byte, []int) {
	return fileDescriptor_c99b4af40468acc1, []int{1}
}
func (m *FeederDelegation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FeederDelegation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeederDelegation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FeederDelegation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeederDelegation.Merge(m, src)
}
func (m *FeederDelegation) XXX_Size() int {
	return m.Size()
}
func (m *FeederDelegation) XXX_DiscardUnknown() {
	xxx_messageInfo_FeederDelegation.DiscardUnknown(m)
}

var xxx_messageInfo_FeederDelegation proto.InternalMessageInfo

// MissCounter defines an miss counter and validator address pair used in
// oracle module's genesis state
type MissCounter struct {
	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	MissCounter      uint64 `protobuf:"varint,2,opt,name=miss_counter,json=missCounter,proto3" json:"miss_counter,omitempty"`
}

func (m *MissCounter) Reset()         { *m = MissCounter{} }
func (m *MissCounter) String() string { return proto.CompactTextString(m) }
func (*MissCounter) ProtoMessage()    {}
func (*MissCounter) Descriptor() ([]byte, []int) {
	return fileDescriptor_c99b4af40468acc1, []int{2}
}
func (m *MissCounter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MissCounter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MissCounter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MissCounter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MissCounter.Merge(m, src)
}
func (m *MissCounter) XXX_Size() int {
	return m.Size()
}
func (m *MissCounter) XXX_DiscardUnknown() {
	xxx_messageInfo_MissCounter.DiscardUnknown(m)
}

var xxx_messageInfo_MissCounter proto.InternalMessageInfo

// Price is an instance of a price "stamp"
type Price struct {
	ExchangeRateTuple ExchangeRateTuple `protobuf:"bytes,1,opt,name=exchange_rate_tuple,json=exchangeRateTuple,proto3" json:"exchange_rate_tuple"`
	BlockNum          uint64            `protobuf:"varint,2,opt,name=block_num,json=blockNum,proto3" json:"block_num,omitempty"`
}

func (m *Price) Reset()         { *m = Price{} }
func (m *Price) String() string { return proto.CompactTextString(m) }
func (*Price) ProtoMessage()    {}
func (*Price) Descriptor() ([]byte, []int) {
	return fileDescriptor_c99b4af40468acc1, []int{3}
}
func (m *Price) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Price) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Price.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Price) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Price.Merge(m, src)
}
func (m *Price) XXX_Size() int {
	return m.Size()
}
func (m *Price) XXX_DiscardUnknown() {
	xxx_messageInfo_Price.DiscardUnknown(m)
}

var xxx_messageInfo_Price proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenesisState)(nil), "umee.oracle.v1.GenesisState")
	proto.RegisterType((*FeederDelegation)(nil), "umee.oracle.v1.FeederDelegation")
	proto.RegisterType((*MissCounter)(nil), "umee.oracle.v1.MissCounter")
	proto.RegisterType((*Price)(nil), "umee.oracle.v1.Price")
}

func init() { proto.RegisterFile("umee/oracle/v1/genesis.proto", fileDescriptor_c99b4af40468acc1) }

var fileDescriptor_c99b4af40468acc1 = []byte{
	// 620 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x4e, 0xd4, 0x4e,
	0x14, 0xdf, 0xf2, 0xb1, 0xc0, 0x2c, 0xec, 0x7f, 0x99, 0xbf, 0x9a, 0xba, 0x48, 0x59, 0x9a, 0x98,
	0x60, 0xd4, 0x36, 0xe0, 0xc7, 0x85, 0x77, 0x20, 0xc2, 0x8d, 0x1a, 0x52, 0xbf, 0x12, 0x13, 0xd3,
	0x0c, 0xed, 0xa1, 0x34, 0xb4, 0x9d, 0x3a, 0x33, 0xad, 0x10, 0xe3, 0x3b, 0xf8, 0x1c, 0x3e, 0x09,
	0x97, 0x5c, 0x7a, 0x85, 0x0a, 0x6f, 0xa0, 0x2f, 0x60, 0x3a, 0x9d, 0xc2, 0xd2, 0x05, 0xe1, 0x6e,
	0xf7, 0x9c, 0xdf, 0xc7, 0x49, 0xe7, 0x77, 0x0e, 0xba, 0x95, 0xc5, 0x00, 0x36, 0x65, 0xc4, 0x8b,
	0xc0, 0xce, 0x17, 0xed, 0x00, 0x12, 0xe0, 0x21, 0xb7, 0x52, 0x46, 0x05, 0xc5, 0xed, 0xa2, 0x6b,
	0x95, 0x5d, 0x2b, 0x5f, 0xec, 0x5e, 0x0b, 0x68, 0x40, 0x65, 0xcb, 0x2e, 0x7e, 0x95, 0xa8, 0xee,
	0x4c, 0x4d, 0x43, 0xe1, 0x65, 0xd3, 0xfc, 0xd3, 0x44, 0x93, 0xeb, 0xa5, 0xe8, 0x2b, 0x41, 0x04,
	0xe0, 0x87, 0xa8, 0x99, 0x12, 0x46, 0x62, 0xae, 0x6b, 0x3d, 0x6d, 0xa1, 0xb5, 0x74, 0xc3, 0x3a,
	0x6b, 0x62, 0x6d, 0xc8, 0xee, 0xca, 0xc8, 0xfe, 0xe1, 0x5c, 0xc3, 0x51, 0x58, 0xfc, 0x06, 0xe1,
	0x2d, 0x00, 0x1f, 0x98, 0xeb, 0x43, 0x04, 0x01, 0x11, 0x21, 0x4d, 0xb8, 0x3e, 0xd4, 0x1b, 0x5e,
	0x68, 0x2d, 0xf5, 0xea, 0x0a, 0x6b, 0x12, 0xb9, 0x7a, 0x02, 0x54, 0x5a, 0xd3, 0x5b, 0xb5, 0x3a,
	0xc7, 0x3e, 0x6a, 0xc3, 0xae, 0xb7, 0x4d, 0x92, 0x00, 0x5c, 0x46, 0x04, 0x70, 0x7d, 0x58, 0x4a,
	0xce, 0xd7, 0x25, 0x9f, 0x29, 0x94, 0x43, 0x04, 0xbc, 0xce, 0xd2, 0x08, 0x56, 0xba, 0x85, 0xe6,
	0xb7, 0x1f, 0x73, 0x78, 0xa0, 0xc5, 0x9d, 0x29, 0xe8, 0xab, 0x71, 0xbc, 0x86, 0xa6, 0xe2, 0x90,
	0x73, 0xd7, 0xa3, 0x59, 0x22, 0x80, 0x71, 0x7d, 0x44, 0x9a, 0xcc, 0xd4, 0x4d, 0x5e, 0x84, 0x9c,
	0x3f, 0x2d, 0x31, 0x6a, 0xe4, 0xc9, 0xf8, 0xb4, 0xc4, 0xf1, 0x67, 0xd4, 0x23, 0x41, 0xc0, 0x8a,
	0xe9, 0xc1, 0x3d, 0x33, 0xb7, 0x9b, 0x32, 0xc8, 0x69, 0x31, 0xff, 0xa8, 0x94, 0xbe, 0x57, 0x97,
	0x5e, 0xae, 0x78, 0xfd, 0xd3, 0x6e, 0x94, 0x24, 0xe5, 0x35, 0x4b, 0xfe, 0x81, 0xe1, 0x98, 0xa1,
	0xd9, 0x8b, 0xcc, 0x4b, 0xe7, 0xa6, 0x74, 0xbe, 0x73, 0x25, 0xe7, 0xb7, 0xa7, 0xb6, 0x5d, 0x72,
	0x11, 0x80, 0xe3, 0x47, 0x68, 0x2c, 0x06, 0x3f, 0x24, 0x09, 0xd7, 0xc7, 0xa4, 0xfa, 0xf5, 0x81,
	0xb0, 0xb0, 0xd0, 0xab, 0x94, 0x2a, 0x2c, 0x5e, 0x45, 0xff, 0x6d, 0x87, 0x5c, 0x50, 0x16, 0x7a,
	0x6e, 0x5a, 0x00, 0xb8, 0x3e, 0x7e, 0x39, 0xbd, 0x5d, 0x71, 0x64, 0x91, 0xe3, 0x75, 0xd4, 0x29,
	0x05, 0x57, 0x21, 0x0f, 0x55, 0xe0, 0x26, 0x2e, 0x97, 0x19, 0x20, 0xe1, 0x8f, 0x08, 0x93, 0x3c,
	0xa8, 0x5e, 0xdf, 0x55, 0xe9, 0x47, 0x32, 0xfd, 0x03, 0xd9, 0x5d, 0xce, 0x03, 0xf5, 0xde, 0x6a,
	0x0f, 0xe6, 0x0b, 0xd5, 0xdf, 0x87, 0x73, 0x37, 0xf7, 0x48, 0x1c, 0x3d, 0x31, 0x07, 0x95, 0x4c,
	0xa7, 0x43, 0x6a, 0x24, 0x73, 0x0b, 0x75, 0xea, 0x4b, 0x80, 0x6f, 0xa3, 0xb6, 0x5a, 0x21, 0xe2,
	0xfb, 0x0c, 0x78, 0xb9, 0x80, 0x13, 0xce, 0x54, 0x59, 0x5d, 0x2e, 0x8b, 0xf8, 0x2e, 0x9a, 0xce,
	0x49, 0x14, 0xfa, 0x44, 0xd0, 0x53, 0xe4, 0x90, 0x44, 0x76, 0x4e, 0x1a, 0x0a, 0x6c, 0x7e, 0x40,
	0xad, 0xbe, 0xd0, 0x9e, 0xcf, 0xd5, 0xce, 0xe7, 0xe2, 0x79, 0x34, 0xd9, 0xbf, 0x15, 0xd2, 0x63,
	0xc4, 0x69, 0xf5, 0x25, 0xde, 0xfc, 0x82, 0x46, 0xe5, 0xa7, 0xc5, 0xef, 0xd0, 0xff, 0x67, 0x23,
	0x27, 0x8a, 0x45, 0x53, 0x17, 0xe4, 0x0a, 0xcb, 0xaa, 0x0e, 0x00, 0xd4, 0x1b, 0x78, 0x06, 0x4d,
	0x6c, 0x46, 0xd4, 0xdb, 0x71, 0x93, 0x2c, 0x56, 0x13, 0x8c, 0xcb, 0xc2, 0xcb, 0x2c, 0x5e, 0x79,
	0xbe, 0xff, 0xcb, 0x68, 0xec, 0x1f, 0x19, 0xda, 0xc1, 0x91, 0xa1, 0xfd, 0x3c, 0x32, 0xb4, 0xaf,
	0xc7, 0x46, 0xe3, 0xe0, 0xd8, 0x68, 0x7c, 0x3f, 0x36, 0x1a, 0xef, 0xad, 0x20, 0x14, 0xdb, 0xd9,
	0xa6, 0xe5, 0xd1, 0xd8, 0x2e, 0x06, 0xb8, 0x9f, 0x80, 0xf8, 0x44, 0xd9, 0x8e, 0xfc, 0x63, 0xe7,
	0x8f, 0xed, 0xdd, 0xea, 0x26, 0x8a, 0xbd, 0x14, 0xf8, 0x66, 0x53, 0x1e, 0xc4, 0x07, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x72, 0xa7, 0x58, 0xb8, 0x73, 0x05, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.AvgCounterParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	if len(m.MedianDeviations) > 0 {
		for iNdEx := len(m.MedianDeviations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MedianDeviations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.HistoricPrices) > 0 {
		for iNdEx := len(m.HistoricPrices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.HistoricPrices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.Medians) > 0 {
		for iNdEx := len(m.Medians) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Medians[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.AggregateExchangeRateVotes) > 0 {
		for iNdEx := len(m.AggregateExchangeRateVotes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AggregateExchangeRateVotes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.AggregateExchangeRatePrevotes) > 0 {
		for iNdEx := len(m.AggregateExchangeRatePrevotes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AggregateExchangeRatePrevotes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.MissCounters) > 0 {
		for iNdEx := len(m.MissCounters) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MissCounters[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ExchangeRates) > 0 {
		for iNdEx := len(m.ExchangeRates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ExchangeRates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.FeederDelegations) > 0 {
		for iNdEx := len(m.FeederDelegations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeederDelegations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *FeederDelegation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeederDelegation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeederDelegation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FeederAddress) > 0 {
		i -= len(m.FeederAddress)
		copy(dAtA[i:], m.FeederAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.FeederAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MissCounter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MissCounter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MissCounter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MissCounter != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MissCounter))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Price) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Price) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Price) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockNum != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.BlockNum))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.ExchangeRateTuple.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.FeederDelegations) > 0 {
		for _, e := range m.FeederDelegations {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ExchangeRates) > 0 {
		for _, e := range m.ExchangeRates {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.MissCounters) > 0 {
		for _, e := range m.MissCounters {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.AggregateExchangeRatePrevotes) > 0 {
		for _, e := range m.AggregateExchangeRatePrevotes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.AggregateExchangeRateVotes) > 0 {
		for _, e := range m.AggregateExchangeRateVotes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Medians) > 0 {
		for _, e := range m.Medians {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.HistoricPrices) > 0 {
		for _, e := range m.HistoricPrices {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.MedianDeviations) > 0 {
		for _, e := range m.MedianDeviations {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.AvgCounterParams.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *FeederDelegation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FeederAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *MissCounter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.MissCounter != 0 {
		n += 1 + sovGenesis(uint64(m.MissCounter))
	}
	return n
}

func (m *Price) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ExchangeRateTuple.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.BlockNum != 0 {
		n += 1 + sovGenesis(uint64(m.BlockNum))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeederDelegations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeederDelegations = append(m.FeederDelegations, FeederDelegation{})
			if err := m.FeederDelegations[len(m.FeederDelegations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExchangeRates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExchangeRates = append(m.ExchangeRates, ExchangeRateTuple{})
			if err := m.ExchangeRates[len(m.ExchangeRates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissCounters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MissCounters = append(m.MissCounters, MissCounter{})
			if err := m.MissCounters[len(m.MissCounters)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregateExchangeRatePrevotes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AggregateExchangeRatePrevotes = append(m.AggregateExchangeRatePrevotes, AggregateExchangeRatePrevote{})
			if err := m.AggregateExchangeRatePrevotes[len(m.AggregateExchangeRatePrevotes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregateExchangeRateVotes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AggregateExchangeRateVotes = append(m.AggregateExchangeRateVotes, AggregateExchangeRateVote{})
			if err := m.AggregateExchangeRateVotes[len(m.AggregateExchangeRateVotes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Medians", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Medians = append(m.Medians, Price{})
			if err := m.Medians[len(m.Medians)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HistoricPrices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HistoricPrices = append(m.HistoricPrices, Price{})
			if err := m.HistoricPrices[len(m.HistoricPrices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MedianDeviations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MedianDeviations = append(m.MedianDeviations, Price{})
			if err := m.MedianDeviations[len(m.MedianDeviations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AvgCounterParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AvgCounterParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FeederDelegation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FeederDelegation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeederDelegation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeederAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeederAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MissCounter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MissCounter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MissCounter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissCounter", wireType)
			}
			m.MissCounter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MissCounter |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Price) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Price: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Price: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExchangeRateTuple", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ExchangeRateTuple.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockNum", wireType)
			}
			m.BlockNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockNum |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
