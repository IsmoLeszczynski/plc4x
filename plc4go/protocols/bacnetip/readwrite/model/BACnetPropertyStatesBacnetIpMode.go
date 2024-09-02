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

// BACnetPropertyStatesBacnetIpMode is the corresponding interface of BACnetPropertyStatesBacnetIpMode
type BACnetPropertyStatesBacnetIpMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetBacnetIpMode returns BacnetIpMode (property field)
	GetBacnetIpMode() BACnetIPModeTagged
}

// _BACnetPropertyStatesBacnetIpMode is the data-structure of this message
type _BACnetPropertyStatesBacnetIpMode struct {
	BACnetPropertyStatesContract
	BacnetIpMode BACnetIPModeTagged
}

var _ BACnetPropertyStatesBacnetIpMode = (*_BACnetPropertyStatesBacnetIpMode)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesBacnetIpMode)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesBacnetIpMode) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesBacnetIpMode) GetBacnetIpMode() BACnetIPModeTagged {
	return m.BacnetIpMode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesBacnetIpMode factory function for _BACnetPropertyStatesBacnetIpMode
func NewBACnetPropertyStatesBacnetIpMode(bacnetIpMode BACnetIPModeTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesBacnetIpMode {
	_result := &_BACnetPropertyStatesBacnetIpMode{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		BacnetIpMode:                 bacnetIpMode,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesBacnetIpMode(structType any) BACnetPropertyStatesBacnetIpMode {
	if casted, ok := structType.(BACnetPropertyStatesBacnetIpMode); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesBacnetIpMode); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesBacnetIpMode) GetTypeName() string {
	return "BACnetPropertyStatesBacnetIpMode"
}

func (m *_BACnetPropertyStatesBacnetIpMode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (bacnetIpMode)
	lengthInBits += m.BacnetIpMode.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesBacnetIpMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesBacnetIpMode) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesBacnetIpMode BACnetPropertyStatesBacnetIpMode, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesBacnetIpMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesBacnetIpMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	bacnetIpMode, err := ReadSimpleField[BACnetIPModeTagged](ctx, "bacnetIpMode", ReadComplex[BACnetIPModeTagged](BACnetIPModeTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bacnetIpMode' field"))
	}
	m.BacnetIpMode = bacnetIpMode

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesBacnetIpMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesBacnetIpMode")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesBacnetIpMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesBacnetIpMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesBacnetIpMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesBacnetIpMode")
		}

		if err := WriteSimpleField[BACnetIPModeTagged](ctx, "bacnetIpMode", m.GetBacnetIpMode(), WriteComplex[BACnetIPModeTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'bacnetIpMode' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesBacnetIpMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesBacnetIpMode")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesBacnetIpMode) IsBACnetPropertyStatesBacnetIpMode() {}

func (m *_BACnetPropertyStatesBacnetIpMode) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
