module github.com/nevnet99/pokedex-cli

go 1.22.2

require internal/commands v1.0.0

replace internal/commands v1.0.0 => ./internal/commands

require internal/types v1.0.0

replace internal/types v1.0.0 => ./internal/types

require internal/pokecache v1.0.0

replace interal/commands v1.0.0 => ./internal/pokecache
