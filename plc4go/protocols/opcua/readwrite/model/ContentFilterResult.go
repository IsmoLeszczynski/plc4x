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

// ContentFilterResult is the corresponding interface of ContentFilterResult
type ContentFilterResult interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNoOfElementResults returns NoOfElementResults (property field)
	GetNoOfElementResults() int32
	// GetElementResults returns ElementResults (property field)
	GetElementResults() []ExtensionObjectDefinition
	// GetNoOfElementDiagnosticInfos returns NoOfElementDiagnosticInfos (property field)
	GetNoOfElementDiagnosticInfos() int32
	// GetElementDiagnosticInfos returns ElementDiagnosticInfos (property field)
	GetElementDiagnosticInfos() []DiagnosticInfo
}

// ContentFilterResultExactly can be used when we want exactly this type and not a type which fulfills ContentFilterResult.
// This is useful for switch cases.
type ContentFilterResultExactly interface {
	ContentFilterResult
	isContentFilterResult() bool
}

// _ContentFilterResult is the data-structure of this message
type _ContentFilterResult struct {
	ExtensionObjectDefinitionContract
	NoOfElementResults         int32
	ElementResults             []ExtensionObjectDefinition
	NoOfElementDiagnosticInfos int32
	ElementDiagnosticInfos     []DiagnosticInfo
}

var _ ContentFilterResult = (*_ContentFilterResult)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ContentFilterResult)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ContentFilterResult) GetIdentifier() string {
	return "609"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ContentFilterResult) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ContentFilterResult) GetNoOfElementResults() int32 {
	return m.NoOfElementResults
}

func (m *_ContentFilterResult) GetElementResults() []ExtensionObjectDefinition {
	return m.ElementResults
}

func (m *_ContentFilterResult) GetNoOfElementDiagnosticInfos() int32 {
	return m.NoOfElementDiagnosticInfos
}

func (m *_ContentFilterResult) GetElementDiagnosticInfos() []DiagnosticInfo {
	return m.ElementDiagnosticInfos
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewContentFilterResult factory function for _ContentFilterResult
func NewContentFilterResult(noOfElementResults int32, elementResults []ExtensionObjectDefinition, noOfElementDiagnosticInfos int32, elementDiagnosticInfos []DiagnosticInfo) *_ContentFilterResult {
	_result := &_ContentFilterResult{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		NoOfElementResults:                noOfElementResults,
		ElementResults:                    elementResults,
		NoOfElementDiagnosticInfos:        noOfElementDiagnosticInfos,
		ElementDiagnosticInfos:            elementDiagnosticInfos,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastContentFilterResult(structType any) ContentFilterResult {
	if casted, ok := structType.(ContentFilterResult); ok {
		return casted
	}
	if casted, ok := structType.(*ContentFilterResult); ok {
		return *casted
	}
	return nil
}

func (m *_ContentFilterResult) GetTypeName() string {
	return "ContentFilterResult"
}

func (m *_ContentFilterResult) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (noOfElementResults)
	lengthInBits += 32

	// Array field
	if len(m.ElementResults) > 0 {
		for _curItem, element := range m.ElementResults {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ElementResults), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfElementDiagnosticInfos)
	lengthInBits += 32

	// Array field
	if len(m.ElementDiagnosticInfos) > 0 {
		for _curItem, element := range m.ElementDiagnosticInfos {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ElementDiagnosticInfos), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_ContentFilterResult) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ContentFilterResult) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__contentFilterResult ContentFilterResult, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ContentFilterResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ContentFilterResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	noOfElementResults, err := ReadSimpleField(ctx, "noOfElementResults", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfElementResults' field"))
	}
	m.NoOfElementResults = noOfElementResults

	elementResults, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "elementResults", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("606")), readBuffer), uint64(noOfElementResults))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'elementResults' field"))
	}
	m.ElementResults = elementResults

	noOfElementDiagnosticInfos, err := ReadSimpleField(ctx, "noOfElementDiagnosticInfos", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfElementDiagnosticInfos' field"))
	}
	m.NoOfElementDiagnosticInfos = noOfElementDiagnosticInfos

	elementDiagnosticInfos, err := ReadCountArrayField[DiagnosticInfo](ctx, "elementDiagnosticInfos", ReadComplex[DiagnosticInfo](DiagnosticInfoParseWithBuffer, readBuffer), uint64(noOfElementDiagnosticInfos))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'elementDiagnosticInfos' field"))
	}
	m.ElementDiagnosticInfos = elementDiagnosticInfos

	if closeErr := readBuffer.CloseContext("ContentFilterResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ContentFilterResult")
	}

	return m, nil
}

func (m *_ContentFilterResult) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ContentFilterResult) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ContentFilterResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ContentFilterResult")
		}

		if err := WriteSimpleField[int32](ctx, "noOfElementResults", m.GetNoOfElementResults(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfElementResults' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "elementResults", m.GetElementResults(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'elementResults' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfElementDiagnosticInfos", m.GetNoOfElementDiagnosticInfos(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfElementDiagnosticInfos' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "elementDiagnosticInfos", m.GetElementDiagnosticInfos(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'elementDiagnosticInfos' field")
		}

		if popErr := writeBuffer.PopContext("ContentFilterResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ContentFilterResult")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ContentFilterResult) isContentFilterResult() bool {
	return true
}

func (m *_ContentFilterResult) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
