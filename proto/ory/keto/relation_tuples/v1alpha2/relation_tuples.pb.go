// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: ory/keto/relation_tuples/v1alpha2/relation_tuples.proto

package rts

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// RelationTuple defines a relation between an Object and a Subject.
type RelationTuple struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The namespace this relation tuple lives in.
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// The object related by this tuple.
	// It is an object in the namespace of the tuple.
	Object string `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
	// The relation between an Object and a Subject.
	Relation string `protobuf:"bytes,3,opt,name=relation,proto3" json:"relation,omitempty"`
	// The subject related by this tuple.
	// A Subject either represents a concrete subject id or
	// a `SubjectSet` that expands to more Subjects.
	Subject       *Subject `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RelationTuple) Reset() {
	*x = RelationTuple{}
	mi := &file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RelationTuple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationTuple) ProtoMessage() {}

func (x *RelationTuple) ProtoReflect() protoreflect.Message {
	mi := &file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationTuple.ProtoReflect.Descriptor instead.
func (*RelationTuple) Descriptor() ([]byte, []int) {
	return file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_rawDescGZIP(), []int{0}
}

func (x *RelationTuple) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *RelationTuple) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *RelationTuple) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

func (x *RelationTuple) GetSubject() *Subject {
	if x != nil {
		return x.Subject
	}
	return nil
}

// The query for listing relationships.
// Clients can specify any optional field to
// partially filter for specific relationships.
//
// Example use cases (namespace is always required):
//   - object only: display a list of all permissions referring to a specific object
//   - relation only: get all groups that have members; get all directories that have content
//   - object & relation: display all subjects that have a specific permission relation
//   - subject & relation: display all groups a subject belongs to; display all objects a subject has access to
//   - object & relation & subject: check whether the relation tuple already exists
type RelationQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The namespace this relation tuple lives in.
	Namespace *string `protobuf:"bytes,1,opt,name=namespace,proto3,oneof" json:"namespace,omitempty"`
	// The object related by this tuple.
	// It is an object in the namespace of the tuple.
	Object *string `protobuf:"bytes,2,opt,name=object,proto3,oneof" json:"object,omitempty"`
	// The relation between an Object and a Subject.
	Relation *string `protobuf:"bytes,3,opt,name=relation,proto3,oneof" json:"relation,omitempty"`
	// The subject related by this tuple.
	// A Subject either represents a concrete subject id or
	// a `SubjectSet` that expands to more Subjects.
	Subject       *Subject `protobuf:"bytes,4,opt,name=subject,proto3,oneof" json:"subject,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RelationQuery) Reset() {
	*x = RelationQuery{}
	mi := &file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RelationQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationQuery) ProtoMessage() {}

func (x *RelationQuery) ProtoReflect() protoreflect.Message {
	mi := &file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationQuery.ProtoReflect.Descriptor instead.
func (*RelationQuery) Descriptor() ([]byte, []int) {
	return file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_rawDescGZIP(), []int{1}
}

func (x *RelationQuery) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

func (x *RelationQuery) GetObject() string {
	if x != nil && x.Object != nil {
		return *x.Object
	}
	return ""
}

func (x *RelationQuery) GetRelation() string {
	if x != nil && x.Relation != nil {
		return *x.Relation
	}
	return ""
}

func (x *RelationQuery) GetSubject() *Subject {
	if x != nil {
		return x.Subject
	}
	return nil
}

// Subject is either a concrete subject id or
// a `SubjectSet` expanding to more Subjects.
type Subject struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The reference of this abstract subject.
	//
	// Types that are valid to be assigned to Ref:
	//
	//	*Subject_Id
	//	*Subject_Set
	Ref           isSubject_Ref `protobuf_oneof:"ref"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Subject) Reset() {
	*x = Subject{}
	mi := &file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Subject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subject) ProtoMessage() {}

func (x *Subject) ProtoReflect() protoreflect.Message {
	mi := &file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subject.ProtoReflect.Descriptor instead.
func (*Subject) Descriptor() ([]byte, []int) {
	return file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_rawDescGZIP(), []int{2}
}

func (x *Subject) GetRef() isSubject_Ref {
	if x != nil {
		return x.Ref
	}
	return nil
}

func (x *Subject) GetId() string {
	if x != nil {
		if x, ok := x.Ref.(*Subject_Id); ok {
			return x.Id
		}
	}
	return ""
}

func (x *Subject) GetSet() *SubjectSet {
	if x != nil {
		if x, ok := x.Ref.(*Subject_Set); ok {
			return x.Set
		}
	}
	return nil
}

type isSubject_Ref interface {
	isSubject_Ref()
}

type Subject_Id struct {
	// A concrete id of the subject.
	Id string `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type Subject_Set struct {
	// A subject set that expands to more Subjects.
	// More information are available under [concepts](../concepts/15_subjects.mdx).
	Set *SubjectSet `protobuf:"bytes,2,opt,name=set,proto3,oneof"`
}

func (*Subject_Id) isSubject_Ref() {}

func (*Subject_Set) isSubject_Ref() {}

// SubjectSet refers to all subjects who have
// the same `relation` on an `object`.
type SubjectSet struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The namespace of the object and relation
	// referenced in this subject set.
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// The object related by this subject set.
	Object string `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
	// The relation between the object and the subjects.
	Relation      string `protobuf:"bytes,3,opt,name=relation,proto3" json:"relation,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubjectSet) Reset() {
	*x = SubjectSet{}
	mi := &file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubjectSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubjectSet) ProtoMessage() {}

func (x *SubjectSet) ProtoReflect() protoreflect.Message {
	mi := &file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubjectSet.ProtoReflect.Descriptor instead.
func (*SubjectSet) Descriptor() ([]byte, []int) {
	return file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_rawDescGZIP(), []int{3}
}

func (x *SubjectSet) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *SubjectSet) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *SubjectSet) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

var File_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto protoreflect.FileDescriptor

const file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_rawDesc = "" +
	"\n" +
	"7ory/keto/relation_tuples/v1alpha2/relation_tuples.proto\x12!ory.keto.relation_tuples.v1alpha2\"\xa7\x01\n" +
	"\rRelationTuple\x12\x1c\n" +
	"\tnamespace\x18\x01 \x01(\tR\tnamespace\x12\x16\n" +
	"\x06object\x18\x02 \x01(\tR\x06object\x12\x1a\n" +
	"\brelation\x18\x03 \x01(\tR\brelation\x12D\n" +
	"\asubject\x18\x04 \x01(\v2*.ory.keto.relation_tuples.v1alpha2.SubjectR\asubject\"\xed\x01\n" +
	"\rRelationQuery\x12!\n" +
	"\tnamespace\x18\x01 \x01(\tH\x00R\tnamespace\x88\x01\x01\x12\x1b\n" +
	"\x06object\x18\x02 \x01(\tH\x01R\x06object\x88\x01\x01\x12\x1f\n" +
	"\brelation\x18\x03 \x01(\tH\x02R\brelation\x88\x01\x01\x12I\n" +
	"\asubject\x18\x04 \x01(\v2*.ory.keto.relation_tuples.v1alpha2.SubjectH\x03R\asubject\x88\x01\x01B\f\n" +
	"\n" +
	"_namespaceB\t\n" +
	"\a_objectB\v\n" +
	"\t_relationB\n" +
	"\n" +
	"\b_subject\"e\n" +
	"\aSubject\x12\x10\n" +
	"\x02id\x18\x01 \x01(\tH\x00R\x02id\x12A\n" +
	"\x03set\x18\x02 \x01(\v2-.ory.keto.relation_tuples.v1alpha2.SubjectSetH\x00R\x03setB\x05\n" +
	"\x03ref\"^\n" +
	"\n" +
	"SubjectSet\x12\x1c\n" +
	"\tnamespace\x18\x01 \x01(\tR\tnamespace\x12\x16\n" +
	"\x06object\x18\x02 \x01(\tR\x06object\x12\x1a\n" +
	"\brelation\x18\x03 \x01(\tR\brelationB\xc4\x01\n" +
	"$sh.ory.keto.relation_tuples.v1alpha2B\x13RelationTuplesProtoP\x01Z?github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2;rts\xaa\x02 Ory.Keto.RelationTuples.v1alpha2\xca\x02 Ory\\Keto\\RelationTuples\\v1alpha2b\x06proto3"

var (
	file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_rawDescOnce sync.Once
	file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_rawDescData []byte
)

func file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_rawDescGZIP() []byte {
	file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_rawDescOnce.Do(func() {
		file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_rawDesc), len(file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_rawDesc)))
	})
	return file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_rawDescData
}

var file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_goTypes = []any{
	(*RelationTuple)(nil), // 0: ory.keto.relation_tuples.v1alpha2.RelationTuple
	(*RelationQuery)(nil), // 1: ory.keto.relation_tuples.v1alpha2.RelationQuery
	(*Subject)(nil),       // 2: ory.keto.relation_tuples.v1alpha2.Subject
	(*SubjectSet)(nil),    // 3: ory.keto.relation_tuples.v1alpha2.SubjectSet
}
var file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_depIdxs = []int32{
	2, // 0: ory.keto.relation_tuples.v1alpha2.RelationTuple.subject:type_name -> ory.keto.relation_tuples.v1alpha2.Subject
	2, // 1: ory.keto.relation_tuples.v1alpha2.RelationQuery.subject:type_name -> ory.keto.relation_tuples.v1alpha2.Subject
	3, // 2: ory.keto.relation_tuples.v1alpha2.Subject.set:type_name -> ory.keto.relation_tuples.v1alpha2.SubjectSet
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_init() }
func file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_init() {
	if File_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto != nil {
		return
	}
	file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_msgTypes[1].OneofWrappers = []any{}
	file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_msgTypes[2].OneofWrappers = []any{
		(*Subject_Id)(nil),
		(*Subject_Set)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_rawDesc), len(file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_goTypes,
		DependencyIndexes: file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_depIdxs,
		MessageInfos:      file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_msgTypes,
	}.Build()
	File_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto = out.File
	file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_goTypes = nil
	file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_depIdxs = nil
}
