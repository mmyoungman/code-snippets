#!/bin/bash
# Black-box smoke test for the procgen Linux build.
#
# Verifies, against the compiled binary only:
#   1. The build produces a runnable executable.
#   2. The program launches and renders a window.
#   3. Pressing 'c' regenerates a new layout (window content changes).
#   4. Pressing 'q' quits the program cleanly (exit 0).
#
# Requires: xdotool, ImageMagick (import, compare), and an X display.

set -u
cd "$(dirname "$0")"

BUILD_SCRIPT="./build_linux.sh"
BINARY="./procgen"
TITLE="Proc Gen"
WORKDIR="$(mktemp -d -t procgen-test.XXXXXX)"
PID=""
CLEANED=0

cleanup() {
    if [[ "${CLEANED}" -eq 0 ]]; then
        CLEANED=1
        if [[ -n "${PID}" ]] && kill -0 "${PID}" 2>/dev/null; then
            kill "${PID}" 2>/dev/null
        fi
        rm -rf "${WORKDIR}"
    fi
    exit $1
}
trap 'cleanup 1' INT TERM EXIT

fail() { echo "FAIL: $1"; cleanup 1; }
need() { command -v "$1" >/dev/null 2>&1 || fail "missing tool: $1"; }
current_pid_is() {
    # returns 0 if the launched procgen PID is still alive
    [[ -n "${PID}" ]] && kill -0 "${PID}" 2>/dev/null
}

need xdotool
need import
need compare

# ---------------------------------------------------------------- 1. build
echo "[1/4] Building..."
rm -f procgen
if ! bash "${BUILD_SCRIPT}" 2>"${WORKDIR}/build_err.log"; then
    echo "=== build output (stderr) ==="
    cat "${WORKDIR}/build_err.log"
    fail "build script failed"
fi
[[ -x "${BINARY}" ]] || fail "build did not produce an executable '${BINARY}'"
echo "OK: build succeeded"

# --------------------------------------------------------------- 2. launch
echo "[2/4] Launching ${BINARY}..."
"${BINARY}" &
PID=$!
WAIT_SLEEP=0.1

WID=""
for _ in $(seq 1 50); do
    WID=$(xdotool search --sync --name "${TITLE}" 2>/dev/null | head -n1)
    [[ -n "${WID}" ]] && break
    sleep "${WAIT_SLEEP}"
done
[[ -n "${WID}" ]] || fail "window '${TITLE}' did not appear"
echo "OK: window found ($WID)"
current_pid_is || fail "process exited before the test could interact"

# let the first frame render
sleep 1

# ------------------------------------------------------ 3. 'c' regenerates
echo "[3/4] Testing 'c' key regeneration..."
shot() { import -window "${WID}" "$1" 2>/dev/null; }
xfocus() { xdotool windowactivate --sync "${WID}" 2>/dev/null || xdotool windowfocus "${WID}" 2>/dev/null; }

xfocus
sleep 0.2
shot "${WORKDIR}/before.png"
[[ -s "${WORKDIR}/before.png" ]] || fail "could not capture the window"

xdotool key --clearmodifiers c
sleep 0.5
shot "${WORKDIR}/after.png"
current_pid_is || fail "process exited after pressing 'c'"

AE_RAW=$(compare -metric AE "${WORKDIR}/before.png" "${WORKDIR}/after.png" null: 2>&1)
AE=$(echo "${AE_RAW}" | grep -oE '^[0-9]+' || echo "0")
if [[ "${AE}" == "0" ]]; then
    fail "'c' did not change the screen content (identical pixels: ${AE_RAW})"
fi
echo "OK: 'c' changed ${AE} pixels (areas regenerated, no crash)"

# ----------------------------------------------------------------- 4. quit
echo "[4/4] Testing 'q' quit..."
xfocus
xdotool key --clearmodifiers q

EXITED=0
for _ in $(seq 1 30); do
    if ! current_pid_is; then EXITED=1; break; fi
    sleep "${WAIT_SLEEP}"
done
[[ "${EXITED}" -eq 1 ]] || fail "process still running 3s after 'q'"

wait "${PID}" 
RC=$?
echo "OK: process exited (code ${RC})"
[[ "${RC}" -eq 0 ]] || fail "process exited with non-zero code ${RC}"

echo
echo "ALL TESTS PASSED"
cleanup 0