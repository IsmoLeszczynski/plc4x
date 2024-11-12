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

// BACnetConstructedDataLastStateChange is the corresponding interface of BACnetConstructedDataLastStateChange
type BACnetConstructedDataLastStateChange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetLastStateChange returns LastStateChange (property field)
	GetLastStateChange() BACnetTimerTransitionTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetTimerTransitionTagged
	// IsBACnetConstructedDataLastStateChange is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLastStateChange()
	// CreateBuilder creates a BACnetConstructedDataLastStateChangeBuilder
	CreateBACnetConstructedDataLastStateChangeBuilder() BACnetConstructedDataLastStateChangeBuilder
}

// _BACnetConstructedDataLastStateChange is the data-structure of this message
type _BACnetConstructedDataLastStateChange struct {
	BACnetConstructedDataContract
	LastStateChange BACnetTimerTransitionTagged
}

var _ BACnetConstructedDataLastStateChange = (*_BACnetConstructedDataLastStateChange)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLastStateChange)(nil)

// NewBACnetConstructedDataLastStateChange factory function for _BACnetConstructedDataLastStateChange
func NewBACnetConstructedDataLastStateChange(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, lastStateChange BACnetTimerTransitionTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLastStateChange {
	if lastStateChange == nil {
		panic("lastStateChange of type BACnetTimerTransitionTagged for BACnetConstructedDataLastStateChange must not be nil")
	}
	_result := &_BACnetConstructedDataLastStateChange{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LastStateChange:               lastStateChange,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLastStateChangeBuilder is a builder for BACnetConstructedDataLastStateChange
type BACnetConstructedDataLastStateChangeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(lastStateChange BACnetTimerTransitionTagged) BACnetConstructedDataLastStateChangeBuilder
	// WithLastStateChange adds LastStateChange (property field)
	WithLastStateChange(BACnetTimerTransitionTagged) BACnetConstructedDataLastStateChangeBuilder
	// WithLastStateChangeBuilder adds LastStateChange (property field) which is build by the builder
	WithLastStateChangeBuilder(func(BACnetTimerTransitionTaggedBuilder) BACnetTimerTransitionTaggedBuilder) BACnetConstructedDataLastStateChangeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataLastStateChange or returns an error if something is wrong
	Build() (BACnetConstructedDataLastStateChange, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLastStateChange
}

// NewBACnetConstructedDataLastStateChangeBuilder() creates a BACnetConstructedDataLastStateChangeBuilder
func NewBACnetConstructedDataLastStateChangeBuilder() BACnetConstructedDataLastStateChangeBuilder {
	return &_BACnetConstructedDataLastStateChangeBuilder{_BACnetConstructedDataLastStateChange: new(_BACnetConstructedDataLastStateChange)}
}

type _BACnetConstructedDataLastStateChangeBuilder struct {
	*_BACnetConstructedDataLastStateChange

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLastStateChangeBuilder) = (*_BACnetConstructedDataLastStateChangeBuilder)(nil)

func (b *_BACnetConstructedDataLastStateChangeBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataLastStateChangeBuilder) WithMandatoryFields(lastStateChange BACnetTimerTransitionTagged) BACnetConstructedDataLastStateChangeBuilder {
	return b.WithLastStateChange(lastStateChange)
}

func (b *_BACnetConstructedDataLastStateChangeBuilder) WithLastStateChange(lastStateChange BACnetTimerTransitionTagged) BACnetConstructedDataLastStateChangeBuilder {
	b.LastStateChange = lastStateChange
	return b
}

func (b *_BACnetConstructedDataLastStateChangeBuilder) WithLastStateChangeBuilder(builderSupplier func(BACnetTimerTransitionTaggedBuilder) BACnetTimerTransitionTaggedBuilder) BACnetConstructedDataLastStateChangeBuilder {
	builder := builderSupplier(b.LastStateChange.CreateBACnetTimerTransitionTaggedBuilder())
	var err error
	b.LastStateChange, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTimerTransitionTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataLastStateChangeBuilder) Build() (BACnetConstructedDataLastStateChange, error) {
	if b.LastStateChange == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'lastStateChange' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLastStateChange.deepCopy(), nil
}

func (b *_BACnetConstructedDataLastStateChangeBuilder) MustBuild() BACnetConstructedDataLastStateChange {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataLastStateChangeBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLastStateChangeBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLastStateChangeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLastStateChangeBuilder().(*_BACnetConstructedDataLastStateChangeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLastStateChangeBuilder creates a BACnetConstructedDataLastStateChangeBuilder
func (b *_BACnetConstructedDataLastStateChange) CreateBACnetConstructedDataLastStateChangeBuilder() BACnetConstructedDataLastStateChangeBuilder {
	if b == nil {
		return NewBACnetConstructedDataLastStateChangeBuilder()
	}
	return &_BACnetConstructedDataLastStateChangeBuilder{_BACnetConstructedDataLastStateChange: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLastStateChange) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLastStateChange) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LAST_STATE_CHANGE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLastStateChange) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLastStateChange) GetLastStateChange() BACnetTimerTransitionTagged {
	return m.LastStateChange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLastStateChange) GetActualValue() BACnetTimerTransitionTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetTimerTransitionTagged(m.GetLastStateChange())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLastStateChange(structType any) BACnetConstructedDataLastStateChange {
	if casted, ok := structType.(BACnetConstructedDataLastStateChange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLastStateChange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLastStateChange) GetTypeName() string {
	return "BACnetConstructedDataLastStateChange"
}

func (m *_BACnetConstructedDataLastStateChange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (lastStateChange)
	lengthInBits += m.LastStateChange.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLastStateChange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLastStateChange) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLastStateChange BACnetConstructedDataLastStateChange, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLastStateChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLastStateChange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lastStateChange, err := ReadSimpleField[BACnetTimerTransitionTagged](ctx, "lastStateChange", ReadComplex[BACnetTimerTransitionTagged](BACnetTimerTransitionTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastStateChange' field"))
	}
	m.LastStateChange = lastStateChange

	actualValue, err := ReadVirtualField[BACnetTimerTransitionTagged](ctx, "actualValue", (*BACnetTimerTransitionTagged)(nil), lastStateChange)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLastStateChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLastStateChange")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLastStateChange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLastStateChange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLastStateChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLastStateChange")
		}

		if err := WriteSimpleField[BACnetTimerTransitionTagged](ctx, "lastStateChange", m.GetLastStateChange(), WriteComplex[BACnetTimerTransitionTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastStateChange' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLastStateChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLastStateChange")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLastStateChange) IsBACnetConstructedDataLastStateChange() {}

func (m *_BACnetConstructedDataLastStateChange) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLastStateChange) deepCopy() *_BACnetConstructedDataLastStateChange {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLastStateChangeCopy := &_BACnetConstructedDataLastStateChange{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetTimerTransitionTagged](m.LastStateChange),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLastStateChangeCopy
}

func (m *_BACnetConstructedDataLastStateChange) String() string {
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
