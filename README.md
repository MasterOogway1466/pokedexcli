# Pokedex CLI ⚡️🐹

A robust, REPL-style command-line interface Pokedex built in Go.

This application interacts with the [PokeAPI](https://pokeapi.co/) to allow users to explore map locations, catch Pokemon with dynamic probabilities, and inspect their stats—all while leveraging a custom-built caching layer for instant responses.

## 🚀 Features

- **REPL Architecture**: A fully interactive shell with command history (Up/Down arrow support).
- **Interactive Exploration**: Navigate through 20+ different location areas in the Pokemon world.
- **Game Mechanics**: Catch Pokemon based on their base experience stats—the stronger the Pokemon, the harder they are to catch!
- **Pokedex Management**: Inspect stats (Health, Attack, Defense) of caught Pokemon.
- **Custom Caching System**: 
    - Implements a thread-safe, in-memory cache to minimize API calls.
    - Automatic background "reaping" to expire old cache entries.
- **Robust Error Handling**: Clean feedback for invalid commands or network issues.

## 🛠 Installation & Usage

### Prerequisites
- [Go](https://go.dev/dl/) (version 1.20 or higher recommended)

### Build and Run
1. **Clone the repository:**
   ```bash
   git clone https://github.com/MasterOogway1466/pokedexcli.git
   cd pokedexcli
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```
   
3. **Run the application:**
   ```bash
   go run .
   ```

## 📖 Command Reference

| Command   | Description                             | Usage                    |
|-----------|-----------------------------------------|--------------------------|
| `help`    | Displays a help message                 | `help`                   |
| `exit`    | Exits the Pokedex                       | `exit`                   |
| `map`     | Displays the next 20 location areas     | `map`                    |
| `mapb`    | Displays the previous 20 location areas | `mapb`                   |
| `explore` | Lists all Pokémon in a specific area    | `explore <area_name>`    |
| `catch`   | Attempt to catch a Pokémon              | `catch <pokemon_name>`   |
| `inspect` | View stats of a caught Pokémon          | `inspect <pokemon_name>` |
| `pokedex` | List all caught Pokémon                 | `pokedex`                |

## 🏗 Technical Highlights

### 1. Concurrency & Caching
To ensure the application remains fast and responsive, I implemented a custom caching package (`internal/pokecache`).
* **Mutex Locking:** Used `sync.Mutex` to ensure the cache map is safe for concurrent access (thread-safe).
* **Background Garbage Collection:** A dedicated Goroutine runs a `reapLoop` that checks every 5 minutes (configurable) to remove expired entries, preventing memory bloat.

### 2. Architecture
The project follows a clean, modular structure:
* `cmd/`: Entry point of the application.
* `internal/pokeapi`: Handles all HTTP communication, JSON unmarshalling, and API logic.
* `internal/pokecache`: The isolated caching logic.
* **REPL Loop:** A central loop that parses user input, sanitizes it, and routes it to the correct handler function.

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
