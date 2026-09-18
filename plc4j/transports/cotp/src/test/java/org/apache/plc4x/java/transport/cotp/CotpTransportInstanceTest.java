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
package org.apache.plc4x.java.transport.cotp;

import org.apache.plc4x.java.cotp.readwrite.COTPPacket;
import org.apache.plc4x.java.cotp.readwrite.COTPPacketConnectionRequest;
import org.apache.plc4x.java.cotp.readwrite.COTPPacketConnectionResponse;
import org.apache.plc4x.java.cotp.readwrite.COTPPacketDisconnectRequest;
import org.apache.plc4x.java.cotp.readwrite.COTPParameterTpduSize;
import org.apache.plc4x.java.cotp.readwrite.COTPProtocolClass;
import org.apache.plc4x.java.cotp.readwrite.COTPTpduSize;
import org.apache.plc4x.java.cotp.readwrite.TPKTPacket;
import org.apache.plc4x.java.spi.buffers.bytebased.ReadBufferByteBased;
import org.apache.plc4x.java.spi.buffers.bytebased.WriteBufferByteBased;
import org.apache.plc4x.java.spi.transports.api.exceptions.TransportException;
import org.apache.plc4x.java.transport.cotp.config.CotpTransportConfiguration;
import org.apache.plc4x.java.transport.tcp.TcpTransportInstance;
import org.apache.plc4x.java.utils.auditlog.api.AuditLog;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.net.InetSocketAddress;
import java.nio.ByteBuffer;
import java.nio.channels.ServerSocketChannel;
import java.nio.channels.SocketChannel;
import java.util.List;
import java.util.concurrent.BlockingQueue;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.LinkedBlockingQueue;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.atomic.AtomicInteger;

import static org.junit.jupiter.api.Assertions.*;

class CotpTransportInstanceTest {

    /** Counts close() calls so a test can tell whether the COTP layer reached the TCP layer. */
    static final class RecordingTcpTransport extends TcpTransportInstance {
        final AtomicInteger closeCalls = new AtomicInteger();

        RecordingTcpTransport(InetSocketAddress remoteAddress, CotpTransportConfiguration config)
                throws TransportException {
            super(remoteAddress, config, AuditLog.builder().build());
        }

        @Override
        public void close() throws TransportException {
            closeCalls.incrementAndGet();
            super.close();
        }
    }

    private ServerSocketChannel serverChannel;
    private InetSocketAddress serverAddress;
    private volatile SocketChannel peer;
    private Thread peerThread;
    private final BlockingQueue<COTPPacket> received = new LinkedBlockingQueue<>();

    private RecordingTcpTransport tcpTransport;
    private CotpTransportInstance transportInstance;

    @BeforeEach
    void setUp() throws Exception {
        serverChannel = ServerSocketChannel.open();
        serverChannel.bind(new InetSocketAddress("localhost", 0));
        serverAddress = new InetSocketAddress("localhost",
            ((InetSocketAddress) serverChannel.getLocalAddress()).getPort());
    }

    @AfterEach
    void tearDown() throws Exception {
        if (transportInstance != null) {
            transportInstance.close();
        }
        if (peer != null) {
            peer.close();
        }
        serverChannel.close();
        if (peerThread != null) {
            peerThread.join(2000);
        }
    }

    /** Accepts one connection and records every COTP packet; answers a Connection Request when told to. */
    private void startPeer(boolean confirmConnection) {
        peerThread = new Thread(() -> {
            try (SocketChannel accepted = serverChannel.accept()) {
                peer = accepted;
                while (true) {
                    COTPPacket packet = readTpkt(accepted).getPayload();
                    received.add(packet);
                    if (confirmConnection && packet instanceof COTPPacketConnectionRequest) {
                        COTPPacketConnectionResponse confirm = new COTPPacketConnectionResponse(
                            List.of(new COTPParameterTpduSize(COTPTpduSize.SIZE_1024)), new byte[0],
                            0x000F, 0x0001, COTPProtocolClass.CLASS_0);
                        accepted.write(ByteBuffer.wrap(serialize(new TPKTPacket(confirm))));
                    }
                }
            } catch (Exception e) {
                // Closed by the test or by the transport; the packets recorded so far are what matter.
            }
        }, "cotp-test-peer");
        peerThread.start();
    }

    private static TPKTPacket readTpkt(SocketChannel channel) throws Exception {
        ByteBuffer header = ByteBuffer.allocate(4);
        readFully(channel, header);
        int length = ((header.get(2) & 0xFF) << 8) | (header.get(3) & 0xFF);
        ByteBuffer packet = ByteBuffer.allocate(length);
        packet.put(header.array());
        readFully(channel, packet);
        return TPKTPacket.staticParse(new ReadBufferByteBased(packet.array()));
    }

    private static void readFully(SocketChannel channel, ByteBuffer buffer) throws IOException {
        while (buffer.hasRemaining()) {
            if (channel.read(buffer) == -1) {
                throw new IOException("Peer closed");
            }
        }
    }

    private static byte[] serialize(TPKTPacket packet) throws Exception {
        WriteBufferByteBased writeBuffer = new WriteBufferByteBased(new byte[packet.getLengthInBytes()]);
        packet.serialize(writeBuffer);
        return writeBuffer.getBytes();
    }

    private CotpTransportConfiguration configuration() {
        CotpTransportConfiguration config = new CotpTransportConfiguration();
        config.receiveBufferSize = 81920;
        return config;
    }

    private void connect() throws Exception {
        CotpTransportConfiguration config = configuration();
        tcpTransport = new RecordingTcpTransport(serverAddress, config);
        transportInstance = new CotpTransportInstance(tcpTransport, config, AuditLog.builder().build());
        assertTrue(transportInstance.isOpen());
        assertInstanceOf(COTPPacketConnectionRequest.class, received.poll(5, TimeUnit.SECONDS));
    }

    @Test
    void testClose_afterTcpDisconnect_closesTcpTransport() throws Exception {
        startPeer(true);
        connect();
        CountDownLatch disconnected = new CountDownLatch(1);
        transportInstance.registerDisconnectListener(cause -> disconnected.countDown());

        peer.close();

        assertTrue(disconnected.await(5, TimeUnit.SECONDS), "TCP disconnect should propagate");
        assertFalse(transportInstance.isOpen());
        // close() used to return early here because the disconnect had already cleared `connected`.
        transportInstance.close();
        assertEquals(1, tcpTransport.closeCalls.get(), "close() must reach the TCP transport");
    }

    @Test
    void testClose_whenConnected_sendsDisconnectRequestAndClosesTcp() throws Exception {
        startPeer(true);
        connect();

        transportInstance.close();

        assertInstanceOf(COTPPacketDisconnectRequest.class, received.poll(5, TimeUnit.SECONDS));
        assertFalse(transportInstance.isOpen());
        assertEquals(1, tcpTransport.closeCalls.get());
    }

    @Test
    void testClose_idempotent() throws Exception {
        startPeer(true);
        connect();

        transportInstance.close();
        transportInstance.close();

        assertEquals(1, tcpTransport.closeCalls.get());
    }

    @Test
    void testConstructor_failedHandshake_closesTcpTransport() throws Exception {
        startPeer(false);
        CotpTransportConfiguration config = configuration();
        config.cotpConnectionTimeout = 200;
        tcpTransport = new RecordingTcpTransport(serverAddress, config);

        assertThrows(TransportException.class,
            () -> new CotpTransportInstance(tcpTransport, config, AuditLog.builder().build()));

        assertEquals(1, tcpTransport.closeCalls.get(), "a failed handshake must not leave the TCP transport open");
        assertFalse(tcpTransport.isOpen());
    }
}
