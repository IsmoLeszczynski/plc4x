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
	BACnetPropertyStatesContract
	LockStatus BACnetLockStatusTagged
}

var _ BACnetPropertyStatesLockStatus = (*_BACnetPropertyStatesLockStatus)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesLockStatus)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesLockStatus) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
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
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		LockStatus:                   lockStatus,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
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
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (lockStatus)
	lengthInBits += m.LockStatus.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesLockStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesLockStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesLockStatus BACnetPropertyStatesLockStatus, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLockStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesLockStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lockStatus, err := ReadSimpleField[BACnetLockStatusTagged](ctx, "lockStatus", ReadComplex[BACnetLockStatusTagged](BACnetLockStatusTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lockStatus' field"))
	}
	m.LockStatus = lockStatus

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLockStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesLockStatus")
	}

	return m, nil
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

		if err := WriteSimpleField[BACnetLockStatusTagged](ctx, "lockStatus", m.GetLockStatus(), WriteComplex[BACnetLockStatusTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lockStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLockStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesLockStatus")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
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
