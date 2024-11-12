/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// DeleteMonitoredItemsRequest is the corresponding interface of DeleteMonitoredItemsRequest
type DeleteMonitoredItemsRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() RequestHeader
	// GetSubscriptionId returns SubscriptionId (property field)
	GetSubscriptionId() uint32
	// GetMonitoredItemIds returns MonitoredItemIds (property field)
	GetMonitoredItemIds() []uint32
	// IsDeleteMonitoredItemsRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDeleteMonitoredItemsRequest()
	// CreateBuilder creates a DeleteMonitoredItemsRequestBuilder
	CreateDeleteMonitoredItemsRequestBuilder() DeleteMonitoredItemsRequestBuilder
}

// _DeleteMonitoredItemsRequest is the data-structure of this message
type _DeleteMonitoredItemsRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader    RequestHeader
	SubscriptionId   uint32
	MonitoredItemIds []uint32
}

var _ DeleteMonitoredItemsRequest = (*_DeleteMonitoredItemsRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_DeleteMonitoredItemsRequest)(nil)

// NewDeleteMonitoredItemsRequest factory function for _DeleteMonitoredItemsRequest
func NewDeleteMonitoredItemsRequest(requestHeader RequestHeader, subscriptionId uint32, monitoredItemIds []uint32) *_DeleteMonitoredItemsRequest {
	if requestHeader == nil {
		panic("requestHeader of type RequestHeader for DeleteMonitoredItemsRequest must not be nil")
	}
	_result := &_DeleteMonitoredItemsRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		SubscriptionId:                    subscriptionId,
		MonitoredItemIds:                  monitoredItemIds,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DeleteMonitoredItemsRequestBuilder is a builder for DeleteMonitoredItemsRequest
type DeleteMonitoredItemsRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestHeader RequestHeader, subscriptionId uint32, monitoredItemIds []uint32) DeleteMonitoredItemsRequestBuilder
	// WithRequestHeader adds RequestHeader (property field)
	WithRequestHeader(RequestHeader) DeleteMonitoredItemsRequestBuilder
	// WithRequestHeaderBuilder adds RequestHeader (property field) which is build by the builder
	WithRequestHeaderBuilder(func(RequestHeaderBuilder) RequestHeaderBuilder) DeleteMonitoredItemsRequestBuilder
	// WithSubscriptionId adds SubscriptionId (property field)
	WithSubscriptionId(uint32) DeleteMonitoredItemsRequestBuilder
	// WithMonitoredItemIds adds MonitoredItemIds (property field)
	WithMonitoredItemIds(...uint32) DeleteMonitoredItemsRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the DeleteMonitoredItemsRequest or returns an error if something is wrong
	Build() (DeleteMonitoredItemsRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DeleteMonitoredItemsRequest
}

// NewDeleteMonitoredItemsRequestBuilder() creates a DeleteMonitoredItemsRequestBuilder
func NewDeleteMonitoredItemsRequestBuilder() DeleteMonitoredItemsRequestBuilder {
	return &_DeleteMonitoredItemsRequestBuilder{_DeleteMonitoredItemsRequest: new(_DeleteMonitoredItemsRequest)}
}

type _DeleteMonitoredItemsRequestBuilder struct {
	*_DeleteMonitoredItemsRequest

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (DeleteMonitoredItemsRequestBuilder) = (*_DeleteMonitoredItemsRequestBuilder)(nil)

func (b *_DeleteMonitoredItemsRequestBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_DeleteMonitoredItemsRequestBuilder) WithMandatoryFields(requestHeader RequestHeader, subscriptionId uint32, monitoredItemIds []uint32) DeleteMonitoredItemsRequestBuilder {
	return b.WithRequestHeader(requestHeader).WithSubscriptionId(subscriptionId).WithMonitoredItemIds(monitoredItemIds...)
}

func (b *_DeleteMonitoredItemsRequestBuilder) WithRequestHeader(requestHeader RequestHeader) DeleteMonitoredItemsRequestBuilder {
	b.RequestHeader = requestHeader
	return b
}

func (b *_DeleteMonitoredItemsRequestBuilder) WithRequestHeaderBuilder(builderSupplier func(RequestHeaderBuilder) RequestHeaderBuilder) DeleteMonitoredItemsRequestBuilder {
	builder := builderSupplier(b.RequestHeader.CreateRequestHeaderBuilder())
	var err error
	b.RequestHeader, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "RequestHeaderBuilder failed"))
	}
	return b
}

func (b *_DeleteMonitoredItemsRequestBuilder) WithSubscriptionId(subscriptionId uint32) DeleteMonitoredItemsRequestBuilder {
	b.SubscriptionId = subscriptionId
	return b
}

func (b *_DeleteMonitoredItemsRequestBuilder) WithMonitoredItemIds(monitoredItemIds ...uint32) DeleteMonitoredItemsRequestBuilder {
	b.MonitoredItemIds = monitoredItemIds
	return b
}

func (b *_DeleteMonitoredItemsRequestBuilder) Build() (DeleteMonitoredItemsRequest, error) {
	if b.RequestHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'requestHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._DeleteMonitoredItemsRequest.deepCopy(), nil
}

func (b *_DeleteMonitoredItemsRequestBuilder) MustBuild() DeleteMonitoredItemsRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_DeleteMonitoredItemsRequestBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_DeleteMonitoredItemsRequestBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_DeleteMonitoredItemsRequestBuilder) DeepCopy() any {
	_copy := b.CreateDeleteMonitoredItemsRequestBuilder().(*_DeleteMonitoredItemsRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDeleteMonitoredItemsRequestBuilder creates a DeleteMonitoredItemsRequestBuilder
func (b *_DeleteMonitoredItemsRequest) CreateDeleteMonitoredItemsRequestBuilder() DeleteMonitoredItemsRequestBuilder {
	if b == nil {
		return NewDeleteMonitoredItemsRequestBuilder()
	}
	return &_DeleteMonitoredItemsRequestBuilder{_DeleteMonitoredItemsRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DeleteMonitoredItemsRequest) GetExtensionId() int32 {
	return int32(781)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DeleteMonitoredItemsRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeleteMonitoredItemsRequest) GetRequestHeader() RequestHeader {
	return m.RequestHeader
}

func (m *_DeleteMonitoredItemsRequest) GetSubscriptionId() uint32 {
	return m.SubscriptionId
}

func (m *_DeleteMonitoredItemsRequest) GetMonitoredItemIds() []uint32 {
	return m.MonitoredItemIds
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDeleteMonitoredItemsRequest(structType any) DeleteMonitoredItemsRequest {
	if casted, ok := structType.(DeleteMonitoredItemsRequest); ok {
		return casted
	}
	if casted, ok := structType.(*DeleteMonitoredItemsRequest); ok {
		return *casted
	}
	return nil
}

func (m *_DeleteMonitoredItemsRequest) GetTypeName() string {
	return "DeleteMonitoredItemsRequest"
}

func (m *_DeleteMonitoredItemsRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (subscriptionId)
	lengthInBits += 32

	// Implicit Field (noOfMonitoredItemIds)
	lengthInBits += 32

	// Array field
	if len(m.MonitoredItemIds) > 0 {
		lengthInBits += 32 * uint16(len(m.MonitoredItemIds))
	}

	return lengthInBits
}

func (m *_DeleteMonitoredItemsRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DeleteMonitoredItemsRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__deleteMonitoredItemsRequest DeleteMonitoredItemsRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DeleteMonitoredItemsRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeleteMonitoredItemsRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[RequestHeader](ctx, "requestHeader", ReadComplex[RequestHeader](ExtensionObjectDefinitionParseWithBufferProducer[RequestHeader]((int32)(int32(391))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	subscriptionId, err := ReadSimpleField(ctx, "subscriptionId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subscriptionId' field"))
	}
	m.SubscriptionId = subscriptionId

	noOfMonitoredItemIds, err := ReadImplicitField[int32](ctx, "noOfMonitoredItemIds", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfMonitoredItemIds' field"))
	}
	_ = noOfMonitoredItemIds

	monitoredItemIds, err := ReadCountArrayField[uint32](ctx, "monitoredItemIds", ReadUnsignedInt(readBuffer, uint8(32)), uint64(noOfMonitoredItemIds))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'monitoredItemIds' field"))
	}
	m.MonitoredItemIds = monitoredItemIds

	if closeErr := readBuffer.CloseContext("DeleteMonitoredItemsRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeleteMonitoredItemsRequest")
	}

	return m, nil
}

func (m *_DeleteMonitoredItemsRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DeleteMonitoredItemsRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DeleteMonitoredItemsRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DeleteMonitoredItemsRequest")
		}

		if err := WriteSimpleField[RequestHeader](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[RequestHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[uint32](ctx, "subscriptionId", m.GetSubscriptionId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'subscriptionId' field")
		}
		noOfMonitoredItemIds := int32(utils.InlineIf(bool((m.GetMonitoredItemIds()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetMonitoredItemIds()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfMonitoredItemIds", noOfMonitoredItemIds, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfMonitoredItemIds' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "monitoredItemIds", m.GetMonitoredItemIds(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'monitoredItemIds' field")
		}

		if popErr := writeBuffer.PopContext("DeleteMonitoredItemsRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DeleteMonitoredItemsRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DeleteMonitoredItemsRequest) IsDeleteMonitoredItemsRequest() {}

func (m *_DeleteMonitoredItemsRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DeleteMonitoredItemsRequest) deepCopy() *_DeleteMonitoredItemsRequest {
	if m == nil {
		return nil
	}
	_DeleteMonitoredItemsRequestCopy := &_DeleteMonitoredItemsRequest{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[RequestHeader](m.RequestHeader),
		m.SubscriptionId,
		utils.DeepCopySlice[uint32, uint32](m.MonitoredItemIds),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _DeleteMonitoredItemsRequestCopy
}

func (m *_DeleteMonitoredItemsRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
