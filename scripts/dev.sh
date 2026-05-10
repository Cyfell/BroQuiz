#!/usr/bin/env bash
# Run pnpm dev with output teed to a timestamped log file under logs/.
# Keeps the most recent ${BROQUIZ_LOG_RETAIN:-10} log files; older ones are deleted.
# A logs/dev-latest.log symlink always points at the current run's file.

set -euo pipefail

log_dir="logs"
retain="${BROQUIZ_LOG_RETAIN:-10}"
mkdir -p "$log_dir"

stamp="$(date +%Y%m%d-%H%M%S)"
logfile="$log_dir/dev-$stamp.log"

# Rotate: keep the ${retain} most recent dev-*.log files.
# `|| true` swallows a non-zero exit from `ls` when no log files exist yet.
{ ls -1t "$log_dir"/dev-*.log 2>/dev/null || true; } | tail -n +$((retain + 1)) | xargs -r rm -f

ln -sfn "dev-$stamp.log" "$log_dir/dev-latest.log"

echo "[dev] logging to $logfile"
echo "[dev]   tail -f $log_dir/dev-latest.log"
echo

# Strip ANSI colour codes from the file copy; keep them in the terminal.
strip_ansi='s/\x1B\[[0-9;]*[A-Za-z]//g'

concurrently -k -n vite,api -c cyan,magenta "vite" "tsx watch server/index.ts" 2>&1 \
  | tee >(sed -u "$strip_ansi" > "$logfile")
