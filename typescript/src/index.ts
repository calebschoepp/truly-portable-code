import { AutoRouter, IRequest } from "itty-router";
import { openDefault } from "@spinframework/spin-kv";

const decoder = new TextDecoder();

// Any unmatched route will automatically return a 404
let router = AutoRouter();
router.get("/:slug", redirect).post("/:slug", shorten);

function redirect(req: IRequest) {
  let slug = req.params.slug;
  let store = openDefault();

  let url = store.get(slug);
  if (!url) {
    return new Response("Not Found", { status: 404 });
  }

  return new Response("", {
    status: 302,
    headers: {
      Location: decoder.decode(url),
    },
  });
}

async function shorten(req: IRequest) {
  let slug = req.params.slug;
  let url = await req.arrayBuffer();
  let store = openDefault();
  store.set(slug, new Uint8Array(url));

  return new Response("", { status: 201 });
}

//@ts-ignore
addEventListener("fetch", (event: FetchEvent) => {
  event.respondWith(router.fetch(event.request));
});
