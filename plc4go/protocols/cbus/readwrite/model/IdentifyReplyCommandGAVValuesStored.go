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

// IdentifyReplyCommandGAVValuesStored is the corresponding interface of IdentifyReplyCommandGAVValuesStored
type IdentifyReplyCommandGAVValuesStored interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	IdentifyReplyCommand
	// GetValues returns Values (property field)
	GetValues() []byte
	// IsIdentifyReplyCommandGAVValuesStored is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsIdentifyReplyCommandGAVValuesStored()
	// CreateBuilder creates a IdentifyReplyCommandGAVValuesStoredBuilder
	CreateIdentifyReplyCommandGAVValuesStoredBuilder() IdentifyReplyCommandGAVValuesStoredBuilder
}

// _IdentifyReplyCommandGAVValuesStored is the data-structure of this message
type _IdentifyReplyCommandGAVValuesStored struct {
	IdentifyReplyCommandContract
	Values []byte
}

var _ IdentifyReplyCommandGAVValuesStored = (*_IdentifyReplyCommandGAVValuesStored)(nil)
var _ IdentifyReplyCommandRequirements = (*_IdentifyReplyCommandGAVValuesStored)(nil)

// NewIdentifyReplyCommandGAVValuesStored factory function for _IdentifyReplyCommandGAVValuesStored
func NewIdentifyReplyCommandGAVValuesStored(values []byte, numBytes uint8) *_IdentifyReplyCommandGAVValuesStored {
	_result := &_IdentifyReplyCommandGAVValuesStored{
		IdentifyReplyCommandContract: NewIdentifyReplyCommand(numBytes),
		Values:                       values,
	}
	_result.IdentifyReplyCommandContract.(*_IdentifyReplyCommand)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// IdentifyReplyCommandGAVValuesStoredBuilder is a builder for IdentifyReplyCommandGAVValuesStored
type IdentifyReplyCommandGAVValuesStoredBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(values []byte) IdentifyReplyCommandGAVValuesStoredBuilder
	// WithValues adds Values (property field)
	WithValues(...byte) IdentifyReplyCommandGAVValuesStoredBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() IdentifyReplyCommandBuilder
	// Build builds the IdentifyReplyCommandGAVValuesStored or returns an error if something is wrong
	Build() (IdentifyReplyCommandGAVValuesStored, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() IdentifyReplyCommandGAVValuesStored
}

// NewIdentifyReplyCommandGAVValuesStoredBuilder() creates a IdentifyReplyCommandGAVValuesStoredBuilder
func NewIdentifyReplyCommandGAVValuesStoredBuilder() IdentifyReplyCommandGAVValuesStoredBuilder {
	return &_IdentifyReplyCommandGAVValuesStoredBuilder{_IdentifyReplyCommandGAVValuesStored: new(_IdentifyReplyCommandGAVValuesStored)}
}

type _IdentifyReplyCommandGAVValuesStoredBuilder struct {
	*_IdentifyReplyCommandGAVValuesStored

	parentBuilder *_IdentifyReplyCommandBuilder

	err *utils.MultiError
}

var _ (IdentifyReplyCommandGAVValuesStoredBuilder) = (*_IdentifyReplyCommandGAVValuesStoredBuilder)(nil)

func (b *_IdentifyReplyCommandGAVValuesStoredBuilder) setParent(contract IdentifyReplyCommandContract) {
	b.IdentifyReplyCommandContract = contract
}

func (b *_IdentifyReplyCommandGAVValuesStoredBuilder) WithMandatoryFields(values []byte) IdentifyReplyCommandGAVValuesStoredBuilder {
	return b.WithValues(values...)
}

func (b *_IdentifyReplyCommandGAVValuesStoredBuilder) WithValues(values ...byte) IdentifyReplyCommandGAVValuesStoredBuilder {
	b.Values = values
	return b
}

func (b *_IdentifyReplyCommandGAVValuesStoredBuilder) Build() (IdentifyReplyCommandGAVValuesStored, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._IdentifyReplyCommandGAVValuesStored.deepCopy(), nil
}

func (b *_IdentifyReplyCommandGAVValuesStoredBuilder) MustBuild() IdentifyReplyCommandGAVValuesStored {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_IdentifyReplyCommandGAVValuesStoredBuilder) Done() IdentifyReplyCommandBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewIdentifyReplyCommandBuilder().(*_IdentifyReplyCommandBuilder)
	}
	return b.parentBuilder
}

func (b *_IdentifyReplyCommandGAVValuesStoredBuilder) buildForIdentifyReplyCommand() (IdentifyReplyCommand, error) {
	return b.Build()
}

func (b *_IdentifyReplyCommandGAVValuesStoredBuilder) DeepCopy() any {
	_copy := b.CreateIdentifyReplyCommandGAVValuesStoredBuilder().(*_IdentifyReplyCommandGAVValuesStoredBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateIdentifyReplyCommandGAVValuesStoredBuilder creates a IdentifyReplyCommandGAVValuesStoredBuilder
func (b *_IdentifyReplyCommandGAVValuesStored) CreateIdentifyReplyCommandGAVValuesStoredBuilder() IdentifyReplyCommandGAVValuesStoredBuilder {
	if b == nil {
		return NewIdentifyReplyCommandGAVValuesStoredBuilder()
	}
	return &_IdentifyReplyCommandGAVValuesStoredBuilder{_IdentifyReplyCommandGAVValuesStored: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandGAVValuesStored) GetAttribute() Attribute {
	return Attribute_GAVValuesStored
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandGAVValuesStored) GetParent() IdentifyReplyCommandContract {
	return m.IdentifyReplyCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandGAVValuesStored) GetValues() []byte {
	return m.Values
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandGAVValuesStored(structType any) IdentifyReplyCommandGAVValuesStored {
	if casted, ok := structType.(IdentifyReplyCommandGAVValuesStored); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandGAVValuesStored); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandGAVValuesStored) GetTypeName() string {
	return "IdentifyReplyCommandGAVValuesStored"
}

func (m *_IdentifyReplyCommandGAVValuesStored) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).getLengthInBits(ctx))

	// Array field
	if len(m.Values) > 0 {
		lengthInBits += 8 * uint16(len(m.Values))
	}

	return lengthInBits
}

func (m *_IdentifyReplyCommandGAVValuesStored) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_IdentifyReplyCommandGAVValuesStored) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_IdentifyReplyCommand, attribute Attribute, numBytes uint8) (__identifyReplyCommandGAVValuesStored IdentifyReplyCommandGAVValuesStored, err error) {
	m.IdentifyReplyCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandGAVValuesStored"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandGAVValuesStored")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	values, err := readBuffer.ReadByteArray("values", int(numBytes))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'values' field"))
	}
	m.Values = values

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandGAVValuesStored"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandGAVValuesStored")
	}

	return m, nil
}

func (m *_IdentifyReplyCommandGAVValuesStored) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IdentifyReplyCommandGAVValuesStored) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandGAVValuesStored"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandGAVValuesStored")
		}

		if err := WriteByteArrayField(ctx, "values", m.GetValues(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'values' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandGAVValuesStored"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandGAVValuesStored")
		}
		return nil
	}
	return m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_IdentifyReplyCommandGAVValuesStored) IsIdentifyReplyCommandGAVValuesStored() {}

func (m *_IdentifyReplyCommandGAVValuesStored) DeepCopy() any {
	return m.deepCopy()
}

func (m *_IdentifyReplyCommandGAVValuesStored) deepCopy() *_IdentifyReplyCommandGAVValuesStored {
	if m == nil {
		return nil
	}
	_IdentifyReplyCommandGAVValuesStoredCopy := &_IdentifyReplyCommandGAVValuesStored{
		m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).deepCopy(),
		utils.DeepCopySlice[byte, byte](m.Values),
	}
	m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand)._SubType = m
	return _IdentifyReplyCommandGAVValuesStoredCopy
}

func (m *_IdentifyReplyCommandGAVValuesStored) String() string {
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
