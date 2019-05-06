# Notes on Cloud Native apps

Three key principles fo Cloud native apps

- Distrupt
- Industrialize

Hyperscale, traffic, data, features etc.

You need app to handle lots of scale

**antifragility** cloud native apps don't break anywhere, no single point of failure.

cloud native --> continous delivery and devops

**opex savings** automation and utilization. Only pay for what you use.

packaged, distributed in containers.

## Design challenges

CHange to reactive programming, needs to be designed to automation. For continous delivery, need automation and elasticity. The apps should not break and be fault tolerant and dynamically scale up and down. Elastic, so they scale up and down dynamically to changes in environment. Also you need complex logging and services to debug when things go wrong.


## Decompositions with microservices

basic tradeoffs with more microservices are latency, higher infra complexity and costs.

## Ops components
What is an ops component? Application packaged inside container, interface that talks I/O like http. COntainer with starting interface (rest api) withj a diagnosis interface for debugging. But you gain better runtime isolation (no single point of failure). Higher resource utilization.

limitations
- app does o use kernel space APIS (oracle)
- app does not listen on random ports (RMI)
- app does not require exotic OS eg z/OS
- used endpoints (DNS name, IPs, Ports) can be configured

## Cloud Native Stack
4 layers (bottom --> top)
- cluster virtualization -- decouples it all from physical hardware
- cluster scheduler - container
- cluster orchestrator talks to scheduler, tells scheduler to do an app
- application platform (runtime env and api for apps)


