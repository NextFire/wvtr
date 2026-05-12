<script setup lang="ts">
    import { inject, onMounted } from 'vue'
    import Header from './components/Header.vue'
    import Body from './components/Body.vue'
    import { NavigationHandler, NavigationStatus } from './tools/navigationHandler.ts'
    import type { VueCookies } from 'vue-cookies'

    const $cookies = inject<VueCookies>('$cookies')
    const navigationHandler = inject<NavigationHandler>('navigationHandler')!
    const navigationStatus = navigationHandler.getNavigationStatus()

    onMounted(async () => {
        await navigationHandler.setup($cookies!)
    })
</script>

<template>
    <div v-if="navigationStatus == NavigationStatus.Connexion" class="page page-shell loading-screen">
        <section class="panel panel-feature loading-card">
            <p class="eyebrow">Synchronizing save crystal</p>
            <h1>Waiventure</h1>
            <p>Loading your guild, currencies, and latest expedition state.</p>
        </section>
    </div>
    <div v-else class="page page-shell">
        <Header/>
        <Body/>
    </div>
</template>
