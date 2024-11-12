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

// BACnetConstructedDataCarMovingDirection is the corresponding interface of BACnetConstructedDataCarMovingDirection
type BACnetConstructedDataCarMovingDirection interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetCarMovingDirection returns CarMovingDirection (property field)
	GetCarMovingDirection() BACnetLiftCarDirectionTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetLiftCarDirectionTagged
	// IsBACnetConstructedDataCarMovingDirection is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataCarMovingDirection()
	// CreateBuilder creates a BACnetConstructedDataCarMovingDirectionBuilder
	CreateBACnetConstructedDataCarMovingDirectionBuilder() BACnetConstructedDataCarMovingDirectionBuilder
}

// _BACnetConstructedDataCarMovingDirection is the data-structure of this message
type _BACnetConstructedDataCarMovingDirection struct {
	BACnetConstructedDataContract
	CarMovingDirection BACnetLiftCarDirectionTagged
}

var _ BACnetConstructedDataCarMovingDirection = (*_BACnetConstructedDataCarMovingDirection)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataCarMovingDirection)(nil)

// NewBACnetConstructedDataCarMovingDirection factory function for _BACnetConstructedDataCarMovingDirection
func NewBACnetConstructedDataCarMovingDirection(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, carMovingDirection BACnetLiftCarDirectionTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCarMovingDirection {
	if carMovingDirection == nil {
		panic("carMovingDirection of type BACnetLiftCarDirectionTagged for BACnetConstructedDataCarMovingDirection must not be nil")
	}
	_result := &_BACnetConstructedDataCarMovingDirection{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		CarMovingDirection:            carMovingDirection,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataCarMovingDirectionBuilder is a builder for BACnetConstructedDataCarMovingDirection
type BACnetConstructedDataCarMovingDirectionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(carMovingDirection BACnetLiftCarDirectionTagged) BACnetConstructedDataCarMovingDirectionBuilder
	// WithCarMovingDirection adds CarMovingDirection (property field)
	WithCarMovingDirection(BACnetLiftCarDirectionTagged) BACnetConstructedDataCarMovingDirectionBuilder
	// WithCarMovingDirectionBuilder adds CarMovingDirection (property field) which is build by the builder
	WithCarMovingDirectionBuilder(func(BACnetLiftCarDirectionTaggedBuilder) BACnetLiftCarDirectionTaggedBuilder) BACnetConstructedDataCarMovingDirectionBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataCarMovingDirection or returns an error if something is wrong
	Build() (BACnetConstructedDataCarMovingDirection, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataCarMovingDirection
}

// NewBACnetConstructedDataCarMovingDirectionBuilder() creates a BACnetConstructedDataCarMovingDirectionBuilder
func NewBACnetConstructedDataCarMovingDirectionBuilder() BACnetConstructedDataCarMovingDirectionBuilder {
	return &_BACnetConstructedDataCarMovingDirectionBuilder{_BACnetConstructedDataCarMovingDirection: new(_BACnetConstructedDataCarMovingDirection)}
}

type _BACnetConstructedDataCarMovingDirectionBuilder struct {
	*_BACnetConstructedDataCarMovingDirection

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataCarMovingDirectionBuilder) = (*_BACnetConstructedDataCarMovingDirectionBuilder)(nil)

func (b *_BACnetConstructedDataCarMovingDirectionBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataCarMovingDirectionBuilder) WithMandatoryFields(carMovingDirection BACnetLiftCarDirectionTagged) BACnetConstructedDataCarMovingDirectionBuilder {
	return b.WithCarMovingDirection(carMovingDirection)
}

func (b *_BACnetConstructedDataCarMovingDirectionBuilder) WithCarMovingDirection(carMovingDirection BACnetLiftCarDirectionTagged) BACnetConstructedDataCarMovingDirectionBuilder {
	b.CarMovingDirection = carMovingDirection
	return b
}

func (b *_BACnetConstructedDataCarMovingDirectionBuilder) WithCarMovingDirectionBuilder(builderSupplier func(BACnetLiftCarDirectionTaggedBuilder) BACnetLiftCarDirectionTaggedBuilder) BACnetConstructedDataCarMovingDirectionBuilder {
	builder := builderSupplier(b.CarMovingDirection.CreateBACnetLiftCarDirectionTaggedBuilder())
	var err error
	b.CarMovingDirection, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetLiftCarDirectionTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataCarMovingDirectionBuilder) Build() (BACnetConstructedDataCarMovingDirection, error) {
	if b.CarMovingDirection == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'carMovingDirection' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataCarMovingDirection.deepCopy(), nil
}

func (b *_BACnetConstructedDataCarMovingDirectionBuilder) MustBuild() BACnetConstructedDataCarMovingDirection {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataCarMovingDirectionBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataCarMovingDirectionBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataCarMovingDirectionBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataCarMovingDirectionBuilder().(*_BACnetConstructedDataCarMovingDirectionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataCarMovingDirectionBuilder creates a BACnetConstructedDataCarMovingDirectionBuilder
func (b *_BACnetConstructedDataCarMovingDirection) CreateBACnetConstructedDataCarMovingDirectionBuilder() BACnetConstructedDataCarMovingDirectionBuilder {
	if b == nil {
		return NewBACnetConstructedDataCarMovingDirectionBuilder()
	}
	return &_BACnetConstructedDataCarMovingDirectionBuilder{_BACnetConstructedDataCarMovingDirection: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCarMovingDirection) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCarMovingDirection) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CAR_MOVING_DIRECTION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCarMovingDirection) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCarMovingDirection) GetCarMovingDirection() BACnetLiftCarDirectionTagged {
	return m.CarMovingDirection
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCarMovingDirection) GetActualValue() BACnetLiftCarDirectionTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetLiftCarDirectionTagged(m.GetCarMovingDirection())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCarMovingDirection(structType any) BACnetConstructedDataCarMovingDirection {
	if casted, ok := structType.(BACnetConstructedDataCarMovingDirection); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCarMovingDirection); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCarMovingDirection) GetTypeName() string {
	return "BACnetConstructedDataCarMovingDirection"
}

func (m *_BACnetConstructedDataCarMovingDirection) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (carMovingDirection)
	lengthInBits += m.CarMovingDirection.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataCarMovingDirection) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataCarMovingDirection) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataCarMovingDirection BACnetConstructedDataCarMovingDirection, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCarMovingDirection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCarMovingDirection")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	carMovingDirection, err := ReadSimpleField[BACnetLiftCarDirectionTagged](ctx, "carMovingDirection", ReadComplex[BACnetLiftCarDirectionTagged](BACnetLiftCarDirectionTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'carMovingDirection' field"))
	}
	m.CarMovingDirection = carMovingDirection

	actualValue, err := ReadVirtualField[BACnetLiftCarDirectionTagged](ctx, "actualValue", (*BACnetLiftCarDirectionTagged)(nil), carMovingDirection)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCarMovingDirection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCarMovingDirection")
	}

	return m, nil
}

func (m *_BACnetConstructedDataCarMovingDirection) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCarMovingDirection) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCarMovingDirection"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCarMovingDirection")
		}

		if err := WriteSimpleField[BACnetLiftCarDirectionTagged](ctx, "carMovingDirection", m.GetCarMovingDirection(), WriteComplex[BACnetLiftCarDirectionTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'carMovingDirection' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCarMovingDirection"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCarMovingDirection")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCarMovingDirection) IsBACnetConstructedDataCarMovingDirection() {}

func (m *_BACnetConstructedDataCarMovingDirection) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataCarMovingDirection) deepCopy() *_BACnetConstructedDataCarMovingDirection {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataCarMovingDirectionCopy := &_BACnetConstructedDataCarMovingDirection{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetLiftCarDirectionTagged](m.CarMovingDirection),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataCarMovingDirectionCopy
}

func (m *_BACnetConstructedDataCarMovingDirection) String() string {
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
