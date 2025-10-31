import os
import re
import yaml

CUSTOM_COMMENT_PREFIX = "# QUIETER:"

def load_config(config_path):
    with open(config_path, 'r') as f:
        return yaml.safe_load(f)

def extract_modded_events_and_comments(mod_lines, custom_event_keys):
    modded_events = {}
    custom_comments = []
    i = 0
    while i < len(mod_lines):
        line = mod_lines[i]
        # Event block
        m = re.match(r'^\s*([a-zA-Z0-9_.]+)\s*=\s*{', line)
        if m:
            event_key = m.group(1)
            if event_key in custom_event_keys:
                # Check for QUIETER comments immediately above
                event_block = []
                j = i - 1
                while j >= 0 and mod_lines[j].strip().startswith(CUSTOM_COMMENT_PREFIX):
                    event_block.insert(0, mod_lines[j])
                    j -= 1
                # Collect the event block itself
                depth = line.count('{') - line.count('}')
                event_block.append(line)
                i += 1
                while i < len(mod_lines) and depth > 0:
                    event_block.append(mod_lines[i])
                    depth += mod_lines[i].count('{')
                    depth -= mod_lines[i].count('}')
                    i += 1
                modded_events[event_key] = event_block
                continue
        # Standalone QUIETER comments (not above a modded event)
        if line.strip().startswith(CUSTOM_COMMENT_PREFIX):
            # Only add if not immediately above a modded event
            if i+1 >= len(mod_lines) or not re.match(r'^\s*([a-zA-Z0-9_.]+)\s*=\s*{', mod_lines[i+1]):
                custom_comments.append(line)
        i += 1
    return modded_events, custom_comments

def merge_file(vanilla_path, mod_path, output_path, custom_event_keys):
    with open(vanilla_path, 'r') as f:
        vanilla_lines = f.readlines()
    with open(mod_path, 'r') as f:
        mod_lines = f.readlines()

    modded_events, custom_comments = extract_modded_events_and_comments(mod_lines, custom_event_keys)

    merged_lines = []
    i = 0
    while i < len(vanilla_lines):
        m = re.match(r'^\s*([a-zA-Z0-9_.]+)\s*=\s*{', vanilla_lines[i])
        if m:
            event_key = m.group(1)
            event_block = []
            depth = vanilla_lines[i].count('{') - vanilla_lines[i].count('}')
            event_block.append(vanilla_lines[i])
            i += 1
            while i < len(vanilla_lines) and depth > 0:
                event_block.append(vanilla_lines[i])
                depth += vanilla_lines[i].count('{')
                depth -= vanilla_lines[i].count('}')
                i += 1
            if event_key in modded_events:
                merged_lines.extend(modded_events[event_key])
            else:
                merged_lines.extend(event_block)
        else:
            merged_lines.append(vanilla_lines[i])
            i += 1

    # Add custom comments from mod file that aren't already present
    for comment in custom_comments:
        if comment not in merged_lines:
            merged_lines.append(comment)

    os.makedirs(os.path.dirname(output_path), exist_ok=True)
    with open(output_path, 'w') as f:
        f.writelines(merged_lines)

def main():
    import sys
    if len(sys.argv) != 2:
        print("Usage: python3 vanilla_merger.py <config.yaml>")
        return
    config = load_config(sys.argv[1])
    game_root = config["game_root"]
    mod_root = mod_root = os.path.abspath(os.path.join(os.path.dirname(__file__), '..'))
    files = config["files"]

    for file_entry in files:
        rel_path = file_entry["relative_path"]
        custom_events = file_entry["modified_entity_keys"]
        vanilla_path = os.path.join(game_root, rel_path)
        mod_path = os.path.join(mod_root, rel_path)
        output_path = os.path.join(mod_root, "vanilla-mod-merger/merge-output", os.path.basename(rel_path))
        print(f"Merging {rel_path}...")
        merge_file(vanilla_path, mod_path, output_path, custom_events)

if __name__ == "__main__":
    main()