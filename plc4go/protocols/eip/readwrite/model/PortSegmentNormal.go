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

// PortSegmentNormal is the corresponding interface of PortSegmentNormal
type PortSegmentNormal interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	PortSegmentType
	// GetPort returns Port (property field)
	GetPort() uint8
	// GetLinkAddress returns LinkAddress (property field)
	GetLinkAddress() uint8
}

// _PortSegmentNormal is the data-structure of this message
type _PortSegmentNormal struct {
	PortSegmentTypeContract
	Port        uint8
	LinkAddress uint8
}

var _ PortSegmentNormal = (*_PortSegmentNormal)(nil)
var _ PortSegmentTypeRequirements = (*_PortSegmentNormal)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_PortSegmentNormal) GetExtendedLinkAddress() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PortSegmentNormal) GetParent() PortSegmentTypeContract {
	return m.PortSegmentTypeContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PortSegmentNormal) GetPort() uint8 {
	return m.Port
}

func (m *_PortSegmentNormal) GetLinkAddress() uint8 {
	return m.LinkAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewPortSegmentNormal factory function for _PortSegmentNormal
func NewPortSegmentNormal(port uint8, linkAddress uint8) *_PortSegmentNormal {
	_result := &_PortSegmentNormal{
		PortSegmentTypeContract: NewPortSegmentType(),
		Port:                    port,
		LinkAddress:             linkAddress,
	}
	_result.PortSegmentTypeContract.(*_PortSegmentType)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastPortSegmentNormal(structType any) PortSegmentNormal {
	if casted, ok := structType.(PortSegmentNormal); ok {
		return casted
	}
	if casted, ok := structType.(*PortSegmentNormal); ok {
		return *casted
	}
	return nil
}

func (m *_PortSegmentNormal) GetTypeName() string {
	return "PortSegmentNormal"
}

func (m *_PortSegmentNormal) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.PortSegmentTypeContract.(*_PortSegmentType).getLengthInBits(ctx))

	// Simple field (port)
	lengthInBits += 4

	// Simple field (linkAddress)
	lengthInBits += 8

	return lengthInBits
}

func (m *_PortSegmentNormal) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_PortSegmentNormal) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_PortSegmentType) (__portSegmentNormal PortSegmentNormal, err error) {
	m.PortSegmentTypeContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("PortSegmentNormal"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PortSegmentNormal")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	port, err := ReadSimpleField(ctx, "port", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'port' field"))
	}
	m.Port = port

	linkAddress, err := ReadSimpleField(ctx, "linkAddress", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'linkAddress' field"))
	}
	m.LinkAddress = linkAddress

	if closeErr := readBuffer.CloseContext("PortSegmentNormal"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PortSegmentNormal")
	}

	return m, nil
}

func (m *_PortSegmentNormal) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PortSegmentNormal) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PortSegmentNormal"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PortSegmentNormal")
		}

		if err := WriteSimpleField[uint8](ctx, "port", m.GetPort(), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'port' field")
		}

		if err := WriteSimpleField[uint8](ctx, "linkAddress", m.GetLinkAddress(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'linkAddress' field")
		}

		if popErr := writeBuffer.PopContext("PortSegmentNormal"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PortSegmentNormal")
		}
		return nil
	}
	return m.PortSegmentTypeContract.(*_PortSegmentType).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_PortSegmentNormal) IsPortSegmentNormal() {}

func (m *_PortSegmentNormal) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
