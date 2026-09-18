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
import org.apache.plc4x.java.spi.transports.api.TransportInstance;
import org.apache.plc4x.java.utils.auditlog.api.AuditLog;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import java.util.List;
import java.util.Optional;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertThrows;

/**
 * When getConnection() fails after the transport has been opened, nothing owns that transport yet,
 * so DriverBase has to close it before letting the exception out.
 */
class DriverBaseFailureCleanupTest {

    public static class StubConfiguration implements Configuration {
    }

    /** Connects over the close-counting transport and fails in the step given. */
    static final class StubDriver extends DriverBase {
        private final Class<? extends Configuration> configurationClass;

        StubDriver(Class<? extends Configuration> configurationClass) {
            this.configurationClass = configurationClass;
        }

        @Override public String getProtocolCode() { return "cleanup"; }
        @Override public String getProtocolName() { return "cleanup"; }
        @Override public Optional<String> getDefaultTransportCode() { return Optional.of("close-counting"); }
        @Override public List<String> getSupportedTransportCodes() { return List.of("close-counting"); }
        @Override protected Class<? extends Configuration> getConfigurationClass() { return configurationClass; }
        @Override protected ConnectionBase<?> getConnection(Configuration c, TransportInstance<?> t, AuditLog a) {
            throw new IllegalStateException("driver refuses to connect");
        }
    }

    @BeforeEach
    void resetTransport() {
        CloseCountingTransport.LAST_INSTANCE.set(null);
    }

    @Test
    void aDriverThatFailsToBuildItsConnectionClosesTheTransport() {
        assertThrows(IllegalStateException.class,
            () -> new StubDriver(StubConfiguration.class).getConnection("cleanup:close-counting://host"));

        CloseCountingTransport.Instance transport = CloseCountingTransport.LAST_INSTANCE.get();
        assertNotNull(transport, "the transport should have been opened before the driver failed");
        assertEquals(1, transport.closeCalls.get(), "the opened transport must be closed with the failure");
        assertFalse(transport.isOpen());
    }

    @Test
    void anUnusableProtocolConfigurationClosesTheTransport() {
        // An interface cannot be instantiated as the protocol configuration.
        assertThrows(Exception.class,
            () -> new StubDriver(Configuration.class).getConnection("cleanup:close-counting://host"));

        CloseCountingTransport.Instance transport = CloseCountingTransport.LAST_INSTANCE.get();
        assertNotNull(transport);
        assertEquals(1, transport.closeCalls.get());
    }
}
