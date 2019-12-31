# goReplicateServers
Automatically replicate a Go server across ports (tested locally).

Given the route implementations, a server can be replicated to be deployed in different ports sequentially using go routines & net/http to create a new replicas of the original server.  
