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

// NLMUpdateKeyDistributionKey is the corresponding interface of NLMUpdateKeyDistributionKey
type NLMUpdateKeyDistributionKey interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NLM
	// GetKeyRevision returns KeyRevision (property field)
	GetKeyRevision() byte
	// GetKey returns Key (property field)
	GetKey() NLMUpdateKeyUpdateKeyEntry
}

// NLMUpdateKeyDistributionKeyExactly can be used when we want exactly this type and not a type which fulfills NLMUpdateKeyDistributionKey.
// This is useful for switch cases.
type NLMUpdateKeyDistributionKeyExactly interface {
	NLMUpdateKeyDistributionKey
	isNLMUpdateKeyDistributionKey() bool
}

// _NLMUpdateKeyDistributionKey is the data-structure of this message
type _NLMUpdateKeyDistributionKey struct {
	NLMContract
	KeyRevision byte
	Key         NLMUpdateKeyUpdateKeyEntry
}

var _ NLMUpdateKeyDistributionKey = (*_NLMUpdateKeyDistributionKey)(nil)
var _ NLMRequirements = (*_NLMUpdateKeyDistributionKey)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMUpdateKeyDistributionKey) GetMessageType() uint8 {
	return 0x0F
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMUpdateKeyDistributionKey) GetParent() NLMContract {
	return m.NLMContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMUpdateKeyDistributionKey) GetKeyRevision() byte {
	return m.KeyRevision
}

func (m *_NLMUpdateKeyDistributionKey) GetKey() NLMUpdateKeyUpdateKeyEntry {
	return m.Key
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLMUpdateKeyDistributionKey factory function for _NLMUpdateKeyDistributionKey
func NewNLMUpdateKeyDistributionKey(keyRevision byte, key NLMUpdateKeyUpdateKeyEntry, apduLength uint16) *_NLMUpdateKeyDistributionKey {
	_result := &_NLMUpdateKeyDistributionKey{
		NLMContract: NewNLM(apduLength),
		KeyRevision: keyRevision,
		Key:         key,
	}
	_result.NLMContract.(*_NLM)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNLMUpdateKeyDistributionKey(structType any) NLMUpdateKeyDistributionKey {
	if casted, ok := structType.(NLMUpdateKeyDistributionKey); ok {
		return casted
	}
	if casted, ok := structType.(*NLMUpdateKeyDistributionKey); ok {
		return *casted
	}
	return nil
}

func (m *_NLMUpdateKeyDistributionKey) GetTypeName() string {
	return "NLMUpdateKeyDistributionKey"
}

func (m *_NLMUpdateKeyDistributionKey) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.NLMContract.(*_NLM).getLengthInBits(ctx))

	// Simple field (keyRevision)
	lengthInBits += 8

	// Simple field (key)
	lengthInBits += m.Key.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_NLMUpdateKeyDistributionKey) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NLMUpdateKeyDistributionKey) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_NLM, apduLength uint16) (__nLMUpdateKeyDistributionKey NLMUpdateKeyDistributionKey, err error) {
	m.NLMContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMUpdateKeyDistributionKey"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMUpdateKeyDistributionKey")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	keyRevision, err := ReadSimpleField(ctx, "keyRevision", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'keyRevision' field"))
	}
	m.KeyRevision = keyRevision

	key, err := ReadSimpleField[NLMUpdateKeyUpdateKeyEntry](ctx, "key", ReadComplex[NLMUpdateKeyUpdateKeyEntry](NLMUpdateKeyUpdateKeyEntryParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'key' field"))
	}
	m.Key = key

	if closeErr := readBuffer.CloseContext("NLMUpdateKeyDistributionKey"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMUpdateKeyDistributionKey")
	}

	return m, nil
}

func (m *_NLMUpdateKeyDistributionKey) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMUpdateKeyDistributionKey) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMUpdateKeyDistributionKey"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMUpdateKeyDistributionKey")
		}

		if err := WriteSimpleField[byte](ctx, "keyRevision", m.GetKeyRevision(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'keyRevision' field")
		}

		if err := WriteSimpleField[NLMUpdateKeyUpdateKeyEntry](ctx, "key", m.GetKey(), WriteComplex[NLMUpdateKeyUpdateKeyEntry](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'key' field")
		}

		if popErr := writeBuffer.PopContext("NLMUpdateKeyDistributionKey"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMUpdateKeyDistributionKey")
		}
		return nil
	}
	return m.NLMContract.(*_NLM).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NLMUpdateKeyDistributionKey) isNLMUpdateKeyDistributionKey() bool {
	return true
}

func (m *_NLMUpdateKeyDistributionKey) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
