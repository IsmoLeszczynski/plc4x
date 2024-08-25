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

// CustomTypes is the corresponding interface of CustomTypes
type CustomTypes interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetCustomString returns CustomString (property field)
	GetCustomString() string
}

// CustomTypesExactly can be used when we want exactly this type and not a type which fulfills CustomTypes.
// This is useful for switch cases.
type CustomTypesExactly interface {
	CustomTypes
	isCustomTypes() bool
}

// _CustomTypes is the data-structure of this message
type _CustomTypes struct {
	CustomString string

	// Arguments.
	NumBytes uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CustomTypes) GetCustomString() string {
	return m.CustomString
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCustomTypes factory function for _CustomTypes
func NewCustomTypes(customString string, numBytes uint8) *_CustomTypes {
	return &_CustomTypes{CustomString: customString, NumBytes: numBytes}
}

// Deprecated: use the interface for direct cast
func CastCustomTypes(structType any) CustomTypes {
	if casted, ok := structType.(CustomTypes); ok {
		return casted
	}
	if casted, ok := structType.(*CustomTypes); ok {
		return *casted
	}
	return nil
}

func (m *_CustomTypes) GetTypeName() string {
	return "CustomTypes"
}

func (m *_CustomTypes) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (customString)
	lengthInBits += uint16(int32(int32(8)) * int32(m.NumBytes))

	return lengthInBits
}

func (m *_CustomTypes) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CustomTypesParse(ctx context.Context, theBytes []byte, numBytes uint8) (CustomTypes, error) {
	return CustomTypesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), numBytes)
}

func CustomTypesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, numBytes uint8) (CustomTypes, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("CustomTypes"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CustomTypes")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (customString)
	_customString, _customStringErr := readBuffer.ReadString("customString", uint32((8)*(numBytes)), "UTF-8")
	if _customStringErr != nil {
		return nil, errors.Wrap(_customStringErr, "Error parsing 'customString' field of CustomTypes")
	}
	customString := _customString

	if closeErr := readBuffer.CloseContext("CustomTypes"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CustomTypes")
	}

	// Create the instance
	return &_CustomTypes{
		NumBytes:     numBytes,
		CustomString: customString,
	}, nil
}

func (m *_CustomTypes) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CustomTypes) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("CustomTypes"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CustomTypes")
	}

	// Simple Field (customString)
	customString := string(m.GetCustomString())
	_customStringErr := writeBuffer.WriteString("customString", uint32((8)*(m.GetNumBytes())), "UTF-8", (customString))
	if _customStringErr != nil {
		return errors.Wrap(_customStringErr, "Error serializing 'customString' field")
	}

	if popErr := writeBuffer.PopContext("CustomTypes"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CustomTypes")
	}
	return nil
}

////
// Arguments Getter

func (m *_CustomTypes) GetNumBytes() uint8 {
	return m.NumBytes
}

//
////

func (m *_CustomTypes) isCustomTypes() bool {
	return true
}

func (m *_CustomTypes) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
