<script setup lang="ts">
    import { inject } from 'vue'
    import type { Team } from '../tools/types.ts'
    import type { NavigationHandler } from '@/tools/navigationHandler.ts'
    import TeamP from './Team.vue'

    defineProps<{
        eteam: Team | undefined;
    }>()

    const navigationHandler = inject<NavigationHandler>('navigationHandler')!
    const user = navigationHandler.getUser()
</script>

<template>
    <section class="encounter-layout encounter-layout-duel">
        <div class="panel encounter-team-panel">
            <div>
                <p class="eyebrow">Allies</p>
                <h3>Your front line</h3>
            </div>
            <TeamP :team="user?.currentTeam"/>
        </div>

        <div class="panel panel-feature encounter-story encounter-fight">
            <p class="eyebrow">Battlefield</p>
            <h3>Fight in progress</h3>
            <p>This live view favors clarity over simulation detail, so use it as a stylish combat snapshot while the timeline updates the real state underneath.</p>
        </div>

        <div class="panel encounter-team-panel opponent-panel">
            <div>
                <p class="eyebrow">Opponents</p>
                <h3>Enemy team</h3>
            </div>
            <TeamP :team="eteam"/>
        </div>
    </section>
</template>
