<script setup lang="ts">
    import { inject, onMounted, ref, watch } from "vue"
    import type { Hero, User } from "../tools/types.ts"
    import Team from "./Team.vue"
    import TeamManagement from "./TeamManagement.vue"
    import ExpeditionsList from "./ExpeditionsList.vue"
    import Waifus from "./Waifus.vue"
    import InspectHero from "./InspectHero.vue"
    import { NavigationStatus, NavigationHandler } from "@/tools/navigationHandler.ts"

        const navigationHandler = inject<NavigationHandler>('navigationHandler')!
    const user = navigationHandler.getUser()

    
    function setHomeStatus (newStatus: NavigationStatus) {
        navigationHandler!.setHomeStatus(newStatus)
    }
</script>

<template>
    <!-- <div class="home"> -->
        <Team :team="user?.currentTeam"/>
        <div class="column">
            <button v-on:click="setHomeStatus(NavigationStatus.TeamManagement)">
            manage Team
            </button>
            <button v-on:click="setHomeStatus(NavigationStatus.ExpeditionManagement)">
            launch expedition
            </button>
            <button v-on:click="setHomeStatus(NavigationStatus.HeroMaker)">
            Check available waifus
            </button>
        </div>
        <TeamManagement v-if="navigationHandler.getNavigationStatus().value == NavigationStatus.TeamManagement"/>
        
        <ExpeditionsList v-else-if="navigationHandler.getNavigationStatus().value == NavigationStatus.ExpeditionManagement"/>
        <Waifus v-else-if="navigationHandler.getNavigationStatus().value == NavigationStatus.HeroMaker"/>
        <InspectHero v-else-if="navigationHandler.getNavigationStatus().value == NavigationStatus.InspectHero"/>
    <!-- </div> -->
</template>


