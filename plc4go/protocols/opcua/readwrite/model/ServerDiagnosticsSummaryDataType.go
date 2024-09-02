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

// ServerDiagnosticsSummaryDataType is the corresponding interface of ServerDiagnosticsSummaryDataType
type ServerDiagnosticsSummaryDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetServerViewCount returns ServerViewCount (property field)
	GetServerViewCount() uint32
	// GetCurrentSessionCount returns CurrentSessionCount (property field)
	GetCurrentSessionCount() uint32
	// GetCumulatedSessionCount returns CumulatedSessionCount (property field)
	GetCumulatedSessionCount() uint32
	// GetSecurityRejectedSessionCount returns SecurityRejectedSessionCount (property field)
	GetSecurityRejectedSessionCount() uint32
	// GetRejectedSessionCount returns RejectedSessionCount (property field)
	GetRejectedSessionCount() uint32
	// GetSessionTimeoutCount returns SessionTimeoutCount (property field)
	GetSessionTimeoutCount() uint32
	// GetSessionAbortCount returns SessionAbortCount (property field)
	GetSessionAbortCount() uint32
	// GetCurrentSubscriptionCount returns CurrentSubscriptionCount (property field)
	GetCurrentSubscriptionCount() uint32
	// GetCumulatedSubscriptionCount returns CumulatedSubscriptionCount (property field)
	GetCumulatedSubscriptionCount() uint32
	// GetPublishingIntervalCount returns PublishingIntervalCount (property field)
	GetPublishingIntervalCount() uint32
	// GetSecurityRejectedRequestsCount returns SecurityRejectedRequestsCount (property field)
	GetSecurityRejectedRequestsCount() uint32
	// GetRejectedRequestsCount returns RejectedRequestsCount (property field)
	GetRejectedRequestsCount() uint32
}

// _ServerDiagnosticsSummaryDataType is the data-structure of this message
type _ServerDiagnosticsSummaryDataType struct {
	ExtensionObjectDefinitionContract
	ServerViewCount               uint32
	CurrentSessionCount           uint32
	CumulatedSessionCount         uint32
	SecurityRejectedSessionCount  uint32
	RejectedSessionCount          uint32
	SessionTimeoutCount           uint32
	SessionAbortCount             uint32
	CurrentSubscriptionCount      uint32
	CumulatedSubscriptionCount    uint32
	PublishingIntervalCount       uint32
	SecurityRejectedRequestsCount uint32
	RejectedRequestsCount         uint32
}

var _ ServerDiagnosticsSummaryDataType = (*_ServerDiagnosticsSummaryDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ServerDiagnosticsSummaryDataType)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ServerDiagnosticsSummaryDataType) GetIdentifier() string {
	return "861"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ServerDiagnosticsSummaryDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ServerDiagnosticsSummaryDataType) GetServerViewCount() uint32 {
	return m.ServerViewCount
}

func (m *_ServerDiagnosticsSummaryDataType) GetCurrentSessionCount() uint32 {
	return m.CurrentSessionCount
}

func (m *_ServerDiagnosticsSummaryDataType) GetCumulatedSessionCount() uint32 {
	return m.CumulatedSessionCount
}

func (m *_ServerDiagnosticsSummaryDataType) GetSecurityRejectedSessionCount() uint32 {
	return m.SecurityRejectedSessionCount
}

func (m *_ServerDiagnosticsSummaryDataType) GetRejectedSessionCount() uint32 {
	return m.RejectedSessionCount
}

func (m *_ServerDiagnosticsSummaryDataType) GetSessionTimeoutCount() uint32 {
	return m.SessionTimeoutCount
}

func (m *_ServerDiagnosticsSummaryDataType) GetSessionAbortCount() uint32 {
	return m.SessionAbortCount
}

func (m *_ServerDiagnosticsSummaryDataType) GetCurrentSubscriptionCount() uint32 {
	return m.CurrentSubscriptionCount
}

func (m *_ServerDiagnosticsSummaryDataType) GetCumulatedSubscriptionCount() uint32 {
	return m.CumulatedSubscriptionCount
}

func (m *_ServerDiagnosticsSummaryDataType) GetPublishingIntervalCount() uint32 {
	return m.PublishingIntervalCount
}

func (m *_ServerDiagnosticsSummaryDataType) GetSecurityRejectedRequestsCount() uint32 {
	return m.SecurityRejectedRequestsCount
}

func (m *_ServerDiagnosticsSummaryDataType) GetRejectedRequestsCount() uint32 {
	return m.RejectedRequestsCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewServerDiagnosticsSummaryDataType factory function for _ServerDiagnosticsSummaryDataType
func NewServerDiagnosticsSummaryDataType(serverViewCount uint32, currentSessionCount uint32, cumulatedSessionCount uint32, securityRejectedSessionCount uint32, rejectedSessionCount uint32, sessionTimeoutCount uint32, sessionAbortCount uint32, currentSubscriptionCount uint32, cumulatedSubscriptionCount uint32, publishingIntervalCount uint32, securityRejectedRequestsCount uint32, rejectedRequestsCount uint32) *_ServerDiagnosticsSummaryDataType {
	_result := &_ServerDiagnosticsSummaryDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ServerViewCount:                   serverViewCount,
		CurrentSessionCount:               currentSessionCount,
		CumulatedSessionCount:             cumulatedSessionCount,
		SecurityRejectedSessionCount:      securityRejectedSessionCount,
		RejectedSessionCount:              rejectedSessionCount,
		SessionTimeoutCount:               sessionTimeoutCount,
		SessionAbortCount:                 sessionAbortCount,
		CurrentSubscriptionCount:          currentSubscriptionCount,
		CumulatedSubscriptionCount:        cumulatedSubscriptionCount,
		PublishingIntervalCount:           publishingIntervalCount,
		SecurityRejectedRequestsCount:     securityRejectedRequestsCount,
		RejectedRequestsCount:             rejectedRequestsCount,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastServerDiagnosticsSummaryDataType(structType any) ServerDiagnosticsSummaryDataType {
	if casted, ok := structType.(ServerDiagnosticsSummaryDataType); ok {
		return casted
	}
	if casted, ok := structType.(*ServerDiagnosticsSummaryDataType); ok {
		return *casted
	}
	return nil
}

func (m *_ServerDiagnosticsSummaryDataType) GetTypeName() string {
	return "ServerDiagnosticsSummaryDataType"
}

func (m *_ServerDiagnosticsSummaryDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (serverViewCount)
	lengthInBits += 32

	// Simple field (currentSessionCount)
	lengthInBits += 32

	// Simple field (cumulatedSessionCount)
	lengthInBits += 32

	// Simple field (securityRejectedSessionCount)
	lengthInBits += 32

	// Simple field (rejectedSessionCount)
	lengthInBits += 32

	// Simple field (sessionTimeoutCount)
	lengthInBits += 32

	// Simple field (sessionAbortCount)
	lengthInBits += 32

	// Simple field (currentSubscriptionCount)
	lengthInBits += 32

	// Simple field (cumulatedSubscriptionCount)
	lengthInBits += 32

	// Simple field (publishingIntervalCount)
	lengthInBits += 32

	// Simple field (securityRejectedRequestsCount)
	lengthInBits += 32

	// Simple field (rejectedRequestsCount)
	lengthInBits += 32

	return lengthInBits
}

func (m *_ServerDiagnosticsSummaryDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ServerDiagnosticsSummaryDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__serverDiagnosticsSummaryDataType ServerDiagnosticsSummaryDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ServerDiagnosticsSummaryDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ServerDiagnosticsSummaryDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	serverViewCount, err := ReadSimpleField(ctx, "serverViewCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serverViewCount' field"))
	}
	m.ServerViewCount = serverViewCount

	currentSessionCount, err := ReadSimpleField(ctx, "currentSessionCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'currentSessionCount' field"))
	}
	m.CurrentSessionCount = currentSessionCount

	cumulatedSessionCount, err := ReadSimpleField(ctx, "cumulatedSessionCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'cumulatedSessionCount' field"))
	}
	m.CumulatedSessionCount = cumulatedSessionCount

	securityRejectedSessionCount, err := ReadSimpleField(ctx, "securityRejectedSessionCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityRejectedSessionCount' field"))
	}
	m.SecurityRejectedSessionCount = securityRejectedSessionCount

	rejectedSessionCount, err := ReadSimpleField(ctx, "rejectedSessionCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'rejectedSessionCount' field"))
	}
	m.RejectedSessionCount = rejectedSessionCount

	sessionTimeoutCount, err := ReadSimpleField(ctx, "sessionTimeoutCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sessionTimeoutCount' field"))
	}
	m.SessionTimeoutCount = sessionTimeoutCount

	sessionAbortCount, err := ReadSimpleField(ctx, "sessionAbortCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sessionAbortCount' field"))
	}
	m.SessionAbortCount = sessionAbortCount

	currentSubscriptionCount, err := ReadSimpleField(ctx, "currentSubscriptionCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'currentSubscriptionCount' field"))
	}
	m.CurrentSubscriptionCount = currentSubscriptionCount

	cumulatedSubscriptionCount, err := ReadSimpleField(ctx, "cumulatedSubscriptionCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'cumulatedSubscriptionCount' field"))
	}
	m.CumulatedSubscriptionCount = cumulatedSubscriptionCount

	publishingIntervalCount, err := ReadSimpleField(ctx, "publishingIntervalCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'publishingIntervalCount' field"))
	}
	m.PublishingIntervalCount = publishingIntervalCount

	securityRejectedRequestsCount, err := ReadSimpleField(ctx, "securityRejectedRequestsCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityRejectedRequestsCount' field"))
	}
	m.SecurityRejectedRequestsCount = securityRejectedRequestsCount

	rejectedRequestsCount, err := ReadSimpleField(ctx, "rejectedRequestsCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'rejectedRequestsCount' field"))
	}
	m.RejectedRequestsCount = rejectedRequestsCount

	if closeErr := readBuffer.CloseContext("ServerDiagnosticsSummaryDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ServerDiagnosticsSummaryDataType")
	}

	return m, nil
}

func (m *_ServerDiagnosticsSummaryDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ServerDiagnosticsSummaryDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ServerDiagnosticsSummaryDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ServerDiagnosticsSummaryDataType")
		}

		if err := WriteSimpleField[uint32](ctx, "serverViewCount", m.GetServerViewCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'serverViewCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "currentSessionCount", m.GetCurrentSessionCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'currentSessionCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "cumulatedSessionCount", m.GetCumulatedSessionCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'cumulatedSessionCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "securityRejectedSessionCount", m.GetSecurityRejectedSessionCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'securityRejectedSessionCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "rejectedSessionCount", m.GetRejectedSessionCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'rejectedSessionCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "sessionTimeoutCount", m.GetSessionTimeoutCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'sessionTimeoutCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "sessionAbortCount", m.GetSessionAbortCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'sessionAbortCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "currentSubscriptionCount", m.GetCurrentSubscriptionCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'currentSubscriptionCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "cumulatedSubscriptionCount", m.GetCumulatedSubscriptionCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'cumulatedSubscriptionCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "publishingIntervalCount", m.GetPublishingIntervalCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'publishingIntervalCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "securityRejectedRequestsCount", m.GetSecurityRejectedRequestsCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'securityRejectedRequestsCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "rejectedRequestsCount", m.GetRejectedRequestsCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'rejectedRequestsCount' field")
		}

		if popErr := writeBuffer.PopContext("ServerDiagnosticsSummaryDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ServerDiagnosticsSummaryDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ServerDiagnosticsSummaryDataType) IsServerDiagnosticsSummaryDataType() {}

func (m *_ServerDiagnosticsSummaryDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
