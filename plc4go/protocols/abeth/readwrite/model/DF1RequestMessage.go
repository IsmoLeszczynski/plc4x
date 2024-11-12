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

// DF1RequestMessage is the corresponding interface of DF1RequestMessage
type DF1RequestMessage interface {
	DF1RequestMessageContract
	DF1RequestMessageRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsDF1RequestMessage is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDF1RequestMessage()
	// CreateBuilder creates a DF1RequestMessageBuilder
	CreateDF1RequestMessageBuilder() DF1RequestMessageBuilder
}

// DF1RequestMessageContract provides a set of functions which can be overwritten by a sub struct
type DF1RequestMessageContract interface {
	// GetDestinationAddress returns DestinationAddress (property field)
	GetDestinationAddress() uint8
	// GetSourceAddress returns SourceAddress (property field)
	GetSourceAddress() uint8
	// GetStatus returns Status (property field)
	GetStatus() uint8
	// GetTransactionCounter returns TransactionCounter (property field)
	GetTransactionCounter() uint16
	// IsDF1RequestMessage is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDF1RequestMessage()
	// CreateBuilder creates a DF1RequestMessageBuilder
	CreateDF1RequestMessageBuilder() DF1RequestMessageBuilder
}

// DF1RequestMessageRequirements provides a set of functions which need to be implemented by a sub struct
type DF1RequestMessageRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetCommandCode returns CommandCode (discriminator field)
	GetCommandCode() uint8
}

// _DF1RequestMessage is the data-structure of this message
type _DF1RequestMessage struct {
	_SubType interface {
		DF1RequestMessageContract
		DF1RequestMessageRequirements
	}
	DestinationAddress uint8
	SourceAddress      uint8
	Status             uint8
	TransactionCounter uint16
	// Reserved Fields
	reservedField0 *uint16
}

var _ DF1RequestMessageContract = (*_DF1RequestMessage)(nil)

// NewDF1RequestMessage factory function for _DF1RequestMessage
func NewDF1RequestMessage(destinationAddress uint8, sourceAddress uint8, status uint8, transactionCounter uint16) *_DF1RequestMessage {
	return &_DF1RequestMessage{DestinationAddress: destinationAddress, SourceAddress: sourceAddress, Status: status, TransactionCounter: transactionCounter}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DF1RequestMessageBuilder is a builder for DF1RequestMessage
type DF1RequestMessageBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(destinationAddress uint8, sourceAddress uint8, status uint8, transactionCounter uint16) DF1RequestMessageBuilder
	// WithDestinationAddress adds DestinationAddress (property field)
	WithDestinationAddress(uint8) DF1RequestMessageBuilder
	// WithSourceAddress adds SourceAddress (property field)
	WithSourceAddress(uint8) DF1RequestMessageBuilder
	// WithStatus adds Status (property field)
	WithStatus(uint8) DF1RequestMessageBuilder
	// WithTransactionCounter adds TransactionCounter (property field)
	WithTransactionCounter(uint16) DF1RequestMessageBuilder
	// AsDF1CommandRequestMessage converts this build to a subType of DF1RequestMessage. It is always possible to return to current builder using Done()
	AsDF1CommandRequestMessage() DF1CommandRequestMessageBuilder
	// Build builds the DF1RequestMessage or returns an error if something is wrong
	PartialBuild() (DF1RequestMessageContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() DF1RequestMessageContract
	// Build builds the DF1RequestMessage or returns an error if something is wrong
	Build() (DF1RequestMessage, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DF1RequestMessage
}

// NewDF1RequestMessageBuilder() creates a DF1RequestMessageBuilder
func NewDF1RequestMessageBuilder() DF1RequestMessageBuilder {
	return &_DF1RequestMessageBuilder{_DF1RequestMessage: new(_DF1RequestMessage)}
}

type _DF1RequestMessageChildBuilder interface {
	utils.Copyable
	setParent(DF1RequestMessageContract)
	buildForDF1RequestMessage() (DF1RequestMessage, error)
}

type _DF1RequestMessageBuilder struct {
	*_DF1RequestMessage

	childBuilder _DF1RequestMessageChildBuilder

	err *utils.MultiError
}

var _ (DF1RequestMessageBuilder) = (*_DF1RequestMessageBuilder)(nil)

func (b *_DF1RequestMessageBuilder) WithMandatoryFields(destinationAddress uint8, sourceAddress uint8, status uint8, transactionCounter uint16) DF1RequestMessageBuilder {
	return b.WithDestinationAddress(destinationAddress).WithSourceAddress(sourceAddress).WithStatus(status).WithTransactionCounter(transactionCounter)
}

func (b *_DF1RequestMessageBuilder) WithDestinationAddress(destinationAddress uint8) DF1RequestMessageBuilder {
	b.DestinationAddress = destinationAddress
	return b
}

func (b *_DF1RequestMessageBuilder) WithSourceAddress(sourceAddress uint8) DF1RequestMessageBuilder {
	b.SourceAddress = sourceAddress
	return b
}

func (b *_DF1RequestMessageBuilder) WithStatus(status uint8) DF1RequestMessageBuilder {
	b.Status = status
	return b
}

func (b *_DF1RequestMessageBuilder) WithTransactionCounter(transactionCounter uint16) DF1RequestMessageBuilder {
	b.TransactionCounter = transactionCounter
	return b
}

func (b *_DF1RequestMessageBuilder) PartialBuild() (DF1RequestMessageContract, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._DF1RequestMessage.deepCopy(), nil
}

func (b *_DF1RequestMessageBuilder) PartialMustBuild() DF1RequestMessageContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_DF1RequestMessageBuilder) AsDF1CommandRequestMessage() DF1CommandRequestMessageBuilder {
	if cb, ok := b.childBuilder.(DF1CommandRequestMessageBuilder); ok {
		return cb
	}
	cb := NewDF1CommandRequestMessageBuilder().(*_DF1CommandRequestMessageBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_DF1RequestMessageBuilder) Build() (DF1RequestMessage, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForDF1RequestMessage()
}

func (b *_DF1RequestMessageBuilder) MustBuild() DF1RequestMessage {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_DF1RequestMessageBuilder) DeepCopy() any {
	_copy := b.CreateDF1RequestMessageBuilder().(*_DF1RequestMessageBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_DF1RequestMessageChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDF1RequestMessageBuilder creates a DF1RequestMessageBuilder
func (b *_DF1RequestMessage) CreateDF1RequestMessageBuilder() DF1RequestMessageBuilder {
	if b == nil {
		return NewDF1RequestMessageBuilder()
	}
	return &_DF1RequestMessageBuilder{_DF1RequestMessage: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DF1RequestMessage) GetDestinationAddress() uint8 {
	return m.DestinationAddress
}

func (m *_DF1RequestMessage) GetSourceAddress() uint8 {
	return m.SourceAddress
}

func (m *_DF1RequestMessage) GetStatus() uint8 {
	return m.Status
}

func (m *_DF1RequestMessage) GetTransactionCounter() uint16 {
	return m.TransactionCounter
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDF1RequestMessage(structType any) DF1RequestMessage {
	if casted, ok := structType.(DF1RequestMessage); ok {
		return casted
	}
	if casted, ok := structType.(*DF1RequestMessage); ok {
		return *casted
	}
	return nil
}

func (m *_DF1RequestMessage) GetTypeName() string {
	return "DF1RequestMessage"
}

func (m *_DF1RequestMessage) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (destinationAddress)
	lengthInBits += 8

	// Simple field (sourceAddress)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 16
	// Discriminator Field (commandCode)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	// Simple field (transactionCounter)
	lengthInBits += 16

	return lengthInBits
}

func (m *_DF1RequestMessage) GetLengthInBits(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx)
}

func (m *_DF1RequestMessage) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func DF1RequestMessageParse[T DF1RequestMessage](ctx context.Context, theBytes []byte) (T, error) {
	return DF1RequestMessageParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func DF1RequestMessageParseWithBufferProducer[T DF1RequestMessage]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := DF1RequestMessageParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func DF1RequestMessageParseWithBuffer[T DF1RequestMessage](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_DF1RequestMessage{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_DF1RequestMessage) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__dF1RequestMessage DF1RequestMessage, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DF1RequestMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DF1RequestMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	destinationAddress, err := ReadSimpleField(ctx, "destinationAddress", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'destinationAddress' field"))
	}
	m.DestinationAddress = destinationAddress

	sourceAddress, err := ReadSimpleField(ctx, "sourceAddress", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sourceAddress' field"))
	}
	m.SourceAddress = sourceAddress

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedShort(readBuffer, uint8(16)), uint16(0x0000))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	commandCode, err := ReadDiscriminatorField[uint8](ctx, "commandCode", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandCode' field"))
	}

	status, err := ReadSimpleField(ctx, "status", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'status' field"))
	}
	m.Status = status

	transactionCounter, err := ReadSimpleField(ctx, "transactionCounter", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transactionCounter' field"))
	}
	m.TransactionCounter = transactionCounter

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child DF1RequestMessage
	switch {
	case commandCode == 0x0F: // DF1CommandRequestMessage
		if _child, err = new(_DF1CommandRequestMessage).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type DF1CommandRequestMessage for type-switch of DF1RequestMessage")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [commandCode=%v]", commandCode)
	}

	if closeErr := readBuffer.CloseContext("DF1RequestMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DF1RequestMessage")
	}

	return _child, nil
}

func (pm *_DF1RequestMessage) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child DF1RequestMessage, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("DF1RequestMessage"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for DF1RequestMessage")
	}

	if err := WriteSimpleField[uint8](ctx, "destinationAddress", m.GetDestinationAddress(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'destinationAddress' field")
	}

	if err := WriteSimpleField[uint8](ctx, "sourceAddress", m.GetSourceAddress(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'sourceAddress' field")
	}

	if err := WriteReservedField[uint16](ctx, "reserved", uint16(0x0000), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 1")
	}

	if err := WriteDiscriminatorField(ctx, "commandCode", m.GetCommandCode(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'commandCode' field")
	}

	if err := WriteSimpleField[uint8](ctx, "status", m.GetStatus(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'status' field")
	}

	if err := WriteSimpleField[uint16](ctx, "transactionCounter", m.GetTransactionCounter(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'transactionCounter' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("DF1RequestMessage"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for DF1RequestMessage")
	}
	return nil
}

func (m *_DF1RequestMessage) IsDF1RequestMessage() {}

func (m *_DF1RequestMessage) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DF1RequestMessage) deepCopy() *_DF1RequestMessage {
	if m == nil {
		return nil
	}
	_DF1RequestMessageCopy := &_DF1RequestMessage{
		nil, // will be set by child
		m.DestinationAddress,
		m.SourceAddress,
		m.Status,
		m.TransactionCounter,
		m.reservedField0,
	}
	return _DF1RequestMessageCopy
}
