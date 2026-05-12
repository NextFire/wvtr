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
                <h3>Road team</h3>
            </div>
            <Team :team="user?.currentTeam"/>
        </div>

        <div class="panel panel-feature encounter-story">
            <p class="eyebrow">Travel event</p>
            <h3>Rolling toward the next surprise</h3>
            <p>The squad is moving between scenes. This quiet stretch is a nice moment to check HP bars before the next incident snaps into view.</p>
        </div>
    </section>
</template>
