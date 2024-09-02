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

// AnnotationDataType is the corresponding interface of AnnotationDataType
type AnnotationDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetAnnotation returns Annotation (property field)
	GetAnnotation() PascalString
	// GetDiscipline returns Discipline (property field)
	GetDiscipline() PascalString
	// GetUri returns Uri (property field)
	GetUri() PascalString
}

// AnnotationDataTypeExactly can be used when we want exactly this type and not a type which fulfills AnnotationDataType.
// This is useful for switch cases.
type AnnotationDataTypeExactly interface {
	AnnotationDataType
	isAnnotationDataType() bool
}

// _AnnotationDataType is the data-structure of this message
type _AnnotationDataType struct {
	ExtensionObjectDefinitionContract
	Annotation PascalString
	Discipline PascalString
	Uri        PascalString
}

var _ AnnotationDataType = (*_AnnotationDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_AnnotationDataType)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AnnotationDataType) GetIdentifier() string {
	return "32436"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AnnotationDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AnnotationDataType) GetAnnotation() PascalString {
	return m.Annotation
}

func (m *_AnnotationDataType) GetDiscipline() PascalString {
	return m.Discipline
}

func (m *_AnnotationDataType) GetUri() PascalString {
	return m.Uri
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAnnotationDataType factory function for _AnnotationDataType
func NewAnnotationDataType(annotation PascalString, discipline PascalString, uri PascalString) *_AnnotationDataType {
	_result := &_AnnotationDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Annotation:                        annotation,
		Discipline:                        discipline,
		Uri:                               uri,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAnnotationDataType(structType any) AnnotationDataType {
	if casted, ok := structType.(AnnotationDataType); ok {
		return casted
	}
	if casted, ok := structType.(*AnnotationDataType); ok {
		return *casted
	}
	return nil
}

func (m *_AnnotationDataType) GetTypeName() string {
	return "AnnotationDataType"
}

func (m *_AnnotationDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (annotation)
	lengthInBits += m.Annotation.GetLengthInBits(ctx)

	// Simple field (discipline)
	lengthInBits += m.Discipline.GetLengthInBits(ctx)

	// Simple field (uri)
	lengthInBits += m.Uri.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_AnnotationDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AnnotationDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__annotationDataType AnnotationDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AnnotationDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AnnotationDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	annotation, err := ReadSimpleField[PascalString](ctx, "annotation", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'annotation' field"))
	}
	m.Annotation = annotation

	discipline, err := ReadSimpleField[PascalString](ctx, "discipline", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'discipline' field"))
	}
	m.Discipline = discipline

	uri, err := ReadSimpleField[PascalString](ctx, "uri", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'uri' field"))
	}
	m.Uri = uri

	if closeErr := readBuffer.CloseContext("AnnotationDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AnnotationDataType")
	}

	return m, nil
}

func (m *_AnnotationDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AnnotationDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AnnotationDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AnnotationDataType")
		}

		if err := WriteSimpleField[PascalString](ctx, "annotation", m.GetAnnotation(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'annotation' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "discipline", m.GetDiscipline(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'discipline' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "uri", m.GetUri(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'uri' field")
		}

		if popErr := writeBuffer.PopContext("AnnotationDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AnnotationDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AnnotationDataType) isAnnotationDataType() bool {
	return true
}

func (m *_AnnotationDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
