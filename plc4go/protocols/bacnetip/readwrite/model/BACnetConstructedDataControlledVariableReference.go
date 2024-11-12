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

// BACnetConstructedDataControlledVariableReference is the corresponding interface of BACnetConstructedDataControlledVariableReference
type BACnetConstructedDataControlledVariableReference interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetControlledVariableReference returns ControlledVariableReference (property field)
	GetControlledVariableReference() BACnetObjectPropertyReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetObjectPropertyReference
	// IsBACnetConstructedDataControlledVariableReference is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataControlledVariableReference()
	// CreateBuilder creates a BACnetConstructedDataControlledVariableReferenceBuilder
	CreateBACnetConstructedDataControlledVariableReferenceBuilder() BACnetConstructedDataControlledVariableReferenceBuilder
}

// _BACnetConstructedDataControlledVariableReference is the data-structure of this message
type _BACnetConstructedDataControlledVariableReference struct {
	BACnetConstructedDataContract
	ControlledVariableReference BACnetObjectPropertyReference
}

var _ BACnetConstructedDataControlledVariableReference = (*_BACnetConstructedDataControlledVariableReference)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataControlledVariableReference)(nil)

// NewBACnetConstructedDataControlledVariableReference factory function for _BACnetConstructedDataControlledVariableReference
func NewBACnetConstructedDataControlledVariableReference(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, controlledVariableReference BACnetObjectPropertyReference, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataControlledVariableReference {
	if controlledVariableReference == nil {
		panic("controlledVariableReference of type BACnetObjectPropertyReference for BACnetConstructedDataControlledVariableReference must not be nil")
	}
	_result := &_BACnetConstructedDataControlledVariableReference{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ControlledVariableReference:   controlledVariableReference,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataControlledVariableReferenceBuilder is a builder for BACnetConstructedDataControlledVariableReference
type BACnetConstructedDataControlledVariableReferenceBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(controlledVariableReference BACnetObjectPropertyReference) BACnetConstructedDataControlledVariableReferenceBuilder
	// WithControlledVariableReference adds ControlledVariableReference (property field)
	WithControlledVariableReference(BACnetObjectPropertyReference) BACnetConstructedDataControlledVariableReferenceBuilder
	// WithControlledVariableReferenceBuilder adds ControlledVariableReference (property field) which is build by the builder
	WithControlledVariableReferenceBuilder(func(BACnetObjectPropertyReferenceBuilder) BACnetObjectPropertyReferenceBuilder) BACnetConstructedDataControlledVariableReferenceBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataControlledVariableReference or returns an error if something is wrong
	Build() (BACnetConstructedDataControlledVariableReference, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataControlledVariableReference
}

// NewBACnetConstructedDataControlledVariableReferenceBuilder() creates a BACnetConstructedDataControlledVariableReferenceBuilder
func NewBACnetConstructedDataControlledVariableReferenceBuilder() BACnetConstructedDataControlledVariableReferenceBuilder {
	return &_BACnetConstructedDataControlledVariableReferenceBuilder{_BACnetConstructedDataControlledVariableReference: new(_BACnetConstructedDataControlledVariableReference)}
}

type _BACnetConstructedDataControlledVariableReferenceBuilder struct {
	*_BACnetConstructedDataControlledVariableReference

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataControlledVariableReferenceBuilder) = (*_BACnetConstructedDataControlledVariableReferenceBuilder)(nil)

func (b *_BACnetConstructedDataControlledVariableReferenceBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataControlledVariableReferenceBuilder) WithMandatoryFields(controlledVariableReference BACnetObjectPropertyReference) BACnetConstructedDataControlledVariableReferenceBuilder {
	return b.WithControlledVariableReference(controlledVariableReference)
}

func (b *_BACnetConstructedDataControlledVariableReferenceBuilder) WithControlledVariableReference(controlledVariableReference BACnetObjectPropertyReference) BACnetConstructedDataControlledVariableReferenceBuilder {
	b.ControlledVariableReference = controlledVariableReference
	return b
}

func (b *_BACnetConstructedDataControlledVariableReferenceBuilder) WithControlledVariableReferenceBuilder(builderSupplier func(BACnetObjectPropertyReferenceBuilder) BACnetObjectPropertyReferenceBuilder) BACnetConstructedDataControlledVariableReferenceBuilder {
	builder := builderSupplier(b.ControlledVariableReference.CreateBACnetObjectPropertyReferenceBuilder())
	var err error
	b.ControlledVariableReference, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetObjectPropertyReferenceBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataControlledVariableReferenceBuilder) Build() (BACnetConstructedDataControlledVariableReference, error) {
	if b.ControlledVariableReference == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'controlledVariableReference' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataControlledVariableReference.deepCopy(), nil
}

func (b *_BACnetConstructedDataControlledVariableReferenceBuilder) MustBuild() BACnetConstructedDataControlledVariableReference {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataControlledVariableReferenceBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataControlledVariableReferenceBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataControlledVariableReferenceBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataControlledVariableReferenceBuilder().(*_BACnetConstructedDataControlledVariableReferenceBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataControlledVariableReferenceBuilder creates a BACnetConstructedDataControlledVariableReferenceBuilder
func (b *_BACnetConstructedDataControlledVariableReference) CreateBACnetConstructedDataControlledVariableReferenceBuilder() BACnetConstructedDataControlledVariableReferenceBuilder {
	if b == nil {
		return NewBACnetConstructedDataControlledVariableReferenceBuilder()
	}
	return &_BACnetConstructedDataControlledVariableReferenceBuilder{_BACnetConstructedDataControlledVariableReference: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataControlledVariableReference) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataControlledVariableReference) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CONTROLLED_VARIABLE_REFERENCE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataControlledVariableReference) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataControlledVariableReference) GetControlledVariableReference() BACnetObjectPropertyReference {
	return m.ControlledVariableReference
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataControlledVariableReference) GetActualValue() BACnetObjectPropertyReference {
	ctx := context.Background()
	_ = ctx
	return CastBACnetObjectPropertyReference(m.GetControlledVariableReference())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataControlledVariableReference(structType any) BACnetConstructedDataControlledVariableReference {
	if casted, ok := structType.(BACnetConstructedDataControlledVariableReference); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataControlledVariableReference); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataControlledVariableReference) GetTypeName() string {
	return "BACnetConstructedDataControlledVariableReference"
}

func (m *_BACnetConstructedDataControlledVariableReference) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (controlledVariableReference)
	lengthInBits += m.ControlledVariableReference.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataControlledVariableReference) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataControlledVariableReference) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataControlledVariableReference BACnetConstructedDataControlledVariableReference, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataControlledVariableReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataControlledVariableReference")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	controlledVariableReference, err := ReadSimpleField[BACnetObjectPropertyReference](ctx, "controlledVariableReference", ReadComplex[BACnetObjectPropertyReference](BACnetObjectPropertyReferenceParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'controlledVariableReference' field"))
	}
	m.ControlledVariableReference = controlledVariableReference

	actualValue, err := ReadVirtualField[BACnetObjectPropertyReference](ctx, "actualValue", (*BACnetObjectPropertyReference)(nil), controlledVariableReference)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataControlledVariableReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataControlledVariableReference")
	}

	return m, nil
}

func (m *_BACnetConstructedDataControlledVariableReference) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataControlledVariableReference) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataControlledVariableReference"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataControlledVariableReference")
		}

		if err := WriteSimpleField[BACnetObjectPropertyReference](ctx, "controlledVariableReference", m.GetControlledVariableReference(), WriteComplex[BACnetObjectPropertyReference](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'controlledVariableReference' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataControlledVariableReference"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataControlledVariableReference")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataControlledVariableReference) IsBACnetConstructedDataControlledVariableReference() {
}

func (m *_BACnetConstructedDataControlledVariableReference) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataControlledVariableReference) deepCopy() *_BACnetConstructedDataControlledVariableReference {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataControlledVariableReferenceCopy := &_BACnetConstructedDataControlledVariableReference{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetObjectPropertyReference](m.ControlledVariableReference),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataControlledVariableReferenceCopy
}

func (m *_BACnetConstructedDataControlledVariableReference) String() string {
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
