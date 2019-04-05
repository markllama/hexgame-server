## Map API queries

### Query all maps [GET]

Request a list of all stored maps

+ Request (application/json)

+ Response 200 (application/json)

    + Headers
    
        Location: /maps
        
    + Body
      
          [
            {
                "id": "1",
                "gameId": "1",
                "name": "samplemap",
                "description": "A Sample Map",
                "size": { "hx": 15, "hy": 22 },
                "origin": { "hx": 0, "hy": 0 },
                "shape": "sawtooth",
                "terrains": [ ]
            },
            {
                "id": "2",
                "gameId": "2",
                "name": "meleemap",
                "description": "A Melee Map",
                "size": { "hx": 3, "hy": 5 },
                "origin": { "hx": 0, "hy": 0 },
                "shape": "megahex",
                "terrains": [ ]
            },
            {
                "id": "3",
                "gameId": "2",
                "name": "wizardmap",
                "description": "A Wizard Map",
                "size": { "hx": 5, "hy": 5 },
                "origin": { "hx": 0, "hy": 0 },
                "shape": "megahex",
                "terrains": [ ]
            },
          ]

### Query a specified map by ID [GET]

Request a single map by id

+ Request (application/json)

+ Response 200 (application/json)

    + Headers
    
        Location: /maps/1

    + Body
    
            {
                "id": "1",
                "gameId": "1",
                "name": "samplemap",
                "description": "A Sample Map",
                "shape": "sawtooth",
                "size": { "hx": 15, "hy": 22 },
                "origin": { "hx": 0, "hy": 0 },
                "terrains": [ ]
            }


### Create a new map [POST]

+ Request (application/json)

        {
            "gameId": "1",
            "name": "newmap",
            "description": "A new map",
            "shape": "sawtooth",
            "size": { "hx": 15, "hy": 22 },
            "origin": { "hx": 0, "hy": 0 },
            "terrains": [ ]
        }

+ Response 201 (application/json)

    + Headers
    
        Location: /maps/{id}
        
    + Body
    
            {
                "id": "1"
                "gameId": "1",
                "name": "newmap",
                "description": "A new map",
                "shape": "sawtooth",
                "size": { "hx": 15, "hy": 22 },
                "origin": { "hx": 0, "hy": 0 },
                "terrains": [ ]
            }
            

