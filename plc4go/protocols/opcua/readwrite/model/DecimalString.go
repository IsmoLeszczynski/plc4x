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

// DecimalString is the corresponding interface of DecimalString
type DecimalString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// DecimalStringExactly can be used when we want exactly this type and not a type which fulfills DecimalString.
// This is useful for switch cases.
type DecimalStringExactly interface {
	DecimalString
	isDecimalString() bool
}

// _DecimalString is the data-structure of this message
type _DecimalString struct {
}

// NewDecimalString factory function for _DecimalString
func NewDecimalString() *_DecimalString {
	return &_DecimalString{}
}

// Deprecated: use the interface for direct cast
func CastDecimalString(structType any) DecimalString {
	if casted, ok := structType.(DecimalString); ok {
		return casted
	}
	if casted, ok := structType.(*DecimalString); ok {
		return *casted
	}
	return nil
}

func (m *_DecimalString) GetTypeName() string {
	return "DecimalString"
}

func (m *_DecimalString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_DecimalString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DecimalStringParse(ctx context.Context, theBytes []byte) (DecimalString, error) {
	return DecimalStringParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func DecimalStringParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DecimalString, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("DecimalString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DecimalString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("DecimalString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DecimalString")
	}

	// Create the instance
	return &_DecimalString{}, nil
}

func (m *_DecimalString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DecimalString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("DecimalString"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for DecimalString")
	}

	if popErr := writeBuffer.PopContext("DecimalString"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for DecimalString")
	}
	return nil
}

func (m *_DecimalString) isDecimalString() bool {
	return true
}

func (m *_DecimalString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
