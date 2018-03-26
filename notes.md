# TODO

Formless notes about next steps.

### `/api/insert` als route erstellen.

In dem `HandlFunc`-Block werden dann die JSON-Daten ausgelesen, die JavaScript-seitig (oder erstmal mit curl) mitgegeben werden. JSON wird im Body uebermittelt, und enthaelt folgende Info:

* UserID
* Amount (Menge die ausgegeben wird)
* Tags (Array von strings)
* Description (bspw. "eating at 'elrado house'")

Beispiel:

```json
{
	"user_id": 1,
	"amount": 15.00,
	"tags": ["eating out", "lunch", "social"],
	"description": "Eating with Felix at Lyst"
}
```

Tag-Strings werden seitens des API-Servers normalisiert. Damit der User ein besseres Feedback bekommt auch JS-seitig. Ich koennte eine Methode via API bereitstellen die das normalisiert.


### Pruefen ob Eintraege in der Datenbank vorhanden sind

Bevor irgendwas damit gemacht wird muss sichergestellt werden, dass auch schon Daten darin vorhanden wird. Ansonsten gibt man bspw. ein leeres JSON zurueck. Falls ich das nicht tue, bekomme ich den Fehler `2018/03/26 19:23:17 http: panic serving [::1]:53615: runtime error: invalid memory address or nil pointer dereference`.


### Include X-Clack-Overhead

Create Middleware that adds that header to all responses


### Login

Create Middleware that, in the first step, checks if a `X-Penunse-Auth-Token` header is set in the response and matches a pre-defined token. This can later be expanded to a full authentication system.
