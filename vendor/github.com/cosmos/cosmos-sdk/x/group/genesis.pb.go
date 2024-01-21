// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/group/v1/genesis.proto

package group

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the group module's genesis state.
type GenesisState struct {
	// group_seq is the group table orm.Sequence,
	// it is used to get the next group ID.
	GroupSeq uint64 `protobuf:"varint,1,opt,name=group_seq,json=groupSeq,proto3" json:"group_seq,omitempty"`
	// groups is the list of groups info.
	Groups []*GroupInfo `protobuf:"bytes,2,rep,name=groups,proto3" json:"groups,omitempty"`
	// group_members is the list of groups members.
	GroupMembers []*GroupMember `protobuf:"bytes,3,rep,name=group_members,json=groupMembers,proto3" json:"group_members,omitempty"`
	// group_policy_seq is the group policy table orm.Sequence,
	// it is used to generate the next group policy account address.
	GroupPolicySeq uint64 `protobuf:"varint,4,opt,name=group_policy_seq,json=groupPolicySeq,proto3" json:"group_policy_seq,omitempty"`
	// group_policies is the list of group policies info.
	GroupPolicies []*GroupPolicyInfo `protobuf:"bytes,5,rep,name=group_policies,json=groupPolicies,proto3" json:"group_policies,omitempty"`
	// proposal_seq is the proposal table orm.Sequence,
	// it is used to get the next proposal ID.
	ProposalSeq uint64 `protobuf:"varint,6,opt,name=proposal_seq,json=proposalSeq,proto3" json:"proposal_seq,omitempty"`
	// proposals is the list of proposals.
	Proposals []*Proposal `protobuf:"bytes,7,rep,name=proposals,proto3" json:"proposals,omitempty"`
	// votes is the list of votes.
	Votes []*Vote `protobuf:"bytes,8,rep,name=votes,proto3" json:"votes,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc6105fe3ef99f06, []int{0}
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

func (m *GenesisState) GetGroupSeq() uint64 {
	if m != nil {
		return m.GroupSeq
	}
	return 0
}

func (m *GenesisState) GetGroups() []*GroupInfo {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *GenesisState) GetGroupMembers() []*GroupMember {
	if m != nil {
		return m.GroupMembers
	}
	return nil
}

func (m *GenesisState) GetGroupPolicySeq() uint64 {
	if m != nil {
		return m.GroupPolicySeq
	}
	return 0
}

func (m *GenesisState) GetGroupPolicies() []*GroupPolicyInfo {
	if m != nil {
		return m.GroupPolicies
	}
	return nil
}

func (m *GenesisState) GetProposalSeq() uint64 {
	if m != nil {
		return m.ProposalSeq
	}
	return 0
}

func (m *GenesisState) GetProposals() []*Proposal {
	if m != nil {
		return m.Proposals
	}
	return nil
}

func (m *GenesisState) GetVotes() []*Vote {
	if m != nil {
		return m.Votes
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "cosmos.group.v1.GenesisState")
}

func init() { proto.RegisterFile("cosmos/group/v1/genesis.proto", fileDescriptor_cc6105fe3ef99f06) }

var fileDescriptor_cc6105fe3ef99f06 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x4e, 0xfa, 0x40,
	0x10, 0xc7, 0xe9, 0x8f, 0x3f, 0x3f, 0x58, 0xfe, 0x68, 0x36, 0x31, 0xa9, 0xa0, 0x0d, 0x1a, 0x0f,
	0x24, 0xc6, 0x36, 0xe0, 0xc1, 0x9b, 0x89, 0x5e, 0x88, 0x07, 0x13, 0x52, 0x12, 0x0f, 0x5e, 0x0c,
	0xe0, 0x58, 0x1b, 0x29, 0x53, 0x3a, 0x0b, 0x91, 0xb7, 0xf0, 0x09, 0x7c, 0x1e, 0x8f, 0x1c, 0x3d,
	0x1a, 0x78, 0x11, 0xc3, 0x6c, 0x49, 0x0d, 0x70, 0xda, 0xdd, 0xd9, 0xcf, 0x77, 0x3e, 0x93, 0x8c,
	0x38, 0x1e, 0x20, 0x05, 0x48, 0x8e, 0x17, 0xe1, 0x24, 0x74, 0xa6, 0x4d, 0xc7, 0x83, 0x11, 0x90,
	0x4f, 0x76, 0x18, 0xa1, 0x42, 0xb9, 0xa7, 0xbf, 0x6d, 0xfe, 0xb6, 0xa7, 0xcd, 0x6a, 0x6d, 0x93,
	0x57, 0xb3, 0x10, 0x62, 0xfa, 0xf4, 0x33, 0x2d, 0x4a, 0x6d, 0x9d, 0xef, 0xaa, 0x9e, 0x02, 0x59,
	0x13, 0x05, 0x06, 0x9f, 0x08, 0xc6, 0xa6, 0x51, 0x37, 0x1a, 0x19, 0x37, 0xcf, 0x85, 0x2e, 0x8c,
	0x65, 0x4b, 0xe4, 0xf8, 0x4e, 0xe6, 0xbf, 0x7a, 0xba, 0x51, 0x6c, 0x55, 0xed, 0x0d, 0x99, 0xdd,
	0x5e, 0x5d, 0xee, 0x46, 0x2f, 0xe8, 0xc6, 0xa4, 0xbc, 0x11, 0x65, 0xdd, 0x30, 0x80, 0xa0, 0x0f,
	0x11, 0x99, 0x69, 0x8e, 0x1e, 0xed, 0x8e, 0xde, 0x33, 0xe4, 0x96, 0xbc, 0xe4, 0x41, 0xb2, 0x21,
	0xf6, 0x75, 0x8b, 0x10, 0x87, 0xfe, 0x60, 0xc6, 0xa3, 0x65, 0x78, 0xb4, 0x0a, 0xd7, 0x3b, 0x5c,
	0x5e, 0x0d, 0xd8, 0x16, 0x95, 0x3f, 0xa4, 0x0f, 0x64, 0x66, 0xd9, 0x56, 0xdf, 0x6d, 0xd3, 0x41,
	0x1e, 0xb7, 0x9c, 0x74, 0xf2, 0x81, 0xe4, 0x89, 0x28, 0x85, 0x11, 0x86, 0x48, 0xbd, 0x21, 0xeb,
	0x72, 0xac, 0x2b, 0xae, 0x6b, 0x2b, 0xd7, 0x95, 0x28, 0xac, 0x9f, 0x64, 0xfe, 0x67, 0xcd, 0xe1,
	0x96, 0xa6, 0x13, 0x13, 0x6e, 0xc2, 0xca, 0x73, 0x91, 0x9d, 0xa2, 0x02, 0x32, 0xf3, 0x1c, 0x3a,
	0xd8, 0x0a, 0x3d, 0xa0, 0x02, 0x57, 0x33, 0xb7, 0xd7, 0x5f, 0x0b, 0xcb, 0x98, 0x2f, 0x2c, 0xe3,
	0x67, 0x61, 0x19, 0x1f, 0x4b, 0x2b, 0x35, 0x5f, 0x5a, 0xa9, 0xef, 0xa5, 0x95, 0x7a, 0x3c, 0xf3,
	0x7c, 0xf5, 0x3a, 0xe9, 0xdb, 0x03, 0x0c, 0x9c, 0x78, 0xc5, 0xfa, 0xb8, 0xa0, 0xe7, 0x37, 0xe7,
	0x5d, 0xef, 0xbb, 0x9f, 0xe3, 0x3d, 0x5f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x99, 0x5b, 0x30,
	0xc4, 0x36, 0x02, 0x00, 0x00,
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
	if len(m.Votes) > 0 {
		for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Votes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Proposals) > 0 {
		for iNdEx := len(m.Proposals) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Proposals[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if m.ProposalSeq != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.ProposalSeq))
		i--
		dAtA[i] = 0x30
	}
	if len(m.GroupPolicies) > 0 {
		for iNdEx := len(m.GroupPolicies) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GroupPolicies[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if m.GroupPolicySeq != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.GroupPolicySeq))
		i--
		dAtA[i] = 0x20
	}
	if len(m.GroupMembers) > 0 {
		for iNdEx := len(m.GroupMembers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GroupMembers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Groups) > 0 {
		for iNdEx := len(m.Groups) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Groups[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if m.GroupSeq != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.GroupSeq))
		i--
		dAtA[i] = 0x8
	}
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
	if m.GroupSeq != 0 {
		n += 1 + sovGenesis(uint64(m.GroupSeq))
	}
	if len(m.Groups) > 0 {
		for _, e := range m.Groups {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.GroupMembers) > 0 {
		for _, e := range m.GroupMembers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.GroupPolicySeq != 0 {
		n += 1 + sovGenesis(uint64(m.GroupPolicySeq))
	}
	if len(m.GroupPolicies) > 0 {
		for _, e := range m.GroupPolicies {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.ProposalSeq != 0 {
		n += 1 + sovGenesis(uint64(m.ProposalSeq))
	}
	if len(m.Proposals) > 0 {
		for _, e := range m.Proposals {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Votes) > 0 {
		for _, e := range m.Votes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupSeq", wireType)
			}
			m.GroupSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GroupSeq |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Groups", wireType)
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
			m.Groups = append(m.Groups, &GroupInfo{})
			if err := m.Groups[len(m.Groups)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupMembers", wireType)
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
			m.GroupMembers = append(m.GroupMembers, &GroupMember{})
			if err := m.GroupMembers[len(m.GroupMembers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupPolicySeq", wireType)
			}
			m.GroupPolicySeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GroupPolicySeq |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupPolicies", wireType)
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
			m.GroupPolicies = append(m.GroupPolicies, &GroupPolicyInfo{})
			if err := m.GroupPolicies[len(m.GroupPolicies)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalSeq", wireType)
			}
			m.ProposalSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalSeq |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposals", wireType)
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
			m.Proposals = append(m.Proposals, &Proposal{})
			if err := m.Proposals[len(m.Proposals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
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
			m.Votes = append(m.Votes, &Vote{})
			if err := m.Votes[len(m.Votes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
