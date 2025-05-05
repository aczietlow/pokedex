# Pokedex

A pokemon api wrapper, pretending to be a pokedex, via the cli

## Useage

To start:
`go run .`

Commands:

pokedex: Consult your pokedex
- returns a list of all pokemon stored in the local pokedex

exit: Exit the Pokedex
- exits the application

help: Displays a help message
- Displays help 

map: List map locations, 20 at a time
- Query a list of 20 map locations, paginated

mapb: Fetch the previous 20 locations
- Quert the previous 20 map locations; subsequent api calls are cached

explore {name-or-id}: Fetch data on a specific area
- Fetch details about a map location

catch {name-or-id}: catch and add pokemon to pokedex
- Chance at catching a pokemon. This app requires a pokemon to be caught before adding them to the pokedex

inspect {name-or-id}: Displays information of pokemon registered in the pokedex
- Displays information about a previously caught pokemon. 
