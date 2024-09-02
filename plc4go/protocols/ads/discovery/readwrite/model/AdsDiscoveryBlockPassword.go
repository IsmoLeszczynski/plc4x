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

// AdsDiscoveryBlockPassword is the corresponding interface of AdsDiscoveryBlockPassword
type AdsDiscoveryBlockPassword interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AdsDiscoveryBlock
	// GetPassword returns Password (property field)
	GetPassword() AmsString
}

// AdsDiscoveryBlockPasswordExactly can be used when we want exactly this type and not a type which fulfills AdsDiscoveryBlockPassword.
// This is useful for switch cases.
type AdsDiscoveryBlockPasswordExactly interface {
	AdsDiscoveryBlockPassword
	isAdsDiscoveryBlockPassword() bool
}

// _AdsDiscoveryBlockPassword is the data-structure of this message
type _AdsDiscoveryBlockPassword struct {
	AdsDiscoveryBlockContract
	Password AmsString
}

var _ AdsDiscoveryBlockPassword = (*_AdsDiscoveryBlockPassword)(nil)
var _ AdsDiscoveryBlockRequirements = (*_AdsDiscoveryBlockPassword)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsDiscoveryBlockPassword) GetBlockType() AdsDiscoveryBlockType {
	return AdsDiscoveryBlockType_PASSWORD
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsDiscoveryBlockPassword) GetParent() AdsDiscoveryBlockContract {
	return m.AdsDiscoveryBlockContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsDiscoveryBlockPassword) GetPassword() AmsString {
	return m.Password
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsDiscoveryBlockPassword factory function for _AdsDiscoveryBlockPassword
func NewAdsDiscoveryBlockPassword(password AmsString) *_AdsDiscoveryBlockPassword {
	_result := &_AdsDiscoveryBlockPassword{
		AdsDiscoveryBlockContract: NewAdsDiscoveryBlock(),
		Password:                  password,
	}
	_result.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsDiscoveryBlockPassword(structType any) AdsDiscoveryBlockPassword {
	if casted, ok := structType.(AdsDiscoveryBlockPassword); ok {
		return casted
	}
	if casted, ok := structType.(*AdsDiscoveryBlockPassword); ok {
		return *casted
	}
	return nil
}

func (m *_AdsDiscoveryBlockPassword) GetTypeName() string {
	return "AdsDiscoveryBlockPassword"
}

func (m *_AdsDiscoveryBlockPassword) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock).getLengthInBits(ctx))

	// Simple field (password)
	lengthInBits += m.Password.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_AdsDiscoveryBlockPassword) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AdsDiscoveryBlockPassword) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AdsDiscoveryBlock) (__adsDiscoveryBlockPassword AdsDiscoveryBlockPassword, err error) {
	m.AdsDiscoveryBlockContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsDiscoveryBlockPassword"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsDiscoveryBlockPassword")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	password, err := ReadSimpleField[AmsString](ctx, "password", ReadComplex[AmsString](AmsStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'password' field"))
	}
	m.Password = password

	if closeErr := readBuffer.CloseContext("AdsDiscoveryBlockPassword"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsDiscoveryBlockPassword")
	}

	return m, nil
}

func (m *_AdsDiscoveryBlockPassword) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsDiscoveryBlockPassword) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsDiscoveryBlockPassword"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsDiscoveryBlockPassword")
		}

		if err := WriteSimpleField[AmsString](ctx, "password", m.GetPassword(), WriteComplex[AmsString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'password' field")
		}

		if popErr := writeBuffer.PopContext("AdsDiscoveryBlockPassword"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsDiscoveryBlockPassword")
		}
		return nil
	}
	return m.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsDiscoveryBlockPassword) isAdsDiscoveryBlockPassword() bool {
	return true
}

func (m *_AdsDiscoveryBlockPassword) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
