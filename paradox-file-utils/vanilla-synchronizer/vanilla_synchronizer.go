package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"paradox-file-utils/parser"

	"github.com/antlr4-go/antlr/v4"
	"github.com/goccy/go-yaml"
)

var (
	EVENT_KEY_PATTERN = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*\.\d+$`)
	SYNC_OUTPUT_DIR   = "sync-output"
)

type Config struct {
	GameRoot            string `yaml:"game_root"`
	ModRoot             string `yaml:"mod_root"`
	CustomCommentPrefix string `yaml:"custom_comment_prefix"`
	Files               []struct {
		RelativePath       string   `yaml:"relative_path"`
		ModifiedEntityKeys []string `yaml:"modified_entity_keys"`
	} `yaml:"files"`
}

func loadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func collectEvents(tree *parser.FileContext) map[string]*parser.AssignmentContext {
	events := make(map[string]*parser.AssignmentContext)
	for _, line := range tree.AllLine() {
		if stmt := line.Statement(); stmt != nil {
			if assign := stmt.Assignment(); assign != nil {
				key := assign.Key().GetText()
				if assign.Value().Block() != nil && EVENT_KEY_PATTERN.MatchString(key) {
					events[key] = assign.(*parser.AssignmentContext)
				}
			}
		}
	}
	return events
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./vanilla_synchronizer <config.yaml>")
		return
	}
	cfg, err := loadConfig(os.Args[1])
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	for _, f := range cfg.Files {
		relPath := f.RelativePath
		modKeys := make(map[string]struct{})
		for _, k := range f.ModifiedEntityKeys {
			modKeys[k] = struct{}{}
		}

		// Parse vanilla file
		vanillaPath := filepath.Join(cfg.GameRoot, relPath)
		vanillaInput, _ := antlr.NewFileStream(vanillaPath)
		vanillaLexer := parser.NewParadoxLexer(vanillaInput)
		vanillaStream := antlr.NewCommonTokenStream(vanillaLexer, antlr.TokenDefaultChannel)
		vanillaParser := parser.NewParadoxParser(vanillaStream)
		vanillaTree := vanillaParser.File()

		// Parse mod file
		modPath := filepath.Join(cfg.ModRoot, relPath)
		modInput, _ := antlr.NewFileStream(modPath)
		modLexer := parser.NewParadoxLexer(modInput)
		modStream := antlr.NewCommonTokenStream(modLexer, antlr.TokenDefaultChannel)
		modParser := parser.NewParadoxParser(modStream)
		modTree := modParser.File()

		// Collect modded events
		modEvents := collectEvents(modTree.(*parser.FileContext))

		// Prepare output file
		outPath := SYNC_OUTPUT_DIR + "/" + filepath.Base(relPath)
		out, err := os.Create(outPath)
		if err != nil {
			fmt.Println("Error writing merged file:", err)
			return
		}

		// Walk vanilla lines, swap event blocks if needed
		for _, line := range vanillaTree.AllLine() {
			if stmt := line.Statement(); stmt != nil {
				if assign := stmt.Assignment(); assign != nil {
					key := assign.Key().GetText()
					if _, ok := modKeys[key]; ok {
						if modAssign, ok := modEvents[key]; ok {
							if _, err := out.WriteString(modAssign.GetText()); err != nil {
								fmt.Println("Error writing to output file:", err)
								out.Close()
								return
							}
							continue
						}
					}
					if _, err := out.WriteString(assign.GetText()); err != nil {
						fmt.Println("Error writing to output file:", err)
						out.Close()
						return
					}
					continue
				}
			}
			// Write comments/whitespace as-is
			if c := line.COMMENT(); c != nil {
				if _, err := out.WriteString(c.GetText()); err != nil {
					fmt.Println("Error writing to output file:", err)
					out.Close()
					return
				}
			}
			if ws := line.WS(); ws != nil {
				if _, err := out.WriteString(ws.GetText()); err != nil {
					fmt.Println("Error writing to output file:", err)
					out.Close()
					return
				}
			}
		}
		fmt.Println("Merged file written to:", outPath)
		defer out.Close()
	}
}
