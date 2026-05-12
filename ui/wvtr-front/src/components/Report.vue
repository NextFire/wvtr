<script setup lang="ts">
    import type { NavigationHandler } from '@/tools/navigationHandler'
    import { EncounterState } from '@/tools/types'
    import { getEncounterStateString, getStringFromFAD } from '@/tools/utils'
    import { inject } from 'vue'

    const navigationHandler = inject<NavigationHandler>('navigationHandler')!
    const report = navigationHandler.getReport()

    async function onclick() {
        await navigationHandler.setGameState(EncounterState.Home)
        await navigationHandler.fetchTeam()
        await navigationHandler.fetchInventory()
    }
</script>

<template>
    <section v-if="report" class="report-screen">
        <section class="panel panel-feature report-summary">
            <div>
                <p class="eyebrow">Mission log</p>
                <h2>{{ report.identifier }}</h2>
                <p>Review the flow of the expedition, from travel beats to combat events, before heading back to the guild hall.</p>
            </div>

            <div class="report-summary-actions">
                <div class="stat-pill">
                    <span>Started</span>
                    <strong>{{ new Date(report.startedAt).toLocaleString() }}</strong>
                </div>
                <div class="stat-pill">
                    <span>XP reward</span>
                    <strong>{{ report.ExpeditionRewards.xp }}</strong>
                </div>
                <button class="primary-button" v-on:click="onclick()">Return home</button>
            </div>
        </section>

        <div class="report-timeline">
            <article v-for="(eventReport, eventIndex) in report.whatHappened" :key="eventIndex" class="panel report-section">
                <div class="panel-heading">
                    <div>
                        <p class="eyebrow">Encounter phase</p>
                        <h3>{{ getEncounterStateString(eventReport.stepState) }}</h3>
                    </div>
                </div>

                <div class="report-events">
                    <div v-for="(entry, entryIndex) in eventReport.timeline" :key="entryIndex" class="report-event">
                        <div class="report-event-time">{{ new Date(entry.when).toLocaleString() }}</div>

                        <div class="report-event-copy">
                            <strong>{{ entry.what }}</strong>

                            <div v-if="entry.whatAction">
                                <p v-for="(txt, txtIndex) in getStringFromFAD(entry.whatAction)" :key="txtIndex">{{ txt }}</p>
                            </div>
                        </div>
                    </div>
                </div>
            </article>
        </div>
    </section>

    <section v-else class="panel report-summary">
        <p class="eyebrow">Mission log</p>
        <h2>Loading report...</h2>
    </section>
</template>
