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

// ProgramDiagnosticDataType is the corresponding interface of ProgramDiagnosticDataType
type ProgramDiagnosticDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetCreateSessionId returns CreateSessionId (property field)
	GetCreateSessionId() NodeId
	// GetCreateClientName returns CreateClientName (property field)
	GetCreateClientName() PascalString
	// GetInvocationCreationTime returns InvocationCreationTime (property field)
	GetInvocationCreationTime() int64
	// GetLastTransitionTime returns LastTransitionTime (property field)
	GetLastTransitionTime() int64
	// GetLastMethodCall returns LastMethodCall (property field)
	GetLastMethodCall() PascalString
	// GetLastMethodSessionId returns LastMethodSessionId (property field)
	GetLastMethodSessionId() NodeId
	// GetNoOfLastMethodInputArguments returns NoOfLastMethodInputArguments (property field)
	GetNoOfLastMethodInputArguments() int32
	// GetLastMethodInputArguments returns LastMethodInputArguments (property field)
	GetLastMethodInputArguments() []ExtensionObjectDefinition
	// GetNoOfLastMethodOutputArguments returns NoOfLastMethodOutputArguments (property field)
	GetNoOfLastMethodOutputArguments() int32
	// GetLastMethodOutputArguments returns LastMethodOutputArguments (property field)
	GetLastMethodOutputArguments() []ExtensionObjectDefinition
	// GetLastMethodCallTime returns LastMethodCallTime (property field)
	GetLastMethodCallTime() int64
	// GetLastMethodReturnStatus returns LastMethodReturnStatus (property field)
	GetLastMethodReturnStatus() ExtensionObjectDefinition
}

// ProgramDiagnosticDataTypeExactly can be used when we want exactly this type and not a type which fulfills ProgramDiagnosticDataType.
// This is useful for switch cases.
type ProgramDiagnosticDataTypeExactly interface {
	ProgramDiagnosticDataType
	isProgramDiagnosticDataType() bool
}

// _ProgramDiagnosticDataType is the data-structure of this message
type _ProgramDiagnosticDataType struct {
	ExtensionObjectDefinitionContract
	CreateSessionId               NodeId
	CreateClientName              PascalString
	InvocationCreationTime        int64
	LastTransitionTime            int64
	LastMethodCall                PascalString
	LastMethodSessionId           NodeId
	NoOfLastMethodInputArguments  int32
	LastMethodInputArguments      []ExtensionObjectDefinition
	NoOfLastMethodOutputArguments int32
	LastMethodOutputArguments     []ExtensionObjectDefinition
	LastMethodCallTime            int64
	LastMethodReturnStatus        ExtensionObjectDefinition
}

var _ ProgramDiagnosticDataType = (*_ProgramDiagnosticDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ProgramDiagnosticDataType)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ProgramDiagnosticDataType) GetIdentifier() string {
	return "896"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ProgramDiagnosticDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ProgramDiagnosticDataType) GetCreateSessionId() NodeId {
	return m.CreateSessionId
}

func (m *_ProgramDiagnosticDataType) GetCreateClientName() PascalString {
	return m.CreateClientName
}

func (m *_ProgramDiagnosticDataType) GetInvocationCreationTime() int64 {
	return m.InvocationCreationTime
}

func (m *_ProgramDiagnosticDataType) GetLastTransitionTime() int64 {
	return m.LastTransitionTime
}

func (m *_ProgramDiagnosticDataType) GetLastMethodCall() PascalString {
	return m.LastMethodCall
}

func (m *_ProgramDiagnosticDataType) GetLastMethodSessionId() NodeId {
	return m.LastMethodSessionId
}

func (m *_ProgramDiagnosticDataType) GetNoOfLastMethodInputArguments() int32 {
	return m.NoOfLastMethodInputArguments
}

func (m *_ProgramDiagnosticDataType) GetLastMethodInputArguments() []ExtensionObjectDefinition {
	return m.LastMethodInputArguments
}

func (m *_ProgramDiagnosticDataType) GetNoOfLastMethodOutputArguments() int32 {
	return m.NoOfLastMethodOutputArguments
}

func (m *_ProgramDiagnosticDataType) GetLastMethodOutputArguments() []ExtensionObjectDefinition {
	return m.LastMethodOutputArguments
}

func (m *_ProgramDiagnosticDataType) GetLastMethodCallTime() int64 {
	return m.LastMethodCallTime
}

func (m *_ProgramDiagnosticDataType) GetLastMethodReturnStatus() ExtensionObjectDefinition {
	return m.LastMethodReturnStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewProgramDiagnosticDataType factory function for _ProgramDiagnosticDataType
func NewProgramDiagnosticDataType(createSessionId NodeId, createClientName PascalString, invocationCreationTime int64, lastTransitionTime int64, lastMethodCall PascalString, lastMethodSessionId NodeId, noOfLastMethodInputArguments int32, lastMethodInputArguments []ExtensionObjectDefinition, noOfLastMethodOutputArguments int32, lastMethodOutputArguments []ExtensionObjectDefinition, lastMethodCallTime int64, lastMethodReturnStatus ExtensionObjectDefinition) *_ProgramDiagnosticDataType {
	_result := &_ProgramDiagnosticDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		CreateSessionId:                   createSessionId,
		CreateClientName:                  createClientName,
		InvocationCreationTime:            invocationCreationTime,
		LastTransitionTime:                lastTransitionTime,
		LastMethodCall:                    lastMethodCall,
		LastMethodSessionId:               lastMethodSessionId,
		NoOfLastMethodInputArguments:      noOfLastMethodInputArguments,
		LastMethodInputArguments:          lastMethodInputArguments,
		NoOfLastMethodOutputArguments:     noOfLastMethodOutputArguments,
		LastMethodOutputArguments:         lastMethodOutputArguments,
		LastMethodCallTime:                lastMethodCallTime,
		LastMethodReturnStatus:            lastMethodReturnStatus,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastProgramDiagnosticDataType(structType any) ProgramDiagnosticDataType {
	if casted, ok := structType.(ProgramDiagnosticDataType); ok {
		return casted
	}
	if casted, ok := structType.(*ProgramDiagnosticDataType); ok {
		return *casted
	}
	return nil
}

func (m *_ProgramDiagnosticDataType) GetTypeName() string {
	return "ProgramDiagnosticDataType"
}

func (m *_ProgramDiagnosticDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (createSessionId)
	lengthInBits += m.CreateSessionId.GetLengthInBits(ctx)

	// Simple field (createClientName)
	lengthInBits += m.CreateClientName.GetLengthInBits(ctx)

	// Simple field (invocationCreationTime)
	lengthInBits += 64

	// Simple field (lastTransitionTime)
	lengthInBits += 64

	// Simple field (lastMethodCall)
	lengthInBits += m.LastMethodCall.GetLengthInBits(ctx)

	// Simple field (lastMethodSessionId)
	lengthInBits += m.LastMethodSessionId.GetLengthInBits(ctx)

	// Simple field (noOfLastMethodInputArguments)
	lengthInBits += 32

	// Array field
	if len(m.LastMethodInputArguments) > 0 {
		for _curItem, element := range m.LastMethodInputArguments {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.LastMethodInputArguments), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfLastMethodOutputArguments)
	lengthInBits += 32

	// Array field
	if len(m.LastMethodOutputArguments) > 0 {
		for _curItem, element := range m.LastMethodOutputArguments {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.LastMethodOutputArguments), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (lastMethodCallTime)
	lengthInBits += 64

	// Simple field (lastMethodReturnStatus)
	lengthInBits += m.LastMethodReturnStatus.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ProgramDiagnosticDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ProgramDiagnosticDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__programDiagnosticDataType ProgramDiagnosticDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ProgramDiagnosticDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ProgramDiagnosticDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	createSessionId, err := ReadSimpleField[NodeId](ctx, "createSessionId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'createSessionId' field"))
	}
	m.CreateSessionId = createSessionId

	createClientName, err := ReadSimpleField[PascalString](ctx, "createClientName", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'createClientName' field"))
	}
	m.CreateClientName = createClientName

	invocationCreationTime, err := ReadSimpleField(ctx, "invocationCreationTime", ReadSignedLong(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'invocationCreationTime' field"))
	}
	m.InvocationCreationTime = invocationCreationTime

	lastTransitionTime, err := ReadSimpleField(ctx, "lastTransitionTime", ReadSignedLong(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastTransitionTime' field"))
	}
	m.LastTransitionTime = lastTransitionTime

	lastMethodCall, err := ReadSimpleField[PascalString](ctx, "lastMethodCall", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastMethodCall' field"))
	}
	m.LastMethodCall = lastMethodCall

	lastMethodSessionId, err := ReadSimpleField[NodeId](ctx, "lastMethodSessionId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastMethodSessionId' field"))
	}
	m.LastMethodSessionId = lastMethodSessionId

	noOfLastMethodInputArguments, err := ReadSimpleField(ctx, "noOfLastMethodInputArguments", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfLastMethodInputArguments' field"))
	}
	m.NoOfLastMethodInputArguments = noOfLastMethodInputArguments

	lastMethodInputArguments, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "lastMethodInputArguments", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("298")), readBuffer), uint64(noOfLastMethodInputArguments))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastMethodInputArguments' field"))
	}
	m.LastMethodInputArguments = lastMethodInputArguments

	noOfLastMethodOutputArguments, err := ReadSimpleField(ctx, "noOfLastMethodOutputArguments", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfLastMethodOutputArguments' field"))
	}
	m.NoOfLastMethodOutputArguments = noOfLastMethodOutputArguments

	lastMethodOutputArguments, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "lastMethodOutputArguments", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("298")), readBuffer), uint64(noOfLastMethodOutputArguments))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastMethodOutputArguments' field"))
	}
	m.LastMethodOutputArguments = lastMethodOutputArguments

	lastMethodCallTime, err := ReadSimpleField(ctx, "lastMethodCallTime", ReadSignedLong(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastMethodCallTime' field"))
	}
	m.LastMethodCallTime = lastMethodCallTime

	lastMethodReturnStatus, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "lastMethodReturnStatus", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("301")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastMethodReturnStatus' field"))
	}
	m.LastMethodReturnStatus = lastMethodReturnStatus

	if closeErr := readBuffer.CloseContext("ProgramDiagnosticDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ProgramDiagnosticDataType")
	}

	return m, nil
}

func (m *_ProgramDiagnosticDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ProgramDiagnosticDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ProgramDiagnosticDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ProgramDiagnosticDataType")
		}

		if err := WriteSimpleField[NodeId](ctx, "createSessionId", m.GetCreateSessionId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'createSessionId' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "createClientName", m.GetCreateClientName(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'createClientName' field")
		}

		if err := WriteSimpleField[int64](ctx, "invocationCreationTime", m.GetInvocationCreationTime(), WriteSignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'invocationCreationTime' field")
		}

		if err := WriteSimpleField[int64](ctx, "lastTransitionTime", m.GetLastTransitionTime(), WriteSignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastTransitionTime' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "lastMethodCall", m.GetLastMethodCall(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastMethodCall' field")
		}

		if err := WriteSimpleField[NodeId](ctx, "lastMethodSessionId", m.GetLastMethodSessionId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastMethodSessionId' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfLastMethodInputArguments", m.GetNoOfLastMethodInputArguments(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfLastMethodInputArguments' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "lastMethodInputArguments", m.GetLastMethodInputArguments(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'lastMethodInputArguments' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfLastMethodOutputArguments", m.GetNoOfLastMethodOutputArguments(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfLastMethodOutputArguments' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "lastMethodOutputArguments", m.GetLastMethodOutputArguments(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'lastMethodOutputArguments' field")
		}

		if err := WriteSimpleField[int64](ctx, "lastMethodCallTime", m.GetLastMethodCallTime(), WriteSignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastMethodCallTime' field")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "lastMethodReturnStatus", m.GetLastMethodReturnStatus(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastMethodReturnStatus' field")
		}

		if popErr := writeBuffer.PopContext("ProgramDiagnosticDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ProgramDiagnosticDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ProgramDiagnosticDataType) isProgramDiagnosticDataType() bool {
	return true
}

func (m *_ProgramDiagnosticDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
