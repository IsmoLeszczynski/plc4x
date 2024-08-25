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

// X509IdentityToken is the corresponding interface of X509IdentityToken
type X509IdentityToken interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	UserIdentityTokenDefinition
	// GetCertificateData returns CertificateData (property field)
	GetCertificateData() PascalByteString
}

// X509IdentityTokenExactly can be used when we want exactly this type and not a type which fulfills X509IdentityToken.
// This is useful for switch cases.
type X509IdentityTokenExactly interface {
	X509IdentityToken
	isX509IdentityToken() bool
}

// _X509IdentityToken is the data-structure of this message
type _X509IdentityToken struct {
	*_UserIdentityTokenDefinition
	CertificateData PascalByteString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_X509IdentityToken) GetIdentifier() string {
	return "certificate"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_X509IdentityToken) InitializeParent(parent UserIdentityTokenDefinition) {}

func (m *_X509IdentityToken) GetParent() UserIdentityTokenDefinition {
	return m._UserIdentityTokenDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_X509IdentityToken) GetCertificateData() PascalByteString {
	return m.CertificateData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewX509IdentityToken factory function for _X509IdentityToken
func NewX509IdentityToken(certificateData PascalByteString) *_X509IdentityToken {
	_result := &_X509IdentityToken{
		CertificateData:              certificateData,
		_UserIdentityTokenDefinition: NewUserIdentityTokenDefinition(),
	}
	_result._UserIdentityTokenDefinition._UserIdentityTokenDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastX509IdentityToken(structType any) X509IdentityToken {
	if casted, ok := structType.(X509IdentityToken); ok {
		return casted
	}
	if casted, ok := structType.(*X509IdentityToken); ok {
		return *casted
	}
	return nil
}

func (m *_X509IdentityToken) GetTypeName() string {
	return "X509IdentityToken"
}

func (m *_X509IdentityToken) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (certificateData)
	lengthInBits += m.CertificateData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_X509IdentityToken) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func X509IdentityTokenParse(ctx context.Context, theBytes []byte, identifier string) (X509IdentityToken, error) {
	return X509IdentityTokenParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func X509IdentityTokenParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (X509IdentityToken, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("X509IdentityToken"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for X509IdentityToken")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (certificateData)
	if pullErr := readBuffer.PullContext("certificateData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for certificateData")
	}
	_certificateData, _certificateDataErr := PascalByteStringParseWithBuffer(ctx, readBuffer)
	if _certificateDataErr != nil {
		return nil, errors.Wrap(_certificateDataErr, "Error parsing 'certificateData' field of X509IdentityToken")
	}
	certificateData := _certificateData.(PascalByteString)
	if closeErr := readBuffer.CloseContext("certificateData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for certificateData")
	}

	if closeErr := readBuffer.CloseContext("X509IdentityToken"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for X509IdentityToken")
	}

	// Create a partially initialized instance
	_child := &_X509IdentityToken{
		_UserIdentityTokenDefinition: &_UserIdentityTokenDefinition{},
		CertificateData:              certificateData,
	}
	_child._UserIdentityTokenDefinition._UserIdentityTokenDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_X509IdentityToken) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_X509IdentityToken) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("X509IdentityToken"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for X509IdentityToken")
		}

		// Simple Field (certificateData)
		if pushErr := writeBuffer.PushContext("certificateData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for certificateData")
		}
		_certificateDataErr := writeBuffer.WriteSerializable(ctx, m.GetCertificateData())
		if popErr := writeBuffer.PopContext("certificateData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for certificateData")
		}
		if _certificateDataErr != nil {
			return errors.Wrap(_certificateDataErr, "Error serializing 'certificateData' field")
		}

		if popErr := writeBuffer.PopContext("X509IdentityToken"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for X509IdentityToken")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_X509IdentityToken) isX509IdentityToken() bool {
	return true
}

func (m *_X509IdentityToken) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
