from spin_sdk.http import IncomingHandler, Request, Response
from spin_sdk import key_value
import re


class IncomingHandler(IncomingHandler):
    def handle_request(self, request: Request) -> Response:
        match = re.match(r"^/([a-zA-Z0-9]+)$", request.uri)
        if not match:
            return Response(
                404, {"content-type": "text/plain"}, bytes("Not Found", "utf-8")
            )
        slug = match.group(1)

        if request.method == "GET":
            return redirect(request, slug)
        elif request.method == "POST":
            return shorten(request, slug)
        else:
            return Response(
                405,
                {"content-type": "text/plain"},
                bytes("Method Not Allowed", "utf-8"),
            )


def redirect(req: Request, slug: str) -> Response:
    store = key_value.open_default()

    url = store.get(slug)
    if url is None:
        return Response(
            404, {"content-type": "text/plain"}, bytes("Not Found", "utf-8")
        )

    return Response(302, {"Location": url.decode("utf-8")}, bytes("", "utf-8"))


def shorten(req: Request, slug: str) -> Response:
    store = key_value.open_default()
    url = req.body
    store.set(slug, url)

    return Response(201, {}, bytes("", "utf-8"))
