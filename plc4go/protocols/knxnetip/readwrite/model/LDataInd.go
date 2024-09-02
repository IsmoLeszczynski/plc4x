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
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// LDataInd is the corresponding interface of LDataInd
type LDataInd interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CEMI
	// GetAdditionalInformationLength returns AdditionalInformationLength (property field)
	GetAdditionalInformationLength() uint8
	// GetAdditionalInformation returns AdditionalInformation (property field)
	GetAdditionalInformation() []CEMIAdditionalInformation
	// GetDataFrame returns DataFrame (property field)
	GetDataFrame() LDataFrame
}

// _LDataInd is the data-structure of this message
type _LDataInd struct {
	CEMIContract
	AdditionalInformationLength uint8
	AdditionalInformation       []CEMIAdditionalInformation
	DataFrame                   LDataFrame
}

var _ LDataInd = (*_LDataInd)(nil)
var _ CEMIRequirements = (*_LDataInd)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_LDataInd) GetMessageCode() uint8 {
	return 0x29
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LDataInd) GetParent() CEMIContract {
	return m.CEMIContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LDataInd) GetAdditionalInformationLength() uint8 {
	return m.AdditionalInformationLength
}

func (m *_LDataInd) GetAdditionalInformation() []CEMIAdditionalInformation {
	return m.AdditionalInformation
}

func (m *_LDataInd) GetDataFrame() LDataFrame {
	return m.DataFrame
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewLDataInd factory function for _LDataInd
func NewLDataInd(additionalInformationLength uint8, additionalInformation []CEMIAdditionalInformation, dataFrame LDataFrame, size uint16) *_LDataInd {
	_result := &_LDataInd{
		CEMIContract:                NewCEMI(size),
		AdditionalInformationLength: additionalInformationLength,
		AdditionalInformation:       additionalInformation,
		DataFrame:                   dataFrame,
	}
	_result.CEMIContract.(*_CEMI)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastLDataInd(structType any) LDataInd {
	if casted, ok := structType.(LDataInd); ok {
		return casted
	}
	if casted, ok := structType.(*LDataInd); ok {
		return *casted
	}
	return nil
}

func (m *_LDataInd) GetTypeName() string {
	return "LDataInd"
}

func (m *_LDataInd) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CEMIContract.(*_CEMI).getLengthInBits(ctx))

	// Simple field (additionalInformationLength)
	lengthInBits += 8

	// Array field
	if len(m.AdditionalInformation) > 0 {
		for _, element := range m.AdditionalInformation {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (dataFrame)
	lengthInBits += m.DataFrame.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_LDataInd) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_LDataInd) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CEMI, size uint16) (__lDataInd LDataInd, err error) {
	m.CEMIContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LDataInd"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LDataInd")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	additionalInformationLength, err := ReadSimpleField(ctx, "additionalInformationLength", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'additionalInformationLength' field"))
	}
	m.AdditionalInformationLength = additionalInformationLength

	additionalInformation, err := ReadLengthArrayField[CEMIAdditionalInformation](ctx, "additionalInformation", ReadComplex[CEMIAdditionalInformation](CEMIAdditionalInformationParseWithBuffer, readBuffer), int(additionalInformationLength))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'additionalInformation' field"))
	}
	m.AdditionalInformation = additionalInformation

	dataFrame, err := ReadSimpleField[LDataFrame](ctx, "dataFrame", ReadComplex[LDataFrame](LDataFrameParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataFrame' field"))
	}
	m.DataFrame = dataFrame

	if closeErr := readBuffer.CloseContext("LDataInd"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LDataInd")
	}

	return m, nil
}

func (m *_LDataInd) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LDataInd) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LDataInd"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LDataInd")
		}

		if err := WriteSimpleField[uint8](ctx, "additionalInformationLength", m.GetAdditionalInformationLength(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'additionalInformationLength' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "additionalInformation", m.GetAdditionalInformation(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'additionalInformation' field")
		}

		if err := WriteSimpleField[LDataFrame](ctx, "dataFrame", m.GetDataFrame(), WriteComplex[LDataFrame](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dataFrame' field")
		}

		if popErr := writeBuffer.PopContext("LDataInd"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LDataInd")
		}
		return nil
	}
	return m.CEMIContract.(*_CEMI).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_LDataInd) IsLDataInd() {}

func (m *_LDataInd) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
