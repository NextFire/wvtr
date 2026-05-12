<script setup lang="ts">
    import { inject } from 'vue'
    import Team from './Team.vue'
    import type { NavigationHandler } from '@/tools/navigationHandler.ts'

    const navigationHandler = inject<NavigationHandler>('navigationHandler')!
    const user = navigationHandler.getUser()
</script>

<template>
    <section class="encounter-layout encounter-layout-single">
        <div class="panel encounter-team-panel">
            <div>
                <p class="eyebrow">Party status</p>
                <h3>Event-ready squad</h3>
            </div>
            <Team :team="user?.currentTeam"/>
        </div>

        <div class="panel panel-feature encounter-story">
            <p class="eyebrow">Neutral event</p>
            <h3>A pause between battles</h3>
            <p>Something odd, lucky, or mildly suspicious is happening. Keep the roster in view while the story moment resolves.</p>
        </div>
    </section>
</template>
