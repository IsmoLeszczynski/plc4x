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

package tests

import (
	"github.com/apache/plc4x/plc4go/spi/options"
	"testing"

	s7IO "github.com/apache/plc4x/plc4go/protocols/s7/readwrite"
	"github.com/apache/plc4x/plc4go/spi/testutils"
)

func TestS7ParserSerializer(t *testing.T) {
	testutils.RunParserSerializerTestsuite(
		t,
		"assets/testing/protocols/s7/ParserSerializerTestsuite.xml",
		s7IO.S7ParserHelper{},
		testutils.WithSkippedTestCases(
			// TODO: ignored due to carcia changes
			"S7 Read PLC Type Request",
			"S7 Read PLC Type Response",
			"S7 Read Request",
			"S7 Read Response",
			"S7 Read Error Response",
			"S7 Write Request",
			"S7 Write Request",
		),
		options.WithCustomLogger(testutils.ProduceTestingLogger(t)),
	)
}
