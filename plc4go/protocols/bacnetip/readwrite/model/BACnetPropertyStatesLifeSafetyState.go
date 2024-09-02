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
	BACnetPropertyStates
	// GetLifeSafetyState returns LifeSafetyState (property field)
	GetLifeSafetyState() BACnetLifeSafetyStateTagged
}

// BACnetPropertyStatesLifeSafetyStateExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesLifeSafetyState.
// This is useful for switch cases.
type BACnetPropertyStatesLifeSafetyStateExactly interface {
	BACnetPropertyStatesLifeSafetyState
	isBACnetPropertyStatesLifeSafetyState() bool
}

// _BACnetPropertyStatesLifeSafetyState is the data-structure of this message
type _BACnetPropertyStatesLifeSafetyState struct {
	BACnetPropertyStatesContract
	LifeSafetyState BACnetLifeSafetyStateTagged
}

var _ BACnetPropertyStatesLifeSafetyState = (*_BACnetPropertyStatesLifeSafetyState)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesLifeSafetyState)(nil)

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

// NewBACnetPropertyStatesLifeSafetyState factory function for _BACnetPropertyStatesLifeSafetyState
func NewBACnetPropertyStatesLifeSafetyState(lifeSafetyState BACnetLifeSafetyStateTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesLifeSafetyState {
	_result := &_BACnetPropertyStatesLifeSafetyState{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		LifeSafetyState:              lifeSafetyState,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

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

func (m *_BACnetPropertyStatesLifeSafetyState) isBACnetPropertyStatesLifeSafetyState() bool {
	return true
}

func (m *_BACnetPropertyStatesLifeSafetyState) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
