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

// NLMInitializeRoutingTablePortMapping is the corresponding interface of NLMInitializeRoutingTablePortMapping
type NLMInitializeRoutingTablePortMapping interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetDestinationNetworkAddress returns DestinationNetworkAddress (property field)
	GetDestinationNetworkAddress() uint16
	// GetPortId returns PortId (property field)
	GetPortId() uint8
	// GetPortInfoLength returns PortInfoLength (property field)
	GetPortInfoLength() uint8
	// GetPortInfo returns PortInfo (property field)
	GetPortInfo() []byte
}

// NLMInitializeRoutingTablePortMappingExactly can be used when we want exactly this type and not a type which fulfills NLMInitializeRoutingTablePortMapping.
// This is useful for switch cases.
type NLMInitializeRoutingTablePortMappingExactly interface {
	NLMInitializeRoutingTablePortMapping
	isNLMInitializeRoutingTablePortMapping() bool
}

// _NLMInitializeRoutingTablePortMapping is the data-structure of this message
type _NLMInitializeRoutingTablePortMapping struct {
	DestinationNetworkAddress uint16
	PortId                    uint8
	PortInfoLength            uint8
	PortInfo                  []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMInitializeRoutingTablePortMapping) GetDestinationNetworkAddress() uint16 {
	return m.DestinationNetworkAddress
}

func (m *_NLMInitializeRoutingTablePortMapping) GetPortId() uint8 {
	return m.PortId
}

func (m *_NLMInitializeRoutingTablePortMapping) GetPortInfoLength() uint8 {
	return m.PortInfoLength
}

func (m *_NLMInitializeRoutingTablePortMapping) GetPortInfo() []byte {
	return m.PortInfo
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLMInitializeRoutingTablePortMapping factory function for _NLMInitializeRoutingTablePortMapping
func NewNLMInitializeRoutingTablePortMapping(destinationNetworkAddress uint16, portId uint8, portInfoLength uint8, portInfo []byte) *_NLMInitializeRoutingTablePortMapping {
	return &_NLMInitializeRoutingTablePortMapping{DestinationNetworkAddress: destinationNetworkAddress, PortId: portId, PortInfoLength: portInfoLength, PortInfo: portInfo}
}

// Deprecated: use the interface for direct cast
func CastNLMInitializeRoutingTablePortMapping(structType any) NLMInitializeRoutingTablePortMapping {
	if casted, ok := structType.(NLMInitializeRoutingTablePortMapping); ok {
		return casted
	}
	if casted, ok := structType.(*NLMInitializeRoutingTablePortMapping); ok {
		return *casted
	}
	return nil
}

func (m *_NLMInitializeRoutingTablePortMapping) GetTypeName() string {
	return "NLMInitializeRoutingTablePortMapping"
}

func (m *_NLMInitializeRoutingTablePortMapping) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (destinationNetworkAddress)
	lengthInBits += 16

	// Simple field (portId)
	lengthInBits += 8

	// Simple field (portInfoLength)
	lengthInBits += 8

	// Array field
	if len(m.PortInfo) > 0 {
		lengthInBits += 8 * uint16(len(m.PortInfo))
	}

	return lengthInBits
}

func (m *_NLMInitializeRoutingTablePortMapping) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NLMInitializeRoutingTablePortMappingParse(ctx context.Context, theBytes []byte) (NLMInitializeRoutingTablePortMapping, error) {
	return NLMInitializeRoutingTablePortMappingParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func NLMInitializeRoutingTablePortMappingParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (NLMInitializeRoutingTablePortMapping, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("NLMInitializeRoutingTablePortMapping"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMInitializeRoutingTablePortMapping")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (destinationNetworkAddress)
	_destinationNetworkAddress, _destinationNetworkAddressErr := readBuffer.ReadUint16("destinationNetworkAddress", 16)
	if _destinationNetworkAddressErr != nil {
		return nil, errors.Wrap(_destinationNetworkAddressErr, "Error parsing 'destinationNetworkAddress' field of NLMInitializeRoutingTablePortMapping")
	}
	destinationNetworkAddress := _destinationNetworkAddress

	// Simple Field (portId)
	_portId, _portIdErr := readBuffer.ReadUint8("portId", 8)
	if _portIdErr != nil {
		return nil, errors.Wrap(_portIdErr, "Error parsing 'portId' field of NLMInitializeRoutingTablePortMapping")
	}
	portId := _portId

	// Simple Field (portInfoLength)
	_portInfoLength, _portInfoLengthErr := readBuffer.ReadUint8("portInfoLength", 8)
	if _portInfoLengthErr != nil {
		return nil, errors.Wrap(_portInfoLengthErr, "Error parsing 'portInfoLength' field of NLMInitializeRoutingTablePortMapping")
	}
	portInfoLength := _portInfoLength
	// Byte Array field (portInfo)
	numberOfBytesportInfo := int(portInfoLength)
	portInfo, _readArrayErr := readBuffer.ReadByteArray("portInfo", numberOfBytesportInfo)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'portInfo' field of NLMInitializeRoutingTablePortMapping")
	}

	if closeErr := readBuffer.CloseContext("NLMInitializeRoutingTablePortMapping"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMInitializeRoutingTablePortMapping")
	}

	// Create the instance
	return &_NLMInitializeRoutingTablePortMapping{
		DestinationNetworkAddress: destinationNetworkAddress,
		PortId:                    portId,
		PortInfoLength:            portInfoLength,
		PortInfo:                  portInfo,
	}, nil
}

func (m *_NLMInitializeRoutingTablePortMapping) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMInitializeRoutingTablePortMapping) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("NLMInitializeRoutingTablePortMapping"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for NLMInitializeRoutingTablePortMapping")
	}

	// Simple Field (destinationNetworkAddress)
	destinationNetworkAddress := uint16(m.GetDestinationNetworkAddress())
	_destinationNetworkAddressErr := writeBuffer.WriteUint16("destinationNetworkAddress", 16, uint16((destinationNetworkAddress)))
	if _destinationNetworkAddressErr != nil {
		return errors.Wrap(_destinationNetworkAddressErr, "Error serializing 'destinationNetworkAddress' field")
	}

	// Simple Field (portId)
	portId := uint8(m.GetPortId())
	_portIdErr := writeBuffer.WriteUint8("portId", 8, uint8((portId)))
	if _portIdErr != nil {
		return errors.Wrap(_portIdErr, "Error serializing 'portId' field")
	}

	// Simple Field (portInfoLength)
	portInfoLength := uint8(m.GetPortInfoLength())
	_portInfoLengthErr := writeBuffer.WriteUint8("portInfoLength", 8, uint8((portInfoLength)))
	if _portInfoLengthErr != nil {
		return errors.Wrap(_portInfoLengthErr, "Error serializing 'portInfoLength' field")
	}

	// Array Field (portInfo)
	// Byte Array field (portInfo)
	if err := writeBuffer.WriteByteArray("portInfo", m.GetPortInfo()); err != nil {
		return errors.Wrap(err, "Error serializing 'portInfo' field")
	}

	if popErr := writeBuffer.PopContext("NLMInitializeRoutingTablePortMapping"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for NLMInitializeRoutingTablePortMapping")
	}
	return nil
}

func (m *_NLMInitializeRoutingTablePortMapping) isNLMInitializeRoutingTablePortMapping() bool {
	return true
}

func (m *_NLMInitializeRoutingTablePortMapping) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
