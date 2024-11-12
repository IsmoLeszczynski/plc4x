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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// AdsDeviceNotificationResponse is the corresponding interface of AdsDeviceNotificationResponse
type AdsDeviceNotificationResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AmsPacket
	// IsAdsDeviceNotificationResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsDeviceNotificationResponse()
	// CreateBuilder creates a AdsDeviceNotificationResponseBuilder
	CreateAdsDeviceNotificationResponseBuilder() AdsDeviceNotificationResponseBuilder
}

// _AdsDeviceNotificationResponse is the data-structure of this message
type _AdsDeviceNotificationResponse struct {
	AmsPacketContract
}

var _ AdsDeviceNotificationResponse = (*_AdsDeviceNotificationResponse)(nil)
var _ AmsPacketRequirements = (*_AdsDeviceNotificationResponse)(nil)

// NewAdsDeviceNotificationResponse factory function for _AdsDeviceNotificationResponse
func NewAdsDeviceNotificationResponse(targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32) *_AdsDeviceNotificationResponse {
	_result := &_AdsDeviceNotificationResponse{
		AmsPacketContract: NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId),
	}
	_result.AmsPacketContract.(*_AmsPacket)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AdsDeviceNotificationResponseBuilder is a builder for AdsDeviceNotificationResponse
type AdsDeviceNotificationResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() AdsDeviceNotificationResponseBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() AmsPacketBuilder
	// Build builds the AdsDeviceNotificationResponse or returns an error if something is wrong
	Build() (AdsDeviceNotificationResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AdsDeviceNotificationResponse
}

// NewAdsDeviceNotificationResponseBuilder() creates a AdsDeviceNotificationResponseBuilder
func NewAdsDeviceNotificationResponseBuilder() AdsDeviceNotificationResponseBuilder {
	return &_AdsDeviceNotificationResponseBuilder{_AdsDeviceNotificationResponse: new(_AdsDeviceNotificationResponse)}
}

type _AdsDeviceNotificationResponseBuilder struct {
	*_AdsDeviceNotificationResponse

	parentBuilder *_AmsPacketBuilder

	err *utils.MultiError
}

var _ (AdsDeviceNotificationResponseBuilder) = (*_AdsDeviceNotificationResponseBuilder)(nil)

func (b *_AdsDeviceNotificationResponseBuilder) setParent(contract AmsPacketContract) {
	b.AmsPacketContract = contract
}

func (b *_AdsDeviceNotificationResponseBuilder) WithMandatoryFields() AdsDeviceNotificationResponseBuilder {
	return b
}

func (b *_AdsDeviceNotificationResponseBuilder) Build() (AdsDeviceNotificationResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AdsDeviceNotificationResponse.deepCopy(), nil
}

func (b *_AdsDeviceNotificationResponseBuilder) MustBuild() AdsDeviceNotificationResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_AdsDeviceNotificationResponseBuilder) Done() AmsPacketBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewAmsPacketBuilder().(*_AmsPacketBuilder)
	}
	return b.parentBuilder
}

func (b *_AdsDeviceNotificationResponseBuilder) buildForAmsPacket() (AmsPacket, error) {
	return b.Build()
}

func (b *_AdsDeviceNotificationResponseBuilder) DeepCopy() any {
	_copy := b.CreateAdsDeviceNotificationResponseBuilder().(*_AdsDeviceNotificationResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAdsDeviceNotificationResponseBuilder creates a AdsDeviceNotificationResponseBuilder
func (b *_AdsDeviceNotificationResponse) CreateAdsDeviceNotificationResponseBuilder() AdsDeviceNotificationResponseBuilder {
	if b == nil {
		return NewAdsDeviceNotificationResponseBuilder()
	}
	return &_AdsDeviceNotificationResponseBuilder{_AdsDeviceNotificationResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsDeviceNotificationResponse) GetCommandId() CommandId {
	return CommandId_ADS_DEVICE_NOTIFICATION
}

func (m *_AdsDeviceNotificationResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsDeviceNotificationResponse) GetParent() AmsPacketContract {
	return m.AmsPacketContract
}

// Deprecated: use the interface for direct cast
func CastAdsDeviceNotificationResponse(structType any) AdsDeviceNotificationResponse {
	if casted, ok := structType.(AdsDeviceNotificationResponse); ok {
		return casted
	}
	if casted, ok := structType.(*AdsDeviceNotificationResponse); ok {
		return *casted
	}
	return nil
}

func (m *_AdsDeviceNotificationResponse) GetTypeName() string {
	return "AdsDeviceNotificationResponse"
}

func (m *_AdsDeviceNotificationResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AmsPacketContract.(*_AmsPacket).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_AdsDeviceNotificationResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AdsDeviceNotificationResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AmsPacket) (__adsDeviceNotificationResponse AdsDeviceNotificationResponse, err error) {
	m.AmsPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsDeviceNotificationResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsDeviceNotificationResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("AdsDeviceNotificationResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsDeviceNotificationResponse")
	}

	return m, nil
}

func (m *_AdsDeviceNotificationResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsDeviceNotificationResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsDeviceNotificationResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsDeviceNotificationResponse")
		}

		if popErr := writeBuffer.PopContext("AdsDeviceNotificationResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsDeviceNotificationResponse")
		}
		return nil
	}
	return m.AmsPacketContract.(*_AmsPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsDeviceNotificationResponse) IsAdsDeviceNotificationResponse() {}

func (m *_AdsDeviceNotificationResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AdsDeviceNotificationResponse) deepCopy() *_AdsDeviceNotificationResponse {
	if m == nil {
		return nil
	}
	_AdsDeviceNotificationResponseCopy := &_AdsDeviceNotificationResponse{
		m.AmsPacketContract.(*_AmsPacket).deepCopy(),
	}
	m.AmsPacketContract.(*_AmsPacket)._SubType = m
	return _AdsDeviceNotificationResponseCopy
}

func (m *_AdsDeviceNotificationResponse) String() string {
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
