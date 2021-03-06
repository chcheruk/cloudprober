// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/probes/http/proto/config.proto

package proto

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

type ProbeConf_ProtocolType int32

const (
	ProbeConf_HTTP  ProbeConf_ProtocolType = 0
	ProbeConf_HTTPS ProbeConf_ProtocolType = 1
)

var ProbeConf_ProtocolType_name = map[int32]string{
	0: "HTTP",
	1: "HTTPS",
}
var ProbeConf_ProtocolType_value = map[string]int32{
	"HTTP":  0,
	"HTTPS": 1,
}

func (x ProbeConf_ProtocolType) Enum() *ProbeConf_ProtocolType {
	p := new(ProbeConf_ProtocolType)
	*p = x
	return p
}
func (x ProbeConf_ProtocolType) String() string {
	return proto.EnumName(ProbeConf_ProtocolType_name, int32(x))
}
func (x *ProbeConf_ProtocolType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ProbeConf_ProtocolType_value, data, "ProbeConf_ProtocolType")
	if err != nil {
		return err
	}
	*x = ProbeConf_ProtocolType(value)
	return nil
}
func (ProbeConf_ProtocolType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_config_d95e376e63169085, []int{0, 0}
}

type ProbeConf_Method int32

const (
	ProbeConf_GET     ProbeConf_Method = 0
	ProbeConf_POST    ProbeConf_Method = 1
	ProbeConf_PUT     ProbeConf_Method = 2
	ProbeConf_HEAD    ProbeConf_Method = 3
	ProbeConf_DELETE  ProbeConf_Method = 4
	ProbeConf_PATCH   ProbeConf_Method = 5
	ProbeConf_OPTIONS ProbeConf_Method = 6
)

var ProbeConf_Method_name = map[int32]string{
	0: "GET",
	1: "POST",
	2: "PUT",
	3: "HEAD",
	4: "DELETE",
	5: "PATCH",
	6: "OPTIONS",
}
var ProbeConf_Method_value = map[string]int32{
	"GET":     0,
	"POST":    1,
	"PUT":     2,
	"HEAD":    3,
	"DELETE":  4,
	"PATCH":   5,
	"OPTIONS": 6,
}

func (x ProbeConf_Method) Enum() *ProbeConf_Method {
	p := new(ProbeConf_Method)
	*p = x
	return p
}
func (x ProbeConf_Method) String() string {
	return proto.EnumName(ProbeConf_Method_name, int32(x))
}
func (x *ProbeConf_Method) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ProbeConf_Method_value, data, "ProbeConf_Method")
	if err != nil {
		return err
	}
	*x = ProbeConf_Method(value)
	return nil
}
func (ProbeConf_Method) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_config_d95e376e63169085, []int{0, 1}
}

type ProbeConf struct {
	// Which HTTP protocol to use
	Protocol *ProbeConf_ProtocolType `protobuf:"varint,1,opt,name=protocol,enum=cloudprober.probes.http.ProbeConf_ProtocolType,def=0" json:"protocol,omitempty"`
	// Relative URL (to append to all targets). Must begin with '/'
	RelativeUrl *string `protobuf:"bytes,2,opt,name=relative_url,json=relativeUrl" json:"relative_url,omitempty"`
	// Port, default is 80 for HTTP and 443 for HTTPS
	Port *int32 `protobuf:"varint,3,opt,name=port" json:"port,omitempty"`
	// Set the source address to send packets from, either by providing an address
	// or a network interface. Note: the source can also be configured dynamically
	// using an environment variable.
	//
	// Types that are valid to be assigned to Source:
	//	*ProbeConf_SourceIp
	//	*ProbeConf_SourceInterface
	Source isProbeConf_Source `protobuf_oneof:"source"`
	// Whether to resolve the target before making the request. If set to false,
	// we hand over the target and relative_url directly to the golang's HTTP
	// module, Otherwise, we resolve the target first to an IP address and
	// make a request using that while passing target name as Host header.
	ResolveFirst *bool `protobuf:"varint,4,opt,name=resolve_first,json=resolveFirst,def=0" json:"resolve_first,omitempty"`
	// Export response (body) count as a metric
	ExportResponseAsMetrics *bool `protobuf:"varint,5,opt,name=export_response_as_metrics,json=exportResponseAsMetrics,def=0" json:"export_response_as_metrics,omitempty"`
	// This field is now deprecated and doesn't do anything (except resulting in
	// a warning). It will be removed in the next release. For data integrity
	// checks, please use integrity validator.
	IntegrityCheckPattern *string `protobuf:"bytes,6,opt,name=integrity_check_pattern,json=integrityCheckPattern" json:"integrity_check_pattern,omitempty"`
	// HTTP request method
	Method *ProbeConf_Method `protobuf:"varint,7,opt,name=method,enum=cloudprober.probes.http.ProbeConf_Method,def=0" json:"method,omitempty"`
	// HTTP request headers
	Headers []*ProbeConf_Header `protobuf:"bytes,8,rep,name=headers" json:"headers,omitempty"`
	// Request body.
	Body *string `protobuf:"bytes,9,opt,name=body" json:"body,omitempty"`
	// Requests per probe
	RequestsPerProbe *int32 `protobuf:"varint,98,opt,name=requests_per_probe,json=requestsPerProbe,def=1" json:"requests_per_probe,omitempty"`
	// How long to wait between two requests to the same target
	RequestsIntervalMsec *int32 `protobuf:"varint,99,opt,name=requests_interval_msec,json=requestsIntervalMsec,def=25" json:"requests_interval_msec,omitempty"`
	// Export stats after these many milliseconds
	StatsExportIntervalMsec *int32   `protobuf:"varint,100,opt,name=stats_export_interval_msec,json=statsExportIntervalMsec,def=10000" json:"stats_export_interval_msec,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *ProbeConf) Reset()         { *m = ProbeConf{} }
func (m *ProbeConf) String() string { return proto.CompactTextString(m) }
func (*ProbeConf) ProtoMessage()    {}
func (*ProbeConf) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_d95e376e63169085, []int{0}
}
func (m *ProbeConf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProbeConf.Unmarshal(m, b)
}
func (m *ProbeConf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProbeConf.Marshal(b, m, deterministic)
}
func (dst *ProbeConf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProbeConf.Merge(dst, src)
}
func (m *ProbeConf) XXX_Size() int {
	return xxx_messageInfo_ProbeConf.Size(m)
}
func (m *ProbeConf) XXX_DiscardUnknown() {
	xxx_messageInfo_ProbeConf.DiscardUnknown(m)
}

var xxx_messageInfo_ProbeConf proto.InternalMessageInfo

const Default_ProbeConf_Protocol ProbeConf_ProtocolType = ProbeConf_HTTP
const Default_ProbeConf_ResolveFirst bool = false
const Default_ProbeConf_ExportResponseAsMetrics bool = false
const Default_ProbeConf_Method ProbeConf_Method = ProbeConf_GET
const Default_ProbeConf_RequestsPerProbe int32 = 1
const Default_ProbeConf_RequestsIntervalMsec int32 = 25
const Default_ProbeConf_StatsExportIntervalMsec int32 = 10000

func (m *ProbeConf) GetProtocol() ProbeConf_ProtocolType {
	if m != nil && m.Protocol != nil {
		return *m.Protocol
	}
	return Default_ProbeConf_Protocol
}

func (m *ProbeConf) GetRelativeUrl() string {
	if m != nil && m.RelativeUrl != nil {
		return *m.RelativeUrl
	}
	return ""
}

func (m *ProbeConf) GetPort() int32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

type isProbeConf_Source interface {
	isProbeConf_Source()
}

type ProbeConf_SourceIp struct {
	SourceIp string `protobuf:"bytes,10,opt,name=source_ip,json=sourceIp,oneof"`
}

type ProbeConf_SourceInterface struct {
	SourceInterface string `protobuf:"bytes,11,opt,name=source_interface,json=sourceInterface,oneof"`
}

func (*ProbeConf_SourceIp) isProbeConf_Source() {}

func (*ProbeConf_SourceInterface) isProbeConf_Source() {}

func (m *ProbeConf) GetSource() isProbeConf_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *ProbeConf) GetSourceIp() string {
	if x, ok := m.GetSource().(*ProbeConf_SourceIp); ok {
		return x.SourceIp
	}
	return ""
}

func (m *ProbeConf) GetSourceInterface() string {
	if x, ok := m.GetSource().(*ProbeConf_SourceInterface); ok {
		return x.SourceInterface
	}
	return ""
}

func (m *ProbeConf) GetResolveFirst() bool {
	if m != nil && m.ResolveFirst != nil {
		return *m.ResolveFirst
	}
	return Default_ProbeConf_ResolveFirst
}

func (m *ProbeConf) GetExportResponseAsMetrics() bool {
	if m != nil && m.ExportResponseAsMetrics != nil {
		return *m.ExportResponseAsMetrics
	}
	return Default_ProbeConf_ExportResponseAsMetrics
}

func (m *ProbeConf) GetIntegrityCheckPattern() string {
	if m != nil && m.IntegrityCheckPattern != nil {
		return *m.IntegrityCheckPattern
	}
	return ""
}

func (m *ProbeConf) GetMethod() ProbeConf_Method {
	if m != nil && m.Method != nil {
		return *m.Method
	}
	return Default_ProbeConf_Method
}

func (m *ProbeConf) GetHeaders() []*ProbeConf_Header {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *ProbeConf) GetBody() string {
	if m != nil && m.Body != nil {
		return *m.Body
	}
	return ""
}

func (m *ProbeConf) GetRequestsPerProbe() int32 {
	if m != nil && m.RequestsPerProbe != nil {
		return *m.RequestsPerProbe
	}
	return Default_ProbeConf_RequestsPerProbe
}

func (m *ProbeConf) GetRequestsIntervalMsec() int32 {
	if m != nil && m.RequestsIntervalMsec != nil {
		return *m.RequestsIntervalMsec
	}
	return Default_ProbeConf_RequestsIntervalMsec
}

func (m *ProbeConf) GetStatsExportIntervalMsec() int32 {
	if m != nil && m.StatsExportIntervalMsec != nil {
		return *m.StatsExportIntervalMsec
	}
	return Default_ProbeConf_StatsExportIntervalMsec
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ProbeConf) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ProbeConf_OneofMarshaler, _ProbeConf_OneofUnmarshaler, _ProbeConf_OneofSizer, []interface{}{
		(*ProbeConf_SourceIp)(nil),
		(*ProbeConf_SourceInterface)(nil),
	}
}

func _ProbeConf_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ProbeConf)
	// source
	switch x := m.Source.(type) {
	case *ProbeConf_SourceIp:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.SourceIp)
	case *ProbeConf_SourceInterface:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.SourceInterface)
	case nil:
	default:
		return fmt.Errorf("ProbeConf.Source has unexpected type %T", x)
	}
	return nil
}

func _ProbeConf_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ProbeConf)
	switch tag {
	case 10: // source.source_ip
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Source = &ProbeConf_SourceIp{x}
		return true, err
	case 11: // source.source_interface
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Source = &ProbeConf_SourceInterface{x}
		return true, err
	default:
		return false, nil
	}
}

func _ProbeConf_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ProbeConf)
	// source
	switch x := m.Source.(type) {
	case *ProbeConf_SourceIp:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.SourceIp)))
		n += len(x.SourceIp)
	case *ProbeConf_SourceInterface:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.SourceInterface)))
		n += len(x.SourceInterface)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ProbeConf_Header struct {
	Name                 *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value                *string  `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProbeConf_Header) Reset()         { *m = ProbeConf_Header{} }
func (m *ProbeConf_Header) String() string { return proto.CompactTextString(m) }
func (*ProbeConf_Header) ProtoMessage()    {}
func (*ProbeConf_Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_d95e376e63169085, []int{0, 0}
}
func (m *ProbeConf_Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProbeConf_Header.Unmarshal(m, b)
}
func (m *ProbeConf_Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProbeConf_Header.Marshal(b, m, deterministic)
}
func (dst *ProbeConf_Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProbeConf_Header.Merge(dst, src)
}
func (m *ProbeConf_Header) XXX_Size() int {
	return xxx_messageInfo_ProbeConf_Header.Size(m)
}
func (m *ProbeConf_Header) XXX_DiscardUnknown() {
	xxx_messageInfo_ProbeConf_Header.DiscardUnknown(m)
}

var xxx_messageInfo_ProbeConf_Header proto.InternalMessageInfo

func (m *ProbeConf_Header) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ProbeConf_Header) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*ProbeConf)(nil), "cloudprober.probes.http.ProbeConf")
	proto.RegisterType((*ProbeConf_Header)(nil), "cloudprober.probes.http.ProbeConf.Header")
	proto.RegisterEnum("cloudprober.probes.http.ProbeConf_ProtocolType", ProbeConf_ProtocolType_name, ProbeConf_ProtocolType_value)
	proto.RegisterEnum("cloudprober.probes.http.ProbeConf_Method", ProbeConf_Method_name, ProbeConf_Method_value)
}

func init() {
	proto.RegisterFile("github.com/google/cloudprober/probes/http/proto/config.proto", fileDescriptor_config_d95e376e63169085)
}

var fileDescriptor_config_d95e376e63169085 = []byte{
	// 577 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x97, 0xb5, 0x49, 0xdb, 0xd3, 0x01, 0x91, 0x35, 0xa8, 0x35, 0x09, 0xa9, 0x8c, 0x9b,
	0x02, 0x52, 0xd3, 0x4d, 0x02, 0xa1, 0x8a, 0x9b, 0xad, 0xeb, 0xe8, 0x24, 0xc6, 0x82, 0x97, 0x5d,
	0x47, 0xa9, 0x7b, 0xda, 0x46, 0xa4, 0x75, 0xb0, 0xdd, 0x8a, 0xbd, 0x05, 0x8f, 0x8c, 0xec, 0x24,
	0x55, 0xb9, 0x40, 0xda, 0x95, 0x8f, 0xfd, 0x7d, 0xbf, 0x23, 0x9f, 0x3f, 0xf0, 0x65, 0x91, 0xea,
	0xe5, 0x66, 0xda, 0xe7, 0x62, 0x15, 0x2c, 0x84, 0x58, 0x64, 0x18, 0xf0, 0x4c, 0x6c, 0x66, 0xb9,
	0x14, 0x53, 0x94, 0x81, 0x3d, 0x54, 0xb0, 0xd4, 0x3a, 0x37, 0xb1, 0x16, 0x01, 0x17, 0xeb, 0x79,
	0xba, 0xe8, 0xdb, 0x0b, 0xe9, 0xec, 0x79, 0xfb, 0x85, 0xb7, 0x6f, 0xbc, 0xa7, 0x7f, 0x1a, 0xd0,
	0x0a, 0xcd, 0x7d, 0x24, 0xd6, 0x73, 0xf2, 0x03, 0x9a, 0xd6, 0xcf, 0x45, 0x46, 0x9d, 0xae, 0xd3,
	0x7b, 0x7e, 0x1e, 0xf4, 0xff, 0x43, 0xf6, 0x77, 0x94, 0x89, 0x2c, 0x12, 0x3d, 0xe6, 0x38, 0xac,
	0x4f, 0xa2, 0x28, 0x64, 0xbb, 0x34, 0xe4, 0x0d, 0x1c, 0x49, 0xcc, 0x12, 0x9d, 0x6e, 0x31, 0xde,
	0xc8, 0x8c, 0x1e, 0x76, 0x9d, 0x5e, 0x8b, 0xb5, 0xab, 0xb7, 0x07, 0x99, 0x11, 0x02, 0xf5, 0x5c,
	0x48, 0x4d, 0x6b, 0x5d, 0xa7, 0xe7, 0x32, 0x1b, 0x93, 0xd7, 0xd0, 0x52, 0x62, 0x23, 0x39, 0xc6,
	0x69, 0x4e, 0xc1, 0x30, 0x93, 0x03, 0xd6, 0x2c, 0x9e, 0x6e, 0x72, 0xf2, 0x01, 0xfc, 0x4a, 0x5e,
	0x6b, 0x94, 0xf3, 0x84, 0x23, 0x6d, 0x97, 0xae, 0x17, 0xa5, 0xab, 0x12, 0xc8, 0x7b, 0x78, 0x26,
	0x51, 0x89, 0x6c, 0x8b, 0xf1, 0x3c, 0x95, 0x4a, 0xd3, 0x7a, 0xd7, 0xe9, 0x35, 0x87, 0xee, 0x3c,
	0xc9, 0x14, 0xb2, 0xa3, 0x52, 0xbb, 0x36, 0x12, 0xb9, 0x84, 0x13, 0xfc, 0x6d, 0x7e, 0x10, 0x4b,
	0x54, 0xb9, 0x58, 0x2b, 0x8c, 0x13, 0x15, 0xaf, 0x50, 0xcb, 0x94, 0x2b, 0xea, 0xee, 0x83, 0x9d,
	0xc2, 0xc8, 0x4a, 0xdf, 0x85, 0xba, 0x2d, 0x5c, 0xe4, 0x13, 0x74, 0xcc, 0xaf, 0x16, 0x32, 0xd5,
	0x8f, 0x31, 0x5f, 0x22, 0xff, 0x19, 0xe7, 0x89, 0xd6, 0x28, 0xd7, 0xd4, 0xb3, 0xd5, 0xbf, 0xdc,
	0xc9, 0x23, 0xa3, 0x86, 0x85, 0x48, 0xae, 0xc1, 0x5b, 0xa1, 0x5e, 0x8a, 0x19, 0x6d, 0xd8, 0xde,
	0xbf, 0x7b, 0x42, 0xef, 0x6f, 0x2d, 0x30, 0xac, 0x7d, 0x1d, 0x47, 0xac, 0xa4, 0xc9, 0x08, 0x1a,
	0x4b, 0x4c, 0x66, 0x28, 0x15, 0x6d, 0x76, 0x6b, 0xbd, 0xf6, 0x93, 0x12, 0x4d, 0x2c, 0xc1, 0x2a,
	0xd2, 0x0c, 0x65, 0x2a, 0x66, 0x8f, 0xb4, 0x65, 0x7f, 0x6c, 0x63, 0x12, 0x00, 0x91, 0xf8, 0x6b,
	0x83, 0x4a, 0xab, 0x38, 0x47, 0x19, 0xdb, 0x4c, 0x74, 0x6a, 0xc6, 0x36, 0x74, 0xce, 0x98, 0x5f,
	0x89, 0x21, 0x4a, 0x9b, 0x98, 0x7c, 0x86, 0x57, 0x3b, 0xc0, 0x0e, 0x6a, 0x9b, 0x64, 0xf1, 0x4a,
	0x21, 0xa7, 0xdc, 0x42, 0x87, 0xe7, 0x1f, 0xd9, 0x71, 0xe5, 0xb8, 0x29, 0x0d, 0xb7, 0x0a, 0xb9,
	0x99, 0x83, 0xd2, 0x89, 0x56, 0x71, 0x39, 0x8d, 0x7f, 0xe9, 0x99, 0xa5, 0xdd, 0xb3, 0xc1, 0x60,
	0x30, 0x60, 0x1d, 0x6b, 0x1c, 0x5b, 0xdf, 0x7e, 0x8e, 0x93, 0x73, 0xf0, 0x8a, 0xaa, 0x4c, 0x31,
	0xeb, 0x64, 0x85, 0x76, 0xa7, 0x5b, 0xcc, 0xc6, 0xe4, 0x18, 0xdc, 0x6d, 0x92, 0x6d, 0xb0, 0xdc,
	0xc8, 0xe2, 0x72, 0xfa, 0x16, 0x8e, 0xf6, 0xd7, 0x99, 0x34, 0xc1, 0x2e, 0xb4, 0x7f, 0x40, 0x5a,
	0xe0, 0x9a, 0xe8, 0xde, 0x77, 0x4e, 0x19, 0x78, 0x45, 0xdf, 0x49, 0x03, 0x4c, 0xe7, 0xfd, 0x03,
	0xe3, 0x0b, 0xef, 0xee, 0x23, 0xdf, 0x31, 0x4f, 0xe1, 0x43, 0xe4, 0x1f, 0x5a, 0x74, 0x7c, 0x71,
	0xe5, 0xd7, 0x08, 0x80, 0x77, 0x35, 0xfe, 0x36, 0x8e, 0xc6, 0x7e, 0xdd, 0xa4, 0x09, 0x2f, 0xa2,
	0xd1, 0xc4, 0x77, 0x49, 0x1b, 0x1a, 0x77, 0x61, 0x74, 0x73, 0xf7, 0xfd, 0xde, 0xf7, 0x2e, 0x9b,
	0xe0, 0x15, 0x7b, 0xfb, 0x37, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x90, 0x9e, 0xe3, 0xea, 0x03, 0x00,
	0x00,
}
