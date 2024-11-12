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

// ModbusPDUReadHoldingRegistersResponse is the corresponding interface of ModbusPDUReadHoldingRegistersResponse
type ModbusPDUReadHoldingRegistersResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ModbusPDU
	// GetValue returns Value (property field)
	GetValue() []byte
	// IsModbusPDUReadHoldingRegistersResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusPDUReadHoldingRegistersResponse()
	// CreateBuilder creates a ModbusPDUReadHoldingRegistersResponseBuilder
	CreateModbusPDUReadHoldingRegistersResponseBuilder() ModbusPDUReadHoldingRegistersResponseBuilder
}

// _ModbusPDUReadHoldingRegistersResponse is the data-structure of this message
type _ModbusPDUReadHoldingRegistersResponse struct {
	ModbusPDUContract
	Value []byte
}

var _ ModbusPDUReadHoldingRegistersResponse = (*_ModbusPDUReadHoldingRegistersResponse)(nil)
var _ ModbusPDURequirements = (*_ModbusPDUReadHoldingRegistersResponse)(nil)

// NewModbusPDUReadHoldingRegistersResponse factory function for _ModbusPDUReadHoldingRegistersResponse
func NewModbusPDUReadHoldingRegistersResponse(value []byte) *_ModbusPDUReadHoldingRegistersResponse {
	_result := &_ModbusPDUReadHoldingRegistersResponse{
		ModbusPDUContract: NewModbusPDU(),
		Value:             value,
	}
	_result.ModbusPDUContract.(*_ModbusPDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ModbusPDUReadHoldingRegistersResponseBuilder is a builder for ModbusPDUReadHoldingRegistersResponse
type ModbusPDUReadHoldingRegistersResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(value []byte) ModbusPDUReadHoldingRegistersResponseBuilder
	// WithValue adds Value (property field)
	WithValue(...byte) ModbusPDUReadHoldingRegistersResponseBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ModbusPDUBuilder
	// Build builds the ModbusPDUReadHoldingRegistersResponse or returns an error if something is wrong
	Build() (ModbusPDUReadHoldingRegistersResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ModbusPDUReadHoldingRegistersResponse
}

// NewModbusPDUReadHoldingRegistersResponseBuilder() creates a ModbusPDUReadHoldingRegistersResponseBuilder
func NewModbusPDUReadHoldingRegistersResponseBuilder() ModbusPDUReadHoldingRegistersResponseBuilder {
	return &_ModbusPDUReadHoldingRegistersResponseBuilder{_ModbusPDUReadHoldingRegistersResponse: new(_ModbusPDUReadHoldingRegistersResponse)}
}

type _ModbusPDUReadHoldingRegistersResponseBuilder struct {
	*_ModbusPDUReadHoldingRegistersResponse

	parentBuilder *_ModbusPDUBuilder

	err *utils.MultiError
}

var _ (ModbusPDUReadHoldingRegistersResponseBuilder) = (*_ModbusPDUReadHoldingRegistersResponseBuilder)(nil)

func (b *_ModbusPDUReadHoldingRegistersResponseBuilder) setParent(contract ModbusPDUContract) {
	b.ModbusPDUContract = contract
}

func (b *_ModbusPDUReadHoldingRegistersResponseBuilder) WithMandatoryFields(value []byte) ModbusPDUReadHoldingRegistersResponseBuilder {
	return b.WithValue(value...)
}

func (b *_ModbusPDUReadHoldingRegistersResponseBuilder) WithValue(value ...byte) ModbusPDUReadHoldingRegistersResponseBuilder {
	b.Value = value
	return b
}

func (b *_ModbusPDUReadHoldingRegistersResponseBuilder) Build() (ModbusPDUReadHoldingRegistersResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ModbusPDUReadHoldingRegistersResponse.deepCopy(), nil
}

func (b *_ModbusPDUReadHoldingRegistersResponseBuilder) MustBuild() ModbusPDUReadHoldingRegistersResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ModbusPDUReadHoldingRegistersResponseBuilder) Done() ModbusPDUBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewModbusPDUBuilder().(*_ModbusPDUBuilder)
	}
	return b.parentBuilder
}

func (b *_ModbusPDUReadHoldingRegistersResponseBuilder) buildForModbusPDU() (ModbusPDU, error) {
	return b.Build()
}

func (b *_ModbusPDUReadHoldingRegistersResponseBuilder) DeepCopy() any {
	_copy := b.CreateModbusPDUReadHoldingRegistersResponseBuilder().(*_ModbusPDUReadHoldingRegistersResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateModbusPDUReadHoldingRegistersResponseBuilder creates a ModbusPDUReadHoldingRegistersResponseBuilder
func (b *_ModbusPDUReadHoldingRegistersResponse) CreateModbusPDUReadHoldingRegistersResponseBuilder() ModbusPDUReadHoldingRegistersResponseBuilder {
	if b == nil {
		return NewModbusPDUReadHoldingRegistersResponseBuilder()
	}
	return &_ModbusPDUReadHoldingRegistersResponseBuilder{_ModbusPDUReadHoldingRegistersResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReadHoldingRegistersResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUReadHoldingRegistersResponse) GetFunctionFlag() uint8 {
	return 0x03
}

func (m *_ModbusPDUReadHoldingRegistersResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReadHoldingRegistersResponse) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReadHoldingRegistersResponse) GetValue() []byte {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastModbusPDUReadHoldingRegistersResponse(structType any) ModbusPDUReadHoldingRegistersResponse {
	if casted, ok := structType.(ModbusPDUReadHoldingRegistersResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadHoldingRegistersResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadHoldingRegistersResponse) GetTypeName() string {
	return "ModbusPDUReadHoldingRegistersResponse"
}

func (m *_ModbusPDUReadHoldingRegistersResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 8 * uint16(len(m.Value))
	}

	return lengthInBits
}

func (m *_ModbusPDUReadHoldingRegistersResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUReadHoldingRegistersResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (__modbusPDUReadHoldingRegistersResponse ModbusPDUReadHoldingRegistersResponse, err error) {
	m.ModbusPDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadHoldingRegistersResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadHoldingRegistersResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

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

	if closeErr := readBuffer.CloseContext("ModbusPDUReadHoldingRegistersResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadHoldingRegistersResponse")
	}

	return m, nil
}

func (m *_ModbusPDUReadHoldingRegistersResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReadHoldingRegistersResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadHoldingRegistersResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadHoldingRegistersResponse")
		}
		byteCount := uint8(uint8(len(m.GetValue())))
		if err := WriteImplicitField(ctx, "byteCount", byteCount, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'byteCount' field")
		}

		if err := WriteByteArrayField(ctx, "value", m.GetValue(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadHoldingRegistersResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReadHoldingRegistersResponse")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUReadHoldingRegistersResponse) IsModbusPDUReadHoldingRegistersResponse() {}

func (m *_ModbusPDUReadHoldingRegistersResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModbusPDUReadHoldingRegistersResponse) deepCopy() *_ModbusPDUReadHoldingRegistersResponse {
	if m == nil {
		return nil
	}
	_ModbusPDUReadHoldingRegistersResponseCopy := &_ModbusPDUReadHoldingRegistersResponse{
		m.ModbusPDUContract.(*_ModbusPDU).deepCopy(),
		utils.DeepCopySlice[byte, byte](m.Value),
	}
	m.ModbusPDUContract.(*_ModbusPDU)._SubType = m
	return _ModbusPDUReadHoldingRegistersResponseCopy
}

func (m *_ModbusPDUReadHoldingRegistersResponse) String() string {
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
