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

// BACnetPropertyStatesNetworkNumberQuality is the corresponding interface of BACnetPropertyStatesNetworkNumberQuality
type BACnetPropertyStatesNetworkNumberQuality interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetNetworkNumberQuality returns NetworkNumberQuality (property field)
	GetNetworkNumberQuality() BACnetNetworkNumberQualityTagged
}

// _BACnetPropertyStatesNetworkNumberQuality is the data-structure of this message
type _BACnetPropertyStatesNetworkNumberQuality struct {
	BACnetPropertyStatesContract
	NetworkNumberQuality BACnetNetworkNumberQualityTagged
}

var _ BACnetPropertyStatesNetworkNumberQuality = (*_BACnetPropertyStatesNetworkNumberQuality)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesNetworkNumberQuality)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesNetworkNumberQuality) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesNetworkNumberQuality) GetNetworkNumberQuality() BACnetNetworkNumberQualityTagged {
	return m.NetworkNumberQuality
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesNetworkNumberQuality factory function for _BACnetPropertyStatesNetworkNumberQuality
func NewBACnetPropertyStatesNetworkNumberQuality(networkNumberQuality BACnetNetworkNumberQualityTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesNetworkNumberQuality {
	_result := &_BACnetPropertyStatesNetworkNumberQuality{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		NetworkNumberQuality:         networkNumberQuality,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesNetworkNumberQuality(structType any) BACnetPropertyStatesNetworkNumberQuality {
	if casted, ok := structType.(BACnetPropertyStatesNetworkNumberQuality); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesNetworkNumberQuality); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesNetworkNumberQuality) GetTypeName() string {
	return "BACnetPropertyStatesNetworkNumberQuality"
}

func (m *_BACnetPropertyStatesNetworkNumberQuality) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (networkNumberQuality)
	lengthInBits += m.NetworkNumberQuality.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesNetworkNumberQuality) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesNetworkNumberQuality) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesNetworkNumberQuality BACnetPropertyStatesNetworkNumberQuality, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesNetworkNumberQuality"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesNetworkNumberQuality")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	networkNumberQuality, err := ReadSimpleField[BACnetNetworkNumberQualityTagged](ctx, "networkNumberQuality", ReadComplex[BACnetNetworkNumberQualityTagged](BACnetNetworkNumberQualityTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkNumberQuality' field"))
	}
	m.NetworkNumberQuality = networkNumberQuality

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesNetworkNumberQuality"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesNetworkNumberQuality")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesNetworkNumberQuality) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesNetworkNumberQuality) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesNetworkNumberQuality"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesNetworkNumberQuality")
		}

		if err := WriteSimpleField[BACnetNetworkNumberQualityTagged](ctx, "networkNumberQuality", m.GetNetworkNumberQuality(), WriteComplex[BACnetNetworkNumberQualityTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'networkNumberQuality' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesNetworkNumberQuality"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesNetworkNumberQuality")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesNetworkNumberQuality) IsBACnetPropertyStatesNetworkNumberQuality() {}

func (m *_BACnetPropertyStatesNetworkNumberQuality) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
