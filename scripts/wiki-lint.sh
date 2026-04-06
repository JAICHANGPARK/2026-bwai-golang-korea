#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
WIKI="$ROOT/wiki"
INDEX="$WIKI/index.md"

status=0

while IFS= read -r file; do
  rel="${file#$WIKI/}"

  if ! rg -Fq -- "$rel" "$INDEX"; then
    echo "MISSING_INDEX_ENTRY $rel"
    status=1
  fi

  if ! rg -q '^---$' "$file"; then
    echo "MISSING_FRONTMATTER $rel"
    status=1
  fi

  if ! rg -q '^updated:' "$file"; then
    echo "MISSING_UPDATED_FIELD $rel"
    status=1
  fi

  if ! rg -q '^type:' "$file"; then
    echo "MISSING_TYPE_FIELD $rel"
    status=1
  fi
done < <(find "$WIKI" -type f -name '*.md' ! -name 'index.md' ! -name 'log.md' | sort)

if [[ "$status" -eq 0 ]]; then
  echo "OK wiki lint passed"
fi
