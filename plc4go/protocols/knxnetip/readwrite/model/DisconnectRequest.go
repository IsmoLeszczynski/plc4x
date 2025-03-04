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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// DisconnectRequest is the corresponding interface of DisconnectRequest
type DisconnectRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	KnxNetIpMessage
	// GetCommunicationChannelId returns CommunicationChannelId (property field)
	GetCommunicationChannelId() uint8
	// GetHpaiControlEndpoint returns HpaiControlEndpoint (property field)
	GetHpaiControlEndpoint() HPAIControlEndpoint
	// IsDisconnectRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDisconnectRequest()
	// CreateBuilder creates a DisconnectRequestBuilder
	CreateDisconnectRequestBuilder() DisconnectRequestBuilder
}

// _DisconnectRequest is the data-structure of this message
type _DisconnectRequest struct {
	KnxNetIpMessageContract
	CommunicationChannelId uint8
	HpaiControlEndpoint    HPAIControlEndpoint
	// Reserved Fields
	reservedField0 *uint8
}

var _ DisconnectRequest = (*_DisconnectRequest)(nil)
var _ KnxNetIpMessageRequirements = (*_DisconnectRequest)(nil)

// NewDisconnectRequest factory function for _DisconnectRequest
func NewDisconnectRequest(communicationChannelId uint8, hpaiControlEndpoint HPAIControlEndpoint) *_DisconnectRequest {
	if hpaiControlEndpoint == nil {
		panic("hpaiControlEndpoint of type HPAIControlEndpoint for DisconnectRequest must not be nil")
	}
	_result := &_DisconnectRequest{
		KnxNetIpMessageContract: NewKnxNetIpMessage(),
		CommunicationChannelId:  communicationChannelId,
		HpaiControlEndpoint:     hpaiControlEndpoint,
	}
	_result.KnxNetIpMessageContract.(*_KnxNetIpMessage)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DisconnectRequestBuilder is a builder for DisconnectRequest
type DisconnectRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(communicationChannelId uint8, hpaiControlEndpoint HPAIControlEndpoint) DisconnectRequestBuilder
	// WithCommunicationChannelId adds CommunicationChannelId (property field)
	WithCommunicationChannelId(uint8) DisconnectRequestBuilder
	// WithHpaiControlEndpoint adds HpaiControlEndpoint (property field)
	WithHpaiControlEndpoint(HPAIControlEndpoint) DisconnectRequestBuilder
	// WithHpaiControlEndpointBuilder adds HpaiControlEndpoint (property field) which is build by the builder
	WithHpaiControlEndpointBuilder(func(HPAIControlEndpointBuilder) HPAIControlEndpointBuilder) DisconnectRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() KnxNetIpMessageBuilder
	// Build builds the DisconnectRequest or returns an error if something is wrong
	Build() (DisconnectRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DisconnectRequest
}

// NewDisconnectRequestBuilder() creates a DisconnectRequestBuilder
func NewDisconnectRequestBuilder() DisconnectRequestBuilder {
	return &_DisconnectRequestBuilder{_DisconnectRequest: new(_DisconnectRequest)}
}

type _DisconnectRequestBuilder struct {
	*_DisconnectRequest

	parentBuilder *_KnxNetIpMessageBuilder

	err *utils.MultiError
}

var _ (DisconnectRequestBuilder) = (*_DisconnectRequestBuilder)(nil)

func (b *_DisconnectRequestBuilder) setParent(contract KnxNetIpMessageContract) {
	b.KnxNetIpMessageContract = contract
	contract.(*_KnxNetIpMessage)._SubType = b._DisconnectRequest
}

func (b *_DisconnectRequestBuilder) WithMandatoryFields(communicationChannelId uint8, hpaiControlEndpoint HPAIControlEndpoint) DisconnectRequestBuilder {
	return b.WithCommunicationChannelId(communicationChannelId).WithHpaiControlEndpoint(hpaiControlEndpoint)
}

func (b *_DisconnectRequestBuilder) WithCommunicationChannelId(communicationChannelId uint8) DisconnectRequestBuilder {
	b.CommunicationChannelId = communicationChannelId
	return b
}

func (b *_DisconnectRequestBuilder) WithHpaiControlEndpoint(hpaiControlEndpoint HPAIControlEndpoint) DisconnectRequestBuilder {
	b.HpaiControlEndpoint = hpaiControlEndpoint
	return b
}

func (b *_DisconnectRequestBuilder) WithHpaiControlEndpointBuilder(builderSupplier func(HPAIControlEndpointBuilder) HPAIControlEndpointBuilder) DisconnectRequestBuilder {
	builder := builderSupplier(b.HpaiControlEndpoint.CreateHPAIControlEndpointBuilder())
	var err error
	b.HpaiControlEndpoint, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "HPAIControlEndpointBuilder failed"))
	}
	return b
}

func (b *_DisconnectRequestBuilder) Build() (DisconnectRequest, error) {
	if b.HpaiControlEndpoint == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'hpaiControlEndpoint' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._DisconnectRequest.deepCopy(), nil
}

func (b *_DisconnectRequestBuilder) MustBuild() DisconnectRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_DisconnectRequestBuilder) Done() KnxNetIpMessageBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewKnxNetIpMessageBuilder().(*_KnxNetIpMessageBuilder)
	}
	return b.parentBuilder
}

func (b *_DisconnectRequestBuilder) buildForKnxNetIpMessage() (KnxNetIpMessage, error) {
	return b.Build()
}

func (b *_DisconnectRequestBuilder) DeepCopy() any {
	_copy := b.CreateDisconnectRequestBuilder().(*_DisconnectRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDisconnectRequestBuilder creates a DisconnectRequestBuilder
func (b *_DisconnectRequest) CreateDisconnectRequestBuilder() DisconnectRequestBuilder {
	if b == nil {
		return NewDisconnectRequestBuilder()
	}
	return &_DisconnectRequestBuilder{_DisconnectRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DisconnectRequest) GetMsgType() uint16 {
	return 0x0209
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DisconnectRequest) GetParent() KnxNetIpMessageContract {
	return m.KnxNetIpMessageContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DisconnectRequest) GetCommunicationChannelId() uint8 {
	return m.CommunicationChannelId
}

func (m *_DisconnectRequest) GetHpaiControlEndpoint() HPAIControlEndpoint {
	return m.HpaiControlEndpoint
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDisconnectRequest(structType any) DisconnectRequest {
	if casted, ok := structType.(DisconnectRequest); ok {
		return casted
	}
	if casted, ok := structType.(*DisconnectRequest); ok {
		return *casted
	}
	return nil
}

func (m *_DisconnectRequest) GetTypeName() string {
	return "DisconnectRequest"
}

func (m *_DisconnectRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.KnxNetIpMessageContract.(*_KnxNetIpMessage).getLengthInBits(ctx))

	// Simple field (communicationChannelId)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (hpaiControlEndpoint)
	lengthInBits += m.HpaiControlEndpoint.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_DisconnectRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DisconnectRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_KnxNetIpMessage) (__disconnectRequest DisconnectRequest, err error) {
	m.KnxNetIpMessageContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DisconnectRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DisconnectRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	communicationChannelId, err := ReadSimpleField(ctx, "communicationChannelId", ReadUnsignedByte(readBuffer, uint8(8)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'communicationChannelId' field"))
	}
	m.CommunicationChannelId = communicationChannelId

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	hpaiControlEndpoint, err := ReadSimpleField[HPAIControlEndpoint](ctx, "hpaiControlEndpoint", ReadComplex[HPAIControlEndpoint](HPAIControlEndpointParseWithBuffer, readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hpaiControlEndpoint' field"))
	}
	m.HpaiControlEndpoint = hpaiControlEndpoint

	if closeErr := readBuffer.CloseContext("DisconnectRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DisconnectRequest")
	}

	return m, nil
}

func (m *_DisconnectRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DisconnectRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DisconnectRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DisconnectRequest")
		}

		if err := WriteSimpleField[uint8](ctx, "communicationChannelId", m.GetCommunicationChannelId(), WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'communicationChannelId' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[HPAIControlEndpoint](ctx, "hpaiControlEndpoint", m.GetHpaiControlEndpoint(), WriteComplex[HPAIControlEndpoint](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'hpaiControlEndpoint' field")
		}

		if popErr := writeBuffer.PopContext("DisconnectRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DisconnectRequest")
		}
		return nil
	}
	return m.KnxNetIpMessageContract.(*_KnxNetIpMessage).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DisconnectRequest) IsDisconnectRequest() {}

func (m *_DisconnectRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DisconnectRequest) deepCopy() *_DisconnectRequest {
	if m == nil {
		return nil
	}
	_DisconnectRequestCopy := &_DisconnectRequest{
		m.KnxNetIpMessageContract.(*_KnxNetIpMessage).deepCopy(),
		m.CommunicationChannelId,
		utils.DeepCopy[HPAIControlEndpoint](m.HpaiControlEndpoint),
		m.reservedField0,
	}
	_DisconnectRequestCopy.KnxNetIpMessageContract.(*_KnxNetIpMessage)._SubType = m
	return _DisconnectRequestCopy
}

func (m *_DisconnectRequest) String() string {
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
