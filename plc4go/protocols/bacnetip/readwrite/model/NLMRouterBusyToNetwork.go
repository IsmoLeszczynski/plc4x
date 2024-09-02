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

// NLMRouterBusyToNetwork is the corresponding interface of NLMRouterBusyToNetwork
type NLMRouterBusyToNetwork interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NLM
	// GetDestinationNetworkAddresses returns DestinationNetworkAddresses (property field)
	GetDestinationNetworkAddresses() []uint16
}

// _NLMRouterBusyToNetwork is the data-structure of this message
type _NLMRouterBusyToNetwork struct {
	NLMContract
	DestinationNetworkAddresses []uint16
}

var _ NLMRouterBusyToNetwork = (*_NLMRouterBusyToNetwork)(nil)
var _ NLMRequirements = (*_NLMRouterBusyToNetwork)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMRouterBusyToNetwork) GetMessageType() uint8 {
	return 0x04
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMRouterBusyToNetwork) GetParent() NLMContract {
	return m.NLMContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMRouterBusyToNetwork) GetDestinationNetworkAddresses() []uint16 {
	return m.DestinationNetworkAddresses
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLMRouterBusyToNetwork factory function for _NLMRouterBusyToNetwork
func NewNLMRouterBusyToNetwork(destinationNetworkAddresses []uint16, apduLength uint16) *_NLMRouterBusyToNetwork {
	_result := &_NLMRouterBusyToNetwork{
		NLMContract:                 NewNLM(apduLength),
		DestinationNetworkAddresses: destinationNetworkAddresses,
	}
	_result.NLMContract.(*_NLM)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNLMRouterBusyToNetwork(structType any) NLMRouterBusyToNetwork {
	if casted, ok := structType.(NLMRouterBusyToNetwork); ok {
		return casted
	}
	if casted, ok := structType.(*NLMRouterBusyToNetwork); ok {
		return *casted
	}
	return nil
}

func (m *_NLMRouterBusyToNetwork) GetTypeName() string {
	return "NLMRouterBusyToNetwork"
}

func (m *_NLMRouterBusyToNetwork) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.NLMContract.(*_NLM).getLengthInBits(ctx))

	// Array field
	if len(m.DestinationNetworkAddresses) > 0 {
		lengthInBits += 16 * uint16(len(m.DestinationNetworkAddresses))
	}

	return lengthInBits
}

func (m *_NLMRouterBusyToNetwork) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NLMRouterBusyToNetwork) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_NLM, apduLength uint16) (__nLMRouterBusyToNetwork NLMRouterBusyToNetwork, err error) {
	m.NLMContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMRouterBusyToNetwork"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMRouterBusyToNetwork")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	destinationNetworkAddresses, err := ReadLengthArrayField[uint16](ctx, "destinationNetworkAddresses", ReadUnsignedShort(readBuffer, uint8(16)), int(int32(apduLength)-int32(int32(1))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'destinationNetworkAddresses' field"))
	}
	m.DestinationNetworkAddresses = destinationNetworkAddresses

	if closeErr := readBuffer.CloseContext("NLMRouterBusyToNetwork"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMRouterBusyToNetwork")
	}

	return m, nil
}

func (m *_NLMRouterBusyToNetwork) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMRouterBusyToNetwork) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMRouterBusyToNetwork"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMRouterBusyToNetwork")
		}

		if err := WriteSimpleTypeArrayField(ctx, "destinationNetworkAddresses", m.GetDestinationNetworkAddresses(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'destinationNetworkAddresses' field")
		}

		if popErr := writeBuffer.PopContext("NLMRouterBusyToNetwork"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMRouterBusyToNetwork")
		}
		return nil
	}
	return m.NLMContract.(*_NLM).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NLMRouterBusyToNetwork) IsNLMRouterBusyToNetwork() {}

func (m *_NLMRouterBusyToNetwork) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
