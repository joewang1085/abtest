package abtest_proto

import (
	context "context"
	reflect "reflect"
	sync "sync"

	"time"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Project is a completed AB test Lab config. Different Projects are independent of each other
// Next ID: 2
type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the Project
	Id                string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" dgraph:"id_2"`
	DgraphId          string     `dgraph:"uid"`
	DgraphType        []string   `dgraph:"dgraph.type"`
	DgraphCreatedDate *time.Time `dgraph:"_created_date @index(month)"`
	DgraphExpiredDate *time.Time `dgraph:"_expired_date @index(month)"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_liulishuo_algorithm_algapi_abtest_abtest_proto_rawDescGZIP(), []int{0}
}

func (x *Project) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Layer is a time dimension in a Project which contains Zones for random diversion of users.
// Next ID: 4
type Layer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the Layer
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" dgraph:"id_3"`
	// ParentZones is where the users of this layer come from
	ParentZones []*Zone `protobuf:"bytes,2,rep,name=parent_zones,json=parentZones,proto3" json:"parent_zones,omitempty" dgraph:"parent_zones"`
	// TotalWeight is the total diversion of the users
	TotalWeight       int32      `protobuf:"varint,3,opt,name=total_weight,json=totalWeight,proto3" json:"total_weight,omitempty" dgraph:"total_weight"`
	DgraphId          string     `dgraph:"uid"`
	DgraphType        []string   `dgraph:"dgraph.type"`
	DgraphCreatedDate *time.Time `dgraph:"_created_date @index(month)"`
	DgraphExpiredDate *time.Time `dgraph:"_expired_date @index(month)"`
}

func (x *Layer) Reset() {
	*x = Layer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Layer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Layer) ProtoMessage() {}

func (x *Layer) ProtoReflect() protoreflect.Message {
	mi := &file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Layer.ProtoReflect.Descriptor instead.
func (*Layer) Descriptor() ([]byte, []int) {
	return file_liulishuo_algorithm_algapi_abtest_abtest_proto_rawDescGZIP(), []int{1}
}

func (x *Layer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Layer) GetParentZones() []*Zone {
	if x != nil {
		return x.ParentZones
	}
	return nil
}

func (x *Layer) GetTotalWeight() int32 {
	if x != nil {
		return x.TotalWeight
	}
	return 0
}

// Zone is a lab in the layer
// Next ID: 8
type Zone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the Zone
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" dgraph:"id_4"`
	// The Project of the Zone
	Project *Project `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty" dgraph:"project"`
	// The Layer of the Zone
	Layer *Layer `protobuf:"bytes,3,opt,name=layer,proto3" json:"layer,omitempty" dgraph:"layer"`
	// The Weight of the Zone
	Weight *Weight `protobuf:"bytes,4,opt,name=weight,proto3" json:"weight,omitempty" dgraph:"weight"`
	// The Label of the Zone which to match the logic code
	Label string `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty" dgraph:"label"`
	// The Description of the Zone
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty" dgraph:"description"`
	// The UserGroups of the Zone
	UserGroups        []string   `protobuf:"bytes,7,rep,name=user_groups,json=userGroups,proto3" json:"user_groups,omitempty" dgraph:"user_groups"`
	DgraphId          string     `dgraph:"uid"`
	DgraphType        []string   `dgraph:"dgraph.type"`
	DgraphCreatedDate *time.Time `dgraph:"_created_date @index(month)"`
	DgraphExpiredDate *time.Time `dgraph:"_expired_date @index(month)"`
}

func (x *Zone) Reset() {
	*x = Zone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Zone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Zone) ProtoMessage() {}

func (x *Zone) ProtoReflect() protoreflect.Message {
	mi := &file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Zone.ProtoReflect.Descriptor instead.
func (*Zone) Descriptor() ([]byte, []int) {
	return file_liulishuo_algorithm_algapi_abtest_abtest_proto_rawDescGZIP(), []int{2}
}

func (x *Zone) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Zone) GetProject() *Project {
	if x != nil {
		return x.Project
	}
	return nil
}

func (x *Zone) GetLayer() *Layer {
	if x != nil {
		return x.Layer
	}
	return nil
}

func (x *Zone) GetWeight() *Weight {
	if x != nil {
		return x.Weight
	}
	return nil
}

func (x *Zone) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Zone) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Zone) GetUserGroups() []string {
	if x != nil {
		return x.UserGroups
	}
	return nil
}

// The Weight of Zone
// Next ID: 3
type Weight struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Min number
	Min int32 `protobuf:"varint,1,opt,name=min,proto3" json:"min,omitempty" dgraph:"min"`
	// Max number
	Max               int32      `protobuf:"varint,2,opt,name=max,proto3" json:"max,omitempty" dgraph:"max"`
	DgraphId          string     `dgraph:"uid"`
	DgraphType        []string   `dgraph:"dgraph.type"`
	DgraphCreatedDate *time.Time `dgraph:"_created_date @index(month)"`
	DgraphExpiredDate *time.Time `dgraph:"_expired_date @index(month)"`
}

func (x *Weight) Reset() {
	*x = Weight{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Weight) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Weight) ProtoMessage() {}

func (x *Weight) ProtoReflect() protoreflect.Message {
	mi := &file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Weight.ProtoReflect.Descriptor instead.
func (*Weight) Descriptor() ([]byte, []int) {
	return file_liulishuo_algorithm_algapi_abtest_abtest_proto_rawDescGZIP(), []int{3}
}

func (x *Weight) GetMin() int32 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *Weight) GetMax() int32 {
	if x != nil {
		return x.Max
	}
	return 0
}

// The request of GetABTest
// Next ID: 4
type GetABTestZoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The ProjectID to match
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty" dgraph:"project_id"`
	// Required. The Hashkey for random diversion
	HashKey string `protobuf:"bytes,2,opt,name=hash_key,json=hashKey,proto3" json:"hash_key,omitempty" dgraph:"hash_key"`
	// Required. The LayerID to match
	LayerId           string     `protobuf:"bytes,3,opt,name=layer_id,json=layerId,proto3" json:"layer_id,omitempty" dgraph:"layer_id"`
	DgraphId          string     `dgraph:"uid"`
	DgraphType        []string   `dgraph:"dgraph.type"`
	DgraphCreatedDate *time.Time `dgraph:"_created_date @index(month)"`
	DgraphExpiredDate *time.Time `dgraph:"_expired_date @index(month)"`
}

func (x *GetABTestZoneRequest) Reset() {
	*x = GetABTestZoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetABTestZoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetABTestZoneRequest) ProtoMessage() {}

func (x *GetABTestZoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetABTestZoneRequest.ProtoReflect.Descriptor instead.
func (*GetABTestZoneRequest) Descriptor() ([]byte, []int) {
	return file_liulishuo_algorithm_algapi_abtest_abtest_proto_rawDescGZIP(), []int{4}
}

func (x *GetABTestZoneRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *GetABTestZoneRequest) GetHashKey() string {
	if x != nil {
		return x.HashKey
	}
	return ""
}

func (x *GetABTestZoneRequest) GetLayerId() string {
	if x != nil {
		return x.LayerId
	}
	return ""
}

// The response of GetABTest
// Next ID: 2
type GetABTestZoneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The target Zone of once random diversion
	Zone              *Zone      `protobuf:"bytes,1,opt,name=zone,proto3" json:"zone,omitempty" dgraph:"zone"`
	DgraphId          string     `dgraph:"uid"`
	DgraphType        []string   `dgraph:"dgraph.type"`
	DgraphCreatedDate *time.Time `dgraph:"_created_date @index(month)"`
	DgraphExpiredDate *time.Time `dgraph:"_expired_date @index(month)"`
}

func (x *GetABTestZoneResponse) Reset() {
	*x = GetABTestZoneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetABTestZoneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetABTestZoneResponse) ProtoMessage() {}

func (x *GetABTestZoneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetABTestZoneResponse.ProtoReflect.Descriptor instead.
func (*GetABTestZoneResponse) Descriptor() ([]byte, []int) {
	return file_liulishuo_algorithm_algapi_abtest_abtest_proto_rawDescGZIP(), []int{5}
}

func (x *GetABTestZoneResponse) GetZone() *Zone {
	if x != nil {
		return x.Zone
	}
	return nil
}

// LabData is a self defined data of the service
// Next ID: 3
type LabData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key is the key of lab data
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty" dgraph:"key"`
	// Value is the value of lab data
	Value             string     `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty" dgraph:"value"`
	DgraphId          string     `dgraph:"uid"`
	DgraphType        []string   `dgraph:"dgraph.type"`
	DgraphCreatedDate *time.Time `dgraph:"_created_date @index(month)"`
	DgraphExpiredDate *time.Time `dgraph:"_expired_date @index(month)"`
}

func (x *LabData) Reset() {
	*x = LabData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabData) ProtoMessage() {}

func (x *LabData) ProtoReflect() protoreflect.Message {
	mi := &file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabData.ProtoReflect.Descriptor instead.
func (*LabData) Descriptor() ([]byte, []int) {
	return file_liulishuo_algorithm_algapi_abtest_abtest_proto_rawDescGZIP(), []int{6}
}

func (x *LabData) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *LabData) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// The request of PushABTestData
// Next ID: 6
type PushABTestDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The ProjectID to match
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty" dgraph:"project_id_1"`
	// Required. The Hashkey for random diversion
	HashKey string `protobuf:"bytes,2,opt,name=hash_key,json=hashKey,proto3" json:"hash_key,omitempty" dgraph:"hash_key_1"`
	// The type of Hashkey
	KeyType string `protobuf:"bytes,3,opt,name=key_type,json=keyType,proto3" json:"key_type,omitempty" dgraph:"key_type"`
	// Data is self defined data
	Data []*LabData `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty" dgraph:"data"`
	// The login ID
	LoginId           string     `protobuf:"bytes,5,opt,name=login_id,json=loginId,proto3" json:"login_id,omitempty" dgraph:"login_id"`
	DgraphId          string     `dgraph:"uid"`
	DgraphType        []string   `dgraph:"dgraph.type"`
	DgraphCreatedDate *time.Time `dgraph:"_created_date @index(month)"`
	DgraphExpiredDate *time.Time `dgraph:"_expired_date @index(month)"`
}

func (x *PushABTestDataRequest) Reset() {
	*x = PushABTestDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushABTestDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushABTestDataRequest) ProtoMessage() {}

func (x *PushABTestDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushABTestDataRequest.ProtoReflect.Descriptor instead.
func (*PushABTestDataRequest) Descriptor() ([]byte, []int) {
	return file_liulishuo_algorithm_algapi_abtest_abtest_proto_rawDescGZIP(), []int{7}
}

func (x *PushABTestDataRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *PushABTestDataRequest) GetHashKey() string {
	if x != nil {
		return x.HashKey
	}
	return ""
}

func (x *PushABTestDataRequest) GetKeyType() string {
	if x != nil {
		return x.KeyType
	}
	return ""
}

func (x *PushABTestDataRequest) GetData() []*LabData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PushABTestDataRequest) GetLoginId() string {
	if x != nil {
		return x.LoginId
	}
	return ""
}

// The response of PushABTestData
// Next ID: 1
type PushABTestDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PushABTestDataResponse) Reset() {
	*x = PushABTestDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushABTestDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushABTestDataResponse) ProtoMessage() {}

func (x *PushABTestDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushABTestDataResponse.ProtoReflect.Descriptor instead.
func (*PushABTestDataResponse) Descriptor() ([]byte, []int) {
	return file_liulishuo_algorithm_algapi_abtest_abtest_proto_rawDescGZIP(), []int{8}
}

// The Message sent to mq
// Next ID: 8
type ABTestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of project
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty" dgraph:"project_id_2"`
	// The hash key
	HashKey string `protobuf:"bytes,2,opt,name=hash_key,json=hashKey,proto3" json:"hash_key,omitempty" dgraph:"hash_key_2"`
	// The targeted Strategy
	Strategy string `protobuf:"bytes,3,opt,name=strategy,proto3" json:"strategy,omitempty" dgraph:"strategy"`
	// The type of the hash key
	KeyType string `protobuf:"bytes,4,opt,name=key_type,json=keyType,proto3" json:"key_type,omitempty" dgraph:"key_type_1"`
	// The login id may usually be userID
	LoginId string `protobuf:"bytes,5,opt,name=login_id,json=loginId,proto3" json:"login_id,omitempty" dgraph:"login_id_1"`
	// The self defined extra data may usually be empty
	Ext string `protobuf:"bytes,6,opt,name=ext,proto3" json:"ext,omitempty" dgraph:"ext"`
	// the created time of message
	CreatedTime       string     `protobuf:"bytes,7,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty" dgraph:"created_time"`
	DgraphId          string     `dgraph:"uid"`
	DgraphType        []string   `dgraph:"dgraph.type"`
	DgraphCreatedDate *time.Time `dgraph:"_created_date @index(month)"`
	DgraphExpiredDate *time.Time `dgraph:"_expired_date @index(month)"`
}

func (x *ABTestMessage) Reset() {
	*x = ABTestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ABTestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ABTestMessage) ProtoMessage() {}

func (x *ABTestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ABTestMessage.ProtoReflect.Descriptor instead.
func (*ABTestMessage) Descriptor() ([]byte, []int) {
	return file_liulishuo_algorithm_algapi_abtest_abtest_proto_rawDescGZIP(), []int{9}
}

func (x *ABTestMessage) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *ABTestMessage) GetHashKey() string {
	if x != nil {
		return x.HashKey
	}
	return ""
}

func (x *ABTestMessage) GetStrategy() string {
	if x != nil {
		return x.Strategy
	}
	return ""
}

func (x *ABTestMessage) GetKeyType() string {
	if x != nil {
		return x.KeyType
	}
	return ""
}

func (x *ABTestMessage) GetLoginId() string {
	if x != nil {
		return x.LoginId
	}
	return ""
}

func (x *ABTestMessage) GetExt() string {
	if x != nil {
		return x.Ext
	}
	return ""
}

func (x *ABTestMessage) GetCreatedTime() string {
	if x != nil {
		return x.CreatedTime
	}
	return ""
}

var File_liulishuo_algorithm_algapi_abtest_abtest_proto protoreflect.FileDescriptor

var file_liulishuo_algorithm_algapi_abtest_abtest_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x6c, 0x69, 0x75, 0x6c, 0x69, 0x73, 0x68, 0x75, 0x6f, 0x2f, 0x61, 0x6c, 0x67, 0x6f,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x2f, 0x61, 0x6c, 0x67, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x62, 0x74,
	0x65, 0x73, 0x74, 0x2f, 0x61, 0x62, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0c, 0x61, 0x62, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x19,
	0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x71, 0x0a, 0x05, 0x4c, 0x61, 0x79,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x35, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x7a, 0x6f, 0x6e,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x62, 0x74, 0x65, 0x73,
	0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x0b, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0xf9, 0x01, 0x0a,
	0x04, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x62, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x62, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x05, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x12, 0x2c, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x61, 0x62, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73,
	0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x2c, 0x0a, 0x06, 0x57, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x22, 0x6b, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x42, 0x54,
	0x65, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x68, 0x61, 0x73, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x42, 0x54, 0x65, 0x73, 0x74,
	0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04,
	0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x62, 0x74,
	0x65, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x04,
	0x7a, 0x6f, 0x6e, 0x65, 0x22, 0x31, 0x0a, 0x07, 0x4c, 0x61, 0x62, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xb2, 0x01, 0x0a, 0x15, 0x50, 0x75, 0x73, 0x68,
	0x41, 0x42, 0x54, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x68, 0x61, 0x73, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x6b,
	0x65, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b,
	0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x62, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x61, 0x62, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x18, 0x0a, 0x16,
	0x50, 0x75, 0x73, 0x68, 0x41, 0x42, 0x54, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xd5, 0x01, 0x0a, 0x12, 0x41, 0x42, 0x54, 0x65, 0x73,
	0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x68, 0x61, 0x73, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x68, 0x61, 0x73, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x78, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xca,
	0x01, 0x0a, 0x0d, 0x41, 0x42, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x5a, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x42, 0x54, 0x65, 0x73, 0x74, 0x5a, 0x6f, 0x6e,
	0x65, 0x12, 0x22, 0x2e, 0x61, 0x62, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x42, 0x54, 0x65, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x62, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x42, 0x54, 0x65, 0x73, 0x74, 0x5a, 0x6f,
	0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x0e,
	0x50, 0x75, 0x73, 0x68, 0x41, 0x42, 0x54, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x23,
	0x2e, 0x61, 0x62, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75,
	0x73, 0x68, 0x41, 0x42, 0x54, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x62, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x41, 0x42, 0x54, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_liulishuo_algorithm_algapi_abtest_abtest_proto_rawDescOnce sync.Once
	file_liulishuo_algorithm_algapi_abtest_abtest_proto_rawDescData = file_liulishuo_algorithm_algapi_abtest_abtest_proto_rawDesc
)

func file_liulishuo_algorithm_algapi_abtest_abtest_proto_rawDescGZIP() []byte {
	file_liulishuo_algorithm_algapi_abtest_abtest_proto_rawDescOnce.Do(func() {
		file_liulishuo_algorithm_algapi_abtest_abtest_proto_rawDescData = protoimpl.X.CompressGZIP(file_liulishuo_algorithm_algapi_abtest_abtest_proto_rawDescData)
	})
	return file_liulishuo_algorithm_algapi_abtest_abtest_proto_rawDescData
}

var file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_liulishuo_algorithm_algapi_abtest_abtest_proto_goTypes = []interface{}{
	(*Project)(nil),                // 0: abtest_proto.Project
	(*Layer)(nil),                  // 1: abtest_proto.Layer
	(*Zone)(nil),                   // 2: abtest_proto.Zone
	(*Weight)(nil),                 // 3: abtest_proto.Weight
	(*GetABTestZoneRequest)(nil),   // 4: abtest_proto.GetABTestZoneRequest
	(*GetABTestZoneResponse)(nil),  // 5: abtest_proto.GetABTestZoneResponse
	(*LabData)(nil),                // 6: abtest_proto.LabData
	(*PushABTestDataRequest)(nil),  // 7: abtest_proto.PushABTestDataRequest
	(*PushABTestDataResponse)(nil), // 8: abtest_proto.PushABTestDataResponse
	(*ABTestMessage)(nil),          // 9: abtest_proto.ABTestMessage
}
var file_liulishuo_algorithm_algapi_abtest_abtest_proto_depIdxs = []int32{
	2, // 0: abtest_proto.Layer.parent_zones:type_name -> abtest_proto.Zone
	0, // 1: abtest_proto.Zone.project:type_name -> abtest_proto.Project
	1, // 2: abtest_proto.Zone.layer:type_name -> abtest_proto.Layer
	3, // 3: abtest_proto.Zone.weight:type_name -> abtest_proto.Weight
	2, // 4: abtest_proto.GetABTestZoneResponse.zone:type_name -> abtest_proto.Zone
	6, // 5: abtest_proto.PushABTestDataRequest.data:type_name -> abtest_proto.LabData
	4, // 6: abtest_proto.ABTestService.GetABTestZone:input_type -> abtest_proto.GetABTestZoneRequest
	7, // 7: abtest_proto.ABTestService.PushABTestData:input_type -> abtest_proto.PushABTestDataRequest
	5, // 8: abtest_proto.ABTestService.GetABTestZone:output_type -> abtest_proto.GetABTestZoneResponse
	8, // 9: abtest_proto.ABTestService.PushABTestData:output_type -> abtest_proto.PushABTestDataResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_liulishuo_algorithm_algapi_abtest_abtest_proto_init() }
func file_liulishuo_algorithm_algapi_abtest_abtest_proto_init() {
	if File_liulishuo_algorithm_algapi_abtest_abtest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
		file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Layer); i {
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
		file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Zone); i {
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
		file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Weight); i {
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
		file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetABTestZoneRequest); i {
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
		file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetABTestZoneResponse); i {
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
		file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabData); i {
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
		file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushABTestDataRequest); i {
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
		file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushABTestDataResponse); i {
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
		file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ABTestMessage); i {
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
			RawDescriptor: file_liulishuo_algorithm_algapi_abtest_abtest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_liulishuo_algorithm_algapi_abtest_abtest_proto_goTypes,
		DependencyIndexes: file_liulishuo_algorithm_algapi_abtest_abtest_proto_depIdxs,
		MessageInfos:      file_liulishuo_algorithm_algapi_abtest_abtest_proto_msgTypes,
	}.Build()
	File_liulishuo_algorithm_algapi_abtest_abtest_proto = out.File
	file_liulishuo_algorithm_algapi_abtest_abtest_proto_rawDesc = nil
	file_liulishuo_algorithm_algapi_abtest_abtest_proto_goTypes = nil
	file_liulishuo_algorithm_algapi_abtest_abtest_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ABTestServiceClient is the client API for ABTestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ABTestServiceClient interface {
	// GetABTestZone is to match the ABTest target Zone
	GetABTestZone(ctx context.Context, in *GetABTestZoneRequest, opts ...grpc.CallOption) (*GetABTestZoneResponse, error)
	// PushABTestData is to push abtest data to mq
	PushABTestData(ctx context.Context, in *PushABTestDataRequest, opts ...grpc.CallOption) (*PushABTestDataResponse, error)
}

type aBTestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewABTestServiceClient(cc grpc.ClientConnInterface) ABTestServiceClient {
	return &aBTestServiceClient{cc}
}

func (c *aBTestServiceClient) GetABTestZone(ctx context.Context, in *GetABTestZoneRequest, opts ...grpc.CallOption) (*GetABTestZoneResponse, error) {
	out := new(GetABTestZoneResponse)
	err := c.cc.Invoke(ctx, "/abtest_proto.ABTestService/GetABTestZone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBTestServiceClient) PushABTestData(ctx context.Context, in *PushABTestDataRequest, opts ...grpc.CallOption) (*PushABTestDataResponse, error) {
	out := new(PushABTestDataResponse)
	err := c.cc.Invoke(ctx, "/abtest_proto.ABTestService/PushABTestData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ABTestServiceServer is the server API for ABTestService service.
type ABTestServiceServer interface {
	// GetABTestZone is to match the ABTest target Zone
	GetABTestZone(context.Context, *GetABTestZoneRequest) (*GetABTestZoneResponse, error)
	// PushABTestData is to push abtest data to mq
	PushABTestData(context.Context, *PushABTestDataRequest) (*PushABTestDataResponse, error)
}

// UnimplementedABTestServiceServer can be embedded to have forward compatible implementations.
type UnimplementedABTestServiceServer struct {
}

func (*UnimplementedABTestServiceServer) GetABTestZone(context.Context, *GetABTestZoneRequest) (*GetABTestZoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetABTestZone not implemented")
}
func (*UnimplementedABTestServiceServer) PushABTestData(context.Context, *PushABTestDataRequest) (*PushABTestDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushABTestData not implemented")
}

func RegisterABTestServiceServer(s *grpc.Server, srv ABTestServiceServer) {
	s.RegisterService(&_ABTestService_serviceDesc, srv)
}

func _ABTestService_GetABTestZone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetABTestZoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABTestServiceServer).GetABTestZone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abtest_proto.ABTestService/GetABTestZone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABTestServiceServer).GetABTestZone(ctx, req.(*GetABTestZoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABTestService_PushABTestData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushABTestDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABTestServiceServer).PushABTestData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abtest_proto.ABTestService/PushABTestData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABTestServiceServer).PushABTestData(ctx, req.(*PushABTestDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ABTestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "abtest_proto.ABTestService",
	HandlerType: (*ABTestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetABTestZone",
			Handler:    _ABTestService_GetABTestZone_Handler,
		},
		{
			MethodName: "PushABTestData",
			Handler:    _ABTestService_PushABTestData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "liulishuo/algorithm/algapi/abtest/abtest.proto",
}
