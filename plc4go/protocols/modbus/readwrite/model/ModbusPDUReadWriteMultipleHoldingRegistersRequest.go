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

// ModbusPDUReadWriteMultipleHoldingRegistersRequest is the corresponding interface of ModbusPDUReadWriteMultipleHoldingRegistersRequest
type ModbusPDUReadWriteMultipleHoldingRegistersRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ModbusPDU
	// GetReadStartingAddress returns ReadStartingAddress (property field)
	GetReadStartingAddress() uint16
	// GetReadQuantity returns ReadQuantity (property field)
	GetReadQuantity() uint16
	// GetWriteStartingAddress returns WriteStartingAddress (property field)
	GetWriteStartingAddress() uint16
	// GetWriteQuantity returns WriteQuantity (property field)
	GetWriteQuantity() uint16
	// GetValue returns Value (property field)
	GetValue() []byte
	// IsModbusPDUReadWriteMultipleHoldingRegistersRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusPDUReadWriteMultipleHoldingRegistersRequest()
	// CreateBuilder creates a ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder
	CreateModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder() ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder
}

// _ModbusPDUReadWriteMultipleHoldingRegistersRequest is the data-structure of this message
type _ModbusPDUReadWriteMultipleHoldingRegistersRequest struct {
	ModbusPDUContract
	ReadStartingAddress  uint16
	ReadQuantity         uint16
	WriteStartingAddress uint16
	WriteQuantity        uint16
	Value                []byte
}

var _ ModbusPDUReadWriteMultipleHoldingRegistersRequest = (*_ModbusPDUReadWriteMultipleHoldingRegistersRequest)(nil)
var _ ModbusPDURequirements = (*_ModbusPDUReadWriteMultipleHoldingRegistersRequest)(nil)

// NewModbusPDUReadWriteMultipleHoldingRegistersRequest factory function for _ModbusPDUReadWriteMultipleHoldingRegistersRequest
func NewModbusPDUReadWriteMultipleHoldingRegistersRequest(readStartingAddress uint16, readQuantity uint16, writeStartingAddress uint16, writeQuantity uint16, value []byte) *_ModbusPDUReadWriteMultipleHoldingRegistersRequest {
	_result := &_ModbusPDUReadWriteMultipleHoldingRegistersRequest{
		ModbusPDUContract:    NewModbusPDU(),
		ReadStartingAddress:  readStartingAddress,
		ReadQuantity:         readQuantity,
		WriteStartingAddress: writeStartingAddress,
		WriteQuantity:        writeQuantity,
		Value:                value,
	}
	_result.ModbusPDUContract.(*_ModbusPDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder is a builder for ModbusPDUReadWriteMultipleHoldingRegistersRequest
type ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(readStartingAddress uint16, readQuantity uint16, writeStartingAddress uint16, writeQuantity uint16, value []byte) ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder
	// WithReadStartingAddress adds ReadStartingAddress (property field)
	WithReadStartingAddress(uint16) ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder
	// WithReadQuantity adds ReadQuantity (property field)
	WithReadQuantity(uint16) ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder
	// WithWriteStartingAddress adds WriteStartingAddress (property field)
	WithWriteStartingAddress(uint16) ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder
	// WithWriteQuantity adds WriteQuantity (property field)
	WithWriteQuantity(uint16) ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder
	// WithValue adds Value (property field)
	WithValue(...byte) ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ModbusPDUBuilder
	// Build builds the ModbusPDUReadWriteMultipleHoldingRegistersRequest or returns an error if something is wrong
	Build() (ModbusPDUReadWriteMultipleHoldingRegistersRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ModbusPDUReadWriteMultipleHoldingRegistersRequest
}

// NewModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder() creates a ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder
func NewModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder() ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder {
	return &_ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder{_ModbusPDUReadWriteMultipleHoldingRegistersRequest: new(_ModbusPDUReadWriteMultipleHoldingRegistersRequest)}
}

type _ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder struct {
	*_ModbusPDUReadWriteMultipleHoldingRegistersRequest

	parentBuilder *_ModbusPDUBuilder

	err *utils.MultiError
}

var _ (ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder) = (*_ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder)(nil)

func (b *_ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder) setParent(contract ModbusPDUContract) {
	b.ModbusPDUContract = contract
}

func (b *_ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder) WithMandatoryFields(readStartingAddress uint16, readQuantity uint16, writeStartingAddress uint16, writeQuantity uint16, value []byte) ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder {
	return b.WithReadStartingAddress(readStartingAddress).WithReadQuantity(readQuantity).WithWriteStartingAddress(writeStartingAddress).WithWriteQuantity(writeQuantity).WithValue(value...)
}

func (b *_ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder) WithReadStartingAddress(readStartingAddress uint16) ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder {
	b.ReadStartingAddress = readStartingAddress
	return b
}

func (b *_ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder) WithReadQuantity(readQuantity uint16) ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder {
	b.ReadQuantity = readQuantity
	return b
}

func (b *_ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder) WithWriteStartingAddress(writeStartingAddress uint16) ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder {
	b.WriteStartingAddress = writeStartingAddress
	return b
}

func (b *_ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder) WithWriteQuantity(writeQuantity uint16) ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder {
	b.WriteQuantity = writeQuantity
	return b
}

func (b *_ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder) WithValue(value ...byte) ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder {
	b.Value = value
	return b
}

func (b *_ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder) Build() (ModbusPDUReadWriteMultipleHoldingRegistersRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ModbusPDUReadWriteMultipleHoldingRegistersRequest.deepCopy(), nil
}

func (b *_ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder) MustBuild() ModbusPDUReadWriteMultipleHoldingRegistersRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder) Done() ModbusPDUBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewModbusPDUBuilder().(*_ModbusPDUBuilder)
	}
	return b.parentBuilder
}

func (b *_ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder) buildForModbusPDU() (ModbusPDU, error) {
	return b.Build()
}

func (b *_ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder) DeepCopy() any {
	_copy := b.CreateModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder().(*_ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder creates a ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder
func (b *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) CreateModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder() ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder {
	if b == nil {
		return NewModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder()
	}
	return &_ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder{_ModbusPDUReadWriteMultipleHoldingRegistersRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetFunctionFlag() uint8 {
	return 0x17
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetReadStartingAddress() uint16 {
	return m.ReadStartingAddress
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetReadQuantity() uint16 {
	return m.ReadQuantity
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetWriteStartingAddress() uint16 {
	return m.WriteStartingAddress
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetWriteQuantity() uint16 {
	return m.WriteQuantity
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetValue() []byte {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastModbusPDUReadWriteMultipleHoldingRegistersRequest(structType any) ModbusPDUReadWriteMultipleHoldingRegistersRequest {
	if casted, ok := structType.(ModbusPDUReadWriteMultipleHoldingRegistersRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadWriteMultipleHoldingRegistersRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetTypeName() string {
	return "ModbusPDUReadWriteMultipleHoldingRegistersRequest"
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	// Simple field (readStartingAddress)
	lengthInBits += 16

	// Simple field (readQuantity)
	lengthInBits += 16

	// Simple field (writeStartingAddress)
	lengthInBits += 16

	// Simple field (writeQuantity)
	lengthInBits += 16

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 8 * uint16(len(m.Value))
	}

	return lengthInBits
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (__modbusPDUReadWriteMultipleHoldingRegistersRequest ModbusPDUReadWriteMultipleHoldingRegistersRequest, err error) {
	m.ModbusPDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadWriteMultipleHoldingRegistersRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadWriteMultipleHoldingRegistersRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	readStartingAddress, err := ReadSimpleField(ctx, "readStartingAddress", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'readStartingAddress' field"))
	}
	m.ReadStartingAddress = readStartingAddress

	readQuantity, err := ReadSimpleField(ctx, "readQuantity", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'readQuantity' field"))
	}
	m.ReadQuantity = readQuantity

	writeStartingAddress, err := ReadSimpleField(ctx, "writeStartingAddress", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'writeStartingAddress' field"))
	}
	m.WriteStartingAddress = writeStartingAddress

	writeQuantity, err := ReadSimpleField(ctx, "writeQuantity", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'writeQuantity' field"))
	}
	m.WriteQuantity = writeQuantity

	byteCount, err := ReadImplicitField[uint8](ctx, "byteCount", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'byteCount' field"))
	}
	_ = byteCount

	value, err := readBuffer.ReadByteArray("value", int(byteCount))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("ModbusPDUReadWriteMultipleHoldingRegistersRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadWriteMultipleHoldingRegistersRequest")
	}

	return m, nil
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadWriteMultipleHoldingRegistersRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadWriteMultipleHoldingRegistersRequest")
		}

		if err := WriteSimpleField[uint16](ctx, "readStartingAddress", m.GetReadStartingAddress(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'readStartingAddress' field")
		}

		if err := WriteSimpleField[uint16](ctx, "readQuantity", m.GetReadQuantity(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'readQuantity' field")
		}

		if err := WriteSimpleField[uint16](ctx, "writeStartingAddress", m.GetWriteStartingAddress(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'writeStartingAddress' field")
		}

		if err := WriteSimpleField[uint16](ctx, "writeQuantity", m.GetWriteQuantity(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'writeQuantity' field")
		}
		byteCount := uint8(uint8(len(m.GetValue())))
		if err := WriteImplicitField(ctx, "byteCount", byteCount, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'byteCount' field")
		}

		if err := WriteByteArrayField(ctx, "value", m.GetValue(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadWriteMultipleHoldingRegistersRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReadWriteMultipleHoldingRegistersRequest")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) IsModbusPDUReadWriteMultipleHoldingRegistersRequest() {
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) deepCopy() *_ModbusPDUReadWriteMultipleHoldingRegistersRequest {
	if m == nil {
		return nil
	}
	_ModbusPDUReadWriteMultipleHoldingRegistersRequestCopy := &_ModbusPDUReadWriteMultipleHoldingRegistersRequest{
		m.ModbusPDUContract.(*_ModbusPDU).deepCopy(),
		m.ReadStartingAddress,
		m.ReadQuantity,
		m.WriteStartingAddress,
		m.WriteQuantity,
		utils.DeepCopySlice[byte, byte](m.Value),
	}
	m.ModbusPDUContract.(*_ModbusPDU)._SubType = m
	return _ModbusPDUReadWriteMultipleHoldingRegistersRequestCopy
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) String() string {
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
