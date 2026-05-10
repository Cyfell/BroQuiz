import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import type { GameState, ServerMessage, Buzz } from '../../shared/types';

const wsUrl = () => {
  const proto = location.protocol === 'https:' ? 'wss' : 'ws';
  return `${proto}://${location.host}/events`;
};

export const useGameStore = defineStore('game', () => {
  const state = ref<GameState | null>(null);
  const lastBuzz = ref<Buzz | null>(null);
  const connected = ref(false);
  let ws: WebSocket | null = null;
  let retry: number | null = null;

  function connect() {
    ws = new WebSocket(wsUrl());
    ws.onopen = () => {
      connected.value = true;
    };
    ws.onmessage = (ev) => {
      const msg = JSON.parse(ev.data) as ServerMessage;
      if (msg.type === 'state') state.value = msg.state;
      else if (msg.type === 'buzz') lastBuzz.value = msg.buzz;
    };
    ws.onclose = () => {
      connected.value = false;
      retry = window.setTimeout(connect, 1000);
    };
    ws.onerror = () => {
      ws?.close();
    };
  }

  function disconnect() {
    if (retry) window.clearTimeout(retry);
    ws?.close();
    ws = null;
  }

  const teams = computed(() => state.value?.teams ?? []);
  const scores = computed(() => state.value?.scores ?? {});
  const currentBuzz = computed(() => state.value?.currentBuzz ?? null);
  const cooldowns = computed(() => state.value?.cooldowns ?? {});
  const locked = computed(() => state.value?.locked ?? false);
  const cooldownMs = computed(() => state.value?.config.cooldownMs ?? 2000);

  return {
    state,
    lastBuzz,
    connected,
    teams,
    scores,
    currentBuzz,
    cooldowns,
    locked,
    cooldownMs,
    connect,
    disconnect,
  };
});
