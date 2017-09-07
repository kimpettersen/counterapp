# counterapp

TODO Deploy:
- Set up private docker registry
- Set up build scripts that pushes image to docker registry
- Set up deploy scripts that pulls from registr and runs container
- Run unit tests
- Run integration/smoke test
on either localhost, Nomad, k8 or swarm?
- Figure out how to run the scripts. Command line?
- Service discovery
- Load balancing

TODO app:
- Create `/status` endpoint
- Create `/get` endpoint
- Set up Redis

## Counter solution
CAP..


### v1 Single instance
- Might be performant enough
- Simple locking and full control by using in app mutexes
- Not scalable

### v2 Multiple Instances distributed lock
- Complicated locking. Need to use something like Redlock
- Could potentially block for a long time and be a bad user experience
- This is scalable over several availability zones
- Likely not performant
- Likely very complicated when it comes to database replication

### v3 Multiple instances publish event on request and have separate reader
- Simple solution, no locking
- You can return fast. Handle request,  publish event then return 201(?)
- Eventual consistency. calling `/status` might be wrong
- DB managment is a lot easier
- Scalable


## How to

### Set up private registry
This is super unsafe
- `docker run -d -p 5000:5000 --restart=always --name registry registry:2`
- `docker build -t counterapp app`
- `docker tag counterapp localhost:5000/counterapp`
- `docker push localhost:5000/counterapp`
- `docker run --publish 6060:8080 --name test --rm localhost:5000/counterapp`

### Build and run container
- Build: `docker build -t counterapp app`
- Run: `docker run --publish 6060:8080 --name test --rm counterapp`


