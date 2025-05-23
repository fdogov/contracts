// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: backend/partnerconsumer/kafka/v1/order.proto

package kafkav1

import (
	_type "github.com/fdogov/contracts/gen/go/google/type"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrderStatus int32

const (
	OrderStatus_ORDER_STATUS_UNSPECIFIED OrderStatus = 0
	OrderStatus_ORDER_STATUS_PENDING     OrderStatus = 1
	OrderStatus_ORDER_STATUS_COMPLETED   OrderStatus = 2
	OrderStatus_ORDER_STATUS_FAILED      OrderStatus = 3
	OrderStatus_ORDER_STATUS_CANCELLED   OrderStatus = 4
)

// Enum value maps for OrderStatus.
var (
	OrderStatus_name = map[int32]string{
		0: "ORDER_STATUS_UNSPECIFIED",
		1: "ORDER_STATUS_PENDING",
		2: "ORDER_STATUS_COMPLETED",
		3: "ORDER_STATUS_FAILED",
		4: "ORDER_STATUS_CANCELLED",
	}
	OrderStatus_value = map[string]int32{
		"ORDER_STATUS_UNSPECIFIED": 0,
		"ORDER_STATUS_PENDING":     1,
		"ORDER_STATUS_COMPLETED":   2,
		"ORDER_STATUS_FAILED":      3,
		"ORDER_STATUS_CANCELLED":   4,
	}
)

func (x OrderStatus) Enum() *OrderStatus {
	p := new(OrderStatus)
	*p = x
	return p
}

func (x OrderStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_backend_partnerconsumer_kafka_v1_order_proto_enumTypes[0].Descriptor()
}

func (OrderStatus) Type() protoreflect.EnumType {
	return &file_backend_partnerconsumer_kafka_v1_order_proto_enumTypes[0]
}

func (x OrderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderStatus.Descriptor instead.
func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return file_backend_partnerconsumer_kafka_v1_order_proto_rawDescGZIP(), []int{0}
}

type OrderEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExtId          string                 `protobuf:"bytes,1,opt,name=ext_id,json=extId,proto3" json:"ext_id,omitempty"`
	ExtAccountId   string                 `protobuf:"bytes,2,opt,name=ext_account_id,json=extAccountId,proto3" json:"ext_account_id,omitempty"`
	Amount         *_type.Decimal         `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency       string                 `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Status         OrderStatus            `protobuf:"varint,6,opt,name=status,proto3,enum=backend.partnerconsumer.kafka.v1.OrderStatus" json:"status,omitempty"`
	Symbol         string                 `protobuf:"bytes,7,opt,name=symbol,proto3" json:"symbol,omitempty"`
	InstrumentId   string                 `protobuf:"bytes,8,opt,name=instrument_id,json=instrumentId,proto3" json:"instrument_id,omitempty"`
	BalanceNew     *_type.Decimal         `protobuf:"bytes,9,opt,name=balance_new,json=balanceNew,proto3" json:"balance_new,omitempty"`
	IdempotencyKey string                 `protobuf:"bytes,10,opt,name=idempotency_key,json=idempotencyKey,proto3" json:"idempotency_key,omitempty"`
}

func (x *OrderEvent) Reset() {
	*x = OrderEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_partnerconsumer_kafka_v1_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderEvent) ProtoMessage() {}

func (x *OrderEvent) ProtoReflect() protoreflect.Message {
	mi := &file_backend_partnerconsumer_kafka_v1_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderEvent.ProtoReflect.Descriptor instead.
func (*OrderEvent) Descriptor() ([]byte, []int) {
	return file_backend_partnerconsumer_kafka_v1_order_proto_rawDescGZIP(), []int{0}
}

func (x *OrderEvent) GetExtId() string {
	if x != nil {
		return x.ExtId
	}
	return ""
}

func (x *OrderEvent) GetExtAccountId() string {
	if x != nil {
		return x.ExtAccountId
	}
	return ""
}

func (x *OrderEvent) GetAmount() *_type.Decimal {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *OrderEvent) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *OrderEvent) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *OrderEvent) GetStatus() OrderStatus {
	if x != nil {
		return x.Status
	}
	return OrderStatus_ORDER_STATUS_UNSPECIFIED
}

func (x *OrderEvent) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *OrderEvent) GetInstrumentId() string {
	if x != nil {
		return x.InstrumentId
	}
	return ""
}

func (x *OrderEvent) GetBalanceNew() *_type.Decimal {
	if x != nil {
		return x.BalanceNew
	}
	return nil
}

func (x *OrderEvent) GetIdempotencyKey() string {
	if x != nil {
		return x.IdempotencyKey
	}
	return ""
}

var File_backend_partnerconsumer_kafka_v1_order_proto protoreflect.FileDescriptor

var file_backend_partnerconsumer_kafka_v1_order_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2f,
	0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64,
	0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x03, 0x0a,
	0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x65,
	0x78, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x78, 0x74,
	0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x65, 0x78, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x78, 0x74, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x45, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x23, 0x0a, 0x0d,
	0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x35, 0x0a, 0x0b, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x65, 0x77,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x0a, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x65, 0x77, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x64, 0x65, 0x6d,
	0x70, 0x6f, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x69, 0x64, 0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x4b, 0x65,
	0x79, 0x2a, 0x96, 0x01, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1c, 0x0a, 0x18, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x18, 0x0a, 0x14, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x4f, 0x52, 0x44,
	0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45,
	0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x1a,
	0x0a, 0x16, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43,
	0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x42, 0xa2, 0x02, 0x0a, 0x24, 0x63,
	0x6f, 0x6d, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61,
	0x2e, 0x76, 0x31, 0x42, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x64,
	0x6f, 0x67, 0x6f, 0x76, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2f, 0x6b, 0x61,
	0x66, 0x6b, 0x61, 0x2f, 0x76, 0x31, 0x3b, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x42, 0x50, 0x4b, 0xaa, 0x02, 0x20, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x50,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x4b,
	0x61, 0x66, 0x6b, 0x61, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x20, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x5c, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x5c, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x2c, 0x42, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x5c, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x72, 0x5c, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x23, 0x42, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x3a, 0x3a, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x72, 0x3a, 0x3a, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_backend_partnerconsumer_kafka_v1_order_proto_rawDescOnce sync.Once
	file_backend_partnerconsumer_kafka_v1_order_proto_rawDescData = file_backend_partnerconsumer_kafka_v1_order_proto_rawDesc
)

func file_backend_partnerconsumer_kafka_v1_order_proto_rawDescGZIP() []byte {
	file_backend_partnerconsumer_kafka_v1_order_proto_rawDescOnce.Do(func() {
		file_backend_partnerconsumer_kafka_v1_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_backend_partnerconsumer_kafka_v1_order_proto_rawDescData)
	})
	return file_backend_partnerconsumer_kafka_v1_order_proto_rawDescData
}

var file_backend_partnerconsumer_kafka_v1_order_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_backend_partnerconsumer_kafka_v1_order_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_backend_partnerconsumer_kafka_v1_order_proto_goTypes = []interface{}{
	(OrderStatus)(0),              // 0: backend.partnerconsumer.kafka.v1.OrderStatus
	(*OrderEvent)(nil),            // 1: backend.partnerconsumer.kafka.v1.OrderEvent
	(*_type.Decimal)(nil),         // 2: google.type.Decimal
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_backend_partnerconsumer_kafka_v1_order_proto_depIdxs = []int32{
	2, // 0: backend.partnerconsumer.kafka.v1.OrderEvent.amount:type_name -> google.type.Decimal
	3, // 1: backend.partnerconsumer.kafka.v1.OrderEvent.created_at:type_name -> google.protobuf.Timestamp
	0, // 2: backend.partnerconsumer.kafka.v1.OrderEvent.status:type_name -> backend.partnerconsumer.kafka.v1.OrderStatus
	2, // 3: backend.partnerconsumer.kafka.v1.OrderEvent.balance_new:type_name -> google.type.Decimal
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_backend_partnerconsumer_kafka_v1_order_proto_init() }
func file_backend_partnerconsumer_kafka_v1_order_proto_init() {
	if File_backend_partnerconsumer_kafka_v1_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_backend_partnerconsumer_kafka_v1_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderEvent); i {
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
			RawDescriptor: file_backend_partnerconsumer_kafka_v1_order_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_backend_partnerconsumer_kafka_v1_order_proto_goTypes,
		DependencyIndexes: file_backend_partnerconsumer_kafka_v1_order_proto_depIdxs,
		EnumInfos:         file_backend_partnerconsumer_kafka_v1_order_proto_enumTypes,
		MessageInfos:      file_backend_partnerconsumer_kafka_v1_order_proto_msgTypes,
	}.Build()
	File_backend_partnerconsumer_kafka_v1_order_proto = out.File
	file_backend_partnerconsumer_kafka_v1_order_proto_rawDesc = nil
	file_backend_partnerconsumer_kafka_v1_order_proto_goTypes = nil
	file_backend_partnerconsumer_kafka_v1_order_proto_depIdxs = nil
}
