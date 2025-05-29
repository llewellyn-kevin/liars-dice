import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import tailwindcss from '@tailwindcss/vite'

// https://vite.dev/config/
export default defineConfig({
    plugins: [
        react(),
        tailwindcss(),
    ],
    build: {
        // generates .vite/manifest.json in outDir
        manifest: true,

        rollupOptions: {
            // overwrite default .html entry and include a secondary
            input: {
                main: "/src/main.tsx",
            },
        },
    },
})
