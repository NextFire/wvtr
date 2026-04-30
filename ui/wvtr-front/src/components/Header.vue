<script setup lang="ts">
    import { inject } from "vue";
import type { NavigationHandler } from "@/tools/navigationHandler.ts";

    const navigationHandler = inject<NavigationHandler>('navigationHandler')!
    const user = navigationHandler.getUser() 

</script>

<template>
    <div class="header">
        <div v-if="user">
            <h1>{{ user.name }}</h1>
        </div>
        <div v-else>
            <h1>Chargement ...</h1>
        </div>
    </div>
</template>