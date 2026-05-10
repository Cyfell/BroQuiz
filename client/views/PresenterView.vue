<script setup lang="ts">
import { onMounted, onUnmounted, ref, computed } from 'vue';
import { useGameStore } from '../composables/useGameStore';
import { api } from '../composables/useApi';

const game = useGameStore();
const now = ref(Date.now());
let nowTimer: number | null = null;

onMounted(() => {
  game.connect();
  nowTimer = window.setInterval(() => {
    now.value = Date.now();
  }, 100);
});
onUnmounted(() => {
  game.disconnect();
  if (nowTimer !== null) window.clearInterval(nowTimer);
});

const newId = ref<number | null>(null);
const newName = ref('');
const cooldownInput = ref<number | null>(null);
const confirmReset = ref(false);

async function doReset() {
  await api.reset();
  confirmReset.value = false;
}

const buzzedTeam = computed(() => {
  const id = game.currentBuzz?.teamId;
  return id !== undefined ? game.teams.find((t) => t.id === id) ?? null : null;
});

async function addTeam() {
  if (newId.value === null || !newName.value.trim()) return;
  await api.createTeam(newId.value, newName.value.trim());
  newId.value = null;
  newName.value = '';
}

async function applyCooldown() {
  if (cooldownInput.value === null || cooldownInput.value < 0) return;
  await api.setCooldown(cooldownInput.value);
}

function cooldownRemaining(teamId: number): number {
  const until = game.cooldowns[teamId];
  if (!until) return 0;
  return Math.max(0, until - now.value);
}
</script>

<template>
  <div class="min-h-screen bg-gray-950 text-white p-6 grid grid-cols-[1fr_360px] gap-6">
    <main class="flex flex-col gap-6">
      <header class="flex items-center justify-between">
        <h1 class="text-2xl font-bold">Presenter</h1>
        <div class="flex items-center gap-3 text-sm">
          <button
            class="breathe w-3 h-3 rounded-full"
            type="button"
            :class="game.connected ? 'bg-green-500' : 'bg-red-500'"
            :style="{ '--breathe-color': game.connected ? '34 197 94' : '239 68 68' }"
            :aria-label="game.connected ? 'connected' : 'disconnected'"
          />
          <span :class="game.connected ? 'text-green-400' : 'text-red-400'">
            {{ game.connected ? 'live' : 'offline' }}
          </span>
          <span :class="game.locked ? 'text-amber-400' : 'text-gray-500'">
            {{ game.locked ? 'LOCKED' : 'open' }}
          </span>
        </div>
      </header>

      <section
        class="rounded-2xl p-8 text-center"
        :style="buzzedTeam ? { background: buzzedTeam.color } : { background: '#111827' }"
      >
        <div v-if="buzzedTeam">
          <div class="text-sm uppercase tracking-widest opacity-70">Active buzz</div>
          <div class="text-5xl font-extrabold mt-2">{{ buzzedTeam.name }}</div>
          <div class="flex gap-3 justify-center mt-6">
            <button class="bg-green-600 hover:bg-green-500 px-6 py-3 rounded-lg font-bold" @click="api.validate()">
              ✓ Validate (+1)
            </button>
            <button class="bg-red-600 hover:bg-red-500 px-6 py-3 rounded-lg font-bold" @click="api.discard()">
              ✗ Discard (cooldown)
            </button>
          </div>
        </div>
        <div v-else class="text-gray-500 text-xl">No active buzz</div>
      </section>

      <section class="bg-gray-900 rounded-xl p-4">
        <div class="flex items-center justify-between mb-3">
          <h2 class="font-semibold">Teams</h2>
          <div class="flex gap-2">
            <button class="text-xs bg-gray-700 hover:bg-gray-600 px-2 py-1 rounded" @click="game.locked ? api.unlock() : api.lock()">
              {{ game.locked ? 'Unlock' : 'Lock' }}
            </button>
            <button class="text-xs bg-red-700 hover:bg-red-600 px-2 py-1 rounded" @click="confirmReset = true">
              Reset all
            </button>
          </div>
        </div>

        <table class="w-full text-sm">
          <thead class="text-gray-400 text-left">
            <tr>
              <th class="py-1">ID</th>
              <th class="py-1">Name</th>
              <th class="py-1">Score</th>
              <th class="py-1">Cooldown</th>
              <th class="py-1"></th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="t in game.teams" :key="t.id" class="border-t border-gray-800">
              <td class="py-2">
                <span class="inline-block w-3 h-3 rounded-full mr-2" :style="{ background: t.color }" />
                {{ t.id }}
              </td>
              <td class="py-2">{{ t.name }}</td>
              <td class="py-2 font-bold">{{ game.scores[t.id] ?? 0 }}</td>
              <td class="py-2 text-amber-400">
                {{ cooldownRemaining(t.id) > 0 ? (cooldownRemaining(t.id) / 1000).toFixed(1) + 's' : '—' }}
              </td>
              <td class="py-2 flex gap-1 justify-end">
                <button class="bg-gray-700 hover:bg-gray-600 px-2 rounded" @click="api.score(t.id, +1)">+1</button>
                <button class="bg-gray-700 hover:bg-gray-600 px-2 rounded" @click="api.score(t.id, -1)">-1</button>
                <button class="bg-blue-700 hover:bg-blue-600 px-2 rounded" @click="api.testBuzz(t.id)">test buzz</button>
                <button class="bg-red-800 hover:bg-red-700 px-2 rounded" @click="api.deleteTeam(t.id)">×</button>
              </td>
            </tr>
            <tr v-if="game.teams.length === 0">
              <td colspan="5" class="py-4 text-center text-gray-500">No teams yet — add one below.</td>
            </tr>
          </tbody>
        </table>
      </section>
    </main>

    <aside class="flex flex-col gap-4">
      <section class="bg-gray-900 rounded-xl p-4">
        <h2 class="font-semibold mb-3">Add team</h2>
        <div class="flex flex-col gap-2">
          <label class="text-xs text-gray-400">Buzzer ID</label>
          <input
            v-model.number="newId"
            type="number"
            class="bg-gray-800 text-white px-3 py-2 rounded outline-none focus:ring-1 focus:ring-brand"
            placeholder="e.g. 1"
          />
          <label class="text-xs text-gray-400 mt-2">Team name</label>
          <input
            v-model="newName"
            class="bg-gray-800 text-white px-3 py-2 rounded outline-none focus:ring-1 focus:ring-brand"
            placeholder="e.g. Red Team"
            @keyup.enter="addTeam"
          />
          <button class="mt-3 bg-brand hover:bg-brand-dark px-3 py-2 rounded font-semibold" @click="addTeam">
            Add team
          </button>
        </div>
      </section>

      <section class="bg-gray-900 rounded-xl p-4">
        <h2 class="font-semibold mb-3">Cooldown</h2>
        <div class="flex flex-col gap-2">
          <label class="text-xs text-gray-400">Wrong-answer lockout (ms)</label>
          <div class="flex gap-2">
            <input
              v-model.number="cooldownInput"
              type="number"
              :placeholder="String(game.cooldownMs)"
              class="bg-gray-800 text-white px-3 py-2 rounded flex-1 outline-none focus:ring-1 focus:ring-brand"
            />
            <button class="bg-brand hover:bg-brand-dark px-3 rounded font-semibold" @click="applyCooldown">
              set
            </button>
          </div>
          <div class="text-xs text-gray-500">current: {{ game.cooldownMs }} ms</div>
        </div>
      </section>
    </aside>

    <Teleport to="body">
      <div
        v-if="confirmReset"
        class="fixed inset-0 bg-black/70 flex items-center justify-center z-50"
        @click.self="confirmReset = false"
        @keydown.esc="confirmReset = false"
      >
        <div class="bg-gray-900 border border-gray-700 rounded-xl p-6 w-[420px] max-w-[90vw] shadow-2xl">
          <h3 class="text-xl font-bold mb-2">Reset everything?</h3>
          <p class="text-gray-400 text-sm mb-5">
            This wipes all teams, scores, cooldowns, and the current round. The snapshot file will be overwritten. This action cannot be undone.
          </p>
          <div class="flex gap-2 justify-end">
            <button
              class="bg-gray-700 hover:bg-gray-600 px-4 py-2 rounded font-semibold"
              @click="confirmReset = false"
            >
              Cancel
            </button>
            <button
              class="bg-red-700 hover:bg-red-600 px-4 py-2 rounded font-semibold"
              @click="doReset"
            >
              Reset all
            </button>
          </div>
        </div>
      </div>
    </Teleport>
  </div>
</template>
