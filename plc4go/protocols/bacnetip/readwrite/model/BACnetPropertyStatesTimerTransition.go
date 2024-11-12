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

// BACnetPropertyStatesTimerTransition is the corresponding interface of BACnetPropertyStatesTimerTransition
type BACnetPropertyStatesTimerTransition interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetTimerTransition returns TimerTransition (property field)
	GetTimerTransition() BACnetTimerTransitionTagged
	// IsBACnetPropertyStatesTimerTransition is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesTimerTransition()
	// CreateBuilder creates a BACnetPropertyStatesTimerTransitionBuilder
	CreateBACnetPropertyStatesTimerTransitionBuilder() BACnetPropertyStatesTimerTransitionBuilder
}

// _BACnetPropertyStatesTimerTransition is the data-structure of this message
type _BACnetPropertyStatesTimerTransition struct {
	BACnetPropertyStatesContract
	TimerTransition BACnetTimerTransitionTagged
}

var _ BACnetPropertyStatesTimerTransition = (*_BACnetPropertyStatesTimerTransition)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesTimerTransition)(nil)

// NewBACnetPropertyStatesTimerTransition factory function for _BACnetPropertyStatesTimerTransition
func NewBACnetPropertyStatesTimerTransition(peekedTagHeader BACnetTagHeader, timerTransition BACnetTimerTransitionTagged) *_BACnetPropertyStatesTimerTransition {
	if timerTransition == nil {
		panic("timerTransition of type BACnetTimerTransitionTagged for BACnetPropertyStatesTimerTransition must not be nil")
	}
	_result := &_BACnetPropertyStatesTimerTransition{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		TimerTransition:              timerTransition,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesTimerTransitionBuilder is a builder for BACnetPropertyStatesTimerTransition
type BACnetPropertyStatesTimerTransitionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(timerTransition BACnetTimerTransitionTagged) BACnetPropertyStatesTimerTransitionBuilder
	// WithTimerTransition adds TimerTransition (property field)
	WithTimerTransition(BACnetTimerTransitionTagged) BACnetPropertyStatesTimerTransitionBuilder
	// WithTimerTransitionBuilder adds TimerTransition (property field) which is build by the builder
	WithTimerTransitionBuilder(func(BACnetTimerTransitionTaggedBuilder) BACnetTimerTransitionTaggedBuilder) BACnetPropertyStatesTimerTransitionBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetPropertyStatesBuilder
	// Build builds the BACnetPropertyStatesTimerTransition or returns an error if something is wrong
	Build() (BACnetPropertyStatesTimerTransition, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesTimerTransition
}

// NewBACnetPropertyStatesTimerTransitionBuilder() creates a BACnetPropertyStatesTimerTransitionBuilder
func NewBACnetPropertyStatesTimerTransitionBuilder() BACnetPropertyStatesTimerTransitionBuilder {
	return &_BACnetPropertyStatesTimerTransitionBuilder{_BACnetPropertyStatesTimerTransition: new(_BACnetPropertyStatesTimerTransition)}
}

type _BACnetPropertyStatesTimerTransitionBuilder struct {
	*_BACnetPropertyStatesTimerTransition

	parentBuilder *_BACnetPropertyStatesBuilder

	err *utils.MultiError
}

var _ (BACnetPropertyStatesTimerTransitionBuilder) = (*_BACnetPropertyStatesTimerTransitionBuilder)(nil)

func (b *_BACnetPropertyStatesTimerTransitionBuilder) setParent(contract BACnetPropertyStatesContract) {
	b.BACnetPropertyStatesContract = contract
}

func (b *_BACnetPropertyStatesTimerTransitionBuilder) WithMandatoryFields(timerTransition BACnetTimerTransitionTagged) BACnetPropertyStatesTimerTransitionBuilder {
	return b.WithTimerTransition(timerTransition)
}

func (b *_BACnetPropertyStatesTimerTransitionBuilder) WithTimerTransition(timerTransition BACnetTimerTransitionTagged) BACnetPropertyStatesTimerTransitionBuilder {
	b.TimerTransition = timerTransition
	return b
}

func (b *_BACnetPropertyStatesTimerTransitionBuilder) WithTimerTransitionBuilder(builderSupplier func(BACnetTimerTransitionTaggedBuilder) BACnetTimerTransitionTaggedBuilder) BACnetPropertyStatesTimerTransitionBuilder {
	builder := builderSupplier(b.TimerTransition.CreateBACnetTimerTransitionTaggedBuilder())
	var err error
	b.TimerTransition, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTimerTransitionTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyStatesTimerTransitionBuilder) Build() (BACnetPropertyStatesTimerTransition, error) {
	if b.TimerTransition == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'timerTransition' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPropertyStatesTimerTransition.deepCopy(), nil
}

func (b *_BACnetPropertyStatesTimerTransitionBuilder) MustBuild() BACnetPropertyStatesTimerTransition {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetPropertyStatesTimerTransitionBuilder) Done() BACnetPropertyStatesBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetPropertyStatesBuilder().(*_BACnetPropertyStatesBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetPropertyStatesTimerTransitionBuilder) buildForBACnetPropertyStates() (BACnetPropertyStates, error) {
	return b.Build()
}

func (b *_BACnetPropertyStatesTimerTransitionBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPropertyStatesTimerTransitionBuilder().(*_BACnetPropertyStatesTimerTransitionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPropertyStatesTimerTransitionBuilder creates a BACnetPropertyStatesTimerTransitionBuilder
func (b *_BACnetPropertyStatesTimerTransition) CreateBACnetPropertyStatesTimerTransitionBuilder() BACnetPropertyStatesTimerTransitionBuilder {
	if b == nil {
		return NewBACnetPropertyStatesTimerTransitionBuilder()
	}
	return &_BACnetPropertyStatesTimerTransitionBuilder{_BACnetPropertyStatesTimerTransition: b.deepCopy()}
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

func (m *_BACnetPropertyStatesTimerTransition) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesTimerTransition) GetTimerTransition() BACnetTimerTransitionTagged {
	return m.TimerTransition
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesTimerTransition(structType any) BACnetPropertyStatesTimerTransition {
	if casted, ok := structType.(BACnetPropertyStatesTimerTransition); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesTimerTransition); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesTimerTransition) GetTypeName() string {
	return "BACnetPropertyStatesTimerTransition"
}

func (m *_BACnetPropertyStatesTimerTransition) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (timerTransition)
	lengthInBits += m.TimerTransition.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesTimerTransition) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesTimerTransition) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesTimerTransition BACnetPropertyStatesTimerTransition, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesTimerTransition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesTimerTransition")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timerTransition, err := ReadSimpleField[BACnetTimerTransitionTagged](ctx, "timerTransition", ReadComplex[BACnetTimerTransitionTagged](BACnetTimerTransitionTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timerTransition' field"))
	}
	m.TimerTransition = timerTransition

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesTimerTransition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesTimerTransition")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesTimerTransition) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesTimerTransition) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesTimerTransition"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesTimerTransition")
		}

		if err := WriteSimpleField[BACnetTimerTransitionTagged](ctx, "timerTransition", m.GetTimerTransition(), WriteComplex[BACnetTimerTransitionTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timerTransition' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesTimerTransition"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesTimerTransition")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesTimerTransition) IsBACnetPropertyStatesTimerTransition() {}

func (m *_BACnetPropertyStatesTimerTransition) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesTimerTransition) deepCopy() *_BACnetPropertyStatesTimerTransition {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesTimerTransitionCopy := &_BACnetPropertyStatesTimerTransition{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		utils.DeepCopy[BACnetTimerTransitionTagged](m.TimerTransition),
	}
	m.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesTimerTransitionCopy
}

func (m *_BACnetPropertyStatesTimerTransition) String() string {
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
