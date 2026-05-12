<script setup lang="ts">
    import { inject, ref } from 'vue'
    import type { Waifu } from '../tools/types.ts'
    import WaifuComp from './WaifuComp.vue'
    import { NavigationHandler } from '@/tools/navigationHandler.ts'

    const navigationHandler = inject<NavigationHandler>('navigationHandler')!
    const userWaifus = navigationHandler.getUserWaifus()
    const selectedWaifu = ref<Waifu | undefined>(undefined)
    const isCreating = ref(false)

    function clickOnWaifu(waifu: Waifu) {
        selectedWaifu.value = selectedWaifu.value == waifu ? undefined : waifu
    }

    async function onclick() {
        if (!selectedWaifu.value || isCreating.value) {
            return
        }

        isCreating.value = true

        try {
            await navigationHandler.createAHeroFromAWaifu(selectedWaifu.value)
            selectedWaifu.value = undefined
        } finally {
            isCreating.value = false
        }
    }
</script>

<template>
    <section v-if="userWaifus" class="selection-panel">
        <div class="panel-heading">
            <div>
                <p class="eyebrow">Summoning room</p>
                <h2>Select a waifu</h2>
                <p>Choose the next recruit you want to elevate into your hero roster.</p>
            </div>

            <div class="selection-toolbar">
                <div class="stat-pill">
                    <span>Chosen</span>
                    <strong>{{ selectedWaifu?.name_user_preferred ?? 'None' }}</strong>
                </div>
                <button class="primary-button" v-on:click="onclick()" :disabled="!selectedWaifu || isCreating">{{ isCreating ? 'Forging hero...' : 'Make a hero' }}</button>
            </div>
        </div>

        <div class="selection-grid waifu-grid">
            <WaifuComp
                v-for="waifu in userWaifus"
                :key="waifu.id"
                :waifu="waifu"
                :class="{ 'is-selected': selectedWaifu == waifu }"
                v-on:click="clickOnWaifu(waifu)"
            />
        </div>
    </section>

    <section v-else class="selection-panel">
        <p class="eyebrow">Summoning room</p>
        <h2>Loading waifus...</h2>
    </section>
</template>
