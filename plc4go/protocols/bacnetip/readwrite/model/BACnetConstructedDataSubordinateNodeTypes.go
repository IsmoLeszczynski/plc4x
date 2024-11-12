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

// BACnetConstructedDataSubordinateNodeTypes is the corresponding interface of BACnetConstructedDataSubordinateNodeTypes
type BACnetConstructedDataSubordinateNodeTypes interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetSubordinateNodeTypes returns SubordinateNodeTypes (property field)
	GetSubordinateNodeTypes() []BACnetNodeTypeTagged
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// IsBACnetConstructedDataSubordinateNodeTypes is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataSubordinateNodeTypes()
	// CreateBuilder creates a BACnetConstructedDataSubordinateNodeTypesBuilder
	CreateBACnetConstructedDataSubordinateNodeTypesBuilder() BACnetConstructedDataSubordinateNodeTypesBuilder
}

// _BACnetConstructedDataSubordinateNodeTypes is the data-structure of this message
type _BACnetConstructedDataSubordinateNodeTypes struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	SubordinateNodeTypes []BACnetNodeTypeTagged
}

var _ BACnetConstructedDataSubordinateNodeTypes = (*_BACnetConstructedDataSubordinateNodeTypes)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataSubordinateNodeTypes)(nil)

// NewBACnetConstructedDataSubordinateNodeTypes factory function for _BACnetConstructedDataSubordinateNodeTypes
func NewBACnetConstructedDataSubordinateNodeTypes(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfDataElements BACnetApplicationTagUnsignedInteger, subordinateNodeTypes []BACnetNodeTypeTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSubordinateNodeTypes {
	_result := &_BACnetConstructedDataSubordinateNodeTypes{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		SubordinateNodeTypes:          subordinateNodeTypes,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataSubordinateNodeTypesBuilder is a builder for BACnetConstructedDataSubordinateNodeTypes
type BACnetConstructedDataSubordinateNodeTypesBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(subordinateNodeTypes []BACnetNodeTypeTagged) BACnetConstructedDataSubordinateNodeTypesBuilder
	// WithNumberOfDataElements adds NumberOfDataElements (property field)
	WithOptionalNumberOfDataElements(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataSubordinateNodeTypesBuilder
	// WithOptionalNumberOfDataElementsBuilder adds NumberOfDataElements (property field) which is build by the builder
	WithOptionalNumberOfDataElementsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataSubordinateNodeTypesBuilder
	// WithSubordinateNodeTypes adds SubordinateNodeTypes (property field)
	WithSubordinateNodeTypes(...BACnetNodeTypeTagged) BACnetConstructedDataSubordinateNodeTypesBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataSubordinateNodeTypes or returns an error if something is wrong
	Build() (BACnetConstructedDataSubordinateNodeTypes, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataSubordinateNodeTypes
}

// NewBACnetConstructedDataSubordinateNodeTypesBuilder() creates a BACnetConstructedDataSubordinateNodeTypesBuilder
func NewBACnetConstructedDataSubordinateNodeTypesBuilder() BACnetConstructedDataSubordinateNodeTypesBuilder {
	return &_BACnetConstructedDataSubordinateNodeTypesBuilder{_BACnetConstructedDataSubordinateNodeTypes: new(_BACnetConstructedDataSubordinateNodeTypes)}
}

type _BACnetConstructedDataSubordinateNodeTypesBuilder struct {
	*_BACnetConstructedDataSubordinateNodeTypes

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataSubordinateNodeTypesBuilder) = (*_BACnetConstructedDataSubordinateNodeTypesBuilder)(nil)

func (b *_BACnetConstructedDataSubordinateNodeTypesBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataSubordinateNodeTypesBuilder) WithMandatoryFields(subordinateNodeTypes []BACnetNodeTypeTagged) BACnetConstructedDataSubordinateNodeTypesBuilder {
	return b.WithSubordinateNodeTypes(subordinateNodeTypes...)
}

func (b *_BACnetConstructedDataSubordinateNodeTypesBuilder) WithOptionalNumberOfDataElements(numberOfDataElements BACnetApplicationTagUnsignedInteger) BACnetConstructedDataSubordinateNodeTypesBuilder {
	b.NumberOfDataElements = numberOfDataElements
	return b
}

func (b *_BACnetConstructedDataSubordinateNodeTypesBuilder) WithOptionalNumberOfDataElementsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataSubordinateNodeTypesBuilder {
	builder := builderSupplier(b.NumberOfDataElements.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.NumberOfDataElements, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataSubordinateNodeTypesBuilder) WithSubordinateNodeTypes(subordinateNodeTypes ...BACnetNodeTypeTagged) BACnetConstructedDataSubordinateNodeTypesBuilder {
	b.SubordinateNodeTypes = subordinateNodeTypes
	return b
}

func (b *_BACnetConstructedDataSubordinateNodeTypesBuilder) Build() (BACnetConstructedDataSubordinateNodeTypes, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataSubordinateNodeTypes.deepCopy(), nil
}

func (b *_BACnetConstructedDataSubordinateNodeTypesBuilder) MustBuild() BACnetConstructedDataSubordinateNodeTypes {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataSubordinateNodeTypesBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataSubordinateNodeTypesBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataSubordinateNodeTypesBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataSubordinateNodeTypesBuilder().(*_BACnetConstructedDataSubordinateNodeTypesBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataSubordinateNodeTypesBuilder creates a BACnetConstructedDataSubordinateNodeTypesBuilder
func (b *_BACnetConstructedDataSubordinateNodeTypes) CreateBACnetConstructedDataSubordinateNodeTypesBuilder() BACnetConstructedDataSubordinateNodeTypesBuilder {
	if b == nil {
		return NewBACnetConstructedDataSubordinateNodeTypesBuilder()
	}
	return &_BACnetConstructedDataSubordinateNodeTypesBuilder{_BACnetConstructedDataSubordinateNodeTypes: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSubordinateNodeTypes) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataSubordinateNodeTypes) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SUBORDINATE_NODE_TYPES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSubordinateNodeTypes) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSubordinateNodeTypes) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataSubordinateNodeTypes) GetSubordinateNodeTypes() []BACnetNodeTypeTagged {
	return m.SubordinateNodeTypes
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataSubordinateNodeTypes) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.GetNumberOfDataElements()
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataSubordinateNodeTypes(structType any) BACnetConstructedDataSubordinateNodeTypes {
	if casted, ok := structType.(BACnetConstructedDataSubordinateNodeTypes); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSubordinateNodeTypes); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSubordinateNodeTypes) GetTypeName() string {
	return "BACnetConstructedDataSubordinateNodeTypes"
}

func (m *_BACnetConstructedDataSubordinateNodeTypes) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.SubordinateNodeTypes) > 0 {
		for _, element := range m.SubordinateNodeTypes {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataSubordinateNodeTypes) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataSubordinateNodeTypes) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataSubordinateNodeTypes BACnetConstructedDataSubordinateNodeTypes, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSubordinateNodeTypes"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSubordinateNodeTypes")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zero, err := ReadVirtualField[uint64](ctx, "zero", (*uint64)(nil), uint64(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zero' field"))
	}
	_ = zero

	var numberOfDataElements BACnetApplicationTagUnsignedInteger
	_numberOfDataElements, err := ReadOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfDataElements' field"))
	}
	if _numberOfDataElements != nil {
		numberOfDataElements = *_numberOfDataElements
		m.NumberOfDataElements = numberOfDataElements
	}

	subordinateNodeTypes, err := ReadTerminatedArrayField[BACnetNodeTypeTagged](ctx, "subordinateNodeTypes", ReadComplex[BACnetNodeTypeTagged](BACnetNodeTypeTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subordinateNodeTypes' field"))
	}
	m.SubordinateNodeTypes = subordinateNodeTypes

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSubordinateNodeTypes"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSubordinateNodeTypes")
	}

	return m, nil
}

func (m *_BACnetConstructedDataSubordinateNodeTypes) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataSubordinateNodeTypes) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSubordinateNodeTypes"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSubordinateNodeTypes")
		}
		// Virtual field
		zero := m.GetZero()
		_ = zero
		if _zeroErr := writeBuffer.WriteVirtual(ctx, "zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		if err := WriteOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", GetRef(m.GetNumberOfDataElements()), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'numberOfDataElements' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "subordinateNodeTypes", m.GetSubordinateNodeTypes(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'subordinateNodeTypes' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSubordinateNodeTypes"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSubordinateNodeTypes")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSubordinateNodeTypes) IsBACnetConstructedDataSubordinateNodeTypes() {}

func (m *_BACnetConstructedDataSubordinateNodeTypes) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataSubordinateNodeTypes) deepCopy() *_BACnetConstructedDataSubordinateNodeTypes {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataSubordinateNodeTypesCopy := &_BACnetConstructedDataSubordinateNodeTypes{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.NumberOfDataElements),
		utils.DeepCopySlice[BACnetNodeTypeTagged, BACnetNodeTypeTagged](m.SubordinateNodeTypes),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataSubordinateNodeTypesCopy
}

func (m *_BACnetConstructedDataSubordinateNodeTypes) String() string {
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
