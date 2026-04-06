#!/usr/bin/env bash
set -euo pipefail

if [[ $# -lt 1 ]]; then
  echo "Usage: scripts/wiki-search.sh \"query\"" >&2
  exit 1
fi

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
QUERY="$*"

echo "# Index matches"
rg -n -i -- "$QUERY" "$ROOT/wiki/index.md" || true
echo
echo "# Page matches"
rg -n -i \
  --glob "*.md" \
  --glob "!index.md" \
  --glob "!log.md" \
  --glob "!README.md" \
  -- "$QUERY" "$ROOT/wiki" || true
