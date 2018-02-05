# Service and Peer Discovery

When you have machines talking to each over over a distributed system, they need a way to find each other.  This module will discuss the various ways to do this and present the pros and cons of each choice.

# Static configuration

An obvious first choice is static configuration.  Host names and ports are kept in a static configuration file and looked up at runtime.

Pros:
- it doesn't get any simpler

Cons:
- static file means changing config and restarting services when hosts/services change
- very low flexibility in an evolving environment

# Multicast DNS

Multicast DNS sends DNS-SD messages across the network.  Services announce their presence by broadcasting  messages periodically.  Clients listen for DNS-SD announcements and choose a server/service from these messages.

Pros:
- Relatively simple code to broadcast and listen.  Easily accomplished with Go's std library
- More dynamic than config files, new services announce themselves.

Cons:
- DNS-SD doesn't route, so confined to one network segment without a special rebroadcasting node spanning segments.
- Clients poll for changes, reconnect servers disappear.

### Example :  
Apple's Bonjour protocol.
- [Hashicorp mdns](https://github.com/hashicorp/mdns)
- [Dave Cheney's version](https://github.com/davecheney/mdns/)

# Service registries

Service Registries store host/port/service information for services running in a distributed system.  Clients look up services using an API, or by using the registry's custom DNS Server (consul, skydns)

Pros:
- Real time updating
- Spans network routers
- registries are distributed for no SPOF
- Some registries provide more services than just Service Discovery, which may be useful

Cons:
- DNS cache can cause trouble, client must respect DNS TTL
- When using DNS to lookup, requires more configuration
- When using API to lookup, requires more code

Example: Consul, SkyDNS, Zookeeper, Kubernetes Services, Service Fabric

# Dynamically generated static configuration

Sometimes you have a client that can't be modified to use one of the other dynamic options.  In a case like this, use a tool that dynamically generates a configuration file for the client.

Pros:
- Allows dynamic config of an otherwise static process

Cons:
- Requires a dynamic SD mechanism, and another tool to watch config values and regenerate config from a template
- Requires a service watchdog to restart the service when configs change

### Example:  
- [confd](https://github.com/kelseyhightower/confd) 

# Hybrid

The truth of the matter is you'll need more than one of these in a complex system.  Many end up using a primary Service Registry, back it up with DNS, and use confd or other configuration generation tools for things like Nginx.

# Exercise

You can find this exercise in folder *material/disco/exercises/contextrpc*

Start a single consul node:

	$ docker run -p 8400:8400 -p 8500:8500 -p 8600:53/udp -h `hostname` progrium/consul -server -bootstrap -ui-dir /ui

(This may not work for you if you're running older docker for mac. //FIXME with new instructions there.)

Register the service in consul:

Get local IP address with `ifconfig`  Find the public ip.

	run the service:
	SERVER_IP=x.x.x.x ./server

Run the `dig` command to see it in Consul's DNS:

	dig @127.0.0.1 -p 8600 hello.service.consul SRV    

Run `curl` to see the consul API's registry version of the same record:

	curl http://127.0.0.1:8500/v1/catalog/services

For more detailed information about the service, `curl`: 

	curl http://127.0.0.1:8500/v1/agent/services

Modify the client to read the server address from the consul registry, then call the service