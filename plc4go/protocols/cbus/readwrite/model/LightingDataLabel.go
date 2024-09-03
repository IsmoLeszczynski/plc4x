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

// LightingDataLabel is the corresponding interface of LightingDataLabel
type LightingDataLabel interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	LightingData
	// GetGroup returns Group (property field)
	GetGroup() byte
	// GetLabelOptions returns LabelOptions (property field)
	GetLabelOptions() LightingLabelOptions
	// GetLanguage returns Language (property field)
	GetLanguage() *Language
	// GetData returns Data (property field)
	GetData() []byte
	// IsLightingDataLabel is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsLightingDataLabel()
}

// _LightingDataLabel is the data-structure of this message
type _LightingDataLabel struct {
	LightingDataContract
	Group        byte
	LabelOptions LightingLabelOptions
	Language     *Language
	Data         []byte
}

var _ LightingDataLabel = (*_LightingDataLabel)(nil)
var _ LightingDataRequirements = (*_LightingDataLabel)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LightingDataLabel) GetParent() LightingDataContract {
	return m.LightingDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LightingDataLabel) GetGroup() byte {
	return m.Group
}

func (m *_LightingDataLabel) GetLabelOptions() LightingLabelOptions {
	return m.LabelOptions
}

func (m *_LightingDataLabel) GetLanguage() *Language {
	return m.Language
}

func (m *_LightingDataLabel) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewLightingDataLabel factory function for _LightingDataLabel
func NewLightingDataLabel(group byte, labelOptions LightingLabelOptions, language *Language, data []byte, commandTypeContainer LightingCommandTypeContainer) *_LightingDataLabel {
	if labelOptions == nil {
		panic("labelOptions of type LightingLabelOptions for LightingDataLabel must not be nil")
	}
	_result := &_LightingDataLabel{
		LightingDataContract: NewLightingData(commandTypeContainer),
		Group:                group,
		LabelOptions:         labelOptions,
		Language:             language,
		Data:                 data,
	}
	_result.LightingDataContract.(*_LightingData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastLightingDataLabel(structType any) LightingDataLabel {
	if casted, ok := structType.(LightingDataLabel); ok {
		return casted
	}
	if casted, ok := structType.(*LightingDataLabel); ok {
		return *casted
	}
	return nil
}

func (m *_LightingDataLabel) GetTypeName() string {
	return "LightingDataLabel"
}

func (m *_LightingDataLabel) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.LightingDataContract.(*_LightingData).getLengthInBits(ctx))

	// Simple field (group)
	lengthInBits += 8

	// Simple field (labelOptions)
	lengthInBits += m.LabelOptions.GetLengthInBits(ctx)

	// Optional Field (language)
	if m.Language != nil {
		lengthInBits += 8
	}

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_LightingDataLabel) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_LightingDataLabel) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_LightingData, commandTypeContainer LightingCommandTypeContainer) (__lightingDataLabel LightingDataLabel, err error) {
	m.LightingDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LightingDataLabel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LightingDataLabel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	group, err := ReadSimpleField(ctx, "group", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'group' field"))
	}
	m.Group = group

	labelOptions, err := ReadSimpleField[LightingLabelOptions](ctx, "labelOptions", ReadComplex[LightingLabelOptions](LightingLabelOptionsParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'labelOptions' field"))
	}
	m.LabelOptions = labelOptions

	var language *Language
	language, err = ReadOptionalField[Language](ctx, "language", ReadEnum(LanguageByValue, ReadUnsignedByte(readBuffer, uint8(8))), bool((labelOptions.GetLabelType()) != (LightingLabelType_LOAD_DYNAMIC_ICON)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'language' field"))
	}
	m.Language = language

	data, err := readBuffer.ReadByteArray("data", int((int32(commandTypeContainer.NumBytes()) - int32((utils.InlineIf((bool((labelOptions.GetLabelType()) != (LightingLabelType_LOAD_DYNAMIC_ICON))), func() any { return int32((int32(3))) }, func() any { return int32((int32(2))) }).(int32))))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("LightingDataLabel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LightingDataLabel")
	}

	return m, nil
}

func (m *_LightingDataLabel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LightingDataLabel) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LightingDataLabel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LightingDataLabel")
		}

		if err := WriteSimpleField[byte](ctx, "group", m.GetGroup(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'group' field")
		}

		if err := WriteSimpleField[LightingLabelOptions](ctx, "labelOptions", m.GetLabelOptions(), WriteComplex[LightingLabelOptions](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'labelOptions' field")
		}

		if err := WriteOptionalEnumField[Language](ctx, "language", "Language", m.GetLanguage(), WriteEnum[Language, uint8](Language.GetValue, Language.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8)), bool((m.GetLabelOptions().GetLabelType()) != (LightingLabelType_LOAD_DYNAMIC_ICON))); err != nil {
			return errors.Wrap(err, "Error serializing 'language' field")
		}

		if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("LightingDataLabel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LightingDataLabel")
		}
		return nil
	}
	return m.LightingDataContract.(*_LightingData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_LightingDataLabel) IsLightingDataLabel() {}

func (m *_LightingDataLabel) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
