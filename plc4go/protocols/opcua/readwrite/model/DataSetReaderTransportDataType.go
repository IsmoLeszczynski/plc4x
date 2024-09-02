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

// DataSetReaderTransportDataType is the corresponding interface of DataSetReaderTransportDataType
type DataSetReaderTransportDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
}

// DataSetReaderTransportDataTypeExactly can be used when we want exactly this type and not a type which fulfills DataSetReaderTransportDataType.
// This is useful for switch cases.
type DataSetReaderTransportDataTypeExactly interface {
	DataSetReaderTransportDataType
	isDataSetReaderTransportDataType() bool
}

// _DataSetReaderTransportDataType is the data-structure of this message
type _DataSetReaderTransportDataType struct {
	ExtensionObjectDefinitionContract
}

var _ DataSetReaderTransportDataType = (*_DataSetReaderTransportDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_DataSetReaderTransportDataType)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DataSetReaderTransportDataType) GetIdentifier() string {
	return "15630"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DataSetReaderTransportDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

// NewDataSetReaderTransportDataType factory function for _DataSetReaderTransportDataType
func NewDataSetReaderTransportDataType() *_DataSetReaderTransportDataType {
	_result := &_DataSetReaderTransportDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDataSetReaderTransportDataType(structType any) DataSetReaderTransportDataType {
	if casted, ok := structType.(DataSetReaderTransportDataType); ok {
		return casted
	}
	if casted, ok := structType.(*DataSetReaderTransportDataType); ok {
		return *casted
	}
	return nil
}

func (m *_DataSetReaderTransportDataType) GetTypeName() string {
	return "DataSetReaderTransportDataType"
}

func (m *_DataSetReaderTransportDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_DataSetReaderTransportDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DataSetReaderTransportDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__dataSetReaderTransportDataType DataSetReaderTransportDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DataSetReaderTransportDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DataSetReaderTransportDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("DataSetReaderTransportDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DataSetReaderTransportDataType")
	}

	return m, nil
}

func (m *_DataSetReaderTransportDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DataSetReaderTransportDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DataSetReaderTransportDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DataSetReaderTransportDataType")
		}

		if popErr := writeBuffer.PopContext("DataSetReaderTransportDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DataSetReaderTransportDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DataSetReaderTransportDataType) isDataSetReaderTransportDataType() bool {
	return true
}

func (m *_DataSetReaderTransportDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
