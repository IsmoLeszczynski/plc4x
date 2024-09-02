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

// SecurityDataStatusReport1 is the corresponding interface of SecurityDataStatusReport1
type SecurityDataStatusReport1 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SecurityData
	// GetArmCodeType returns ArmCodeType (property field)
	GetArmCodeType() SecurityArmCode
	// GetTamperStatus returns TamperStatus (property field)
	GetTamperStatus() TamperStatus
	// GetPanicStatus returns PanicStatus (property field)
	GetPanicStatus() PanicStatus
	// GetZoneStatus returns ZoneStatus (property field)
	GetZoneStatus() []ZoneStatus
}

// SecurityDataStatusReport1Exactly can be used when we want exactly this type and not a type which fulfills SecurityDataStatusReport1.
// This is useful for switch cases.
type SecurityDataStatusReport1Exactly interface {
	SecurityDataStatusReport1
	isSecurityDataStatusReport1() bool
}

// _SecurityDataStatusReport1 is the data-structure of this message
type _SecurityDataStatusReport1 struct {
	SecurityDataContract
	ArmCodeType  SecurityArmCode
	TamperStatus TamperStatus
	PanicStatus  PanicStatus
	ZoneStatus   []ZoneStatus
}

var _ SecurityDataStatusReport1 = (*_SecurityDataStatusReport1)(nil)
var _ SecurityDataRequirements = (*_SecurityDataStatusReport1)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataStatusReport1) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataStatusReport1) GetArmCodeType() SecurityArmCode {
	return m.ArmCodeType
}

func (m *_SecurityDataStatusReport1) GetTamperStatus() TamperStatus {
	return m.TamperStatus
}

func (m *_SecurityDataStatusReport1) GetPanicStatus() PanicStatus {
	return m.PanicStatus
}

func (m *_SecurityDataStatusReport1) GetZoneStatus() []ZoneStatus {
	return m.ZoneStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSecurityDataStatusReport1 factory function for _SecurityDataStatusReport1
func NewSecurityDataStatusReport1(armCodeType SecurityArmCode, tamperStatus TamperStatus, panicStatus PanicStatus, zoneStatus []ZoneStatus, commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataStatusReport1 {
	_result := &_SecurityDataStatusReport1{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
		ArmCodeType:          armCodeType,
		TamperStatus:         tamperStatus,
		PanicStatus:          panicStatus,
		ZoneStatus:           zoneStatus,
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataStatusReport1(structType any) SecurityDataStatusReport1 {
	if casted, ok := structType.(SecurityDataStatusReport1); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataStatusReport1); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataStatusReport1) GetTypeName() string {
	return "SecurityDataStatusReport1"
}

func (m *_SecurityDataStatusReport1) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	// Simple field (armCodeType)
	lengthInBits += m.ArmCodeType.GetLengthInBits(ctx)

	// Simple field (tamperStatus)
	lengthInBits += m.TamperStatus.GetLengthInBits(ctx)

	// Simple field (panicStatus)
	lengthInBits += m.PanicStatus.GetLengthInBits(ctx)

	// Array field
	if len(m.ZoneStatus) > 0 {
		for _curItem, element := range m.ZoneStatus {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ZoneStatus), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_SecurityDataStatusReport1) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataStatusReport1) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataStatusReport1 SecurityDataStatusReport1, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataStatusReport1"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataStatusReport1")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	armCodeType, err := ReadSimpleField[SecurityArmCode](ctx, "armCodeType", ReadComplex[SecurityArmCode](SecurityArmCodeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'armCodeType' field"))
	}
	m.ArmCodeType = armCodeType

	tamperStatus, err := ReadSimpleField[TamperStatus](ctx, "tamperStatus", ReadComplex[TamperStatus](TamperStatusParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'tamperStatus' field"))
	}
	m.TamperStatus = tamperStatus

	panicStatus, err := ReadSimpleField[PanicStatus](ctx, "panicStatus", ReadComplex[PanicStatus](PanicStatusParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'panicStatus' field"))
	}
	m.PanicStatus = panicStatus

	zoneStatus, err := ReadCountArrayField[ZoneStatus](ctx, "zoneStatus", ReadComplex[ZoneStatus](ZoneStatusParseWithBuffer, readBuffer), uint64(int32(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneStatus' field"))
	}
	m.ZoneStatus = zoneStatus

	if closeErr := readBuffer.CloseContext("SecurityDataStatusReport1"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataStatusReport1")
	}

	return m, nil
}

func (m *_SecurityDataStatusReport1) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataStatusReport1) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataStatusReport1"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataStatusReport1")
		}

		if err := WriteSimpleField[SecurityArmCode](ctx, "armCodeType", m.GetArmCodeType(), WriteComplex[SecurityArmCode](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'armCodeType' field")
		}

		if err := WriteSimpleField[TamperStatus](ctx, "tamperStatus", m.GetTamperStatus(), WriteComplex[TamperStatus](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'tamperStatus' field")
		}

		if err := WriteSimpleField[PanicStatus](ctx, "panicStatus", m.GetPanicStatus(), WriteComplex[PanicStatus](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'panicStatus' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "zoneStatus", m.GetZoneStatus(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneStatus' field")
		}

		if popErr := writeBuffer.PopContext("SecurityDataStatusReport1"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataStatusReport1")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataStatusReport1) isSecurityDataStatusReport1() bool {
	return true
}

func (m *_SecurityDataStatusReport1) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
