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

// DataSetReaderMessageDataType is the corresponding interface of DataSetReaderMessageDataType
type DataSetReaderMessageDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
}

// DataSetReaderMessageDataTypeExactly can be used when we want exactly this type and not a type which fulfills DataSetReaderMessageDataType.
// This is useful for switch cases.
type DataSetReaderMessageDataTypeExactly interface {
	DataSetReaderMessageDataType
	isDataSetReaderMessageDataType() bool
}

// _DataSetReaderMessageDataType is the data-structure of this message
type _DataSetReaderMessageDataType struct {
	ExtensionObjectDefinitionContract
}

var _ DataSetReaderMessageDataType = (*_DataSetReaderMessageDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_DataSetReaderMessageDataType)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DataSetReaderMessageDataType) GetIdentifier() string {
	return "15631"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DataSetReaderMessageDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

// NewDataSetReaderMessageDataType factory function for _DataSetReaderMessageDataType
func NewDataSetReaderMessageDataType() *_DataSetReaderMessageDataType {
	_result := &_DataSetReaderMessageDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDataSetReaderMessageDataType(structType any) DataSetReaderMessageDataType {
	if casted, ok := structType.(DataSetReaderMessageDataType); ok {
		return casted
	}
	if casted, ok := structType.(*DataSetReaderMessageDataType); ok {
		return *casted
	}
	return nil
}

func (m *_DataSetReaderMessageDataType) GetTypeName() string {
	return "DataSetReaderMessageDataType"
}

func (m *_DataSetReaderMessageDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_DataSetReaderMessageDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DataSetReaderMessageDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__dataSetReaderMessageDataType DataSetReaderMessageDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DataSetReaderMessageDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DataSetReaderMessageDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("DataSetReaderMessageDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DataSetReaderMessageDataType")
	}

	return m, nil
}

func (m *_DataSetReaderMessageDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DataSetReaderMessageDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DataSetReaderMessageDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DataSetReaderMessageDataType")
		}

		if popErr := writeBuffer.PopContext("DataSetReaderMessageDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DataSetReaderMessageDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DataSetReaderMessageDataType) isDataSetReaderMessageDataType() bool {
	return true
}

func (m *_DataSetReaderMessageDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
