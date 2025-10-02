# Local Spin

- [Local Spin](#local-spin)
  - [1. Spin New - Create a Spin application and choose a template](#1-spin-new---create-a-spin-application-and-choose-a-template)
  - [2. Spin Build and Spin Up - Build and run the application](#2-spin-build-and-spin-up---build-and-run-the-application)
  - [3. Writing the URL Shortener](#3-writing-the-url-shortener)
  - [Learning Summary](#learning-summary)
  - [Navigation](#navigation)

In this module we'll explore how to build an application using Spin, an open-source developer tool for building and running serverless applications with WebAssembly (Wasm).

Spin uses Wasm because of its portability, sandboxed execution environment, non-existent cold starts, and near-native speed. More and more languages have support for WebAssembly, so you should be able to use your favorite language to build your first serverless application with Wasm.

In this workshop, we've provided working examples for Rust, TypeScript, Go, and Python. You can find these examples in their respective directories.

We'll be showing Rust code throughout, but feel free to use whatever language you're most comfortable with and copy the code from the example directories.

In this workshop we'll be taking the role of a developer tasked with the job of building a serverless URL shortener. The service will allow users to create short URLs (slugs) that redirect to longer URLs.

First, let's learn how to scaffold a new Spin application.

## 1. Spin New - Create a Spin application and choose a template

`spin new` is the command you'll use to initialize a new Spin application. A Spin application can consist of multiple components, which are triggered independently.

When you run `spin new`, you are provided with a list of available templates for a Spin component. Please pick any template, which starts with `http-` prefix from the list.

> **Note:** We recommend Rust, Python, Go, or TS. If choosing any other template, it is not guaranteed that all modules of this workshop can be completed successfully.

```bash
$ spin new
Pick a template to start your application with:
> http-c (HTTP request handler using C and the Zig toolchain)
  http-empty (HTTP application with no components)
  http-go (HTTP request handler using (Tiny)Go)
  http-grain (HTTP request handler using Grain)
  http-js (HTTP request handler using JavaScript)
  http-php (HTTP request handler using PHP)
  http-py (HTTP request handler using Python)
  http-rust (HTTP request handler using Rust)
  http-swift (HTTP request handler using SwiftWasm)
  http-ts (HTTP request handler using TypeScript)
  http-zig (HTTP request handler using Zig)
...
```

You can also provide the `--template <template-name>` parameter to the `spin new` command, to choose a template.

E.g. For using Rust:

```bash
$ spin new url-shortener --template http-rust
Description: Shortens URLs
HTTP path: /...
```

Once the application is created, a sub-directory is created with the name corresponding to the application name you provided. Depending on the programming language you choose for your template, the directory layout may differ.

E.g. having used the `http-rust` template:

```bash
$ cd url-shortener
$ tree
rust
├── Cargo.lock
├── Cargo.toml
├── spin.toml.    <-- Spin manifest
├── src
│   └── lib.rs.   <-- Source file for the first component
└── target
```

Let's explore the `spin.toml` file. This is the Spin manifest file, which tells Spin what events should trigger what components. In this case our trigger is HTTP, for a web application, and we have only one component, at the route `/...` — a wildcard expression that matches any request sent to this application. Spin applications can consists of many components, where you can define which components that are triggered for requests on different routes.

```toml
spin_manifest_version = 2

[application]
name = "url-shortener"
version = "0.1.0"
authors = ["Caleb Schoepp <caleb.schoepp@fermyon.com>"]
description = "Shorten URLs"

[[trigger.http]]
route = "/..."
component = "url-shortener"

[component.url-shortener]
source = "target/wasm32-wasip1/release/url_shortener.wasm"
allowed_outbound_hosts = []
key_value_stores = ["default"]
[component.url-shortener.build]
command = "cargo build --target wasm32-wasip1 --release"
watch = ["src/**/*.rs", "Cargo.toml"]
```

> **Note**: You can [learn more about the Spin manifest file in the Spin documentation](https://spinframework.dev/v3/writing-apps).

### A few words about choosing a programming language

For this workshop, it doesn't matter what language you choose, as we've provided all the source code needed to complete the workshop. However, for any project you would like to endeavour on yourself, it's important to know the state of support for a giving programming language in the context ow WebAssembly, and more specifically [WASI](https://wasi.dev/).

You can find a good overview of the support on the [WebAssembly language support matrix](https://developer.fermyon.com/wasm-languages/webassembly-language-support).
The Spin documentation also provides guides for the most popular languages, which you can find here: [Spin Language Guides](https://spinframework.dev/v2/language-support-overview).

## 2. Spin Build and Spin Up - Build and run the application

You are now ready to build your application using `spin build`, which will invoke each component's `[component.MYAPP.build.command]` from `spin.toml`. This compiles your app to WebAssembly.

E.g.

```bash
spin build
Building component url-shortener with `cargo build --target wasm32-wasip1 --release`
   Compiling proc-macro2 v1.0.101
   Compiling unicode-ident v1.0.19
   ...
   Compiling wit-bindgen v0.43.0
   Compiling url-shortener v0.1.0 (/Users/caleb/.Trash/rust-old)
    Finished `release` profile [optimized] target(s) in 11.49s
Finished building all Spin components
```

> **Note**: If you are having issues building your application, refer to the [troubleshooting guide from the setup document](./00-setup.md#troubleshooting).

You can now start your application using `spin up`. This will run the WebAssembly version of your app locally with Spin acting as a WebAssembly runtime.

```bash
$ spin up
Logging component stdio to ".spin/logs/"

Serving http://127.0.0.1:3000
Available Routes:
  url-shortener: http://127.0.0.1:3000 (wildcard)
```

The command will start Spin on port `3000` (use `--listen <ADDRESS>` to change the address and port - e.g., `--listen 0.0.0.0:3030`). You can now access the application by navigating to `localhost:3000` in your browser, or by using `curl`:

```bash
$ curl localhost:3000
Hello World!
```

## 3. Writing the URL Shortener

Now that you have a Spin application running locally we can write the URL shortener. We want it to have the following functionality:

- `GET /<slug>`
  - `302` redirect to the URL configured for that slug
  - `404` if the slug was not configured
- `POST /<slug>` with a body containing a URL
  - `201` if the slug was successfully created
- Anything else
  - `404` not found

> **Note:** Remember that if you're not using Rust you'll be able to see what the code looks like in your language in the example directories.

Let's use the router from the Spin SDK to define our routes.

```rust
#[http_component]
fn handle_url_shortener(req: Request) -> anyhow::Result<impl IntoResponse> {
    let mut router = Router::new();

    router.get("/:slug", redirect);
    router.post("/:slug", shorten);

    Ok(router.handle(req))
}
```

Now when the component is invoked the handler will route the request to the appropriate function based on the HTTP method and path.

Since memory is not shared between each invocation of our component we'll need somewhere to persist the mappings between slugs and URLs. Spin provides an interface for you to persist data in a key value store managed by Spin. This key value store allows Spin developers to persist non-relational data across application invocations. The store is backed differently depending on the environment where you run Spin. Locally it is backed by a SQLite database.

To use the key value store in your application you must explicitly request access to it in your application manifest. We'll request access to the default store (you can use multiple stores in an application if you want).

```toml
[component.url-shortener]
key_value_stores = ["default"]
```

> **Note:** You can learn more about key value stores in the [Spin documentation](https://spinframework.dev/v3/kv-store-api-guide).

The Spin SDK provides a simple interface for interacting with the key value store. Let's implement the `shorten` function to store the mapping between a slug and a URL.

```rust
fn shorten(req: Request, params: Params) -> anyhow::Result<impl IntoResponse> {
    let slug = params.get("slug").unwrap();
    let url = req.body();
    let store = Store::open_default()?;
    store.set(slug, url)?;

    Ok(Response::builder().status(201).build())
}
```

Now we can implement the `redirect` function to look up the URL for a given slug and redirect to it. If it doesn't find a mapping it will return a `404`.

```rust
fn redirect(_req: Request, params: Params) -> anyhow::Result<impl IntoResponse> {
    let slug = params.get("slug").unwrap();
    let store = Store::open_default()?;

    if let Some(url) = store.get(slug)? {
        return Ok(Response::builder()
            .status(302)
            .header("Location", String::from_utf8(url)?)
            .build());
    }

    Ok(Response::builder().status(404).body("Not Found").build())
}
```

> **Note:** If you're using Python you'll need to install your dependencies in a virtual environment. Instructions [here](https://spinframework.dev/v3/python-components#system-housekeeping-use-a-virtual-environment).

> **Note:** If you're using TypeScript you'll need to `npm install @spinframework/spin-kv`.

Go ahead and rebuild the application and run it:

```bash
$ spin build --up
```

Let's test it out!

```bash
$ curl localhost:3000/foo -i
HTTP/1.1 404 Not Found

$ curl localhost:3000/foo -i --data 'https://wikipedia.org'
HTTP/1.1 201 Created

$ curl localhost:3000/foo -i
HTTP/1.1 302 Found
location: https://wikipedia.org
```

> **Note:** If you run `curl localhost:3000/foo -L` it will follow the redirect and return the HTML from Wikipedia.

Great! You've built a simple URL shortener using Spin and WebAssembly.

## Learning Summary

In this module you learned how to:

- Scaffold a new Spin application using `spin new`
- Build and run a Spin application using `spin build` and `spin up`
- Use the Spin KV store to persist data
- Build a simple URL shortener application

## Navigation

- Go back to [Setup](00-setup.md) if you still have questions about the previous section
- Otherwise, proceed to [Deploying to SpinKube](02-spinkube.md).

If you have any feedback for this module please open an issue on the GitHub repo [here](https://github.com/calebschoepp/truly-portable-code/issues/new).
