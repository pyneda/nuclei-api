// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pkg/service/service.proto

package service

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

// https://github.com/projectdiscovery/nuclei/blob/bb98eced070f4ae137b8cd2a7f887611bc1b9c93/v2/pkg/types/types.go#L13
type ScanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Targets           []string `protobuf:"bytes,1,rep,name=targets,proto3" json:"targets,omitempty"`
	AutomaticScan     bool     `protobuf:"varint,2,opt,name=automatic_scan,json=automaticScan,proto3" json:"automatic_scan,omitempty"`
	Tags              []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	ExcludeTags       []string `protobuf:"bytes,4,rep,name=exclude_tags,json=excludeTags,proto3" json:"exclude_tags,omitempty"`
	Workflows         []string `protobuf:"bytes,5,rep,name=workflows,proto3" json:"workflows,omitempty"`
	WorkflowUrls      []string `protobuf:"bytes,6,rep,name=workflow_urls,json=workflowUrls,proto3" json:"workflow_urls,omitempty"`
	Templates         []string `protobuf:"bytes,7,rep,name=templates,proto3" json:"templates,omitempty"`
	TemplateUrls      []string `protobuf:"bytes,8,rep,name=template_urls,json=templateUrls,proto3" json:"template_urls,omitempty"`
	ExcludedTemplates []string `protobuf:"bytes,9,rep,name=excluded_templates,json=excludedTemplates,proto3" json:"excluded_templates,omitempty"`
	ExcludeMatchers   []string `protobuf:"bytes,10,rep,name=exclude_matchers,json=excludeMatchers,proto3" json:"exclude_matchers,omitempty"`
	CustomHeaders     []string `protobuf:"bytes,11,rep,name=custom_headers,json=customHeaders,proto3" json:"custom_headers,omitempty"`
	Severities        []string `protobuf:"bytes,12,rep,name=severities,proto3" json:"severities,omitempty"`
	ExcludeSeverities []string `protobuf:"bytes,13,rep,name=exclude_severities,json=excludeSeverities,proto3" json:"exclude_severities,omitempty"`
	Authors           []string `protobuf:"bytes,14,rep,name=authors,proto3" json:"authors,omitempty"`
	Protocols         []string `protobuf:"bytes,15,rep,name=protocols,proto3" json:"protocols,omitempty"`
	ExcludeProtocols  []string `protobuf:"bytes,16,rep,name=exclude_protocols,json=excludeProtocols,proto3" json:"exclude_protocols,omitempty"`
	IncludeTags       []string `protobuf:"bytes,17,rep,name=include_tags,json=includeTags,proto3" json:"include_tags,omitempty"`
	IncludeTemplates  []string `protobuf:"bytes,18,rep,name=include_templates,json=includeTemplates,proto3" json:"include_templates,omitempty"`
	IncludeIds        []string `protobuf:"bytes,19,rep,name=include_ids,json=includeIds,proto3" json:"include_ids,omitempty"`
	ExcludeIds        []string `protobuf:"bytes,20,rep,name=exclude_ids,json=excludeIds,proto3" json:"exclude_ids,omitempty"`
	Headless          bool     `protobuf:"varint,21,opt,name=headless,proto3" json:"headless,omitempty"`
	NewTemplates      bool     `protobuf:"varint,22,opt,name=new_templates,json=newTemplates,proto3" json:"new_templates,omitempty"`
}

func (x *ScanRequest) Reset() {
	*x = ScanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_service_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanRequest) ProtoMessage() {}

func (x *ScanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_service_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanRequest.ProtoReflect.Descriptor instead.
func (*ScanRequest) Descriptor() ([]byte, []int) {
	return file_pkg_service_service_proto_rawDescGZIP(), []int{0}
}

func (x *ScanRequest) GetTargets() []string {
	if x != nil {
		return x.Targets
	}
	return nil
}

func (x *ScanRequest) GetAutomaticScan() bool {
	if x != nil {
		return x.AutomaticScan
	}
	return false
}

func (x *ScanRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *ScanRequest) GetExcludeTags() []string {
	if x != nil {
		return x.ExcludeTags
	}
	return nil
}

func (x *ScanRequest) GetWorkflows() []string {
	if x != nil {
		return x.Workflows
	}
	return nil
}

func (x *ScanRequest) GetWorkflowUrls() []string {
	if x != nil {
		return x.WorkflowUrls
	}
	return nil
}

func (x *ScanRequest) GetTemplates() []string {
	if x != nil {
		return x.Templates
	}
	return nil
}

func (x *ScanRequest) GetTemplateUrls() []string {
	if x != nil {
		return x.TemplateUrls
	}
	return nil
}

func (x *ScanRequest) GetExcludedTemplates() []string {
	if x != nil {
		return x.ExcludedTemplates
	}
	return nil
}

func (x *ScanRequest) GetExcludeMatchers() []string {
	if x != nil {
		return x.ExcludeMatchers
	}
	return nil
}

func (x *ScanRequest) GetCustomHeaders() []string {
	if x != nil {
		return x.CustomHeaders
	}
	return nil
}

func (x *ScanRequest) GetSeverities() []string {
	if x != nil {
		return x.Severities
	}
	return nil
}

func (x *ScanRequest) GetExcludeSeverities() []string {
	if x != nil {
		return x.ExcludeSeverities
	}
	return nil
}

func (x *ScanRequest) GetAuthors() []string {
	if x != nil {
		return x.Authors
	}
	return nil
}

func (x *ScanRequest) GetProtocols() []string {
	if x != nil {
		return x.Protocols
	}
	return nil
}

func (x *ScanRequest) GetExcludeProtocols() []string {
	if x != nil {
		return x.ExcludeProtocols
	}
	return nil
}

func (x *ScanRequest) GetIncludeTags() []string {
	if x != nil {
		return x.IncludeTags
	}
	return nil
}

func (x *ScanRequest) GetIncludeTemplates() []string {
	if x != nil {
		return x.IncludeTemplates
	}
	return nil
}

func (x *ScanRequest) GetIncludeIds() []string {
	if x != nil {
		return x.IncludeIds
	}
	return nil
}

func (x *ScanRequest) GetExcludeIds() []string {
	if x != nil {
		return x.ExcludeIds
	}
	return nil
}

func (x *ScanRequest) GetHeadless() bool {
	if x != nil {
		return x.Headless
	}
	return false
}

func (x *ScanRequest) GetNewTemplates() bool {
	if x != nil {
		return x.NewTemplates
	}
	return false
}

type ScanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ScanResponse) Reset() {
	*x = ScanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_service_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanResponse) ProtoMessage() {}

func (x *ScanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_service_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanResponse.ProtoReflect.Descriptor instead.
func (*ScanResponse) Descriptor() ([]byte, []int) {
	return file_pkg_service_service_proto_rawDescGZIP(), []int{1}
}

func (x *ScanResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// https://github.com/projectdiscovery/nuclei/blob/bb98eced070f4ae137b8cd2a7f887611bc1b9c93/v2/pkg/output/output.go#L103
type ScanResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId       string          `protobuf:"bytes,1,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	Template         string          `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
	Info             *ScanResultInfo `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	MatcherName      string          `protobuf:"bytes,4,opt,name=matcher_name,json=matcherName,proto3" json:"matcher_name,omitempty"`
	ExtractorName    string          `protobuf:"bytes,5,opt,name=extractor_name,json=extractorName,proto3" json:"extractor_name,omitempty"`
	Type             string          `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	Host             string          `protobuf:"bytes,7,opt,name=host,proto3" json:"host,omitempty"`
	Path             string          `protobuf:"bytes,8,opt,name=path,proto3" json:"path,omitempty"`
	Matched          string          `protobuf:"bytes,9,opt,name=matched,proto3" json:"matched,omitempty"`
	ExtractedResults []string        `protobuf:"bytes,10,rep,name=extracted_results,json=extractedResults,proto3" json:"extracted_results,omitempty"`
	Request          []byte          `protobuf:"bytes,11,opt,name=request,proto3" json:"request,omitempty"`
	Response         []byte          `protobuf:"bytes,12,opt,name=response,proto3" json:"response,omitempty"`
	Ip               string          `protobuf:"bytes,13,opt,name=ip,proto3" json:"ip,omitempty"`
	Timestamp        string          `protobuf:"bytes,14,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Interaction      *Interaction    `protobuf:"bytes,15,opt,name=interaction,proto3" json:"interaction,omitempty"`
	CurlCommand      string          `protobuf:"bytes,16,opt,name=curl_command,json=curlCommand,proto3" json:"curl_command,omitempty"`
	MatcherStatus    bool            `protobuf:"varint,17,opt,name=matcher_status,json=matcherStatus,proto3" json:"matcher_status,omitempty"` // metadata
}

func (x *ScanResult) Reset() {
	*x = ScanResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_service_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanResult) ProtoMessage() {}

func (x *ScanResult) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_service_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanResult.ProtoReflect.Descriptor instead.
func (*ScanResult) Descriptor() ([]byte, []int) {
	return file_pkg_service_service_proto_rawDescGZIP(), []int{2}
}

func (x *ScanResult) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

func (x *ScanResult) GetTemplate() string {
	if x != nil {
		return x.Template
	}
	return ""
}

func (x *ScanResult) GetInfo() *ScanResultInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *ScanResult) GetMatcherName() string {
	if x != nil {
		return x.MatcherName
	}
	return ""
}

func (x *ScanResult) GetExtractorName() string {
	if x != nil {
		return x.ExtractorName
	}
	return ""
}

func (x *ScanResult) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ScanResult) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ScanResult) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ScanResult) GetMatched() string {
	if x != nil {
		return x.Matched
	}
	return ""
}

func (x *ScanResult) GetExtractedResults() []string {
	if x != nil {
		return x.ExtractedResults
	}
	return nil
}

func (x *ScanResult) GetRequest() []byte {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *ScanResult) GetResponse() []byte {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *ScanResult) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *ScanResult) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *ScanResult) GetInteraction() *Interaction {
	if x != nil {
		return x.Interaction
	}
	return nil
}

func (x *ScanResult) GetCurlCommand() string {
	if x != nil {
		return x.CurlCommand
	}
	return ""
}

func (x *ScanResult) GetMatcherStatus() bool {
	if x != nil {
		return x.MatcherStatus
	}
	return false
}

// https://github.com/projectdiscovery/nuclei/blob/main/v2/pkg/model/model.go#L9
type ScanResultInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description    string                    `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Severity       string                    `protobuf:"bytes,3,opt,name=severity,proto3" json:"severity,omitempty"`
	Remediation    string                    `protobuf:"bytes,4,opt,name=remediation,proto3" json:"remediation,omitempty"`
	Tags           []string                  `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	References     []string                  `protobuf:"bytes,6,rep,name=references,proto3" json:"references,omitempty"`
	Classification *ScanResultClassification `protobuf:"bytes,7,opt,name=classification,proto3" json:"classification,omitempty"` // metadata
}

func (x *ScanResultInfo) Reset() {
	*x = ScanResultInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_service_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanResultInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanResultInfo) ProtoMessage() {}

func (x *ScanResultInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_service_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanResultInfo.ProtoReflect.Descriptor instead.
func (*ScanResultInfo) Descriptor() ([]byte, []int) {
	return file_pkg_service_service_proto_rawDescGZIP(), []int{3}
}

func (x *ScanResultInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ScanResultInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ScanResultInfo) GetSeverity() string {
	if x != nil {
		return x.Severity
	}
	return ""
}

func (x *ScanResultInfo) GetRemediation() string {
	if x != nil {
		return x.Remediation
	}
	return ""
}

func (x *ScanResultInfo) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *ScanResultInfo) GetReferences() []string {
	if x != nil {
		return x.References
	}
	return nil
}

func (x *ScanResultInfo) GetClassification() *ScanResultClassification {
	if x != nil {
		return x.Classification
	}
	return nil
}

type ScanResultClassification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cves       []string `protobuf:"bytes,1,rep,name=cves,proto3" json:"cves,omitempty"`
	Cwes       []string `protobuf:"bytes,2,rep,name=cwes,proto3" json:"cwes,omitempty"`
	Cpe        string   `protobuf:"bytes,3,opt,name=cpe,proto3" json:"cpe,omitempty"`
	CvssVector string   `protobuf:"bytes,4,opt,name=cvss_vector,json=cvssVector,proto3" json:"cvss_vector,omitempty"`
	CvssScore  float64  `protobuf:"fixed64,5,opt,name=cvss_score,json=cvssScore,proto3" json:"cvss_score,omitempty"`
}

func (x *ScanResultClassification) Reset() {
	*x = ScanResultClassification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_service_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanResultClassification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanResultClassification) ProtoMessage() {}

func (x *ScanResultClassification) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_service_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanResultClassification.ProtoReflect.Descriptor instead.
func (*ScanResultClassification) Descriptor() ([]byte, []int) {
	return file_pkg_service_service_proto_rawDescGZIP(), []int{4}
}

func (x *ScanResultClassification) GetCves() []string {
	if x != nil {
		return x.Cves
	}
	return nil
}

func (x *ScanResultClassification) GetCwes() []string {
	if x != nil {
		return x.Cwes
	}
	return nil
}

func (x *ScanResultClassification) GetCpe() string {
	if x != nil {
		return x.Cpe
	}
	return ""
}

func (x *ScanResultClassification) GetCvssVector() string {
	if x != nil {
		return x.CvssVector
	}
	return ""
}

func (x *ScanResultClassification) GetCvssScore() float64 {
	if x != nil {
		return x.CvssScore
	}
	return 0
}

// https://github.com/projectdiscovery/interactsh/blob/539c2ddea675a4198663a1ef9d0a15f86587ad54/pkg/server/server.go#L13
type Interaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protocol      string `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	UniqueId      string `protobuf:"bytes,2,opt,name=unique_id,json=uniqueId,proto3" json:"unique_id,omitempty"`
	FullId        string `protobuf:"bytes,3,opt,name=full_id,json=fullId,proto3" json:"full_id,omitempty"`
	Qtype         string `protobuf:"bytes,4,opt,name=qtype,proto3" json:"qtype,omitempty"`
	RawRequest    []byte `protobuf:"bytes,5,opt,name=raw_request,json=rawRequest,proto3" json:"raw_request,omitempty"`
	RawResponse   []byte `protobuf:"bytes,6,opt,name=raw_response,json=rawResponse,proto3" json:"raw_response,omitempty"`
	SmtpFrom      string `protobuf:"bytes,7,opt,name=smtp_from,json=smtpFrom,proto3" json:"smtp_from,omitempty"`
	RemoteAddress string `protobuf:"bytes,8,opt,name=remote_address,json=remoteAddress,proto3" json:"remote_address,omitempty"`
	Timestamp     string `protobuf:"bytes,9,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Interaction) Reset() {
	*x = Interaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_service_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Interaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Interaction) ProtoMessage() {}

func (x *Interaction) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_service_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Interaction.ProtoReflect.Descriptor instead.
func (*Interaction) Descriptor() ([]byte, []int) {
	return file_pkg_service_service_proto_rawDescGZIP(), []int{5}
}

func (x *Interaction) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *Interaction) GetUniqueId() string {
	if x != nil {
		return x.UniqueId
	}
	return ""
}

func (x *Interaction) GetFullId() string {
	if x != nil {
		return x.FullId
	}
	return ""
}

func (x *Interaction) GetQtype() string {
	if x != nil {
		return x.Qtype
	}
	return ""
}

func (x *Interaction) GetRawRequest() []byte {
	if x != nil {
		return x.RawRequest
	}
	return nil
}

func (x *Interaction) GetRawResponse() []byte {
	if x != nil {
		return x.RawResponse
	}
	return nil
}

func (x *Interaction) GetSmtpFrom() string {
	if x != nil {
		return x.SmtpFrom
	}
	return ""
}

func (x *Interaction) GetRemoteAddress() string {
	if x != nil {
		return x.RemoteAddress
	}
	return ""
}

func (x *Interaction) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

var File_pkg_service_service_proto protoreflect.FileDescriptor

var file_pkg_service_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x06, 0x0a, 0x0b,
	0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x69, 0x63, 0x5f, 0x73, 0x63, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x53, 0x63, 0x61, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x54,
	0x61, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x73, 0x12, 0x23, 0x0a, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x75, 0x72,
	0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x65, 0x78, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x64, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x78, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65,
	0x76, 0x65, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x65, 0x78,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x53,
	0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73,
	0x18, 0x0f, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x73, 0x12, 0x2b, 0x0a, 0x11, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x65, 0x78,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x12, 0x21,
	0x0a, 0x0c, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x11,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x54, 0x61, 0x67,
	0x73, 0x12, 0x2b, 0x0a, 0x11, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x18, 0x12, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x69, 0x6e,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x13, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x49, 0x64, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x14,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x49, 0x64, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x68, 0x65, 0x61, 0x64, 0x6c, 0x65, 0x73, 0x73, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x68, 0x65, 0x61, 0x64, 0x6c, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d,
	0x6e, 0x65, 0x77, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x18, 0x16, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x6e, 0x65, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x73, 0x22, 0x1e, 0x0a, 0x0c, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x99, 0x04, 0x0a, 0x0a, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x53, 0x63,
	0x61, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65,
	0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x65, 0x64, 0x5f,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x65,
	0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x2e, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x6c, 0x5f, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x6c, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xfb, 0x01,
	0x0a, 0x0e, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69,
	0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x0e, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x94, 0x01, 0x0a, 0x18,
	0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x76, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x63, 0x76, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x77, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x63, 0x77, 0x65, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63,
	0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x76, 0x73, 0x73, 0x5f, 0x76, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x76, 0x73, 0x73, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x76, 0x73, 0x73, 0x5f, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x63, 0x76, 0x73, 0x73, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x22, 0x9b, 0x02, 0x0a, 0x0b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x1b,
	0x0a, 0x09, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x66,
	0x75, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x75,
	0x6c, 0x6c, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x61,
	0x77, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0a, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72,
	0x61, 0x77, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0b, 0x72, 0x61, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x6d, 0x74, 0x70, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x6d, 0x74, 0x70, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x25, 0x0a, 0x0e, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x32, 0x32, 0x0a, 0x09, 0x4e, 0x75, 0x63, 0x6c, 0x65, 0x69, 0x41, 0x70, 0x69, 0x12, 0x25, 0x0a,
	0x04, 0x53, 0x63, 0x61, 0x6e, 0x12, 0x0c, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x00, 0x30, 0x01, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x79, 0x6e, 0x65, 0x64, 0x61, 0x2f, 0x6e, 0x75, 0x63, 0x6c, 0x65, 0x69,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_service_service_proto_rawDescOnce sync.Once
	file_pkg_service_service_proto_rawDescData = file_pkg_service_service_proto_rawDesc
)

func file_pkg_service_service_proto_rawDescGZIP() []byte {
	file_pkg_service_service_proto_rawDescOnce.Do(func() {
		file_pkg_service_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_service_service_proto_rawDescData)
	})
	return file_pkg_service_service_proto_rawDescData
}

var file_pkg_service_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pkg_service_service_proto_goTypes = []interface{}{
	(*ScanRequest)(nil),              // 0: ScanRequest
	(*ScanResponse)(nil),             // 1: ScanResponse
	(*ScanResult)(nil),               // 2: ScanResult
	(*ScanResultInfo)(nil),           // 3: ScanResultInfo
	(*ScanResultClassification)(nil), // 4: ScanResultClassification
	(*Interaction)(nil),              // 5: Interaction
}
var file_pkg_service_service_proto_depIdxs = []int32{
	3, // 0: ScanResult.info:type_name -> ScanResultInfo
	5, // 1: ScanResult.interaction:type_name -> Interaction
	4, // 2: ScanResultInfo.classification:type_name -> ScanResultClassification
	0, // 3: NucleiApi.Scan:input_type -> ScanRequest
	2, // 4: NucleiApi.Scan:output_type -> ScanResult
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_service_service_proto_init() }
func file_pkg_service_service_proto_init() {
	if File_pkg_service_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_service_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanRequest); i {
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
		file_pkg_service_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanResponse); i {
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
		file_pkg_service_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanResult); i {
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
		file_pkg_service_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanResultInfo); i {
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
		file_pkg_service_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanResultClassification); i {
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
		file_pkg_service_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Interaction); i {
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
			RawDescriptor: file_pkg_service_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_service_service_proto_goTypes,
		DependencyIndexes: file_pkg_service_service_proto_depIdxs,
		MessageInfos:      file_pkg_service_service_proto_msgTypes,
	}.Build()
	File_pkg_service_service_proto = out.File
	file_pkg_service_service_proto_rawDesc = nil
	file_pkg_service_service_proto_goTypes = nil
	file_pkg_service_service_proto_depIdxs = nil
}
