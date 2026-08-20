# Go Practice Projects
Repository for toy projects done to learn Go

---

## Projects & Concepts 

### 1. CLI Quiz Game
* **Overview:** Reads a csv file containing questions and answers, prompts the user via the command line, and tracks score within a configurable time limit.
* **Concepts Practiced:**
  * File I/O (`os.Open`, `encoding/csv`)
  * Command-line flags (`flag` package)
  * Concurrent timer handling using channels and `select` (`time.NewTimer`)
  * Reading user input (`bufio.NewReader`, `bufio.ReadString`)

---

### 2. URL Shortener
* **Overview:** Builds an `http.Handler` that redirects short paths to target URLs using fallback handlers and routing rules parsed from JSON or YAML files.
* **Concepts Practiced:**
  * HTTP routing and handler chaining (`net/http`)
  * JSON and YAML parsing (`encoding/json`, `gopkg.in/yaml.v3`)
  * Middleware design patterns in Go