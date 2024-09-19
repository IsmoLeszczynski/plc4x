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

package object

import (
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/basetypes"
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/comp"
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/primitivedata"
)

type AccessUserObject struct {
	Object
	objectType           string // TODO: migrateme
	properties           []Property
	_object_supports_cov bool
}

func NewAccessUserObject(arg Arg) (*AccessUserObject, error) {
	o := &AccessUserObject{
		objectType: "accessUser",
		properties: []Property{
			NewWritableProperty("globalIdentifier", V2P(NewUnsigned)),
			NewReadableProperty("statusFlags", V2P(NewStatusFlags)),
			NewReadableProperty("reliability", V2P(NewReliability)),
			NewReadableProperty("userType", V2P(NewAccessUserType)),
			NewOptionalProperty("userName", V2P(NewCharacterString)),
			NewOptionalProperty("userExternalIdentifier", V2P(NewCharacterString)),
			NewOptionalProperty("userInformationReference", V2P(NewCharacterString)),
			NewOptionalProperty("members", ListOfP(NewDeviceObjectReference)),
			NewOptionalProperty("memberOf", ListOfP(NewDeviceObjectReference)),
			NewReadableProperty("credentials", ListOfP(NewDeviceObjectReference)),
			NewOptionalProperty("reliabilityEvaluationInhibit", V2P(NewBoolean)),
		},
	}
	// TODO: @register_object_type
	return o, nil
}
