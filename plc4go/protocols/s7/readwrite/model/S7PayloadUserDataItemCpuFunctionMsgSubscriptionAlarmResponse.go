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

// S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse is the corresponding interface of S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse
type S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	S7PayloadUserDataItem
	// GetResult returns Result (property field)
	GetResult() uint8
	// GetReserved01 returns Reserved01 (property field)
	GetReserved01() uint8
	// GetAlarmType returns AlarmType (property field)
	GetAlarmType() AlarmType
	// GetReserved02 returns Reserved02 (property field)
	GetReserved02() uint8
	// GetReserved03 returns Reserved03 (property field)
	GetReserved03() uint8
	// IsS7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse()
	// CreateBuilder creates a S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder
	CreateS7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder() S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder
}

// _S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse is the data-structure of this message
type _S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse struct {
	S7PayloadUserDataItemContract
	Result     uint8
	Reserved01 uint8
	AlarmType  AlarmType
	Reserved02 uint8
	Reserved03 uint8
}

var _ S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse = (*_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse)(nil)
var _ S7PayloadUserDataItemRequirements = (*_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse)(nil)

// NewS7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse factory function for _S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse
func NewS7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse(returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16, result uint8, reserved01 uint8, alarmType AlarmType, reserved02 uint8, reserved03 uint8) *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse {
	_result := &_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse{
		S7PayloadUserDataItemContract: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
		Result:                        result,
		Reserved01:                    reserved01,
		AlarmType:                     alarmType,
		Reserved02:                    reserved02,
		Reserved03:                    reserved03,
	}
	_result.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder is a builder for S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse
type S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(result uint8, reserved01 uint8, alarmType AlarmType, reserved02 uint8, reserved03 uint8) S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder
	// WithResult adds Result (property field)
	WithResult(uint8) S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder
	// WithReserved01 adds Reserved01 (property field)
	WithReserved01(uint8) S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder
	// WithAlarmType adds AlarmType (property field)
	WithAlarmType(AlarmType) S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder
	// WithReserved02 adds Reserved02 (property field)
	WithReserved02(uint8) S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder
	// WithReserved03 adds Reserved03 (property field)
	WithReserved03(uint8) S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() S7PayloadUserDataItemBuilder
	// Build builds the S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse or returns an error if something is wrong
	Build() (S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse
}

// NewS7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder() creates a S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder
func NewS7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder() S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder {
	return &_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder{_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse: new(_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse)}
}

type _S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder struct {
	*_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse

	parentBuilder *_S7PayloadUserDataItemBuilder

	err *utils.MultiError
}

var _ (S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder) = (*_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder)(nil)

func (b *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder) setParent(contract S7PayloadUserDataItemContract) {
	b.S7PayloadUserDataItemContract = contract
	contract.(*_S7PayloadUserDataItem)._SubType = b._S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse
}

func (b *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder) WithMandatoryFields(result uint8, reserved01 uint8, alarmType AlarmType, reserved02 uint8, reserved03 uint8) S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder {
	return b.WithResult(result).WithReserved01(reserved01).WithAlarmType(alarmType).WithReserved02(reserved02).WithReserved03(reserved03)
}

func (b *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder) WithResult(result uint8) S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder {
	b.Result = result
	return b
}

func (b *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder) WithReserved01(reserved01 uint8) S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder {
	b.Reserved01 = reserved01
	return b
}

func (b *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder) WithAlarmType(alarmType AlarmType) S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder {
	b.AlarmType = alarmType
	return b
}

func (b *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder) WithReserved02(reserved02 uint8) S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder {
	b.Reserved02 = reserved02
	return b
}

func (b *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder) WithReserved03(reserved03 uint8) S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder {
	b.Reserved03 = reserved03
	return b
}

func (b *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder) Build() (S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse.deepCopy(), nil
}

func (b *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder) MustBuild() S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder) Done() S7PayloadUserDataItemBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewS7PayloadUserDataItemBuilder().(*_S7PayloadUserDataItemBuilder)
	}
	return b.parentBuilder
}

func (b *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder) buildForS7PayloadUserDataItem() (S7PayloadUserDataItem, error) {
	return b.Build()
}

func (b *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder) DeepCopy() any {
	_copy := b.CreateS7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder().(*_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateS7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder creates a S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder
func (b *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) CreateS7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder() S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder {
	if b == nil {
		return NewS7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder()
	}
	return &_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseBuilder{_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetCpuFunctionGroup() uint8 {
	return 0x04
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetCpuFunctionType() uint8 {
	return 0x08
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetCpuSubfunction() uint8 {
	return 0x02
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetParent() S7PayloadUserDataItemContract {
	return m.S7PayloadUserDataItemContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetResult() uint8 {
	return m.Result
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetReserved01() uint8 {
	return m.Reserved01
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetAlarmType() AlarmType {
	return m.AlarmType
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetReserved02() uint8 {
	return m.Reserved02
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetReserved03() uint8 {
	return m.Reserved03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse(structType any) S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse {
	if casted, ok := structType.(S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetTypeName() string {
	return "S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse"
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).getLengthInBits(ctx))

	// Simple field (result)
	lengthInBits += 8

	// Simple field (reserved01)
	lengthInBits += 8

	// Simple field (alarmType)
	lengthInBits += 8

	// Simple field (reserved02)
	lengthInBits += 8

	// Simple field (reserved03)
	lengthInBits += 8

	return lengthInBits
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7PayloadUserDataItem, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (__s7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse, err error) {
	m.S7PayloadUserDataItemContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	result, err := ReadSimpleField(ctx, "result", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'result' field"))
	}
	m.Result = result

	reserved01, err := ReadSimpleField(ctx, "reserved01", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reserved01' field"))
	}
	m.Reserved01 = reserved01

	alarmType, err := ReadEnumField[AlarmType](ctx, "alarmType", "AlarmType", ReadEnum(AlarmTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'alarmType' field"))
	}
	m.AlarmType = alarmType

	reserved02, err := ReadSimpleField(ctx, "reserved02", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reserved02' field"))
	}
	m.Reserved02 = reserved02

	reserved03, err := ReadSimpleField(ctx, "reserved03", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reserved03' field"))
	}
	m.Reserved03 = reserved03

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse")
	}

	return m, nil
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse")
		}

		if err := WriteSimpleField[uint8](ctx, "result", m.GetResult(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'result' field")
		}

		if err := WriteSimpleField[uint8](ctx, "reserved01", m.GetReserved01(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved01' field")
		}

		if err := WriteSimpleEnumField[AlarmType](ctx, "alarmType", "AlarmType", m.GetAlarmType(), WriteEnum[AlarmType, uint8](AlarmType.GetValue, AlarmType.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'alarmType' field")
		}

		if err := WriteSimpleField[uint8](ctx, "reserved02", m.GetReserved02(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved02' field")
		}

		if err := WriteSimpleField[uint8](ctx, "reserved03", m.GetReserved03(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved03' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse")
		}
		return nil
	}
	return m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) IsS7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse() {
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) deepCopy() *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse {
	if m == nil {
		return nil
	}
	_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseCopy := &_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse{
		m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).deepCopy(),
		m.Result,
		m.Reserved01,
		m.AlarmType,
		m.Reserved02,
		m.Reserved03,
	}
	_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseCopy.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = m
	return _S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseCopy
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
