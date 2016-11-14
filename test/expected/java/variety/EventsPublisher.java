/**
 * Autogenerated by Frugal Compiler (2.0.0-RC2)
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *
 * @generated
 */

package variety.java;

import com.workiva.frugal.middleware.InvocationHandler;
import com.workiva.frugal.middleware.ServiceMiddleware;
import com.workiva.frugal.protocol.*;
import com.workiva.frugal.provider.FScopeProvider;
import com.workiva.frugal.transport.FPublisherTransport;
import com.workiva.frugal.transport.FSubscriberTransport;
import com.workiva.frugal.transport.FSubscription;
import com.workiva.frugal.transport.TMemoryOutputBuffer;
import org.apache.thrift.TException;
import org.apache.thrift.TApplicationException;
import org.apache.thrift.transport.TTransport;
import org.apache.thrift.transport.TTransportException;
import org.apache.thrift.protocol.*;

import javax.annotation.Generated;
import java.util.logging.Logger;




@Generated(value = "Autogenerated by Frugal Compiler (2.0.0-RC2)", date = "2015-11-24")
public class EventsPublisher {

	/**
	 * This docstring gets added to the generated code because it has
	 * the @ sign. Prefix specifies topic prefix tokens, which can be static or
	 * variable.
	 */
	public interface Iface {
		public void open() throws TException;

		public void close() throws TException;

		/**
		 * This is a docstring.
		 */
		public void publishEventCreated(FContext ctx, String user, Event req) throws TException;

	}

	/**
	 * This docstring gets added to the generated code because it has
	 * the @ sign. Prefix specifies topic prefix tokens, which can be static or
	 * variable.
	 */
	public static class Client implements Iface {
		private static final String DELIMITER = ".";

		private final Iface target;
		private final Iface proxy;

		public Client(FScopeProvider provider, ServiceMiddleware... middleware) {
			target = new InternalEventsPublisher(provider);
			proxy = InvocationHandler.composeMiddleware(target, Iface.class, middleware);
		}

		public void open() throws TException {
			target.open();
		}

		public void close() throws TException {
			target.close();
		}

		/**
		 * This is a docstring.
		 */
		public void publishEventCreated(FContext ctx, String user, Event req) throws TException {
			proxy.publishEventCreated(ctx, user, req);
		}

		protected static class InternalEventsPublisher implements Iface {

			private FScopeProvider provider;
			private FPublisherTransport transport;
			private FProtocolFactory protocolFactory;

			protected InternalEventsPublisher() {
			}

			public InternalEventsPublisher(FScopeProvider provider) {
				this.provider = provider;
			}

			public void open() throws TException {
				FScopeProvider.Publisher publisher = provider.buildPublisher();
				transport = publisher.getTransport();
				protocolFactory = publisher.getProtocolFactory();
				transport.open();
			}

			public void close() throws TException {
				transport.close();
			}

			/**
			 * This is a docstring.
			 */
			public void publishEventCreated(FContext ctx, String user, Event req) throws TException {
				String op = "EventCreated";
				String prefix = String.format("foo.%s.", user);
				String topic = String.format("%sEvents%s%s", prefix, DELIMITER, op);
				TMemoryOutputBuffer memoryBuffer = new TMemoryOutputBuffer(transport.getPublishSizeLimit());
				FProtocol protocol = protocolFactory.getProtocol(memoryBuffer);
				protocol.writeRequestHeader(ctx);
				protocol.writeMessageBegin(new TMessage(op, TMessageType.CALL, 0));
				req.write(protocol);
				protocol.writeMessageEnd();
				transport.publish(topic, memoryBuffer.getWriteBytes());
			}
		}
	}
}
