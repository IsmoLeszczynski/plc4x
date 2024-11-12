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

// CloseSessionRequest is the corresponding interface of CloseSessionRequest
type CloseSessionRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() RequestHeader
	// GetDeleteSubscriptions returns DeleteSubscriptions (property field)
	GetDeleteSubscriptions() bool
	// IsCloseSessionRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCloseSessionRequest()
	// CreateBuilder creates a CloseSessionRequestBuilder
	CreateCloseSessionRequestBuilder() CloseSessionRequestBuilder
}

// _CloseSessionRequest is the data-structure of this message
type _CloseSessionRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader       RequestHeader
	DeleteSubscriptions bool
	// Reserved Fields
	reservedField0 *uint8
}

var _ CloseSessionRequest = (*_CloseSessionRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_CloseSessionRequest)(nil)

// NewCloseSessionRequest factory function for _CloseSessionRequest
func NewCloseSessionRequest(requestHeader RequestHeader, deleteSubscriptions bool) *_CloseSessionRequest {
	if requestHeader == nil {
		panic("requestHeader of type RequestHeader for CloseSessionRequest must not be nil")
	}
	_result := &_CloseSessionRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		DeleteSubscriptions:               deleteSubscriptions,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CloseSessionRequestBuilder is a builder for CloseSessionRequest
type CloseSessionRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestHeader RequestHeader, deleteSubscriptions bool) CloseSessionRequestBuilder
	// WithRequestHeader adds RequestHeader (property field)
	WithRequestHeader(RequestHeader) CloseSessionRequestBuilder
	// WithRequestHeaderBuilder adds RequestHeader (property field) which is build by the builder
	WithRequestHeaderBuilder(func(RequestHeaderBuilder) RequestHeaderBuilder) CloseSessionRequestBuilder
	// WithDeleteSubscriptions adds DeleteSubscriptions (property field)
	WithDeleteSubscriptions(bool) CloseSessionRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the CloseSessionRequest or returns an error if something is wrong
	Build() (CloseSessionRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CloseSessionRequest
}

// NewCloseSessionRequestBuilder() creates a CloseSessionRequestBuilder
func NewCloseSessionRequestBuilder() CloseSessionRequestBuilder {
	return &_CloseSessionRequestBuilder{_CloseSessionRequest: new(_CloseSessionRequest)}
}

type _CloseSessionRequestBuilder struct {
	*_CloseSessionRequest

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (CloseSessionRequestBuilder) = (*_CloseSessionRequestBuilder)(nil)

func (b *_CloseSessionRequestBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_CloseSessionRequestBuilder) WithMandatoryFields(requestHeader RequestHeader, deleteSubscriptions bool) CloseSessionRequestBuilder {
	return b.WithRequestHeader(requestHeader).WithDeleteSubscriptions(deleteSubscriptions)
}

func (b *_CloseSessionRequestBuilder) WithRequestHeader(requestHeader RequestHeader) CloseSessionRequestBuilder {
	b.RequestHeader = requestHeader
	return b
}

func (b *_CloseSessionRequestBuilder) WithRequestHeaderBuilder(builderSupplier func(RequestHeaderBuilder) RequestHeaderBuilder) CloseSessionRequestBuilder {
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

func (b *_CloseSessionRequestBuilder) WithDeleteSubscriptions(deleteSubscriptions bool) CloseSessionRequestBuilder {
	b.DeleteSubscriptions = deleteSubscriptions
	return b
}

func (b *_CloseSessionRequestBuilder) Build() (CloseSessionRequest, error) {
	if b.RequestHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'requestHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CloseSessionRequest.deepCopy(), nil
}

func (b *_CloseSessionRequestBuilder) MustBuild() CloseSessionRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CloseSessionRequestBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_CloseSessionRequestBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_CloseSessionRequestBuilder) DeepCopy() any {
	_copy := b.CreateCloseSessionRequestBuilder().(*_CloseSessionRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCloseSessionRequestBuilder creates a CloseSessionRequestBuilder
func (b *_CloseSessionRequest) CreateCloseSessionRequestBuilder() CloseSessionRequestBuilder {
	if b == nil {
		return NewCloseSessionRequestBuilder()
	}
	return &_CloseSessionRequestBuilder{_CloseSessionRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CloseSessionRequest) GetExtensionId() int32 {
	return int32(473)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CloseSessionRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CloseSessionRequest) GetRequestHeader() RequestHeader {
	return m.RequestHeader
}

func (m *_CloseSessionRequest) GetDeleteSubscriptions() bool {
	return m.DeleteSubscriptions
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCloseSessionRequest(structType any) CloseSessionRequest {
	if casted, ok := structType.(CloseSessionRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CloseSessionRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CloseSessionRequest) GetTypeName() string {
	return "CloseSessionRequest"
}

func (m *_CloseSessionRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (deleteSubscriptions)
	lengthInBits += 1

	return lengthInBits
}

func (m *_CloseSessionRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CloseSessionRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__closeSessionRequest CloseSessionRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CloseSessionRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CloseSessionRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[RequestHeader](ctx, "requestHeader", ReadComplex[RequestHeader](ExtensionObjectDefinitionParseWithBufferProducer[RequestHeader]((int32)(int32(391))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	deleteSubscriptions, err := ReadSimpleField(ctx, "deleteSubscriptions", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deleteSubscriptions' field"))
	}
	m.DeleteSubscriptions = deleteSubscriptions

	if closeErr := readBuffer.CloseContext("CloseSessionRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CloseSessionRequest")
	}

	return m, nil
}

func (m *_CloseSessionRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CloseSessionRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CloseSessionRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CloseSessionRequest")
		}

		if err := WriteSimpleField[RequestHeader](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[RequestHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "deleteSubscriptions", m.GetDeleteSubscriptions(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'deleteSubscriptions' field")
		}

		if popErr := writeBuffer.PopContext("CloseSessionRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CloseSessionRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CloseSessionRequest) IsCloseSessionRequest() {}

func (m *_CloseSessionRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CloseSessionRequest) deepCopy() *_CloseSessionRequest {
	if m == nil {
		return nil
	}
	_CloseSessionRequestCopy := &_CloseSessionRequest{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[RequestHeader](m.RequestHeader),
		m.DeleteSubscriptions,
		m.reservedField0,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _CloseSessionRequestCopy
}

func (m *_CloseSessionRequest) String() string {
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
