// Code generated by protoc-gen-gogo.
// source: pb/geo/geo.proto
// DO NOT EDIT!

package pbgeo

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import pb_money "go.pedge.io/pb/gogo/pb/money"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type StreetDirection int32

const (
	StreetDirection_STREET_DIRECTION_NONE StreetDirection = 0
	StreetDirection_STREET_DIRECTION_N    StreetDirection = 1
	StreetDirection_STREET_DIRECTION_S    StreetDirection = 2
	StreetDirection_STREET_DIRECTION_E    StreetDirection = 3
	StreetDirection_STREET_DIRECTION_W    StreetDirection = 4
	StreetDirection_STREET_DIRECTION_NE   StreetDirection = 5
	StreetDirection_STREET_DIRECTION_SE   StreetDirection = 6
	StreetDirection_STREET_DIRECTION_NW   StreetDirection = 7
	StreetDirection_STREET_DIRECTION_SW   StreetDirection = 8
)

var StreetDirection_name = map[int32]string{
	0: "STREET_DIRECTION_NONE",
	1: "STREET_DIRECTION_N",
	2: "STREET_DIRECTION_S",
	3: "STREET_DIRECTION_E",
	4: "STREET_DIRECTION_W",
	5: "STREET_DIRECTION_NE",
	6: "STREET_DIRECTION_SE",
	7: "STREET_DIRECTION_NW",
	8: "STREET_DIRECTION_SW",
}
var StreetDirection_value = map[string]int32{
	"STREET_DIRECTION_NONE": 0,
	"STREET_DIRECTION_N":    1,
	"STREET_DIRECTION_S":    2,
	"STREET_DIRECTION_E":    3,
	"STREET_DIRECTION_W":    4,
	"STREET_DIRECTION_NE":   5,
	"STREET_DIRECTION_SE":   6,
	"STREET_DIRECTION_NW":   7,
	"STREET_DIRECTION_SW":   8,
}

func (x StreetDirection) String() string {
	return proto.EnumName(StreetDirection_name, int32(x))
}
func (StreetDirection) EnumDescriptor() ([]byte, []int) { return fileDescriptorGeo, []int{0} }

type Country struct {
	Name         string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Alpha_2Code  CountryAlpha2Code     `protobuf:"varint,2,opt,name=alpha_2_code,json=alpha2Code,proto3,enum=pb.geo.CountryAlpha2Code" json:"alpha_2_code,omitempty"`
	Alpha_3Code  CountryAlpha3Code     `protobuf:"varint,3,opt,name=alpha_3_code,json=alpha3Code,proto3,enum=pb.geo.CountryAlpha3Code" json:"alpha_3_code,omitempty"`
	NumericCode  uint32                `protobuf:"varint,4,opt,name=numeric_code,json=numericCode,proto3" json:"numeric_code,omitempty"`
	CurrencyCode pb_money.CurrencyCode `protobuf:"varint,5,opt,name=currency_code,json=currencyCode,proto3,enum=pb.money.CurrencyCode" json:"currency_code,omitempty"`
}

func (m *Country) Reset()                    { *m = Country{} }
func (m *Country) String() string            { return proto.CompactTextString(m) }
func (*Country) ProtoMessage()               {}
func (*Country) Descriptor() ([]byte, []int) { return fileDescriptorGeo, []int{0} }

type LatLng struct {
	LatPicos  int64 `protobuf:"varint,1,opt,name=lat_picos,json=latPicos,proto3" json:"lat_picos,omitempty"`
	LongPicos int64 `protobuf:"varint,2,opt,name=long_picos,json=longPicos,proto3" json:"long_picos,omitempty"`
}

func (m *LatLng) Reset()                    { *m = LatLng{} }
func (m *LatLng) String() string            { return proto.CompactTextString(m) }
func (*LatLng) ProtoMessage()               {}
func (*LatLng) Descriptor() ([]byte, []int) { return fileDescriptorGeo, []int{1} }

type PostalAddress struct {
	StreetNumber                     uint64            `protobuf:"varint,1,opt,name=street_number,json=streetNumber,proto3" json:"street_number,omitempty"`
	StreetNumberPostfix              string            `protobuf:"bytes,2,opt,name=street_number_postfix,json=streetNumberPostfix,proto3" json:"street_number_postfix,omitempty"`
	StreetName                       string            `protobuf:"bytes,3,opt,name=street_name,json=streetName,proto3" json:"street_name,omitempty"`
	PreStreetDirection               StreetDirection   `protobuf:"varint,4,opt,name=pre_street_direction,json=preStreetDirection,proto3,enum=pb.geo.StreetDirection" json:"pre_street_direction,omitempty"`
	PostStreetDirection              StreetDirection   `protobuf:"varint,5,opt,name=post_street_direction,json=postStreetDirection,proto3,enum=pb.geo.StreetDirection" json:"post_street_direction,omitempty"`
	StreetTypeAbbreviation           string            `protobuf:"bytes,6,opt,name=street_type_abbreviation,json=streetTypeAbbreviation,proto3" json:"street_type_abbreviation,omitempty"`
	SecondaryAddressTypeAbbreviation string            `protobuf:"bytes,7,opt,name=secondary_address_type_abbreviation,json=secondaryAddressTypeAbbreviation,proto3" json:"secondary_address_type_abbreviation,omitempty"`
	SecondaryAddressValue            string            `protobuf:"bytes,8,opt,name=secondary_address_value,json=secondaryAddressValue,proto3" json:"secondary_address_value,omitempty"`
	LocalityName                     string            `protobuf:"bytes,9,opt,name=locality_name,json=localityName,proto3" json:"locality_name,omitempty"`
	RegionCode                       string            `protobuf:"bytes,10,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	PostalCode                       string            `protobuf:"bytes,11,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	CountryAlpha_2Code               CountryAlpha2Code `protobuf:"varint,12,opt,name=country_alpha_2_code,json=countryAlpha2Code,proto3,enum=pb.geo.CountryAlpha2Code" json:"country_alpha_2_code,omitempty"`
}

func (m *PostalAddress) Reset()                    { *m = PostalAddress{} }
func (m *PostalAddress) String() string            { return proto.CompactTextString(m) }
func (*PostalAddress) ProtoMessage()               {}
func (*PostalAddress) Descriptor() ([]byte, []int) { return fileDescriptorGeo, []int{2} }

type SimplePostalAddress struct {
	StreetAddress      string            `protobuf:"bytes,1,opt,name=street_address,json=streetAddress,proto3" json:"street_address,omitempty"`
	StreetAddress_2    string            `protobuf:"bytes,2,opt,name=street_address_2,json=streetAddress2,proto3" json:"street_address_2,omitempty"`
	LocalityName       string            `protobuf:"bytes,3,opt,name=locality_name,json=localityName,proto3" json:"locality_name,omitempty"`
	RegionCode         string            `protobuf:"bytes,4,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	PostalCode         string            `protobuf:"bytes,5,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	CountryAlpha_2Code CountryAlpha2Code `protobuf:"varint,6,opt,name=country_alpha_2_code,json=countryAlpha2Code,proto3,enum=pb.geo.CountryAlpha2Code" json:"country_alpha_2_code,omitempty"`
}

func (m *SimplePostalAddress) Reset()                    { *m = SimplePostalAddress{} }
func (m *SimplePostalAddress) String() string            { return proto.CompactTextString(m) }
func (*SimplePostalAddress) ProtoMessage()               {}
func (*SimplePostalAddress) Descriptor() ([]byte, []int) { return fileDescriptorGeo, []int{3} }

func init() {
	proto.RegisterType((*Country)(nil), "pb.geo.Country")
	proto.RegisterType((*LatLng)(nil), "pb.geo.LatLng")
	proto.RegisterType((*PostalAddress)(nil), "pb.geo.PostalAddress")
	proto.RegisterType((*SimplePostalAddress)(nil), "pb.geo.SimplePostalAddress")
	proto.RegisterEnum("pb.geo.StreetDirection", StreetDirection_name, StreetDirection_value)
}

var fileDescriptorGeo = []byte{
	// 648 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xcf, 0x6f, 0xd3, 0x3e,
	0x14, 0xff, 0xf6, 0x57, 0xba, 0xbe, 0xa5, 0xfd, 0x06, 0x6f, 0xdd, 0xba, 0x21, 0xc4, 0xe8, 0x84,
	0x34, 0x71, 0x28, 0x52, 0x26, 0x21, 0x24, 0x4e, 0xa3, 0xcb, 0x61, 0x30, 0xba, 0x29, 0xad, 0x98,
	0xc4, 0xc5, 0x72, 0x52, 0x53, 0x22, 0xa5, 0x71, 0xe4, 0xa4, 0x13, 0xf9, 0x43, 0x38, 0xf0, 0x5f,
	0x72, 0xe3, 0x8a, 0x63, 0xbb, 0x5b, 0xdb, 0x04, 0x06, 0x87, 0x4d, 0xf6, 0xe7, 0x97, 0xe3, 0xf7,
	0x9e, 0x0b, 0x56, 0xec, 0xbd, 0x9c, 0x51, 0x96, 0xff, 0x0d, 0x62, 0xce, 0x52, 0x86, 0x8c, 0xd8,
	0x1b, 0x88, 0xdd, 0xe1, 0xee, 0x0a, 0x33, 0xa3, 0x91, 0x62, 0x0f, 0x7b, 0x02, 0x9d, 0xb3, 0x88,
	0x66, 0xea, 0xff, 0x3d, 0xd3, 0xff, 0x59, 0x81, 0xe6, 0x90, 0x2d, 0xa2, 0x94, 0x67, 0x08, 0x41,
	0x3d, 0x22, 0x73, 0xda, 0xab, 0x1c, 0x55, 0x4e, 0x5a, 0xae, 0x5c, 0xa3, 0x37, 0x60, 0x92, 0x30,
	0xfe, 0x42, 0xb0, 0x8d, 0x7d, 0x36, 0xa5, 0xbd, 0xaa, 0xe0, 0x3a, 0xf6, 0xc1, 0x40, 0x1d, 0x37,
	0xd0, 0xd6, 0xb3, 0x5c, 0x62, 0x0f, 0x85, 0xc0, 0x05, 0x72, 0xb7, 0xbe, 0x37, 0x9f, 0x2a, 0x73,
	0xed, 0xf7, 0xe6, 0xd3, 0x15, 0xb3, 0x5c, 0xa3, 0x67, 0x60, 0x46, 0x8b, 0x39, 0xe5, 0x81, 0xaf,
	0xcc, 0x75, 0x61, 0x6e, 0xbb, 0xdb, 0x1a, 0xd3, 0xf9, 0x6d, 0x7f, 0xc1, 0x39, 0x8d, 0xfc, 0x4c,
	0x69, 0x1a, 0xf2, 0x80, 0xbd, 0xfc, 0x00, 0x75, 0xd1, 0xa1, 0xa6, 0x65, 0xba, 0xe9, 0xaf, 0xec,
	0xfa, 0xe7, 0x60, 0x5c, 0x92, 0xf4, 0x32, 0x9a, 0xa1, 0xc7, 0xd0, 0x0a, 0x49, 0x8a, 0xe3, 0xc0,
	0x67, 0x89, 0xbc, 0x7c, 0xcd, 0xdd, 0x12, 0xc0, 0x75, 0xbe, 0x47, 0x4f, 0x00, 0x42, 0x16, 0xcd,
	0x34, 0x5b, 0x95, 0x6c, 0x2b, 0x47, 0x24, 0xdd, 0xff, 0xd6, 0x80, 0xf6, 0x35, 0x4b, 0x52, 0x12,
	0x9e, 0x4d, 0xa7, 0x9c, 0x26, 0x09, 0x3a, 0x86, 0x76, 0x92, 0x72, 0x4a, 0x53, 0x2c, 0x3e, 0xd5,
	0xa3, 0x5c, 0x26, 0xd6, 0x5d, 0x53, 0x81, 0x23, 0x89, 0x21, 0x1b, 0xba, 0x6b, 0x22, 0x1c, 0x8b,
	0x8c, 0xcf, 0xc1, 0x57, 0x79, 0x40, 0xcb, 0xdd, 0x59, 0x15, 0x5f, 0x2b, 0x0a, 0x3d, 0x85, 0xed,
	0xa5, 0x27, 0xef, 0x52, 0x4d, 0x2a, 0x41, 0x2b, 0xf3, 0x5e, 0x5d, 0xc0, 0x6e, 0xcc, 0x29, 0xd6,
	0xa2, 0x69, 0xc0, 0xa9, 0x9f, 0x06, 0x2c, 0x92, 0x95, 0xeb, 0xd8, 0xfb, 0xcb, 0xb2, 0x8f, 0x25,
	0x7f, 0xbe, 0xa4, 0x5d, 0x24, 0x4c, 0x1b, 0x18, 0x7a, 0x0f, 0xdd, 0xfc, 0x8b, 0x8a, 0x59, 0x8d,
	0x3f, 0x67, 0xed, 0xe4, 0xae, 0xcd, 0xb0, 0xd7, 0xd0, 0xd3, 0x39, 0x69, 0x16, 0x53, 0x4c, 0x3c,
	0x8f, 0xd3, 0xdb, 0x80, 0xc8, 0x3c, 0x43, 0xde, 0x62, 0x4f, 0xf1, 0x13, 0x41, 0x9f, 0xad, 0xb0,
	0xe8, 0x03, 0x1c, 0x27, 0xd4, 0x67, 0xd1, 0x94, 0xf0, 0x0c, 0x13, 0x55, 0xe0, 0x92, 0x90, 0xa6,
	0x0c, 0x39, 0xba, 0x93, 0xea, 0x56, 0x14, 0xe2, 0x5e, 0xc1, 0x7e, 0x31, 0xee, 0x96, 0x84, 0x0b,
	0xda, 0xdb, 0x92, 0x11, 0xdd, 0xcd, 0x88, 0x8f, 0x39, 0x99, 0xb7, 0x34, 0x64, 0x3e, 0x09, 0x83,
	0x34, 0x53, 0xb5, 0x6f, 0x49, 0xb5, 0xb9, 0x04, 0x65, 0xf5, 0x45, 0x7b, 0x38, 0x9d, 0x89, 0x63,
	0xd4, 0x28, 0x82, 0x6a, 0x8f, 0x82, 0xe4, 0xb4, 0x0a, 0x41, 0x2c, 0x27, 0x45, 0x09, 0xb6, 0x95,
	0x40, 0x41, 0x52, 0xf0, 0x0e, 0x76, 0x7d, 0xf5, 0x24, 0xf0, 0xda, 0x9b, 0x33, 0x1f, 0x7a, 0x73,
	0x8f, 0xfc, 0x4d, 0xa8, 0xff, 0xbd, 0x0a, 0x3b, 0xe3, 0x60, 0x1e, 0x87, 0x74, 0x7d, 0x3a, 0x9f,
	0x43, 0x47, 0xf7, 0x42, 0xdf, 0x5f, 0xbf, 0x76, 0x3d, 0xb3, 0x4b, 0xd9, 0x09, 0x58, 0xeb, 0x32,
	0x6c, 0xeb, 0xd1, 0xec, 0xac, 0x09, 0xed, 0x62, 0x6d, 0x6a, 0x0f, 0xd7, 0xa6, 0xfe, 0x50, 0x6d,
	0x1a, 0x7f, 0x5d, 0x1b, 0xe3, 0xdf, 0x6b, 0xf3, 0xe2, 0x47, 0x05, 0xfe, 0xdf, 0x9c, 0xd1, 0x03,
	0xe8, 0x8e, 0x27, 0xae, 0xe3, 0x4c, 0xf0, 0xf9, 0x85, 0xeb, 0x0c, 0x27, 0x17, 0x57, 0x23, 0x3c,
	0xba, 0x1a, 0x39, 0xd6, 0x7f, 0x68, 0x0f, 0x50, 0x91, 0xb2, 0x2a, 0xa5, 0xf8, 0xd8, 0xaa, 0x96,
	0xe2, 0x8e, 0x55, 0x2b, 0xc5, 0x6f, 0xac, 0x3a, 0xda, 0x17, 0x9d, 0x2a, 0xe4, 0x3b, 0x56, 0xa3,
	0x94, 0x18, 0x3b, 0x96, 0x51, 0xee, 0xb8, 0xb1, 0x9a, 0xe5, 0x8e, 0x1b, 0x6b, 0xeb, 0x6d, 0xf3,
	0x53, 0x23, 0xf6, 0x44, 0x81, 0x3c, 0x43, 0xfe, 0xec, 0x9f, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff,
	0xce, 0xc5, 0x0c, 0x01, 0x42, 0x06, 0x00, 0x00,
}