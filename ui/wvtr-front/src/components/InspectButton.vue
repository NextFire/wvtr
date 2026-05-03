<script setup lang="ts">
    import { inject } from "vue";
    import type { Hero } from "../tools/types.ts"
    import { global } from "../tools/utils.ts"
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
    <button class="inspection-button" v-on:click="onClickInspectButton()">
        <img :src="global.DOMAIN_NAME + '/imgs/magnifyingglassicon.png'" width="25px"/>
    </button>
</template>