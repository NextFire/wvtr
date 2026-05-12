<script setup lang="ts">
    import { inject } from 'vue'
    import type { NavigationHandler } from '@/tools/navigationHandler.ts'

    const navigationHandler = inject<NavigationHandler>('navigationHandler')!
    const user = navigationHandler.getUser()
    const gameFaviconSrc = '/imgs/favicon.ico'
</script>

<template>
    <header v-if="user" class="header quest-header panel panel-raised">
        <div class="header-brand">
            <div class="header-crest">
                <img class="header-crest-icon" :src="gameFaviconSrc" alt="Waiventure favicon"/>
            </div>
            <div>
                <p class="eyebrow">Playful quest command center</p>
                <h1>Waiventure</h1>
                <p class="header-tagline">Captain {{ user.name }} is ready for another bright little disaster.</p>
            </div>
        </div>

        <div class="header-stats">
            <div class="stat-pill">
                <span>Heroes</span>
                <strong>{{ user.ownedHeroes.length }}</strong>
            </div>
            <div class="stat-pill">
                <span>Party Slots</span>
                <strong>{{ user.currentTeam?.heroes.length ?? 0 }}/3</strong>
            </div>
        </div>

        <div class="currency-strip">
            <div
                v-if="user.inventory && user.inventory.currencies"
                v-for="(c, index) in user.inventory.currencies"
                :key="index"
                class="currency-chip"
            >
                <span class="currency-icon">
                    <img :src="c.currency.iconURL" :alt="`Currency ${index + 1}`" width="18"/>
                </span>
                <strong>{{ c.numberOwned }}</strong>
            </div>
        </div>
    </header>
    <section v-else class="header panel">
        <p class="eyebrow">Preparing guild hall</p>
        <h1>Loading player profile...</h1>
    </section>
</template>
