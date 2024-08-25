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

// BACnetPropertyStatesLockStatus is the corresponding interface of BACnetPropertyStatesLockStatus
type BACnetPropertyStatesLockStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetLockStatus returns LockStatus (property field)
	GetLockStatus() BACnetLockStatusTagged
}

// BACnetPropertyStatesLockStatusExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesLockStatus.
// This is useful for switch cases.
type BACnetPropertyStatesLockStatusExactly interface {
	BACnetPropertyStatesLockStatus
	isBACnetPropertyStatesLockStatus() bool
}

// _BACnetPropertyStatesLockStatus is the data-structure of this message
type _BACnetPropertyStatesLockStatus struct {
	*_BACnetPropertyStates
	LockStatus BACnetLockStatusTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesLockStatus) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesLockStatus) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesLockStatus) GetLockStatus() BACnetLockStatusTagged {
	return m.LockStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesLockStatus factory function for _BACnetPropertyStatesLockStatus
func NewBACnetPropertyStatesLockStatus(lockStatus BACnetLockStatusTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesLockStatus {
	_result := &_BACnetPropertyStatesLockStatus{
		LockStatus:            lockStatus,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesLockStatus(structType any) BACnetPropertyStatesLockStatus {
	if casted, ok := structType.(BACnetPropertyStatesLockStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesLockStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesLockStatus) GetTypeName() string {
	return "BACnetPropertyStatesLockStatus"
}

func (m *_BACnetPropertyStatesLockStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (lockStatus)
	lengthInBits += m.LockStatus.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesLockStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyStatesLockStatusParse(ctx context.Context, theBytes []byte, peekedTagNumber uint8) (BACnetPropertyStatesLockStatus, error) {
	return BACnetPropertyStatesLockStatusParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), peekedTagNumber)
}

func BACnetPropertyStatesLockStatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesLockStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLockStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesLockStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lockStatus)
	if pullErr := readBuffer.PullContext("lockStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lockStatus")
	}
	_lockStatus, _lockStatusErr := BACnetLockStatusTaggedParseWithBuffer(ctx, readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _lockStatusErr != nil {
		return nil, errors.Wrap(_lockStatusErr, "Error parsing 'lockStatus' field of BACnetPropertyStatesLockStatus")
	}
	lockStatus := _lockStatus.(BACnetLockStatusTagged)
	if closeErr := readBuffer.CloseContext("lockStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lockStatus")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLockStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesLockStatus")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesLockStatus{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		LockStatus:            lockStatus,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesLockStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesLockStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesLockStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesLockStatus")
		}

		// Simple Field (lockStatus)
		if pushErr := writeBuffer.PushContext("lockStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lockStatus")
		}
		_lockStatusErr := writeBuffer.WriteSerializable(ctx, m.GetLockStatus())
		if popErr := writeBuffer.PopContext("lockStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lockStatus")
		}
		if _lockStatusErr != nil {
			return errors.Wrap(_lockStatusErr, "Error serializing 'lockStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLockStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesLockStatus")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesLockStatus) isBACnetPropertyStatesLockStatus() bool {
	return true
}

func (m *_BACnetPropertyStatesLockStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
