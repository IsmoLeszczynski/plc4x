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

// DataChangeNotification is the corresponding interface of DataChangeNotification
type DataChangeNotification interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetMonitoredItems returns MonitoredItems (property field)
	GetMonitoredItems() []MonitoredItemNotification
	// GetDiagnosticInfos returns DiagnosticInfos (property field)
	GetDiagnosticInfos() []DiagnosticInfo
	// IsDataChangeNotification is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDataChangeNotification()
	// CreateBuilder creates a DataChangeNotificationBuilder
	CreateDataChangeNotificationBuilder() DataChangeNotificationBuilder
}

// _DataChangeNotification is the data-structure of this message
type _DataChangeNotification struct {
	ExtensionObjectDefinitionContract
	MonitoredItems  []MonitoredItemNotification
	DiagnosticInfos []DiagnosticInfo
}

var _ DataChangeNotification = (*_DataChangeNotification)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_DataChangeNotification)(nil)

// NewDataChangeNotification factory function for _DataChangeNotification
func NewDataChangeNotification(monitoredItems []MonitoredItemNotification, diagnosticInfos []DiagnosticInfo) *_DataChangeNotification {
	_result := &_DataChangeNotification{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		MonitoredItems:                    monitoredItems,
		DiagnosticInfos:                   diagnosticInfos,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DataChangeNotificationBuilder is a builder for DataChangeNotification
type DataChangeNotificationBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(monitoredItems []MonitoredItemNotification, diagnosticInfos []DiagnosticInfo) DataChangeNotificationBuilder
	// WithMonitoredItems adds MonitoredItems (property field)
	WithMonitoredItems(...MonitoredItemNotification) DataChangeNotificationBuilder
	// WithDiagnosticInfos adds DiagnosticInfos (property field)
	WithDiagnosticInfos(...DiagnosticInfo) DataChangeNotificationBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the DataChangeNotification or returns an error if something is wrong
	Build() (DataChangeNotification, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DataChangeNotification
}

// NewDataChangeNotificationBuilder() creates a DataChangeNotificationBuilder
func NewDataChangeNotificationBuilder() DataChangeNotificationBuilder {
	return &_DataChangeNotificationBuilder{_DataChangeNotification: new(_DataChangeNotification)}
}

type _DataChangeNotificationBuilder struct {
	*_DataChangeNotification

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (DataChangeNotificationBuilder) = (*_DataChangeNotificationBuilder)(nil)

func (b *_DataChangeNotificationBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_DataChangeNotificationBuilder) WithMandatoryFields(monitoredItems []MonitoredItemNotification, diagnosticInfos []DiagnosticInfo) DataChangeNotificationBuilder {
	return b.WithMonitoredItems(monitoredItems...).WithDiagnosticInfos(diagnosticInfos...)
}

func (b *_DataChangeNotificationBuilder) WithMonitoredItems(monitoredItems ...MonitoredItemNotification) DataChangeNotificationBuilder {
	b.MonitoredItems = monitoredItems
	return b
}

func (b *_DataChangeNotificationBuilder) WithDiagnosticInfos(diagnosticInfos ...DiagnosticInfo) DataChangeNotificationBuilder {
	b.DiagnosticInfos = diagnosticInfos
	return b
}

func (b *_DataChangeNotificationBuilder) Build() (DataChangeNotification, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._DataChangeNotification.deepCopy(), nil
}

func (b *_DataChangeNotificationBuilder) MustBuild() DataChangeNotification {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_DataChangeNotificationBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_DataChangeNotificationBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_DataChangeNotificationBuilder) DeepCopy() any {
	_copy := b.CreateDataChangeNotificationBuilder().(*_DataChangeNotificationBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDataChangeNotificationBuilder creates a DataChangeNotificationBuilder
func (b *_DataChangeNotification) CreateDataChangeNotificationBuilder() DataChangeNotificationBuilder {
	if b == nil {
		return NewDataChangeNotificationBuilder()
	}
	return &_DataChangeNotificationBuilder{_DataChangeNotification: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DataChangeNotification) GetExtensionId() int32 {
	return int32(811)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DataChangeNotification) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DataChangeNotification) GetMonitoredItems() []MonitoredItemNotification {
	return m.MonitoredItems
}

func (m *_DataChangeNotification) GetDiagnosticInfos() []DiagnosticInfo {
	return m.DiagnosticInfos
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDataChangeNotification(structType any) DataChangeNotification {
	if casted, ok := structType.(DataChangeNotification); ok {
		return casted
	}
	if casted, ok := structType.(*DataChangeNotification); ok {
		return *casted
	}
	return nil
}

func (m *_DataChangeNotification) GetTypeName() string {
	return "DataChangeNotification"
}

func (m *_DataChangeNotification) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Implicit Field (noOfMonitoredItems)
	lengthInBits += 32

	// Array field
	if len(m.MonitoredItems) > 0 {
		for _curItem, element := range m.MonitoredItems {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.MonitoredItems), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Implicit Field (noOfDiagnosticInfos)
	lengthInBits += 32

	// Array field
	if len(m.DiagnosticInfos) > 0 {
		for _curItem, element := range m.DiagnosticInfos {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.DiagnosticInfos), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_DataChangeNotification) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DataChangeNotification) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__dataChangeNotification DataChangeNotification, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DataChangeNotification"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DataChangeNotification")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	noOfMonitoredItems, err := ReadImplicitField[int32](ctx, "noOfMonitoredItems", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfMonitoredItems' field"))
	}
	_ = noOfMonitoredItems

	monitoredItems, err := ReadCountArrayField[MonitoredItemNotification](ctx, "monitoredItems", ReadComplex[MonitoredItemNotification](ExtensionObjectDefinitionParseWithBufferProducer[MonitoredItemNotification]((int32)(int32(808))), readBuffer), uint64(noOfMonitoredItems))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'monitoredItems' field"))
	}
	m.MonitoredItems = monitoredItems

	noOfDiagnosticInfos, err := ReadImplicitField[int32](ctx, "noOfDiagnosticInfos", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfDiagnosticInfos' field"))
	}
	_ = noOfDiagnosticInfos

	diagnosticInfos, err := ReadCountArrayField[DiagnosticInfo](ctx, "diagnosticInfos", ReadComplex[DiagnosticInfo](DiagnosticInfoParseWithBuffer, readBuffer), uint64(noOfDiagnosticInfos))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'diagnosticInfos' field"))
	}
	m.DiagnosticInfos = diagnosticInfos

	if closeErr := readBuffer.CloseContext("DataChangeNotification"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DataChangeNotification")
	}

	return m, nil
}

func (m *_DataChangeNotification) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DataChangeNotification) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DataChangeNotification"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DataChangeNotification")
		}
		noOfMonitoredItems := int32(utils.InlineIf(bool((m.GetMonitoredItems()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetMonitoredItems()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfMonitoredItems", noOfMonitoredItems, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfMonitoredItems' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "monitoredItems", m.GetMonitoredItems(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'monitoredItems' field")
		}
		noOfDiagnosticInfos := int32(utils.InlineIf(bool((m.GetDiagnosticInfos()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetDiagnosticInfos()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfDiagnosticInfos", noOfDiagnosticInfos, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfDiagnosticInfos' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "diagnosticInfos", m.GetDiagnosticInfos(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'diagnosticInfos' field")
		}

		if popErr := writeBuffer.PopContext("DataChangeNotification"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DataChangeNotification")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DataChangeNotification) IsDataChangeNotification() {}

func (m *_DataChangeNotification) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DataChangeNotification) deepCopy() *_DataChangeNotification {
	if m == nil {
		return nil
	}
	_DataChangeNotificationCopy := &_DataChangeNotification{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopySlice[MonitoredItemNotification, MonitoredItemNotification](m.MonitoredItems),
		utils.DeepCopySlice[DiagnosticInfo, DiagnosticInfo](m.DiagnosticInfos),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _DataChangeNotificationCopy
}

func (m *_DataChangeNotification) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
