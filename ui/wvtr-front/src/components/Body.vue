<script setup lang="ts">
    import type { User } from "../tools/types.ts"
    import { EncounterState } from "../tools/types.ts"
    import Home from "./Home.vue"
    import Expedition from "./Expedition.vue"
import { NavigationHandler } from "@/tools/navigationHandler.ts"
import { inject } from "vue"
import Report from "./Report.vue"

    const navigationHandler = inject<NavigationHandler>('navigationHandler')!
    console.log(navigationHandler.getUserState())
</script>

<template>
    <div v-if="navigationHandler.getUserState() != EncounterState.Error" class="body">
        <Home v-if="navigationHandler.getUserState() == EncounterState.Home"/>
        <Report v-else-if="navigationHandler.getUserState() == EncounterState.Report"/>
        <Expedition v-else/>
    </div>
    <div v-else class="body">
        there is a problem with user...
    </div>
</template>