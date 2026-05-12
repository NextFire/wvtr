<script setup lang="ts">
    import { inject } from 'vue'
    import { EncounterState } from '../tools/types.ts'
    import Home from './Home.vue'
    import Expedition from './Expedition.vue'
    import { NavigationHandler } from '@/tools/navigationHandler.ts'
    import Report from './Report.vue'

    const navigationHandler = inject<NavigationHandler>('navigationHandler')!
</script>

<template>
    <main v-if="navigationHandler.getUserState() != EncounterState.Error" class="body body-shell">
        <Home v-if="navigationHandler.getUserState() == EncounterState.Home"/>
        <Report v-else-if="navigationHandler.getUserState() == EncounterState.Report"/>
        <Expedition v-else/>
    </main>
    <main v-else class="body body-shell">
        <section class="panel panel-danger state-panel">
            <p class="eyebrow">Signal lost</p>
            <h2>There is a problem with the current user state.</h2>
            <p>Refresh the session or reconnect the player before launching the next adventure.</p>
        </section>
    </main>
</template>
