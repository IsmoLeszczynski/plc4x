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

import org.apache.plc4x.java.spi.transports.api.Transport;
import org.apache.plc4x.java.spi.transports.api.TransportInstance;
import org.apache.plc4x.java.spi.transports.api.config.TransportConfiguration;
import org.apache.plc4x.java.spi.transports.api.exceptions.TransportException;
import org.apache.plc4x.java.utils.auditlog.api.AuditLog;

import java.util.concurrent.atomic.AtomicInteger;
import java.util.concurrent.atomic.AtomicReference;

/**
 * Test transport, registered through META-INF/services, whose instances count their close() calls.
 */
public class CloseCountingTransport implements Transport<CloseCountingTransport.Config> {

    public static class Config implements TransportConfiguration {
    }

    /** The most recently created instance, for tests to inspect. */
    public static final AtomicReference<Instance> LAST_INSTANCE = new AtomicReference<>();

    public static final class Instance implements TransportInstance<Config> {
        public final AtomicInteger closeCalls = new AtomicInteger();
        private final Config configuration;
        private volatile boolean open = true;

        Instance(Config configuration) {
            this.configuration = configuration;
        }

        @Override
        public Config getConfiguration() {
            return configuration;
        }

        @Override
        public boolean isOpen() {
            return open;
        }

        @Override
        public int getNumBytesAvailable() {
            return 0;
        }

        @Override
        public byte[] peekReadableBytes(int numBytes) {
            return new byte[0];
        }

        @Override
        public byte[] read(int numBytes) {
            return new byte[0];
        }

        @Override
        public void write(byte[] bytes) {
        }

        @Override
        public void close() {
            closeCalls.incrementAndGet();
            open = false;
        }
    }

    @Override
    public String getTransportCode() {
        return "close-counting";
    }

    @Override
    public String getTransportName() {
        return "Test transport that counts close() calls";
    }

    @Override
    public Class<Config> getTransportConfigType() {
        return Config.class;
    }

    @Override
    public TransportInstance<Config> createTransportInstance(String transportUrl, TransportConfiguration configuration,
            AuditLog auditLog) throws TransportException {
        Instance instance = new Instance((Config) configuration);
        LAST_INSTANCE.set(instance);
        return instance;
    }
}
