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

// BACnetBDTEntry is the corresponding interface of BACnetBDTEntry
type BACnetBDTEntry interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetBbmdAddress returns BbmdAddress (property field)
	GetBbmdAddress() BACnetHostNPortEnclosed
	// GetBroadcastMask returns BroadcastMask (property field)
	GetBroadcastMask() BACnetContextTagOctetString
}

// _BACnetBDTEntry is the data-structure of this message
type _BACnetBDTEntry struct {
	BbmdAddress   BACnetHostNPortEnclosed
	BroadcastMask BACnetContextTagOctetString
}

var _ BACnetBDTEntry = (*_BACnetBDTEntry)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetBDTEntry) GetBbmdAddress() BACnetHostNPortEnclosed {
	return m.BbmdAddress
}

func (m *_BACnetBDTEntry) GetBroadcastMask() BACnetContextTagOctetString {
	return m.BroadcastMask
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetBDTEntry factory function for _BACnetBDTEntry
func NewBACnetBDTEntry(bbmdAddress BACnetHostNPortEnclosed, broadcastMask BACnetContextTagOctetString) *_BACnetBDTEntry {
	return &_BACnetBDTEntry{BbmdAddress: bbmdAddress, BroadcastMask: broadcastMask}
}

// Deprecated: use the interface for direct cast
func CastBACnetBDTEntry(structType any) BACnetBDTEntry {
	if casted, ok := structType.(BACnetBDTEntry); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetBDTEntry); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetBDTEntry) GetTypeName() string {
	return "BACnetBDTEntry"
}

func (m *_BACnetBDTEntry) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (bbmdAddress)
	lengthInBits += m.BbmdAddress.GetLengthInBits(ctx)

	// Optional Field (broadcastMask)
	if m.BroadcastMask != nil {
		lengthInBits += m.BroadcastMask.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetBDTEntry) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetBDTEntryParse(ctx context.Context, theBytes []byte) (BACnetBDTEntry, error) {
	return BACnetBDTEntryParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetBDTEntryParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetBDTEntry, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetBDTEntry, error) {
		return BACnetBDTEntryParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetBDTEntryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetBDTEntry, error) {
	v, err := (&_BACnetBDTEntry{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetBDTEntry) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetBDTEntry BACnetBDTEntry, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetBDTEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetBDTEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	bbmdAddress, err := ReadSimpleField[BACnetHostNPortEnclosed](ctx, "bbmdAddress", ReadComplex[BACnetHostNPortEnclosed](BACnetHostNPortEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bbmdAddress' field"))
	}
	m.BbmdAddress = bbmdAddress

	var broadcastMask BACnetContextTagOctetString
	_broadcastMask, err := ReadOptionalField[BACnetContextTagOctetString](ctx, "broadcastMask", ReadComplex[BACnetContextTagOctetString](BACnetContextTagParseWithBufferProducer[BACnetContextTagOctetString]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_OCTET_STRING)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'broadcastMask' field"))
	}
	if _broadcastMask != nil {
		broadcastMask = *_broadcastMask
		m.BroadcastMask = broadcastMask
	}

	if closeErr := readBuffer.CloseContext("BACnetBDTEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetBDTEntry")
	}

	return m, nil
}

func (m *_BACnetBDTEntry) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetBDTEntry) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetBDTEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetBDTEntry")
	}

	if err := WriteSimpleField[BACnetHostNPortEnclosed](ctx, "bbmdAddress", m.GetBbmdAddress(), WriteComplex[BACnetHostNPortEnclosed](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'bbmdAddress' field")
	}

	if err := WriteOptionalField[BACnetContextTagOctetString](ctx, "broadcastMask", GetRef(m.GetBroadcastMask()), WriteComplex[BACnetContextTagOctetString](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'broadcastMask' field")
	}

	if popErr := writeBuffer.PopContext("BACnetBDTEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetBDTEntry")
	}
	return nil
}

func (m *_BACnetBDTEntry) IsBACnetBDTEntry() {}

func (m *_BACnetBDTEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
