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
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetPropertyStatesAction is the corresponding interface of BACnetPropertyStatesAction
type BACnetPropertyStatesAction interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetAction returns Action (property field)
	GetAction() BACnetActionTagged
}

// _BACnetPropertyStatesAction is the data-structure of this message
type _BACnetPropertyStatesAction struct {
	BACnetPropertyStatesContract
	Action BACnetActionTagged
}

var _ BACnetPropertyStatesAction = (*_BACnetPropertyStatesAction)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesAction)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesAction) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesAction) GetAction() BACnetActionTagged {
	return m.Action
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesAction factory function for _BACnetPropertyStatesAction
func NewBACnetPropertyStatesAction(action BACnetActionTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesAction {
	_result := &_BACnetPropertyStatesAction{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		Action:                       action,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesAction(structType any) BACnetPropertyStatesAction {
	if casted, ok := structType.(BACnetPropertyStatesAction); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesAction); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesAction) GetTypeName() string {
	return "BACnetPropertyStatesAction"
}

func (m *_BACnetPropertyStatesAction) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (action)
	lengthInBits += m.Action.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesAction) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesAction) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesAction BACnetPropertyStatesAction, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesAction"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesAction")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	action, err := ReadSimpleField[BACnetActionTagged](ctx, "action", ReadComplex[BACnetActionTagged](BACnetActionTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'action' field"))
	}
	m.Action = action

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesAction"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesAction")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesAction) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesAction) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesAction"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesAction")
		}

		if err := WriteSimpleField[BACnetActionTagged](ctx, "action", m.GetAction(), WriteComplex[BACnetActionTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'action' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesAction"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesAction")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesAction) IsBACnetPropertyStatesAction() {}

func (m *_BACnetPropertyStatesAction) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
