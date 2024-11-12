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

// BACnetConstructedDataModelName is the corresponding interface of BACnetConstructedDataModelName
type BACnetConstructedDataModelName interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetModelName returns ModelName (property field)
	GetModelName() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
	// IsBACnetConstructedDataModelName is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataModelName()
	// CreateBuilder creates a BACnetConstructedDataModelNameBuilder
	CreateBACnetConstructedDataModelNameBuilder() BACnetConstructedDataModelNameBuilder
}

// _BACnetConstructedDataModelName is the data-structure of this message
type _BACnetConstructedDataModelName struct {
	BACnetConstructedDataContract
	ModelName BACnetApplicationTagCharacterString
}

var _ BACnetConstructedDataModelName = (*_BACnetConstructedDataModelName)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataModelName)(nil)

// NewBACnetConstructedDataModelName factory function for _BACnetConstructedDataModelName
func NewBACnetConstructedDataModelName(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, modelName BACnetApplicationTagCharacterString, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataModelName {
	if modelName == nil {
		panic("modelName of type BACnetApplicationTagCharacterString for BACnetConstructedDataModelName must not be nil")
	}
	_result := &_BACnetConstructedDataModelName{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ModelName:                     modelName,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataModelNameBuilder is a builder for BACnetConstructedDataModelName
type BACnetConstructedDataModelNameBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(modelName BACnetApplicationTagCharacterString) BACnetConstructedDataModelNameBuilder
	// WithModelName adds ModelName (property field)
	WithModelName(BACnetApplicationTagCharacterString) BACnetConstructedDataModelNameBuilder
	// WithModelNameBuilder adds ModelName (property field) which is build by the builder
	WithModelNameBuilder(func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetConstructedDataModelNameBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataModelName or returns an error if something is wrong
	Build() (BACnetConstructedDataModelName, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataModelName
}

// NewBACnetConstructedDataModelNameBuilder() creates a BACnetConstructedDataModelNameBuilder
func NewBACnetConstructedDataModelNameBuilder() BACnetConstructedDataModelNameBuilder {
	return &_BACnetConstructedDataModelNameBuilder{_BACnetConstructedDataModelName: new(_BACnetConstructedDataModelName)}
}

type _BACnetConstructedDataModelNameBuilder struct {
	*_BACnetConstructedDataModelName

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataModelNameBuilder) = (*_BACnetConstructedDataModelNameBuilder)(nil)

func (b *_BACnetConstructedDataModelNameBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataModelNameBuilder) WithMandatoryFields(modelName BACnetApplicationTagCharacterString) BACnetConstructedDataModelNameBuilder {
	return b.WithModelName(modelName)
}

func (b *_BACnetConstructedDataModelNameBuilder) WithModelName(modelName BACnetApplicationTagCharacterString) BACnetConstructedDataModelNameBuilder {
	b.ModelName = modelName
	return b
}

func (b *_BACnetConstructedDataModelNameBuilder) WithModelNameBuilder(builderSupplier func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetConstructedDataModelNameBuilder {
	builder := builderSupplier(b.ModelName.CreateBACnetApplicationTagCharacterStringBuilder())
	var err error
	b.ModelName, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagCharacterStringBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataModelNameBuilder) Build() (BACnetConstructedDataModelName, error) {
	if b.ModelName == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'modelName' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataModelName.deepCopy(), nil
}

func (b *_BACnetConstructedDataModelNameBuilder) MustBuild() BACnetConstructedDataModelName {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataModelNameBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataModelNameBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataModelNameBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataModelNameBuilder().(*_BACnetConstructedDataModelNameBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataModelNameBuilder creates a BACnetConstructedDataModelNameBuilder
func (b *_BACnetConstructedDataModelName) CreateBACnetConstructedDataModelNameBuilder() BACnetConstructedDataModelNameBuilder {
	if b == nil {
		return NewBACnetConstructedDataModelNameBuilder()
	}
	return &_BACnetConstructedDataModelNameBuilder{_BACnetConstructedDataModelName: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataModelName) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataModelName) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MODEL_NAME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataModelName) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataModelName) GetModelName() BACnetApplicationTagCharacterString {
	return m.ModelName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataModelName) GetActualValue() BACnetApplicationTagCharacterString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagCharacterString(m.GetModelName())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataModelName(structType any) BACnetConstructedDataModelName {
	if casted, ok := structType.(BACnetConstructedDataModelName); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataModelName); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataModelName) GetTypeName() string {
	return "BACnetConstructedDataModelName"
}

func (m *_BACnetConstructedDataModelName) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (modelName)
	lengthInBits += m.ModelName.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataModelName) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataModelName) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataModelName BACnetConstructedDataModelName, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataModelName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataModelName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	modelName, err := ReadSimpleField[BACnetApplicationTagCharacterString](ctx, "modelName", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'modelName' field"))
	}
	m.ModelName = modelName

	actualValue, err := ReadVirtualField[BACnetApplicationTagCharacterString](ctx, "actualValue", (*BACnetApplicationTagCharacterString)(nil), modelName)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataModelName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataModelName")
	}

	return m, nil
}

func (m *_BACnetConstructedDataModelName) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataModelName) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataModelName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataModelName")
		}

		if err := WriteSimpleField[BACnetApplicationTagCharacterString](ctx, "modelName", m.GetModelName(), WriteComplex[BACnetApplicationTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'modelName' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataModelName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataModelName")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataModelName) IsBACnetConstructedDataModelName() {}

func (m *_BACnetConstructedDataModelName) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataModelName) deepCopy() *_BACnetConstructedDataModelName {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataModelNameCopy := &_BACnetConstructedDataModelName{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagCharacterString](m.ModelName),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataModelNameCopy
}

func (m *_BACnetConstructedDataModelName) String() string {
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
