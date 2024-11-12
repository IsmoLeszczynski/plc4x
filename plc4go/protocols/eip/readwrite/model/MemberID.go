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

// MemberID is the corresponding interface of MemberID
type MemberID interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	LogicalSegmentType
	// GetFormat returns Format (property field)
	GetFormat() uint8
	// GetInstance returns Instance (property field)
	GetInstance() uint8
	// IsMemberID is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMemberID()
	// CreateBuilder creates a MemberIDBuilder
	CreateMemberIDBuilder() MemberIDBuilder
}

// _MemberID is the data-structure of this message
type _MemberID struct {
	LogicalSegmentTypeContract
	Format   uint8
	Instance uint8
}

var _ MemberID = (*_MemberID)(nil)
var _ LogicalSegmentTypeRequirements = (*_MemberID)(nil)

// NewMemberID factory function for _MemberID
func NewMemberID(format uint8, instance uint8) *_MemberID {
	_result := &_MemberID{
		LogicalSegmentTypeContract: NewLogicalSegmentType(),
		Format:                     format,
		Instance:                   instance,
	}
	_result.LogicalSegmentTypeContract.(*_LogicalSegmentType)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MemberIDBuilder is a builder for MemberID
type MemberIDBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(format uint8, instance uint8) MemberIDBuilder
	// WithFormat adds Format (property field)
	WithFormat(uint8) MemberIDBuilder
	// WithInstance adds Instance (property field)
	WithInstance(uint8) MemberIDBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() LogicalSegmentTypeBuilder
	// Build builds the MemberID or returns an error if something is wrong
	Build() (MemberID, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MemberID
}

// NewMemberIDBuilder() creates a MemberIDBuilder
func NewMemberIDBuilder() MemberIDBuilder {
	return &_MemberIDBuilder{_MemberID: new(_MemberID)}
}

type _MemberIDBuilder struct {
	*_MemberID

	parentBuilder *_LogicalSegmentTypeBuilder

	err *utils.MultiError
}

var _ (MemberIDBuilder) = (*_MemberIDBuilder)(nil)

func (b *_MemberIDBuilder) setParent(contract LogicalSegmentTypeContract) {
	b.LogicalSegmentTypeContract = contract
}

func (b *_MemberIDBuilder) WithMandatoryFields(format uint8, instance uint8) MemberIDBuilder {
	return b.WithFormat(format).WithInstance(instance)
}

func (b *_MemberIDBuilder) WithFormat(format uint8) MemberIDBuilder {
	b.Format = format
	return b
}

func (b *_MemberIDBuilder) WithInstance(instance uint8) MemberIDBuilder {
	b.Instance = instance
	return b
}

func (b *_MemberIDBuilder) Build() (MemberID, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._MemberID.deepCopy(), nil
}

func (b *_MemberIDBuilder) MustBuild() MemberID {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_MemberIDBuilder) Done() LogicalSegmentTypeBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewLogicalSegmentTypeBuilder().(*_LogicalSegmentTypeBuilder)
	}
	return b.parentBuilder
}

func (b *_MemberIDBuilder) buildForLogicalSegmentType() (LogicalSegmentType, error) {
	return b.Build()
}

func (b *_MemberIDBuilder) DeepCopy() any {
	_copy := b.CreateMemberIDBuilder().(*_MemberIDBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateMemberIDBuilder creates a MemberIDBuilder
func (b *_MemberID) CreateMemberIDBuilder() MemberIDBuilder {
	if b == nil {
		return NewMemberIDBuilder()
	}
	return &_MemberIDBuilder{_MemberID: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MemberID) GetLogicalSegmentType() uint8 {
	return 0x02
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MemberID) GetParent() LogicalSegmentTypeContract {
	return m.LogicalSegmentTypeContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MemberID) GetFormat() uint8 {
	return m.Format
}

func (m *_MemberID) GetInstance() uint8 {
	return m.Instance
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastMemberID(structType any) MemberID {
	if casted, ok := structType.(MemberID); ok {
		return casted
	}
	if casted, ok := structType.(*MemberID); ok {
		return *casted
	}
	return nil
}

func (m *_MemberID) GetTypeName() string {
	return "MemberID"
}

func (m *_MemberID) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.LogicalSegmentTypeContract.(*_LogicalSegmentType).getLengthInBits(ctx))

	// Simple field (format)
	lengthInBits += 2

	// Simple field (instance)
	lengthInBits += 8

	return lengthInBits
}

func (m *_MemberID) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MemberID) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_LogicalSegmentType) (__memberID MemberID, err error) {
	m.LogicalSegmentTypeContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MemberID"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MemberID")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	format, err := ReadSimpleField(ctx, "format", ReadUnsignedByte(readBuffer, uint8(2)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'format' field"))
	}
	m.Format = format

	instance, err := ReadSimpleField(ctx, "instance", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'instance' field"))
	}
	m.Instance = instance

	if closeErr := readBuffer.CloseContext("MemberID"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MemberID")
	}

	return m, nil
}

func (m *_MemberID) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MemberID) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MemberID"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MemberID")
		}

		if err := WriteSimpleField[uint8](ctx, "format", m.GetFormat(), WriteUnsignedByte(writeBuffer, 2)); err != nil {
			return errors.Wrap(err, "Error serializing 'format' field")
		}

		if err := WriteSimpleField[uint8](ctx, "instance", m.GetInstance(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'instance' field")
		}

		if popErr := writeBuffer.PopContext("MemberID"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MemberID")
		}
		return nil
	}
	return m.LogicalSegmentTypeContract.(*_LogicalSegmentType).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MemberID) IsMemberID() {}

func (m *_MemberID) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MemberID) deepCopy() *_MemberID {
	if m == nil {
		return nil
	}
	_MemberIDCopy := &_MemberID{
		m.LogicalSegmentTypeContract.(*_LogicalSegmentType).deepCopy(),
		m.Format,
		m.Instance,
	}
	m.LogicalSegmentTypeContract.(*_LogicalSegmentType)._SubType = m
	return _MemberIDCopy
}

func (m *_MemberID) String() string {
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
