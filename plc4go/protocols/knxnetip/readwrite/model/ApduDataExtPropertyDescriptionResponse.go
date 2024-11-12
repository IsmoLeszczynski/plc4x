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

// ApduDataExtPropertyDescriptionResponse is the corresponding interface of ApduDataExtPropertyDescriptionResponse
type ApduDataExtPropertyDescriptionResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduDataExt
	// GetObjectIndex returns ObjectIndex (property field)
	GetObjectIndex() uint8
	// GetPropertyId returns PropertyId (property field)
	GetPropertyId() uint8
	// GetIndex returns Index (property field)
	GetIndex() uint8
	// GetWriteEnabled returns WriteEnabled (property field)
	GetWriteEnabled() bool
	// GetPropertyDataType returns PropertyDataType (property field)
	GetPropertyDataType() KnxPropertyDataType
	// GetMaxNrOfElements returns MaxNrOfElements (property field)
	GetMaxNrOfElements() uint16
	// GetReadLevel returns ReadLevel (property field)
	GetReadLevel() AccessLevel
	// GetWriteLevel returns WriteLevel (property field)
	GetWriteLevel() AccessLevel
	// IsApduDataExtPropertyDescriptionResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataExtPropertyDescriptionResponse()
	// CreateBuilder creates a ApduDataExtPropertyDescriptionResponseBuilder
	CreateApduDataExtPropertyDescriptionResponseBuilder() ApduDataExtPropertyDescriptionResponseBuilder
}

// _ApduDataExtPropertyDescriptionResponse is the data-structure of this message
type _ApduDataExtPropertyDescriptionResponse struct {
	ApduDataExtContract
	ObjectIndex      uint8
	PropertyId       uint8
	Index            uint8
	WriteEnabled     bool
	PropertyDataType KnxPropertyDataType
	MaxNrOfElements  uint16
	ReadLevel        AccessLevel
	WriteLevel       AccessLevel
	// Reserved Fields
	reservedField0 *uint8
	reservedField1 *uint8
}

var _ ApduDataExtPropertyDescriptionResponse = (*_ApduDataExtPropertyDescriptionResponse)(nil)
var _ ApduDataExtRequirements = (*_ApduDataExtPropertyDescriptionResponse)(nil)

// NewApduDataExtPropertyDescriptionResponse factory function for _ApduDataExtPropertyDescriptionResponse
func NewApduDataExtPropertyDescriptionResponse(objectIndex uint8, propertyId uint8, index uint8, writeEnabled bool, propertyDataType KnxPropertyDataType, maxNrOfElements uint16, readLevel AccessLevel, writeLevel AccessLevel, length uint8) *_ApduDataExtPropertyDescriptionResponse {
	_result := &_ApduDataExtPropertyDescriptionResponse{
		ApduDataExtContract: NewApduDataExt(length),
		ObjectIndex:         objectIndex,
		PropertyId:          propertyId,
		Index:               index,
		WriteEnabled:        writeEnabled,
		PropertyDataType:    propertyDataType,
		MaxNrOfElements:     maxNrOfElements,
		ReadLevel:           readLevel,
		WriteLevel:          writeLevel,
	}
	_result.ApduDataExtContract.(*_ApduDataExt)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ApduDataExtPropertyDescriptionResponseBuilder is a builder for ApduDataExtPropertyDescriptionResponse
type ApduDataExtPropertyDescriptionResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(objectIndex uint8, propertyId uint8, index uint8, writeEnabled bool, propertyDataType KnxPropertyDataType, maxNrOfElements uint16, readLevel AccessLevel, writeLevel AccessLevel) ApduDataExtPropertyDescriptionResponseBuilder
	// WithObjectIndex adds ObjectIndex (property field)
	WithObjectIndex(uint8) ApduDataExtPropertyDescriptionResponseBuilder
	// WithPropertyId adds PropertyId (property field)
	WithPropertyId(uint8) ApduDataExtPropertyDescriptionResponseBuilder
	// WithIndex adds Index (property field)
	WithIndex(uint8) ApduDataExtPropertyDescriptionResponseBuilder
	// WithWriteEnabled adds WriteEnabled (property field)
	WithWriteEnabled(bool) ApduDataExtPropertyDescriptionResponseBuilder
	// WithPropertyDataType adds PropertyDataType (property field)
	WithPropertyDataType(KnxPropertyDataType) ApduDataExtPropertyDescriptionResponseBuilder
	// WithMaxNrOfElements adds MaxNrOfElements (property field)
	WithMaxNrOfElements(uint16) ApduDataExtPropertyDescriptionResponseBuilder
	// WithReadLevel adds ReadLevel (property field)
	WithReadLevel(AccessLevel) ApduDataExtPropertyDescriptionResponseBuilder
	// WithWriteLevel adds WriteLevel (property field)
	WithWriteLevel(AccessLevel) ApduDataExtPropertyDescriptionResponseBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ApduDataExtBuilder
	// Build builds the ApduDataExtPropertyDescriptionResponse or returns an error if something is wrong
	Build() (ApduDataExtPropertyDescriptionResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ApduDataExtPropertyDescriptionResponse
}

// NewApduDataExtPropertyDescriptionResponseBuilder() creates a ApduDataExtPropertyDescriptionResponseBuilder
func NewApduDataExtPropertyDescriptionResponseBuilder() ApduDataExtPropertyDescriptionResponseBuilder {
	return &_ApduDataExtPropertyDescriptionResponseBuilder{_ApduDataExtPropertyDescriptionResponse: new(_ApduDataExtPropertyDescriptionResponse)}
}

type _ApduDataExtPropertyDescriptionResponseBuilder struct {
	*_ApduDataExtPropertyDescriptionResponse

	parentBuilder *_ApduDataExtBuilder

	err *utils.MultiError
}

var _ (ApduDataExtPropertyDescriptionResponseBuilder) = (*_ApduDataExtPropertyDescriptionResponseBuilder)(nil)

func (b *_ApduDataExtPropertyDescriptionResponseBuilder) setParent(contract ApduDataExtContract) {
	b.ApduDataExtContract = contract
}

func (b *_ApduDataExtPropertyDescriptionResponseBuilder) WithMandatoryFields(objectIndex uint8, propertyId uint8, index uint8, writeEnabled bool, propertyDataType KnxPropertyDataType, maxNrOfElements uint16, readLevel AccessLevel, writeLevel AccessLevel) ApduDataExtPropertyDescriptionResponseBuilder {
	return b.WithObjectIndex(objectIndex).WithPropertyId(propertyId).WithIndex(index).WithWriteEnabled(writeEnabled).WithPropertyDataType(propertyDataType).WithMaxNrOfElements(maxNrOfElements).WithReadLevel(readLevel).WithWriteLevel(writeLevel)
}

func (b *_ApduDataExtPropertyDescriptionResponseBuilder) WithObjectIndex(objectIndex uint8) ApduDataExtPropertyDescriptionResponseBuilder {
	b.ObjectIndex = objectIndex
	return b
}

func (b *_ApduDataExtPropertyDescriptionResponseBuilder) WithPropertyId(propertyId uint8) ApduDataExtPropertyDescriptionResponseBuilder {
	b.PropertyId = propertyId
	return b
}

func (b *_ApduDataExtPropertyDescriptionResponseBuilder) WithIndex(index uint8) ApduDataExtPropertyDescriptionResponseBuilder {
	b.Index = index
	return b
}

func (b *_ApduDataExtPropertyDescriptionResponseBuilder) WithWriteEnabled(writeEnabled bool) ApduDataExtPropertyDescriptionResponseBuilder {
	b.WriteEnabled = writeEnabled
	return b
}

func (b *_ApduDataExtPropertyDescriptionResponseBuilder) WithPropertyDataType(propertyDataType KnxPropertyDataType) ApduDataExtPropertyDescriptionResponseBuilder {
	b.PropertyDataType = propertyDataType
	return b
}

func (b *_ApduDataExtPropertyDescriptionResponseBuilder) WithMaxNrOfElements(maxNrOfElements uint16) ApduDataExtPropertyDescriptionResponseBuilder {
	b.MaxNrOfElements = maxNrOfElements
	return b
}

func (b *_ApduDataExtPropertyDescriptionResponseBuilder) WithReadLevel(readLevel AccessLevel) ApduDataExtPropertyDescriptionResponseBuilder {
	b.ReadLevel = readLevel
	return b
}

func (b *_ApduDataExtPropertyDescriptionResponseBuilder) WithWriteLevel(writeLevel AccessLevel) ApduDataExtPropertyDescriptionResponseBuilder {
	b.WriteLevel = writeLevel
	return b
}

func (b *_ApduDataExtPropertyDescriptionResponseBuilder) Build() (ApduDataExtPropertyDescriptionResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ApduDataExtPropertyDescriptionResponse.deepCopy(), nil
}

func (b *_ApduDataExtPropertyDescriptionResponseBuilder) MustBuild() ApduDataExtPropertyDescriptionResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ApduDataExtPropertyDescriptionResponseBuilder) Done() ApduDataExtBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewApduDataExtBuilder().(*_ApduDataExtBuilder)
	}
	return b.parentBuilder
}

func (b *_ApduDataExtPropertyDescriptionResponseBuilder) buildForApduDataExt() (ApduDataExt, error) {
	return b.Build()
}

func (b *_ApduDataExtPropertyDescriptionResponseBuilder) DeepCopy() any {
	_copy := b.CreateApduDataExtPropertyDescriptionResponseBuilder().(*_ApduDataExtPropertyDescriptionResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateApduDataExtPropertyDescriptionResponseBuilder creates a ApduDataExtPropertyDescriptionResponseBuilder
func (b *_ApduDataExtPropertyDescriptionResponse) CreateApduDataExtPropertyDescriptionResponseBuilder() ApduDataExtPropertyDescriptionResponseBuilder {
	if b == nil {
		return NewApduDataExtPropertyDescriptionResponseBuilder()
	}
	return &_ApduDataExtPropertyDescriptionResponseBuilder{_ApduDataExtPropertyDescriptionResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtPropertyDescriptionResponse) GetExtApciType() uint8 {
	return 0x19
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtPropertyDescriptionResponse) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ApduDataExtPropertyDescriptionResponse) GetObjectIndex() uint8 {
	return m.ObjectIndex
}

func (m *_ApduDataExtPropertyDescriptionResponse) GetPropertyId() uint8 {
	return m.PropertyId
}

func (m *_ApduDataExtPropertyDescriptionResponse) GetIndex() uint8 {
	return m.Index
}

func (m *_ApduDataExtPropertyDescriptionResponse) GetWriteEnabled() bool {
	return m.WriteEnabled
}

func (m *_ApduDataExtPropertyDescriptionResponse) GetPropertyDataType() KnxPropertyDataType {
	return m.PropertyDataType
}

func (m *_ApduDataExtPropertyDescriptionResponse) GetMaxNrOfElements() uint16 {
	return m.MaxNrOfElements
}

func (m *_ApduDataExtPropertyDescriptionResponse) GetReadLevel() AccessLevel {
	return m.ReadLevel
}

func (m *_ApduDataExtPropertyDescriptionResponse) GetWriteLevel() AccessLevel {
	return m.WriteLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastApduDataExtPropertyDescriptionResponse(structType any) ApduDataExtPropertyDescriptionResponse {
	if casted, ok := structType.(ApduDataExtPropertyDescriptionResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtPropertyDescriptionResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtPropertyDescriptionResponse) GetTypeName() string {
	return "ApduDataExtPropertyDescriptionResponse"
}

func (m *_ApduDataExtPropertyDescriptionResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	// Simple field (objectIndex)
	lengthInBits += 8

	// Simple field (propertyId)
	lengthInBits += 8

	// Simple field (index)
	lengthInBits += 8

	// Simple field (writeEnabled)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (propertyDataType)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 4

	// Simple field (maxNrOfElements)
	lengthInBits += 12

	// Simple field (readLevel)
	lengthInBits += 4

	// Simple field (writeLevel)
	lengthInBits += 4

	return lengthInBits
}

func (m *_ApduDataExtPropertyDescriptionResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtPropertyDescriptionResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (__apduDataExtPropertyDescriptionResponse ApduDataExtPropertyDescriptionResponse, err error) {
	m.ApduDataExtContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtPropertyDescriptionResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtPropertyDescriptionResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectIndex, err := ReadSimpleField(ctx, "objectIndex", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIndex' field"))
	}
	m.ObjectIndex = objectIndex

	propertyId, err := ReadSimpleField(ctx, "propertyId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyId' field"))
	}
	m.PropertyId = propertyId

	index, err := ReadSimpleField(ctx, "index", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'index' field"))
	}
	m.Index = index

	writeEnabled, err := ReadSimpleField(ctx, "writeEnabled", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'writeEnabled' field"))
	}
	m.WriteEnabled = writeEnabled

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(1)), uint8(0x0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	propertyDataType, err := ReadEnumField[KnxPropertyDataType](ctx, "propertyDataType", "KnxPropertyDataType", ReadEnum(KnxPropertyDataTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyDataType' field"))
	}
	m.PropertyDataType = propertyDataType

	reservedField1, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(4)), uint8(0x0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField1 = reservedField1

	maxNrOfElements, err := ReadSimpleField(ctx, "maxNrOfElements", ReadUnsignedShort(readBuffer, uint8(12)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxNrOfElements' field"))
	}
	m.MaxNrOfElements = maxNrOfElements

	readLevel, err := ReadEnumField[AccessLevel](ctx, "readLevel", "AccessLevel", ReadEnum(AccessLevelByValue, ReadUnsignedByte(readBuffer, uint8(4))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'readLevel' field"))
	}
	m.ReadLevel = readLevel

	writeLevel, err := ReadEnumField[AccessLevel](ctx, "writeLevel", "AccessLevel", ReadEnum(AccessLevelByValue, ReadUnsignedByte(readBuffer, uint8(4))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'writeLevel' field"))
	}
	m.WriteLevel = writeLevel

	if closeErr := readBuffer.CloseContext("ApduDataExtPropertyDescriptionResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtPropertyDescriptionResponse")
	}

	return m, nil
}

func (m *_ApduDataExtPropertyDescriptionResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtPropertyDescriptionResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtPropertyDescriptionResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtPropertyDescriptionResponse")
		}

		if err := WriteSimpleField[uint8](ctx, "objectIndex", m.GetObjectIndex(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'objectIndex' field")
		}

		if err := WriteSimpleField[uint8](ctx, "propertyId", m.GetPropertyId(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'propertyId' field")
		}

		if err := WriteSimpleField[uint8](ctx, "index", m.GetIndex(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'index' field")
		}

		if err := WriteSimpleField[bool](ctx, "writeEnabled", m.GetWriteEnabled(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'writeEnabled' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x0), WriteUnsignedByte(writeBuffer, 1)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleEnumField[KnxPropertyDataType](ctx, "propertyDataType", "KnxPropertyDataType", m.GetPropertyDataType(), WriteEnum[KnxPropertyDataType, uint8](KnxPropertyDataType.GetValue, KnxPropertyDataType.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'propertyDataType' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x0), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 2")
		}

		if err := WriteSimpleField[uint16](ctx, "maxNrOfElements", m.GetMaxNrOfElements(), WriteUnsignedShort(writeBuffer, 12)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxNrOfElements' field")
		}

		if err := WriteSimpleEnumField[AccessLevel](ctx, "readLevel", "AccessLevel", m.GetReadLevel(), WriteEnum[AccessLevel, uint8](AccessLevel.GetValue, AccessLevel.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 4))); err != nil {
			return errors.Wrap(err, "Error serializing 'readLevel' field")
		}

		if err := WriteSimpleEnumField[AccessLevel](ctx, "writeLevel", "AccessLevel", m.GetWriteLevel(), WriteEnum[AccessLevel, uint8](AccessLevel.GetValue, AccessLevel.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 4))); err != nil {
			return errors.Wrap(err, "Error serializing 'writeLevel' field")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtPropertyDescriptionResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtPropertyDescriptionResponse")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtPropertyDescriptionResponse) IsApduDataExtPropertyDescriptionResponse() {}

func (m *_ApduDataExtPropertyDescriptionResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataExtPropertyDescriptionResponse) deepCopy() *_ApduDataExtPropertyDescriptionResponse {
	if m == nil {
		return nil
	}
	_ApduDataExtPropertyDescriptionResponseCopy := &_ApduDataExtPropertyDescriptionResponse{
		m.ApduDataExtContract.(*_ApduDataExt).deepCopy(),
		m.ObjectIndex,
		m.PropertyId,
		m.Index,
		m.WriteEnabled,
		m.PropertyDataType,
		m.MaxNrOfElements,
		m.ReadLevel,
		m.WriteLevel,
		m.reservedField0,
		m.reservedField1,
	}
	m.ApduDataExtContract.(*_ApduDataExt)._SubType = m
	return _ApduDataExtPropertyDescriptionResponseCopy
}

func (m *_ApduDataExtPropertyDescriptionResponse) String() string {
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
