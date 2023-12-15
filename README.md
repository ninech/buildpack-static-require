# buildpack-static-require

This buildpack complements the
[`buildpack-static-confgen`](https://github.com/ninech/buildpack-static-confgen)
buildpack, see the `README.md` on why this can't be in a single buildpack.
This simply requires nginx if an `index.html` can be found in the workspace or
if the env var `BP_STATIC_WEBROOT` is configured. The build phase is a noop.

To test the build locally, checkout this repository and then build it with:

```bash
pack build static --path ./integration/testdata/default_app/ \
  --builder paketobuildpacks/builder-jammy-base \
  --buildpack ninech/buildpack-static-confgen \
  --buildpack paketo-buildpacks/web-servers \
  --buildpack .
```
