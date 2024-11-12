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

// BACnetConstructedDataTrendLogLogDeviceObjectProperty is the corresponding interface of BACnetConstructedDataTrendLogLogDeviceObjectProperty
type BACnetConstructedDataTrendLogLogDeviceObjectProperty interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetLogDeviceObjectProperty returns LogDeviceObjectProperty (property field)
	GetLogDeviceObjectProperty() BACnetDeviceObjectPropertyReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDeviceObjectPropertyReference
	// IsBACnetConstructedDataTrendLogLogDeviceObjectProperty is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataTrendLogLogDeviceObjectProperty()
	// CreateBuilder creates a BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder
	CreateBACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder() BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder
}

// _BACnetConstructedDataTrendLogLogDeviceObjectProperty is the data-structure of this message
type _BACnetConstructedDataTrendLogLogDeviceObjectProperty struct {
	BACnetConstructedDataContract
	LogDeviceObjectProperty BACnetDeviceObjectPropertyReference
}

var _ BACnetConstructedDataTrendLogLogDeviceObjectProperty = (*_BACnetConstructedDataTrendLogLogDeviceObjectProperty)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataTrendLogLogDeviceObjectProperty)(nil)

// NewBACnetConstructedDataTrendLogLogDeviceObjectProperty factory function for _BACnetConstructedDataTrendLogLogDeviceObjectProperty
func NewBACnetConstructedDataTrendLogLogDeviceObjectProperty(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, logDeviceObjectProperty BACnetDeviceObjectPropertyReference, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTrendLogLogDeviceObjectProperty {
	if logDeviceObjectProperty == nil {
		panic("logDeviceObjectProperty of type BACnetDeviceObjectPropertyReference for BACnetConstructedDataTrendLogLogDeviceObjectProperty must not be nil")
	}
	_result := &_BACnetConstructedDataTrendLogLogDeviceObjectProperty{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LogDeviceObjectProperty:       logDeviceObjectProperty,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder is a builder for BACnetConstructedDataTrendLogLogDeviceObjectProperty
type BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(logDeviceObjectProperty BACnetDeviceObjectPropertyReference) BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder
	// WithLogDeviceObjectProperty adds LogDeviceObjectProperty (property field)
	WithLogDeviceObjectProperty(BACnetDeviceObjectPropertyReference) BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder
	// WithLogDeviceObjectPropertyBuilder adds LogDeviceObjectProperty (property field) which is build by the builder
	WithLogDeviceObjectPropertyBuilder(func(BACnetDeviceObjectPropertyReferenceBuilder) BACnetDeviceObjectPropertyReferenceBuilder) BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataTrendLogLogDeviceObjectProperty or returns an error if something is wrong
	Build() (BACnetConstructedDataTrendLogLogDeviceObjectProperty, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataTrendLogLogDeviceObjectProperty
}

// NewBACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder() creates a BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder
func NewBACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder() BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder {
	return &_BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder{_BACnetConstructedDataTrendLogLogDeviceObjectProperty: new(_BACnetConstructedDataTrendLogLogDeviceObjectProperty)}
}

type _BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder struct {
	*_BACnetConstructedDataTrendLogLogDeviceObjectProperty

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder) = (*_BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder)(nil)

func (b *_BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder) WithMandatoryFields(logDeviceObjectProperty BACnetDeviceObjectPropertyReference) BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder {
	return b.WithLogDeviceObjectProperty(logDeviceObjectProperty)
}

func (b *_BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder) WithLogDeviceObjectProperty(logDeviceObjectProperty BACnetDeviceObjectPropertyReference) BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder {
	b.LogDeviceObjectProperty = logDeviceObjectProperty
	return b
}

func (b *_BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder) WithLogDeviceObjectPropertyBuilder(builderSupplier func(BACnetDeviceObjectPropertyReferenceBuilder) BACnetDeviceObjectPropertyReferenceBuilder) BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder {
	builder := builderSupplier(b.LogDeviceObjectProperty.CreateBACnetDeviceObjectPropertyReferenceBuilder())
	var err error
	b.LogDeviceObjectProperty, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDeviceObjectPropertyReferenceBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder) Build() (BACnetConstructedDataTrendLogLogDeviceObjectProperty, error) {
	if b.LogDeviceObjectProperty == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'logDeviceObjectProperty' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataTrendLogLogDeviceObjectProperty.deepCopy(), nil
}

func (b *_BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder) MustBuild() BACnetConstructedDataTrendLogLogDeviceObjectProperty {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder().(*_BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder creates a BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder
func (b *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) CreateBACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder() BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder {
	if b == nil {
		return NewBACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder()
	}
	return &_BACnetConstructedDataTrendLogLogDeviceObjectPropertyBuilder{_BACnetConstructedDataTrendLogLogDeviceObjectProperty: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_TREND_LOG
}

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOG_DEVICE_OBJECT_PROPERTY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetLogDeviceObjectProperty() BACnetDeviceObjectPropertyReference {
	return m.LogDeviceObjectProperty
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetActualValue() BACnetDeviceObjectPropertyReference {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDeviceObjectPropertyReference(m.GetLogDeviceObjectProperty())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTrendLogLogDeviceObjectProperty(structType any) BACnetConstructedDataTrendLogLogDeviceObjectProperty {
	if casted, ok := structType.(BACnetConstructedDataTrendLogLogDeviceObjectProperty); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTrendLogLogDeviceObjectProperty); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetTypeName() string {
	return "BACnetConstructedDataTrendLogLogDeviceObjectProperty"
}

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (logDeviceObjectProperty)
	lengthInBits += m.LogDeviceObjectProperty.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataTrendLogLogDeviceObjectProperty BACnetConstructedDataTrendLogLogDeviceObjectProperty, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTrendLogLogDeviceObjectProperty"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTrendLogLogDeviceObjectProperty")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	logDeviceObjectProperty, err := ReadSimpleField[BACnetDeviceObjectPropertyReference](ctx, "logDeviceObjectProperty", ReadComplex[BACnetDeviceObjectPropertyReference](BACnetDeviceObjectPropertyReferenceParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'logDeviceObjectProperty' field"))
	}
	m.LogDeviceObjectProperty = logDeviceObjectProperty

	actualValue, err := ReadVirtualField[BACnetDeviceObjectPropertyReference](ctx, "actualValue", (*BACnetDeviceObjectPropertyReference)(nil), logDeviceObjectProperty)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTrendLogLogDeviceObjectProperty"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTrendLogLogDeviceObjectProperty")
	}

	return m, nil
}

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTrendLogLogDeviceObjectProperty"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTrendLogLogDeviceObjectProperty")
		}

		if err := WriteSimpleField[BACnetDeviceObjectPropertyReference](ctx, "logDeviceObjectProperty", m.GetLogDeviceObjectProperty(), WriteComplex[BACnetDeviceObjectPropertyReference](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'logDeviceObjectProperty' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTrendLogLogDeviceObjectProperty"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTrendLogLogDeviceObjectProperty")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) IsBACnetConstructedDataTrendLogLogDeviceObjectProperty() {
}

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) deepCopy() *_BACnetConstructedDataTrendLogLogDeviceObjectProperty {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataTrendLogLogDeviceObjectPropertyCopy := &_BACnetConstructedDataTrendLogLogDeviceObjectProperty{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetDeviceObjectPropertyReference](m.LogDeviceObjectProperty),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataTrendLogLogDeviceObjectPropertyCopy
}

func (m *_BACnetConstructedDataTrendLogLogDeviceObjectProperty) String() string {
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
