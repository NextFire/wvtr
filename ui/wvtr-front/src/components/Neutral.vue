<script setup lang="ts">
    import { inject } from "vue";
import type { User } from "../tools/types.ts"
    import { global } from "../tools/utils.ts"
import Team from "./Team.vue";
import type { NavigationHandler } from "@/tools/navigationHandler.ts";

        const navigationHandler = inject<NavigationHandler>('navigationHandler')!
    const user = navigationHandler.getUser()

</script>

<template>
    <div class="row">
        <Team :team="user?.currentTeam"/>
        <div>
            <h1>Neutral</h1>
        </div>
    </div>
</template>