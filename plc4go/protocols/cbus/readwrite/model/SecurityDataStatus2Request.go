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

// SecurityDataStatus2Request is the corresponding interface of SecurityDataStatus2Request
type SecurityDataStatus2Request interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SecurityData
}

// SecurityDataStatus2RequestExactly can be used when we want exactly this type and not a type which fulfills SecurityDataStatus2Request.
// This is useful for switch cases.
type SecurityDataStatus2RequestExactly interface {
	SecurityDataStatus2Request
	isSecurityDataStatus2Request() bool
}

// _SecurityDataStatus2Request is the data-structure of this message
type _SecurityDataStatus2Request struct {
	SecurityDataContract
}

var _ SecurityDataStatus2Request = (*_SecurityDataStatus2Request)(nil)
var _ SecurityDataRequirements = (*_SecurityDataStatus2Request)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataStatus2Request) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

// NewSecurityDataStatus2Request factory function for _SecurityDataStatus2Request
func NewSecurityDataStatus2Request(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataStatus2Request {
	_result := &_SecurityDataStatus2Request{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataStatus2Request(structType any) SecurityDataStatus2Request {
	if casted, ok := structType.(SecurityDataStatus2Request); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataStatus2Request); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataStatus2Request) GetTypeName() string {
	return "SecurityDataStatus2Request"
}

func (m *_SecurityDataStatus2Request) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataStatus2Request) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataStatus2Request) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataStatus2Request SecurityDataStatus2Request, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataStatus2Request"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataStatus2Request")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataStatus2Request"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataStatus2Request")
	}

	return m, nil
}

func (m *_SecurityDataStatus2Request) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataStatus2Request) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataStatus2Request"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataStatus2Request")
		}

		if popErr := writeBuffer.PopContext("SecurityDataStatus2Request"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataStatus2Request")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataStatus2Request) isSecurityDataStatus2Request() bool {
	return true
}

func (m *_SecurityDataStatus2Request) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
