<script setup lang="ts">
    import { inject, ref, watch } from 'vue'
    import type { Team } from "../tools/types.ts"
    import Hero from './Hero.vue'
    import { global } from "../tools/utils.ts"
import type { NavigationHandler } from '@/tools/navigationHandler.ts'

    const props = defineProps<{
        team: Team | undefined;
    }>();
    
    // const navigationHandler = inject<NavigationHandler>('navigationHandler')!

    // const user =  navigationHandler.getUser()

</script>

<template>
    <div class="team">
        <Hero v-if="props.team" :hero="props.team.heroes[0]"/>
        <Hero v-else :hero="undefined"/>
        <Hero v-if="props.team" :hero="props.team.heroes[1]"/>
        <Hero v-else :hero="undefined"/>
        <Hero v-if="props.team" :hero="props.team.heroes[2]"/>
        <Hero v-else :hero="undefined"/>
    </div>
</template>
