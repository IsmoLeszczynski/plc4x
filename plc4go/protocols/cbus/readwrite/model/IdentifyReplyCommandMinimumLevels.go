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

// IdentifyReplyCommandMinimumLevels is the corresponding interface of IdentifyReplyCommandMinimumLevels
type IdentifyReplyCommandMinimumLevels interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	IdentifyReplyCommand
	// GetMinimumLevels returns MinimumLevels (property field)
	GetMinimumLevels() []byte
	// IsIdentifyReplyCommandMinimumLevels is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsIdentifyReplyCommandMinimumLevels()
	// CreateBuilder creates a IdentifyReplyCommandMinimumLevelsBuilder
	CreateIdentifyReplyCommandMinimumLevelsBuilder() IdentifyReplyCommandMinimumLevelsBuilder
}

// _IdentifyReplyCommandMinimumLevels is the data-structure of this message
type _IdentifyReplyCommandMinimumLevels struct {
	IdentifyReplyCommandContract
	MinimumLevels []byte
}

var _ IdentifyReplyCommandMinimumLevels = (*_IdentifyReplyCommandMinimumLevels)(nil)
var _ IdentifyReplyCommandRequirements = (*_IdentifyReplyCommandMinimumLevels)(nil)

// NewIdentifyReplyCommandMinimumLevels factory function for _IdentifyReplyCommandMinimumLevels
func NewIdentifyReplyCommandMinimumLevels(minimumLevels []byte, numBytes uint8) *_IdentifyReplyCommandMinimumLevels {
	_result := &_IdentifyReplyCommandMinimumLevels{
		IdentifyReplyCommandContract: NewIdentifyReplyCommand(numBytes),
		MinimumLevels:                minimumLevels,
	}
	_result.IdentifyReplyCommandContract.(*_IdentifyReplyCommand)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// IdentifyReplyCommandMinimumLevelsBuilder is a builder for IdentifyReplyCommandMinimumLevels
type IdentifyReplyCommandMinimumLevelsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(minimumLevels []byte) IdentifyReplyCommandMinimumLevelsBuilder
	// WithMinimumLevels adds MinimumLevels (property field)
	WithMinimumLevels(...byte) IdentifyReplyCommandMinimumLevelsBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() IdentifyReplyCommandBuilder
	// Build builds the IdentifyReplyCommandMinimumLevels or returns an error if something is wrong
	Build() (IdentifyReplyCommandMinimumLevels, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() IdentifyReplyCommandMinimumLevels
}

// NewIdentifyReplyCommandMinimumLevelsBuilder() creates a IdentifyReplyCommandMinimumLevelsBuilder
func NewIdentifyReplyCommandMinimumLevelsBuilder() IdentifyReplyCommandMinimumLevelsBuilder {
	return &_IdentifyReplyCommandMinimumLevelsBuilder{_IdentifyReplyCommandMinimumLevels: new(_IdentifyReplyCommandMinimumLevels)}
}

type _IdentifyReplyCommandMinimumLevelsBuilder struct {
	*_IdentifyReplyCommandMinimumLevels

	parentBuilder *_IdentifyReplyCommandBuilder

	err *utils.MultiError
}

var _ (IdentifyReplyCommandMinimumLevelsBuilder) = (*_IdentifyReplyCommandMinimumLevelsBuilder)(nil)

func (b *_IdentifyReplyCommandMinimumLevelsBuilder) setParent(contract IdentifyReplyCommandContract) {
	b.IdentifyReplyCommandContract = contract
}

func (b *_IdentifyReplyCommandMinimumLevelsBuilder) WithMandatoryFields(minimumLevels []byte) IdentifyReplyCommandMinimumLevelsBuilder {
	return b.WithMinimumLevels(minimumLevels...)
}

func (b *_IdentifyReplyCommandMinimumLevelsBuilder) WithMinimumLevels(minimumLevels ...byte) IdentifyReplyCommandMinimumLevelsBuilder {
	b.MinimumLevels = minimumLevels
	return b
}

func (b *_IdentifyReplyCommandMinimumLevelsBuilder) Build() (IdentifyReplyCommandMinimumLevels, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._IdentifyReplyCommandMinimumLevels.deepCopy(), nil
}

func (b *_IdentifyReplyCommandMinimumLevelsBuilder) MustBuild() IdentifyReplyCommandMinimumLevels {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_IdentifyReplyCommandMinimumLevelsBuilder) Done() IdentifyReplyCommandBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewIdentifyReplyCommandBuilder().(*_IdentifyReplyCommandBuilder)
	}
	return b.parentBuilder
}

func (b *_IdentifyReplyCommandMinimumLevelsBuilder) buildForIdentifyReplyCommand() (IdentifyReplyCommand, error) {
	return b.Build()
}

func (b *_IdentifyReplyCommandMinimumLevelsBuilder) DeepCopy() any {
	_copy := b.CreateIdentifyReplyCommandMinimumLevelsBuilder().(*_IdentifyReplyCommandMinimumLevelsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateIdentifyReplyCommandMinimumLevelsBuilder creates a IdentifyReplyCommandMinimumLevelsBuilder
func (b *_IdentifyReplyCommandMinimumLevels) CreateIdentifyReplyCommandMinimumLevelsBuilder() IdentifyReplyCommandMinimumLevelsBuilder {
	if b == nil {
		return NewIdentifyReplyCommandMinimumLevelsBuilder()
	}
	return &_IdentifyReplyCommandMinimumLevelsBuilder{_IdentifyReplyCommandMinimumLevels: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandMinimumLevels) GetAttribute() Attribute {
	return Attribute_MinimumLevels
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandMinimumLevels) GetParent() IdentifyReplyCommandContract {
	return m.IdentifyReplyCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandMinimumLevels) GetMinimumLevels() []byte {
	return m.MinimumLevels
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandMinimumLevels(structType any) IdentifyReplyCommandMinimumLevels {
	if casted, ok := structType.(IdentifyReplyCommandMinimumLevels); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandMinimumLevels); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandMinimumLevels) GetTypeName() string {
	return "IdentifyReplyCommandMinimumLevels"
}

func (m *_IdentifyReplyCommandMinimumLevels) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).getLengthInBits(ctx))

	// Array field
	if len(m.MinimumLevels) > 0 {
		lengthInBits += 8 * uint16(len(m.MinimumLevels))
	}

	return lengthInBits
}

func (m *_IdentifyReplyCommandMinimumLevels) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_IdentifyReplyCommandMinimumLevels) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_IdentifyReplyCommand, attribute Attribute, numBytes uint8) (__identifyReplyCommandMinimumLevels IdentifyReplyCommandMinimumLevels, err error) {
	m.IdentifyReplyCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandMinimumLevels"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandMinimumLevels")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	minimumLevels, err := readBuffer.ReadByteArray("minimumLevels", int(numBytes))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'minimumLevels' field"))
	}
	m.MinimumLevels = minimumLevels

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandMinimumLevels"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandMinimumLevels")
	}

	return m, nil
}

func (m *_IdentifyReplyCommandMinimumLevels) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IdentifyReplyCommandMinimumLevels) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandMinimumLevels"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandMinimumLevels")
		}

		if err := WriteByteArrayField(ctx, "minimumLevels", m.GetMinimumLevels(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'minimumLevels' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandMinimumLevels"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandMinimumLevels")
		}
		return nil
	}
	return m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_IdentifyReplyCommandMinimumLevels) IsIdentifyReplyCommandMinimumLevels() {}

func (m *_IdentifyReplyCommandMinimumLevels) DeepCopy() any {
	return m.deepCopy()
}

func (m *_IdentifyReplyCommandMinimumLevels) deepCopy() *_IdentifyReplyCommandMinimumLevels {
	if m == nil {
		return nil
	}
	_IdentifyReplyCommandMinimumLevelsCopy := &_IdentifyReplyCommandMinimumLevels{
		m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).deepCopy(),
		utils.DeepCopySlice[byte, byte](m.MinimumLevels),
	}
	m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand)._SubType = m
	return _IdentifyReplyCommandMinimumLevelsCopy
}

func (m *_IdentifyReplyCommandMinimumLevels) String() string {
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
