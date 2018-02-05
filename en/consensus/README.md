# Consensus

References:
[Wikipedia Consensus](https://en.wikipedia.org/wiki/Consensus_(computer_science))

- Sharing a common state across systems
- Agreeing on changes across systems
- Distributed consensus
- Raft, Paxos, Zab

# CAP theorem

- [Consistency](https://en.wikipedia.org/wiki/Consistency_(database_systems))
- [Availability](https://en.wikipedia.org/wiki/Availability)
- [Partition Tolerance](https://en.wikipedia.org/wiki/Network_partition)

References:
- [Wikipedia Cap Theorem](https://en.wikipedia.org/wiki/CAP_theorem)

# Building your own consensus using Go libraries

[raft](https://raft.github.io)

We will use the HashiCorp raft library for our solution.

raft is:
- [Consistent](https://en.wikipedia.org/wiki/Consistency_(database_systems))
- [Partition Tolerant](https://en.wikipedia.org/wiki/Network_partition)
- [Byzantine fault tolerant](https://en.wikipedia.org/wiki/Byzantine_fault_tolerance)

raft is NOT:
- [Always Available](https://en.wikipedia.org/wiki/Availability) 
	- but shouldn't be a problem with proper cluster design

References:
- [HashiCorp Raft Source](https://github.com/hashicorp/raft) 
- [HashiCorp Raft Documentation](https://godoc.org/github.com/hashicorp/raft) 

# Design Considerations - How many nodes?

There are a lot of good articles on how many nodes to have in a consensus cluster, like the one from [etcd](https://coreos.com/etcd/docs/latest/admin_guide.html#optimal-cluster-size)

The short version is you need at least 3 for fault tolerance, and good practice is to always use odd numbers to avoid a split vote (Byzantine failures).

# Common mistakes

- Reading from the finite state machine on a node that isn't the leader (this is a `dirty read`)
- Not calling `raft.Barrier` before reading to ensure all pending writes have applied.

# Exercise

Go through:

    src/consensus/demos/raft

Run it on as single server, then as three nodes (with different ports)

### Using Etcd and Consul for distributed consensus

You don't have to write your own consensus system.  You can use Consul or Etcd or Zookeeper as your storage engine for state.

But if you use Raft, it isn't that hard and it possibly removes a layer of complexity from your system.

### Prove it works

You should write your distributed software with the assumption your network is always failing.

Netflix forces their development team to address this:

- [Chaos Monkey (now written in go - see below)](https://github.com/Netflix/SimianArmy/wiki/Chaos-Monkey)
- [New chaos monkey](https://github.com/Netflix/chaosmonkey) 
