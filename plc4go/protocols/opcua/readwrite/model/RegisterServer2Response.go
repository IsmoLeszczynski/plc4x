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
	ExtensionObjectDefinitionContract
	ResponseHeader           ExtensionObjectDefinition
	NoOfConfigurationResults int32
	ConfigurationResults     []StatusCode
	NoOfDiagnosticInfos      int32
	DiagnosticInfos          []DiagnosticInfo
}

var _ RegisterServer2Response = (*_RegisterServer2Response)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_RegisterServer2Response)(nil)

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

func (m *_RegisterServer2Response) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
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
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ResponseHeader:                    responseHeader,
		NoOfConfigurationResults:          noOfConfigurationResults,
		ConfigurationResults:              configurationResults,
		NoOfDiagnosticInfos:               noOfDiagnosticInfos,
		DiagnosticInfos:                   diagnosticInfos,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
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
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

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

func (m *_RegisterServer2Response) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__registerServer2Response RegisterServer2Response, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RegisterServer2Response"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RegisterServer2Response")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	responseHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("394")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'responseHeader' field"))
	}
	m.ResponseHeader = responseHeader

	noOfConfigurationResults, err := ReadSimpleField(ctx, "noOfConfigurationResults", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfConfigurationResults' field"))
	}
	m.NoOfConfigurationResults = noOfConfigurationResults

	configurationResults, err := ReadCountArrayField[StatusCode](ctx, "configurationResults", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer), uint64(noOfConfigurationResults))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'configurationResults' field"))
	}
	m.ConfigurationResults = configurationResults

	noOfDiagnosticInfos, err := ReadSimpleField(ctx, "noOfDiagnosticInfos", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfDiagnosticInfos' field"))
	}
	m.NoOfDiagnosticInfos = noOfDiagnosticInfos

	diagnosticInfos, err := ReadCountArrayField[DiagnosticInfo](ctx, "diagnosticInfos", ReadComplex[DiagnosticInfo](DiagnosticInfoParseWithBuffer, readBuffer), uint64(noOfDiagnosticInfos))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'diagnosticInfos' field"))
	}
	m.DiagnosticInfos = diagnosticInfos

	if closeErr := readBuffer.CloseContext("RegisterServer2Response"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RegisterServer2Response")
	}

	return m, nil
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

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", m.GetResponseHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'responseHeader' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfConfigurationResults", m.GetNoOfConfigurationResults(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfConfigurationResults' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "configurationResults", m.GetConfigurationResults(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'configurationResults' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfDiagnosticInfos", m.GetNoOfDiagnosticInfos(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfDiagnosticInfos' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "diagnosticInfos", m.GetDiagnosticInfos(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'diagnosticInfos' field")
		}

		if popErr := writeBuffer.PopContext("RegisterServer2Response"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RegisterServer2Response")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
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
