package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"github.com/goccy/go-yaml"
)

type FileEntry struct {
	RelativePath       string   `yaml:"relative_path"`
	ModifiedEntityKeys []string `yaml:"modified_entity_keys"`
}

type Config struct {
	GameRoot            string      `yaml:"game_root"`
	Files               []FileEntry `yaml:"files"`
	CustomCommentPrefix string      `yaml:"custom_comment_prefix"`
}

func loadConfig(configPath string) (*Config, error) {
	f, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var cfg Config
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}
	if cfg.CustomCommentPrefix == "" {
		cfg.CustomCommentPrefix = "# QUIETER:"
	}
	return &cfg, nil
}

func extractEvents(lines []string, modKeys []string, commentPrefix string) map[string][]string {
	events := make(map[string][]string)
	eventRe := regexp.MustCompile(`^\s*([a-zA-Z0-9_.]+)\s*=\s*{`)
	for i := 0; i < len(lines); {
		line := lines[i]
		m := eventRe.FindStringSubmatch(line)
		if m != nil {
			key := m[1]
			if slices.Contains(modKeys, key) {
				block := []string{}
				// Collect comments above
				for j := i - 1; j >= 0 && strings.HasPrefix(strings.TrimSpace(lines[j]), commentPrefix); j-- {
					block = append([]string{lines[j]}, block...)
				}
				depth := strings.Count(line, "{") - strings.Count(line, "}")
				block = append(block, line)
				i++
				for i < len(lines) && depth > 0 {
					block = append(block, lines[i])
					depth += strings.Count(lines[i], "{")
					depth -= strings.Count(lines[i], "}")
					i++
				}
				events[key] = block
				continue
			}
		}
		i++
	}
	return events
}

func merge(vanilla, mod []string, modKeys []string, commentPrefix string) []string {
	modEvents := extractEvents(mod, modKeys, commentPrefix)
	eventRe := regexp.MustCompile(`^\s*([a-zA-Z0-9_.]+)\s*=\s*{`)
	var out []string
	for i := 0; i < len(vanilla); {
		line := vanilla[i]
		m := eventRe.FindStringSubmatch(line)
		if m != nil {
			key := m[1]
			if block, ok := modEvents[key]; ok {
				out = append(out, block...)
				// Skip vanilla event block
				depth := strings.Count(line, "{") - strings.Count(line, "}")
				i++
				for i < len(vanilla) && depth > 0 {
					depth += strings.Count(vanilla[i], "{")
					depth -= strings.Count(vanilla[i], "}")
					i++
				}
				continue
			}
		}
		out = append(out, line)
		i++
	}
	return out
}

func readLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var lines []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		lines = append(lines, sc.Text()+"\n")
	}
	return lines, sc.Err()
}

func writeLines(path string, lines []string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, line := range lines {
		if _, err := f.WriteString(line); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: vanilla_merger <config.yaml>")
		return
	}
	cfg, err := loadConfig(os.Args[1])
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	// TODO: Feigure out this relative path issue when debugging vs running binary from different dirs. Maybe move to config?
	modRoot, _ := filepath.Abs(filepath.Join(filepath.Dir(os.Args[0]), ".."))
	for _, fileEntry := range cfg.Files {
		relPath := fileEntry.RelativePath
		vanilla, _ := readLines(filepath.Join(cfg.GameRoot, relPath))
		mod, _ := readLines(filepath.Join(modRoot, relPath))
		keys := fileEntry.ModifiedEntityKeys
		merged := merge(vanilla, mod, keys, cfg.CustomCommentPrefix)
		if err := writeLines(filepath.Join(modRoot, "vanilla-mod-merger/merge-output", filepath.Base(relPath)), merged); err != nil {
			fmt.Println("Error writing merged file:", err)
			return
		}
	}
}
