<script setup lang="ts">
    import type { Team } from '../tools/types.ts'
    import Hero from './Hero.vue'

    const props = defineProps<{
        team: Team | undefined;
    }>()
</script>

<template>
    <div class="team">
        <Hero v-for="slot in 3" :key="slot" :hero="props.team?.heroes[slot - 1]"/>
    </div>
</template>
