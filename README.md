# goReplicateServers
[WIP] Automatically replicate a Go server across ports.

Given the route implementations, a server can be replicated to be deployed in different ports sequentially (tested locally) using go routines & net/http to create a new replicas of the original server.  
