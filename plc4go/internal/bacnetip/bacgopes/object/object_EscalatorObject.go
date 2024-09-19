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

type EscalatorObject struct {
	Object
	objectType           string // TODO: migrateme
	properties           []Property
	_object_supports_cov bool
}

func NewEscalatorObject(arg Arg) (*EscalatorObject, error) {
	o := &EscalatorObject{
		objectType: "escalator",
		properties: []Property{
			NewReadableProperty("statusFlags", V2P(NewStatusFlags)),
			NewReadableProperty("elevatorGroup", Vs2P(NewObjectIdentifier)),
			NewReadableProperty("groupID", V2P(NewUnsigned8)),
			NewReadableProperty("installationID", V2P(NewUnsigned8)),
			NewOptionalProperty("powerMode", V2P(NewBoolean)),
			NewReadableProperty("operationDirection", V2P(NewEscalatorOperationDirection)),
			NewOptionalProperty("escalatorMode", V2P(NewEscalatorMode)),
			NewOptionalProperty("energyMeter", V2P(NewReal)),
			NewOptionalProperty("energyMeterRef", V2P(NewDeviceObjectReference)),
			NewOptionalProperty("reliability", V2P(NewReliability)),
			NewOptionalProperty("outOfService", V2P(NewBoolean)),
			NewOptionalProperty("faultSignals", ListOfP(NewLiftFault)),
			NewReadableProperty("passengerAlarm", V2P(NewBoolean)),
			NewOptionalProperty("timeDelay", V2P(NewUnsigned)),
			NewOptionalProperty("timeDelayNormal", V2P(NewUnsigned)),
			NewOptionalProperty("eventDetectionEnable", V2P(NewBoolean)),
			NewOptionalProperty("notificationClass", V2P(NewUnsigned)),
			NewOptionalProperty("eventEnable", V2P(NewEventTransitionBits)),
			NewOptionalProperty("eventState", V2P(NewEventState)),
			NewOptionalProperty("ackedTransitions", V2P(NewEventTransitionBits)),
			NewOptionalProperty("notifyType", V2P(NewNotifyType)),
			NewOptionalProperty("eventTimeStamps", ArrayOfP(NewTimeStamp, 3, 0)),
			NewOptionalProperty("eventMessageTexts", ArrayOfP(NewCharacterString, 3, 0)),
			NewOptionalProperty("eventMessageTextsConfig", ArrayOfP(NewCharacterString, 3, 0)),
			NewOptionalProperty("eventAlgorithmInhibit", V2P(NewBoolean)),
			NewOptionalProperty("eventAlgorithmInhibitRef", V2P(NewObjectPropertyReference)),
			NewOptionalProperty("reliabilityEvaluationInhibit", V2P(NewBoolean)),
		},
	}
	// TODO: @register_object_type
	return o, nil
}
