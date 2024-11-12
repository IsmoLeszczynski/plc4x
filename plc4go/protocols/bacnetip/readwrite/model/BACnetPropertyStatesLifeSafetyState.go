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

// BACnetPropertyStatesLifeSafetyState is the corresponding interface of BACnetPropertyStatesLifeSafetyState
type BACnetPropertyStatesLifeSafetyState interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetLifeSafetyState returns LifeSafetyState (property field)
	GetLifeSafetyState() BACnetLifeSafetyStateTagged
	// IsBACnetPropertyStatesLifeSafetyState is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesLifeSafetyState()
	// CreateBuilder creates a BACnetPropertyStatesLifeSafetyStateBuilder
	CreateBACnetPropertyStatesLifeSafetyStateBuilder() BACnetPropertyStatesLifeSafetyStateBuilder
}

// _BACnetPropertyStatesLifeSafetyState is the data-structure of this message
type _BACnetPropertyStatesLifeSafetyState struct {
	BACnetPropertyStatesContract
	LifeSafetyState BACnetLifeSafetyStateTagged
}

var _ BACnetPropertyStatesLifeSafetyState = (*_BACnetPropertyStatesLifeSafetyState)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesLifeSafetyState)(nil)

// NewBACnetPropertyStatesLifeSafetyState factory function for _BACnetPropertyStatesLifeSafetyState
func NewBACnetPropertyStatesLifeSafetyState(peekedTagHeader BACnetTagHeader, lifeSafetyState BACnetLifeSafetyStateTagged) *_BACnetPropertyStatesLifeSafetyState {
	if lifeSafetyState == nil {
		panic("lifeSafetyState of type BACnetLifeSafetyStateTagged for BACnetPropertyStatesLifeSafetyState must not be nil")
	}
	_result := &_BACnetPropertyStatesLifeSafetyState{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		LifeSafetyState:              lifeSafetyState,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesLifeSafetyStateBuilder is a builder for BACnetPropertyStatesLifeSafetyState
type BACnetPropertyStatesLifeSafetyStateBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(lifeSafetyState BACnetLifeSafetyStateTagged) BACnetPropertyStatesLifeSafetyStateBuilder
	// WithLifeSafetyState adds LifeSafetyState (property field)
	WithLifeSafetyState(BACnetLifeSafetyStateTagged) BACnetPropertyStatesLifeSafetyStateBuilder
	// WithLifeSafetyStateBuilder adds LifeSafetyState (property field) which is build by the builder
	WithLifeSafetyStateBuilder(func(BACnetLifeSafetyStateTaggedBuilder) BACnetLifeSafetyStateTaggedBuilder) BACnetPropertyStatesLifeSafetyStateBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetPropertyStatesBuilder
	// Build builds the BACnetPropertyStatesLifeSafetyState or returns an error if something is wrong
	Build() (BACnetPropertyStatesLifeSafetyState, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesLifeSafetyState
}

// NewBACnetPropertyStatesLifeSafetyStateBuilder() creates a BACnetPropertyStatesLifeSafetyStateBuilder
func NewBACnetPropertyStatesLifeSafetyStateBuilder() BACnetPropertyStatesLifeSafetyStateBuilder {
	return &_BACnetPropertyStatesLifeSafetyStateBuilder{_BACnetPropertyStatesLifeSafetyState: new(_BACnetPropertyStatesLifeSafetyState)}
}

type _BACnetPropertyStatesLifeSafetyStateBuilder struct {
	*_BACnetPropertyStatesLifeSafetyState

	parentBuilder *_BACnetPropertyStatesBuilder

	err *utils.MultiError
}

var _ (BACnetPropertyStatesLifeSafetyStateBuilder) = (*_BACnetPropertyStatesLifeSafetyStateBuilder)(nil)

func (b *_BACnetPropertyStatesLifeSafetyStateBuilder) setParent(contract BACnetPropertyStatesContract) {
	b.BACnetPropertyStatesContract = contract
}

func (b *_BACnetPropertyStatesLifeSafetyStateBuilder) WithMandatoryFields(lifeSafetyState BACnetLifeSafetyStateTagged) BACnetPropertyStatesLifeSafetyStateBuilder {
	return b.WithLifeSafetyState(lifeSafetyState)
}

func (b *_BACnetPropertyStatesLifeSafetyStateBuilder) WithLifeSafetyState(lifeSafetyState BACnetLifeSafetyStateTagged) BACnetPropertyStatesLifeSafetyStateBuilder {
	b.LifeSafetyState = lifeSafetyState
	return b
}

func (b *_BACnetPropertyStatesLifeSafetyStateBuilder) WithLifeSafetyStateBuilder(builderSupplier func(BACnetLifeSafetyStateTaggedBuilder) BACnetLifeSafetyStateTaggedBuilder) BACnetPropertyStatesLifeSafetyStateBuilder {
	builder := builderSupplier(b.LifeSafetyState.CreateBACnetLifeSafetyStateTaggedBuilder())
	var err error
	b.LifeSafetyState, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetLifeSafetyStateTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyStatesLifeSafetyStateBuilder) Build() (BACnetPropertyStatesLifeSafetyState, error) {
	if b.LifeSafetyState == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'lifeSafetyState' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPropertyStatesLifeSafetyState.deepCopy(), nil
}

func (b *_BACnetPropertyStatesLifeSafetyStateBuilder) MustBuild() BACnetPropertyStatesLifeSafetyState {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetPropertyStatesLifeSafetyStateBuilder) Done() BACnetPropertyStatesBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetPropertyStatesBuilder().(*_BACnetPropertyStatesBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetPropertyStatesLifeSafetyStateBuilder) buildForBACnetPropertyStates() (BACnetPropertyStates, error) {
	return b.Build()
}

func (b *_BACnetPropertyStatesLifeSafetyStateBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPropertyStatesLifeSafetyStateBuilder().(*_BACnetPropertyStatesLifeSafetyStateBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPropertyStatesLifeSafetyStateBuilder creates a BACnetPropertyStatesLifeSafetyStateBuilder
func (b *_BACnetPropertyStatesLifeSafetyState) CreateBACnetPropertyStatesLifeSafetyStateBuilder() BACnetPropertyStatesLifeSafetyStateBuilder {
	if b == nil {
		return NewBACnetPropertyStatesLifeSafetyStateBuilder()
	}
	return &_BACnetPropertyStatesLifeSafetyStateBuilder{_BACnetPropertyStatesLifeSafetyState: b.deepCopy()}
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

func (m *_BACnetPropertyStatesLifeSafetyState) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesLifeSafetyState) GetLifeSafetyState() BACnetLifeSafetyStateTagged {
	return m.LifeSafetyState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesLifeSafetyState(structType any) BACnetPropertyStatesLifeSafetyState {
	if casted, ok := structType.(BACnetPropertyStatesLifeSafetyState); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesLifeSafetyState); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesLifeSafetyState) GetTypeName() string {
	return "BACnetPropertyStatesLifeSafetyState"
}

func (m *_BACnetPropertyStatesLifeSafetyState) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (lifeSafetyState)
	lengthInBits += m.LifeSafetyState.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesLifeSafetyState) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesLifeSafetyState) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesLifeSafetyState BACnetPropertyStatesLifeSafetyState, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLifeSafetyState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesLifeSafetyState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lifeSafetyState, err := ReadSimpleField[BACnetLifeSafetyStateTagged](ctx, "lifeSafetyState", ReadComplex[BACnetLifeSafetyStateTagged](BACnetLifeSafetyStateTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lifeSafetyState' field"))
	}
	m.LifeSafetyState = lifeSafetyState

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLifeSafetyState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesLifeSafetyState")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesLifeSafetyState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesLifeSafetyState) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesLifeSafetyState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesLifeSafetyState")
		}

		if err := WriteSimpleField[BACnetLifeSafetyStateTagged](ctx, "lifeSafetyState", m.GetLifeSafetyState(), WriteComplex[BACnetLifeSafetyStateTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lifeSafetyState' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLifeSafetyState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesLifeSafetyState")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesLifeSafetyState) IsBACnetPropertyStatesLifeSafetyState() {}

func (m *_BACnetPropertyStatesLifeSafetyState) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesLifeSafetyState) deepCopy() *_BACnetPropertyStatesLifeSafetyState {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesLifeSafetyStateCopy := &_BACnetPropertyStatesLifeSafetyState{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		utils.DeepCopy[BACnetLifeSafetyStateTagged](m.LifeSafetyState),
	}
	m.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesLifeSafetyStateCopy
}

func (m *_BACnetPropertyStatesLifeSafetyState) String() string {
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
