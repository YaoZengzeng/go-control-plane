// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/admin/v3/server_info.proto

package envoy_admin_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/go-control-plane/envoy/annotations"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type ServerInfo_State int32

const (
	ServerInfo_LIVE             ServerInfo_State = 0
	ServerInfo_DRAINING         ServerInfo_State = 1
	ServerInfo_PRE_INITIALIZING ServerInfo_State = 2
	ServerInfo_INITIALIZING     ServerInfo_State = 3
)

var ServerInfo_State_name = map[int32]string{
	0: "LIVE",
	1: "DRAINING",
	2: "PRE_INITIALIZING",
	3: "INITIALIZING",
}

var ServerInfo_State_value = map[string]int32{
	"LIVE":             0,
	"DRAINING":         1,
	"PRE_INITIALIZING": 2,
	"INITIALIZING":     3,
}

func (x ServerInfo_State) String() string {
	return proto.EnumName(ServerInfo_State_name, int32(x))
}

func (ServerInfo_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8736ae14312a45ee, []int{0, 0}
}

type CommandLineOptions_IpVersion int32

const (
	CommandLineOptions_v4 CommandLineOptions_IpVersion = 0
	CommandLineOptions_v6 CommandLineOptions_IpVersion = 1
)

var CommandLineOptions_IpVersion_name = map[int32]string{
	0: "v4",
	1: "v6",
}

var CommandLineOptions_IpVersion_value = map[string]int32{
	"v4": 0,
	"v6": 1,
}

func (x CommandLineOptions_IpVersion) String() string {
	return proto.EnumName(CommandLineOptions_IpVersion_name, int32(x))
}

func (CommandLineOptions_IpVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8736ae14312a45ee, []int{1, 0}
}

type CommandLineOptions_Mode int32

const (
	CommandLineOptions_Serve    CommandLineOptions_Mode = 0
	CommandLineOptions_Validate CommandLineOptions_Mode = 1
	CommandLineOptions_InitOnly CommandLineOptions_Mode = 2
)

var CommandLineOptions_Mode_name = map[int32]string{
	0: "Serve",
	1: "Validate",
	2: "InitOnly",
}

var CommandLineOptions_Mode_value = map[string]int32{
	"Serve":    0,
	"Validate": 1,
	"InitOnly": 2,
}

func (x CommandLineOptions_Mode) String() string {
	return proto.EnumName(CommandLineOptions_Mode_name, int32(x))
}

func (CommandLineOptions_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8736ae14312a45ee, []int{1, 1}
}

type ServerInfo struct {
	Version              string              `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	State                ServerInfo_State    `protobuf:"varint,2,opt,name=state,proto3,enum=envoy.admin.v3.ServerInfo_State" json:"state,omitempty"`
	UptimeCurrentEpoch   *duration.Duration  `protobuf:"bytes,3,opt,name=uptime_current_epoch,json=uptimeCurrentEpoch,proto3" json:"uptime_current_epoch,omitempty"`
	UptimeAllEpochs      *duration.Duration  `protobuf:"bytes,4,opt,name=uptime_all_epochs,json=uptimeAllEpochs,proto3" json:"uptime_all_epochs,omitempty"`
	HotRestartVersion    string              `protobuf:"bytes,5,opt,name=hot_restart_version,json=hotRestartVersion,proto3" json:"hot_restart_version,omitempty"`
	CommandLineOptions   *CommandLineOptions `protobuf:"bytes,6,opt,name=command_line_options,json=commandLineOptions,proto3" json:"command_line_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ServerInfo) Reset()         { *m = ServerInfo{} }
func (m *ServerInfo) String() string { return proto.CompactTextString(m) }
func (*ServerInfo) ProtoMessage()    {}
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8736ae14312a45ee, []int{0}
}

func (m *ServerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerInfo.Unmarshal(m, b)
}
func (m *ServerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerInfo.Marshal(b, m, deterministic)
}
func (m *ServerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerInfo.Merge(m, src)
}
func (m *ServerInfo) XXX_Size() int {
	return xxx_messageInfo_ServerInfo.Size(m)
}
func (m *ServerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ServerInfo proto.InternalMessageInfo

func (m *ServerInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ServerInfo) GetState() ServerInfo_State {
	if m != nil {
		return m.State
	}
	return ServerInfo_LIVE
}

func (m *ServerInfo) GetUptimeCurrentEpoch() *duration.Duration {
	if m != nil {
		return m.UptimeCurrentEpoch
	}
	return nil
}

func (m *ServerInfo) GetUptimeAllEpochs() *duration.Duration {
	if m != nil {
		return m.UptimeAllEpochs
	}
	return nil
}

func (m *ServerInfo) GetHotRestartVersion() string {
	if m != nil {
		return m.HotRestartVersion
	}
	return ""
}

func (m *ServerInfo) GetCommandLineOptions() *CommandLineOptions {
	if m != nil {
		return m.CommandLineOptions
	}
	return nil
}

type CommandLineOptions struct {
	BaseId                             uint64                       `protobuf:"varint,1,opt,name=base_id,json=baseId,proto3" json:"base_id,omitempty"`
	Concurrency                        uint32                       `protobuf:"varint,2,opt,name=concurrency,proto3" json:"concurrency,omitempty"`
	ConfigPath                         string                       `protobuf:"bytes,3,opt,name=config_path,json=configPath,proto3" json:"config_path,omitempty"`
	ConfigYaml                         string                       `protobuf:"bytes,4,opt,name=config_yaml,json=configYaml,proto3" json:"config_yaml,omitempty"`
	AllowUnknownStaticFields           bool                         `protobuf:"varint,5,opt,name=allow_unknown_static_fields,json=allowUnknownStaticFields,proto3" json:"allow_unknown_static_fields,omitempty"`
	RejectUnknownDynamicFields         bool                         `protobuf:"varint,26,opt,name=reject_unknown_dynamic_fields,json=rejectUnknownDynamicFields,proto3" json:"reject_unknown_dynamic_fields,omitempty"`
	AdminAddressPath                   string                       `protobuf:"bytes,6,opt,name=admin_address_path,json=adminAddressPath,proto3" json:"admin_address_path,omitempty"`
	LocalAddressIpVersion              CommandLineOptions_IpVersion `protobuf:"varint,7,opt,name=local_address_ip_version,json=localAddressIpVersion,proto3,enum=envoy.admin.v3.CommandLineOptions_IpVersion" json:"local_address_ip_version,omitempty"`
	LogLevel                           string                       `protobuf:"bytes,8,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	ComponentLogLevel                  string                       `protobuf:"bytes,9,opt,name=component_log_level,json=componentLogLevel,proto3" json:"component_log_level,omitempty"`
	LogFormat                          string                       `protobuf:"bytes,10,opt,name=log_format,json=logFormat,proto3" json:"log_format,omitempty"`
	LogFormatEscaped                   bool                         `protobuf:"varint,27,opt,name=log_format_escaped,json=logFormatEscaped,proto3" json:"log_format_escaped,omitempty"`
	LogPath                            string                       `protobuf:"bytes,11,opt,name=log_path,json=logPath,proto3" json:"log_path,omitempty"`
	ServiceCluster                     string                       `protobuf:"bytes,13,opt,name=service_cluster,json=serviceCluster,proto3" json:"service_cluster,omitempty"`
	ServiceNode                        string                       `protobuf:"bytes,14,opt,name=service_node,json=serviceNode,proto3" json:"service_node,omitempty"`
	ServiceZone                        string                       `protobuf:"bytes,15,opt,name=service_zone,json=serviceZone,proto3" json:"service_zone,omitempty"`
	FileFlushInterval                  *duration.Duration           `protobuf:"bytes,16,opt,name=file_flush_interval,json=fileFlushInterval,proto3" json:"file_flush_interval,omitempty"`
	DrainTime                          *duration.Duration           `protobuf:"bytes,17,opt,name=drain_time,json=drainTime,proto3" json:"drain_time,omitempty"`
	ParentShutdownTime                 *duration.Duration           `protobuf:"bytes,18,opt,name=parent_shutdown_time,json=parentShutdownTime,proto3" json:"parent_shutdown_time,omitempty"`
	Mode                               CommandLineOptions_Mode      `protobuf:"varint,19,opt,name=mode,proto3,enum=envoy.admin.v3.CommandLineOptions_Mode" json:"mode,omitempty"`
	DisableHotRestart                  bool                         `protobuf:"varint,22,opt,name=disable_hot_restart,json=disableHotRestart,proto3" json:"disable_hot_restart,omitempty"`
	EnableMutexTracing                 bool                         `protobuf:"varint,23,opt,name=enable_mutex_tracing,json=enableMutexTracing,proto3" json:"enable_mutex_tracing,omitempty"`
	RestartEpoch                       uint32                       `protobuf:"varint,24,opt,name=restart_epoch,json=restartEpoch,proto3" json:"restart_epoch,omitempty"`
	CpusetThreads                      bool                         `protobuf:"varint,25,opt,name=cpuset_threads,json=cpusetThreads,proto3" json:"cpuset_threads,omitempty"`
	DisabledExtensions                 []string                     `protobuf:"bytes,28,rep,name=disabled_extensions,json=disabledExtensions,proto3" json:"disabled_extensions,omitempty"`
	HiddenEnvoyDeprecatedMaxStats      uint64                       `protobuf:"varint,20,opt,name=hidden_envoy_deprecated_max_stats,json=hiddenEnvoyDeprecatedMaxStats,proto3" json:"hidden_envoy_deprecated_max_stats,omitempty"`                    // Deprecated: Do not use.
	HiddenEnvoyDeprecatedMaxObjNameLen uint64                       `protobuf:"varint,21,opt,name=hidden_envoy_deprecated_max_obj_name_len,json=hiddenEnvoyDeprecatedMaxObjNameLen,proto3" json:"hidden_envoy_deprecated_max_obj_name_len,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral               struct{}                     `json:"-"`
	XXX_unrecognized                   []byte                       `json:"-"`
	XXX_sizecache                      int32                        `json:"-"`
}

func (m *CommandLineOptions) Reset()         { *m = CommandLineOptions{} }
func (m *CommandLineOptions) String() string { return proto.CompactTextString(m) }
func (*CommandLineOptions) ProtoMessage()    {}
func (*CommandLineOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_8736ae14312a45ee, []int{1}
}

func (m *CommandLineOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandLineOptions.Unmarshal(m, b)
}
func (m *CommandLineOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandLineOptions.Marshal(b, m, deterministic)
}
func (m *CommandLineOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandLineOptions.Merge(m, src)
}
func (m *CommandLineOptions) XXX_Size() int {
	return xxx_messageInfo_CommandLineOptions.Size(m)
}
func (m *CommandLineOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandLineOptions.DiscardUnknown(m)
}

var xxx_messageInfo_CommandLineOptions proto.InternalMessageInfo

func (m *CommandLineOptions) GetBaseId() uint64 {
	if m != nil {
		return m.BaseId
	}
	return 0
}

func (m *CommandLineOptions) GetConcurrency() uint32 {
	if m != nil {
		return m.Concurrency
	}
	return 0
}

func (m *CommandLineOptions) GetConfigPath() string {
	if m != nil {
		return m.ConfigPath
	}
	return ""
}

func (m *CommandLineOptions) GetConfigYaml() string {
	if m != nil {
		return m.ConfigYaml
	}
	return ""
}

func (m *CommandLineOptions) GetAllowUnknownStaticFields() bool {
	if m != nil {
		return m.AllowUnknownStaticFields
	}
	return false
}

func (m *CommandLineOptions) GetRejectUnknownDynamicFields() bool {
	if m != nil {
		return m.RejectUnknownDynamicFields
	}
	return false
}

func (m *CommandLineOptions) GetAdminAddressPath() string {
	if m != nil {
		return m.AdminAddressPath
	}
	return ""
}

func (m *CommandLineOptions) GetLocalAddressIpVersion() CommandLineOptions_IpVersion {
	if m != nil {
		return m.LocalAddressIpVersion
	}
	return CommandLineOptions_v4
}

func (m *CommandLineOptions) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

func (m *CommandLineOptions) GetComponentLogLevel() string {
	if m != nil {
		return m.ComponentLogLevel
	}
	return ""
}

func (m *CommandLineOptions) GetLogFormat() string {
	if m != nil {
		return m.LogFormat
	}
	return ""
}

func (m *CommandLineOptions) GetLogFormatEscaped() bool {
	if m != nil {
		return m.LogFormatEscaped
	}
	return false
}

func (m *CommandLineOptions) GetLogPath() string {
	if m != nil {
		return m.LogPath
	}
	return ""
}

func (m *CommandLineOptions) GetServiceCluster() string {
	if m != nil {
		return m.ServiceCluster
	}
	return ""
}

func (m *CommandLineOptions) GetServiceNode() string {
	if m != nil {
		return m.ServiceNode
	}
	return ""
}

func (m *CommandLineOptions) GetServiceZone() string {
	if m != nil {
		return m.ServiceZone
	}
	return ""
}

func (m *CommandLineOptions) GetFileFlushInterval() *duration.Duration {
	if m != nil {
		return m.FileFlushInterval
	}
	return nil
}

func (m *CommandLineOptions) GetDrainTime() *duration.Duration {
	if m != nil {
		return m.DrainTime
	}
	return nil
}

func (m *CommandLineOptions) GetParentShutdownTime() *duration.Duration {
	if m != nil {
		return m.ParentShutdownTime
	}
	return nil
}

func (m *CommandLineOptions) GetMode() CommandLineOptions_Mode {
	if m != nil {
		return m.Mode
	}
	return CommandLineOptions_Serve
}

func (m *CommandLineOptions) GetDisableHotRestart() bool {
	if m != nil {
		return m.DisableHotRestart
	}
	return false
}

func (m *CommandLineOptions) GetEnableMutexTracing() bool {
	if m != nil {
		return m.EnableMutexTracing
	}
	return false
}

func (m *CommandLineOptions) GetRestartEpoch() uint32 {
	if m != nil {
		return m.RestartEpoch
	}
	return 0
}

func (m *CommandLineOptions) GetCpusetThreads() bool {
	if m != nil {
		return m.CpusetThreads
	}
	return false
}

func (m *CommandLineOptions) GetDisabledExtensions() []string {
	if m != nil {
		return m.DisabledExtensions
	}
	return nil
}

// Deprecated: Do not use.
func (m *CommandLineOptions) GetHiddenEnvoyDeprecatedMaxStats() uint64 {
	if m != nil {
		return m.HiddenEnvoyDeprecatedMaxStats
	}
	return 0
}

// Deprecated: Do not use.
func (m *CommandLineOptions) GetHiddenEnvoyDeprecatedMaxObjNameLen() uint64 {
	if m != nil {
		return m.HiddenEnvoyDeprecatedMaxObjNameLen
	}
	return 0
}

func init() {
	proto.RegisterEnum("envoy.admin.v3.ServerInfo_State", ServerInfo_State_name, ServerInfo_State_value)
	proto.RegisterEnum("envoy.admin.v3.CommandLineOptions_IpVersion", CommandLineOptions_IpVersion_name, CommandLineOptions_IpVersion_value)
	proto.RegisterEnum("envoy.admin.v3.CommandLineOptions_Mode", CommandLineOptions_Mode_name, CommandLineOptions_Mode_value)
	proto.RegisterType((*ServerInfo)(nil), "envoy.admin.v3.ServerInfo")
	proto.RegisterType((*CommandLineOptions)(nil), "envoy.admin.v3.CommandLineOptions")
}

func init() { proto.RegisterFile("envoy/admin/v3/server_info.proto", fileDescriptor_8736ae14312a45ee) }

var fileDescriptor_8736ae14312a45ee = []byte{
	// 1115 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0xdf, 0x6e, 0x1b, 0xc5,
	0x17, 0xc7, 0x6b, 0xc7, 0x71, 0xed, 0xc9, 0xbf, 0xcd, 0x34, 0xfd, 0x75, 0x9b, 0x34, 0xfd, 0xb9,
	0xae, 0x4a, 0x7d, 0xd1, 0xda, 0xa8, 0x45, 0x15, 0x14, 0x71, 0x91, 0x36, 0x6e, 0x59, 0x70, 0xd3,
	0x68, 0x13, 0x2a, 0xd1, 0x0b, 0x46, 0xe3, 0x9d, 0xb1, 0x3d, 0x61, 0x76, 0x66, 0xb5, 0x3b, 0xeb,
	0xc6, 0x5c, 0x71, 0xc9, 0x33, 0xf0, 0x24, 0x88, 0x7b, 0x24, 0x6e, 0xe1, 0x29, 0xb8, 0x02, 0x89,
	0x27, 0x40, 0x73, 0x66, 0xfd, 0x27, 0x09, 0x10, 0xae, 0xac, 0x3d, 0xe7, 0x73, 0xbe, 0x9e, 0x39,
	0xe7, 0xec, 0x77, 0x51, 0x83, 0xab, 0xb1, 0x9e, 0x74, 0x28, 0x8b, 0x85, 0xea, 0x8c, 0x1f, 0x77,
	0x32, 0x9e, 0x8e, 0x79, 0x4a, 0x84, 0x1a, 0xe8, 0x76, 0x92, 0x6a, 0xa3, 0xf1, 0x3a, 0x10, 0x6d,
	0x20, 0xda, 0xe3, 0xc7, 0xdb, 0xb7, 0x87, 0x5a, 0x0f, 0x25, 0xef, 0x40, 0xb6, 0x9f, 0x0f, 0x3a,
	0x2c, 0x4f, 0xa9, 0x11, 0x5a, 0x39, 0x7e, 0xfb, 0x6e, 0xa1, 0xa8, 0x94, 0x36, 0x10, 0xcf, 0x3a,
	0x8c, 0x27, 0x29, 0x8f, 0x16, 0xa1, 0xdd, 0x9c, 0x25, 0xf4, 0x0c, 0x93, 0x19, 0x6a, 0xf2, 0xac,
	0x48, 0xdf, 0xb9, 0x90, 0x1e, 0xf3, 0x34, 0x13, 0x5a, 0x09, 0x35, 0x74, 0x48, 0xf3, 0xb7, 0x25,
	0x84, 0x8e, 0xe0, 0xb0, 0x81, 0x1a, 0x68, 0xec, 0xa3, 0xab, 0x05, 0xe2, 0x97, 0x1a, 0xa5, 0x56,
	0x3d, 0x9c, 0x3e, 0xe2, 0x27, 0x68, 0xd9, 0x6a, 0x73, 0xbf, 0xdc, 0x28, 0xb5, 0xd6, 0x1f, 0x35,
	0xda, 0x67, 0xef, 0xd3, 0x9e, 0x8b, 0xb4, 0x8f, 0x2c, 0x17, 0x3a, 0x1c, 0x7f, 0x8e, 0xb6, 0xf2,
	0xc4, 0x88, 0x98, 0x93, 0x28, 0x4f, 0x53, 0xae, 0x0c, 0xe1, 0x89, 0x8e, 0x46, 0xfe, 0x52, 0xa3,
	0xd4, 0x5a, 0x79, 0x74, 0xb3, 0xed, 0xda, 0xd0, 0x9e, 0xb6, 0xa1, 0xbd, 0x5f, 0xb4, 0x21, 0xc4,
	0xae, 0xec, 0xb9, 0xab, 0xea, 0xda, 0x22, 0xdc, 0x45, 0x9b, 0x85, 0x18, 0x95, 0xd2, 0x09, 0x65,
	0x7e, 0xe5, 0x32, 0xa5, 0x0d, 0x57, 0xb3, 0x27, 0x25, 0xa8, 0x64, 0xb8, 0x8d, 0xae, 0x8d, 0xb4,
	0x21, 0x29, 0xcf, 0x0c, 0x4d, 0x0d, 0x99, 0xde, 0x78, 0x19, 0x6e, 0xbc, 0x39, 0xd2, 0x26, 0x74,
	0x99, 0x37, 0xc5, 0xdd, 0x8f, 0xd1, 0x56, 0xa4, 0xe3, 0x98, 0x2a, 0x46, 0xa4, 0x50, 0x9c, 0xe8,
	0x04, 0xba, 0xe9, 0x57, 0xe1, 0x9f, 0x9b, 0xe7, 0x5b, 0xf1, 0xdc, 0xb1, 0x3d, 0xa1, 0xf8, 0x6b,
	0x47, 0x86, 0x38, 0xba, 0x10, 0x6b, 0xbe, 0x44, 0xcb, 0xd0, 0x29, 0x5c, 0x43, 0x95, 0x5e, 0xf0,
	0xa6, 0xeb, 0x5d, 0xc1, 0xab, 0xa8, 0xb6, 0x1f, 0xee, 0x05, 0x07, 0xc1, 0xc1, 0x4b, 0xaf, 0x84,
	0xb7, 0x90, 0x77, 0x18, 0x76, 0x49, 0x70, 0x10, 0x1c, 0x07, 0x7b, 0xbd, 0xe0, 0xad, 0x8d, 0x96,
	0xb1, 0x87, 0x56, 0xcf, 0x44, 0x96, 0x9e, 0xde, 0xfb, 0xfe, 0xa7, 0xef, 0x6e, 0x37, 0xd0, 0xed,
	0x33, 0xc7, 0x78, 0x44, 0x65, 0x32, 0xa2, 0x0b, 0x63, 0x69, 0xfe, 0xb1, 0x82, 0xf0, 0xc5, 0xa3,
	0xe1, 0x1b, 0xe8, 0x6a, 0x9f, 0x66, 0x9c, 0x08, 0x06, 0x23, 0xaf, 0x84, 0x55, 0xfb, 0x18, 0x30,
	0xdc, 0x40, 0x2b, 0x91, 0x56, 0x6e, 0x6a, 0xd1, 0x04, 0xe6, 0xbe, 0x16, 0x2e, 0x86, 0xf0, 0xff,
	0x81, 0x18, 0x88, 0x21, 0x49, 0xa8, 0x71, 0x23, 0xad, 0x87, 0xc8, 0x85, 0x0e, 0xa9, 0x19, 0x2d,
	0x00, 0x13, 0x1a, 0x4b, 0x98, 0xd4, 0x0c, 0xf8, 0x92, 0xc6, 0x12, 0x7f, 0x82, 0x76, 0xa8, 0x94,
	0xfa, 0x1d, 0xc9, 0xd5, 0xd7, 0x4a, 0xbf, 0x53, 0xc4, 0x2e, 0x8d, 0x88, 0xc8, 0x40, 0x70, 0xc9,
	0x32, 0x98, 0x48, 0x2d, 0xf4, 0x01, 0xf9, 0xc2, 0x11, 0x47, 0x00, 0xbc, 0x80, 0x3c, 0xde, 0x43,
	0xbb, 0x29, 0x3f, 0xe1, 0x91, 0x99, 0xd5, 0xb3, 0x89, 0xa2, 0xf1, 0x5c, 0x60, 0x1b, 0x04, 0xb6,
	0x1d, 0x54, 0x28, 0xec, 0x3b, 0xa4, 0x90, 0x78, 0x80, 0x30, 0x74, 0x8c, 0x50, 0xc6, 0x52, 0x9e,
	0x65, 0xee, 0x2a, 0x55, 0x38, 0xa9, 0x07, 0x99, 0x3d, 0x97, 0x80, 0x0b, 0x71, 0xe4, 0x4b, 0x1d,
	0x51, 0x39, 0xa3, 0x45, 0x32, 0x5b, 0x9f, 0xab, 0xf0, 0x62, 0x3c, 0xb8, 0x7c, 0x1b, 0xda, 0x41,
	0x52, 0x6c, 0x56, 0x78, 0x1d, 0xd4, 0x8a, 0x7f, 0x98, 0x85, 0xf1, 0x0e, 0xaa, 0x4b, 0x3d, 0x24,
	0x92, 0x8f, 0xb9, 0xf4, 0x6b, 0x70, 0x96, 0x9a, 0xd4, 0xc3, 0x9e, 0x7d, 0xb6, 0xdb, 0x1b, 0xe9,
	0x38, 0xd1, 0xca, 0xbe, 0x4c, 0x73, 0xac, 0xee, 0xb6, 0x77, 0x96, 0xea, 0x4d, 0xf9, 0x5d, 0x84,
	0x2c, 0x35, 0xd0, 0x69, 0x4c, 0x8d, 0x8f, 0x00, 0xb3, 0xf2, 0x2f, 0x20, 0x60, 0x1b, 0x30, 0x4f,
	0x13, 0x9e, 0x45, 0x34, 0xe1, 0xcc, 0xdf, 0x81, 0xc6, 0x79, 0x33, 0xac, 0xeb, 0xe2, 0xf8, 0x26,
	0xb2, 0x07, 0x71, 0x4d, 0x5a, 0x71, 0x0e, 0x21, 0xb5, 0x1b, 0xf6, 0x7d, 0xb4, 0x61, 0x6d, 0x4f,
	0x44, 0x9c, 0x44, 0x32, 0xcf, 0x0c, 0x4f, 0xfd, 0x35, 0x20, 0xd6, 0x8b, 0xf0, 0x73, 0x17, 0xc5,
	0x77, 0xd0, 0xea, 0x14, 0x54, 0x9a, 0x71, 0x7f, 0x1d, 0xa8, 0x95, 0x22, 0x76, 0xa0, 0x19, 0x5f,
	0x44, 0xbe, 0xd1, 0x8a, 0xfb, 0x1b, 0x67, 0x90, 0xb7, 0x5a, 0x71, 0x1c, 0xa0, 0x6b, 0x03, 0x21,
	0x39, 0x19, 0xc8, 0x3c, 0x1b, 0x11, 0xa1, 0x0c, 0x4f, 0xc7, 0x54, 0xfa, 0xde, 0x65, 0x6e, 0xb0,
	0x69, 0xab, 0x5e, 0xd8, 0xa2, 0xa0, 0xa8, 0xc1, 0x1f, 0x22, 0xc4, 0x52, 0x2a, 0x14, 0xb1, 0x36,
	0xe1, 0x6f, 0x5e, 0xa6, 0x50, 0x07, 0xf8, 0x58, 0xc4, 0xe0, 0x6e, 0x09, 0x05, 0x57, 0xcb, 0x46,
	0xb9, 0x61, 0x76, 0x03, 0x41, 0x03, 0x5f, 0xea, 0x6e, 0xae, 0xec, 0xa8, 0xa8, 0x02, 0xb1, 0x8f,
	0x51, 0x25, 0xb6, 0xfd, 0xb8, 0x06, 0x8b, 0x74, 0xff, 0x3f, 0x2c, 0xd2, 0x2b, 0xcd, 0x78, 0x08,
	0x45, 0x76, 0x2b, 0x98, 0xc8, 0x68, 0x5f, 0x72, 0xb2, 0xe0, 0x6d, 0xfe, 0xff, 0x60, 0x8e, 0x9b,
	0x45, 0xea, 0xd3, 0x99, 0xb5, 0xe1, 0xf7, 0xd1, 0x16, 0x57, 0x80, 0xc7, 0xb9, 0xe1, 0xa7, 0xc4,
	0xa4, 0x34, 0x12, 0x6a, 0xe8, 0xdf, 0x80, 0x02, 0xec, 0x72, 0xaf, 0x6c, 0xea, 0xd8, 0x65, 0xf0,
	0x5d, 0xb4, 0x36, 0x75, 0x4c, 0x67, 0xe1, 0x3e, 0x38, 0xc2, 0x6a, 0x11, 0x74, 0x0e, 0x7d, 0x0f,
	0xad, 0x47, 0x49, 0x9e, 0x71, 0x43, 0xcc, 0x28, 0xe5, 0x94, 0x65, 0xfe, 0x4d, 0x10, 0x5c, 0x73,
	0xd1, 0x63, 0x17, 0xc4, 0x9d, 0xd9, 0x69, 0x19, 0xe1, 0xa7, 0x86, 0xab, 0x0c, 0x0c, 0xf5, 0x56,
	0x63, 0xa9, 0x55, 0x0f, 0xf1, 0x34, 0xd5, 0x9d, 0x65, 0x70, 0x88, 0xee, 0x8c, 0x04, 0x63, 0x5c,
	0x11, 0xe8, 0x0a, 0x99, 0x7e, 0x0b, 0x39, 0x23, 0x31, 0x3d, 0x05, 0xdb, 0xc8, 0xfc, 0x2d, 0xeb,
	0x5f, 0xcf, 0x6a, 0x3f, 0xfc, 0xfe, 0xe7, 0xaf, 0xcb, 0x25, 0xbf, 0x14, 0xee, 0xba, 0x92, 0xae,
	0xad, 0xd8, 0x9f, 0x15, 0xbc, 0xa2, 0xa7, 0xd6, 0x44, 0x32, 0xfc, 0x15, 0x6a, 0xfd, 0x9b, 0xa6,
	0xee, 0x9f, 0x10, 0x45, 0x63, 0x4e, 0x24, 0x57, 0xfe, 0xf5, 0x73, 0xd2, 0xcd, 0x7f, 0x92, 0x7e,
	0xdd, 0x3f, 0x39, 0xa0, 0x31, 0xef, 0x71, 0xd5, 0xdc, 0x41, 0xf5, 0xf9, 0x2b, 0x5d, 0x45, 0xe5,
	0xf1, 0x07, 0xde, 0x15, 0xf8, 0x7d, 0xe2, 0x95, 0x9a, 0x0f, 0x51, 0xc5, 0x4e, 0x0f, 0xd7, 0xd1,
	0x32, 0x78, 0xb4, 0x73, 0xff, 0x37, 0x54, 0x0a, 0x46, 0x0d, 0xf7, 0x4a, 0xf6, 0x29, 0x50, 0xc2,
	0xbc, 0x56, 0x72, 0xe2, 0x95, 0x9f, 0x3e, 0xb4, 0x1e, 0xdf, 0x42, 0xef, 0xfd, 0x9d, 0xc7, 0x5f,
	0x5c, 0x8c, 0xcf, 0x2a, 0xb5, 0x55, 0x6f, 0xed, 0xd9, 0x47, 0x3f, 0x7e, 0xfb, 0xf3, 0x2f, 0xd5,
	0xb2, 0x57, 0x46, 0xb7, 0x84, 0x76, 0xeb, 0x94, 0xa4, 0xfa, 0x74, 0x72, 0x6e, 0xb3, 0x9e, 0x6d,
	0xcc, 0xbf, 0x12, 0x87, 0x76, 0x53, 0x0f, 0x4b, 0xfd, 0x2a, 0xac, 0xec, 0xe3, 0xbf, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xaa, 0x16, 0xc5, 0x10, 0xd9, 0x08, 0x00, 0x00,
}
