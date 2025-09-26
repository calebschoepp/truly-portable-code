# Deploying to Fermyon Wasm Functions

- [Deploying to Fermyon Wasm Functions](#deploying-to-fermyon-wasm-functions)
  - [1. Install `aka` Plugin for Spin](#1-install-aka-plugin-for-spin)
  - [2. Deploy the URL Shortener to Fermyon Wasm Functions](#2-deploy-the-url-shortener-to-fermyon-wasm-functions)
  - [Learning Summary](#learning-summary)
  - [Navigation](#navigation)

In this module we'll explore how to deploy our URL shortening Spin application to Fermyon Wasm Functions.

Fermyon Wasm Functions is a globally distributed PaaS (Platform-as-a-Service) running edge native applications with Spin on top of Akamai’s Connected Cloud. It offers fast and resilient hosting for Spin applications. This means that Fermyon Wasm Functions is tailored towards hosting applications, which need fast execution, and availability in multiple regions across the World. A perfect fit for our URL shortener application.

Fermyon Wasm Functions does not require any operational effort on infrastructure from you as a user of the platform. Upon deploying a Spin application to the platform, the application is automatically distributed and made available in multiple regions within the service. A URL is provided for the application, which will persist across deployments.

## 1. Install `aka` Plugin for Spin

To deploy to Fermyon Wasm Functions you'll need to install the `aka` plugin for Spin. You can do so by running:

```bash
$ spin plugins install aka
```

Run the following command to verify the plugin is installed:

```bash
$ spin aka -V
```

## 2. Deploy the URL Shortener to Fermyon Wasm Functions

Once you’ve installed the aka plugin for Spin, you must log in to Fermyon Wasm Functions, which requires your GitHub account to sign in:

```bash
$ spin aka login
Go to https://login.infra.fermyon.tech/realms/neutrino/device?user_code=AAAA-BBBB and follow the prompts.

Don't worry, we'll wait here for you. You got this.
```

Your Github account needs to have been given access to Fermyon Wasm Functions before logging in. If you haven't done this already see the instructions in the [Setup](00-setup.md#request-access-to-fermyon-wasm-functions) module.

Now deploying is as easy as running:

```bash
$ spin aka deploy
Name of new app: url-shortener
Creating new app url-shortener in account calebschoepp
Note: If you would instead like to deploy to an existing app, cancel this deploy and link this workspace to the app with `spin aka app link`
OK to continue? yes
Workspace linked to app url-shortener
Waiting for app to be ready... ready

App Routes:
- url-shortener: <App URL> (wildcard)
```

It will ask you for a name for the application (defaulting to `url-shortener`) and then it will ask for confirmation that you want to deploy.

> **Note:** Deploying a Spin application to Fermyon Wasm Functions includes packaging the application and all the required files, uploading it to an OCI registry, as well as instantiating the application on Fermyon Wasm Functions.

Now let's test it out:

```bash
$ curl <App URL>/foo -i
HTTP/1.1 404 Not Found

$ curl <App URL>/foo -i --data 'https://wikipedia.org'
HTTP/1.1 201 Created

$ curl <App URL>/foo -i
HTTP/1.1 302 Found
location: https://wikipedia.org
```

Feel free to open the URL in your browser as well.

Notice how we didn't need to configure anything for key value to work. Fermyon Wasm Functions provisions and manages the key value store on your behalf, handling the heavy lifting for you. It is low-latency, persistent, and globally replicated with read-your-writes behavior within a request.

> **Note:** Find more on the KV store [here](https://developer.fermyon.com/wasm-functions/using-key-value-store).

## Learning Summary

In this module you learned how to:

- Use the `aka` plugin for Spin
- Deploy a Spin application to Fermyon Wasm Functions

## Navigation

- Go back to [Deploying to SpinKube](02-spinkube.md) if you still have questions about the previous section
- Otherwise, proceed to [Bonus Exercises](04-bonus.md).

If you have any feedback for this module please open an issue on the GitHub repo [here](https://github.com/calebschoepp/truly-portable-code/issues/new).
