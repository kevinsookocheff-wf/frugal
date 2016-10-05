package frugal

import (
	"fmt"
	"net"
	"testing"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/nats-io/gnatsd/server"
	"github.com/nats-io/nats"
	"github.com/stretchr/testify/assert"
)

var defaultOptions = server.Options{
	Host:   "localhost",
	Port:   11222,
	NoLog:  true,
	NoSigs: true,
}

func runServer(opts *server.Options) *server.Server {
	if opts == nil {
		opts = &defaultOptions
	}
	s := server.New(opts)
	if s == nil {
		panic("No NATS Server object returned.")
	}

	// Run server in Go routine.
	go s.Start()

	end := time.Now().Add(10 * time.Second)
	for time.Now().Before(end) {
		addr := s.GetListenEndpoint()
		if addr == "" {
			time.Sleep(10 * time.Millisecond)
			// Retry. We might take a little while to open a connection.
			continue
		}
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			// Retry after 50ms
			time.Sleep(50 * time.Millisecond)
			continue
		}
		conn.Close()
		// Wait a bit to give a chance to the server to remove this
		// "client" from its state, which may otherwise interfere with
		// some tests.
		time.Sleep(25 * time.Millisecond)

		return s
	}
	panic("Unable to start NATS Server in Go Routine")

}

// Ensures LockTopic returns nil when a publisher successfully locks a topic.
// Subsequent calls will wait on the mutex. Unlock releases the topic.
func TestNatsPublisherLockUnlockTopic(t *testing.T) {
	tr := NewNatsFPublisherTransport(nil)
	assert.Nil(t, tr.LockTopic("foo"))
	acquired := make(chan bool)
	go func() {
		assert.Nil(t, tr.LockTopic("bar"))
		assert.Equal(t, "bar", tr.(*fNatsPublisherTransport).subject)
		acquired <- true
	}()
	assert.Equal(t, "foo", tr.(*fNatsPublisherTransport).subject)
	tr.UnlockTopic()
	<-acquired

	tr.UnlockTopic()
	assert.Equal(t, "", tr.(*fNatsPublisherTransport).subject)
}

// Ensures Subscribe subscribes to the topic on NATS and puts received frames
// on the read buffer and received in Read calls. All subscribers receive the
// message when they aren't subscribed to a queue.
func TestNatsSubscriberSubscribe(t *testing.T) {
	s := runServer(nil)
	defer s.Shutdown()
	conn, err := nats.Connect(fmt.Sprintf("nats://localhost:%d", defaultOptions.Port))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	tr1 := NewNatsFSubscriberTransport(conn)
	tr2 := NewNatsFSubscriberTransport(conn)

	cb1Called := make(chan bool, 1)
	cb1 := func(transport thrift.TTransport) error {
		cb1Called <- true
		return nil
	}
	cb2Called := make(chan bool, 1)
	cb2 := func(transport thrift.TTransport) error {
		cb2Called <- true
		return nil
	}

	assert.Nil(t, tr1.Subscribe("foo", cb1))
	assert.Nil(t, tr2.Subscribe("foo", cb2))

	frame := make([]byte, 50)
	assert.Nil(t, conn.Publish("frugal.foo", append(make([]byte, 4), frame...))) // Add 4 bytes for frame size

	select {
	case <-cb1Called:
	case <-time.After(time.Second):
		assert.True(t, false, "Callback1 was not called")
	}

	select {
	case <-cb2Called:
	case <-time.After(time.Second):
		assert.True(t, false, "Callback2 was not called")
	}
}

// Ensures Subscribe subscribes to the topic on NATS and puts received frames
// on the read buffer. If the transport specifies a queue, only one member of
// the queue group receives the message.
func TestNatsSubscriberSubscribeQueue(t *testing.T) {
	s := runServer(nil)
	defer s.Shutdown()
	conn, err := nats.Connect(fmt.Sprintf("nats://localhost:%d", defaultOptions.Port))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	tr1 := NewNatsFSubscriberTransportWithQueue(conn, "fooqueue")
	tr2 := NewNatsFSubscriberTransportWithQueue(conn, "fooqueue")

	cb1Called := make(chan bool, 1)
	cb1 := func(transport thrift.TTransport) error {
		cb1Called <- true
		return nil
	}
	cb2Called := make(chan bool, 1)
	cb2 := func(transport thrift.TTransport) error {
		cb2Called <- true
		return nil
	}
	assert.Nil(t, tr1.Subscribe("foo", cb1))
	assert.Nil(t, tr2.Subscribe("foo", cb2))

	frame := make([]byte, 50)
	assert.Nil(t, conn.Publish("frugal.foo", append(make([]byte, 4), frame...))) // Add 4 bytes for frame size
	conn.Flush()
	time.Sleep(10 * time.Millisecond)

	// Only one of the two transports should have received the frame.

	assert.True(t, len(cb1Called) != len(cb2Called))
}

// Ensures Subscribe returns an error if the NATS connection is not open.
func TestNatsSubscriberSubscribeConnectionNotOpen(t *testing.T) {
	s := runServer(nil)
	defer s.Shutdown()
	conn, err := nats.Connect(fmt.Sprintf("nats://localhost:%d", defaultOptions.Port))
	if err != nil {
		t.Fatal(err)
	}
	tr := NewNatsFSubscriberTransport(conn)
	conn.Close()

	assert.Error(t, tr.Subscribe("foo", func(thrift.TTransport) error {
		return nil
	}))
}

// Ensures Open returns nil on success and writes work.
func TestNatsPublisherOpenWriteFlush(t *testing.T) {
	s := runServer(nil)
	defer s.Shutdown()
	conn, err := nats.Connect(fmt.Sprintf("nats://localhost:%d", defaultOptions.Port))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	tr := NewNatsFPublisherTransport(conn)

	assert.Nil(t, tr.Open())
	assert.True(t, tr.IsOpen())
	assert.Nil(t, tr.LockTopic("foo"))
	frame := make([]byte, 10)
	n, err := tr.Write(frame)
	assert.Nil(t, err)
	assert.Equal(t, 10, n)
	received := make(chan bool)
	_, err = conn.Subscribe("frugal.foo", func(msg *nats.Msg) {
		assert.Equal(t, append([]byte{0, 0, 0, 10}, frame...), msg.Data)
		received <- true
	})
	assert.Nil(t, err)
	assert.Nil(t, tr.Flush())
	select {
	case <-received:
	case <-time.After(time.Second):
		t.Fatal("expected to receive frame")
	}
}

// Ensures Open returns an ALREADY_OPEN TTransportException if the transport is
// already open.
func TestNatsPublisherOpenAlreadyOpen(t *testing.T) {
	s := runServer(nil)
	defer s.Shutdown()
	conn, err := nats.Connect(fmt.Sprintf("nats://localhost:%d", defaultOptions.Port))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	tr := NewNatsFPublisherTransport(conn)

	assert.Nil(t, tr.Open())

	err = tr.Open()

	trErr := err.(thrift.TTransportException)
	assert.Equal(t, thrift.ALREADY_OPEN, trErr.TypeId())
}

// Ensures subscribers discard invalid frames (size < 4).
func TestNatsSubscriberDiscardInvalidFrame(t *testing.T) {
	s := runServer(nil)
	defer s.Shutdown()
	conn, err := nats.Connect(fmt.Sprintf("nats://localhost:%d", defaultOptions.Port))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	tr := NewNatsFSubscriberTransport(conn)

	cbCalled := false
	cb := func(transport thrift.TTransport) error {
		cbCalled = true
		return nil
	}
	assert.Nil(t, tr.Subscribe("blah", cb))

	assert.Nil(t, conn.Publish("frugal.blah", make([]byte, 2)))
	assert.Nil(t, conn.Flush())
	time.Sleep(10 * time.Millisecond)
	assert.Nil(t, tr.Unsubscribe())
	assert.False(t, cbCalled)
}

// Ensures Close returns nil if the transport is not open.
func TestNatsPublisherCloseNotOpen(t *testing.T) {
	s := runServer(nil)
	defer s.Shutdown()
	conn, err := nats.Connect(fmt.Sprintf("nats://localhost:%d", defaultOptions.Port))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	tr := NewNatsFPublisherTransport(conn)
	assert.Nil(t, tr.Close())
	assert.False(t, tr.IsOpen())
}

// Ensures Close closes the publisher transport and returns nil.
func TestNatsPublisherClosePublisher(t *testing.T) {
	s := runServer(nil)
	defer s.Shutdown()
	conn, err := nats.Connect(fmt.Sprintf("nats://localhost:%d", defaultOptions.Port))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	tr := NewNatsFPublisherTransport(conn)
	assert.Nil(t, tr.LockTopic("foo"))
	assert.Nil(t, tr.Open())
	assert.True(t, tr.IsOpen())
	assert.Nil(t, tr.Close())
	assert.False(t, tr.IsOpen())
}

// Ensures Close returns an error if the unsubscribe fails.
func TestNatsCloseSubscriberError(t *testing.T) {
	s := runServer(nil)
	defer s.Shutdown()
	conn, err := nats.Connect(fmt.Sprintf("nats://localhost:%d", defaultOptions.Port))
	if err != nil {
		t.Fatal(err)
	}
	tr := NewNatsFSubscriberTransport(conn)
	assert.Nil(t, tr.Subscribe("foo", func(thrift.TTransport) error {
		return nil
	}))
	conn.Close()
	assert.Error(t, tr.Unsubscribe())
}

// Ensures Write returns an ErrTooLarge if the written frame exceeds 1MB.
func TestNatsPublisherWriteTooLarge(t *testing.T) {
	s := runServer(nil)
	defer s.Shutdown()
	conn, err := nats.Connect(fmt.Sprintf("nats://localhost:%d", defaultOptions.Port))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	tr := NewNatsFPublisherTransport(conn)
	assert.Nil(t, tr.Open())

	n, err := tr.Write(make([]byte, 5))
	assert.Equal(t, 5, n)
	assert.Nil(t, err)
	n, err = tr.Write(make([]byte, 1024*1024+10))
	assert.Equal(t, 0, n)
	assert.Equal(t, ErrTooLarge, err)
	assert.Equal(t, 0, tr.(*fNatsPublisherTransport).writeBuffer.Len())
}

// Ensures Flush returns an error if the transport is not open.
func TestNatsPublisherFlushNotOpen(t *testing.T) {
	s := runServer(nil)
	defer s.Shutdown()
	conn, err := nats.Connect(fmt.Sprintf("nats://localhost:%d", defaultOptions.Port))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	tr := NewNatsFPublisherTransport(conn)

	err = tr.Flush()
	trErr := err.(thrift.TTransportException)
	assert.Equal(t, thrift.NOT_OPEN, trErr.TypeId())
}

// Ensures Flush returns nil and nothing is sent to NATS when there is no data
// to flush.
func TestNatsPublisherFlushNoData(t *testing.T) {
	s := runServer(nil)
	defer s.Shutdown()
	conn, err := nats.Connect(fmt.Sprintf("nats://localhost:%d", defaultOptions.Port))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	tr := NewNatsFPublisherTransport(conn)
	assert.Nil(t, tr.Open())
	assert.Nil(t, tr.LockTopic("foo"))
	_, err = conn.Subscribe("frugal.foo", func(msg *nats.Msg) {
		t.Fatal("did not expect to receive message")
	})
	assert.Nil(t, err)
	assert.Nil(t, tr.Flush())
	assert.Nil(t, conn.Flush())
	time.Sleep(10 * time.Millisecond)
}
