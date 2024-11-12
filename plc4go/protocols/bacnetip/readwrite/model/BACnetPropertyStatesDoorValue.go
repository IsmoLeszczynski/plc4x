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

// BACnetPropertyStatesDoorValue is the corresponding interface of BACnetPropertyStatesDoorValue
type BACnetPropertyStatesDoorValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetDoorValue returns DoorValue (property field)
	GetDoorValue() BACnetDoorValueTagged
	// IsBACnetPropertyStatesDoorValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesDoorValue()
	// CreateBuilder creates a BACnetPropertyStatesDoorValueBuilder
	CreateBACnetPropertyStatesDoorValueBuilder() BACnetPropertyStatesDoorValueBuilder
}

// _BACnetPropertyStatesDoorValue is the data-structure of this message
type _BACnetPropertyStatesDoorValue struct {
	BACnetPropertyStatesContract
	DoorValue BACnetDoorValueTagged
}

var _ BACnetPropertyStatesDoorValue = (*_BACnetPropertyStatesDoorValue)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesDoorValue)(nil)

// NewBACnetPropertyStatesDoorValue factory function for _BACnetPropertyStatesDoorValue
func NewBACnetPropertyStatesDoorValue(peekedTagHeader BACnetTagHeader, doorValue BACnetDoorValueTagged) *_BACnetPropertyStatesDoorValue {
	if doorValue == nil {
		panic("doorValue of type BACnetDoorValueTagged for BACnetPropertyStatesDoorValue must not be nil")
	}
	_result := &_BACnetPropertyStatesDoorValue{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		DoorValue:                    doorValue,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesDoorValueBuilder is a builder for BACnetPropertyStatesDoorValue
type BACnetPropertyStatesDoorValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(doorValue BACnetDoorValueTagged) BACnetPropertyStatesDoorValueBuilder
	// WithDoorValue adds DoorValue (property field)
	WithDoorValue(BACnetDoorValueTagged) BACnetPropertyStatesDoorValueBuilder
	// WithDoorValueBuilder adds DoorValue (property field) which is build by the builder
	WithDoorValueBuilder(func(BACnetDoorValueTaggedBuilder) BACnetDoorValueTaggedBuilder) BACnetPropertyStatesDoorValueBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetPropertyStatesBuilder
	// Build builds the BACnetPropertyStatesDoorValue or returns an error if something is wrong
	Build() (BACnetPropertyStatesDoorValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesDoorValue
}

// NewBACnetPropertyStatesDoorValueBuilder() creates a BACnetPropertyStatesDoorValueBuilder
func NewBACnetPropertyStatesDoorValueBuilder() BACnetPropertyStatesDoorValueBuilder {
	return &_BACnetPropertyStatesDoorValueBuilder{_BACnetPropertyStatesDoorValue: new(_BACnetPropertyStatesDoorValue)}
}

type _BACnetPropertyStatesDoorValueBuilder struct {
	*_BACnetPropertyStatesDoorValue

	parentBuilder *_BACnetPropertyStatesBuilder

	err *utils.MultiError
}

var _ (BACnetPropertyStatesDoorValueBuilder) = (*_BACnetPropertyStatesDoorValueBuilder)(nil)

func (b *_BACnetPropertyStatesDoorValueBuilder) setParent(contract BACnetPropertyStatesContract) {
	b.BACnetPropertyStatesContract = contract
}

func (b *_BACnetPropertyStatesDoorValueBuilder) WithMandatoryFields(doorValue BACnetDoorValueTagged) BACnetPropertyStatesDoorValueBuilder {
	return b.WithDoorValue(doorValue)
}

func (b *_BACnetPropertyStatesDoorValueBuilder) WithDoorValue(doorValue BACnetDoorValueTagged) BACnetPropertyStatesDoorValueBuilder {
	b.DoorValue = doorValue
	return b
}

func (b *_BACnetPropertyStatesDoorValueBuilder) WithDoorValueBuilder(builderSupplier func(BACnetDoorValueTaggedBuilder) BACnetDoorValueTaggedBuilder) BACnetPropertyStatesDoorValueBuilder {
	builder := builderSupplier(b.DoorValue.CreateBACnetDoorValueTaggedBuilder())
	var err error
	b.DoorValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDoorValueTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyStatesDoorValueBuilder) Build() (BACnetPropertyStatesDoorValue, error) {
	if b.DoorValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'doorValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPropertyStatesDoorValue.deepCopy(), nil
}

func (b *_BACnetPropertyStatesDoorValueBuilder) MustBuild() BACnetPropertyStatesDoorValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetPropertyStatesDoorValueBuilder) Done() BACnetPropertyStatesBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetPropertyStatesBuilder().(*_BACnetPropertyStatesBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetPropertyStatesDoorValueBuilder) buildForBACnetPropertyStates() (BACnetPropertyStates, error) {
	return b.Build()
}

func (b *_BACnetPropertyStatesDoorValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPropertyStatesDoorValueBuilder().(*_BACnetPropertyStatesDoorValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPropertyStatesDoorValueBuilder creates a BACnetPropertyStatesDoorValueBuilder
func (b *_BACnetPropertyStatesDoorValue) CreateBACnetPropertyStatesDoorValueBuilder() BACnetPropertyStatesDoorValueBuilder {
	if b == nil {
		return NewBACnetPropertyStatesDoorValueBuilder()
	}
	return &_BACnetPropertyStatesDoorValueBuilder{_BACnetPropertyStatesDoorValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesDoorValue) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesDoorValue) GetDoorValue() BACnetDoorValueTagged {
	return m.DoorValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesDoorValue(structType any) BACnetPropertyStatesDoorValue {
	if casted, ok := structType.(BACnetPropertyStatesDoorValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesDoorValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesDoorValue) GetTypeName() string {
	return "BACnetPropertyStatesDoorValue"
}

func (m *_BACnetPropertyStatesDoorValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (doorValue)
	lengthInBits += m.DoorValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesDoorValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesDoorValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesDoorValue BACnetPropertyStatesDoorValue, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesDoorValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesDoorValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	doorValue, err := ReadSimpleField[BACnetDoorValueTagged](ctx, "doorValue", ReadComplex[BACnetDoorValueTagged](BACnetDoorValueTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'doorValue' field"))
	}
	m.DoorValue = doorValue

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesDoorValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesDoorValue")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesDoorValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesDoorValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesDoorValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesDoorValue")
		}

		if err := WriteSimpleField[BACnetDoorValueTagged](ctx, "doorValue", m.GetDoorValue(), WriteComplex[BACnetDoorValueTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'doorValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesDoorValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesDoorValue")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesDoorValue) IsBACnetPropertyStatesDoorValue() {}

func (m *_BACnetPropertyStatesDoorValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesDoorValue) deepCopy() *_BACnetPropertyStatesDoorValue {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesDoorValueCopy := &_BACnetPropertyStatesDoorValue{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		utils.DeepCopy[BACnetDoorValueTagged](m.DoorValue),
	}
	m.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesDoorValueCopy
}

func (m *_BACnetPropertyStatesDoorValue) String() string {
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
