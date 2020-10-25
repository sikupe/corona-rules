# Corona-Rules
Mit diesem Projekt soll ein Überblick über die aktuelle Regelungslage je Bundesland gegeben werden.

Dazu wird aktuell der Gerätestandort bestimmt, um anhand der Inzidenz im Landkreis in dem man sich aufhält, die geltenden Regeln zu bestimmen.

Dabei werden aktuell nicht Einzelregelungen für/von Landkreisen beachtet, die nicht auf der allgemeinen Landesverordnung beruhen (z.B. Lockdown im Berchtesgadener Land).

## Regelungen hinzufügen bzw. aktualisieren
Um Regelungen hinzuzufügen oder zu aktualisieren muss die Datei `./frontend/src/assets/restrictions.json` bearbeitet werden.

Das verwendete Format sieht folgendermaßen aus:
```json
[
  {
    "land": "Hessen", // Land der Verordnung
    "source": "http://meine-corona-regel.de", // Quelle der Regeln
    "name": "FNP", // Name der Quelle
    "stand": "2020-10-25", // Stand der Quelle/Datum der Abfrage
    "restrictions": {
      "0": [ // Inzidenz ab der die Regeln gelten
        { // Regel-"Karte": Gruppierung nach bspw. Veranstaltungsregelungen, Maskenpflicht
          "title": "Veranstaltungen", // Überschrift der Regel-"Karte"
          "description": [ // Einzelne Regeln. Ein Arrayeintrag wird als ein Listenpunkt dargestellt
            "Öffentliche Veranstaltungen bis maximal 250 Personen",
            "Private Veranstaltungen im öffentlichen Raum bis maximal 50 Personen"
          ]
        },
        { // Weitere Regel-"Karte"
          "title": "Maskenpflicht",
          "description": [
            "Keine Maskenpflicht vorgeschrieben!"
          ]
        }
      ],
      "35": [ // Diese Regeln werden ab einer Inzidenz von 35 in einem Hessischen Landkreis angezeigt werden
        // Weitere Regeln
      ]
    }
  },
  {
    // Regeln eines weiteren Bundeslandes im gleichen Schema
  }
]
``` 