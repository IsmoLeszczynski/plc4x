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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// SecurityDataArmFailedCleared is the corresponding interface of SecurityDataArmFailedCleared
type SecurityDataArmFailedCleared interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SecurityData
}

// _SecurityDataArmFailedCleared is the data-structure of this message
type _SecurityDataArmFailedCleared struct {
	SecurityDataContract
}

var _ SecurityDataArmFailedCleared = (*_SecurityDataArmFailedCleared)(nil)
var _ SecurityDataRequirements = (*_SecurityDataArmFailedCleared)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataArmFailedCleared) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

// NewSecurityDataArmFailedCleared factory function for _SecurityDataArmFailedCleared
func NewSecurityDataArmFailedCleared(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataArmFailedCleared {
	_result := &_SecurityDataArmFailedCleared{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataArmFailedCleared(structType any) SecurityDataArmFailedCleared {
	if casted, ok := structType.(SecurityDataArmFailedCleared); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataArmFailedCleared); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataArmFailedCleared) GetTypeName() string {
	return "SecurityDataArmFailedCleared"
}

func (m *_SecurityDataArmFailedCleared) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataArmFailedCleared) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataArmFailedCleared) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataArmFailedCleared SecurityDataArmFailedCleared, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataArmFailedCleared"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataArmFailedCleared")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataArmFailedCleared"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataArmFailedCleared")
	}

	return m, nil
}

func (m *_SecurityDataArmFailedCleared) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataArmFailedCleared) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataArmFailedCleared"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataArmFailedCleared")
		}

		if popErr := writeBuffer.PopContext("SecurityDataArmFailedCleared"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataArmFailedCleared")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataArmFailedCleared) IsSecurityDataArmFailedCleared() {}

func (m *_SecurityDataArmFailedCleared) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
