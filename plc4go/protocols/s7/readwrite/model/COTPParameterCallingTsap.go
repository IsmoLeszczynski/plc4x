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

// COTPParameterCallingTsap is the corresponding interface of COTPParameterCallingTsap
type COTPParameterCallingTsap interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	COTPParameter
	// GetTsapId returns TsapId (property field)
	GetTsapId() uint16
	// IsCOTPParameterCallingTsap is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCOTPParameterCallingTsap()
	// CreateBuilder creates a COTPParameterCallingTsapBuilder
	CreateCOTPParameterCallingTsapBuilder() COTPParameterCallingTsapBuilder
}

// _COTPParameterCallingTsap is the data-structure of this message
type _COTPParameterCallingTsap struct {
	COTPParameterContract
	TsapId uint16
}

var _ COTPParameterCallingTsap = (*_COTPParameterCallingTsap)(nil)
var _ COTPParameterRequirements = (*_COTPParameterCallingTsap)(nil)

// NewCOTPParameterCallingTsap factory function for _COTPParameterCallingTsap
func NewCOTPParameterCallingTsap(tsapId uint16, rest uint8) *_COTPParameterCallingTsap {
	_result := &_COTPParameterCallingTsap{
		COTPParameterContract: NewCOTPParameter(rest),
		TsapId:                tsapId,
	}
	_result.COTPParameterContract.(*_COTPParameter)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// COTPParameterCallingTsapBuilder is a builder for COTPParameterCallingTsap
type COTPParameterCallingTsapBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(tsapId uint16) COTPParameterCallingTsapBuilder
	// WithTsapId adds TsapId (property field)
	WithTsapId(uint16) COTPParameterCallingTsapBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() COTPParameterBuilder
	// Build builds the COTPParameterCallingTsap or returns an error if something is wrong
	Build() (COTPParameterCallingTsap, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() COTPParameterCallingTsap
}

// NewCOTPParameterCallingTsapBuilder() creates a COTPParameterCallingTsapBuilder
func NewCOTPParameterCallingTsapBuilder() COTPParameterCallingTsapBuilder {
	return &_COTPParameterCallingTsapBuilder{_COTPParameterCallingTsap: new(_COTPParameterCallingTsap)}
}

type _COTPParameterCallingTsapBuilder struct {
	*_COTPParameterCallingTsap

	parentBuilder *_COTPParameterBuilder

	err *utils.MultiError
}

var _ (COTPParameterCallingTsapBuilder) = (*_COTPParameterCallingTsapBuilder)(nil)

func (b *_COTPParameterCallingTsapBuilder) setParent(contract COTPParameterContract) {
	b.COTPParameterContract = contract
	contract.(*_COTPParameter)._SubType = b._COTPParameterCallingTsap
}

func (b *_COTPParameterCallingTsapBuilder) WithMandatoryFields(tsapId uint16) COTPParameterCallingTsapBuilder {
	return b.WithTsapId(tsapId)
}

func (b *_COTPParameterCallingTsapBuilder) WithTsapId(tsapId uint16) COTPParameterCallingTsapBuilder {
	b.TsapId = tsapId
	return b
}

func (b *_COTPParameterCallingTsapBuilder) Build() (COTPParameterCallingTsap, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._COTPParameterCallingTsap.deepCopy(), nil
}

func (b *_COTPParameterCallingTsapBuilder) MustBuild() COTPParameterCallingTsap {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_COTPParameterCallingTsapBuilder) Done() COTPParameterBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewCOTPParameterBuilder().(*_COTPParameterBuilder)
	}
	return b.parentBuilder
}

func (b *_COTPParameterCallingTsapBuilder) buildForCOTPParameter() (COTPParameter, error) {
	return b.Build()
}

func (b *_COTPParameterCallingTsapBuilder) DeepCopy() any {
	_copy := b.CreateCOTPParameterCallingTsapBuilder().(*_COTPParameterCallingTsapBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCOTPParameterCallingTsapBuilder creates a COTPParameterCallingTsapBuilder
func (b *_COTPParameterCallingTsap) CreateCOTPParameterCallingTsapBuilder() COTPParameterCallingTsapBuilder {
	if b == nil {
		return NewCOTPParameterCallingTsapBuilder()
	}
	return &_COTPParameterCallingTsapBuilder{_COTPParameterCallingTsap: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_COTPParameterCallingTsap) GetParameterType() uint8 {
	return 0xC1
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_COTPParameterCallingTsap) GetParent() COTPParameterContract {
	return m.COTPParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_COTPParameterCallingTsap) GetTsapId() uint16 {
	return m.TsapId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCOTPParameterCallingTsap(structType any) COTPParameterCallingTsap {
	if casted, ok := structType.(COTPParameterCallingTsap); ok {
		return casted
	}
	if casted, ok := structType.(*COTPParameterCallingTsap); ok {
		return *casted
	}
	return nil
}

func (m *_COTPParameterCallingTsap) GetTypeName() string {
	return "COTPParameterCallingTsap"
}

func (m *_COTPParameterCallingTsap) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.COTPParameterContract.(*_COTPParameter).getLengthInBits(ctx))

	// Simple field (tsapId)
	lengthInBits += 16

	return lengthInBits
}

func (m *_COTPParameterCallingTsap) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_COTPParameterCallingTsap) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_COTPParameter, rest uint8) (__cOTPParameterCallingTsap COTPParameterCallingTsap, err error) {
	m.COTPParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("COTPParameterCallingTsap"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for COTPParameterCallingTsap")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	tsapId, err := ReadSimpleField(ctx, "tsapId", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'tsapId' field"))
	}
	m.TsapId = tsapId

	if closeErr := readBuffer.CloseContext("COTPParameterCallingTsap"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for COTPParameterCallingTsap")
	}

	return m, nil
}

func (m *_COTPParameterCallingTsap) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_COTPParameterCallingTsap) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("COTPParameterCallingTsap"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for COTPParameterCallingTsap")
		}

		if err := WriteSimpleField[uint16](ctx, "tsapId", m.GetTsapId(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'tsapId' field")
		}

		if popErr := writeBuffer.PopContext("COTPParameterCallingTsap"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for COTPParameterCallingTsap")
		}
		return nil
	}
	return m.COTPParameterContract.(*_COTPParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_COTPParameterCallingTsap) IsCOTPParameterCallingTsap() {}

func (m *_COTPParameterCallingTsap) DeepCopy() any {
	return m.deepCopy()
}

func (m *_COTPParameterCallingTsap) deepCopy() *_COTPParameterCallingTsap {
	if m == nil {
		return nil
	}
	_COTPParameterCallingTsapCopy := &_COTPParameterCallingTsap{
		m.COTPParameterContract.(*_COTPParameter).deepCopy(),
		m.TsapId,
	}
	_COTPParameterCallingTsapCopy.COTPParameterContract.(*_COTPParameter)._SubType = m
	return _COTPParameterCallingTsapCopy
}

func (m *_COTPParameterCallingTsap) String() string {
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
