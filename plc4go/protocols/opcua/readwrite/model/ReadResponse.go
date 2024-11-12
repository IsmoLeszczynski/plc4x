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

// ReadResponse is the corresponding interface of ReadResponse
type ReadResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ResponseHeader
	// GetResults returns Results (property field)
	GetResults() []DataValue
	// GetDiagnosticInfos returns DiagnosticInfos (property field)
	GetDiagnosticInfos() []DiagnosticInfo
	// IsReadResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsReadResponse()
	// CreateBuilder creates a ReadResponseBuilder
	CreateReadResponseBuilder() ReadResponseBuilder
}

// _ReadResponse is the data-structure of this message
type _ReadResponse struct {
	ExtensionObjectDefinitionContract
	ResponseHeader  ResponseHeader
	Results         []DataValue
	DiagnosticInfos []DiagnosticInfo
}

var _ ReadResponse = (*_ReadResponse)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ReadResponse)(nil)

// NewReadResponse factory function for _ReadResponse
func NewReadResponse(responseHeader ResponseHeader, results []DataValue, diagnosticInfos []DiagnosticInfo) *_ReadResponse {
	if responseHeader == nil {
		panic("responseHeader of type ResponseHeader for ReadResponse must not be nil")
	}
	_result := &_ReadResponse{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ResponseHeader:                    responseHeader,
		Results:                           results,
		DiagnosticInfos:                   diagnosticInfos,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ReadResponseBuilder is a builder for ReadResponse
type ReadResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(responseHeader ResponseHeader, results []DataValue, diagnosticInfos []DiagnosticInfo) ReadResponseBuilder
	// WithResponseHeader adds ResponseHeader (property field)
	WithResponseHeader(ResponseHeader) ReadResponseBuilder
	// WithResponseHeaderBuilder adds ResponseHeader (property field) which is build by the builder
	WithResponseHeaderBuilder(func(ResponseHeaderBuilder) ResponseHeaderBuilder) ReadResponseBuilder
	// WithResults adds Results (property field)
	WithResults(...DataValue) ReadResponseBuilder
	// WithDiagnosticInfos adds DiagnosticInfos (property field)
	WithDiagnosticInfos(...DiagnosticInfo) ReadResponseBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the ReadResponse or returns an error if something is wrong
	Build() (ReadResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ReadResponse
}

// NewReadResponseBuilder() creates a ReadResponseBuilder
func NewReadResponseBuilder() ReadResponseBuilder {
	return &_ReadResponseBuilder{_ReadResponse: new(_ReadResponse)}
}

type _ReadResponseBuilder struct {
	*_ReadResponse

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (ReadResponseBuilder) = (*_ReadResponseBuilder)(nil)

func (b *_ReadResponseBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_ReadResponseBuilder) WithMandatoryFields(responseHeader ResponseHeader, results []DataValue, diagnosticInfos []DiagnosticInfo) ReadResponseBuilder {
	return b.WithResponseHeader(responseHeader).WithResults(results...).WithDiagnosticInfos(diagnosticInfos...)
}

func (b *_ReadResponseBuilder) WithResponseHeader(responseHeader ResponseHeader) ReadResponseBuilder {
	b.ResponseHeader = responseHeader
	return b
}

func (b *_ReadResponseBuilder) WithResponseHeaderBuilder(builderSupplier func(ResponseHeaderBuilder) ResponseHeaderBuilder) ReadResponseBuilder {
	builder := builderSupplier(b.ResponseHeader.CreateResponseHeaderBuilder())
	var err error
	b.ResponseHeader, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ResponseHeaderBuilder failed"))
	}
	return b
}

func (b *_ReadResponseBuilder) WithResults(results ...DataValue) ReadResponseBuilder {
	b.Results = results
	return b
}

func (b *_ReadResponseBuilder) WithDiagnosticInfos(diagnosticInfos ...DiagnosticInfo) ReadResponseBuilder {
	b.DiagnosticInfos = diagnosticInfos
	return b
}

func (b *_ReadResponseBuilder) Build() (ReadResponse, error) {
	if b.ResponseHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'responseHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ReadResponse.deepCopy(), nil
}

func (b *_ReadResponseBuilder) MustBuild() ReadResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ReadResponseBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_ReadResponseBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_ReadResponseBuilder) DeepCopy() any {
	_copy := b.CreateReadResponseBuilder().(*_ReadResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateReadResponseBuilder creates a ReadResponseBuilder
func (b *_ReadResponse) CreateReadResponseBuilder() ReadResponseBuilder {
	if b == nil {
		return NewReadResponseBuilder()
	}
	return &_ReadResponseBuilder{_ReadResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ReadResponse) GetExtensionId() int32 {
	return int32(634)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ReadResponse) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ReadResponse) GetResponseHeader() ResponseHeader {
	return m.ResponseHeader
}

func (m *_ReadResponse) GetResults() []DataValue {
	return m.Results
}

func (m *_ReadResponse) GetDiagnosticInfos() []DiagnosticInfo {
	return m.DiagnosticInfos
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastReadResponse(structType any) ReadResponse {
	if casted, ok := structType.(ReadResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ReadResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ReadResponse) GetTypeName() string {
	return "ReadResponse"
}

func (m *_ReadResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Implicit Field (noOfResults)
	lengthInBits += 32

	// Array field
	if len(m.Results) > 0 {
		for _curItem, element := range m.Results {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Results), _curItem)
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

func (m *_ReadResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ReadResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__readResponse ReadResponse, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ReadResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ReadResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	responseHeader, err := ReadSimpleField[ResponseHeader](ctx, "responseHeader", ReadComplex[ResponseHeader](ExtensionObjectDefinitionParseWithBufferProducer[ResponseHeader]((int32)(int32(394))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'responseHeader' field"))
	}
	m.ResponseHeader = responseHeader

	noOfResults, err := ReadImplicitField[int32](ctx, "noOfResults", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfResults' field"))
	}
	_ = noOfResults

	results, err := ReadCountArrayField[DataValue](ctx, "results", ReadComplex[DataValue](DataValueParseWithBuffer, readBuffer), uint64(noOfResults))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'results' field"))
	}
	m.Results = results

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

	if closeErr := readBuffer.CloseContext("ReadResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ReadResponse")
	}

	return m, nil
}

func (m *_ReadResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ReadResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ReadResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ReadResponse")
		}

		if err := WriteSimpleField[ResponseHeader](ctx, "responseHeader", m.GetResponseHeader(), WriteComplex[ResponseHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'responseHeader' field")
		}
		noOfResults := int32(utils.InlineIf(bool((m.GetResults()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetResults()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfResults", noOfResults, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfResults' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "results", m.GetResults(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'results' field")
		}
		noOfDiagnosticInfos := int32(utils.InlineIf(bool((m.GetDiagnosticInfos()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetDiagnosticInfos()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfDiagnosticInfos", noOfDiagnosticInfos, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfDiagnosticInfos' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "diagnosticInfos", m.GetDiagnosticInfos(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'diagnosticInfos' field")
		}

		if popErr := writeBuffer.PopContext("ReadResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ReadResponse")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ReadResponse) IsReadResponse() {}

func (m *_ReadResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ReadResponse) deepCopy() *_ReadResponse {
	if m == nil {
		return nil
	}
	_ReadResponseCopy := &_ReadResponse{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[ResponseHeader](m.ResponseHeader),
		utils.DeepCopySlice[DataValue, DataValue](m.Results),
		utils.DeepCopySlice[DiagnosticInfo, DiagnosticInfo](m.DiagnosticInfos),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _ReadResponseCopy
}

func (m *_ReadResponse) String() string {
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
