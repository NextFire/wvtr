<script setup lang="ts">
    import { onMounted, ref, watch } from "vue";
    import { EncounterState, type CurrentStepRequestMessage, type ExpeditionStepResolveInfo, type User } from "../tools/types.ts"
    import { formatTextTimeFromTimeMS, getCurrentExpeditionStepResolveInfo, postRequest, RequestType } from "../tools/utils.ts"
    import Team from "./Team.vue";
</script>

<template>
    <div>
        <h1>Traveling</h1>
    </div>
</template>
