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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BVLCReadBroadcastDistributionTableAck is the corresponding interface of BVLCReadBroadcastDistributionTableAck
type BVLCReadBroadcastDistributionTableAck interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BVLC
	// GetTable returns Table (property field)
	GetTable() []BVLCBroadcastDistributionTableEntry
}

// BVLCReadBroadcastDistributionTableAckExactly can be used when we want exactly this type and not a type which fulfills BVLCReadBroadcastDistributionTableAck.
// This is useful for switch cases.
type BVLCReadBroadcastDistributionTableAckExactly interface {
	BVLCReadBroadcastDistributionTableAck
	isBVLCReadBroadcastDistributionTableAck() bool
}

// _BVLCReadBroadcastDistributionTableAck is the data-structure of this message
type _BVLCReadBroadcastDistributionTableAck struct {
	BVLCContract
	Table []BVLCBroadcastDistributionTableEntry

	// Arguments.
	BvlcPayloadLength uint16
}

var _ BVLCReadBroadcastDistributionTableAck = (*_BVLCReadBroadcastDistributionTableAck)(nil)
var _ BVLCRequirements = (*_BVLCReadBroadcastDistributionTableAck)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BVLCReadBroadcastDistributionTableAck) GetBvlcFunction() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BVLCReadBroadcastDistributionTableAck) GetParent() BVLCContract {
	return m.BVLCContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BVLCReadBroadcastDistributionTableAck) GetTable() []BVLCBroadcastDistributionTableEntry {
	return m.Table
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBVLCReadBroadcastDistributionTableAck factory function for _BVLCReadBroadcastDistributionTableAck
func NewBVLCReadBroadcastDistributionTableAck(table []BVLCBroadcastDistributionTableEntry, bvlcPayloadLength uint16) *_BVLCReadBroadcastDistributionTableAck {
	_result := &_BVLCReadBroadcastDistributionTableAck{
		BVLCContract: NewBVLC(),
		Table:        table,
	}
	_result.BVLCContract.(*_BVLC)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBVLCReadBroadcastDistributionTableAck(structType any) BVLCReadBroadcastDistributionTableAck {
	if casted, ok := structType.(BVLCReadBroadcastDistributionTableAck); ok {
		return casted
	}
	if casted, ok := structType.(*BVLCReadBroadcastDistributionTableAck); ok {
		return *casted
	}
	return nil
}

func (m *_BVLCReadBroadcastDistributionTableAck) GetTypeName() string {
	return "BVLCReadBroadcastDistributionTableAck"
}

func (m *_BVLCReadBroadcastDistributionTableAck) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BVLCContract.(*_BVLC).getLengthInBits(ctx))

	// Array field
	if len(m.Table) > 0 {
		for _, element := range m.Table {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BVLCReadBroadcastDistributionTableAck) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BVLCReadBroadcastDistributionTableAck) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BVLC, bvlcPayloadLength uint16) (__bVLCReadBroadcastDistributionTableAck BVLCReadBroadcastDistributionTableAck, err error) {
	m.BVLCContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BVLCReadBroadcastDistributionTableAck"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BVLCReadBroadcastDistributionTableAck")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	table, err := ReadLengthArrayField[BVLCBroadcastDistributionTableEntry](ctx, "table", ReadComplex[BVLCBroadcastDistributionTableEntry](BVLCBroadcastDistributionTableEntryParseWithBuffer, readBuffer), int(bvlcPayloadLength), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'table' field"))
	}
	m.Table = table

	if closeErr := readBuffer.CloseContext("BVLCReadBroadcastDistributionTableAck"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BVLCReadBroadcastDistributionTableAck")
	}

	return m, nil
}

func (m *_BVLCReadBroadcastDistributionTableAck) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BVLCReadBroadcastDistributionTableAck) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCReadBroadcastDistributionTableAck"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BVLCReadBroadcastDistributionTableAck")
		}

		if err := WriteComplexTypeArrayField(ctx, "table", m.GetTable(), writeBuffer, codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'table' field")
		}

		if popErr := writeBuffer.PopContext("BVLCReadBroadcastDistributionTableAck"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BVLCReadBroadcastDistributionTableAck")
		}
		return nil
	}
	return m.BVLCContract.(*_BVLC).serializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_BVLCReadBroadcastDistributionTableAck) GetBvlcPayloadLength() uint16 {
	return m.BvlcPayloadLength
}

//
////

func (m *_BVLCReadBroadcastDistributionTableAck) isBVLCReadBroadcastDistributionTableAck() bool {
	return true
}

func (m *_BVLCReadBroadcastDistributionTableAck) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
