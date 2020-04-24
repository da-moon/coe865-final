# coe865-final

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io#https://github.com/da-moon/coe865-final)

## gossip agent : overview

the agent has the following sub-components :

- `sentry` : it stores the agents private RSA key and it's job is to confirm message signatures,
generate new nonce and encrypt/decrypt payloads based on my `DARE` ( data at rest encryption) refrence implementation
- `swarm` : holds a snapshot of all the gossip nodes the agent is directly connected to. it has an upper limit which controls the maximum number of nodes
the agent has established a direct connection to
- `??` : holds information about the state of the network and uses djikstra's algorithm calculate best path
- `scheduler` : a background task scheduler that fires an event, based on given cron specs

### gossip agent : init

initially , when every gossip agent daemon comes online, starts a listener and waits for incomming connection.
it also takes in a list of initial bootstrap peers to connect to which it tries to establish direct connection to.

swarm's listener spawns a new goroutine per incoming connection. the newly created incomming connection 
is passed down to underlying `swarm` component for handling the new peer that established connection to it. 
`swarm` component , checks to see if it still has room for more direct peers. it replies back with a `hello message`
in which it indicates whether the connection was accepted or dropped based on the number of direct
connections it was configured to accept. if the connection was dropped , 
it returns back with an `error` to agent and agent, by default ,drops connections 
if there is an error from swarm. in case the connection was accepted, `swarm` replies back with addresses
of peers it is directly connected to so that the other agent can start infecting them.
when agent recieve a `hello message` that shows the other agent they are trying to connect to 
accepted the connection, they reply back with their own `hello message` which contains
list of nodes they are directly connected to which the recieving agent would try to 
establish direct connection to them.

### gossip agent : rumour

the gossip agent, with the help of `scheduler` component tries to spread a `rumor` at set time
quantums. the `rumor` is result of `??` calculating best path using `djikstra's` algorithm.
once the message is prepared, it is sent to `sentry` component. `sentry` would encrypt the payload
with `DARE`, i.e. generate a new nonce and encrypt the message with the said nonce and original key.
it also would sign the message and attach the signature before returning the result to agent.
then agent would pass in the message to `swarm` component so that it can encode the message with json
and pushes it on the wire for all nodes it is directly connected to

once `swarm` recieves a rumor, it decodes it from json and through it's internal channel with agent , sends
the decode message to agent for processing. once agent picks up a message from channel , it would send the message
to `sentry`. `sentry` would confirm message authenticity and reply back to agent. at this point based on `sentry's` reply, 
the agent can take two actions : reject the message or infect other nodes connected to it.

in case `sentry` replies back with an error, the agent would reject the message. it also keeps track of how
many messages it has rejected from a particular node. if it passes a certain threshold, it would ask `swarm`
component to remove the node from it's address book and broadcast to others that the node is behaving maliciously
so that other nodes remove the bad actor from their own swarm, effectively kicking it out from the network

in case the message was authentic, `sentry` would generate a new nonce and re-encrypts the message with new nonce
before returning it back to agent . this nonce chaining scheme is why my `DARE` implementation can prevent 
replay attacks.

### gossip agent : leave
