<script setup lang="ts">
    import { inject, onMounted, ref, watch } from "vue";
    import { EncounterState, type CurrentStepRequestMessage, type ExpeditionStepResolveInfo, type User } from "../tools/types.ts"
    import { formatTextTimeFromTimeMS, getCurrentExpeditionStepResolveInfo, postRequest, RequestType } from "../tools/utils.ts"
    import Team from "./Team.vue";
import type { NavigationHandler } from "@/tools/navigationHandler.ts";

            const navigationHandler = inject<NavigationHandler>('navigationHandler')!
    const user = navigationHandler.getUser()
</script>

<template>
    <div class="row">
        <Team :team="user?.currentTeam"/>
        <div>
            <h1>Traveling</h1>
        </div>
    </div>
</template>
