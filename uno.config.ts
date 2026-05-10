import { defineConfig, presetUno, presetIcons } from 'unocss';

export default defineConfig({
  presets: [presetUno(), presetIcons()],
  theme: {
    colors: {
      brand: {
        DEFAULT: '#6366f1',
        dark: '#4f46e5',
      },
    },
  },
});
