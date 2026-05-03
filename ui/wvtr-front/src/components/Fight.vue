<script setup lang="ts">
    import { inject } from "vue";
import type { Team, User } from "../tools/types.ts"
    import { global } from "../tools/utils.ts"
import type { NavigationHandler } from "@/tools/navigationHandler.ts";
import TeamP from "./Team.vue";
    
    const props = defineProps<{
        eteam: Team | undefined;
    }>();

    const navigationHandler = inject<NavigationHandler>('navigationHandler')!
    const user = navigationHandler.getUser()
    // const currentExpStep = navigationHandler.getCurrentExpeditionStepResolveInfo();
</script>

<template>
    <div class="row">
        <div>
            <TeamP :team="user?.currentTeam"/>
        </div>
        <div>
            <h1>Fight</h1>
            <div>This display is experimental for now and might not reflect the exacte state of the user team.</div>
        </div>
        <div>
            <TeamP :team="eteam"/>
        </div>
    </div>

</template>