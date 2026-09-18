/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.java.spi.drivers;

import org.apache.plc4x.java.spi.config.Configuration;
import org.apache.plc4x.java.spi.drivers.tags.PlcTagHandler;
import org.apache.plc4x.java.spi.values.PlcValueHandler;
import org.apache.plc4x.java.utils.auditlog.api.AuditLog;
import org.apache.plc4x.java.utils.auditlog.api.AuditLogEventType;
import org.apache.plc4x.java.utils.auditlog.api.config.AuditLogConfiguration;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class ConnectionBaseCloseTest {

    /** Disabled audit log that only counts close() calls. */
    static final class RecordingAuditLog extends AuditLog {
        int closeCalls;

        RecordingAuditLog() {
            super(new AuditLogConfiguration(), "test");
        }

        @Override
        public void write(AuditLogEventType eventType, String message) {
        }

        @Override
        public void close() {
            closeCalls++;
        }
    }

    static final class TestConnection extends ConnectionBase<Configuration> {
        TestConnection(AuditLog auditLog) {
            super(null, null, auditLog);
        }

        @Override
        protected PlcTagHandler getTagHandler() {
            return null;
        }

        @Override
        protected PlcValueHandler getValueHandler() {
            return null;
        }
    }

    @Test
    void closeReleasesTheAuditLog() throws Exception {
        RecordingAuditLog auditLog = new RecordingAuditLog();

        new TestConnection(auditLog).close();

        // DriverBase creates one per connection and nothing else closes it.
        assertEquals(1, auditLog.closeCalls);
    }
}
