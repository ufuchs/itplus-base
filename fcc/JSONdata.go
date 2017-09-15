package fcc

var aletaJSON = map[string]string{
	"2017-05-15T202418Z": `{
      "name": "aleta",
      "modifiedOn": "2017-05-15T202418Z",
      "outputscheme": {
          "show_unknown_devices": false,
          "interval": 60
      },
      "devices": [
          {
              "id": 1,
              "num": 200,
              "addr": 7,
              "type": "TX29TDH-IT",
              "alias": "R 2.7 - Schreibtisch",
              "absent": false,
              "locked": false,
              "lon": 19.31,
              "lat": 19.32,
              "alt": 19.33
          },
          {
              "id": 2,
              "num": 194,
              "addr": 44,
              "type": "TX29TDH-IT",
              "alias": "R 2.7 - Kühlschrank",
              "absent": false,
              "locked": false,
              "lon": 19.41,
              "lat": 19.42,
              "alt": 19.43
          },
          {
              "id": 3,
              "num": 195,
              "addr": 48,
              "type": "TX29TDH-IT",
              "alias": "R 2.7 - Wand",
              "absent": false,
              "locked": false,
              "lon": 19.51,
              "lat": 19.52,
              "alt": 19.53
          },
          {
              "id": 4,
              "num": 7,
              "addr": 24,
              "type": "TX25TP-IT",
              "alias": "R 2.7 - Fenster:Draußen - AUCOTEAM Nordseite",
              "absent": false,
              "locked": false,
              "lon": 7.1,
              "lat": 7.2,
              "alt": 7.3
          }
      ]
  }
  `,
}

var HomePureJSON = `{
  "name": "home",
  "modifiedon": "2017-05-18T160934Z",
  "outputscheme": {
    "show_unknown_devices": true,
    "interval": 30
  },
  "devices": [
    {
      "id": 1,
      "num": 0,
      "addr": 0,
      "type": "JeeLink",
      "alias": "Schreibtisch",
      "absent": false,
      "locked": true,
      "lon": 0,
      "lat": 0,
      "alt": 0
    },
    {
      "id": 2,
      "num": 1,
      "addr": 46,
      "type": "TX29TDH-IT",
      "alias": "Bad",
      "absent": false,
      "locked": false,
      "lon": 1.1,
      "lat": 1.2,
      "alt": 1.3
    },
    {
      "id": 3,
      "num": 2,
      "addr": 42,
      "type": "TX29TDH-IT",
      "alias": "Küche",
      "absent": false,
      "locked": false,
      "lon": 2.1,
      "lat": 2.2,
      "alt": 2.3
    },
    {
      "id": 4,
      "num": 3,
      "addr": 11,
      "type": "TX29TDH-IT",
      "alias": "Kühlschrank",
      "absent": false,
      "locked": false,
      "lon": 3.1,
      "lat": 3.2,
      "alt": 3.3
    },
    {
      "id": 5,
      "num": 4,
      "addr": 50,
      "type": "TX29TDH-IT",
      "alias": "Schlafzimmer",
      "absent": false,
      "locked": false,
      "lon": 4.1,
      "lat": 4.2,
      "alt": 4.3
    },
    {
      "id": 6,
      "num": 5,
      "addr": 57,
      "type": "TX29TDH-IT",
      "alias": "Wohnzimmer",
      "absent": false,
      "locked": false,
      "lon": 5.1,
      "lat": 5.2,
      "alt": 5.3
    },
    {
      "id": 7,
      "num": 6,
      "addr": 12,
      "type": "TX29TDH-IT",
      "alias": "Bei Mama",
      "absent": true,
      "locked": false,
      "lon": 6.1,
      "lat": 6.2,
      "alt": 6.3
    },
    {
      "id": 8,
      "num": 7,
      "addr": 24,
      "type": "TX25TP-IT",
      "alias": "Bei Aucoteam",
      "absent": false,
      "locked": false,
      "lon": 7.1,
      "lat": 7.2,
      "alt": 7.3
    },
    {
      "id": 9,
      "num": 8,
      "addr": 6,
      "type": "TX29TDH-IT",
      "alias": "Other-08",
      "absent": false,
      "locked": false,
      "lon": 8.1,
      "lat": 8.2,
      "alt": 8.3
    },
    {
      "id": 10,
      "num": 9,
      "addr": 14,
      "type": "TX29TDH-IT",
      "alias": "Other-09",
      "absent": false,
      "locked": false,
      "lon": 9.1,
      "lat": 9.2,
      "alt": 9.3
    },
    {
      "id": 11,
      "num": 10,
      "addr": 43,
      "type": "TX29TDH-IT",
      "alias": "Other-10",
      "absent": false,
      "locked": false,
      "lon": 10.1,
      "lat": 10.2,
      "alt": 10.3
    },
    {
      "id": 12,
      "num": 11,
      "addr": 28,
      "type": "TX29TDH-IT",
      "alias": "Other-11",
      "absent": false,
      "locked": false,
      "lon": 11.1,
      "lat": 11.2,
      "alt": 11.3
    },
    {
      "id": 13,
      "num": 12,
      "addr": 26,
      "type": "TX29TDH-IT",
      "alias": "Other-12",
      "absent": false,
      "locked": false,
      "lon": 12.1,
      "lat": 12.2,
      "alt": 12.3
    }
  ]
}
`

var HomePureJSON_Half_1 = `{
    "name": "home",
    "modifiedon": "2017-05-18T160934Z",
    "outputscheme": {
        "show_unknown_devices": true,
        "interval": 30
    },
    "devices": [
        {
            "id": 1,
            "num": 0,
            "addr": 0,
            "type": "JeeLink",
            "alias": "Schreibtisch",
            "absent": false,
            "locked": true,
            "lon": 0,
            "lat": 0,
            "alt": 0
        },
        {
            "id": 2,
            "num": 1,
            "addr": 46,
            "type": "TX29TDH-IT",
            "alias": "Bad",
            "absent": false,
            "locked": false,
            "lon": 1.1,
            "lat": 1.2,
            "alt": 1.3
        },
        {
            "id": 3,
            "num": 2,
            "addr": 42,
            "type": "TX29TDH-IT",
            "alias": "Küche",
            "absent": false,
            "locked": false,
            "lon": 2.1,
            "lat": 2.2,
            "alt": 2.3
        },
        {
            "id": 4,
            "num": 3,
            "addr": 11,
            "type": "TX29TDH-IT",
            "alias": "Kühlschrank",
            "absent": false,
            "locked": false,
            "lon": 3.1,
            "lat": 3.2,
            "alt": 3.3
        },
        {
            "id": 5,
            "num": 4,
            "addr": 50,
            "type": "TX29TDH-IT",
            "alias": "Schlafzimmer",
            "absent": false,
            "locked": false,
            "lon": 4.1,
            "lat": 4.2,
            "alt": 4.3
        },
        {
            "id": 6,
            "num": 5,
            "addr": 57,
            "type": "TX29TDH-IT",
            "alias": "Wohnzimmer",
            "absent": false,
            "locked": false,
            "lon": 5.1,
            "lat": 5.2,
            "alt": 5.3
        },
        {
            "id": 7,
            "num": 6,
            "addr": 12,
            "type": "TX29TDH-IT",
            "alias": "Bei Mama",
            "absent": true,
            "locked": false,
            "lon": 6.1,
            "lat": 6.2,
            "alt": 6.3
        }
    ]
}
`

var HomePureJSON_Half_2 = `{
  "name": "home",
  "modifiedon": "2017-05-18T160934Z",
  "outputscheme": {
    "show_unknown_devices": true,
    "interval": 30
  },
  "devices": [
    {
      "id": 8,
      "num": 7,
      "addr": 24,
      "type": "TX25TP-IT",
      "alias": "Bei Aucoteam",
      "absent": false,
      "locked": false,
      "lon": 7.1,
      "lat": 7.2,
      "alt": 7.3
    },
    {
      "id": 9,
      "num": 8,
      "addr": 6,
      "type": "TX29TDH-IT",
      "alias": "Other-08",
      "absent": false,
      "locked": false,
      "lon": 8.1,
      "lat": 8.2,
      "alt": 8.3
    },
    {
      "id": 10,
      "num": 9,
      "addr": 14,
      "type": "TX29TDH-IT",
      "alias": "Other-09",
      "absent": false,
      "locked": false,
      "lon": 9.1,
      "lat": 9.2,
      "alt": 9.3
    },
    {
      "id": 11,
      "num": 10,
      "addr": 43,
      "type": "TX29TDH-IT",
      "alias": "Other-10",
      "absent": false,
      "locked": false,
      "lon": 10.1,
      "lat": 10.2,
      "alt": 10.3
    },
    {
      "id": 12,
      "num": 11,
      "addr": 28,
      "type": "TX29TDH-IT",
      "alias": "Other-11",
      "absent": false,
      "locked": false,
      "lon": 11.1,
      "lat": 11.2,
      "alt": 11.3
    },
    {
      "id": 13,
      "num": 12,
      "addr": 26,
      "type": "TX29TDH-IT",
      "alias": "Other-12",
      "absent": false,
      "locked": false,
      "lon": 12.1,
      "lat": 12.2,
      "alt": 12.3
    }
  ]
}
`




