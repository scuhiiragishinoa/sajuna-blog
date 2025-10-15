import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { fileURLToPath, URL } from 'node:url'

export default defineConfig({
    plugins: [vue()],

    server: {
        host: '0.0.0.0',
        port: 3000,
        proxy: {
            '/api': {
                target: 'http://backend:8080',
                changeOrigin: true,
                secure: false,
            },
        },
    },

    resolve: {
        alias: {
            '@': fileURLToPath(new URL('./src', import.meta.url)),
            '@components': fileURLToPath(new URL('./src/components', import.meta.url)),
            '@views': fileURLToPath(new URL('./src/views', import.meta.url)),
            '@stores': fileURLToPath(new URL('./src/stores', import.meta.url)),
            '@utils': fileURLToPath(new URL('./src/utils', import.meta.url)),
        },
    },

    css: {
        preprocessorOptions: {
            scss: {
                additionalData: `@import "@/styles/variables.scss";`,
            },
        },
    },

    define: {
        'process.env': {},
    },
})