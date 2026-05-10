# BroQuiz

A LAN buzzer game with a slidev-style dual UI: an **audience** view (big screen, scores + who buzzed first) and a **presenter** view (controls, team management, validate / discard). Hardware buzzers send HTTP POSTs; the UIs subscribe to live state via WebSocket.

## Repository layout

```
api/
  openapi.yaml          OpenAPI 3 spec for the HTTP + WebSocket API
buzzer/                 Hardware buzzer firmware (separate domain — not part of the Node package)
client/                 Vue 3 SPA: audience view + presenter view
server/                 Fastify server: routes, WebSocket, state, snapshots
shared/                 Types and constants shared by client and server
data/                   Runtime snapshot file (created on first state change, gitignored)
package.json            The Node package lives at the repo root
```

## Prerequisites

- Node.js 22+
- pnpm 9+

## Install

```bash
pnpm install
```

## Run (development)

```bash
pnpm dev
```

Two processes start:

- **Vite** on `http://localhost:5173` — serves the SPA with HMR. API + WebSocket calls are proxied to Fastify.
- **Fastify** on `http://localhost:8080` — HTTP API and WebSocket.

Open:

- Audience: <http://localhost:5173/>
- Presenter: <http://localhost:5173/presenter>

Hardware buzzers should target the Fastify port directly (`:8080`).

### Dev logs

`pnpm dev` mirrors all output (Vite + Fastify) to a timestamped file under `logs/`:

```
logs/dev-20260510-164230.log
logs/dev-latest.log         → symlink to the current run's file
```

Tail the active run from another terminal:

```bash
tail -f logs/dev-latest.log
```

The script keeps the **10 most recent** runs and deletes older ones. Override with `BROQUIZ_LOG_RETAIN`:

```bash
BROQUIZ_LOG_RETAIN=50 pnpm dev
```

The terminal still shows colored output; the file copy has ANSI codes stripped. The `logs/` directory is gitignored.

## Run (production)

```bash
pnpm build       # vue-tsc + vite build → dist/
pnpm start       # Fastify on :8080 serves the built SPA + API + WS
```

In production a single port (`8080`) serves everything.

## Configuration

Environment variables (all optional):

| Variable           | Default          | Purpose                              |
| ------------------ | ---------------- | ------------------------------------ |
| `HOST`             | `0.0.0.0`        | Bind address                         |
| `PORT`             | `8080`           | HTTP / WebSocket port                |
| `LOG_LEVEL`        | `info`           | Fastify log level                    |
| `BROQUIZ_API_URL`  | `http://localhost:8080` | Used by Vite dev proxy        |
| `BROQUIZ_LOG_RETAIN` | `10`           | How many `pnpm dev` log files to keep |

Cooldown duration (after a wrong answer) is set at runtime via the presenter UI or `POST /round/cooldown`. Default is 2000 ms.

## Game flow

1. **Presenter** creates teams. Each team's `id` must match the hardware buzzer that belongs to it (buzzer id == team id).
2. A buzzer is pressed → `POST /buzzer {id}` → if accepted, that team becomes the "current buzz" and further presses are rejected until the round is resolved.
3. The presenter resolves the buzz:
   - **Validate** — the team's score is incremented by 1, the snapshot is written, the round is cleared, and any other-team cooldowns are lifted.
   - **Discard** — no score change; the team that buzzed is put on cooldown (default 2 s); other teams can buzz immediately for the same question.
4. Repeat. Use **Lock** to globally pause buzzing during reveals; **Reset all** wipes teams, scores, and round state.

## Hardware buzzer contract

```
POST http://<server>:8080/buzzer
Content-Type: application/json

{ "id": <buzzer_id> }
```

Responses:

- `200 { "accepted": true }` — first press of the round.
- `409 { "accepted": false, "reason": "locked" | "already_buzzed" | "cooldown" | "unknown_team" }` — rejected. The hardware can use this to drive a feedback LED (e.g. blink red on cooldown).

## Persistence

Teams, scores, lock state, and cooldown configuration are written atomically to `data/snapshot.json` on every state mutation and restored on startup. Transient state (current buzz, active cooldowns) is **not** persisted — a crash mid-round resets to "no active buzz".

## API reference

Full HTTP + WebSocket reference: [`api/openapi.yaml`](api/openapi.yaml).

WebSocket envelope on `/events`:

```json
{ "type": "state", "state": { /* full GameState */ } }
{ "type": "buzz",  "buzz":  { "teamId": 1, "at": 1715000000000 } }
```

A `state` message is sent immediately on connect, then again on every change. A `buzz` message is emitted each time a press is accepted (the following `state` message reflects the new `currentBuzz`).

## Scripts

| Command           | What it does                                    |
| ----------------- | ----------------------------------------------- |
| `pnpm dev`        | Vite + Fastify in watch mode (logs to `logs/`)  |
| `pnpm build`      | Type-check and build the SPA to `dist/`         |
| `pnpm start`      | Production server (serves `dist/` + API + WS)   |
| `pnpm typecheck`  | `vue-tsc --noEmit`                              |
