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

// BACnetConstructedDataObjectName is the corresponding interface of BACnetConstructedDataObjectName
type BACnetConstructedDataObjectName interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetObjectName returns ObjectName (property field)
	GetObjectName() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
	// IsBACnetConstructedDataObjectName is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataObjectName()
	// CreateBuilder creates a BACnetConstructedDataObjectNameBuilder
	CreateBACnetConstructedDataObjectNameBuilder() BACnetConstructedDataObjectNameBuilder
}

// _BACnetConstructedDataObjectName is the data-structure of this message
type _BACnetConstructedDataObjectName struct {
	BACnetConstructedDataContract
	ObjectName BACnetApplicationTagCharacterString
}

var _ BACnetConstructedDataObjectName = (*_BACnetConstructedDataObjectName)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataObjectName)(nil)

// NewBACnetConstructedDataObjectName factory function for _BACnetConstructedDataObjectName
func NewBACnetConstructedDataObjectName(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, objectName BACnetApplicationTagCharacterString, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataObjectName {
	if objectName == nil {
		panic("objectName of type BACnetApplicationTagCharacterString for BACnetConstructedDataObjectName must not be nil")
	}
	_result := &_BACnetConstructedDataObjectName{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ObjectName:                    objectName,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataObjectNameBuilder is a builder for BACnetConstructedDataObjectName
type BACnetConstructedDataObjectNameBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(objectName BACnetApplicationTagCharacterString) BACnetConstructedDataObjectNameBuilder
	// WithObjectName adds ObjectName (property field)
	WithObjectName(BACnetApplicationTagCharacterString) BACnetConstructedDataObjectNameBuilder
	// WithObjectNameBuilder adds ObjectName (property field) which is build by the builder
	WithObjectNameBuilder(func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetConstructedDataObjectNameBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataObjectName or returns an error if something is wrong
	Build() (BACnetConstructedDataObjectName, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataObjectName
}

// NewBACnetConstructedDataObjectNameBuilder() creates a BACnetConstructedDataObjectNameBuilder
func NewBACnetConstructedDataObjectNameBuilder() BACnetConstructedDataObjectNameBuilder {
	return &_BACnetConstructedDataObjectNameBuilder{_BACnetConstructedDataObjectName: new(_BACnetConstructedDataObjectName)}
}

type _BACnetConstructedDataObjectNameBuilder struct {
	*_BACnetConstructedDataObjectName

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataObjectNameBuilder) = (*_BACnetConstructedDataObjectNameBuilder)(nil)

func (b *_BACnetConstructedDataObjectNameBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataObjectNameBuilder) WithMandatoryFields(objectName BACnetApplicationTagCharacterString) BACnetConstructedDataObjectNameBuilder {
	return b.WithObjectName(objectName)
}

func (b *_BACnetConstructedDataObjectNameBuilder) WithObjectName(objectName BACnetApplicationTagCharacterString) BACnetConstructedDataObjectNameBuilder {
	b.ObjectName = objectName
	return b
}

func (b *_BACnetConstructedDataObjectNameBuilder) WithObjectNameBuilder(builderSupplier func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetConstructedDataObjectNameBuilder {
	builder := builderSupplier(b.ObjectName.CreateBACnetApplicationTagCharacterStringBuilder())
	var err error
	b.ObjectName, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagCharacterStringBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataObjectNameBuilder) Build() (BACnetConstructedDataObjectName, error) {
	if b.ObjectName == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'objectName' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataObjectName.deepCopy(), nil
}

func (b *_BACnetConstructedDataObjectNameBuilder) MustBuild() BACnetConstructedDataObjectName {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataObjectNameBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataObjectNameBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataObjectNameBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataObjectNameBuilder().(*_BACnetConstructedDataObjectNameBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataObjectNameBuilder creates a BACnetConstructedDataObjectNameBuilder
func (b *_BACnetConstructedDataObjectName) CreateBACnetConstructedDataObjectNameBuilder() BACnetConstructedDataObjectNameBuilder {
	if b == nil {
		return NewBACnetConstructedDataObjectNameBuilder()
	}
	return &_BACnetConstructedDataObjectNameBuilder{_BACnetConstructedDataObjectName: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataObjectName) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataObjectName) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OBJECT_NAME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataObjectName) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataObjectName) GetObjectName() BACnetApplicationTagCharacterString {
	return m.ObjectName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataObjectName) GetActualValue() BACnetApplicationTagCharacterString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagCharacterString(m.GetObjectName())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataObjectName(structType any) BACnetConstructedDataObjectName {
	if casted, ok := structType.(BACnetConstructedDataObjectName); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataObjectName); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataObjectName) GetTypeName() string {
	return "BACnetConstructedDataObjectName"
}

func (m *_BACnetConstructedDataObjectName) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (objectName)
	lengthInBits += m.ObjectName.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataObjectName) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataObjectName) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataObjectName BACnetConstructedDataObjectName, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataObjectName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataObjectName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectName, err := ReadSimpleField[BACnetApplicationTagCharacterString](ctx, "objectName", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectName' field"))
	}
	m.ObjectName = objectName

	actualValue, err := ReadVirtualField[BACnetApplicationTagCharacterString](ctx, "actualValue", (*BACnetApplicationTagCharacterString)(nil), objectName)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataObjectName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataObjectName")
	}

	return m, nil
}

func (m *_BACnetConstructedDataObjectName) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataObjectName) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataObjectName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataObjectName")
		}

		if err := WriteSimpleField[BACnetApplicationTagCharacterString](ctx, "objectName", m.GetObjectName(), WriteComplex[BACnetApplicationTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'objectName' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataObjectName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataObjectName")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataObjectName) IsBACnetConstructedDataObjectName() {}

func (m *_BACnetConstructedDataObjectName) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataObjectName) deepCopy() *_BACnetConstructedDataObjectName {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataObjectNameCopy := &_BACnetConstructedDataObjectName{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagCharacterString](m.ObjectName),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataObjectNameCopy
}

func (m *_BACnetConstructedDataObjectName) String() string {
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
