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

// BACnetConstructedDataLastCredentialAdded is the corresponding interface of BACnetConstructedDataLastCredentialAdded
type BACnetConstructedDataLastCredentialAdded interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetLastCredentialAdded returns LastCredentialAdded (property field)
	GetLastCredentialAdded() BACnetDeviceObjectReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDeviceObjectReference
	// IsBACnetConstructedDataLastCredentialAdded is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLastCredentialAdded()
	// CreateBuilder creates a BACnetConstructedDataLastCredentialAddedBuilder
	CreateBACnetConstructedDataLastCredentialAddedBuilder() BACnetConstructedDataLastCredentialAddedBuilder
}

// _BACnetConstructedDataLastCredentialAdded is the data-structure of this message
type _BACnetConstructedDataLastCredentialAdded struct {
	BACnetConstructedDataContract
	LastCredentialAdded BACnetDeviceObjectReference
}

var _ BACnetConstructedDataLastCredentialAdded = (*_BACnetConstructedDataLastCredentialAdded)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLastCredentialAdded)(nil)

// NewBACnetConstructedDataLastCredentialAdded factory function for _BACnetConstructedDataLastCredentialAdded
func NewBACnetConstructedDataLastCredentialAdded(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, lastCredentialAdded BACnetDeviceObjectReference, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLastCredentialAdded {
	if lastCredentialAdded == nil {
		panic("lastCredentialAdded of type BACnetDeviceObjectReference for BACnetConstructedDataLastCredentialAdded must not be nil")
	}
	_result := &_BACnetConstructedDataLastCredentialAdded{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LastCredentialAdded:           lastCredentialAdded,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLastCredentialAddedBuilder is a builder for BACnetConstructedDataLastCredentialAdded
type BACnetConstructedDataLastCredentialAddedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(lastCredentialAdded BACnetDeviceObjectReference) BACnetConstructedDataLastCredentialAddedBuilder
	// WithLastCredentialAdded adds LastCredentialAdded (property field)
	WithLastCredentialAdded(BACnetDeviceObjectReference) BACnetConstructedDataLastCredentialAddedBuilder
	// WithLastCredentialAddedBuilder adds LastCredentialAdded (property field) which is build by the builder
	WithLastCredentialAddedBuilder(func(BACnetDeviceObjectReferenceBuilder) BACnetDeviceObjectReferenceBuilder) BACnetConstructedDataLastCredentialAddedBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataLastCredentialAdded or returns an error if something is wrong
	Build() (BACnetConstructedDataLastCredentialAdded, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLastCredentialAdded
}

// NewBACnetConstructedDataLastCredentialAddedBuilder() creates a BACnetConstructedDataLastCredentialAddedBuilder
func NewBACnetConstructedDataLastCredentialAddedBuilder() BACnetConstructedDataLastCredentialAddedBuilder {
	return &_BACnetConstructedDataLastCredentialAddedBuilder{_BACnetConstructedDataLastCredentialAdded: new(_BACnetConstructedDataLastCredentialAdded)}
}

type _BACnetConstructedDataLastCredentialAddedBuilder struct {
	*_BACnetConstructedDataLastCredentialAdded

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLastCredentialAddedBuilder) = (*_BACnetConstructedDataLastCredentialAddedBuilder)(nil)

func (b *_BACnetConstructedDataLastCredentialAddedBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataLastCredentialAddedBuilder) WithMandatoryFields(lastCredentialAdded BACnetDeviceObjectReference) BACnetConstructedDataLastCredentialAddedBuilder {
	return b.WithLastCredentialAdded(lastCredentialAdded)
}

func (b *_BACnetConstructedDataLastCredentialAddedBuilder) WithLastCredentialAdded(lastCredentialAdded BACnetDeviceObjectReference) BACnetConstructedDataLastCredentialAddedBuilder {
	b.LastCredentialAdded = lastCredentialAdded
	return b
}

func (b *_BACnetConstructedDataLastCredentialAddedBuilder) WithLastCredentialAddedBuilder(builderSupplier func(BACnetDeviceObjectReferenceBuilder) BACnetDeviceObjectReferenceBuilder) BACnetConstructedDataLastCredentialAddedBuilder {
	builder := builderSupplier(b.LastCredentialAdded.CreateBACnetDeviceObjectReferenceBuilder())
	var err error
	b.LastCredentialAdded, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDeviceObjectReferenceBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataLastCredentialAddedBuilder) Build() (BACnetConstructedDataLastCredentialAdded, error) {
	if b.LastCredentialAdded == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'lastCredentialAdded' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLastCredentialAdded.deepCopy(), nil
}

func (b *_BACnetConstructedDataLastCredentialAddedBuilder) MustBuild() BACnetConstructedDataLastCredentialAdded {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataLastCredentialAddedBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLastCredentialAddedBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLastCredentialAddedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLastCredentialAddedBuilder().(*_BACnetConstructedDataLastCredentialAddedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLastCredentialAddedBuilder creates a BACnetConstructedDataLastCredentialAddedBuilder
func (b *_BACnetConstructedDataLastCredentialAdded) CreateBACnetConstructedDataLastCredentialAddedBuilder() BACnetConstructedDataLastCredentialAddedBuilder {
	if b == nil {
		return NewBACnetConstructedDataLastCredentialAddedBuilder()
	}
	return &_BACnetConstructedDataLastCredentialAddedBuilder{_BACnetConstructedDataLastCredentialAdded: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLastCredentialAdded) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLastCredentialAdded) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LAST_CREDENTIAL_ADDED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLastCredentialAdded) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLastCredentialAdded) GetLastCredentialAdded() BACnetDeviceObjectReference {
	return m.LastCredentialAdded
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLastCredentialAdded) GetActualValue() BACnetDeviceObjectReference {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDeviceObjectReference(m.GetLastCredentialAdded())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLastCredentialAdded(structType any) BACnetConstructedDataLastCredentialAdded {
	if casted, ok := structType.(BACnetConstructedDataLastCredentialAdded); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLastCredentialAdded); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLastCredentialAdded) GetTypeName() string {
	return "BACnetConstructedDataLastCredentialAdded"
}

func (m *_BACnetConstructedDataLastCredentialAdded) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (lastCredentialAdded)
	lengthInBits += m.LastCredentialAdded.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLastCredentialAdded) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLastCredentialAdded) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLastCredentialAdded BACnetConstructedDataLastCredentialAdded, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLastCredentialAdded"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLastCredentialAdded")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lastCredentialAdded, err := ReadSimpleField[BACnetDeviceObjectReference](ctx, "lastCredentialAdded", ReadComplex[BACnetDeviceObjectReference](BACnetDeviceObjectReferenceParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastCredentialAdded' field"))
	}
	m.LastCredentialAdded = lastCredentialAdded

	actualValue, err := ReadVirtualField[BACnetDeviceObjectReference](ctx, "actualValue", (*BACnetDeviceObjectReference)(nil), lastCredentialAdded)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLastCredentialAdded"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLastCredentialAdded")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLastCredentialAdded) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLastCredentialAdded) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLastCredentialAdded"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLastCredentialAdded")
		}

		if err := WriteSimpleField[BACnetDeviceObjectReference](ctx, "lastCredentialAdded", m.GetLastCredentialAdded(), WriteComplex[BACnetDeviceObjectReference](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastCredentialAdded' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLastCredentialAdded"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLastCredentialAdded")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLastCredentialAdded) IsBACnetConstructedDataLastCredentialAdded() {}

func (m *_BACnetConstructedDataLastCredentialAdded) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLastCredentialAdded) deepCopy() *_BACnetConstructedDataLastCredentialAdded {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLastCredentialAddedCopy := &_BACnetConstructedDataLastCredentialAdded{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetDeviceObjectReference](m.LastCredentialAdded),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLastCredentialAddedCopy
}

func (m *_BACnetConstructedDataLastCredentialAdded) String() string {
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
