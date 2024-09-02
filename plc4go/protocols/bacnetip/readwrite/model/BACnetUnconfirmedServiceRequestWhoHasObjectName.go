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

// BACnetUnconfirmedServiceRequestWhoHasObjectName is the corresponding interface of BACnetUnconfirmedServiceRequestWhoHasObjectName
type BACnetUnconfirmedServiceRequestWhoHasObjectName interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetUnconfirmedServiceRequestWhoHasObject
	// GetObjectName returns ObjectName (property field)
	GetObjectName() BACnetContextTagCharacterString
}

// _BACnetUnconfirmedServiceRequestWhoHasObjectName is the data-structure of this message
type _BACnetUnconfirmedServiceRequestWhoHasObjectName struct {
	BACnetUnconfirmedServiceRequestWhoHasObjectContract
	ObjectName BACnetContextTagCharacterString
}

var _ BACnetUnconfirmedServiceRequestWhoHasObjectName = (*_BACnetUnconfirmedServiceRequestWhoHasObjectName)(nil)
var _ BACnetUnconfirmedServiceRequestWhoHasObjectRequirements = (*_BACnetUnconfirmedServiceRequestWhoHasObjectName)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectName) GetParent() BACnetUnconfirmedServiceRequestWhoHasObjectContract {
	return m.BACnetUnconfirmedServiceRequestWhoHasObjectContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectName) GetObjectName() BACnetContextTagCharacterString {
	return m.ObjectName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetUnconfirmedServiceRequestWhoHasObjectName factory function for _BACnetUnconfirmedServiceRequestWhoHasObjectName
func NewBACnetUnconfirmedServiceRequestWhoHasObjectName(objectName BACnetContextTagCharacterString, peekedTagHeader BACnetTagHeader) *_BACnetUnconfirmedServiceRequestWhoHasObjectName {
	_result := &_BACnetUnconfirmedServiceRequestWhoHasObjectName{
		BACnetUnconfirmedServiceRequestWhoHasObjectContract: NewBACnetUnconfirmedServiceRequestWhoHasObject(peekedTagHeader),
		ObjectName: objectName,
	}
	_result.BACnetUnconfirmedServiceRequestWhoHasObjectContract.(*_BACnetUnconfirmedServiceRequestWhoHasObject)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetUnconfirmedServiceRequestWhoHasObjectName(structType any) BACnetUnconfirmedServiceRequestWhoHasObjectName {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestWhoHasObjectName); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestWhoHasObjectName); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectName) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestWhoHasObjectName"
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectName) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetUnconfirmedServiceRequestWhoHasObjectContract.(*_BACnetUnconfirmedServiceRequestWhoHasObject).getLengthInBits(ctx))

	// Simple field (objectName)
	lengthInBits += m.ObjectName.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectName) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectName) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetUnconfirmedServiceRequestWhoHasObject) (__bACnetUnconfirmedServiceRequestWhoHasObjectName BACnetUnconfirmedServiceRequestWhoHasObjectName, err error) {
	m.BACnetUnconfirmedServiceRequestWhoHasObjectContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestWhoHasObjectName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetUnconfirmedServiceRequestWhoHasObjectName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectName, err := ReadSimpleField[BACnetContextTagCharacterString](ctx, "objectName", ReadComplex[BACnetContextTagCharacterString](BACnetContextTagParseWithBufferProducer[BACnetContextTagCharacterString]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_CHARACTER_STRING)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectName' field"))
	}
	m.ObjectName = objectName

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestWhoHasObjectName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetUnconfirmedServiceRequestWhoHasObjectName")
	}

	return m, nil
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectName) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectName) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestWhoHasObjectName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetUnconfirmedServiceRequestWhoHasObjectName")
		}

		if err := WriteSimpleField[BACnetContextTagCharacterString](ctx, "objectName", m.GetObjectName(), WriteComplex[BACnetContextTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'objectName' field")
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestWhoHasObjectName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetUnconfirmedServiceRequestWhoHasObjectName")
		}
		return nil
	}
	return m.BACnetUnconfirmedServiceRequestWhoHasObjectContract.(*_BACnetUnconfirmedServiceRequestWhoHasObject).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectName) IsBACnetUnconfirmedServiceRequestWhoHasObjectName() {
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectName) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
