# etcd部署
```
docker-compose up -d

docker ps 
CONTAINER ID        IMAGE                 COMMAND                  CREATED              STATUS              PORTS                                              NAMES
3241574326e9        quay.io/coreos/etcd   "etcd -name etcd3 -a…"   About a minute ago   Up 4 seconds        0.0.0.0:32779->2379/tcp, 0.0.0.0:32777->2380/tcp   etcd3
24b3533caa6c        quay.io/coreos/etcd   "etcd -name etcd2 -a…"   About a minute ago   Up 4 seconds        0.0.0.0:32775->2379/tcp, 0.0.0.0:32774->2380/tcp   etcd2
d0c51d97b19b        quay.io/coreos/etcd   "etcd -name etcd1 -a…"   About a minute ago   Up 4 seconds        0.0.0.0:32778->2379/tcp, 0.0.0.0:32776->2380/tcp   etcd1
```
```
curl -L http://127.0.0.1:32778/v2/members
{"members":[{"id":"ade526d28b1f92f7","name":"etcd1","peerURLs":["http://etcd1:2380"],"clientURLs":["http://0.0.0.0:2379"]},{"id":"bd388e7810915853","name":"etcd3","peerURLs":["http://etcd3:2380"],"clientURLs":["http://0.0.0.0:2379"]},{"id":"d282ac2ce600c1ce","name":"etcd2","peerURLs":["http://etcd2:2380"],"clientURLs":["http://0.0.0.0:2379"]}]}

curl -L http://127.0.0.1:32775/v2/members
{"members":[{"id":"ade526d28b1f92f7","name":"etcd1","peerURLs":["http://etcd1:2380"],"clientURLs":["http://0.0.0.0:2379"]},{"id":"bd388e7810915853","name":"etcd3","peerURLs":["http://etcd3:2380"],"clientURLs":["http://0.0.0.0:2379"]},{"id":"d282ac2ce600c1ce","name":"etcd2","peerURLs":["http://etcd2:2380"],"clientURLs":["http://0.0.0.0:2379"]}]}

curl -L http://127.0.0.1:32779/v2/members
{"members":[{"id":"ade526d28b1f92f7","name":"etcd1","peerURLs":["http://etcd1:2380"],"clientURLs":["http://0.0.0.0:2379"]},{"id":"bd388e7810915853","name":"etcd3","peerURLs":["http://etcd3:2380"],"clientURLs":["http://0.0.0.0:2379"]},{"id":"d282ac2ce600c1ce","name":"etcd2","peerURLs":["http://etcd2:2380"],"clientURLs":["http://0.0.0.0:2379"]}]}
```
```
docker exec -t etcd1 etcdctl member list
ade526d28b1f92f7: name=etcd1 peerURLs=http://etcd1:2380 clientURLs=http://0.0.0.0:2379 isLeader=false
bd388e7810915853: name=etcd3 peerURLs=http://etcd3:2380 clientURLs=http://0.0.0.0:2379 isLeader=true
d282ac2ce600c1ce: name=etcd2 peerURLs=http://etcd2:2380 clientURLs=http://0.0.0.0:2379 isLeader=false

```
