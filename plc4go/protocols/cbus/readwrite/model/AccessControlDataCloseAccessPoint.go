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

// AccessControlDataCloseAccessPoint is the corresponding interface of AccessControlDataCloseAccessPoint
type AccessControlDataCloseAccessPoint interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AccessControlData
}

// _AccessControlDataCloseAccessPoint is the data-structure of this message
type _AccessControlDataCloseAccessPoint struct {
	AccessControlDataContract
}

var _ AccessControlDataCloseAccessPoint = (*_AccessControlDataCloseAccessPoint)(nil)
var _ AccessControlDataRequirements = (*_AccessControlDataCloseAccessPoint)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AccessControlDataCloseAccessPoint) GetParent() AccessControlDataContract {
	return m.AccessControlDataContract
}

// NewAccessControlDataCloseAccessPoint factory function for _AccessControlDataCloseAccessPoint
func NewAccessControlDataCloseAccessPoint(commandTypeContainer AccessControlCommandTypeContainer, networkId byte, accessPointId byte) *_AccessControlDataCloseAccessPoint {
	_result := &_AccessControlDataCloseAccessPoint{
		AccessControlDataContract: NewAccessControlData(commandTypeContainer, networkId, accessPointId),
	}
	_result.AccessControlDataContract.(*_AccessControlData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAccessControlDataCloseAccessPoint(structType any) AccessControlDataCloseAccessPoint {
	if casted, ok := structType.(AccessControlDataCloseAccessPoint); ok {
		return casted
	}
	if casted, ok := structType.(*AccessControlDataCloseAccessPoint); ok {
		return *casted
	}
	return nil
}

func (m *_AccessControlDataCloseAccessPoint) GetTypeName() string {
	return "AccessControlDataCloseAccessPoint"
}

func (m *_AccessControlDataCloseAccessPoint) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AccessControlDataContract.(*_AccessControlData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_AccessControlDataCloseAccessPoint) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AccessControlDataCloseAccessPoint) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AccessControlData) (__accessControlDataCloseAccessPoint AccessControlDataCloseAccessPoint, err error) {
	m.AccessControlDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AccessControlDataCloseAccessPoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AccessControlDataCloseAccessPoint")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("AccessControlDataCloseAccessPoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AccessControlDataCloseAccessPoint")
	}

	return m, nil
}

func (m *_AccessControlDataCloseAccessPoint) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AccessControlDataCloseAccessPoint) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AccessControlDataCloseAccessPoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AccessControlDataCloseAccessPoint")
		}

		if popErr := writeBuffer.PopContext("AccessControlDataCloseAccessPoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AccessControlDataCloseAccessPoint")
		}
		return nil
	}
	return m.AccessControlDataContract.(*_AccessControlData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AccessControlDataCloseAccessPoint) IsAccessControlDataCloseAccessPoint() {}

func (m *_AccessControlDataCloseAccessPoint) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
