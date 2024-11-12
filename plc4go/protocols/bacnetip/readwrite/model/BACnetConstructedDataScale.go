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

// BACnetConstructedDataScale is the corresponding interface of BACnetConstructedDataScale
type BACnetConstructedDataScale interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetScale returns Scale (property field)
	GetScale() BACnetScale
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetScale
	// IsBACnetConstructedDataScale is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataScale()
	// CreateBuilder creates a BACnetConstructedDataScaleBuilder
	CreateBACnetConstructedDataScaleBuilder() BACnetConstructedDataScaleBuilder
}

// _BACnetConstructedDataScale is the data-structure of this message
type _BACnetConstructedDataScale struct {
	BACnetConstructedDataContract
	Scale BACnetScale
}

var _ BACnetConstructedDataScale = (*_BACnetConstructedDataScale)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataScale)(nil)

// NewBACnetConstructedDataScale factory function for _BACnetConstructedDataScale
func NewBACnetConstructedDataScale(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, scale BACnetScale, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataScale {
	if scale == nil {
		panic("scale of type BACnetScale for BACnetConstructedDataScale must not be nil")
	}
	_result := &_BACnetConstructedDataScale{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Scale:                         scale,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataScaleBuilder is a builder for BACnetConstructedDataScale
type BACnetConstructedDataScaleBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(scale BACnetScale) BACnetConstructedDataScaleBuilder
	// WithScale adds Scale (property field)
	WithScale(BACnetScale) BACnetConstructedDataScaleBuilder
	// WithScaleBuilder adds Scale (property field) which is build by the builder
	WithScaleBuilder(func(BACnetScaleBuilder) BACnetScaleBuilder) BACnetConstructedDataScaleBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataScale or returns an error if something is wrong
	Build() (BACnetConstructedDataScale, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataScale
}

// NewBACnetConstructedDataScaleBuilder() creates a BACnetConstructedDataScaleBuilder
func NewBACnetConstructedDataScaleBuilder() BACnetConstructedDataScaleBuilder {
	return &_BACnetConstructedDataScaleBuilder{_BACnetConstructedDataScale: new(_BACnetConstructedDataScale)}
}

type _BACnetConstructedDataScaleBuilder struct {
	*_BACnetConstructedDataScale

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataScaleBuilder) = (*_BACnetConstructedDataScaleBuilder)(nil)

func (b *_BACnetConstructedDataScaleBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataScaleBuilder) WithMandatoryFields(scale BACnetScale) BACnetConstructedDataScaleBuilder {
	return b.WithScale(scale)
}

func (b *_BACnetConstructedDataScaleBuilder) WithScale(scale BACnetScale) BACnetConstructedDataScaleBuilder {
	b.Scale = scale
	return b
}

func (b *_BACnetConstructedDataScaleBuilder) WithScaleBuilder(builderSupplier func(BACnetScaleBuilder) BACnetScaleBuilder) BACnetConstructedDataScaleBuilder {
	builder := builderSupplier(b.Scale.CreateBACnetScaleBuilder())
	var err error
	b.Scale, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetScaleBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataScaleBuilder) Build() (BACnetConstructedDataScale, error) {
	if b.Scale == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'scale' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataScale.deepCopy(), nil
}

func (b *_BACnetConstructedDataScaleBuilder) MustBuild() BACnetConstructedDataScale {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataScaleBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataScaleBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataScaleBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataScaleBuilder().(*_BACnetConstructedDataScaleBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataScaleBuilder creates a BACnetConstructedDataScaleBuilder
func (b *_BACnetConstructedDataScale) CreateBACnetConstructedDataScaleBuilder() BACnetConstructedDataScaleBuilder {
	if b == nil {
		return NewBACnetConstructedDataScaleBuilder()
	}
	return &_BACnetConstructedDataScaleBuilder{_BACnetConstructedDataScale: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataScale) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataScale) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SCALE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataScale) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataScale) GetScale() BACnetScale {
	return m.Scale
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataScale) GetActualValue() BACnetScale {
	ctx := context.Background()
	_ = ctx
	return CastBACnetScale(m.GetScale())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataScale(structType any) BACnetConstructedDataScale {
	if casted, ok := structType.(BACnetConstructedDataScale); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataScale); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataScale) GetTypeName() string {
	return "BACnetConstructedDataScale"
}

func (m *_BACnetConstructedDataScale) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (scale)
	lengthInBits += m.Scale.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataScale) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataScale) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataScale BACnetConstructedDataScale, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataScale"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataScale")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	scale, err := ReadSimpleField[BACnetScale](ctx, "scale", ReadComplex[BACnetScale](BACnetScaleParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'scale' field"))
	}
	m.Scale = scale

	actualValue, err := ReadVirtualField[BACnetScale](ctx, "actualValue", (*BACnetScale)(nil), scale)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataScale"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataScale")
	}

	return m, nil
}

func (m *_BACnetConstructedDataScale) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataScale) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataScale"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataScale")
		}

		if err := WriteSimpleField[BACnetScale](ctx, "scale", m.GetScale(), WriteComplex[BACnetScale](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'scale' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataScale"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataScale")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataScale) IsBACnetConstructedDataScale() {}

func (m *_BACnetConstructedDataScale) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataScale) deepCopy() *_BACnetConstructedDataScale {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataScaleCopy := &_BACnetConstructedDataScale{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetScale](m.Scale),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataScaleCopy
}

func (m *_BACnetConstructedDataScale) String() string {
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
