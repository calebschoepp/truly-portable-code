Resources:

# Bonus Exercises

- [Bonus Exercises](#bonus-exercises)
- [1. Spin Exercises](#1-spin-exercises)
- [2. SpinKube Exercises](#2-spinkube-exercises)
- [3. Fermyon Wasm Function Exercises](#3-fermyon-wasm-function-exercises)
- [Navigation](#navigation)

The core content of this workshop is complete. This module will provide some bonus exercises for how you expand your knowledge of the technologies we've worked with today.

## 1. Spin Exercises

- Protect the shorten endpoint with authentication.
- Add validation to the shorten endpoint to ensure the provided URL is valid.
- Add a delete endpoint to allow users to delete a shortened URL.
- Add a list endpoint to allow users to list all their shortened URLs.

You may find the [Spin documentation](https://spinframework.dev/) useful for these exercises.

## 2. SpinKube Exercises

- Setup autoscaling so that the Spin app can scale up and down based on demand ([hint](https://www.spinkube.dev/docs/topics/autoscaling/)).
- Setup an Ingress controller to provide a custom domain for the Spin app.
- Monitor your application with OpenTelemetry traces ([hint](https://www.spinkube.dev/docs/topics/monitoring-your-app/)).

You may find the [SpinKube documentation](https://www.spinkube.dev/docs/) useful for these exercises.

## 3. Fermyon Wasm Function Exercises

- Introduce a time to live (TTL) for shortened URLs so that they expire after a certain amount of time. Create a cron job that periodically cleans up expired URLs ([hint](https://developer.fermyon.com/wasm-functions/using-cron-jobs)).

You may find the [Fermyon Wasm Functions documentation](https://developer.fermyon.com/wasm-functions/index) useful for these exercises.

## Navigation

- Go back to [Deploying to Fermyon Wasm Functions](03-fwf.md) if you still have questions about the previous section
- Otherwise, you're all done.

If you have any feedback for this module please open an issue on the GitHub repo [here](https://github.com/calebschoepp/truly-portable-code/issues/new).
