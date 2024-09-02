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

// BACnetSecurityKeySet is the corresponding interface of BACnetSecurityKeySet
type BACnetSecurityKeySet interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetKeyRevision returns KeyRevision (property field)
	GetKeyRevision() BACnetContextTagUnsignedInteger
	// GetActivationTime returns ActivationTime (property field)
	GetActivationTime() BACnetDateTimeEnclosed
	// GetExpirationTime returns ExpirationTime (property field)
	GetExpirationTime() BACnetDateTimeEnclosed
	// GetKeyIds returns KeyIds (property field)
	GetKeyIds() BACnetSecurityKeySetKeyIds
}

// _BACnetSecurityKeySet is the data-structure of this message
type _BACnetSecurityKeySet struct {
	KeyRevision    BACnetContextTagUnsignedInteger
	ActivationTime BACnetDateTimeEnclosed
	ExpirationTime BACnetDateTimeEnclosed
	KeyIds         BACnetSecurityKeySetKeyIds
}

var _ BACnetSecurityKeySet = (*_BACnetSecurityKeySet)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetSecurityKeySet) GetKeyRevision() BACnetContextTagUnsignedInteger {
	return m.KeyRevision
}

func (m *_BACnetSecurityKeySet) GetActivationTime() BACnetDateTimeEnclosed {
	return m.ActivationTime
}

func (m *_BACnetSecurityKeySet) GetExpirationTime() BACnetDateTimeEnclosed {
	return m.ExpirationTime
}

func (m *_BACnetSecurityKeySet) GetKeyIds() BACnetSecurityKeySetKeyIds {
	return m.KeyIds
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetSecurityKeySet factory function for _BACnetSecurityKeySet
func NewBACnetSecurityKeySet(keyRevision BACnetContextTagUnsignedInteger, activationTime BACnetDateTimeEnclosed, expirationTime BACnetDateTimeEnclosed, keyIds BACnetSecurityKeySetKeyIds) *_BACnetSecurityKeySet {
	return &_BACnetSecurityKeySet{KeyRevision: keyRevision, ActivationTime: activationTime, ExpirationTime: expirationTime, KeyIds: keyIds}
}

// Deprecated: use the interface for direct cast
func CastBACnetSecurityKeySet(structType any) BACnetSecurityKeySet {
	if casted, ok := structType.(BACnetSecurityKeySet); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetSecurityKeySet); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetSecurityKeySet) GetTypeName() string {
	return "BACnetSecurityKeySet"
}

func (m *_BACnetSecurityKeySet) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (keyRevision)
	lengthInBits += m.KeyRevision.GetLengthInBits(ctx)

	// Simple field (activationTime)
	lengthInBits += m.ActivationTime.GetLengthInBits(ctx)

	// Simple field (expirationTime)
	lengthInBits += m.ExpirationTime.GetLengthInBits(ctx)

	// Simple field (keyIds)
	lengthInBits += m.KeyIds.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetSecurityKeySet) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetSecurityKeySetParse(ctx context.Context, theBytes []byte) (BACnetSecurityKeySet, error) {
	return BACnetSecurityKeySetParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetSecurityKeySetParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetSecurityKeySet, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetSecurityKeySet, error) {
		return BACnetSecurityKeySetParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetSecurityKeySetParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetSecurityKeySet, error) {
	v, err := (&_BACnetSecurityKeySet{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetSecurityKeySet) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetSecurityKeySet BACnetSecurityKeySet, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetSecurityKeySet"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetSecurityKeySet")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	keyRevision, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "keyRevision", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'keyRevision' field"))
	}
	m.KeyRevision = keyRevision

	activationTime, err := ReadSimpleField[BACnetDateTimeEnclosed](ctx, "activationTime", ReadComplex[BACnetDateTimeEnclosed](BACnetDateTimeEnclosedParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'activationTime' field"))
	}
	m.ActivationTime = activationTime

	expirationTime, err := ReadSimpleField[BACnetDateTimeEnclosed](ctx, "expirationTime", ReadComplex[BACnetDateTimeEnclosed](BACnetDateTimeEnclosedParseWithBufferProducer((uint8)(uint8(2))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'expirationTime' field"))
	}
	m.ExpirationTime = expirationTime

	keyIds, err := ReadSimpleField[BACnetSecurityKeySetKeyIds](ctx, "keyIds", ReadComplex[BACnetSecurityKeySetKeyIds](BACnetSecurityKeySetKeyIdsParseWithBufferProducer((uint8)(uint8(3))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'keyIds' field"))
	}
	m.KeyIds = keyIds

	if closeErr := readBuffer.CloseContext("BACnetSecurityKeySet"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetSecurityKeySet")
	}

	return m, nil
}

func (m *_BACnetSecurityKeySet) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetSecurityKeySet) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetSecurityKeySet"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetSecurityKeySet")
	}

	if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "keyRevision", m.GetKeyRevision(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'keyRevision' field")
	}

	if err := WriteSimpleField[BACnetDateTimeEnclosed](ctx, "activationTime", m.GetActivationTime(), WriteComplex[BACnetDateTimeEnclosed](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'activationTime' field")
	}

	if err := WriteSimpleField[BACnetDateTimeEnclosed](ctx, "expirationTime", m.GetExpirationTime(), WriteComplex[BACnetDateTimeEnclosed](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'expirationTime' field")
	}

	if err := WriteSimpleField[BACnetSecurityKeySetKeyIds](ctx, "keyIds", m.GetKeyIds(), WriteComplex[BACnetSecurityKeySetKeyIds](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'keyIds' field")
	}

	if popErr := writeBuffer.PopContext("BACnetSecurityKeySet"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetSecurityKeySet")
	}
	return nil
}

func (m *_BACnetSecurityKeySet) IsBACnetSecurityKeySet() {}

func (m *_BACnetSecurityKeySet) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
