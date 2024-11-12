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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataCalendarAll is the corresponding interface of BACnetConstructedDataCalendarAll
type BACnetConstructedDataCalendarAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataCalendarAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataCalendarAll()
	// CreateBuilder creates a BACnetConstructedDataCalendarAllBuilder
	CreateBACnetConstructedDataCalendarAllBuilder() BACnetConstructedDataCalendarAllBuilder
}

// _BACnetConstructedDataCalendarAll is the data-structure of this message
type _BACnetConstructedDataCalendarAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataCalendarAll = (*_BACnetConstructedDataCalendarAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataCalendarAll)(nil)

// NewBACnetConstructedDataCalendarAll factory function for _BACnetConstructedDataCalendarAll
func NewBACnetConstructedDataCalendarAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCalendarAll {
	_result := &_BACnetConstructedDataCalendarAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataCalendarAllBuilder is a builder for BACnetConstructedDataCalendarAll
type BACnetConstructedDataCalendarAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataCalendarAllBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataCalendarAll or returns an error if something is wrong
	Build() (BACnetConstructedDataCalendarAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataCalendarAll
}

// NewBACnetConstructedDataCalendarAllBuilder() creates a BACnetConstructedDataCalendarAllBuilder
func NewBACnetConstructedDataCalendarAllBuilder() BACnetConstructedDataCalendarAllBuilder {
	return &_BACnetConstructedDataCalendarAllBuilder{_BACnetConstructedDataCalendarAll: new(_BACnetConstructedDataCalendarAll)}
}

type _BACnetConstructedDataCalendarAllBuilder struct {
	*_BACnetConstructedDataCalendarAll

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataCalendarAllBuilder) = (*_BACnetConstructedDataCalendarAllBuilder)(nil)

func (b *_BACnetConstructedDataCalendarAllBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataCalendarAllBuilder) WithMandatoryFields() BACnetConstructedDataCalendarAllBuilder {
	return b
}

func (b *_BACnetConstructedDataCalendarAllBuilder) Build() (BACnetConstructedDataCalendarAll, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataCalendarAll.deepCopy(), nil
}

func (b *_BACnetConstructedDataCalendarAllBuilder) MustBuild() BACnetConstructedDataCalendarAll {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataCalendarAllBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataCalendarAllBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataCalendarAllBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataCalendarAllBuilder().(*_BACnetConstructedDataCalendarAllBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataCalendarAllBuilder creates a BACnetConstructedDataCalendarAllBuilder
func (b *_BACnetConstructedDataCalendarAll) CreateBACnetConstructedDataCalendarAllBuilder() BACnetConstructedDataCalendarAllBuilder {
	if b == nil {
		return NewBACnetConstructedDataCalendarAllBuilder()
	}
	return &_BACnetConstructedDataCalendarAllBuilder{_BACnetConstructedDataCalendarAll: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCalendarAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_CALENDAR
}

func (m *_BACnetConstructedDataCalendarAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCalendarAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCalendarAll(structType any) BACnetConstructedDataCalendarAll {
	if casted, ok := structType.(BACnetConstructedDataCalendarAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCalendarAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCalendarAll) GetTypeName() string {
	return "BACnetConstructedDataCalendarAll"
}

func (m *_BACnetConstructedDataCalendarAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataCalendarAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataCalendarAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataCalendarAll BACnetConstructedDataCalendarAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCalendarAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCalendarAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCalendarAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCalendarAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataCalendarAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCalendarAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCalendarAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCalendarAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCalendarAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCalendarAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCalendarAll) IsBACnetConstructedDataCalendarAll() {}

func (m *_BACnetConstructedDataCalendarAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataCalendarAll) deepCopy() *_BACnetConstructedDataCalendarAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataCalendarAllCopy := &_BACnetConstructedDataCalendarAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataCalendarAllCopy
}

func (m *_BACnetConstructedDataCalendarAll) String() string {
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
