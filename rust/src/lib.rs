use spin_sdk::http::{IntoResponse, Params, Request, Response, Router};
use spin_sdk::http_component;
use spin_sdk::key_value::Store;

#[http_component]
fn handle_url_shortener(req: Request) -> anyhow::Result<impl IntoResponse> {
    let mut router = Router::new();

    router.get("/:slug", redirect);
    router.post("/:slug", shorten);

    Ok(router.handle(req))
}

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

fn shorten(req: Request, params: Params) -> anyhow::Result<impl IntoResponse> {
    let slug = params.get("slug").unwrap();
    let url = req.body();
    let store = Store::open_default()?;
    store.set(slug, url)?;

    Ok(Response::builder().status(201).build())
}
