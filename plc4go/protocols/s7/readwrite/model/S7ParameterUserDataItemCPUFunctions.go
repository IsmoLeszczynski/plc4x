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

// S7ParameterUserDataItemCPUFunctions is the corresponding interface of S7ParameterUserDataItemCPUFunctions
type S7ParameterUserDataItemCPUFunctions interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	S7ParameterUserDataItem
	// GetMethod returns Method (property field)
	GetMethod() uint8
	// GetCpuFunctionType returns CpuFunctionType (property field)
	GetCpuFunctionType() uint8
	// GetCpuFunctionGroup returns CpuFunctionGroup (property field)
	GetCpuFunctionGroup() uint8
	// GetCpuSubfunction returns CpuSubfunction (property field)
	GetCpuSubfunction() uint8
	// GetSequenceNumber returns SequenceNumber (property field)
	GetSequenceNumber() uint8
	// GetDataUnitReferenceNumber returns DataUnitReferenceNumber (property field)
	GetDataUnitReferenceNumber() *uint8
	// GetLastDataUnit returns LastDataUnit (property field)
	GetLastDataUnit() *uint8
	// GetErrorCode returns ErrorCode (property field)
	GetErrorCode() *uint16
	// IsS7ParameterUserDataItemCPUFunctions is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7ParameterUserDataItemCPUFunctions()
	// CreateBuilder creates a S7ParameterUserDataItemCPUFunctionsBuilder
	CreateS7ParameterUserDataItemCPUFunctionsBuilder() S7ParameterUserDataItemCPUFunctionsBuilder
}

// _S7ParameterUserDataItemCPUFunctions is the data-structure of this message
type _S7ParameterUserDataItemCPUFunctions struct {
	S7ParameterUserDataItemContract
	Method                  uint8
	CpuFunctionType         uint8
	CpuFunctionGroup        uint8
	CpuSubfunction          uint8
	SequenceNumber          uint8
	DataUnitReferenceNumber *uint8
	LastDataUnit            *uint8
	ErrorCode               *uint16
}

var _ S7ParameterUserDataItemCPUFunctions = (*_S7ParameterUserDataItemCPUFunctions)(nil)
var _ S7ParameterUserDataItemRequirements = (*_S7ParameterUserDataItemCPUFunctions)(nil)

// NewS7ParameterUserDataItemCPUFunctions factory function for _S7ParameterUserDataItemCPUFunctions
func NewS7ParameterUserDataItemCPUFunctions(method uint8, cpuFunctionType uint8, cpuFunctionGroup uint8, cpuSubfunction uint8, sequenceNumber uint8, dataUnitReferenceNumber *uint8, lastDataUnit *uint8, errorCode *uint16) *_S7ParameterUserDataItemCPUFunctions {
	_result := &_S7ParameterUserDataItemCPUFunctions{
		S7ParameterUserDataItemContract: NewS7ParameterUserDataItem(),
		Method:                          method,
		CpuFunctionType:                 cpuFunctionType,
		CpuFunctionGroup:                cpuFunctionGroup,
		CpuSubfunction:                  cpuSubfunction,
		SequenceNumber:                  sequenceNumber,
		DataUnitReferenceNumber:         dataUnitReferenceNumber,
		LastDataUnit:                    lastDataUnit,
		ErrorCode:                       errorCode,
	}
	_result.S7ParameterUserDataItemContract.(*_S7ParameterUserDataItem)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// S7ParameterUserDataItemCPUFunctionsBuilder is a builder for S7ParameterUserDataItemCPUFunctions
type S7ParameterUserDataItemCPUFunctionsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(method uint8, cpuFunctionType uint8, cpuFunctionGroup uint8, cpuSubfunction uint8, sequenceNumber uint8) S7ParameterUserDataItemCPUFunctionsBuilder
	// WithMethod adds Method (property field)
	WithMethod(uint8) S7ParameterUserDataItemCPUFunctionsBuilder
	// WithCpuFunctionType adds CpuFunctionType (property field)
	WithCpuFunctionType(uint8) S7ParameterUserDataItemCPUFunctionsBuilder
	// WithCpuFunctionGroup adds CpuFunctionGroup (property field)
	WithCpuFunctionGroup(uint8) S7ParameterUserDataItemCPUFunctionsBuilder
	// WithCpuSubfunction adds CpuSubfunction (property field)
	WithCpuSubfunction(uint8) S7ParameterUserDataItemCPUFunctionsBuilder
	// WithSequenceNumber adds SequenceNumber (property field)
	WithSequenceNumber(uint8) S7ParameterUserDataItemCPUFunctionsBuilder
	// WithDataUnitReferenceNumber adds DataUnitReferenceNumber (property field)
	WithOptionalDataUnitReferenceNumber(uint8) S7ParameterUserDataItemCPUFunctionsBuilder
	// WithLastDataUnit adds LastDataUnit (property field)
	WithOptionalLastDataUnit(uint8) S7ParameterUserDataItemCPUFunctionsBuilder
	// WithErrorCode adds ErrorCode (property field)
	WithOptionalErrorCode(uint16) S7ParameterUserDataItemCPUFunctionsBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() S7ParameterUserDataItemBuilder
	// Build builds the S7ParameterUserDataItemCPUFunctions or returns an error if something is wrong
	Build() (S7ParameterUserDataItemCPUFunctions, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() S7ParameterUserDataItemCPUFunctions
}

// NewS7ParameterUserDataItemCPUFunctionsBuilder() creates a S7ParameterUserDataItemCPUFunctionsBuilder
func NewS7ParameterUserDataItemCPUFunctionsBuilder() S7ParameterUserDataItemCPUFunctionsBuilder {
	return &_S7ParameterUserDataItemCPUFunctionsBuilder{_S7ParameterUserDataItemCPUFunctions: new(_S7ParameterUserDataItemCPUFunctions)}
}

type _S7ParameterUserDataItemCPUFunctionsBuilder struct {
	*_S7ParameterUserDataItemCPUFunctions

	parentBuilder *_S7ParameterUserDataItemBuilder

	err *utils.MultiError
}

var _ (S7ParameterUserDataItemCPUFunctionsBuilder) = (*_S7ParameterUserDataItemCPUFunctionsBuilder)(nil)

func (b *_S7ParameterUserDataItemCPUFunctionsBuilder) setParent(contract S7ParameterUserDataItemContract) {
	b.S7ParameterUserDataItemContract = contract
}

func (b *_S7ParameterUserDataItemCPUFunctionsBuilder) WithMandatoryFields(method uint8, cpuFunctionType uint8, cpuFunctionGroup uint8, cpuSubfunction uint8, sequenceNumber uint8) S7ParameterUserDataItemCPUFunctionsBuilder {
	return b.WithMethod(method).WithCpuFunctionType(cpuFunctionType).WithCpuFunctionGroup(cpuFunctionGroup).WithCpuSubfunction(cpuSubfunction).WithSequenceNumber(sequenceNumber)
}

func (b *_S7ParameterUserDataItemCPUFunctionsBuilder) WithMethod(method uint8) S7ParameterUserDataItemCPUFunctionsBuilder {
	b.Method = method
	return b
}

func (b *_S7ParameterUserDataItemCPUFunctionsBuilder) WithCpuFunctionType(cpuFunctionType uint8) S7ParameterUserDataItemCPUFunctionsBuilder {
	b.CpuFunctionType = cpuFunctionType
	return b
}

func (b *_S7ParameterUserDataItemCPUFunctionsBuilder) WithCpuFunctionGroup(cpuFunctionGroup uint8) S7ParameterUserDataItemCPUFunctionsBuilder {
	b.CpuFunctionGroup = cpuFunctionGroup
	return b
}

func (b *_S7ParameterUserDataItemCPUFunctionsBuilder) WithCpuSubfunction(cpuSubfunction uint8) S7ParameterUserDataItemCPUFunctionsBuilder {
	b.CpuSubfunction = cpuSubfunction
	return b
}

func (b *_S7ParameterUserDataItemCPUFunctionsBuilder) WithSequenceNumber(sequenceNumber uint8) S7ParameterUserDataItemCPUFunctionsBuilder {
	b.SequenceNumber = sequenceNumber
	return b
}

func (b *_S7ParameterUserDataItemCPUFunctionsBuilder) WithOptionalDataUnitReferenceNumber(dataUnitReferenceNumber uint8) S7ParameterUserDataItemCPUFunctionsBuilder {
	b.DataUnitReferenceNumber = &dataUnitReferenceNumber
	return b
}

func (b *_S7ParameterUserDataItemCPUFunctionsBuilder) WithOptionalLastDataUnit(lastDataUnit uint8) S7ParameterUserDataItemCPUFunctionsBuilder {
	b.LastDataUnit = &lastDataUnit
	return b
}

func (b *_S7ParameterUserDataItemCPUFunctionsBuilder) WithOptionalErrorCode(errorCode uint16) S7ParameterUserDataItemCPUFunctionsBuilder {
	b.ErrorCode = &errorCode
	return b
}

func (b *_S7ParameterUserDataItemCPUFunctionsBuilder) Build() (S7ParameterUserDataItemCPUFunctions, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._S7ParameterUserDataItemCPUFunctions.deepCopy(), nil
}

func (b *_S7ParameterUserDataItemCPUFunctionsBuilder) MustBuild() S7ParameterUserDataItemCPUFunctions {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_S7ParameterUserDataItemCPUFunctionsBuilder) Done() S7ParameterUserDataItemBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewS7ParameterUserDataItemBuilder().(*_S7ParameterUserDataItemBuilder)
	}
	return b.parentBuilder
}

func (b *_S7ParameterUserDataItemCPUFunctionsBuilder) buildForS7ParameterUserDataItem() (S7ParameterUserDataItem, error) {
	return b.Build()
}

func (b *_S7ParameterUserDataItemCPUFunctionsBuilder) DeepCopy() any {
	_copy := b.CreateS7ParameterUserDataItemCPUFunctionsBuilder().(*_S7ParameterUserDataItemCPUFunctionsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateS7ParameterUserDataItemCPUFunctionsBuilder creates a S7ParameterUserDataItemCPUFunctionsBuilder
func (b *_S7ParameterUserDataItemCPUFunctions) CreateS7ParameterUserDataItemCPUFunctionsBuilder() S7ParameterUserDataItemCPUFunctionsBuilder {
	if b == nil {
		return NewS7ParameterUserDataItemCPUFunctionsBuilder()
	}
	return &_S7ParameterUserDataItemCPUFunctionsBuilder{_S7ParameterUserDataItemCPUFunctions: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7ParameterUserDataItemCPUFunctions) GetItemType() uint8 {
	return 0x12
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7ParameterUserDataItemCPUFunctions) GetParent() S7ParameterUserDataItemContract {
	return m.S7ParameterUserDataItemContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7ParameterUserDataItemCPUFunctions) GetMethod() uint8 {
	return m.Method
}

func (m *_S7ParameterUserDataItemCPUFunctions) GetCpuFunctionType() uint8 {
	return m.CpuFunctionType
}

func (m *_S7ParameterUserDataItemCPUFunctions) GetCpuFunctionGroup() uint8 {
	return m.CpuFunctionGroup
}

func (m *_S7ParameterUserDataItemCPUFunctions) GetCpuSubfunction() uint8 {
	return m.CpuSubfunction
}

func (m *_S7ParameterUserDataItemCPUFunctions) GetSequenceNumber() uint8 {
	return m.SequenceNumber
}

func (m *_S7ParameterUserDataItemCPUFunctions) GetDataUnitReferenceNumber() *uint8 {
	return m.DataUnitReferenceNumber
}

func (m *_S7ParameterUserDataItemCPUFunctions) GetLastDataUnit() *uint8 {
	return m.LastDataUnit
}

func (m *_S7ParameterUserDataItemCPUFunctions) GetErrorCode() *uint16 {
	return m.ErrorCode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastS7ParameterUserDataItemCPUFunctions(structType any) S7ParameterUserDataItemCPUFunctions {
	if casted, ok := structType.(S7ParameterUserDataItemCPUFunctions); ok {
		return casted
	}
	if casted, ok := structType.(*S7ParameterUserDataItemCPUFunctions); ok {
		return *casted
	}
	return nil
}

func (m *_S7ParameterUserDataItemCPUFunctions) GetTypeName() string {
	return "S7ParameterUserDataItemCPUFunctions"
}

func (m *_S7ParameterUserDataItemCPUFunctions) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7ParameterUserDataItemContract.(*_S7ParameterUserDataItem).getLengthInBits(ctx))

	// Implicit Field (itemLength)
	lengthInBits += 8

	// Simple field (method)
	lengthInBits += 8

	// Simple field (cpuFunctionType)
	lengthInBits += 4

	// Simple field (cpuFunctionGroup)
	lengthInBits += 4

	// Simple field (cpuSubfunction)
	lengthInBits += 8

	// Simple field (sequenceNumber)
	lengthInBits += 8

	// Optional Field (dataUnitReferenceNumber)
	if m.DataUnitReferenceNumber != nil {
		lengthInBits += 8
	}

	// Optional Field (lastDataUnit)
	if m.LastDataUnit != nil {
		lengthInBits += 8
	}

	// Optional Field (errorCode)
	if m.ErrorCode != nil {
		lengthInBits += 16
	}

	return lengthInBits
}

func (m *_S7ParameterUserDataItemCPUFunctions) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7ParameterUserDataItemCPUFunctions) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7ParameterUserDataItem) (__s7ParameterUserDataItemCPUFunctions S7ParameterUserDataItemCPUFunctions, err error) {
	m.S7ParameterUserDataItemContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7ParameterUserDataItemCPUFunctions"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7ParameterUserDataItemCPUFunctions")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	itemLength, err := ReadImplicitField[uint8](ctx, "itemLength", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemLength' field"))
	}
	_ = itemLength

	method, err := ReadSimpleField(ctx, "method", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'method' field"))
	}
	m.Method = method

	cpuFunctionType, err := ReadSimpleField(ctx, "cpuFunctionType", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'cpuFunctionType' field"))
	}
	m.CpuFunctionType = cpuFunctionType

	cpuFunctionGroup, err := ReadSimpleField(ctx, "cpuFunctionGroup", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'cpuFunctionGroup' field"))
	}
	m.CpuFunctionGroup = cpuFunctionGroup

	cpuSubfunction, err := ReadSimpleField(ctx, "cpuSubfunction", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'cpuSubfunction' field"))
	}
	m.CpuSubfunction = cpuSubfunction

	sequenceNumber, err := ReadSimpleField(ctx, "sequenceNumber", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sequenceNumber' field"))
	}
	m.SequenceNumber = sequenceNumber

	var dataUnitReferenceNumber *uint8
	dataUnitReferenceNumber, err = ReadOptionalField[uint8](ctx, "dataUnitReferenceNumber", ReadUnsignedByte(readBuffer, uint8(8)), bool((bool((cpuFunctionType) == (8)))) || bool((bool((bool((cpuFunctionType) == (0)))) && bool((bool((cpuFunctionGroup) == (2)))))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataUnitReferenceNumber' field"))
	}
	m.DataUnitReferenceNumber = dataUnitReferenceNumber

	var lastDataUnit *uint8
	lastDataUnit, err = ReadOptionalField[uint8](ctx, "lastDataUnit", ReadUnsignedByte(readBuffer, uint8(8)), bool((bool((cpuFunctionType) == (8)))) || bool((bool((bool((cpuFunctionType) == (0)))) && bool((bool((cpuFunctionGroup) == (2)))))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastDataUnit' field"))
	}
	m.LastDataUnit = lastDataUnit

	var errorCode *uint16
	errorCode, err = ReadOptionalField[uint16](ctx, "errorCode", ReadUnsignedShort(readBuffer, uint8(16)), bool((bool((cpuFunctionType) == (8)))) || bool((bool((bool((cpuFunctionType) == (0)))) && bool((bool((cpuFunctionGroup) == (2)))))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'errorCode' field"))
	}
	m.ErrorCode = errorCode

	if closeErr := readBuffer.CloseContext("S7ParameterUserDataItemCPUFunctions"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7ParameterUserDataItemCPUFunctions")
	}

	return m, nil
}

func (m *_S7ParameterUserDataItemCPUFunctions) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7ParameterUserDataItemCPUFunctions) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7ParameterUserDataItemCPUFunctions"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7ParameterUserDataItemCPUFunctions")
		}
		itemLength := uint8(uint8(uint8(m.GetLengthInBytes(ctx))) - uint8(uint8(2)))
		if err := WriteImplicitField(ctx, "itemLength", itemLength, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'itemLength' field")
		}

		if err := WriteSimpleField[uint8](ctx, "method", m.GetMethod(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'method' field")
		}

		if err := WriteSimpleField[uint8](ctx, "cpuFunctionType", m.GetCpuFunctionType(), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'cpuFunctionType' field")
		}

		if err := WriteSimpleField[uint8](ctx, "cpuFunctionGroup", m.GetCpuFunctionGroup(), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'cpuFunctionGroup' field")
		}

		if err := WriteSimpleField[uint8](ctx, "cpuSubfunction", m.GetCpuSubfunction(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'cpuSubfunction' field")
		}

		if err := WriteSimpleField[uint8](ctx, "sequenceNumber", m.GetSequenceNumber(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'sequenceNumber' field")
		}

		if err := WriteOptionalField[uint8](ctx, "dataUnitReferenceNumber", m.GetDataUnitReferenceNumber(), WriteUnsignedByte(writeBuffer, 8), true); err != nil {
			return errors.Wrap(err, "Error serializing 'dataUnitReferenceNumber' field")
		}

		if err := WriteOptionalField[uint8](ctx, "lastDataUnit", m.GetLastDataUnit(), WriteUnsignedByte(writeBuffer, 8), true); err != nil {
			return errors.Wrap(err, "Error serializing 'lastDataUnit' field")
		}

		if err := WriteOptionalField[uint16](ctx, "errorCode", m.GetErrorCode(), WriteUnsignedShort(writeBuffer, 16), true); err != nil {
			return errors.Wrap(err, "Error serializing 'errorCode' field")
		}

		if popErr := writeBuffer.PopContext("S7ParameterUserDataItemCPUFunctions"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7ParameterUserDataItemCPUFunctions")
		}
		return nil
	}
	return m.S7ParameterUserDataItemContract.(*_S7ParameterUserDataItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7ParameterUserDataItemCPUFunctions) IsS7ParameterUserDataItemCPUFunctions() {}

func (m *_S7ParameterUserDataItemCPUFunctions) DeepCopy() any {
	return m.deepCopy()
}

func (m *_S7ParameterUserDataItemCPUFunctions) deepCopy() *_S7ParameterUserDataItemCPUFunctions {
	if m == nil {
		return nil
	}
	_S7ParameterUserDataItemCPUFunctionsCopy := &_S7ParameterUserDataItemCPUFunctions{
		m.S7ParameterUserDataItemContract.(*_S7ParameterUserDataItem).deepCopy(),
		m.Method,
		m.CpuFunctionType,
		m.CpuFunctionGroup,
		m.CpuSubfunction,
		m.SequenceNumber,
		utils.CopyPtr[uint8](m.DataUnitReferenceNumber),
		utils.CopyPtr[uint8](m.LastDataUnit),
		utils.CopyPtr[uint16](m.ErrorCode),
	}
	m.S7ParameterUserDataItemContract.(*_S7ParameterUserDataItem)._SubType = m
	return _S7ParameterUserDataItemCPUFunctionsCopy
}

func (m *_S7ParameterUserDataItemCPUFunctions) String() string {
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
