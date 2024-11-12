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

// BACnetConstructedDataTimeOfActiveTimeReset is the corresponding interface of BACnetConstructedDataTimeOfActiveTimeReset
type BACnetConstructedDataTimeOfActiveTimeReset interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetTimeOfActiveTimeReset returns TimeOfActiveTimeReset (property field)
	GetTimeOfActiveTimeReset() BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDateTime
	// IsBACnetConstructedDataTimeOfActiveTimeReset is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataTimeOfActiveTimeReset()
	// CreateBuilder creates a BACnetConstructedDataTimeOfActiveTimeResetBuilder
	CreateBACnetConstructedDataTimeOfActiveTimeResetBuilder() BACnetConstructedDataTimeOfActiveTimeResetBuilder
}

// _BACnetConstructedDataTimeOfActiveTimeReset is the data-structure of this message
type _BACnetConstructedDataTimeOfActiveTimeReset struct {
	BACnetConstructedDataContract
	TimeOfActiveTimeReset BACnetDateTime
}

var _ BACnetConstructedDataTimeOfActiveTimeReset = (*_BACnetConstructedDataTimeOfActiveTimeReset)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataTimeOfActiveTimeReset)(nil)

// NewBACnetConstructedDataTimeOfActiveTimeReset factory function for _BACnetConstructedDataTimeOfActiveTimeReset
func NewBACnetConstructedDataTimeOfActiveTimeReset(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, timeOfActiveTimeReset BACnetDateTime, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTimeOfActiveTimeReset {
	if timeOfActiveTimeReset == nil {
		panic("timeOfActiveTimeReset of type BACnetDateTime for BACnetConstructedDataTimeOfActiveTimeReset must not be nil")
	}
	_result := &_BACnetConstructedDataTimeOfActiveTimeReset{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		TimeOfActiveTimeReset:         timeOfActiveTimeReset,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataTimeOfActiveTimeResetBuilder is a builder for BACnetConstructedDataTimeOfActiveTimeReset
type BACnetConstructedDataTimeOfActiveTimeResetBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(timeOfActiveTimeReset BACnetDateTime) BACnetConstructedDataTimeOfActiveTimeResetBuilder
	// WithTimeOfActiveTimeReset adds TimeOfActiveTimeReset (property field)
	WithTimeOfActiveTimeReset(BACnetDateTime) BACnetConstructedDataTimeOfActiveTimeResetBuilder
	// WithTimeOfActiveTimeResetBuilder adds TimeOfActiveTimeReset (property field) which is build by the builder
	WithTimeOfActiveTimeResetBuilder(func(BACnetDateTimeBuilder) BACnetDateTimeBuilder) BACnetConstructedDataTimeOfActiveTimeResetBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataTimeOfActiveTimeReset or returns an error if something is wrong
	Build() (BACnetConstructedDataTimeOfActiveTimeReset, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataTimeOfActiveTimeReset
}

// NewBACnetConstructedDataTimeOfActiveTimeResetBuilder() creates a BACnetConstructedDataTimeOfActiveTimeResetBuilder
func NewBACnetConstructedDataTimeOfActiveTimeResetBuilder() BACnetConstructedDataTimeOfActiveTimeResetBuilder {
	return &_BACnetConstructedDataTimeOfActiveTimeResetBuilder{_BACnetConstructedDataTimeOfActiveTimeReset: new(_BACnetConstructedDataTimeOfActiveTimeReset)}
}

type _BACnetConstructedDataTimeOfActiveTimeResetBuilder struct {
	*_BACnetConstructedDataTimeOfActiveTimeReset

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataTimeOfActiveTimeResetBuilder) = (*_BACnetConstructedDataTimeOfActiveTimeResetBuilder)(nil)

func (b *_BACnetConstructedDataTimeOfActiveTimeResetBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataTimeOfActiveTimeResetBuilder) WithMandatoryFields(timeOfActiveTimeReset BACnetDateTime) BACnetConstructedDataTimeOfActiveTimeResetBuilder {
	return b.WithTimeOfActiveTimeReset(timeOfActiveTimeReset)
}

func (b *_BACnetConstructedDataTimeOfActiveTimeResetBuilder) WithTimeOfActiveTimeReset(timeOfActiveTimeReset BACnetDateTime) BACnetConstructedDataTimeOfActiveTimeResetBuilder {
	b.TimeOfActiveTimeReset = timeOfActiveTimeReset
	return b
}

func (b *_BACnetConstructedDataTimeOfActiveTimeResetBuilder) WithTimeOfActiveTimeResetBuilder(builderSupplier func(BACnetDateTimeBuilder) BACnetDateTimeBuilder) BACnetConstructedDataTimeOfActiveTimeResetBuilder {
	builder := builderSupplier(b.TimeOfActiveTimeReset.CreateBACnetDateTimeBuilder())
	var err error
	b.TimeOfActiveTimeReset, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDateTimeBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataTimeOfActiveTimeResetBuilder) Build() (BACnetConstructedDataTimeOfActiveTimeReset, error) {
	if b.TimeOfActiveTimeReset == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'timeOfActiveTimeReset' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataTimeOfActiveTimeReset.deepCopy(), nil
}

func (b *_BACnetConstructedDataTimeOfActiveTimeResetBuilder) MustBuild() BACnetConstructedDataTimeOfActiveTimeReset {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataTimeOfActiveTimeResetBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataTimeOfActiveTimeResetBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataTimeOfActiveTimeResetBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataTimeOfActiveTimeResetBuilder().(*_BACnetConstructedDataTimeOfActiveTimeResetBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataTimeOfActiveTimeResetBuilder creates a BACnetConstructedDataTimeOfActiveTimeResetBuilder
func (b *_BACnetConstructedDataTimeOfActiveTimeReset) CreateBACnetConstructedDataTimeOfActiveTimeResetBuilder() BACnetConstructedDataTimeOfActiveTimeResetBuilder {
	if b == nil {
		return NewBACnetConstructedDataTimeOfActiveTimeResetBuilder()
	}
	return &_BACnetConstructedDataTimeOfActiveTimeResetBuilder{_BACnetConstructedDataTimeOfActiveTimeReset: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TIME_OF_ACTIVE_TIME_RESET
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) GetTimeOfActiveTimeReset() BACnetDateTime {
	return m.TimeOfActiveTimeReset
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) GetActualValue() BACnetDateTime {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDateTime(m.GetTimeOfActiveTimeReset())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTimeOfActiveTimeReset(structType any) BACnetConstructedDataTimeOfActiveTimeReset {
	if casted, ok := structType.(BACnetConstructedDataTimeOfActiveTimeReset); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimeOfActiveTimeReset); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) GetTypeName() string {
	return "BACnetConstructedDataTimeOfActiveTimeReset"
}

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (timeOfActiveTimeReset)
	lengthInBits += m.TimeOfActiveTimeReset.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataTimeOfActiveTimeReset BACnetConstructedDataTimeOfActiveTimeReset, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimeOfActiveTimeReset"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTimeOfActiveTimeReset")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timeOfActiveTimeReset, err := ReadSimpleField[BACnetDateTime](ctx, "timeOfActiveTimeReset", ReadComplex[BACnetDateTime](BACnetDateTimeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeOfActiveTimeReset' field"))
	}
	m.TimeOfActiveTimeReset = timeOfActiveTimeReset

	actualValue, err := ReadVirtualField[BACnetDateTime](ctx, "actualValue", (*BACnetDateTime)(nil), timeOfActiveTimeReset)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimeOfActiveTimeReset"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTimeOfActiveTimeReset")
	}

	return m, nil
}

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimeOfActiveTimeReset"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTimeOfActiveTimeReset")
		}

		if err := WriteSimpleField[BACnetDateTime](ctx, "timeOfActiveTimeReset", m.GetTimeOfActiveTimeReset(), WriteComplex[BACnetDateTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeOfActiveTimeReset' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimeOfActiveTimeReset"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTimeOfActiveTimeReset")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) IsBACnetConstructedDataTimeOfActiveTimeReset() {
}

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) deepCopy() *_BACnetConstructedDataTimeOfActiveTimeReset {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataTimeOfActiveTimeResetCopy := &_BACnetConstructedDataTimeOfActiveTimeReset{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetDateTime](m.TimeOfActiveTimeReset),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataTimeOfActiveTimeResetCopy
}

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) String() string {
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
