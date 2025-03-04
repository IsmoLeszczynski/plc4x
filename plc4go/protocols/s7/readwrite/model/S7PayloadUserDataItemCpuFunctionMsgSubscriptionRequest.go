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

// S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest is the corresponding interface of S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest
type S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	S7PayloadUserDataItem
	// GetSubscription returns Subscription (property field)
	GetSubscription() uint8
	// GetMagicKey returns MagicKey (property field)
	GetMagicKey() string
	// GetAlarmtype returns Alarmtype (property field)
	GetAlarmtype() *AlarmStateType
	// GetReserve returns Reserve (property field)
	GetReserve() *uint8
	// IsS7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest()
	// CreateBuilder creates a S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder
	CreateS7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder() S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder
}

// _S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest is the data-structure of this message
type _S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest struct {
	S7PayloadUserDataItemContract
	Subscription uint8
	MagicKey     string
	Alarmtype    *AlarmStateType
	Reserve      *uint8
	// Reserved Fields
	reservedField0 *uint8
}

var _ S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest = (*_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest)(nil)
var _ S7PayloadUserDataItemRequirements = (*_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest)(nil)

// NewS7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest factory function for _S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest
func NewS7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest(returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16, subscription uint8, magicKey string, alarmtype *AlarmStateType, reserve *uint8) *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest {
	_result := &_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest{
		S7PayloadUserDataItemContract: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
		Subscription:                  subscription,
		MagicKey:                      magicKey,
		Alarmtype:                     alarmtype,
		Reserve:                       reserve,
	}
	_result.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder is a builder for S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest
type S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(subscription uint8, magicKey string) S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder
	// WithSubscription adds Subscription (property field)
	WithSubscription(uint8) S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder
	// WithMagicKey adds MagicKey (property field)
	WithMagicKey(string) S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder
	// WithAlarmtype adds Alarmtype (property field)
	WithOptionalAlarmtype(AlarmStateType) S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder
	// WithReserve adds Reserve (property field)
	WithOptionalReserve(uint8) S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() S7PayloadUserDataItemBuilder
	// Build builds the S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest or returns an error if something is wrong
	Build() (S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest
}

// NewS7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder() creates a S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder
func NewS7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder() S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder {
	return &_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder{_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest: new(_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest)}
}

type _S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder struct {
	*_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest

	parentBuilder *_S7PayloadUserDataItemBuilder

	err *utils.MultiError
}

var _ (S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder) = (*_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder)(nil)

func (b *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder) setParent(contract S7PayloadUserDataItemContract) {
	b.S7PayloadUserDataItemContract = contract
	contract.(*_S7PayloadUserDataItem)._SubType = b._S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest
}

func (b *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder) WithMandatoryFields(subscription uint8, magicKey string) S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder {
	return b.WithSubscription(subscription).WithMagicKey(magicKey)
}

func (b *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder) WithSubscription(subscription uint8) S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder {
	b.Subscription = subscription
	return b
}

func (b *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder) WithMagicKey(magicKey string) S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder {
	b.MagicKey = magicKey
	return b
}

func (b *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder) WithOptionalAlarmtype(alarmtype AlarmStateType) S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder {
	b.Alarmtype = &alarmtype
	return b
}

func (b *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder) WithOptionalReserve(reserve uint8) S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder {
	b.Reserve = &reserve
	return b
}

func (b *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder) Build() (S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest.deepCopy(), nil
}

func (b *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder) MustBuild() S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder) Done() S7PayloadUserDataItemBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewS7PayloadUserDataItemBuilder().(*_S7PayloadUserDataItemBuilder)
	}
	return b.parentBuilder
}

func (b *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder) buildForS7PayloadUserDataItem() (S7PayloadUserDataItem, error) {
	return b.Build()
}

func (b *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder) DeepCopy() any {
	_copy := b.CreateS7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder().(*_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateS7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder creates a S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder
func (b *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) CreateS7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder() S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder {
	if b == nil {
		return NewS7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder()
	}
	return &_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestBuilder{_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) GetCpuFunctionGroup() uint8 {
	return 0x04
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) GetCpuFunctionType() uint8 {
	return 0x04
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) GetCpuSubfunction() uint8 {
	return 0x02
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) GetParent() S7PayloadUserDataItemContract {
	return m.S7PayloadUserDataItemContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) GetSubscription() uint8 {
	return m.Subscription
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) GetMagicKey() string {
	return m.MagicKey
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) GetAlarmtype() *AlarmStateType {
	return m.Alarmtype
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) GetReserve() *uint8 {
	return m.Reserve
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest(structType any) S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest {
	if casted, ok := structType.(S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) GetTypeName() string {
	return "S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest"
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).getLengthInBits(ctx))

	// Simple field (subscription)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (magicKey)
	lengthInBits += 64

	// Optional Field (alarmtype)
	if m.Alarmtype != nil {
		lengthInBits += 8
	}

	// Optional Field (reserve)
	if m.Reserve != nil {
		lengthInBits += 8
	}

	return lengthInBits
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7PayloadUserDataItem, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (__s7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest, err error) {
	m.S7PayloadUserDataItemContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	subscription, err := ReadSimpleField(ctx, "subscription", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subscription' field"))
	}
	m.Subscription = subscription

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	magicKey, err := ReadSimpleField(ctx, "magicKey", ReadString(readBuffer, uint32(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'magicKey' field"))
	}
	m.MagicKey = magicKey

	var alarmtype *AlarmStateType
	alarmtype, err = ReadOptionalField[AlarmStateType](ctx, "alarmtype", ReadEnum(AlarmStateTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))), bool((subscription) >= (128)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'alarmtype' field"))
	}
	m.Alarmtype = alarmtype

	var reserve *uint8
	reserve, err = ReadOptionalField[uint8](ctx, "reserve", ReadUnsignedByte(readBuffer, uint8(8)), bool((subscription) >= (128)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reserve' field"))
	}
	m.Reserve = reserve

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest")
	}

	return m, nil
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest")
		}

		if err := WriteSimpleField[uint8](ctx, "subscription", m.GetSubscription(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'subscription' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[string](ctx, "magicKey", m.GetMagicKey(), WriteString(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'magicKey' field")
		}

		if err := WriteOptionalEnumField[AlarmStateType](ctx, "alarmtype", "AlarmStateType", m.GetAlarmtype(), WriteEnum[AlarmStateType, uint8](AlarmStateType.GetValue, AlarmStateType.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8)), bool((m.GetSubscription()) >= (128))); err != nil {
			return errors.Wrap(err, "Error serializing 'alarmtype' field")
		}

		if err := WriteOptionalField[uint8](ctx, "reserve", m.GetReserve(), WriteUnsignedByte(writeBuffer, 8), true); err != nil {
			return errors.Wrap(err, "Error serializing 'reserve' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest")
		}
		return nil
	}
	return m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) IsS7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest() {
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) deepCopy() *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest {
	if m == nil {
		return nil
	}
	_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestCopy := &_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest{
		m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).deepCopy(),
		m.Subscription,
		m.MagicKey,
		utils.CopyPtr[AlarmStateType](m.Alarmtype),
		utils.CopyPtr[uint8](m.Reserve),
		m.reservedField0,
	}
	_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestCopy.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = m
	return _S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestCopy
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest) String() string {
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
