export default {
    async fetch(req, env, ctx) {
        const target = new URL(req.url);

        // Overwrite host
        target.host = env.HOST;

        // Add prefix if revision tag is provided
        const tag = req.headers.get("X-Revision-Tag");
        if (tag) {
            target.host = `${tag}---${target.host}`;
        }

        // Proxy the request
        let response = await fetch(target, req);

        // Stream the response
        let { readable, writable } = new TransformStream();
        response.body.pipeTo(writable).then();

        return new Response(readable, response);
    }
}
