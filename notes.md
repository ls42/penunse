# TODO

Formless notes about next steps.

# Tags

Tag-Strings werden seitens des API-Servers normalisiert. Damit der User ein besseres Feedback bekommt auch JS-seitig. Ich koennte eine Methode via API bereitstellen die das normalisiert.


### Include X-Clack-Overhead

Create Middleware that adds that header to all responses


### Login

Create Middleware that, in the first step, checks if a `X-Penunse-Auth-Token` header is set in the response and matches a pre-defined token. This can later be expanded to a full authentication system.
