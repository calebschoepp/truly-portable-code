# Setup

- [Setup](#setup)
  - [Request Access to Fermyon Wasm Functions](#request-access-to-fermyon-wasm-functions)
  - [Install Spin](#install-spin)
  - [Troubleshooting](#troubleshooting)
    - [Q: I cannot build my Rust application with `spin build`.](#q-i-cannot-build-my-rust-application-with-spin-build)
    - [Q: I cannot build my JavaScript or TypeScript application with `spin build`.](#q-i-cannot-build-my-javascript-or-typescript-application-with-spin-build)
  - [Learning Summary](#learning-summary)
  - [Navigation](#navigation)

This module will guide you through the pre-requisites needed to complete the rest of the workshop.

## Request Access to Fermyon Wasm Functions

Fermyon Wasm Functions is currently in public preview. As such you'll need to request access to deploy to it later in this workshop. When it asks for the reason for wanting access enter something based on the conference you're at e.g. `TXLF 2025 workshop`. Fill out the form [here](https://fibsu0jcu2g.typeform.com/fwf-preview).

It's important to do this at the start of the workshop so that you can be granted access in time for you to use Fermyon Wasm Functions at the end of the workshop.

## Install Spin

First, you have to install [Spin](https://spinframework.dev/). To do so follow the instructions [here](https://spinframework.dev/v3/install). If you did not install via the installer script or homebrew you'll also need to [install the templates](https://spinframework.dev/v3/install#templates).

This workshop assumes you're using Spin `3.4.1`. You can check your version of Spin by running `spin -V`.

## Troubleshooting

The following are some common issues you may run into while going through the modules. Refer back to this as needed.

### Q: I cannot build my Rust application with `spin build`.

A: Make sure you have [configured your Rust toolchain](https://www.rust-lang.org/tools/install), and have added the `wasm32-wasi` target using `rustup target add wasm32-wasi`.

### Q: I cannot build my JavaScript or TypeScript application with `spin build`.

A: Make sure you have [configured Node.js and `npm`](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm), and that you have executed `npm install` in the directory with your component's `package.json` file that contains the dependencies.

### Q: Something else is going wrong.

A: Please ask questions on the CNCF Slack [`#spin`](https://cloud-native.slack.com/archives/C089NJ9G1V0) and [`#spinkube`](https://cloud-native.slack.com/archives/C06PC7JA1EE) channels.

## Learning Summary

In this module you learned how to:

- Install Spin
- Request access to Fermyon Wasm Functions

## Navigation

- Proceed to [Local Spin](01-spin.md).

If you have any feedback for this module please open an issue on the GitHub repo [here](https://github.com/calebschoepp/truly-portable-code/issues/new).
