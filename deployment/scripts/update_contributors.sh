#!/bin/bash

# Define markers to identify where to insert contributor stats
START_MARKER="<!-- START_CONTRIBUTOR_STATS -->"
END_MARKER="<!-- END_CONTRIBUTOR_STATS -->"

# Get contribution stats (commits, lines of code, files changed)
CONTRIB_STATS=$(git fame --format markdown)

# Get commit count sorted from highest to lowest
COMMIT_COUNT=$(git shortlog -sne | sort -nr | awk '{print "| " $2 " " $3 " | " $1 " |"}')

# Get complexity solved (Lines of Code per user)
COMPLEXITY_STATS=$(cloc . --by-author --csv | tail -n +2 | column -t -s,)

# Create new stats section
NEW_STATS=$(cat <<EOF
$START_MARKER

## 📊 Contributor Statistics

### 🔥 Commits, Lines of Code & Files Changed
$CONTRIB_STATS

### 📈 Commits by Contributor (Sorted by Count)
| Contributor | Commits |
|------------|---------|
$COMMIT_COUNT

### ⚡ Complexity Solved (Lines of Code by User)
\`\`\`
$COMPLEXITY_STATS
\`\`\`

$END_MARKER
EOF
)

# Remove old stats from README and insert new stats
sed -i "/$START_MARKER/,/$END_MARKER/d" README.md
echo "$NEW_STATS" >> README.md

# Commit and push changes
git config --global user.email "github-actions@github.com"
git config --global user.name "GitHub Actions"
git add README.md
git commit -m "Updated contributor statistics" || exit 0
git push
