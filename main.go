package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
	"net/http"
)

// ── Colors ────────────────────────────────────────────────────────────────────
const (
	ColorRed   = "\033[31m"
	ColorGreen = "\033[32m"
	ColorReset = "\033[0m"
)

// ── Checkers ──────────────────────────────────────────────────────────────────

// CheckWeb handles HTTP/HTTPS URLs
func CheckWeb(link string, c chan string) {
	res, err := http.Get(link)
	if err != nil {
		c <- fmt.Sprintf("❌ %s ha fallado", link)
		return
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		c <- fmt.Sprintf("✅ %s está OK! (HTTP %d)", link, res.StatusCode)
	} else {
		c <- fmt.Sprintf("⚠️  %s devolvió código %d", link, res.StatusCode)
	}
}

// CheckGeneric handles raw IPs and host:port targets via TCP
func CheckGeneric(target string, c chan string) {
	conn, err := net.DialTimeout("tcp", target, 2*time.Second)
	if err != nil {
		c <- fmt.Sprintf("%s❌ %s está inalcanzable%s", ColorRed, target, ColorReset)
		return
	}
	defer conn.Close()
	c <- fmt.Sprintf("%s✅ %s conexión exitosa!%s", ColorGreen, target, ColorReset)
}

// ── Input loaders ─────────────────────────────────────────────────────────────

// loadFromEnv reads the URLS environment variable (comma-separated)
// Example: URLS="https://google.com,https://myapp.com"
func loadFromEnv() []string {
	raw := os.Getenv("URLS")
	if raw == "" {
		return nil
	}
	return splitAndClean(raw, ",")
}

// loadFromFile reads targets line by line from a file
// Skips empty lines and lines starting with #
func loadFromFile(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error al leer %s: %w", path, err)
	}
	return splitAndClean(string(data), "\n"), nil
}

// splitAndClean splits a string by sep, trims spaces, and drops empty/comment lines
func splitAndClean(input, sep string) []string {
	var result []string
	for _, item := range strings.Split(input, sep) {
		item = strings.TrimSpace(item)
		if item != "" && !strings.HasPrefix(item, "#") {
			result = append(result, item)
		}
	}
	return result
}

// ── Dispatcher ────────────────────────────────────────────────────────────────

// isHTTP returns true if the target looks like an HTTP/HTTPS URL
func isHTTP(target string) bool {
	return strings.HasPrefix(target, "http://") || strings.HasPrefix(target, "https://")
}

// run executes concurrent checks and prints results
func run(targets []string) {
	if len(targets) == 0 {
		fmt.Println("⚠️  No hay targets. Usa URLS='...' o monta un archivo urls.txt")
		return
	}

	fmt.Printf("\n🔍 Verificando %d target(s)...\n\n", len(targets))

	c := make(chan string, len(targets))

	for _, target := range targets {
		if isHTTP(target) {
			go CheckWeb(target, c)
		} else {
			go CheckGeneric(target, c)
		}
	}

	for range targets {
		fmt.Println(<-c)
	}

	fmt.Println()
}

// ── Main ──────────────────────────────────────────────────────────────────────
func main() {
	var targets []string

	// Priority 1: ENV variable
	// docker run -e URLS="https://google.com,https://myapp.com" instantcheck
	if envTargets := loadFromEnv(); len(envTargets) > 0 {
		fmt.Println("📡 Fuente: variable de entorno URLS")
		targets = envTargets

	// Priority 2: FILE (default path or custom via FILE_PATH env)
	// docker run -v $(pwd)/urls.txt:/app/urls.txt instantcheck
	} else {
		filePath := os.Getenv("FILE_PATH")
		if filePath == "" {
			filePath = "urls.txt"
		}

		if fileTargets, err := loadFromFile(filePath); err == nil {
			fmt.Printf("📂 Fuente: archivo %s\n", filePath)
			targets = fileTargets
		} else {
			// Priority 3: fallback for local development
			fmt.Println("🛠️  Fuente: lista de desarrollo (hardcoded)")
			targets = []string{
				"https://www.google.com",
				"https://www.youtube.com",
				"https://discord.com",
				"1.1.1.1:80",        // Cloudflare DNS via TCP
				"http://github.com", // 301 redirect
				"https://httpbin.org/status/404",
				"https://httpbin.org/status/500",
			}
		}
	}

	run(targets)
}
