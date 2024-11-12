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

// ModifyMonitoredItemsRequest is the corresponding interface of ModifyMonitoredItemsRequest
type ModifyMonitoredItemsRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() RequestHeader
	// GetSubscriptionId returns SubscriptionId (property field)
	GetSubscriptionId() uint32
	// GetTimestampsToReturn returns TimestampsToReturn (property field)
	GetTimestampsToReturn() TimestampsToReturn
	// GetItemsToModify returns ItemsToModify (property field)
	GetItemsToModify() []MonitoredItemModifyRequest
	// IsModifyMonitoredItemsRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModifyMonitoredItemsRequest()
	// CreateBuilder creates a ModifyMonitoredItemsRequestBuilder
	CreateModifyMonitoredItemsRequestBuilder() ModifyMonitoredItemsRequestBuilder
}

// _ModifyMonitoredItemsRequest is the data-structure of this message
type _ModifyMonitoredItemsRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader      RequestHeader
	SubscriptionId     uint32
	TimestampsToReturn TimestampsToReturn
	ItemsToModify      []MonitoredItemModifyRequest
}

var _ ModifyMonitoredItemsRequest = (*_ModifyMonitoredItemsRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ModifyMonitoredItemsRequest)(nil)

// NewModifyMonitoredItemsRequest factory function for _ModifyMonitoredItemsRequest
func NewModifyMonitoredItemsRequest(requestHeader RequestHeader, subscriptionId uint32, timestampsToReturn TimestampsToReturn, itemsToModify []MonitoredItemModifyRequest) *_ModifyMonitoredItemsRequest {
	if requestHeader == nil {
		panic("requestHeader of type RequestHeader for ModifyMonitoredItemsRequest must not be nil")
	}
	_result := &_ModifyMonitoredItemsRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		SubscriptionId:                    subscriptionId,
		TimestampsToReturn:                timestampsToReturn,
		ItemsToModify:                     itemsToModify,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ModifyMonitoredItemsRequestBuilder is a builder for ModifyMonitoredItemsRequest
type ModifyMonitoredItemsRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestHeader RequestHeader, subscriptionId uint32, timestampsToReturn TimestampsToReturn, itemsToModify []MonitoredItemModifyRequest) ModifyMonitoredItemsRequestBuilder
	// WithRequestHeader adds RequestHeader (property field)
	WithRequestHeader(RequestHeader) ModifyMonitoredItemsRequestBuilder
	// WithRequestHeaderBuilder adds RequestHeader (property field) which is build by the builder
	WithRequestHeaderBuilder(func(RequestHeaderBuilder) RequestHeaderBuilder) ModifyMonitoredItemsRequestBuilder
	// WithSubscriptionId adds SubscriptionId (property field)
	WithSubscriptionId(uint32) ModifyMonitoredItemsRequestBuilder
	// WithTimestampsToReturn adds TimestampsToReturn (property field)
	WithTimestampsToReturn(TimestampsToReturn) ModifyMonitoredItemsRequestBuilder
	// WithItemsToModify adds ItemsToModify (property field)
	WithItemsToModify(...MonitoredItemModifyRequest) ModifyMonitoredItemsRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the ModifyMonitoredItemsRequest or returns an error if something is wrong
	Build() (ModifyMonitoredItemsRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ModifyMonitoredItemsRequest
}

// NewModifyMonitoredItemsRequestBuilder() creates a ModifyMonitoredItemsRequestBuilder
func NewModifyMonitoredItemsRequestBuilder() ModifyMonitoredItemsRequestBuilder {
	return &_ModifyMonitoredItemsRequestBuilder{_ModifyMonitoredItemsRequest: new(_ModifyMonitoredItemsRequest)}
}

type _ModifyMonitoredItemsRequestBuilder struct {
	*_ModifyMonitoredItemsRequest

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (ModifyMonitoredItemsRequestBuilder) = (*_ModifyMonitoredItemsRequestBuilder)(nil)

func (b *_ModifyMonitoredItemsRequestBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_ModifyMonitoredItemsRequestBuilder) WithMandatoryFields(requestHeader RequestHeader, subscriptionId uint32, timestampsToReturn TimestampsToReturn, itemsToModify []MonitoredItemModifyRequest) ModifyMonitoredItemsRequestBuilder {
	return b.WithRequestHeader(requestHeader).WithSubscriptionId(subscriptionId).WithTimestampsToReturn(timestampsToReturn).WithItemsToModify(itemsToModify...)
}

func (b *_ModifyMonitoredItemsRequestBuilder) WithRequestHeader(requestHeader RequestHeader) ModifyMonitoredItemsRequestBuilder {
	b.RequestHeader = requestHeader
	return b
}

func (b *_ModifyMonitoredItemsRequestBuilder) WithRequestHeaderBuilder(builderSupplier func(RequestHeaderBuilder) RequestHeaderBuilder) ModifyMonitoredItemsRequestBuilder {
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

func (b *_ModifyMonitoredItemsRequestBuilder) WithSubscriptionId(subscriptionId uint32) ModifyMonitoredItemsRequestBuilder {
	b.SubscriptionId = subscriptionId
	return b
}

func (b *_ModifyMonitoredItemsRequestBuilder) WithTimestampsToReturn(timestampsToReturn TimestampsToReturn) ModifyMonitoredItemsRequestBuilder {
	b.TimestampsToReturn = timestampsToReturn
	return b
}

func (b *_ModifyMonitoredItemsRequestBuilder) WithItemsToModify(itemsToModify ...MonitoredItemModifyRequest) ModifyMonitoredItemsRequestBuilder {
	b.ItemsToModify = itemsToModify
	return b
}

func (b *_ModifyMonitoredItemsRequestBuilder) Build() (ModifyMonitoredItemsRequest, error) {
	if b.RequestHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'requestHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ModifyMonitoredItemsRequest.deepCopy(), nil
}

func (b *_ModifyMonitoredItemsRequestBuilder) MustBuild() ModifyMonitoredItemsRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ModifyMonitoredItemsRequestBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_ModifyMonitoredItemsRequestBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_ModifyMonitoredItemsRequestBuilder) DeepCopy() any {
	_copy := b.CreateModifyMonitoredItemsRequestBuilder().(*_ModifyMonitoredItemsRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateModifyMonitoredItemsRequestBuilder creates a ModifyMonitoredItemsRequestBuilder
func (b *_ModifyMonitoredItemsRequest) CreateModifyMonitoredItemsRequestBuilder() ModifyMonitoredItemsRequestBuilder {
	if b == nil {
		return NewModifyMonitoredItemsRequestBuilder()
	}
	return &_ModifyMonitoredItemsRequestBuilder{_ModifyMonitoredItemsRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModifyMonitoredItemsRequest) GetExtensionId() int32 {
	return int32(763)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModifyMonitoredItemsRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModifyMonitoredItemsRequest) GetRequestHeader() RequestHeader {
	return m.RequestHeader
}

func (m *_ModifyMonitoredItemsRequest) GetSubscriptionId() uint32 {
	return m.SubscriptionId
}

func (m *_ModifyMonitoredItemsRequest) GetTimestampsToReturn() TimestampsToReturn {
	return m.TimestampsToReturn
}

func (m *_ModifyMonitoredItemsRequest) GetItemsToModify() []MonitoredItemModifyRequest {
	return m.ItemsToModify
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastModifyMonitoredItemsRequest(structType any) ModifyMonitoredItemsRequest {
	if casted, ok := structType.(ModifyMonitoredItemsRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModifyMonitoredItemsRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModifyMonitoredItemsRequest) GetTypeName() string {
	return "ModifyMonitoredItemsRequest"
}

func (m *_ModifyMonitoredItemsRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (subscriptionId)
	lengthInBits += 32

	// Simple field (timestampsToReturn)
	lengthInBits += 32

	// Implicit Field (noOfItemsToModify)
	lengthInBits += 32

	// Array field
	if len(m.ItemsToModify) > 0 {
		for _curItem, element := range m.ItemsToModify {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ItemsToModify), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_ModifyMonitoredItemsRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModifyMonitoredItemsRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__modifyMonitoredItemsRequest ModifyMonitoredItemsRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModifyMonitoredItemsRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModifyMonitoredItemsRequest")
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

	timestampsToReturn, err := ReadEnumField[TimestampsToReturn](ctx, "timestampsToReturn", "TimestampsToReturn", ReadEnum(TimestampsToReturnByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timestampsToReturn' field"))
	}
	m.TimestampsToReturn = timestampsToReturn

	noOfItemsToModify, err := ReadImplicitField[int32](ctx, "noOfItemsToModify", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfItemsToModify' field"))
	}
	_ = noOfItemsToModify

	itemsToModify, err := ReadCountArrayField[MonitoredItemModifyRequest](ctx, "itemsToModify", ReadComplex[MonitoredItemModifyRequest](ExtensionObjectDefinitionParseWithBufferProducer[MonitoredItemModifyRequest]((int32)(int32(757))), readBuffer), uint64(noOfItemsToModify))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemsToModify' field"))
	}
	m.ItemsToModify = itemsToModify

	if closeErr := readBuffer.CloseContext("ModifyMonitoredItemsRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModifyMonitoredItemsRequest")
	}

	return m, nil
}

func (m *_ModifyMonitoredItemsRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModifyMonitoredItemsRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModifyMonitoredItemsRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModifyMonitoredItemsRequest")
		}

		if err := WriteSimpleField[RequestHeader](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[RequestHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[uint32](ctx, "subscriptionId", m.GetSubscriptionId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'subscriptionId' field")
		}

		if err := WriteSimpleEnumField[TimestampsToReturn](ctx, "timestampsToReturn", "TimestampsToReturn", m.GetTimestampsToReturn(), WriteEnum[TimestampsToReturn, uint32](TimestampsToReturn.GetValue, TimestampsToReturn.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'timestampsToReturn' field")
		}
		noOfItemsToModify := int32(utils.InlineIf(bool((m.GetItemsToModify()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetItemsToModify()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfItemsToModify", noOfItemsToModify, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfItemsToModify' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "itemsToModify", m.GetItemsToModify(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'itemsToModify' field")
		}

		if popErr := writeBuffer.PopContext("ModifyMonitoredItemsRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModifyMonitoredItemsRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModifyMonitoredItemsRequest) IsModifyMonitoredItemsRequest() {}

func (m *_ModifyMonitoredItemsRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModifyMonitoredItemsRequest) deepCopy() *_ModifyMonitoredItemsRequest {
	if m == nil {
		return nil
	}
	_ModifyMonitoredItemsRequestCopy := &_ModifyMonitoredItemsRequest{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[RequestHeader](m.RequestHeader),
		m.SubscriptionId,
		m.TimestampsToReturn,
		utils.DeepCopySlice[MonitoredItemModifyRequest, MonitoredItemModifyRequest](m.ItemsToModify),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _ModifyMonitoredItemsRequestCopy
}

func (m *_ModifyMonitoredItemsRequest) String() string {
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
