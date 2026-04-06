#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
COUNT="${1:-5}"

if ! [[ "$COUNT" =~ ^[0-9]+$ ]]; then
  echo "Usage: scripts/wiki-recent.sh [count]" >&2
  exit 1
fi

rg '^## \[' "$ROOT/wiki/log.md" | tail -n "$COUNT"
