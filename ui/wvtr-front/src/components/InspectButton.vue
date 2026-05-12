<script setup lang="ts">
    import { inject } from "vue";
    import type { Hero } from "../tools/types.ts"
    import { NavigationStatus, NavigationHandler } from "../tools/navigationHandler.ts";

    const props = defineProps<{
        hero: Hero;
    }>();

    const navigationHandler = inject<NavigationHandler>('navigationHandler')!
    function onClickInspectButton() {
        navigationHandler.setHomeStatus(NavigationStatus.InspectHero)
        navigationHandler.setHeroToInspect(props.hero)
    }
</script>

<template>
    <button class="inspection-button" aria-label="Inspect hero" v-on:click.stop="onClickInspectButton()">
        <svg class="inspection-icon" viewBox="0 0 24 24" aria-hidden="true">
            <circle cx="11" cy="11" r="6.5" fill="none" stroke="currentColor" stroke-width="2.2"/>
            <path d="M16 16l4.5 4.5" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round"/>
        </svg>
    </button>
</template>
