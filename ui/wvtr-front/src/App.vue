<script setup lang="ts">
    import { ref, onMounted, inject } from 'vue'
    import Header from "./components/Header.vue"
    import Body from "./components/Body.vue"
    import { NavigationHandler } from './tools/navigationHandler.ts'
    import type { User } from "./tools/types.ts"
    import type { VueCookies } from 'vue-cookies'

    let $cookies = inject<VueCookies>('$cookies');
    

    //const user = ref<User|undefined>(undefined)
    const navigationHandler = inject<NavigationHandler>('navigationHandler')!
    let user = navigationHandler.getUser()
    onMounted(async () => {
        user = await navigationHandler.setup($cookies!)
    })



</script>

<template>
<div v-if="!user" class="page">
    <!-- <a v-if="authUrl" :href="authUrl">Login with OIDC</a> -->
    <p>loading auth...</p>
    <!-- <p v-else>loading auth...</p> -->
</div>
<div v-if="user" class="page">
    <Header/>
    <Body/>
</div>
</template>

