# Truly Portable Code: Serverless WebAssembly in a Distributed World

> What if you could build serverless applications that cold-start in under a millisecond, run anywhere—from your laptop to Kubernetes to the edge—and require no changes to move between environments? This workshop introduces Spin, a CNCF open-source WebAssembly (Wasm) developer toolkit designed for performance, portability, and simplicity. Attendees will learn how to build a Spin app, write polyglot WebAssembly functions with sub-millisecond cold starts, and run them locally using the Spin CLI. The same app will then be deployed to Azure Kubernetes Service with SpinKube, the open-source Spin runtime for Kubernetes, and to Fermyon Wasm Functions, Akamai’s multi-tenant, globally distributed PaaS — all without rewriting or cross-compilation. The workshop teaches how WebAssembly and Spin enable true portability across the compute continuum, letting developers build once and run anywhere with no vendor lock-in. This tutorial demonstrates how Spin is reshaping what serverless can be.

In this workshop we'll be taking the role of a developer tasked with the job of building a serverless URL shortener. The service will allow users to create short URLs (slugs) that redirect to longer URLs. The service will be built using Spin and WebAssembly, and will be deployed to multiple environments including local Spin, SpinKube on Kubernetes, and Fermyon Wasm Functions.

## Workshop Modules

0. [Setup](00-setup.md)
1. [Local Spin](01-spin.md)
2. [Deploying to SpinKube](02-spinkube.md)
3. [Deploying to Fermyon Wasm Functions](03-fwf.md)
4. [Bonus](04-bonus.md)

## Acknowledgements

Thank you to the Fermyon [Wasm and Containers workshop](https://github.com/fermyon/workshops/blob/main/wasm-and-containers/README.md) for providing a base of content for this workshop.
