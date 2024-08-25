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

// RegisterServer2Response is the corresponding interface of RegisterServer2Response
type RegisterServer2Response interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
	// GetNoOfConfigurationResults returns NoOfConfigurationResults (property field)
	GetNoOfConfigurationResults() int32
	// GetConfigurationResults returns ConfigurationResults (property field)
	GetConfigurationResults() []StatusCode
	// GetNoOfDiagnosticInfos returns NoOfDiagnosticInfos (property field)
	GetNoOfDiagnosticInfos() int32
	// GetDiagnosticInfos returns DiagnosticInfos (property field)
	GetDiagnosticInfos() []DiagnosticInfo
}

// RegisterServer2ResponseExactly can be used when we want exactly this type and not a type which fulfills RegisterServer2Response.
// This is useful for switch cases.
type RegisterServer2ResponseExactly interface {
	RegisterServer2Response
	isRegisterServer2Response() bool
}

// _RegisterServer2Response is the data-structure of this message
type _RegisterServer2Response struct {
	*_ExtensionObjectDefinition
	ResponseHeader           ExtensionObjectDefinition
	NoOfConfigurationResults int32
	ConfigurationResults     []StatusCode
	NoOfDiagnosticInfos      int32
	DiagnosticInfos          []DiagnosticInfo
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_RegisterServer2Response) GetIdentifier() string {
	return "12196"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RegisterServer2Response) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_RegisterServer2Response) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RegisterServer2Response) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

func (m *_RegisterServer2Response) GetNoOfConfigurationResults() int32 {
	return m.NoOfConfigurationResults
}

func (m *_RegisterServer2Response) GetConfigurationResults() []StatusCode {
	return m.ConfigurationResults
}

func (m *_RegisterServer2Response) GetNoOfDiagnosticInfos() int32 {
	return m.NoOfDiagnosticInfos
}

func (m *_RegisterServer2Response) GetDiagnosticInfos() []DiagnosticInfo {
	return m.DiagnosticInfos
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewRegisterServer2Response factory function for _RegisterServer2Response
func NewRegisterServer2Response(responseHeader ExtensionObjectDefinition, noOfConfigurationResults int32, configurationResults []StatusCode, noOfDiagnosticInfos int32, diagnosticInfos []DiagnosticInfo) *_RegisterServer2Response {
	_result := &_RegisterServer2Response{
		ResponseHeader:             responseHeader,
		NoOfConfigurationResults:   noOfConfigurationResults,
		ConfigurationResults:       configurationResults,
		NoOfDiagnosticInfos:        noOfDiagnosticInfos,
		DiagnosticInfos:            diagnosticInfos,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastRegisterServer2Response(structType any) RegisterServer2Response {
	if casted, ok := structType.(RegisterServer2Response); ok {
		return casted
	}
	if casted, ok := structType.(*RegisterServer2Response); ok {
		return *casted
	}
	return nil
}

func (m *_RegisterServer2Response) GetTypeName() string {
	return "RegisterServer2Response"
}

func (m *_RegisterServer2Response) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Simple field (noOfConfigurationResults)
	lengthInBits += 32

	// Array field
	if len(m.ConfigurationResults) > 0 {
		for _curItem, element := range m.ConfigurationResults {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ConfigurationResults), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfDiagnosticInfos)
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

func (m *_RegisterServer2Response) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func RegisterServer2ResponseParse(ctx context.Context, theBytes []byte, identifier string) (RegisterServer2Response, error) {
	return RegisterServer2ResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func RegisterServer2ResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (RegisterServer2Response, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("RegisterServer2Response"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RegisterServer2Response")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (responseHeader)
	if pullErr := readBuffer.PullContext("responseHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for responseHeader")
	}
	_responseHeader, _responseHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("394"))
	if _responseHeaderErr != nil {
		return nil, errors.Wrap(_responseHeaderErr, "Error parsing 'responseHeader' field of RegisterServer2Response")
	}
	responseHeader := _responseHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("responseHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for responseHeader")
	}

	// Simple Field (noOfConfigurationResults)
	_noOfConfigurationResults, _noOfConfigurationResultsErr := readBuffer.ReadInt32("noOfConfigurationResults", 32)
	if _noOfConfigurationResultsErr != nil {
		return nil, errors.Wrap(_noOfConfigurationResultsErr, "Error parsing 'noOfConfigurationResults' field of RegisterServer2Response")
	}
	noOfConfigurationResults := _noOfConfigurationResults

	// Array field (configurationResults)
	if pullErr := readBuffer.PullContext("configurationResults", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for configurationResults")
	}
	// Count array
	configurationResults := make([]StatusCode, max(noOfConfigurationResults, 0))
	// This happens when the size is set conditional to 0
	if len(configurationResults) == 0 {
		configurationResults = nil
	}
	{
		_numItems := uint16(max(noOfConfigurationResults, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := StatusCodeParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'configurationResults' field of RegisterServer2Response")
			}
			configurationResults[_curItem] = _item.(StatusCode)
		}
	}
	if closeErr := readBuffer.CloseContext("configurationResults", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for configurationResults")
	}

	// Simple Field (noOfDiagnosticInfos)
	_noOfDiagnosticInfos, _noOfDiagnosticInfosErr := readBuffer.ReadInt32("noOfDiagnosticInfos", 32)
	if _noOfDiagnosticInfosErr != nil {
		return nil, errors.Wrap(_noOfDiagnosticInfosErr, "Error parsing 'noOfDiagnosticInfos' field of RegisterServer2Response")
	}
	noOfDiagnosticInfos := _noOfDiagnosticInfos

	// Array field (diagnosticInfos)
	if pullErr := readBuffer.PullContext("diagnosticInfos", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for diagnosticInfos")
	}
	// Count array
	diagnosticInfos := make([]DiagnosticInfo, max(noOfDiagnosticInfos, 0))
	// This happens when the size is set conditional to 0
	if len(diagnosticInfos) == 0 {
		diagnosticInfos = nil
	}
	{
		_numItems := uint16(max(noOfDiagnosticInfos, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := DiagnosticInfoParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'diagnosticInfos' field of RegisterServer2Response")
			}
			diagnosticInfos[_curItem] = _item.(DiagnosticInfo)
		}
	}
	if closeErr := readBuffer.CloseContext("diagnosticInfos", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for diagnosticInfos")
	}

	if closeErr := readBuffer.CloseContext("RegisterServer2Response"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RegisterServer2Response")
	}

	// Create a partially initialized instance
	_child := &_RegisterServer2Response{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		ResponseHeader:             responseHeader,
		NoOfConfigurationResults:   noOfConfigurationResults,
		ConfigurationResults:       configurationResults,
		NoOfDiagnosticInfos:        noOfDiagnosticInfos,
		DiagnosticInfos:            diagnosticInfos,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_RegisterServer2Response) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RegisterServer2Response) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RegisterServer2Response"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RegisterServer2Response")
		}

		// Simple Field (responseHeader)
		if pushErr := writeBuffer.PushContext("responseHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for responseHeader")
		}
		_responseHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetResponseHeader())
		if popErr := writeBuffer.PopContext("responseHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for responseHeader")
		}
		if _responseHeaderErr != nil {
			return errors.Wrap(_responseHeaderErr, "Error serializing 'responseHeader' field")
		}

		// Simple Field (noOfConfigurationResults)
		noOfConfigurationResults := int32(m.GetNoOfConfigurationResults())
		_noOfConfigurationResultsErr := writeBuffer.WriteInt32("noOfConfigurationResults", 32, int32((noOfConfigurationResults)))
		if _noOfConfigurationResultsErr != nil {
			return errors.Wrap(_noOfConfigurationResultsErr, "Error serializing 'noOfConfigurationResults' field")
		}

		// Array Field (configurationResults)
		if pushErr := writeBuffer.PushContext("configurationResults", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for configurationResults")
		}
		for _curItem, _element := range m.GetConfigurationResults() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetConfigurationResults()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'configurationResults' field")
			}
		}
		if popErr := writeBuffer.PopContext("configurationResults", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for configurationResults")
		}

		// Simple Field (noOfDiagnosticInfos)
		noOfDiagnosticInfos := int32(m.GetNoOfDiagnosticInfos())
		_noOfDiagnosticInfosErr := writeBuffer.WriteInt32("noOfDiagnosticInfos", 32, int32((noOfDiagnosticInfos)))
		if _noOfDiagnosticInfosErr != nil {
			return errors.Wrap(_noOfDiagnosticInfosErr, "Error serializing 'noOfDiagnosticInfos' field")
		}

		// Array Field (diagnosticInfos)
		if pushErr := writeBuffer.PushContext("diagnosticInfos", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for diagnosticInfos")
		}
		for _curItem, _element := range m.GetDiagnosticInfos() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetDiagnosticInfos()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'diagnosticInfos' field")
			}
		}
		if popErr := writeBuffer.PopContext("diagnosticInfos", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for diagnosticInfos")
		}

		if popErr := writeBuffer.PopContext("RegisterServer2Response"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RegisterServer2Response")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RegisterServer2Response) isRegisterServer2Response() bool {
	return true
}

func (m *_RegisterServer2Response) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
