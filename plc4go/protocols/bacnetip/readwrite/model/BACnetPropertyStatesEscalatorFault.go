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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetPropertyStatesEscalatorFault is the corresponding interface of BACnetPropertyStatesEscalatorFault
type BACnetPropertyStatesEscalatorFault interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetEscalatorFault returns EscalatorFault (property field)
	GetEscalatorFault() BACnetEscalatorFaultTagged
}

// BACnetPropertyStatesEscalatorFaultExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesEscalatorFault.
// This is useful for switch cases.
type BACnetPropertyStatesEscalatorFaultExactly interface {
	BACnetPropertyStatesEscalatorFault
	isBACnetPropertyStatesEscalatorFault() bool
}

// _BACnetPropertyStatesEscalatorFault is the data-structure of this message
type _BACnetPropertyStatesEscalatorFault struct {
	*_BACnetPropertyStates
	EscalatorFault BACnetEscalatorFaultTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesEscalatorFault) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesEscalatorFault) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesEscalatorFault) GetEscalatorFault() BACnetEscalatorFaultTagged {
	return m.EscalatorFault
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesEscalatorFault factory function for _BACnetPropertyStatesEscalatorFault
func NewBACnetPropertyStatesEscalatorFault(escalatorFault BACnetEscalatorFaultTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesEscalatorFault {
	_result := &_BACnetPropertyStatesEscalatorFault{
		EscalatorFault:        escalatorFault,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesEscalatorFault(structType any) BACnetPropertyStatesEscalatorFault {
	if casted, ok := structType.(BACnetPropertyStatesEscalatorFault); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesEscalatorFault); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesEscalatorFault) GetTypeName() string {
	return "BACnetPropertyStatesEscalatorFault"
}

func (m *_BACnetPropertyStatesEscalatorFault) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (escalatorFault)
	lengthInBits += m.EscalatorFault.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesEscalatorFault) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyStatesEscalatorFaultParse(ctx context.Context, theBytes []byte, peekedTagNumber uint8) (BACnetPropertyStatesEscalatorFault, error) {
	return BACnetPropertyStatesEscalatorFaultParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), peekedTagNumber)
}

func BACnetPropertyStatesEscalatorFaultParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesEscalatorFault, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesEscalatorFault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesEscalatorFault")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (escalatorFault)
	if pullErr := readBuffer.PullContext("escalatorFault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for escalatorFault")
	}
	_escalatorFault, _escalatorFaultErr := BACnetEscalatorFaultTaggedParseWithBuffer(ctx, readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _escalatorFaultErr != nil {
		return nil, errors.Wrap(_escalatorFaultErr, "Error parsing 'escalatorFault' field of BACnetPropertyStatesEscalatorFault")
	}
	escalatorFault := _escalatorFault.(BACnetEscalatorFaultTagged)
	if closeErr := readBuffer.CloseContext("escalatorFault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for escalatorFault")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesEscalatorFault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesEscalatorFault")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesEscalatorFault{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		EscalatorFault:        escalatorFault,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesEscalatorFault) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesEscalatorFault) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesEscalatorFault"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesEscalatorFault")
		}

		// Simple Field (escalatorFault)
		if pushErr := writeBuffer.PushContext("escalatorFault"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for escalatorFault")
		}
		_escalatorFaultErr := writeBuffer.WriteSerializable(ctx, m.GetEscalatorFault())
		if popErr := writeBuffer.PopContext("escalatorFault"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for escalatorFault")
		}
		if _escalatorFaultErr != nil {
			return errors.Wrap(_escalatorFaultErr, "Error serializing 'escalatorFault' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesEscalatorFault"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesEscalatorFault")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesEscalatorFault) isBACnetPropertyStatesEscalatorFault() bool {
	return true
}

func (m *_BACnetPropertyStatesEscalatorFault) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
