<script setup lang="ts">
import { onMounted, onUnmounted, computed } from 'vue';
import { useGameStore } from '../composables/useGameStore';

const game = useGameStore();
onMounted(() => game.connect());
onUnmounted(() => game.disconnect());

const buzzedTeam = computed(() => {
  const id = game.currentBuzz?.teamId;
  return id !== undefined ? game.teams.find((t) => t.id === id) ?? null : null;
});

const cols = computed(() => {
  const n = game.teams.length || 1;
  if (n <= 1) return 1;
  if (n <= 4) return 2;
  if (n <= 9) return 3;
  if (n <= 16) return 4;
  return Math.ceil(Math.sqrt(n));
});
</script>

<template>
  <div
    class="h-screen w-screen overflow-hidden bg-gray-950 text-white p-4 grid gap-3"
    style="grid-template-rows: minmax(0, 22vh) minmax(0, 1fr); box-sizing: border-box;"
  >
    <section
      v-if="buzzedTeam"
      class="rounded-2xl text-center flex flex-col items-center justify-center min-h-0 overflow-hidden"
      :style="{ background: buzzedTeam.color }"
    >
      <div class="uppercase tracking-widest opacity-70" style="font-size: clamp(0.75rem, 1.6vmin, 1.25rem)">Buzzed</div>
      <div
        class="font-extrabold mt-1 truncate w-full px-4"
        style="font-size: clamp(2rem, 10vmin, 6rem); line-height: 1.05"
      >
        {{ buzzedTeam.name }}
      </div>
    </section>
    <section
      v-else
      class="rounded-2xl bg-gray-900 border border-gray-800 flex items-center justify-center min-h-0 overflow-hidden"
    >
      <div class="text-gray-500" style="font-size: clamp(1rem, 3vmin, 2rem)">Waiting for a buzz…</div>
    </section>

    <section
      class="grid gap-3 min-h-0 overflow-hidden"
      :style="{
        gridTemplateColumns: `repeat(${cols}, minmax(0, 1fr))`,
        gridAutoRows: 'minmax(0, 1fr)',
      }"
    >
      <div
        v-for="team in game.teams"
        :key="team.id"
        class="rounded-xl border-4 transition flex flex-col items-center justify-center min-h-0 min-w-0 overflow-hidden p-2"
        :class="game.currentBuzz?.teamId === team.id ? 'scale-[1.02]' : 'border-transparent'"
        :style="{
          background: team.color + '22',
          borderColor: game.currentBuzz?.teamId === team.id ? team.color : 'transparent',
        }"
      >
        <div
          class="font-semibold truncate w-full text-center"
          :style="{ color: team.color, fontSize: 'clamp(0.875rem, 2.8vmin, 1.75rem)' }"
        >
          {{ team.name }}
        </div>
        <div
          class="font-black"
          :style="{ fontSize: 'clamp(1.5rem, 7vmin, 5rem)', lineHeight: 1 }"
        >
          {{ game.scores[team.id] ?? 0 }}
        </div>
      </div>
    </section>
  </div>
</template>
