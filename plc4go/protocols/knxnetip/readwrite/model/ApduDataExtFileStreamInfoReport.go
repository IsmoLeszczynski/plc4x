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

// ApduDataExtFileStreamInfoReport is the corresponding interface of ApduDataExtFileStreamInfoReport
type ApduDataExtFileStreamInfoReport interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ApduDataExt
}

// _ApduDataExtFileStreamInfoReport is the data-structure of this message
type _ApduDataExtFileStreamInfoReport struct {
	ApduDataExtContract
}

var _ ApduDataExtFileStreamInfoReport = (*_ApduDataExtFileStreamInfoReport)(nil)
var _ ApduDataExtRequirements = (*_ApduDataExtFileStreamInfoReport)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtFileStreamInfoReport) GetExtApciType() uint8 {
	return 0x30
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtFileStreamInfoReport) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

// NewApduDataExtFileStreamInfoReport factory function for _ApduDataExtFileStreamInfoReport
func NewApduDataExtFileStreamInfoReport(length uint8) *_ApduDataExtFileStreamInfoReport {
	_result := &_ApduDataExtFileStreamInfoReport{
		ApduDataExtContract: NewApduDataExt(length),
	}
	_result.ApduDataExtContract.(*_ApduDataExt)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtFileStreamInfoReport(structType any) ApduDataExtFileStreamInfoReport {
	if casted, ok := structType.(ApduDataExtFileStreamInfoReport); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtFileStreamInfoReport); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtFileStreamInfoReport) GetTypeName() string {
	return "ApduDataExtFileStreamInfoReport"
}

func (m *_ApduDataExtFileStreamInfoReport) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtFileStreamInfoReport) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtFileStreamInfoReport) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (__apduDataExtFileStreamInfoReport ApduDataExtFileStreamInfoReport, err error) {
	m.ApduDataExtContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtFileStreamInfoReport"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtFileStreamInfoReport")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtFileStreamInfoReport"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtFileStreamInfoReport")
	}

	return m, nil
}

func (m *_ApduDataExtFileStreamInfoReport) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtFileStreamInfoReport) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtFileStreamInfoReport"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtFileStreamInfoReport")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtFileStreamInfoReport"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtFileStreamInfoReport")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtFileStreamInfoReport) IsApduDataExtFileStreamInfoReport() {}

func (m *_ApduDataExtFileStreamInfoReport) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
