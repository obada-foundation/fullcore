// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: obit/v1/nft.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type NFT struct {
	ClassId string     `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	Id      string     `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Uri     string     `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	UriHash string     `protobuf:"bytes,4,opt,name=uri_hash,json=uriHash,proto3" json:"uri_hash,omitempty"`
	Data    *types.Any `protobuf:"bytes,10,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *NFT) Reset()         { *m = NFT{} }
func (m *NFT) String() string { return proto.CompactTextString(m) }
func (*NFT) ProtoMessage()    {}
func (*NFT) Descriptor() ([]byte, []int) {
	return fileDescriptor_d18d26ebeaa370a5, []int{0}
}
func (m *NFT) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NFT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NFT.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NFT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NFT.Merge(m, src)
}
func (m *NFT) XXX_Size() int {
	return m.Size()
}
func (m *NFT) XXX_DiscardUnknown() {
	xxx_messageInfo_NFT.DiscardUnknown(m)
}

var xxx_messageInfo_NFT proto.InternalMessageInfo

func (m *NFT) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

func (m *NFT) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NFT) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *NFT) GetUriHash() string {
	if m != nil {
		return m.UriHash
	}
	return ""
}

func (m *NFT) GetData() *types.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

type NFTData struct {
	Usn string `protobuf:"bytes,1,opt,name=usn,proto3" json:"usn,omitempty"`
}

func (m *NFTData) Reset()         { *m = NFTData{} }
func (m *NFTData) String() string { return proto.CompactTextString(m) }
func (*NFTData) ProtoMessage()    {}
func (*NFTData) Descriptor() ([]byte, []int) {
	return fileDescriptor_d18d26ebeaa370a5, []int{1}
}
func (m *NFTData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NFTData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NFTData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NFTData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NFTData.Merge(m, src)
}
func (m *NFTData) XXX_Size() int {
	return m.Size()
}
func (m *NFTData) XXX_DiscardUnknown() {
	xxx_messageInfo_NFTData.DiscardUnknown(m)
}

var xxx_messageInfo_NFTData proto.InternalMessageInfo

func (m *NFTData) GetUsn() string {
	if m != nil {
		return m.Usn
	}
	return ""
}

func init() {
	proto.RegisterType((*NFT)(nil), "obadafoundation.fullcore.obit.NFT")
	proto.RegisterType((*NFTData)(nil), "obadafoundation.fullcore.obit.NFTData")
}

func init() { proto.RegisterFile("obit/v1/nft.proto", fileDescriptor_d18d26ebeaa370a5) }

var fileDescriptor_d18d26ebeaa370a5 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x18, 0xc4, 0xbb, 0x49, 0xb1, 0xba, 0x82, 0x68, 0xe8, 0x21, 0xad, 0xb8, 0x94, 0x9e, 0x72, 0x71,
	0x97, 0xea, 0x13, 0x28, 0x52, 0x14, 0xa1, 0x87, 0xd2, 0x93, 0x97, 0xb2, 0xc9, 0xe6, 0xcf, 0x42,
	0xdc, 0xaf, 0x64, 0x77, 0xc5, 0x3c, 0x81, 0x57, 0x1f, 0xcb, 0x63, 0x8f, 0x1e, 0x25, 0x79, 0x11,
	0xc9, 0xc6, 0xe2, 0x6d, 0xbe, 0x61, 0x86, 0x6f, 0xf8, 0xe1, 0x0b, 0x88, 0xa5, 0x61, 0x6f, 0x0b,
	0xa6, 0x32, 0x43, 0x77, 0x15, 0x18, 0x08, 0xae, 0x20, 0xe6, 0x82, 0x67, 0x60, 0x95, 0xe0, 0x46,
	0x82, 0xa2, 0x99, 0x2d, 0xcb, 0x04, 0xaa, 0x94, 0x76, 0xd9, 0xe9, 0x24, 0x07, 0xc8, 0xcb, 0x94,
	0xb9, 0x70, 0x6c, 0x33, 0xc6, 0x55, 0xdd, 0x37, 0xa7, 0xe3, 0x1c, 0x72, 0x70, 0x92, 0x75, 0xaa,
	0x77, 0xe7, 0x1f, 0x08, 0xfb, 0xab, 0xe5, 0x26, 0x98, 0xe0, 0xe3, 0xa4, 0xe4, 0x5a, 0x6f, 0xa5,
	0x08, 0xd1, 0x0c, 0x45, 0x27, 0xeb, 0x91, 0xbb, 0x9f, 0x44, 0x70, 0x86, 0x3d, 0x29, 0x42, 0xcf,
	0x99, 0x9e, 0x14, 0xc1, 0x39, 0xf6, 0x6d, 0x25, 0x43, 0xdf, 0x19, 0x9d, 0xec, 0xca, 0xb6, 0x92,
	0xdb, 0x82, 0xeb, 0x22, 0x1c, 0xf6, 0x65, 0x5b, 0xc9, 0x47, 0xae, 0x8b, 0x20, 0xc2, 0x43, 0xc1,
	0x0d, 0x0f, 0xf1, 0x0c, 0x45, 0xa7, 0x37, 0x63, 0xda, 0xef, 0xa3, 0x87, 0x7d, 0xf4, 0x4e, 0xd5,
	0x6b, 0x97, 0x98, 0x5f, 0xe2, 0xd1, 0x6a, 0xb9, 0x79, 0xe0, 0x86, 0xbb, 0x0f, 0x5a, 0xfd, 0xed,
	0xe8, 0xe4, 0xfd, 0xf3, 0x57, 0x43, 0xd0, 0xbe, 0x21, 0xe8, 0xa7, 0x21, 0xe8, 0xb3, 0x25, 0x83,
	0x7d, 0x4b, 0x06, 0xdf, 0x2d, 0x19, 0xbc, 0x2c, 0x72, 0x69, 0x0a, 0x1b, 0xd3, 0x04, 0x5e, 0x99,
	0x63, 0x73, 0xfd, 0x0f, 0x87, 0x1d, 0xe0, 0xb0, 0x77, 0xe6, 0x50, 0x9a, 0x7a, 0x97, 0xea, 0xf8,
	0xc8, 0x7d, 0xbf, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x2a, 0xf1, 0x5a, 0xc6, 0x5f, 0x01, 0x00,
	0x00,
}

func (m *NFT) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NFT) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NFT) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintNft(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	if len(m.UriHash) > 0 {
		i -= len(m.UriHash)
		copy(dAtA[i:], m.UriHash)
		i = encodeVarintNft(dAtA, i, uint64(len(m.UriHash)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Uri) > 0 {
		i -= len(m.Uri)
		copy(dAtA[i:], m.Uri)
		i = encodeVarintNft(dAtA, i, uint64(len(m.Uri)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintNft(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClassId) > 0 {
		i -= len(m.ClassId)
		copy(dAtA[i:], m.ClassId)
		i = encodeVarintNft(dAtA, i, uint64(len(m.ClassId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *NFTData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NFTData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NFTData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Usn) > 0 {
		i -= len(m.Usn)
		copy(dAtA[i:], m.Usn)
		i = encodeVarintNft(dAtA, i, uint64(len(m.Usn)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNft(dAtA []byte, offset int, v uint64) int {
	offset -= sovNft(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NFT) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClassId)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	l = len(m.Uri)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	l = len(m.UriHash)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovNft(uint64(l))
	}
	return n
}

func (m *NFTData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Usn)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	return n
}

func sovNft(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNft(x uint64) (n int) {
	return sovNft(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NFT) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNft
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
			return fmt.Errorf("proto: NFT: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NFT: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UriHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UriHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &types.Any{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNft
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
func (m *NFTData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNft
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
			return fmt.Errorf("proto: NFTData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NFTData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Usn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Usn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNft
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
func skipNft(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNft
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
					return 0, ErrIntOverflowNft
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
					return 0, ErrIntOverflowNft
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
				return 0, ErrInvalidLengthNft
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNft
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNft
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNft        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNft          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNft = fmt.Errorf("proto: unexpected end of group")
)