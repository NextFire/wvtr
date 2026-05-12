<script setup lang="ts">
    import { inject, ref, watch } from 'vue'
    import type { ExpeditionStepResolveInfo, ExpToGetFromBack } from '../tools/types.ts'
    import { formatTextTimeFromTimeMS } from '../tools/utils.ts'
    import type { NavigationHandler } from '@/tools/navigationHandler.ts'

    const navigationHandler = inject<NavigationHandler>('navigationHandler')!

    const expeditions = navigationHandler.getAvailableExpedition()
    const user = navigationHandler.getUser()
    const selectedExp = ref('')
    const errorMsg = ref('')
    const expStepInfo = ref<ExpeditionStepResolveInfo | undefined>(undefined)

    function isSelected(expedition: ExpToGetFromBack) {
        return selectedExp.value === expedition.key
    }

    function clickOnExpedition(expedition: ExpToGetFromBack) {
        if (!expedition.canBeLaunched) {
            return
        }

        selectedExp.value = selectedExp.value === expedition.key ? '' : expedition.key
        errorMsg.value = ''
    }

    async function onclick() {
        if (selectedExp.value === '') {
            errorMsg.value = 'Pick an expedition card before launching.'
            return
        }

        if (user.value!.currentTeam.heroes.length > 0) {
            errorMsg.value = ''
            await navigationHandler.launchExpedition(expStepInfo, selectedExp.value)
            return
        }

        errorMsg.value = 'You should have at least one hero in your team before starting an expedition.'
    }

    watch(expStepInfo, (newExpInfo) => {
        if (newExpInfo) {
            user.value!.state.state = newExpInfo.stepState
        }
    })
</script>

<template>
    <section v-if="expeditions" class="selection-panel">
        <div class="panel-heading">
            <div>
                <p class="eyebrow">Quest counter</p>
                <h2>Choose an expedition</h2>
                <p>Pick a mission card to compare duration, cost, and launch readiness.</p>
            </div>

            <div class="selection-toolbar">
                <div class="stat-pill">
                    <span>Selected</span>
                    <strong>{{ selectedExp || 'None' }}</strong>
                </div>
                <button class="primary-button" v-on:click="onclick()" :disabled="selectedExp === ''">Launch expedition</button>
            </div>
        </div>

        <p v-if="errorMsg" class="inline-error">{{ errorMsg }}</p>

        <div class="selection-grid expedition-grid">
            <article
                v-for="expedition in expeditions"
                :key="expedition.key"
                class="selectable-card expedition-card"
                :class="{ 'is-selected': isSelected(expedition), 'is-disabled': !expedition.canBeLaunched }"
                v-on:click="clickOnExpedition(expedition)"
            >
                <div class="expedition-art">
                    <img :src="expedition.imgURL" :alt="expedition.key"/>
                </div>

                <div class="selectable-copy">
                    <div class="card-heading-row">
                        <h3>{{ expedition.key }}</h3>
                        <span class="badge-pill">{{ formatTextTimeFromTimeMS(expedition.duration / 1000000) }}</span>
                    </div>

                    <p>
                        {{ expedition.canBeLaunched ? 'Ready to launch with your current resources.' : 'Locked until the required conditions are met.' }}
                    </p>

                    <div class="card-meta-row">
                        <span class="badge-pill badge-pill-soft">
                            {{ expedition.costName != '' ? `Cost ${expedition.costName} x${expedition.costNumber}` : 'No launch cost' }}
                        </span>
                        <span v-if="!expedition.canBeLaunched" class="badge-pill badge-pill-warning">Locked</span>
                    </div>
                </div>
            </article>
        </div>
    </section>

    <section v-else class="selection-panel">
        <p class="eyebrow">Quest counter</p>
        <h2>Loading expeditions...</h2>
    </section>
</template>
