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

// MFuncPropCon is the corresponding interface of MFuncPropCon
type MFuncPropCon interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CEMI
}

// MFuncPropConExactly can be used when we want exactly this type and not a type which fulfills MFuncPropCon.
// This is useful for switch cases.
type MFuncPropConExactly interface {
	MFuncPropCon
	isMFuncPropCon() bool
}

// _MFuncPropCon is the data-structure of this message
type _MFuncPropCon struct {
	CEMIContract
}

var _ MFuncPropCon = (*_MFuncPropCon)(nil)
var _ CEMIRequirements = (*_MFuncPropCon)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MFuncPropCon) GetMessageCode() uint8 {
	return 0xFA
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MFuncPropCon) GetParent() CEMIContract {
	return m.CEMIContract
}

// NewMFuncPropCon factory function for _MFuncPropCon
func NewMFuncPropCon(size uint16) *_MFuncPropCon {
	_result := &_MFuncPropCon{
		CEMIContract: NewCEMI(size),
	}
	_result.CEMIContract.(*_CEMI)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMFuncPropCon(structType any) MFuncPropCon {
	if casted, ok := structType.(MFuncPropCon); ok {
		return casted
	}
	if casted, ok := structType.(*MFuncPropCon); ok {
		return *casted
	}
	return nil
}

func (m *_MFuncPropCon) GetTypeName() string {
	return "MFuncPropCon"
}

func (m *_MFuncPropCon) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CEMIContract.(*_CEMI).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_MFuncPropCon) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MFuncPropCon) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CEMI, size uint16) (__mFuncPropCon MFuncPropCon, err error) {
	m.CEMIContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MFuncPropCon"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MFuncPropCon")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MFuncPropCon"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MFuncPropCon")
	}

	return m, nil
}

func (m *_MFuncPropCon) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MFuncPropCon) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MFuncPropCon"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MFuncPropCon")
		}

		if popErr := writeBuffer.PopContext("MFuncPropCon"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MFuncPropCon")
		}
		return nil
	}
	return m.CEMIContract.(*_CEMI).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MFuncPropCon) isMFuncPropCon() bool {
	return true
}

func (m *_MFuncPropCon) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
